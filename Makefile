start:
	@docker-compose up --build

cleanVolumes:
	@docker-compose down
	@docker volume rm artforintrovert_testex_api artforintrovert_testex_mongodb-data
