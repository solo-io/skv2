package parse

import (
	"fmt"
	"io/fs"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/solo-io/skv2/pkg/crdutils"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
)

func ParseCRDs(f fs.FS) (crdutils.CRDMetadata, error) {
	var ret crdutils.CRDMetadata
	// read all crds from the specified directory
	// generate a file with all the versions and hashes of the crds
	// package the generated file into the binary
	dirList, err := fs.ReadDir(f, ".")
	if err != nil {
		return ret, err
	}
	var errs error

	for _, dir := range dirList {
		if !strings.HasSuffix(dir.Name(), ".yaml") {
			continue
		}

		crdInfo, ver, err := getCRDInfo(f, dir.Name())
		if err == nil {
			ret.CRDS = append(ret.CRDS, crdInfo)
			if ret.Version == "" {
				ret.Version = ver
			} else if ret.Version != ver {
				// the version should be the same for all CRDs
				errs = multierror.Append(errs, fmt.Errorf("CRDs version mismatch: %s %s != %s", crdInfo.Name, ret.Version, ver))
			}

		} else {
			errs = multierror.Append(err, errs)
		}
	}
	return ret, errs
}

func getCRDInfo(f fs.FS, file string) (crdutils.CRDAnnotations, string, error) {
	var ret crdutils.CRDAnnotations
	crdRaw, err := fs.ReadFile(f, file)
	if err != nil {
		return ret, "", err
	}

	decoder := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)

	crd := &unstructured.Unstructured{}
	_, _, err = decoder.Decode([]byte(crdRaw), nil, crd)
	if err != nil {
		return ret, "", err
	}

	ret.Name = crd.GetName()
	var v string
	if m := crd.GetAnnotations(); m != nil {
		ret.Hash = m[crdutils.CRDSpecHashKey]
		v = m[crdutils.CRDVersionKey]
	}
	return ret, v, nil
}
