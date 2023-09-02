# logsService

Service for save logs.

**What is contained:**
1. Kafka consumer implementation
2. MongoDB implementation
3. logrus implementation
4. Graceful shutdown
5. Docker compose file
6. Make file


**How To start**
1. add to .env file variables for DB 
```dotenv
MONGO_DB_HOST: localhost
MONGO_DB_PORT: 27017
MONGO_DB_USERNAME: root
MONGO_DB_PASSWORD: myPassword
MONGO_DB_DATABASE: myDB
MONGO_DB_COLLECTION: myCollection
```
2. run makefile with command - "make run"
