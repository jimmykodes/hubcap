BIN = ./var/docker/artifacts/vehicle_maintenance
HEROKU_APP ?= vehicle-logs
DUMP_FILE = ./var/db/dumps/latest.dump

.PHONY: all
all: vendor binary

${BIN}:
	@GOOS=linux GOARCH=amd64 go build -o ${BIN} .

.PHONY: binary
binary: ${BIN}

.PHONY: vendor
vendor:
	@go mod vendor

${DUMP_FILE}:
	@heroku pg:backups:capture --app ${HEROKU_APP}
	@heroku pg:backups:download --app ${HEROKU_APP}
	@mv latest.dump ${DUMP_FILE}

.PHONY dump:
dump: ${DUMP_FILE}
