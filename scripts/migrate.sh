#!/bin/sh

export MIGRATION_DIR=./migrations

if [ "$1" = "--localhost" ]; then
    export POSTGRES_HOST="127.0.0.1"
    export POSTGRES_PORT="5432"
    export POSTGRES_DB="easy_crypto_wallet_service"
    export POSTGRES_USER="easy_crypto_wallet_service"
    export POSTGRES_PASSWORD="password"
fi

goose -dir ${MIGRATION_DIR} postgres "host=${POSTGRES_HOST} port=${POSTGRES_PORT} dbname=${POSTGRES_DB} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} sslmode=disable" up