package directive

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/hashicorp/go-multierror"
	"github.com/rotisserie/eris"
	"github.com/solo-io/skv2/pkg/api/discovery.multicluster.solo.io/v1alpha1"
	"github.com/solo-io/skv2/pkg/multicluster/discovery/cloud/clients"
	"github.com/solo-io/skv2/pkg/utils"
	"k8s.io/apimachinery/pkg/labels"
)

//go:generate mockgen -source ./aws.go -destination ./mocks/mock_aws.go

// Resolve the set of AWS resources selected by provided list of AwsDiscoveryDirectiveSpec.
type AwsDiscoveryResolver interface {
	Resolve(
		ctx context.Context,
		creds *credentials.Credentials,
		awsDirectives []*v1alpha1.AwsDiscoveryDirectiveSpec,
	) AwsResources
}

var UnknownSelectorType = func(selector *v1alpha1.AwsDiscoveryDirectiveSpec_ResourceSelector) error {
	return eris.Errorf("Unknown SettingsSpec_AwsAccount_ResourceSelector type: %+v", selector)
}

type AwsResources struct {
	EksArns []string
}

type awsDiscoveryResolver struct {
}

func (a *awsDiscoveryResolver) Resolve(
	ctx context.Context,
	creds *credentials.Credentials,
	awsDirectives []*v1alpha1.AwsDiscoveryDirectiveSpec,
) (AwsResources, error) {
	var multierr *multierror.Error

	eksArns, err := resolveEks(ctx, creds, awsDirectives)
	if err != nil {
		multierr = multierror.Append(multierr, err)
	}

	return AwsResources{
		EksArns: eksArns,
	}, multierr.ErrorOrNil()
}

func resolveEks(
	ctx context.Context,
	creds *credentials.Credentials,
	awsDirectives []*v1alpha1.AwsDiscoveryDirectiveSpec,
) ([]string, error) {
	var eksArns []string

	selectors := []*v1alpha1.AwsDiscoveryDirectiveSpec_ResourceSelector{}
	for _, directive := range awsDirectives {
		if directive.EksSelectors != nil {
			selectors = append(selectors, directive.EksSelectors...)
		}
	}
	selectorsByRegion, err := groupSelectorsByRegion(selectors)
	if err != nil {
		return nil, err
	}

	for region, selectors := range selectorsByRegion {
		eksClient, err := clients.NewEksClient(region, creds)
		if err != nil {
			return nil, err
		}

		var eksNames []string
		eksClient.ListClusters(ctx, func(output *eks.ListClustersOutput) {
			for _, clusterName := range output.Clusters {
				eksNames = append(eksNames, aws.StringValue(clusterName))
			}
		})

		for _, eksName := range eksNames {
			eks, err := eksClient.DescribeCluster(ctx, eksName)
			if err != nil {
				return nil, err
			}
			arnString := aws.StringValue(eks.Arn)
			matched, err := resourceMatchedBySelectors(arnString, aws.StringValueMap(eks.Tags), selectors)
			if err != nil {
				return nil, err
			}
			if matched {
				eksArns = append(eksArns, arnString)
			}
		}
	}
	return eksArns, nil
}

// Return true if AWS resource is selected by any ResourceSelector.
func resourceMatchedBySelectors(
	arnString string,
	tags map[string]string,
	selectors []*v1alpha1.AwsDiscoveryDirectiveSpec_ResourceSelector,
) (bool, error) {
	parsedArn, err := arn.Parse(arnString)
	if err != nil {
		return false, err
	}
	for _, selector := range selectors {
		switch selector.GetSelectorType().(type) {
		case *v1alpha1.AwsDiscoveryDirectiveSpec_ResourceSelector_Matcher_:
			accountIds := selector.GetMatcher().AccountIds
			selectorTags := selector.GetMatcher().Tags
			// Empty accountIds allows matching on any account ID.
			if accountIds != nil && !utils.ContainsString(accountIds, parsedArn.AccountID) {
				continue
			}
			// Empty tags allows matching on any tags.
			if selectorTags != nil && !labels.AreLabelsInWhiteList(tags, selectorTags) {
				continue
			}
			return true, nil
		case *v1alpha1.AwsDiscoveryDirectiveSpec_ResourceSelector_Arn:
			if selector.GetArn() == arnString {
				return true, nil
			}
		default:
			return false, UnknownSelectorType(selector)
		}
	}
	return false, nil
}

type awsSelectorsByRegion map[string][]*v1alpha1.AwsDiscoveryDirectiveSpec_ResourceSelector

// Group AwsDiscoveryDirectives by region. This is needed because AWS sessions must be scoped to a particular AWS region.
// Reference: https://github.com/aws/aws-sdk-go/blob/f324f9f20f565f497063770e710805717c05ebaa/aws/config.go#L70
func groupSelectorsByRegion(
	awsSelectors []*v1alpha1.AwsDiscoveryDirectiveSpec_ResourceSelector,
) (awsSelectorsByRegion, error) {
	selectorsByRegion := awsSelectorsByRegion{}
	for _, selector := range awsSelectors {
		switch selector.GetSelectorType().(type) {
		case *v1alpha1.AwsDiscoveryDirectiveSpec_ResourceSelector_Matcher_:
			// If Regions is empty, select across all regions.
			if selector.GetMatcher().Regions == nil {
				for region, _ := range endpoints.AwsPartition().Regions() {
					addSelectorToRegion(selectorsByRegion, region, selector)
				}
				continue
			}
			for _, region := range selector.GetMatcher().GetRegions() {
				addSelectorToRegion(selectorsByRegion, region, selector)
			}
		case *v1alpha1.AwsDiscoveryDirectiveSpec_ResourceSelector_Arn:
			parsedArn, err := arn.Parse(selector.GetArn())
			if err != nil {
				return nil, err
			}
			addSelectorToRegion(selectorsByRegion, parsedArn.Region, selector)
		default:
			return nil, UnknownSelectorType(selector)
		}
	}
	return selectorsByRegion, nil
}

func addSelectorToRegion(
	selectorsByRegion awsSelectorsByRegion,
	region string,
	selector *v1alpha1.AwsDiscoveryDirectiveSpec_ResourceSelector,
) {
	regionalSelectors, ok := selectorsByRegion[region]
	if !ok {
		regionalSelectors = []*v1alpha1.AwsDiscoveryDirectiveSpec_ResourceSelector{}
		selectorsByRegion[region] = regionalSelectors
	}
	regionalSelectors = append(regionalSelectors, selector)
}
