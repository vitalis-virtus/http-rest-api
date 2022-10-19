RESTful API for the book author entity. Entity data is stored in MySQL.

App server is listening on port 8080
Db server is listening on port 3306


Commands to run:

- make build - build main.go

- make run - run main.go

- make test - run test for app

- make lint - run golangci lint

- make lint-fast - run golangci lint fast

- make docker-build - run container in development mode

- make docker-compose - spin up the project

- make docker-stop - stop running containers

- make docker-rm - stop and remove running containers


Routes:

/author/ - get all authors via GET request method
/author/ - create new author via POST request method
/author/{id} - get author by id via GET request method
/author/{id} - update author by id via PUT request method
/author/{id} - delete author by id via DELETE request method