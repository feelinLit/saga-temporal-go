BINDIR=${CURDIR}/bin
MIGRATIONS_DIR=./db/migrations/
DB_HOST=localhost
DB_PORT=5432
LOCAL_DB_DSN=postgresql://temporal:temporal@${DB_HOST}:${DB_PORT}/orders?sslmode=disable&connect_timeout=10

bindir:
	mkdir -p ${BINDIR}

.PHONY: install-bins
install-bins: bindir
    GOBIN=$(BINDIR) go install github.com/pressly/goose/v3/cmd/goose@v3.24.1

run-migrations:
	${BINDIR}/goose -dir ${MIGRATIONS_DIR} postgres "${LOCAL_DB_DSN}" up

rollback-migration:
	${BINDIR}/goose -dir ${MIGRATIONS_DIR} postgres "${LOCAL_DB_DSN}" down

create-migration:
	${BINDIR}/goose -dir ${MIGRATIONS_DIR} create $(n) sql
