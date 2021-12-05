build-redis:
	docker-compose build app-redis db-redis

up-with-redis:
	docker-compose up -d app-redis

start-with-redis: up-with-redis

restart-redis: stop start-with-redis

manual-test-redis: 
	curl --silent -X "POST" -i "http://localhost:18000/shorten" -d "{\"initial\":\"$(ARGS)\"}" -H 'Content-Type: application/json'| tail -1 | xargs -I {} echo "http://localhost:18000/{${cat}}"