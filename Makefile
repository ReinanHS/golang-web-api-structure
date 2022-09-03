server:
	go run cmd/server/main.go
up:
	docker compose up
down:
	docker compose down -v
mysql:
	docker exec -it golang-web-api-mysql mysql -u admin -p"root" database
artisan:
	@go run cmd/tools/artisan.go