# !/bin/bash

# Load the .env file to environment variables
export \
$(grep 'DB_HOST' .env | xargs -0) \
$(grep 'DB_PORT' .env | xargs -0) \

# Wait for postgres to be ready
wait-for "${DB_HOST}:${DB_PORT}" -- "$@"

# Execute the command passed to the entrypoint
./main