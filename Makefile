run:
	docker-compose build
	docker-compose up -d
	docker exec -ti tangelo_challenge_app_1 bash -c 'go run main.go'

test:
	docker-compose build
	docker-compose up -d
	docker exec -ti tangelo_challenge_app_1 bash -c 'go test'

testApi:
	docker-compose build
	docker-compose up -d
	docker exec -ti tangelo_challenge_app_1 bash -c 'go test -v -run Test_requestAPI'

clean:
	docker-compose down -v