package stats

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"gopkg.in/yaml.v2"
)

// Could be good to enhance this interface to use more defined types/structs. Right now they are mostly generics which relies on a lot of type assertion.
// If we did maintain the format here it would need to be checked when set
/*
INPUT FORMAT
{
    "cert-issuer": {
        "<resourceType>": [...],
        "<resourceType>": [...]
    },
    "translator": {
        "<resourceType>": [resources .....],
	}
}

OUTPUT FORMAT
{
    "roottrust": {...},
    "translator": {
        "<cluster>/<namespace>": {
            "<name>~<namespace>~<cluster>~<resourceType>": {
                "Err": null,
                "Outputs": {
                    "<resourceType>": [resources .....]
                },
                "Warnings": null
            },
		}
	}
}
*/
type SnapshotHistory interface {
	// SetInput Sets the input snapshot for the given component
	SetInput(id string, latestInput json.Marshaler)
	// GetInputCopy gets an in-memory copy of the output snapshot for all components.
	GetInputCopy() (map[string]interface{}, error)
	// GetInput gets the input snapshot for all components.
	GetInput() ([]byte, error)
	// GetFilteredInput gets the input snapshot for all components and applies filters
	GetFilteredInput(format string, filters Filters) ([]byte, error)
	// SetOutput Sets the output snapshot for the given component
	SetOutput(id string, latestOutput json.Marshaler)
	// GetOutputCopy gets an in-memory copy of the output snapshot for all components.
	GetOutputCopy() (map[string]interface{}, error)
	// GetOutput gets the output snapshot for all component.
	GetOutput() ([]byte, error)
	// GetFilteredOutput gets the output snapshot for all components and applies filters
	GetFilteredOutput(format string, filters Filters) ([]byte, error)
}
type snapshotHistory struct {
	lock         sync.RWMutex
	latestInput  map[string]json.Marshaler
	latestOutput map[string]json.Marshaler
}

func NewSnapshotHistory() SnapshotHistory {
	return &snapshotHistory{
		latestInput:  map[string]json.Marshaler{},
		latestOutput: map[string]json.Marshaler{},
	}
}

func (h *snapshotHistory) SetInput(id string, latestInput json.Marshaler) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.latestInput[id] = latestInput
}

func (h *snapshotHistory) GetInputCopy() (map[string]interface{}, error) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	if h.latestInput == nil {
		return map[string]interface{}{}, nil
	}
	genericMaps, err := getGenericMaps(h.latestInput)
	if err != nil {
		return nil, err
	}
	return genericMaps, nil
}

func (h *snapshotHistory) GetInput() ([]byte, error) {
	return h.GetFilteredInput("json_compact", NewFilters(nil, nil, nil))
}

func (h *snapshotHistory) GetFilteredInput(format string, filters Filters) ([]byte, error) {
	input, err := h.GetInputCopy()
	if err != nil {
		return nil, err
	}

	// short circuit if no filters
	if !filters.clusters.Exists() && !filters.namespaces.Exists() && !filters.resourceTypes.Exists() {
		return formatMap(format, input)
	}

	translator := input["translator"].(map[string]interface{})
	// Filter resource types
	if filters.resourceTypes.Exists() {
		for resourceType := range translator {
			if !filters.resourceTypes.Contains(resourceType) {
				delete(translator, resourceType)
			}
		}
	}

	// Filter cluster and namespaces
	if filters.clusters.Exists() || filters.namespaces.Exists() {
		for k, resources := range translator {
			newResources := []interface{}{}
			for _, resource := range resources.([]interface{}) {
				resource := resource.(map[string]interface{})
				metadata := resource["metadata"].(map[string]interface{})
				cluster, ok := metadata["clusterName"].(string)
				include := true
				if filters.clusters.Exists() && ok && !filters.clusters.Contains(cluster) {
					include = false
				}
				namespace, ok := metadata["namespace"].(string)
				if k == "/v1, Kind=Namespace" {
					namespace, ok = metadata["name"].(string)
				}
				if filters.namespaces.Exists() && ok && !filters.namespaces.Contains(namespace) {
					include = false
				}
				if include {
					newResources = append(newResources, resource)
				}
			}
			if len(newResources) == 0 {
				delete(translator, k)
			} else {
				translator[k] = newResources
			}
		}
	}

	return formatMap(format, translator)
}

func (h *snapshotHistory) SetOutput(id string, latestOutput json.Marshaler) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.latestOutput[id] = latestOutput
}

func (h *snapshotHistory) GetOutputCopy() (map[string]interface{}, error) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	if h.latestOutput == nil {
		return map[string]interface{}{}, nil
	}
	genericMaps, err := getGenericMaps(h.latestOutput)
	if err != nil {
		return nil, err
	}
	return genericMaps, nil
}

func (h *snapshotHistory) GetOutput() ([]byte, error) {
	return h.GetFilteredOutput("json_compact", NewFilters(nil, nil, nil))
}

func (h *snapshotHistory) GetFilteredOutput(format string, filters Filters) ([]byte, error) {
	output, err := h.GetOutputCopy()
	if err != nil {
		return nil, err
	}

	// short circuit if no filters
	if !filters.clusters.Exists() && !filters.namespaces.Exists() && !filters.resourceTypes.Exists() {
		return formatMap(format, output)
	}

	translator := output["translator"].(map[string]interface{})
	// Filter resource types
	if filters.clusters.Exists() || filters.namespaces.Exists() {
		for k := range translator {
			splitKey := strings.Split(k, "/")
			cluster := splitKey[0]
			namespace := splitKey[1]
			include := true
			if filters.clusters.Exists() && !filters.clusters.Contains(cluster) {
				include = false
			}
			if filters.namespaces.Exists() && !filters.namespaces.Contains(namespace) {
				include = false
			}
			if !include {
				delete(translator, k)
			}
		}
	}

	// Filter cluster and namespaces
	if filters.resourceTypes.Exists() {
		for cluster_ns, resources := range translator {
			for k := range resources.(map[string]interface{}) {
				splitKey := strings.Split(k, "~")
				resourceType := splitKey[len(splitKey)-1]
				if !filters.resourceTypes.Contains(resourceType) {
					delete(translator[cluster_ns].(map[string]interface{}), k)
				}
			}
			if len(translator[cluster_ns].(map[string]interface{})) == 0 {
				delete(translator, cluster_ns)
			}
		}
	}

	return formatMap(format, translator)

}

type filter interface {
	// Check if filter constains this value
	Contains(s string) bool
	// Check if filter exists
	Exists() bool
}

type filterMap map[string]bool

func newFilter(selectedFilters []string) filter {
	var fMap filterMap
	fMap = make(map[string]bool)
	for _, f := range selectedFilters {
		fMap[f] = true
	}
	return fMap
}

type Filters struct {
	clusters      filter
	namespaces    filter
	resourceTypes filter
}

func NewFilters(includedClusters []string, includedNamespaces []string, includedResourceTypes []string) Filters {
	return Filters{
		clusters:      newFilter(includedClusters),
		namespaces:    newFilter(includedNamespaces),
		resourceTypes: newFilter(includedResourceTypes),
	}
}

func (f filterMap) Contains(s string) bool {
	if _, ok := f[s]; ok {
		return true
	} else {
		return false
	}
}

func (f filterMap) Exists() bool {
	if len(f) != 0 {
		return true
	} else {
		return false
	}
}

func AddSnapshots(mux *http.ServeMux, history SnapshotHistory) {

	mux.HandleFunc(
		"/snapshots/input", func(w http.ResponseWriter, r *http.Request) {

			b, err := history.GetFilteredInput(getUrlParams(r))
			// Gives download correct file extension
			w.Header().Set("Content-Type", getContentType(r))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			_, _ = fmt.Fprint(w, string(b))
		},
	)

	mux.HandleFunc(
		"/snapshots/output", func(w http.ResponseWriter, r *http.Request) {
			b, err := history.GetFilteredOutput(getUrlParams(r))
			// Gives download correct file extension
			w.Header().Set("Content-Type", getContentType(r))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			_, _ = fmt.Fprint(w, string(b))
		},
	)
}

func getUrlParams(r *http.Request) (string, Filters) {
	clustersS := r.FormValue("clusters")
	namespacesS := r.FormValue("namespaces")
	resourceTypesS := r.FormValue("resourceTypes")
	var clusters []string
	var namespaces []string
	var resourceTypes []string
	if clustersS != "" {
		clusters = strings.Split(clustersS, "::")
	}
	if namespacesS != "" {
		namespaces = strings.Split(namespacesS, "::")
	}
	if resourceTypesS != "" {
		resourceTypes = strings.Split(resourceTypesS, "::")
	}
	return r.FormValue("format"), NewFilters(clusters, namespaces, resourceTypes)
}

func getGenericMaps(snapshot map[string]json.Marshaler) (map[string]interface{}, error) {
	genericMaps := map[string]interface{}{}
	for id, obj := range snapshot {
		jsn, err := obj.MarshalJSON()
		if err != nil {
			return nil, err
		}
		genericMap := map[string]interface{}{}
		if err := json.Unmarshal(jsn, &genericMap); err != nil {
			return nil, err
		}
		genericMaps[id] = genericMap
	}
	return genericMaps, nil
}

func formatMap(format string, genericMaps map[string]interface{}) ([]byte, error) {
	switch format {
	case "json":
		return json.MarshalIndent(genericMaps, "", "    ")
	case "", "json_compact":
		return json.Marshal(genericMaps)
	case "yaml":
		return yaml.Marshal(genericMaps)
	default:
		return nil, fmt.Errorf("Invalid format of %s", format)
	}

}

func getContentType(r *http.Request) string {
	switch r.FormValue("format") {
	case "", "json", "json_compact":
		return "application/json"
	case "yaml":
		return "text/x-yaml"
	default:
		return "application/json"
	}
}
