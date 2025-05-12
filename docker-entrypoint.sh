#!/bin/sh

# Wait for PostgreSQL to be ready
echo "Waiting for PostgreSQL to be ready..."
until nc -z postgres 5432; do
  sleep 1
done

# Wait for MongoDB to be ready
echo "Waiting for MongoDB to be ready..."
until nc -z mongodb 27017; do
  sleep 1
done

# Run migrations
echo "Running PostgreSQL migrations..."
./goravel-app artisan migrate

echo "Running MongoDB migrations..."
# Add MongoDB specific migrations if needed
# ./goravel-app artisan migrate:mongodb

# Start the application
echo "Starting the application..."
exec ./goravel-app
