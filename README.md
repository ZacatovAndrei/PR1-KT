# Kitchen
This repositoty is the home of the kitchen part of the PR laboratory work #1, first being the [Dining hall](https://github.com/zahatikoff/PR1-DH)

### Overview

Kitchen is a web server accepting HTTP POST requests at the `/order` path on port **8087** from the Dining Hall, consisting of orders that have to be cooked by the cooks by using _(or not using)_ some particular kitchen equipment, like an oven,stove, microwave, etc.
After preparing the order it is sent back to the Dining Hall via an HTTP POST request, where it is ranked depending by the time it took to complete.
    
### Available commands
Some commands are available from a makefile as a simplification:
- `make build` - builds an executable `Kitchen`
- `make docker` - creates a docker container with the name zacatov/pr1kitchen
### Running a container:
Containers should be run on the same docker network that one can create via a `docker network create`,
since it allows using the names of the containers as their names as their IP addresses,
so the sample command would look like
`docker run --name Kitchen --network restaurant zacatov/pr1kitchen`