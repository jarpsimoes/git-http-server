# GIT-HttpServer

This is a simple HTTP server to provide "NoOps" to the frontend applications.
The content must be provided from a git repository. To synchronize content, exist
two approach's. 
- Pull Job: can be defined an interval to check updates on origin repository, and every cycle will be run ```git pull``` command to update content
- Web Hook: can be defined an endpoint to be consumed by CI/CD tools or GIT Server, every time of hook has called, the sync command will be performed

## Configuration

TBD

## Easy implementation

TBD