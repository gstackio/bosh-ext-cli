package visualize

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
	"home": {
		Title:        "My Dashboard",
		BodyTemplate: homeTemplate,
		CustomJS:     homeJS,
	},
	"deployments": {
		Title:        "Deployments",
		BodyTemplate: deploymentsTemplate,
		CustomJS:     deploymentsJS,
	},
	"events": {
		Title:        "Events",
		BodyTemplate: eventsTemplate,
		CustomJS:     eventsJS,
	},
	"releases": {
		Title:        "Releases",
		BodyTemplate: releasesTemplate,
		CustomJS:     releasesJS,
	},
	"link-providers": {
		Title:        "Link Providers",
		BodyTemplate: linkProvidersTemplate,
		CustomJS:     linkProvidersJS,
	},
	"link-consumers-detailed": {
		Title:        "Link Consumers",
		BodyTemplate: linkConsumersDetailedTemplate,
		CustomJS:     linkConsumersDetailedJS,
	},
	"links-deployment-dependencies": {
		Title:        "Links Deployment Dependencies",
		BodyTemplate: linksDeploymentDependenciesTemplate,
		CustomJS:     linksDeploymentDependenciesJS,
	},
	"tasks-logs": {
		Title:        "Task Logs",
		BodyTemplate: logsViewerTemplate,
		CustomJS:     logsViewerJS,
	},
	"tasks": {
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
