package stats

import (
	"html/template"
	"net/http"
	"sort"
)

var profileDescriptions = map[string]string{}

type Index struct {
}

func NewIndex() Index {
	return Index{}
}

var indexTmpl = template.Must(template.New("index").Parse(`<!DOCTYPE html><html>
<head>
<script>
const change_urls = () => {
	var e = document.getElementById("format");
	var format = e.options[e.selectedIndex].text;
	var input = document.getElementById('input_url');
	var output = document.getElementById('output_url');
	if (format != '') {
		input.href = "/snapshots/input?format=" + format
		output.href = "/snapshots/output?format=" + format
	}
	else {
		input.href = "/snapshots/input"
		output.href = "/snapshots/output"
	}
};
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
<label for="format">Choose a format:</label>
<select name="format" id="format" onchange="change_urls()">
{{range .Formats}}
  <option value="{{.}}">{{.}}</option>
{{end}}
</select>
<h2><a href="/snapshots/input" id="input_url">/snapshots/input</a></h2>
<p>
Latest Input Snapshot
</p>
<h2><a href="/snapshots/output" id="output_url">/snapshots/output</a></h2>
<p>
Latest Output Snapshot
</p>
</body>
</html>
`))

func (p Index) Generate(w http.ResponseWriter, r *http.Request) {
	type profile struct {
		Name string
		Href string
		Desc string
	}
	type pageData struct {
		Profiles []profile
		Formats  []string
	}
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

	data := pageData{
		Profiles: profiles,
		Formats:  []string{"", "json", "yaml"},
	}

	indexTmpl.Execute(w, data)
}
