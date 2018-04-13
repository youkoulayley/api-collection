# API Management Collection

# Config File

|Name|Description|Value Type|
|----|-----------|----------|
|Version|Version of the application|string|
|Port|Port for the application to listen on|string|
|AuthorizedHosts|CORS Header for authorized hosts|[]string|
|Database.Host|Host of the database server|string|
|Database.Port|Port of the database server|string|
|Database.User|Host of the database|string|
|Database.Password|Password of the database|string|
|Database.Database|Name of the database|string|
|Log.Formatter|Format of the logs|ascii/json|
|Log.Output|Where the logs will be|stdout/stderr|
|Log.Level|Level of logs|debug/info/warn/crit|

# Prerequisites
Two prerequisites :
  1. Go >= 1.10
  2. Postgres >= 10
  3. Glide

# Dev installation
## Database
```
docker run -it -d -p 5432:5432 --restart=always -e POSTGRES_PASSWORD="apicollection" -e POSTGRES_USER="apicollection" -e POSTGRES_DB="api_collection" --name=postgres postgres
```

## App
```
git clone https://github.com/youkoulayley/api-collection
cd api-collection
glide install
cp conf.example.json conf.json
cp conf.example.json conf_test.json
```

First, clone the repository, then go to the folder and fetch all dependencies. You are now ready to work.
Finally you have to copy the conf example twice, one for the app and one for the test. Fill those files with correct values.

# Credits

Based on the work of Sorastro.
