package stats

import (
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
)

var profileDescriptions = map[string]string{}

type Index struct {
	history SnapshotHistory
}

func NewIndex(history SnapshotHistory) Index {
	return Index{
		history: history,
	}
}

type profile struct {
	Name string
	Href string
	Desc string
}
type pageData struct {
	Profiles      []profile
	Formats       []string
	ResourceTypes []string
	Clusters      []string
	Namespaces    []string
}

//go:embed html/index.html
var indexHtml string

//go:embed html/fallback.html
var fallbackHtml string

func (index Index) Generate(w http.ResponseWriter, r *http.Request) {

	var profiles []profile

	// Adding other profiles exposed from within this package
	for p, pd := range profileDescriptions {
		profiles = append(profiles, profile{
			Name: p,
			Href: p,
			Desc: pd,
		})
	}

	sort.Slice(profiles, func(i, j int) bool {
		return profiles[i].Name < profiles[j].Name
	})

	template, data := getTemplate(index, profiles)
	templateErr := template.Execute(w, data)
	if templateErr != nil {
		http.Error(w, templateErr.Error(), http.StatusInternalServerError)
		return
	}
}

type getIndexTemplateResult struct {
	Template *template.Template
	PageData pageData
	Error    error
}

func getTemplate(index Index, profiles []profile) (*template.Template, pageData) {
	result := make(chan getIndexTemplateResult, 1)
	go func() {
		result <- getIndexTemplate(index, profiles)
	}()
	select {
	case <-time.After(time.Duration(1.5 * float64(time.Second))):
		log.Println("Timed out getting filters. Revert to fallback template.")
		return getFallbackTemplate(index, profiles)
	case result := <-result:
		if result.Error != nil {
			log.Printf("Unhandled error. Revert to fallback template. Error: %s\n", result.Error.Error())
			return getFallbackTemplate(index, profiles)
		}
		return result.Template, result.PageData

	}
}

func getIndexTemplate(index Index, profiles []profile) (result getIndexTemplateResult) {
	defer func() {
		if r := recover(); r != nil {
			result = getIndexTemplateResult{
				Error: fmt.Errorf("Recovered. Error: %s", r),
			}
		}
	}()

	clusters := make(map[string]bool)
	namespaces := make(map[string]bool)

	// Capture clusters, namespaces, resourceTypes from input.
	// We can use these as filters for both output & input
	input, err := index.history.GetInputCopy()
	if err != nil {
		return getIndexTemplateResult{
			Error: err,
		}
	}
	translator := input["translator"].(map[string]interface{})
	resourceTypes := make([]string, 0, len(translator))
	for id := range translator {
		resourceTypes = append(resourceTypes, id)
	}

	for resourceType, resources := range translator {
		for _, resource := range resources.([]interface{}) {
			resource := resource.(map[string]interface{})
			metadata := resource["metadata"].(map[string]interface{})
			cluster, ok := metadata["clusterName"].(string)
			if ok && cluster != "" {
				clusters[cluster] = true
			}
			namespace, ok := metadata["namespace"].(string)
			if resourceType == "/v1, Kind=Namespace" {
				namespace, ok = metadata["name"].(string)
			}
			if ok && namespace != "" {
				namespaces[namespace] = true
			}
		}
	}
	clusterSlice := make([]string, 0, len(clusters))
	for k := range clusters {
		clusterSlice = append(clusterSlice, k)
	}
	namespaceSlice := make([]string, 0, len(namespaces))
	for k := range namespaces {
		namespaceSlice = append(namespaceSlice, k)
	}

	sort.Strings(resourceTypes)
	sort.Strings(clusterSlice)
	sort.Strings(namespaceSlice)

	return getIndexTemplateResult{
		Template: template.Must(template.New("index").Parse(indexHtml)),
		PageData: pageData{
			Profiles:      profiles,
			Formats:       []string{"", "json", "yaml"},
			ResourceTypes: resourceTypes,
			Clusters:      clusterSlice,
			Namespaces:    namespaceSlice,
		},
		Error: nil,
	}
}

func getFallbackTemplate(index Index, profiles []profile) (*template.Template, pageData) {
	data := pageData{
		Profiles: profiles,
		Formats:  []string{"", "json", "yaml"},
	}
	return template.Must(template.New("index").Parse(fallbackHtml)), data
}
