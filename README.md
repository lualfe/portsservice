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