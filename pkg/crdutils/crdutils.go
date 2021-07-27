package crdutils

import (
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/rotisserie/eris"
	apiextv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

// Chart provides the input data structure for generating Helm charts from the skv2 chart "meta-templates" (templates whose outputs are templates and other files used by generated Helm charts)
const (
	CRDVersionKey  = "crd.solo.io/version"
	CRDSpecHashKey = "crd.solo.io/specHash"
	CRDMetadataKey = "crd.solo.io/crdMetadata"
)

type CRDMetadata struct {
	CRDS    []CRDAnnotations `json:"crds"`
	Version string           `json:"version"`
}

type CRDAnnotations struct {
	Name string `json:"name"`
	Hash string `json:"hash"`
}

type ErrMap map[string]error
type CrdNeedsUpgrade struct {
	CRDName string
}

func (e *CrdNeedsUpgrade) Error() string {
	return fmt.Sprintf("CRD %s needs to be upgraded", e.CRDName)
}

type CrdNotFound struct {
	CRDName string
}

func (e *CrdNotFound) Error() string {
	return fmt.Sprintf("CRD %s not found. Is it unused?", e.CRDName)
}

func DoCrdsNeedUpgrade(newProdCrdInfo CRDMetadata, deployedInClusterCrds []apiextv1beta1.CustomResourceDefinition) ErrMap {
	crdMap := make(map[string]string)
	for _, crd := range newProdCrdInfo.CRDS {
		crdMap[crd.Name] = crd.Hash
	}
	ret := ErrMap{}
	for _, ourCrd := range deployedInClusterCrds {
		if hash, ok := crdMap[ourCrd.Name]; !ok {
			ret[ourCrd.Name] = &CrdNotFound{CRDName: ourCrd.Name}
		} else {
			needUpgrade, err := DoesCrdNeedUpgrade(newProdCrdInfo.Version, hash, ourCrd.Annotations)

			if err != nil {
				ret[ourCrd.Name] = err
			} else if needUpgrade {
				ret[ourCrd.Name] = &CrdNeedsUpgrade{CRDName: ourCrd.Name}
			}
		}
	}
	if len(ret) == 0 {
		return nil
	}
	return ret
}

/**
 * The idea behind this function is that we only want to ugprade a CRD if the version of the new CRD
 * is higher than version of the CRD in the cluser **and** the specHash is different.
 * The reasoning:
 * if the spec hash is the same, then the CRD's json schema is the same, and it is pointless to upgrade it.
 * if the spec hash is different, then we want to make sure we only update a CRD when it is backwards compatible.
 * under the assumption that we don't make backwards incompatible changes, then this means that we
 * should only upgrade a CRD if it's version is higher than what's deployed on the cluster.
 */
func DoesCrdNeedUpgrade(newProductVersion, newCrdHash string, deployedCrdAnnotations map[string]string) (bool, error) {

	if newProductVersion == "" || newCrdHash == "" {
		return false, eris.New(fmt.Sprintf("Cannot determine if CRDs need an upgrade, missing internal data: %s %s", newProductVersion, newCrdHash))
	}

	crdVersion, ok := deployedCrdAnnotations[CRDVersionKey]
	if !ok {
		return false, eris.New(fmt.Sprintf("Cannot determine crd product version from CRD annotations: %v", deployedCrdAnnotations))
	}
	crdSpecHash, ok := deployedCrdAnnotations[CRDSpecHashKey]
	if !ok {
		return false, eris.New(fmt.Sprintf("Cannot determine crd spec hash from CRD annotations: %v", deployedCrdAnnotations))
	}

	if newCrdHash == crdSpecHash {
		return false, nil
	}

	newProductVersionSemver, err := semver.NewVersion(newProductVersion)
	if err != nil {
		return false, eris.Wrapf(err, "Cannot parse new product version: %s", newProductVersion)
	}

	// parse semver of the current product version
	currentCrdVersionSemver, err := semver.NewVersion(crdVersion)
	if err != nil {
		return false, eris.Wrapf(err, "Cannot parse current crd version: %s", crdVersion)
	}

	// If the new product version is greater than the current crd version, the CRD needs to be upgraded.
	if currentCrdVersionSemver.Compare(newProductVersionSemver) <= 0 {
		return true, nil
	}
	return false, nil
}
