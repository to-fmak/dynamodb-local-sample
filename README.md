# Overview
This is a sample project that creates dynamodb-local and dynamodb-admin with docker-compose, allowing creation and retrieval of DynamoDB tables and items using GUI, CLI, and Python.

# Usage
Start the containers with docker-compose:
```
docker-compose up -d
```
Access the DynamoDB admin interface at http://localhost:8001 via a web browser.

Run scripts under the cli-sample or python-sample directories to create and get tables and items.
