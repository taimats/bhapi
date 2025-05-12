#!/usr/bin/env bash
source ${HOME}/bhapi/.env
migrate -path ${HOME}/bhapi/infra/migrations -database "${PSQL_DSN}" -verbose down