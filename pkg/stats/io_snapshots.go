package stats

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"gopkg.in/yaml.v2"
)

type SnapshotHistory interface {
	// SetInput Sets the input snapshot for the given component
	SetInput(id string, latestInput json.Marshaler)
	// GetInput gets the input snapshot for all components
	GetInput(format string) ([]byte, error)
	// SetOutput Sets the output snapshot for the given component
	SetOutput(id string, latestOutput json.Marshaler)
	// GetOutput gets the output snapshot for all component
	GetOutput(format string) ([]byte, error)
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

func (h *snapshotHistory) GetInput(format string) ([]byte, error) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	if h.latestInput == nil {
		return []byte{}, nil
	}
	genericMaps, err := getGenericMaps(h.latestInput)
	if err != nil {
		return nil, err
	}
	return formatMap(format, genericMaps)
}

func (h *snapshotHistory) SetOutput(id string, latestOutput json.Marshaler) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.latestOutput[id] = latestOutput
}

func (h *snapshotHistory) GetOutput(format string) ([]byte, error) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	if h.latestOutput == nil {
		return []byte{}, nil
	}
	genericMaps, err := getGenericMaps(h.latestOutput)
	if err != nil {
		return nil, err
	}
	return formatMap(format, genericMaps)
}

func AddSnapshots(mux *http.ServeMux, history SnapshotHistory) {

	mux.HandleFunc(
		"/snapshots/input", func(w http.ResponseWriter, r *http.Request) {

			b, err := history.GetInput(r.FormValue("format"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			_, _ = fmt.Fprint(w, string(b))
		},
	)

	mux.HandleFunc(
		"/snapshots/output", func(w http.ResponseWriter, r *http.Request) {
			b, err := history.GetOutput(r.FormValue("format"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			_, _ = fmt.Fprint(w, string(b))
		},
	)
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
