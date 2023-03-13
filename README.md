## Auth Service

### Description

This services handles anything security/auth related. This service will determine what role a user holds, does that user have access to given resources, has their session expired or not, etc.

### Identities

Identities will be used to store information regarding login details and roles. An identity consists of an email address and password. In the future, what role that identity holds.

### Running the service

1. Make init will run `go mod tidy` to clean up and unused packages from the go.mod, and add any that are used to it.
```shell 
make init
```

2. Make platform will spin up any required services for `auth` to run. This includes databases, redis instances, rabbitmq, etc.
```shell
make platform
```

3. Make start and make run will start the service. Make start will spin up the service in a docker container, while make run will simply run the go service locally.

```shell
make start
``` 
or 
```shell 
make run
```

4. Make test (has not been set up yet) will run any go tests that have been written.
``` shell
make test
```
