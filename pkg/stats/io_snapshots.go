package stats

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"gopkg.in/yaml.v2"
)

type SnapshotHistory interface {
	// SetInput Sets the input snapshot for the given component
	SetInput(id string, latestInput json.Marshaler)
	// GetInput gets the input snapshot for all components
	GetInput(format string, clusters []string, namespaces []string, resourceTypes []string) ([]byte, error)
	// GetMapInput gets the input snapshot for all components as in in-memory map
	GetMapInput() (map[string]interface{}, error)
	// SetOutput Sets the output snapshot for the given component
	SetOutput(id string, latestOutput json.Marshaler)
	// GetOutput gets the output snapshot for all component
	GetOutput(format string, clusters []string, namespaces []string, resourceTypes []string) ([]byte, error)
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

func (h *snapshotHistory) GetInput(format string, clusters []string, namespaces []string, resourceTypes []string) ([]byte, error) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	if h.latestInput == nil {
		return []byte{}, nil
	}
	genericMaps, err := getGenericMaps(h.latestInput)
	if err != nil {
		return nil, err
	}

	translator := genericMaps["translator"].(map[string]interface{})
	// Filter resource types
	if resourceTypes != nil {
		for k := range translator {
			if !contains(resourceTypes, k) {
				delete(translator, k)
			}
		}
	}

	// Filter cluster and namespaces
	if clusters != nil || namespaces != nil {
		for k, resources := range translator {
			newResources := []interface{}{}
			for _, resource := range resources.([]interface{}) {
				resource := resource.(map[string]interface{})
				metadata := resource["metadata"].(map[string]interface{})
				cluster, ok := metadata["clusterName"].(string)
				include := true
				if clusters != nil && ok && !contains(clusters, cluster) {
					include = false
				}
				namespace, ok := metadata["namespace"].(string)
				if k == "/v1, Kind=Namespace" {
					namespace, ok = metadata["name"].(string)
				}
				if namespaces != nil && ok && !contains(namespaces, namespace) {
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

	// If filters enabled only return translator section
	if clusters != nil || namespaces != nil || resourceTypes != nil {
		return formatMap(format, translator)
	} else {
		return formatMap(format, genericMaps)
	}

}

func (h *snapshotHistory) GetMapInput() (map[string]interface{}, error) {
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

func (h *snapshotHistory) SetOutput(id string, latestOutput json.Marshaler) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.latestOutput[id] = latestOutput
}

func (h *snapshotHistory) GetOutput(format string, clusters []string, namespaces []string, resourceTypes []string) ([]byte, error) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	if h.latestOutput == nil {
		return []byte{}, nil
	}
	genericMaps, err := getGenericMaps(h.latestOutput)
	if err != nil {
		return nil, err
	}

	translator := genericMaps["translator"].(map[string]interface{})
	// Filter resource types
	if clusters != nil || namespaces != nil {
		for k := range translator {
			splitKey := strings.Split(k, "/")
			cluster := splitKey[0]
			namespace := splitKey[1]
			include := true
			if clusters != nil && !contains(clusters, cluster) {
				include = false
			}
			if namespaces != nil && !contains(namespaces, namespace) {
				include = false
			}
			if !include {
				delete(translator, k)
			}
		}
	}

	// Filter cluster and namespaces
	if resourceTypes != nil {
		for cluster_ns, resources := range translator {
			for k := range resources.(map[string]interface{}) {
				splitKey := strings.Split(k, "~")
				resourceType := splitKey[len(splitKey)-1]
				if !contains(resourceTypes, resourceType) {
					delete(translator[cluster_ns].(map[string]interface{}), k)
				}
			}
			if len(translator[cluster_ns].(map[string]interface{})) == 0 {
				delete(translator, cluster_ns)
			}
		}
	}

	// If filters enabled only return translator section
	if clusters != nil || namespaces != nil || resourceTypes != nil {
		return formatMap(format, translator)
	} else {
		return formatMap(format, genericMaps)
	}
}

func AddSnapshots(mux *http.ServeMux, history SnapshotHistory, index Index) {

	mux.HandleFunc(
		"/snapshots/input", func(w http.ResponseWriter, r *http.Request) {

			b, err := history.GetInput(getUrlParams(r))
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
			b, err := history.GetOutput(getUrlParams(r))
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

func getUrlParams(r *http.Request) (string, []string, []string, []string) {
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
	return r.FormValue("format"), clusters, namespaces, resourceTypes
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

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
