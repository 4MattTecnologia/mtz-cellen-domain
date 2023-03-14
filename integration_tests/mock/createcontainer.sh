#!/usr/bin/env bash

# creating psql container
docker run --name mockdb -p 8080:5432 -e POSTGRES_PASSWORD=pwd -d postgres
