package render

import (
	"github.com/solo-io/skv2/codegen/model"
)

type ChartRenderer struct {
	templateRenderer
}

var defaultChartInputs = inputTemplates{
	"chart/namespace.yamltmpl": {
		Path: "templates/namespace.yaml",
	},
	"chart/operator-deployment.yamltmpl": {
		Path: "templates/deployment.yaml",
	},
	"chart/operator-configmap.yamltmpl": {
		Path: "templates/configmap.yaml",
	},
	"chart/operator-rbac.yamltmpl": {
		Path: "templates/rbac.yaml",
	},
	"chart/values.yamltmpl": {
		Path: "values.yaml",
	},
	"chart/Chart.yamltmpl": {
		Path: "Chart.yaml",
	},
}

func RenderChart(chart model.Chart) ([]OutFile, error) {
	renderer := DefaultTemplateRenderer

	// when rendering helm charts, we need
	// to use a custom delimiter
	renderer.left = "[["
	renderer.right = "]]"

	defaultChartRenderer := ChartRenderer{
		templateRenderer: renderer,
	}
	return defaultChartRenderer.Render(chart)
}

func (r ChartRenderer) Render(chart model.Chart) ([]OutFile, error) {
	templatesToRender := defaultChartInputs

	files, err := r.renderCoreTemplates(templatesToRender, chart)
	if err != nil {
		return nil, err
	}

	customFiles, err := r.RenderCustomTemplates(chart.CustomTemplates.Templates, chart.CustomTemplates.Funcs, chart)
	if err != nil {
		return nil, err
	}

	files = append(files, customFiles...)

	if chart.FilterTemplate != nil {
		var filteredFiles []OutFile
		for _, file := range files {
			if chart.FilterTemplate(file.Path) {
				continue
			}
			filteredFiles = append(filteredFiles, file)
		}
		files = filteredFiles
	}

	return files, nil
}
