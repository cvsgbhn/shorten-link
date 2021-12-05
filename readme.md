## Run project with postgres
```
make build-postgres
make start-with-postgres
make migrations
```

## Run project with redis
```
make build-redis
make start-with-redis
```

## Examples of manual testing
### Postgres:
```
make -s manual-test-postgres https://www.google.com
```
### Redis:
```
make -s manual-test-redis https://www.google.com
```
### Expected output:
```
http://localhost:18001/z3p3O2Ma21
```

## Shut down all
```
make stop
```