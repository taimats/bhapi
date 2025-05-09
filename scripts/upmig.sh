#!/bin/bash
source ../.env
migrate -path ./migrations -database "${PSQL_DSN}" -verbose up