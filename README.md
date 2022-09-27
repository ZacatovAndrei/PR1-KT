# Kitchen
This repositoty is the home of the kitchen part of the PR laboratory work #1, first being the [Dining hall](https://github.com/zahatikoff/PR1-DH)

### Overview
Kitchen is a web server accepting HTTP POST requests at the `/order` path on port **8087** from the Dining Hall, consisting of orders that have to be cooked by the cooks by using _(or not using)_ some particular kitchen equipment, like an oven,stove, microwave, etc.
After preparing the order it is sent back to the Dining Hall via an HTTP POST request, where it is ranked depending by the time it took to complete.

### Makefile
Most of the commands that will be needed will be accessible via a Makefile as a mean of shortening

- `make build` - makes an executable named Kitchen
- `make docker` - runs a `docker build` command removing an older image an creating a new one called *"zacatov/pr1kitchen"*

### Running the container
Since the containers allow their names to be used as addresses this container has to be launched with the `--name Kitchen` option passed to the `docker run`
