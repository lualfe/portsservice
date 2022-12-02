# Ports Service
Ports service is a service that executes a job to extract ports data 
from a json file. 
## Help
This project has a Makefile. To get help, just run `make` and the
documentation for the commands will show.
## Execute
To execute the docker compose of this project, just run `make run-app` and
it will download the necessary images and run.
## How it works
The code reads a file that is saved in the container and saves its content
to a Redis db. The file is `ports.json` in `internal/usecase/testdata`.
This was done for the sake of simplicity. The next step would be reading
the file from an API (a S3 bucket maybe).
## Validation
To check if the data was saved to Redis, you can use a GUI like
[this one](https://github.com/ekvedaras/redis-gui).
## Next steps
- The json file is probably located in some storage or an API. In this case  
it'd be necessary to change the implementation to add http requests. The 
read logic would remain the same. For local development, adding a file server
to docker compose would help;
- Implementing a more robust database layer. Redis is just a temporary
solution, so it's important to implement something more adequate like
Postgres, etc.
- This service would potentially be part of a microservices architecture,
so I'd assume that there'd be other services interested in new ports or
updated ports. Generating events for these would be a great addition 
(it's necessary to consider the required consistency).
- CI/CD - implementing a pipeline for the service for validating tests,
static analysis and deploy.