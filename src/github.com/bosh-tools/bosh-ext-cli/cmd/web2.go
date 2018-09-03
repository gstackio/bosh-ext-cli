package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	boshui "github.com/cloudfoundry/bosh-cli/ui"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"

	"github.com/bosh-tools/bosh-ext-cli/web2"
)

type Web2Cmd struct {
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

func NewWeb2Cmd(cmdRunner boshsys.CmdRunner, ui boshui.UI, logger boshlog.Logger) Web2Cmd {
	return Web2Cmd{
		cmdRunner: cmdRunner,
		ui:        ui,

		logTag: "WebCmd",
		logger: logger,

		allowedCmds: map[string][]apiOpt{
			"env":         []apiOpt{},
			"deployments": []apiOpt{},
			"instances": []apiOpt{
				apiOpt{Name: "deployment"},
				apiOpt{Name: "ps", WithoutValue: true},
				apiOpt{Name: "details", WithoutValue: true},
			},
			"releases": []apiOpt{},
			"tasks": []apiOpt{
				apiOpt{Name: "recent", WithoutValue: true},
				apiOpt{Name: "all", WithoutValue: true},
			},
			"task": []apiOpt{
				apiOpt{Name: "id", Positional: true},
				apiOpt{Name: "debug", WithoutValue: true},
			},
			"events": []apiOpt{
				apiOpt{Name: "action", EqualsSign: true},
				apiOpt{Name: "deployment", EqualsSign: true},
				apiOpt{Name: "instance", EqualsSign: true},
				apiOpt{Name: "object-name", EqualsSign: true},
				apiOpt{Name: "object-type", EqualsSign: true},
				apiOpt{Name: "task", EqualsSign: true},
				apiOpt{Name: "event-user", EqualsSign: true},
				apiOpt{Name: "before", EqualsSign: true},
				apiOpt{Name: "after", EqualsSign: true},
				apiOpt{Name: "before-id", EqualsSign: true},
			},
		},
	}
}

func (c Web2Cmd) Run(opts Web2Opts) error {
	http.HandleFunc("/", c.serveHomePage)
	http.HandleFunc("/deployments", c.serveDeploymentsPage)
	http.HandleFunc("/events", c.serveEventsPage)
	http.HandleFunc("/tasks-logs", c.serveTasksLogsPage)
	http.HandleFunc("/releases", c.serveReleasesPage)
	http.HandleFunc("/tasks", c.serveTasksPage)
	http.HandleFunc("/css/sb-admin.css", c.serveCSS)
	http.HandleFunc("/js/sb-admin.min.js", c.serveJS)
	http.HandleFunc("/api/command", c.serveAPICommand)

	addr := fmt.Sprintf("%s:%d", opts.ListenAddr, opts.ListenPort)
	c.ui.PrintLinef("Starting server on http://%s", addr)

	return http.ListenAndServe(addr, nil)
}

func (c Web2Cmd) serveHomePage(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug(c.logTag, "Serving Home Page")
	renderedPage, _ := web2.GenerateBOSHPage("home")
	fmt.Fprintf(w, renderedPage)
}

func (c Web2Cmd) serveDeploymentsPage(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug(c.logTag, "Serving Deployments Page")
	renderedPage, _ := web2.GenerateBOSHPage("deployments")
	fmt.Fprintf(w, renderedPage)
}

func (c Web2Cmd) serveEventsPage(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug(c.logTag, "Serving Events Page")
	renderedPage, _ := web2.GenerateBOSHPage("events")
	fmt.Fprintf(w, renderedPage)
}

func (c Web2Cmd) serveTasksLogsPage(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug(c.logTag, "Serving Tasks Logs Page")
	renderedPage, _ := web2.GenerateBOSHPage("tasks-logs")
	fmt.Fprintf(w, renderedPage)
}

func (c Web2Cmd) serveReleasesPage(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug(c.logTag, "Serving releases Page")
	renderedPage, _ := web2.GenerateBOSHPage("releases")
	fmt.Fprintf(w, renderedPage)
}

func (c Web2Cmd) serveTasksPage(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug(c.logTag, "Serving tasks Page")
	renderedPage, _ := web2.GenerateBOSHPage("tasks")
	fmt.Fprintf(w, renderedPage)
}

func (c Web2Cmd) serveCSS(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug(c.logTag, "Serving CSS")
	w.Header().Add("Content-Type", "text/css")

	fmt.Fprintf(w, web2.AdminCSS)
}

func (c Web2Cmd) serveJS(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug(c.logTag, "Serving JS")
	w.Header().Add("Content-Type", "application/javascript")
	fmt.Fprintf(w, web2.AdminJS)
}

func (c Web2Cmd) serveAPICommand(w http.ResponseWriter, r *http.Request) {
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

	apiOpts, found := c.allowedCmds[cmdName]
	if !found {
		c.logger.Error(c.logTag, "Disallowed cmd '%s'", cmdName)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmd := boshsys.Command{
		Name: "bosh",
		Args: []string{cmdName, "--json"},
	}

	providedOpts := theRequest.Arguments

	for _, providedOpt := range providedOpts {

		if len(providedOpt.Name) == 0 {
			continue
		}
		builtinOpt, found := c.fetchCmdOption(apiOpts, providedOpt.Name)
		if !found {
			continue
		}

		providedVal := providedOpt.Value

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

func (c Web2Cmd) fetchCmdOption(cmdOptions []apiOpt, optName string) (apiOpt, bool) {
	for _, opt := range cmdOptions {
		if optName == opt.Name {
			return opt, true
		}
	}
	return apiOpt{}, false
}
