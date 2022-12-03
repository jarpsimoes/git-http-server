# GIT-HttpServer

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/b0fde431e29c4e3ba47560a973279fef)](https://www.codacy.com/gh/jarpsimoes/git-http-server/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=jarpsimoes/git-http-server&amp;utm_campaign=Badge_Grade)
[![codecov](https://codecov.io/gh/jarpsimoes/git-http-server/branch/main/graph/badge.svg?token=CCRRRCYLM1)](https://codecov.io/gh/jarpsimoes/git-http-server)

This is a simple HTTP server to provide "NoOps" to the frontend applications.
The content must be provided from a git repository. 

Every times an application is started, the configured repository is cloned on 
selected branch (in variable REPO_BRANCH) and can be pulled new version on defined PATH_PULL.

## Authentication Methods

The GIT-HttpServer only support basic authentication on repositories by protocol HTTPS

## Configuration

### Environment Variables
| Name                      | Description                                              | Default                                           | Mandatory |
|---------------------------|----------------------------------------------------------|---------------------------------------------------|-----------|
| PATH_CLONE                | Set clone path                                           | _clone                                            | Yes       |
| PATH_PULL                 | Set pull path                                            | _pull                                             | Yes       |
| PATH_VERSION              | Set get git commit version path                          | _version                                          | Yes       |
| PATH_WEBHOOK              | Set webhook path                                         | _hook                                             | Yes       |
| PATH_HEALTH               | Set health check path                                    | _health                                           | Yes       |
| REPO_BRANCH               | Set default branch to clone content                      | main                                              | No        |
| REPO_TARGET_FOLDER        | Set folder to clone source                               | target-git                                        | No        |
| REPO_URL                  | Set url as a source origin                               | https://github.com/jarpsimoes/git-http-server.git | No        |
| REPO_USERNAME             | Set username or token identifier to basic authentication | N/D                                               | No        |
| REPO_PASSWORD             | Set password or token to basic authentication            | N/D                                               | No        |
| HTTP_PORT                 | Set port to expose content                               | 8081                                              | Yes       |
| GHS_CUSTOM_PATH_<path>    | Custom path to work as a proxy server                    | N/D                                               | No        |
| GHS_CUSTOM_REWRITE_<path> | Set to remove from proxy request base path               | N/D                                               | No        |
| FOLDER_ROOT               | Select root folder inside cloned repository              | $REPO_TARGET_FOLDER/                              | No        |


## Implementation

### Simple implementation

The most simple implementation of git-http-server
```shell
$ docker run \ 
    -p 8081:8081 \
    -e REPO_URL=[URL REPOSITORY] \
    -e REPO_BRANCH=[DEFAULT BRANCH] \
    ghcr.io/jarpsimoes/git_http_server
```
Test content
```shell
$ curl http://localhost:8080
```

Update Repository
````shell
$ curl http://localhost:8080/_pull
````

### Implementation with Basic Authentication

The implementation with basic authentication can be used with username and password method, or PAT(Personal Access Token) method (like Gitlab Token).
For the PAT approach the username is Token Identifier and Password is token.

```shell
$ docker run \ 
    -p 8081:8081 \
    -e REPO_URL=[URL REPOSITORY] \
    -e REPO_BRANCH=[DEFAULT BRANCH] \
    -e REPO_USERNAME=[Token Identifier or Username] \
    -e REPO_PASSWORD=[Password or Token]
    ghcr.io/jarpsimoes/git_http_server
```

### Implementation with Proxy redirect

Git Http Server support routes to redirect for another URL. Should be defined path to redirect and target. 
The environment variables must be defined as bellow:

- GHS_CUSTOM_PATH_example/path/redirect: "https://example.org"
- GHS_CUSTOM_REWRITE_example/path/redirect: true # If you need remove the "example/path/redirect" from the proxy request must be defined:

```shell
$ docker run \ 
    -p 8081:8081 \
    -e REPO_URL=[URL REPOSITORY] \
    -e REPO_BRANCH=[DEFAULT BRANCH] \
    -e REPO_USERNAME=[Token Identifier or Username] \
    -e REPO_PASSWORD=[Password or Token] \
    -e GHS_CUSTOM_PATH_google/redirect=https://www.google.pt
    -e GHS_CUSTOM_REWRITE_google/redirect=true
    ghcr.io/jarpsimoes/git_http_server
```

Test content
```shell
$ curl http://localhost:8080
```

Update Repository
````shell
$ curl http://localhost:8080/_health
````

### Kubernetes implementation

The git-http-server is Kubernetes Ready, can be deployed as simple deployment. See bellow an example of kubernetes deployment.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: simple-content-deployment
spec:
  replicas: 1
  strategy:
    type: RollingUpdate
  selector:
    matchLabels:
      app: git-http-server-example-1
  template:
    metadata:
      labels:
        app: git-http-server-example-1
    spec:
      containers:
        - ports:
            - name: http
              containerPort: 8081
              protocol: TCP
          resources:
            limits: {}
            requests: {}
          env:
            - name: REPO_URL
              value: [REPOSITORY_URL]
          name: git-http-server
          image: ghcr.io/jarpsimoes/git_http_server:v0.0.1-beta
          imagePullPolicy: Always
          livenessProbe:
            httpGet:
              port: 8081
              path: /_health
            failureThreshold: 3
            periodSeconds: 10
          startupProbe:
            httpGet:
              port: 8081
              path: /_health
            failureThreshold: 3
            periodSeconds: 10
```
