# Overview

Simple Backend Application for the [Octacat App Chatting](https://github.com/OctacatApp/octacat-app-frontend.git) application made with Go programming language with a little Coffee and Clean Architecture

# Tech stack

- [sqlc](https://github.com/sqlc-dev/sqlc): sqlc generates type-safe code from SQL
- [gqlgen](https://github.com/99designs/gqlgen.git): gqlgen is a Go library for building GraphQL servers without any fuss.
- [go-argon2](https://github.com/irdaislakhuafa/go-argon2.git): simple golang code to implement argon2 hashing with standard format

# How to run this app

## Run With Docker Compose

Run this app with Docker Compose and Docker will magically prepare the environment to run the application

```bash
$ docker-compose up -d
```

## Run with local environment

If you want to run this app in local environment you can use command below

```bash
$ go run src/cmd/main.go -env local
```

the option `-env` will search environment configuration in `etc/cfg/*`

## Deployed App

You can access deployed app [here](https://octacat-app-backend.fly.dev/). If you want to use this API send your GraphQL query to `https://octacat-app-backend.fly.dev/query`
