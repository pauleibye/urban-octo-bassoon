run:
	go run main.go

docker-build: 
	docker build --tag pauleibye/urban-octo-bassoon .

docker-run:
	docker run -p 8080:8080 pauleibye/urban-octo-bassoon

docker-push:
	docker push pauleibye/urban-octo-bassoon

docker-delete:
	docker rmi pauleibye/urban-octo-bassoon -f