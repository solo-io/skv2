package stats

import (
	"html/template"
	"net/http"
	"sort"
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

// TODO - Use https://pkg.go.dev/embed once skv2 is upgraded to > go1.16
var indexTmpl = template.Must(template.New("index").Parse(`<!DOCTYPE html><html>
<head>
<script>
var format = "";
var clusters = [];
var namespaces = [];
var resourceTypes = [];

const init = () => {
	changeFormat()
	changeCluster()
	changeNamespace()
	changeResourceType()
};

const getSelectValues = (select) => {
	var result = [];
	var options = select && select.options;
	var opt;

	for (var i = 0, iLen = options.length; i < iLen; i++) {
		opt = options[i];

		if (opt.selected) {
			result.push(opt.value || opt.text);
		}
	}
	return result;
}
const generateUrls = () => {
	var input = document.getElementById('input_url');
	var output = document.getElementById('output_url');
	var input_d = document.getElementById('input_url_download');
	var output_d = document.getElementById('output_url_download');
	if (format === '' && clusters.length === 0 && namespaces.length === 0  && resourceTypes.length === 0) {
		input.href = "/snapshots/input";
		output.href = "/snapshots/output";
		input_d.href = "/snapshots/input";
		output_d.href = "/snapshots/output";
		return;
	}

	var params = [];
	if (format !== '') {
		params.push("format=" + format);
		input_d.download = "input." + format
		output_d.download = "output." + format
	} else {
		input_d.download = "input.json"
		output_d.download = "output.json"
	}
	if (clusters.length !== 0) {
		params.push("clusters=" + clusters.join("::"));
	}
	if (namespaces.length !== 0) {
		params.push("namespaces=" + namespaces.join("::"));
	}
	if (resourceTypes.length !== 0) {
		params.push("resourceTypes=" + resourceTypes.join("::"));
	}

	input.href = encodeURI("/snapshots/input?" + params.join("&"));
	output.href = encodeURI("/snapshots/output?" + params.join("&"));
	input_d.href = encodeURI("/snapshots/input?" + params.join("&"));
	output_d.href = encodeURI("/snapshots/output?" + params.join("&"));
}
const changeFormat = () => {
	var e = document.getElementById("format");
	if (e && e.options) {
		format = e.options[e.selectedIndex].text;
	}
	generateUrls();
};
const changeCluster = () => {
	var e = document.getElementById("cluster");
	clusters = getSelectValues(e);
	generateUrls();
};
const changeNamespace= () => {
	var e = document.getElementById("namespace");
	namespaces = getSelectValues(e);
	generateUrls();
};
const changeResourceType = () => {
	var e = document.getElementById("resourceType");
	resourceTypes = getSelectValues(e);
	generateUrls();
};
init()
</script>

<title>/debug/pprof/</title>
<style>
.profile-name{
	display:inline-block;
	width:6rem;
}
</style>
</head>
<body>
Things to do now:
{{range .Profiles}}
<h2><a href={{.Href}}>{{.Name}}</a></h2>
<p>
{{.Desc}}
</p>
{{end}}
<br>
<br>
<h2>Snapshot Format & Filters</h2>
<p>All options apply to both snapshots. Filters are inclusive.</p>
<br>
<label style="vertical-align:top" for="format">Choose a format:</label>
<select style="vertical-align:top" name="format" id="format" onchange="changeFormat()">
{{range .Formats}}
  <option value="{{.}}">{{.}}</option>
{{end}}
</select>
<label style="vertical-align:top" for="cluster">Choose cluster(s):</label>
<select style="height: 250px; vertical-align:top" for="format" name="cluster" id="cluster" multiple="multiple" onchange="changeCluster()">
{{range .Clusters}}
  <option value="{{.}}">{{.}}</option>
{{end}}
</select>
<label style="vertical-align:top" for="namespace">Choose namspace(s):</label>
<select style="height: 250px; vertical-align:top" for="format" name="namespace" id="namespace" multiple="multiple" onchange="changeNamespace()">
{{range .Namespaces}}
  <option value="{{.}}">{{.}}</option>
{{end}}
</select>
<label style="vertical-align:top" for="resourceType">Choose resource type(s):</label>
<select style="height: 250px; vertical-align:top" for="format" name="resourceType" id="resourceType" multiple="multiple" onchange="changeResourceType()">
{{range .ResourceTypes}}
  <option value="{{.}}">{{.}}</option>
{{end}}
</select>
<h2>Input</h2>
<a href="/snapshots/input" id="input_url">View</a>
<a href="/snapshots/input" id="input_url_download" download="input.json">Download</a>
<h2>Output</h2>
<a href="/snapshots/output" id="output_url">View</a>
<a href="/snapshots/output" id="output_url_download" download="output.json">Download</a>
</body>
</html>
`))

func (index Index) Generate(w http.ResponseWriter, r *http.Request) {
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
	var profiles []profile
	clusters := make(map[string]bool)
	namespaces := make(map[string]bool)

	// Capture clusters, namespaces, resourceTypes from input.
	// We can use these as filters for both output & input
	input, err := index.history.GetInputCopy()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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

	data := pageData{
		Profiles:      profiles,
		Formats:       []string{"", "json", "yaml"},
		ResourceTypes: resourceTypes,
		Clusters:      clusterSlice,
		Namespaces:    namespaceSlice,
	}

	templateErr := indexTmpl.Execute(w, data)
	if templateErr != nil {
		http.Error(w, templateErr.Error(), http.StatusInternalServerError)
		return
	}
}
