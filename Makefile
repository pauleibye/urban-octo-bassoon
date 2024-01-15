run-app:
	go run main.go

start-pg:
	docker compose -f build/docker-compose.yml up --build -d postgres

stop-pg:
	docker compose -f build/docker-compose.yml down postgres

docker-build: 
	docker build --tag pauleibye/urban-octo-bassoon .

docker-run:
	docker run -p 8080:8080 pauleibye/urban-octo-bassoon

docker-push:
	docker push pauleibye/urban-octo-bassoon

docker-delete:
	docker rmi pauleibye/urban-octo-bassoon -f