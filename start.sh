#!/bin/sh

set -e

echo "run db migration"
source /app/app.env
echo "read app.env"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up
echo "finish migrate"

echo "start the app"
exec "$@"
