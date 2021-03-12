package stats

import (
	"fmt"
	"log"
	"net/http"
)

func MustStartServerBackground(snapshotHistory *SnapshotHistory, port uint32, addHandlers ...func(mux *http.ServeMux, profiles map[string]string)) {
	go func() {
		if err := StartServer(snapshotHistory, port, addHandlers...); err != nil {
			log.Fatal(err)
		}
	}()
}

func StartServer(snapshotHistory *SnapshotHistory, port uint32, addHandlers ...func(mux *http.ServeMux, profiles map[string]string)) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	AddPprof(mux)
	AddMetrics(mux)
	AddSnapshots(mux, snapshotHistory)
	for _, h := range addHandlers {
		h(mux, profileDescriptions)
	}
	return http.ListenAndServe(fmt.Sprintf(":%v", port), mux)
}
