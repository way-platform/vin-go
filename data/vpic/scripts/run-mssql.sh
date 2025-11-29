#!/bin/bash
set -e

# Configuration
DB_PASSWORD="Admin@123"
DB_CONTAINER_NAME="mssql_local"
BACKUP_FILE_NAME="../VPICList_lite_2025_11.bak"
DB_NAME="vpic"

# Stop and remove existing container if it exists
if [ $(docker ps -a -q -f name=$DB_CONTAINER_NAME) ]; then
    echo "Stopping and removing existing container..."
    docker stop $DB_CONTAINER_NAME
    docker rm $DB_CONTAINER_NAME
fi

# 1. Pull the Docker image
echo "Pulling MS SQL Server 2025 image..."
docker pull mcr.microsoft.com/mssql/server:2022-latest

# 2. Run the Docker container
echo "Starting MS SQL Server container..."
docker run -e "ACCEPT_EULA=Y" -e "SA_PASSWORD=$DB_PASSWORD" \
    -p 1433:1433 --name $DB_CONTAINER_NAME -d \
    mcr.microsoft.com/mssql/server:2022-latest

# Wait a moment for the container to initialize
sleep 5

# Check if container is still running
if [ "$(docker ps -q -f name=$DB_CONTAINER_NAME)" = "" ]; then
    echo "Container failed to start. See logs below:"
    docker logs $DB_CONTAINER_NAME
    exit 1
fi

# 3. Wait for the database to be ready
echo "Waiting for MS SQL Server to start..."
db_ready=false
for i in {1..50}; do
    if /usr/bin/docker exec $DB_CONTAINER_NAME /opt/mssql-tools18/bin/sqlcmd \
        -S localhost -U sa -P "$DB_PASSWORD" -Q "SELECT 1" -b -C -No &> /dev/null; then
        echo "MS SQL Server is up!"
        db_ready=true
        break
    else
        echo "Waiting for MS SQL Server to start...attempt $i"
        sleep 5
    fi
done

if [ "$db_ready" = false ]; then
    echo "MS SQL Server did not start in time."
    docker logs $DB_CONTAINER_NAME
    exit 1
fi

# 4. Copy the backup file to the container
echo "Copying backup file to the container..."
docker cp "$BACKUP_FILE_NAME" $DB_CONTAINER_NAME:/var/opt/mssql/data/vpic.bak

# 5. Restore the database
echo "Restoring the database..."
docker exec $DB_CONTAINER_NAME /opt/mssql-tools18/bin/sqlcmd \
    -S localhost -U sa -P "$DB_PASSWORD" -b -C -No \
    -Q "RESTORE DATABASE [$DB_NAME] FROM DISK = N'/var/opt/mssql/data/vpic.bak' WITH FILE = 1,
MOVE N'vPICList_Lite' TO N'/var/opt/mssql/data/vPICList_Lite.mdf',
MOVE N'vPICList_Lite_log' TO N'/var/opt/mssql/data/vPICList_Lite_log.ldf',
NOUNLOAD, REPLACE, STATS = 5"

echo "Database '$DB_NAME' restored successfully!"
echo ""
echo "You can connect to the database using:"
echo "  Server: localhost,1433"
echo "  Username: sa"
echo "  Password: $DB_PASSWORD"
echo ""
echo "On Linux, you can use sqlcmd via Docker:
"
echo "  docker exec -it $DB_CONTAINER_NAME /opt/mssql-tools18/bin/sqlcmd -S localhost -U sa -P \"$DB_PASSWORD\" -C -No"
