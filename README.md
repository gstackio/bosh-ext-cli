# bosh-ext-cli

BOSH Extended CLI is a tool that provides additional set of commands for interacting with BOSH and release authoring

## Commands

- Lint release directory and show suggestions

```
$ out/bosh-ext lint-release
$ out/bosh-ext lint-release --dir ~/workspace/whatever-release
```

- Lint manifest

```
$ out/bosh-ext lint-manifest manifest.yml
```

- Show detailed Director task debug information

```
$ bosh task X --debug | out/bosh-ext debug-task -
$ bosh task X --debug | out/bosh-ext debug-task - -a
$ bosh task X --debug | out/bosh-ext debug-task - -a -s duration
$ bosh task X --debug | out/bosh-ext debug-task - -l
```

- Web view (useful with information dense commands such as `bosh events`)

```
$ export BOSH_ENVIRONMENT=vbox
$ out/bosh-ext web
```

- Visualize Deployments, Links, Releases, Tasks, Events, Configs and Variables
  in a web browser

```
$ out/bosh-ext visualize
Starting server on http://127.0.0.1:9090
```

The visualization features have been presented in
“(_Making Sense of BOSH Links Through Visualization_)[https://youtu.be/OnZePLee_Po]”,
a talk given by Jamil Shamy & Saman Alvi in October 2018, at 2nd CF European Summit in Basel.


## Build & Development

```
$ git clone ...
$ cd bosh-ext-cli
$ source .envrc
$ bin/build
$ out/bosh-ext -v
```

## Todo

- linting
  - release: notice common props between release jobs
  - release: multiple jobs?
  - release: camelCase vs snake case
  - `set -e` in packaging?
  - pre_packaging presence?
  - job description presence
  - `type: password` annotation
  - extrapolate links usage
  - discover_external_ip presence
  - logrotate presence
  - todo markers
  - consolidate explanation
  - greedy blob inclusion?
- web
  - reload events
  - output for currently running tasks
  - bosh task (errored task) -> no error
- debug-task
  - columns
