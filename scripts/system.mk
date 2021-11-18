##@ System:

up: ## start project
	docker-compose --env-file configs/.env -f deploy/prod/docker-compose.yml up -d

down: ## stop project
	docker-compose --env-file configs/.env -f deploy/prod/docker-compose.yml down

watch: ## watch project
	docker-compose --env-file configs/.env -f deploy/dev/docker-compose.yml up 

state: ## show state
	docker-compose --env-file configs/.env -f deploy/prod/docker-compose.yml ps

logs: ## show last 100 lines of logs
	docker-compose --env-file configs/.env -f deploy/prod/docker-compose.yml logs --tail=100 $(ARGS)


