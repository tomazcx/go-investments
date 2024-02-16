 help: ## Prints available commands
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[.a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

MIGRATION_NAME =
migrate.new: ## Create new migration
	migrate create -ext=sql -dir=./internal/infra/database/migrations $(MIGRATION_NAME)

migrate.up:  ## Run migrations up
	migrate -source file:./internal/infra/db/migrations/ -database "postgres://root:root@db:5432/go-investments" up

migrate.down: ## Run migrations down 1
	migrate -source file:./internal/infra/db/migrations/ -database "postgres://root:root@db:5432/go-investments" down 1
