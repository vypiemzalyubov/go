DC_DB:=docker-compose -f docker-compose.db.yaml -p wallet
DC_APP:=docker-compose -f docker-compose.app.yaml -f docker-compose.db.yaml -p wallet
DC_APP:=docker-compose -f docker-compose.app.yaml -f docker-compose.db.yaml -f docker-compose.observability.yaml -p wallet

dc-db-up:
	$(DC_DB) up -d

dc-db-down:
	$(DC_DB) down --remove-orphans -v -t0

dc-db-reup:
	make dc-db-down
	make dc-db-up

###########
dc-app-build: dc-app-down
	$(DC_APP) up -d --no-deps --build --force-recreate wallet

dc-app-up:
	$(DC_APP) up -d
	sleep 1
	make dc-app-status

dc-app-down:
	$(DC_APP) down --remove-orphans -v -t0
	make dc-app-status

dc-app-reup: dc-app-down dc-app-up

dc-app-rebuild-reup: dc-app-build dc-app-up

###########

dc-app-status:
	sleep 1
	docker-compose -p wallet ps -a

dc-logs:
	docker-compose -p wallet logs -f wallet

###########

colima:
	colima stop
	colima start --cpu 2 --memory 6 --disk 50
	colima list
	docker ps -a
