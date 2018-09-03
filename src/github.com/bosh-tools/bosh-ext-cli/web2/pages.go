package web2

import (
	"bytes"
	"text/template"
)

type BOSHPage struct {
	Title        string
	BodyTemplate string
	Navigation   string
	CustomJS     string
}

var boshPages = map[string]BOSHPage{
	"home": BOSHPage{
		Title:        "My Dashboard",
		BodyTemplate: homeTemplate,
		CustomJS:     homeJS,
	},
	"deployments": BOSHPage{
		Title:        "Deployments",
		BodyTemplate: deploymentsTemplate,
		CustomJS:     deploymentsJS,
	},
	"events": BOSHPage{
		Title:        "Events",
		BodyTemplate: eventsTemplate,
		CustomJS:     eventsJS,
	},
	"releases": BOSHPage{
		Title:        "Releases",
		BodyTemplate: releasesTemplate,
		CustomJS:     releasesJS,
	},
	"tasks-logs": BOSHPage{
		Title:        "Task Logs",
		BodyTemplate: logsViewerTemplate,
		CustomJS:     logsViewerJS,
	},
	"tasks": BOSHPage{
		Title:        "Tasks",
		BodyTemplate: tasksTemplate,
		CustomJS:     tasksJS,
	},
}

func GenerateBOSHPage(pageName string) (string, error) {
	tmpl, err := template.New("tmp").Parse(
		generateTemplate(
			boshPages[pageName].BodyTemplate,
			boshPages[pageName].CustomJS,
		),
	)

	if err != nil {
		return "", err
	}

	var renderedTemplate bytes.Buffer
	err = tmpl.Execute(&renderedTemplate, boshPages[pageName])
	if err != nil {
		return "", err
	}

	return renderedTemplate.String(), nil
}
