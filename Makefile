all: generate start-local-database run

generate:
	go generate github.com/eveld/graphql/postgres/
	go run scripts/gqlgen.go -v

run:
	GCLOUD_PROJECT=instruqt-dev \
	DB_NAME=test \
	DB_USER=test \
	DB_PASSWORD=123456 \
	go run main.go

start-local-database:
	docker start test-db &>/dev/null || \
	docker run -d -p 5432:5432 \
		-e POSTGRES_USER=test \
		-e POSTGRES_DB=test \
		-e POSTGRES_PASSWORD=123456 \
		--name test-db \
		postgres

stop-local-database:
	docker stop test-db