DC_APP:=docker-compose -f compose.yml -p wallet

###########
app-up:
	$(DC_APP) up -d
	sleep 1
	make app-status

app-down:
	$(DC_APP) down --remove-orphans -v -t0
	make app-status

app-reup: app-down app-up

###########

app-status:
	sleep 1
	docker-compose -p wallet ps -a

dc-logs:
	docker-compose -p wallet logs wallet

###########

colima:
	colima stop
	colima start --cpu 2 --memory 6 --disk 50
	colima list
	docker ps -a
