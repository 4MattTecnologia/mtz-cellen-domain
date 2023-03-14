The go test for database_test.go requires previous setup. The machine which tries to run it must have Docker installed.
In order to set up the mock database, please do the following:
1. cd into "mock/"
2. run:
    source createcontainer.sh
This utilizes the docker tool to create a new psql container with specified port, user and password
3. run:
    source setup.sh
This executes the create_db.sql script to properly populate the mock database with the required data for the tests. The command may fail if the containr is not yet healthy and running, in which case one may simply retry until success is achieved.
3. run:
    vars.sh
This will set up the environment variables necessary for a DatabaseFactory() function to create the proper connection.
4. cd back into "integration_tests/" and run:
    go test
5. After executing the tests, run:
    source mock/cleanup.sh
or:
    cd mock/
    source cleanup.sh
So that the mock database is cleaned away from the machine.
