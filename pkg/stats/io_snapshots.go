package stats

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type SnapshotHistory interface {
	// SetInput Sets the input snapshot for the given component
	SetInput(id string, latestInput json.Marshaler)
	// GetInput gets the input snapshot for all components
	GetInput() ([]byte, error)
	// SetOutput Sets the output snapshot for the given component
	SetOutput(id string, latestOutput json.Marshaler)
	// GetOutput gets the output snapshot for all component
	GetOutput() ([]byte, error)
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

func (h *snapshotHistory) GetInput() ([]byte, error) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	if h.latestInput == nil {
		return []byte{}, nil
	}
	result := map[string]interface{}{}
	for id, obj := range h.latestInput {
		jsn, err := obj.MarshalJSON()
		if err != nil {
			return nil, err
		}
		genericMap := map[string]interface{}{}
		if err := json.Unmarshal(jsn, &genericMap); err != nil {
			return nil, err
		}
		result[id] = genericMap
	}
	return json.Marshal(result)
}

func (h *snapshotHistory) SetOutput(id string, latestOutput json.Marshaler) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.latestOutput[id] = latestOutput
}

func (h *snapshotHistory) GetOutput() ([]byte, error) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	if h.latestOutput == nil {
		return []byte{}, nil
	}
	result := map[string]interface{}{}
	for id, obj := range h.latestOutput {
		jsn, err := obj.MarshalJSON()
		if err != nil {
			return nil, err
		}
		genericMap := map[string]interface{}{}
		if err := json.Unmarshal(jsn, &genericMap); err != nil {
			return nil, err
		}
		result[id] = genericMap
	}
	return json.Marshal(result)
}

func AddSnapshots(mux *http.ServeMux, history SnapshotHistory) {

	profileDescriptions["/snapshots/input"] = "Latest Input Snapshot"
	profileDescriptions["/snapshots/output"] = "Latest Output Snapshot"

	mux.HandleFunc(
		"/snapshots/input", func(w http.ResponseWriter, r *http.Request) {
			b, err := history.GetInput()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			_, _ = fmt.Fprint(w, string(b))
		},
	)

	mux.HandleFunc(
		"/snapshots/output", func(w http.ResponseWriter, r *http.Request) {
			b, err := history.GetOutput()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			_, _ = fmt.Fprint(w, string(b))
		},
	)
}
