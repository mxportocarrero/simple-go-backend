# simple-go-backend
this project implements a basic rest api with Go lang.

## Running local server
it requires a local mongoDB server running host at port `"mongodb://localhost:27017"`. Check .env file

    go run cmd/server/main.go

there is also a first glance of integration with docker and docker-compose. Check next commands

    docker-compose build server
    docker-compose up server

## Endpoints
### Create User (Working)
    http://localhost:8000/api/v1/users (POST)

body params:

    {
    "name": "Maxaassf1",
    "phone": "999876153",
    "email": "mxads@google.com"
    }

### Get all Users (Not Working)
    http://localhost:8000/api/v1/users (GET)

query params:

    {
    "limit": int,
    "filterKey": string(name, phone, email),
    "filterValue": string
    }

### Update User (Not Working)
    http://localhost:8000/api/v1/users/{id} (PUT)

body params:

    {
    "name": "Maxaassf1",
    "phone": "999876153",
    "email": "mxads@google.com"
    }

## File Structure

 - cmd/server: main executable dir. Server initializtion goes here
 - internal: hosts most the application packages
   - api: host all versions of the api 
     - utils: various utils functions
     - v1: first release, contains all handlers used by router
    - config: alternative file to hold variables
    - database: database definition and connection
    - model: all entities comming from the DB
 - vendor: auxiliar folder for dependedcy cnsistency
 - docker-compose-yml/Dockerfile integration config for Docker
 - go.mod/go.sum: depency gestor comming from go modules
 - Makefile: compilation file for building application image