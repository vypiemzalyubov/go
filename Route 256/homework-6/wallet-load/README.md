#  Wallet

## Run app
```shell
 # create db in docker-compose
 make dc-db-up
 # migration up
 make db-up
 # run app
 make run
 ```

## Run app in docker-compose
 ```shell
 make dc-app-up
 ```

## Run tests
```shell
# unit-tests
make test
# unit-tests with coverage
make test-cov
# unit-test and CBR tests
make test-cbr
# unit-tests and tests CBR with coverage
make test-cov-cbr
```

## Generate gRPC server and client
```shell
make deps
make generate
```

## Links
 - [swagger](http://localhost:8080/swagger/)
 - [graylog](http://localhost:9001/search)
 - [grafana](http://localhost:3000/)
 - [victoria](http://localhost:8428)