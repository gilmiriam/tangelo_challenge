run: ## Start the server application
	docker-compose build
	docker-compose up -d
	docker exec -ti tangelo_challenge_app_1 bash -c 'nohup go run main.go &> /tmp/output & sleep 1'

test: ## Run the test suite
	docker exec -ti tangelo_challenge_app_1 bash -c 'go test ./...'

testApi: ## Run only the test corresponding to the response from the API
	docker exec -ti tangelo_challenge_app_1 bash -c 'go test -v -run Test_requestAPI'

clean: ## Clean up build artifacts
	docker-compose down -v