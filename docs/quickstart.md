# Quickstart

## Host Machine Setup

In order to run horseville, your machine must have: TODO: minimum requirements

and the following installed:

- Docker
- Docker Compose
- Make

## Checkout

Check the repo out somewhere. This can be on $GOPATH or not depending on your needs. Open it up in a shell to get the basic setup running:

```console
$ make run
```

This will build the docker image and then start its dependencies in Docker. The Postgres database will have migrations applied to it when the app boots and connects to postgres. 

This exposes TCP port 8084 for HTTP traffic. Open http://127.0.0.1:8084 to get started.

## Cookbook

### Update dependencies

Rebuild the docker images so everything is in sync:

```console
$ make docker
```

Enter a shell with your source code bind-mounted over where the copied code would otherwise be:

```console
$ docker run --rm -it -e TERM=$TERM -v $(pwd):/root/go/github.com/horseville/horseville horseville/core
(ctr)#
```

Now rerun dep

```console
(ctr)# make tools dep
```

### Re-generate protobufs

Rebuild the docker images so everything is in sync:

```console
$ make docker
```

Enter a shell with your source code bind-mounted over where the copied code would otherwise be:

```console
$ docker run --rm -it -e TERM=$TERM -v $(pwd):/root/go/github.com/horseville/horseville horseville/core
(ctr)#
```

Now rerun the generators:

```console
(ctr)# make generate
```
