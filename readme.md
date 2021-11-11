## Run project
`go run cmd/app/main.go`

## Run API documentation:
`godoc -goroot /home/path/to/the/project`

find the documentation here: `http://localhost:6060/pkg/shorten-link/`

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
### ToDos:
- [ ] tests
- [ ] migrations
- [ ] dockerize
- [ ] grpc

