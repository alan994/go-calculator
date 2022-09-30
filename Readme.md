# About the project

This project is build for interviewing purposes.

# Prerequisites
* [Docker](https://docs.docker.com/get-docker/)

# Instructions on how to run the project

This project has dependency to Redis so most practical solution is to run everything with `docker-compose`.

* Position command line to root directory
* run `docker-compose up -d` command

To test endpoints run `curl` or `wget` with url like this `http://localhost:8080/add?x=2&y=5` or use [RestClient](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) Visual Studio Code extension and run `api-endpoints.http` file located in `/src/test` directory