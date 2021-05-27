BIN = ./var/docker/artifacts/vehicle_maintenance

.PHONY: all
all: vendor binary

${BIN}:
	@GOOS=linux GOARCH=amd64 go build -o ${BIN} .

.PHONY: binary
binary: ${BIN}

.PHONY: vendor
vendor:
	@go mod vendor
