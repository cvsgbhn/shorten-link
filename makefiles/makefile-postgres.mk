build-postgres:
	docker-compose build app-postgres db-postgres

up-with-postgres:
	docker-compose up -d app-postgres

migrations: docker-compose up -d migrate

start-with-postgres: up-with-postgres migrations

restart-postgres: stop start-with-postgres

manual-test-postgres: 
	curl --silent -X "POST" -i "http://localhost:18001/shorten" -d "{\"initial\":\"$(ARGS)\"}" -H 'Content-Type: application/json' | tail -1 | xargs -I {} echo "http://localhost:18001/{${cat}}"