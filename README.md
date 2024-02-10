<div align="center">

# cup **(under development 🚧)**

Run **docker compose** from everywhere.

[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=adrien-chinour_cup&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=adrien-chinour_cup)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=adrien-chinour_cup&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=adrien-chinour_cup)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=adrien-chinour_cup&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=adrien-chinour_cup)

</div>


## Usage

Cup is a standalone binary that allow to run `docker compose (up|down|stop)` without need to go to the directory.

Cup will scan all directories listed by `CUPDIR` environnement variable.

| Without cup 😱                                | With cup 🤩 |
|-----------------------------------------------|-------------|
| `cd ~/code/blog ; docker compose up -d; cd -` | `cup blog`  |

```shell
$ cup help

run docker compose up in detach mode on project

Usage:
  cup [flags]
  cup [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  down        run docker compose down on project
  help        Help about any command
  stop        run docker compose stop on project
  up          run docker compose up in detach mode on project

Flags:
  -h, --help      help for cup
  -v, --version   version for cup

Use "cup [command] --help" for more information about a command.
```
