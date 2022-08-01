# GIT-HttpServer

This is a simple HTTP server to provide "NoOps" to the frontend applications.
The content must be provided from a git repository. 

Every times its application started, the configured repository is cloned on selected 
branch defined in variable REPO_BRANCH and can be pulled new version on defined PATH_PULL.

- Git Authentication [WIP]
- Cron job to preform content update [WIP]
- 

## Configuration

#### Environment Variables
| Name               | Description                                              | Default                                           | Mandatory |
|--------------------|----------------------------------------------------------|---------------------------------------------------|-----------|
| PATH_CLONE         | Set clone path                                           | _clone                                            | Yes       |
| PATH_PULL          | Set pull path                                            | _pull                                             | Yes       |
| PATH_VERSION       | Set get git commit version path                          | _version                                          | Yes       |
| PATH_WEBHOOK       | Set webhook path                                         | _hook                                             | Yes       |
| REPO_BRANCH        | Set default branch to clone content                      | main                                              | Yes       |
| REPO_TARGET_FOLDER | Set folder to clone source                               | target-git                                        | Yes       |
| REPO_URL           | Set url as a source origin                               | https://github.com/jarpsimoes/git-http-server.git | Yes       |
| HTTP_PORT          | Set port to expose content                               | 8081                                              | Yes       |
| REPO_USERNAME      | Set username or token identifier to basic authentication | N/D                                               | No        |
| REPO_PASSWORD      | Set password or token to basic authentication            | N/D                                               | No        |



## Implementation

#### Simple implementation

The simple implementation is based with minimal configuration
```shell
$ docker run \ 
    -p 8080:8080 \
    -e REPO_URL=[URL REPOSITORY] \
    -e REPO_BRANCH=[DEFAULT BRANCH] \
    jarpsimoes/git_http_server
```
Test content
```shell
$ curl http://localhost:8080
```

Update Repository
````shell
$ curl http://localhost:8080/_pull
````

#### Implementation with Basic Authentication

The implementation with basic authentication can be used with username and password method, or PAT(Personal Access Token) method (like Gitlab Token).
For the PAT approach the username is Token Identifier and Password is token.

```shell
$ docker run \ 
    -p 8080:8080 \
    -e REPO_URL=[URL REPOSITORY] \
    -e REPO_BRANCH=[DEFAULT BRANCH] \
    -e REPO_USERNAME=[Token Identifier or Username] \
    -e REPO_PASSWORD=[Password or Token]
    jarpsimoes/git_http_server
```

Test content
```shell
$ curl http://localhost:8080
```

Update Repository
````shell
$ curl http://localhost:8080/_pull
````

#### Kubernetes implementation
TBD
