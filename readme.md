## Run project
`go run cmd/app/main.go`

## Run API documentation:
`godoc -goroot /home/path/to/the/project`

find the documentation here: `http://localhost:6060/pkg/shorten-link/`

## Run Tests
`go test -v shorten-link/tests`

## Example of running migrations with [golang-migrate](https://github.com/golang-migrate/migrate) tool
`migrate -database postgres://dev:dev@localhost:5432/shorter -path pkg/db/migrations up 1`

## Plan:
### Done
- [X] skeleton
- [X] receive and respond with full link
- [X] research hash
- [X] respond with shortened link
- [X] add db
- [X] check if link has already been shortened
- [X] handle errors: non-existing shortened link
- [X] handle errors: empty original link
- [X] handle errors: check valid url
- [X] godoc documentation
- [X] simple test
### ToDos:
- [ ] migrations
- [ ] dockerize
- [ ] more complicated tests with mocking or creating/killing specific db
- [ ] grpc

