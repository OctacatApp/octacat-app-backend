MAIN_FILE = ./src/cmd/main.go
run:
	@ go run ${MAIN_FILE}

run-watch:
	@ if ! [ "$(command -v gow)" == "" ]; then go install github.com/mitranim/gow@latest; fi
	@ gow run ${MAIN_FILE}

run-pretty:
	@ if ! [ "$(command -v gojq)" == "" ]; then go install github.com/itchyny/gojq/cmd/gojq@latest; fi
	@ go run ${MAIN_FILE} | gojq

run-pretty-watch:
	@ if ! [ "$(command -v gow)" == "" ]; then go install github.com/mitranim/gow@latest; fi
	@ if ! [ "$(command -v gojq)" == "" ]; then go install github.com/itchyny/gojq/cmd/gojq@latest; fi
	@ gow -e=go,mod,env run ${MAIN_FILE} | gojq

sqlc-gen:
	@ if ! [ "$(command -v sqlc)" == "" ]; then go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest; fi
	@ sqlc generate --file ./sqlc.yaml
