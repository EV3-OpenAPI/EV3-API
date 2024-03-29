[![Go binary](https://github.com/EV3-OpenAPI/EV3-API/actions/workflows/build.yaml/badge.svg)](https://github.com/EV3-OpenAPI/EV3-API/actions/workflows/build.yaml)

![ZHAW-logo](https://upload.wikimedia.org/wikipedia/commons/thumb/e/e6/ZHAW_Logo.svg/206px-ZHAW_Logo.svg.png)

# LEGO® MINDSTORMS® EV3-REST

This project has the goal of designing a system for the LEGO® MINDSTORMS® EV3 robot to promote programming knowledge for non-computer science students, in such a way that the implemented methods in the server are callable over several programming languages. The implemented methods should be callable via an application programming interface.
The server which runs on the LEGO® robot is implemented with [OpenAPI Specification](https://www.openapis.org/), which allows to automatically generate libraries for the client in several programming languages.

This GitHub project is the starting point for the LEGO® MINDSTORMS® EV3 robot REST implementation. Here you will find the information you need about the LEGO® Robot, simple examples of what it looks like, and some general information regarding the project.

## Installation

1. Connect to your EV3 robot with SSH or with a directly attached keyboard and monitor.
2. Ensure that the EV3 has an internet connection
3. Run the install-script with elevated privileges 

```bash
curl -sf -L https://raw.githubusercontent.com/EV3-OpenAPI/EV3-API/master/scripts/install.sh | sudo sh
```

## Tools and Libraries

For this project, **OpenAPI generator** is used and the **EV3 Golang library**.

* [OpenAPI Generator](https://openapi-generator.tech/)
* [EV3 Golang library](https://github.com/ev3go)

## Setup for Students

If you want to code with **Java**, then use this instruction:

* [Setup for Java](./SetupJava.md)

If you want to code with **Python**, then use this instruction:

* [Setup for Python](./SetupPython.md)

# Participation

The current process for development of the LEGO® robot is described in [Development Guidelines](https://github.com/EV3-OpenAPI/EV3-API/blob/master/DEVELOPMENT.md). Development of the next version of the LEGO® robot is guided by the [ZHAW School of Engineering](https://www.zhaw.ch/en/engineering/). This group of committers bring their API expertise, incorporate feedback from the community, and expand the group of committers as appropriate. All development activity on the future specification will be performed as features and merged into this branch. Upon release of the future specification, this branch will be merged to <code>main</code>.

The LEGO® EV3 robot encourages participation from individuals and companies alike. If you want to participate in the evolution of the LEGO® robot, consider taking the following actions:

* Review the current specification. The human-readable markdown file is the source of truth for the specification.
* Review the [development](https://github.com/EV3-OpenAPI/EV3-API/blob/master/DEVELOPMENT.md) process, so you understand how the spec is evolving.
* Check the [issues](https://github.com/EV3-OpenAPI/EV3-API/issues) and [pull requests](https://github.com/EV3-OpenAPI/EV3-API/pulls) to see if someone has already documented your idea or feedback on the specification. You can follow an existing conversation by subscribing to the existing issue or PR.
* Create an issue to describe a new concern. If possible, propose a solution.

Not all feedback can be accommodated and there may be solid arguments for or against a change being appropriate for the specification.

## Setup development

### Prerequisites

Install GO: [Offical website](https://go.dev/dl/), version >= 1.18  
Install either [Docker](https://docs.docker.com/engine/install/) or [Openapi-Generator](https://openapi-generator.tech/docs/installation) for generating the server code.  
Install `goimports` tool: `go install golang.org/x/tools/cmd/goimports@latest` for fixing unused imports in generated code.

#### Generate Server Code

Using Docker:

```bash
docker run --rm \
  --name openapi-generator -u 1000 -w /local \
  -v $PWD:/local openapitools/openapi-generator-cli generate \
  -i /local/openapi/spec.yaml \
  -g go-server \
  -o /local/internal/gen \
  -c /local/openapi/server-config.yml
```

Using Openapi-Generator:

```bash
openapi-generator generate -i openapi/spec.yaml -o internal/gen -g go-server -c openapi/server-config.yml
```

Clean up unused imports (go refuses to compile if you don't):

```bash
goimports -l -w internal/gen/openapi
```

Clean up generated code (because why not?): 

```bash
gofmt -l -w internal/gen/openapi
```

#### Compile Server Binary

The `ldflags` option will remove debug code and result in a smaller binary.  

```bash
go build -o server -ldflags="-s -w" EV3-API/cmd
```

#### Generate Java Client Code

Using Docker:

```bash
docker run --rm \
  --name openapi-generator -u 1000 -w /local \
  -v $PWD:/local openapitools/openapi-generator-cli generate \
  -i /local/openapi/spec.yaml \
  -g java \
  -o clients/ev3-java/ev3api \
  -c openapi/java-client-config.yaml
```

Using Openapi-Generator:

```bash
openapi-generator generate -i openapi/spec.yaml -o clients/ev3-java/ev3api -g java -c openapi/java-client-config.yaml
```

#### Generate Python Client Code

Using Docker:

```bash
docker run --rm \
  --name openapi-generator -u 1000 -w /local \
  -v $PWD:/local openapitools/openapi-generator-cli generate \
  -i /local/openapi/spec.yaml \
  -g python \
  -o clients/ev3-python \
  -c openapi/python-client-config.yaml
```

Using Openapi-Generator:

```bash
openapi-generator generate -i openapi/spec.yaml -o clients/ev3-python -g python -c openapi/python-client-config.yaml
```

## Release and Deployment

**Important**: The versioning follows [Go semver](https://pkg.go.dev/golang.org/x/mod/semver) specifications, meaning versions have to be in the format "vMAJOR.MINOR.BUGFIX" including the leading "v". This is important when creating a release tag.

If a new release is to be created, the new version number must first be entered in the following files:

* openapi/spec.yaml
* clients/ev3-java/build.gradle
* clients/ev3-python/setup.py

After that, a new release can be published via the [GitHub GUI](https://github.com/EV3-OpenAPI/EV3-API/releases/new). The repository has a GitHub action that automatically creates the Golang, Java and Python artifacts and attaches them to the release.

Every time the LEGO robot is started, the [GitHub API](https://api.github.com/repos/EV3-OpenAPI/EV3-API/releases/latest) is used to check whether the latest release is newer than the currently installed version. If this is the case, the latest GO binary will be downloaded via the Assets API. This replaces the old binary and will be executed in its place in the future.
