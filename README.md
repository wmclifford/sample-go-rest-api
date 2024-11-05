# sample-go-rest-api

First-time `golang` attempt; a sample Dockerized REST API written in `go`.

* Exposes a couple of simple REST endpoints for creation and retrieval.
* Uses a PostgreSQL database backend.
* Includes tests.
* Includes `Dockerfile` that builds the application and installs into a Docker image.
* Includes a Docker compose example which spins up a PostgreSQL container alongside the application container.
  * Database is exposed on `localhost:3306` for debugging.
  * Application is exposed on `localhost:8080`.

May or may not extend this, but sharing for anyone else who has decided to look into `golang` as an alternative to Java
and Python for developing microservices. Comments and suggestions are welcome. This is literally the first consequential
code that I have written in `go`, having gone through the online "Getting Started" tutorial and then jumping both feet
first into this endeavor.
