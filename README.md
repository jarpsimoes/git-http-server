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
| Name               | Description                         | Default                                           | Mandatory |
|--------------------|-------------------------------------|---------------------------------------------------|-----------|
| PATH_CLONE         | Set clone path                      | _clone                                            | Yes       |
| PATH_PULL          | Set pull path                       | _pull                                             | Yes       |
| PATH_VERSION       | Set get git commit version path     | _version                                          | Yes       |
| PATH_WEBHOOK       | Set webhook path                    | _hook                                             | Yes       |
| REPO_BRANCH        | Set default branch to clone content | main                                              | Yes       |
| REPO_TARGET_FOLDER | Set folder to clone source          | target-git                                        | Yes       |
| REPO_URL           | Set url as a source origin          | https://github.com/jarpsimoes/git-http-server.git | Yes       |



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

#### Kubernetes implementation
TBD
