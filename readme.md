## Run project
`go run cmd/app/main.go`

## Run API documentation:
`godoc -goroot /home/path/to/the/project`

find the documentation here: `http://localhost:6060/pkg/shorten-link/`

## Run Tests
`go test -v shorten-link/tests`

## Example of running migrations with [golang-migrate](https://github.com/golang-migrate/migrate) tool
`migrate -database postgres://dev:dev@localhost:5432/shorter -path pkg/db/migrations up 1`

## Run app in Docker (DB is not included in image)
```
docker build --tag docker-app . --build-arg db=p
docker run -p 8080:8080 docker-app

```

## Run with docker-compose
```
docker-compose -f docker-compose.yml up --build app-postgres
docker-compose -f docker-compose.yml up -d --force-recreate db-postgres

dokcer ps
docker exec -it <container_id> bash
docker-compose down --volumes ; docker-compose up -d

docker-compose up -d app-postgres
```

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
