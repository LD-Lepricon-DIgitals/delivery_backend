#!/bin/bash

# Load environment variables from ./config/local.env
ENV_FILE="./config/local.env"
if [[ -f "$ENV_FILE" ]]; then
  export $(grep -v '^#' "$ENV_FILE" | xargs)
else
  echo "Error: Environment file $ENV_FILE not found."
  exit 1
fi

# Check if Container_id is set
if [[ -z "$CONTAINER_ID" ]]; then
  echo "Error: Environment variable 'Container_id' is not set."
  exit 1
fi

# Check required environment variables
if [[ -z "$DB_USERNAME" || -z "$DB_PASSWORD" || -z "$DB_NAME" ]]; then
  echo "Error: Missing one or more required variables in $ENV_FILE:"
  echo "  - DB_USER: Database username"
  echo "  - DB_PASSWORD: Database password"
  echo "  - DB_NAME: Database name"
  exit 1
fi

# Path to the schema file
SCHEMA_FILE="./db/schema/000001_init.up.sql"

# Check if the schema file exists
if [[ ! -f "$SCHEMA_FILE" ]]; then
  echo "Error: Schema file $SCHEMA_FILE not found."
  exit 1
fi

# Ensure the Docker container is running
if [[ "$(docker ps -q -f name=$CONTAINER_ID)" == "" ]]; then
  echo "Error: PostgreSQL container $CONTAINER_ID is not running."
  exit 1
fi

# Copy the schema file to the container
echo "Copying schema file to container..."
docker cp "$SCHEMA_FILE" "$CONTAINER_ID:/schema.sql"

# Apply the schema to the database
echo "Applying schema to the database..."
docker exec -i "$CONTAINER_ID" psql -U "$DB_USERNAME" -d "$DB_NAME" -f /schema.sql

# Cleanup the schema file from the container
echo "Cleaning up schema file from container..."
docker exec -i "$CONTAINER_ID" rm /schema.sql

echo "Database initialization complete!"
