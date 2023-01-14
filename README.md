# gqlgen-todos

gqlgen-todos is a pet project, to create a to-do list and save records in mongodb. 

uses

- docker compose to create 
  - reverse proxy 
  - graphql server
  - mongodb   
- sonarqube to run static analysis
- github actions to run tests in race mode. 

## GraphQL server in golang

### Steps: 

1.- generate schema from gqlgen-todos/graph/schema.graphqls
```
gqlgen generate
```

2.- generate resolvers: over graph/resolver.go run command
```
go:generate ./...
```

## SonarQube

For local static analysis we use `sonar-scanner` and review results in `sonar GUI`.

Instalation:

* `brew install sonar` - Installs sonarqube 
* `brew install sonar-scanner` - Installs sonar-scanner
* `export SONAR_HOME=/usr/local/Cellar/sonar-scanner/{version}/lib exec`
* `export SONAR=$SONAR_HOME/bin` 
* `export PATH=$SONAR:$PATH` 
* `/usr/local/opt/sonarqube/bin/sonar console` - Starts SonarQube with main paige on `http://localhost:9000/`. Login as admin/admin

    Source: https://sajidrahman.github.io/2018-06-06-install-sonar-on-macosx/

Project Configuration

* Add `sonar-project.properties` file to project root.

```bash
sonar.host.url={sonar-url}

sonar.projectKey={sonar-internalKey}
sonar.projectName={projectNameInGui}
sonar.projectVersion={projectVersionInGui}
sonar.language=go
sonar.sources=.
sonar.sourceEncoding=UTF-8
sonar.tests=.
sonar.test.inclusions=**/*_test.go
sonar.go.coverage.reportPaths=**/coverage.out
```

* Example

```bash
sonar.host.url=http://localhost:9000

sonar.projectKey=gqlgen-todos
sonar.projectName=gqlgen-todos
sonar.projectVersion=1.2
sonar.language=go
sonar.sources=.
sonar.sourceEncoding=UTF-8
sonar.tests=.
sonar.test.inclusions=**/*_test.go
sonar.coverage.exclusions=**/mocks/**
sonar.go.coverage.reportPaths=**/coverage.out
```  

Usage

* Run command `make sonar` - Excecute test, create coverage report and analiye code + report with sonar-scanner.

* Review results in SonarQube Gui - `http://localhost:9000/`
