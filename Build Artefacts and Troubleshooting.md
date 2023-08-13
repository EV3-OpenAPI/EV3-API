# Build Artefacts Commands

## EV3 Server Tipps

### Connect to the EV3

```bash 
ssh robot@IP-ADDRESS-VISIBLE-ON-DISPLAY (Passwort `maker`)
```

### Troubleshooting

```bash 
sudo systemctl stop ev3api-server
sudo systemctl start ev3api-server
```

## Build Artefacts

### Install OpenAPI Tools

```bash
npm install @openapitools/openapi-generator-cli -g
```

### Go Server

```bash
openapi-generator-cli version-manager set latest
openapi-generator-cli generate -i openapi/spec.yaml -g go-server -o internal/gen -c openapi/server-config.yml

go get EV3-API
go install golang.org/x/tools/cmd/goimports@latest

goimports -l -w internal/gen/openapi
gofmt -l -w internal/gen/openapi

go build -v -o ev3api-server -ldflags="-s -w" EV3-API/cmd
```

### Java Client

```bash
sdk use java 17.0.7-tem

export JAVA_POST_PROCESS_FILE="/usr/local/bin/clang-format -i"

openapi-generator-cli version-manager set latest
openapi-generator-cli generate -i openapi/spec.yaml -o clients/ev3-java/ev3api -g java -c openapi/java-client-config.yaml

cd ./clients/ev3-java
./gradlew shadowJar
```

### Python Client

```bash
sdk use java 11.0.20-tem

# Downgrade for Python needed, since in never version files are missing
openapi-generator-cli version-manager set 6.1.0
openapi-generator-cli generate -i openapi/spec.yaml -o clients/ev3-python -g python -c openapi/python-client-config.yaml

cd ./clients/ev3-python

python -m pip install --upgrade pip
pip install build

python -m build
```
