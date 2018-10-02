package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	boshui "github.com/cloudfoundry/bosh-cli/ui"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"

	"github.com/bosh-tools/bosh-ext-cli/visualize"
)

type VisualizeCmd struct {
	cmdRunner boshsys.CmdRunner
	ui        boshui.UI

	logTag string
	logger boshlog.Logger

	allowedCmds map[string][]apiOpt
}

type RequestArgument struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type RequestBody struct {
	Command   string            `json:"command"`
	Arguments []RequestArgument `json:"arguments"`
}

func NewVisualizeCmd(cmdRunner boshsys.CmdRunner, ui boshui.UI, logger boshlog.Logger) VisualizeCmd {
	return VisualizeCmd{
		cmdRunner: cmdRunner,
		ui:        ui,

		logTag: "WebCmd",
		logger: logger,

		allowedCmds: map[string][]apiOpt{
			"env":         {},
			"deployments": {},
			"instances": {
				{Name: "deployment"},
				{Name: "ps", WithoutValue: true},
				{Name: "details", WithoutValue: true},
			},
			"curl": {
				{Name: "path", Curl: true, Positional: true},
			},
			"releases": {},
			"tasks": {
				{Name: "recent", WithoutValue: true},
				{Name: "all", WithoutValue: true},
			},
			"task": {
				{Name: "id", Positional: true},
				{Name: "debug", WithoutValue: true},
			},
			"events": {
				{Name: "action", EqualsSign: true},
				{Name: "deployment", EqualsSign: true},
				{Name: "instance", EqualsSign: true},
				{Name: "object-name", EqualsSign: true},
				{Name: "object-type", EqualsSign: true},
				{Name: "task", EqualsSign: true},
				{Name: "event-user", EqualsSign: true},
				{Name: "before", EqualsSign: true},
				{Name: "after", EqualsSign: true},
				{Name: "before-id", EqualsSign: true},
			},
		},
	}
}

func (c VisualizeCmd) Run(opts VisualizeOpts) error {
	http.HandleFunc("/", c.serveHomePage)
	http.HandleFunc("/deployments", c.serveDeploymentsPage)
	http.HandleFunc("/events", c.serveEventsPage)
	http.HandleFunc("/tasks-logs", c.serveTasksLogsPage)
	http.HandleFunc("/releases", c.serveReleasesPage)
	http.HandleFunc("/link-providers", c.serveLinkProvidersPage)
	http.HandleFunc("/tasks", c.serveTasksPage)
	http.HandleFunc("/css/sb-admin.css", c.serveCSS)
	http.HandleFunc("/js/sb-admin.min.js", c.serveJS)
	http.HandleFunc("/api/command", c.serveAPICommand)

	addr := fmt.Sprintf("%s:%d", opts.ListenAddr, opts.ListenPort)
	c.ui.PrintLinef("Starting server on http://%s", addr)

	return http.ListenAndServe(addr, nil)
}

func (c VisualizeCmd) serveHomePage(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug(c.logTag, "Serving Home Page")
	renderedPage, _ := visualize.GenerateBOSHPage("home")
	fmt.Fprintf(w, renderedPage)
}

func (c VisualizeCmd) serveDeploymentsPage(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug(c.logTag, "Serving Deployments Page")
	renderedPage, _ := visualize.GenerateBOSHPage("deployments")
	fmt.Fprintf(w, renderedPage)
}

func (c VisualizeCmd) serveEventsPage(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug(c.logTag, "Serving Events Page")
	renderedPage, _ := visualize.GenerateBOSHPage("events")
	fmt.Fprintf(w, renderedPage)
}

func (c VisualizeCmd) serveTasksLogsPage(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug(c.logTag, "Serving Tasks Logs Page")
	renderedPage, _ := visualize.GenerateBOSHPage("tasks-logs")
	fmt.Fprintf(w, renderedPage)
}

func (c VisualizeCmd) serveReleasesPage(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug(c.logTag, "Serving releases Page")
	renderedPage, _ := visualize.GenerateBOSHPage("releases")
	fmt.Fprintf(w, renderedPage)
}

func (c VisualizeCmd) serveLinkProvidersPage(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug(c.logTag, "Serving Link Providers Page")
	renderedPage, _ := visualize.GenerateBOSHPage("link-providers")
	fmt.Fprintf(w, renderedPage)
}

func (c VisualizeCmd) serveTasksPage(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug(c.logTag, "Serving tasks Page")
	renderedPage, _ := visualize.GenerateBOSHPage("tasks")
	fmt.Fprintf(w, renderedPage)
}

func (c VisualizeCmd) serveCSS(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug(c.logTag, "Serving CSS")
	w.Header().Add("Content-Type", "text/css")

	fmt.Fprintf(w, visualize.AdminCSS)
}

func (c VisualizeCmd) serveJS(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug(c.logTag, "Serving JS")
	w.Header().Add("Content-Type", "application/javascript")
	fmt.Fprintf(w, visualize.AdminJS)
}

func (c VisualizeCmd) serveAPICommand(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug(c.logTag, "Serving API command")

	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var theRequest RequestBody
	err = json.Unmarshal(b, &theRequest)
	if err != nil {
		panic(err)
		http.Error(w, err.Error(), 500)
		return
	}

	c.logger.Debug(c.logTag, "Form submitted: %#v", r.Form)

	cmdName := theRequest.Command
	if len(cmdName) == 0 {
		c.logger.Error(c.logTag, "Empty command")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmdAllowedApiOpts, found := c.allowedCmds[cmdName]
	if !found {
		c.logger.Error(c.logTag, "Disallowed cmd '%s'", cmdName)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmd := boshsys.Command{
		Name: "bosh",
		Args: []string{cmdName},
	}

	if cmdName != "curl" {
		cmd.Args = append(cmd.Args, "--json")
	}

	requestPassedInOpts := theRequest.Arguments

	for _, requestProvidedOpt := range requestPassedInOpts {

		if len(requestProvidedOpt.Name) == 0 {
			continue
		}

		builtinOpt, found := c.fetchCmdOption(cmdAllowedApiOpts, requestProvidedOpt.Name)
		if !found {
			continue
		}

		providedVal := requestProvidedOpt.Value

		if builtinOpt.WithoutValue {
			if len(providedVal) > 0 {
				c.logger.Error(c.logTag, "Expected opt '%s' to not have value", builtinOpt.Name)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			cmd.Args = append(cmd.Args, "--"+builtinOpt.Name)
		} else {

			if len(providedVal) == 0 {
				continue
			}

			if builtinOpt.Positional {
				cmd.Args = append(cmd.Args, providedVal)
			} else if builtinOpt.EqualsSign {
				cmd.Args = append(cmd.Args, "--"+builtinOpt.Name+"="+providedVal)
			} else {
				cmd.Args = append(cmd.Args, "--"+builtinOpt.Name, providedVal)
			}
		}
	}

	stdout, _, _, err := c.cmdRunner.RunComplexCommand(cmd)
	if err != nil && cmdName != "task" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// if err != nil && cmdName == "task" {
	// 	stdout = stderr
	// }

	w.Header().Add("Content-Type", "application/json")

	_, err = w.Write([]byte(stdout))
	if err != nil {
		c.logger.Error(c.logTag, "Failed to write API events response")
	}
}

func (c VisualizeCmd) fetchCmdOption(cmdOptions []apiOpt, optName string) (apiOpt, bool) {
	for _, opt := range cmdOptions {
		if optName == opt.Name {
			return opt, true
		}
	}
	return apiOpt{}, false
}
