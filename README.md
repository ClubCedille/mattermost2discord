# mattermost2discord

[![](https://github.com/cguertin14/advent-of-code-2020/workflows/CI/badge.svg)](https://github.com/ClubCedille/mattermost2discord/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/ClubCedille/mattermost2discord)](https://goreportcard.com/report/github.com/ClubCedille/mattermost2discord)

REST API that forwards messages from Mattermost to Discord, using a callback.

## Depencencies

* [Go 1.16](https://golang.org/dl/)
* [Docker](https://docs.docker.com/get-docker/)
* [Docker-Compose](https://docs.docker.com/compose/install/)

## Usage

### With Kubernetes

To use `mattermost2discord` with Kubernetes, simply declare a `kustomization.yml` configuration file like so:
```yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

resources:
- https://github.com/ClubCedille/mattermost2discord?ref=v1.0.0 # Example release
```

Then, you can build your kustomize configuration like this:
```bash
$ kustomize build
...
```

## Development

### Running the API - using Docker

To run the API, you can use the following command:
```bash
$ make docker
...
```

The service will then run on the port 3000.

### Running the API - without Docker

You can also run the API without Docker, but make sure the `PORT` environment variable is set on your machine to something like `3000` or any other port available on your computer. You can run it with the following command:
```bash
$ make run
...
```

### Running the tests

To run the tests, simply execute this command:
```bash
$ make test
...
```
### Git hooks

To share the same versioned across the team, execute this command:
```bash
$ git config core.hooksPath .githooks
```
