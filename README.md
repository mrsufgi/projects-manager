# PROJECT List API (WIP)

This is still pretty much WIP, needs more tests (missing integration tests), fixing live tests (they work but will yield inconsistent results, becuase they live).

- [ ] protobuf support
- [ ] fix docker-compose for integration-test
- [ ] fix migration for docker-compose (wait for postgres script)
- [ ] add more API test (handler)
- [ ] add more unit tests cases (error path)
- [ ] cleanup makefile
- [ ] fix lint issues
- [ ] improve docs

### Prerequisites

1. Docker
2. Postgres (for tests)
3. Go 1.13+

## Getting Started

If you just want to run it, I've created a docker-compose that will install everything for you and start the server.

To start, run:

```bash
make start
```

it will create everything you need and run it.
the server is running on: `localhost:4000`
caveat: you will need to run migrations (actually rerun the migration container):

```bash
docker start migrate
```

to stop and cleanup:

```bash
make stop
```

in fact, this setup is running in development mode and if you want to devlop the container will refresh itself and compile the bin again.
(caveat: sometimes issues with too many concurrent rebuilds and used port)

_NOTE:_
`make stop` runs `docker-compose down`, and it will <b>remove</b> the postgres container. the setup is not meant to keep the data of this container.
if you want to test "restarts" you can just restart `projects` container :)

## Using the API

The easiest way would be importing the POSTMAN json file that describe the api.
Otherwise the api is CRUD using REST:

(assuming you ran it using `make start` the host is <b>http://localhost:4000</b>)

```http
GET /projects/?q=
GET /projects/:id
PUT /projects/:id {done: bool, name: string, details: string}
POST /projects {done: bool, name: string, details: string}
DELETE /projects/:id
```

## Local Devlopment and Tests

if you want to devlop locally or run tests. start with

```bash
make install
```

### Unit tests

```bash
make unit-tests
```

(Live test are ignored)

### Integration tests

PROJECT
