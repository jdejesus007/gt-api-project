# gt-api-project
GT Go API

## Swagger
View OpenAPI swagger docs at http://{{hostname}}:3000/docs/index.html#/
local: http://localhost:3000/docs/index.html

## Setup local db manually
- install postgresql via homebrew and start services
  brew install postgresql
- create root user with example pwd
- create gt db (createdb gt)
- GRANT ALL PRIVILEGES ON DATABASE gt TO root;

  1. psql
  2. \du to show current users
  3. CREATE ROLE root WITH LOGIN SUPERUSER PASSWORD 'example';
  4. \l show all dbs
  5. CREATE DATABASE gt;


## Run api
make dev
make run-dev

## Run Tests
make dev
make run-tests
