# Project Detail

- This is an Expense tracking system using Golang with REST API
- This tracking system accepts only 4 operations which are
  - Create
  - Query a single row
  - Query all rows
  - Update a specific row

## How to start an application

- go run server.go

## How to run unit-test

- go test --tags=unit -v ./...

## How to run integration-test using docker-compose testing sandbox

- docker-compose -f docker-compose.yml up --build --abort-on-container-exit --exit-code-from expense_tracking

## How to build a docker image

- docker build -t assessment .

## How to run a docker image

- docker run \
  -e PORT=:2565 \
  -e DATABASE_URL=postgres://mjjvhixr:ZkGSsm9jMBIY3x37X8A_hFUBPTIlgw-g@john.db.elephantsql.com/mjjvhixr \
  -p 2565:2565 \
  assessment
