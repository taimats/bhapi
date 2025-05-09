#!/usr/bin/env bash
source ../.env
migrate -path ./migrations -database "${PSQL_DSN}" -verbose down