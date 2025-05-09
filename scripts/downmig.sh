#!/usr/bin/env bash
source ../.env
migrate -path ../infra/migrations -database "${PSQL_DSN}" -verbose down