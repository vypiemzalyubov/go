db-create-migration: NAME=$NAME
db-create-migration:
	$(LOCAL_BIN)/goose -dir ./migrations postgres $(PG_DSN) create "${NAME}" sql

db-status: $(eval $(call db_env))
	$(LOCAL_BIN)/goose -dir ./migrations postgres $(PG_DSN) status

db-up: $(eval $(call db_env))
	$(LOCAL_BIN)/goose -dir ./migrations postgres $(PG_DSN) up

db-down: $(eval $(call db_env))
	$(LOCAL_BIN)/goose -dir ./migrations postgres $(PG_DSN) down

db-up-ci: $(eval $(call db_env))
	$(LOCAL_BIN)/goose -dir ./migrations postgres "$(PG_DSN_CI)" up
