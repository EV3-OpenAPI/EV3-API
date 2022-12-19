[![Go binary](https://github.com/EV3-OpenAPI/EV3-API/actions/workflows/build.yaml/badge.svg)](https://github.com/PA-arslasel-machitic/EV3-API/actions/workflows/build.yaml)

![ZHAW-logo](https://upload.wikimedia.org/wikipedia/commons/thumb/e/e6/ZHAW_Logo.svg/206px-ZHAW_Logo.svg.png)


# LEGO MINDSTORMS EV3-REST

This project deals with the goal to design the Lego Robot of the type mindstorm EV3 for the promotion of programming knowledge for non-informaticians, in such a way that the implemented methods in the server are callable over several programming technologies. The implemented methods should be callable via an application programming interface.
The server which is the Lego robot is implemented with [OpenAPI Specification](https://www.openapis.org/), which allows to automatically generate libraries for the client in several programming languages.

This GitHub project is the starting point for the Lego Robot Mindstorm EV3-RESt. Here you will find the information you need about the Lego Robot, simple examples of what it looks like, and some general information regarding the project.

# Tools and Libraries

For this project, **OpenAPI generator** is used and the **EV3 Golang library**.

* [OpenAPI Generator](https://openapi-generator.tech/)

* [EV3 Golang library](https://github.com/ev3go)

# Setup for students
If you want to code with **Java**, then use this instruction:

* [Setup for Java](./SetupJava.md)

If you want to code with **Python**, then use this instruction:

* [Setup for Python](./SetupJava.md) TODO MATTHEW

# Participation

The current process for development of the Lego Robot is described in [Development Guidelines](https://github.com/PA-arslasel-machitic/EV3-API/blob/master/DEVELOPMENT.md). Development of the next version of the Lego Robot is guided by the [ZHAW School of Engineering](https://www.zhaw.ch/en/engineering/). This group of committers bring their API expertise, incorporate feedback from the community, and expand the group of committers as appropriate. All development activity on the future specification will be performed as features and merged into this branch. Upon release of the future specification, this branch will be merged to <code>main</code>.

The ZHAW holds every 2 weeks web conferences to review to code and open pull requests and discuss open issues related to the evolving Lego Robot. 

The Lego Robot EV3 encourages participation from individuals and companies alike. If you want to participate in the evolution of the Lego Robot, consider taking the following actions:

* Review the current specification. The human-readable markdown file is the source of truth for the specification.
* Review the [development](https://github.com/EV3-OpenAPI/EV3-API/blob/master/DEVELOPMENT.md) process, so you understand how the spec is evolving.
* Check the [issues](https://github.com/EV3-OpenAPI/EV3-API/issues) and [pull requests](https://github.com/EV3-OpenAPI/EV3-API/pulls) to see if someone has already documented your idea or feedback on the specification. You can follow an existing conversation by subscribing to the existing issue or PR.
* Create an issue to describe a new concern. If possible, propose a solution.

Not all feedback can be accommodated and there may be solid arguments for or against a change being appropriate for the specification.


## Setup development

### Prerequisites

Install GO: [Offical website](https://go.dev/dl/), version >= 1.18  
Install either [Docker](https://docs.docker.com/engine/install/) or [Openapi-Generator](https://openapi-generator.tech/docs/installation) for genrating the server code.  
Install goimports tool: `go install golang.org/x/tools/cmd/goimports@latest` for fixing unused imports ing enerated code.

#### Generate server code

Using Docker: `docker run -v ${pwd}:/local --name openapi-generator -u 1000 -w /local openapitools/openapi-generator-cli:latest generate -i openapi/spec.yaml -o internal/gen -g go-server -c openapi/server-config.yml`  
Using Openapi-Generator: `openapi-generator generate -i openapi/spec.yaml -o internal/gen -g go-server -c openapi/server-config.yml`

Clean up unused imports (go refuses to compile if you don't): `goimports -l -w internal/openapi`  
Clean up generated code (because why not?): `gofmt -l -w internal/openapi`

#### Compile server binary

The ldflags option will remove debug code and result in a smaller binary.  
`go build -o server -ldflags="-s -w" EV3-API/cmd`


#### Generate java client code

Using Docker: `docker run -v ${pwd}:/local --name openapi-generator -u 1000 -w /local openapitools/openapi-generator-cli:latest generate -i openapi/spec.yaml -o clients/ev3-java/ev3api -g java -c openapi/java-client-config.yaml`  
Using Openapi-Generator: `openapi-generator generate -i openapi/spec.yaml -o clients/ev3-java/ev3api -g java -c openapi/java-client-config.yaml`


#### Generate python client code

Using Docker: `docker run -v ${pwd}:/local --name openapi-generator -u 1000 -w /local openapitools/openapi-generator-cli:latest generate -i openapi/spec.yaml -o clients/ev3-python -g python -c openapi/python-client-config.yaml`  
Using Openapi-Generator: `openapi-generator generate -i openapi/spec.yaml -o clients/ev3-python -g java -c openapi/python-client-config.yaml`

## Release and Deployment

If a new release is to be created, the new version number must first be entered in the following files:

* openapi/spec.yaml
* clients/ev3-java/build.gradle
* clients/ev3-python/setup.py

After that, a new release can be published via the GitHub. The repository has a GitHub action that automatically creates the Golang, Java and Python artifacts and attaches the release.
Every time the LEGO robot is started, the [GitHub](https://api.github.com/repos/EV3-OpenAPI/EV3-API/releases/latest) is used to check whether the latest release is newer than the currently installed version. If this is the case, the latest GO binary can be downloaded via the Assets API. This replaces the old binary and will be executed in its place in the future.
