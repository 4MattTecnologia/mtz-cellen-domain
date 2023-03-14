#!/usr/bin/env bash

# populating database
export PGPASSWORD=pwd
psql -h localhost -p 8080 -U postgres -f ./create_db.sql -q
