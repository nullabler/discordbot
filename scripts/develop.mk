##@ Develop:

sh: ## start shell in backend
	docker-compose --env-file configs/.env -f deploy/docker-compose.yml exec discordbot sh
