#!/bin/bash

# https://dev.to/nietzscheson/multiples-postgres-databases-in-one-service-with-docker-compose-4fdf

if [ ! -d ./env ]; then
    mkdir env
fi

if [ ! -f ./env/postgres.env ]; then
    touch env/postgres.env
    echo "POSTGRES_PORT=5432" >> ./env/postgres.env
    echo "POSTGRES_USER=postgres" >> ./env/postgres.env
    echo "POSTGRES_PASSWORD=postgres" >> ./env/postgres.env
    echo "POSTGRES_MULTIPLE_DATABASES=todos, todos_test" >> ./env/postgres.env
fi

if [ ! -f ./env/app.env ]; then
    touch env/app.env
    echo "DB_ADDR=postgres://postgres:postgres@localhost:5432/todos?sslmode=disable" >> ./env/app.env
    echo "TEST_DB_ADDR=postgres://postgres:postgres@localhost:5432/todos_test?sslmode=disable" >> ./env/app.env
fi