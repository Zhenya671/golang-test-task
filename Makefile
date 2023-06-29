start:
	docker-compose up -d --force-recreate --build && docker image prune -f
stop:
	docker-compose down && docker image prune -f
log_app:
	docker logs --tail 50 --follow --timestamps rest_api
log_db:
	docker logs --tail 50 --follow --timestamps db_postgres
