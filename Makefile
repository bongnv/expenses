mysql:
	docker run --name mysql8 -p 3306:3306  -e MYSQL_ROOT_PASSWORD=secret -d mysql:8

server:
	cd ./backend/cmd/server && go run main.go

.PHONY: mysql
