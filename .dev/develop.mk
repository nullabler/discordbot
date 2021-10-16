##@ Develop:

bot: ## start shell in backend
	docker-compose exec bot sh

mongo: ## connect for db
	docker-compose exec mongo mongosh
