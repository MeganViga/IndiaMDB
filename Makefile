migratecreate:
	migrate -ext sql -dir db/migration -seq init_schema
migrateup:
	 migrate -path db/migration -database "postgres://root:secret@localhost:5432/indiamdb?sslmode=disable" -verbose up

migratedown:
	 migrate -path db/migration -database "postgres://root:secret@localhost:5432/indiamdb?sslmode=disable" -verbose down
createdb:
	docker exec -it postgres createdb indiamdb
dropdb:
	docker exec -it postgres dropdb indiamdb
accessdb:
	docker exec -it postgres psql  indiamdb

.PHONY: accessdb migrateup migratedown migratecreate createdb dropdb