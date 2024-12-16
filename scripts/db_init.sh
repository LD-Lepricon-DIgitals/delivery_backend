#!/bin/bash

# Load environment variables from the local.env file
ENV_FILE="../config/local.env"
if [[ -f "$ENV_FILE" ]]; then
  export $(grep -v '^#' "$ENV_FILE" | xargs)
else
  echo "Error: Environment file $ENV_FILE not found."
  exit 1
fi

# Check for required environment variables
if [[ -z "$DB_USER" || -z "$DB_PASSWORD" || -z "$DB_NAME" ]]; then
  echo "Error: Missing one or more required environment variables:"
  echo "  - DB_USER: Database username"
  echo "  - DB_PASSWORD: Database password"
  echo "  - DB_NAME: Database name"
  exit 1
fi

# Check if container ID is provided
if [[ -z "$1" ]]; then
  echo "Error: Missing required container ID as the first argument."
  echo "Usage: ./init_db.sh <container_id>"
  exit 1
fi

DB_CONTAINER_NAME="$1"

# Path to the schema SQL file
SCHEMA_FILE="../db/schema/000001_init.up.sql"

# Check if the schema file exists
if [[ ! -f "$SCHEMA_FILE" ]]; then
  echo "Error: Schema file $SCHEMA_FILE not found."
  exit 1
fi

# Start PostgreSQL container if not running
if [[ "$(docker ps -q -f name=$DB_CONTAINER_NAME)" == "" ]]; then
  echo "Error: PostgreSQL container $DB_CONTAINER_NAME is not running."
  exit 1
fi

# Copy schema.sql to the container
echo "Copying schema file to container..."
docker cp "$SCHEMA_FILE" "$DB_CONTAINER_NAME:/schema.sql"

# Apply the schema to the database
echo "Applying schema..."
docker exec -i "$DB_CONTAINER_NAME" psql -U "$DB_USER" -d "$DB_NAME" -f /schema.sql

# Cleanup
echo "Cleaning up container schema file..."
docker exec -i "$DB_CONTAINER_NAME" rm /schema.sql

echo "Database initialization complete!"
