# API Management Collection

# Config File

|Name|Description|Value Type|
|----|-----------|----------|
|Port|Port for the application to listen on|string|
|AuthorizedHosts|CORS Header for authorized hosts|[]string|
|Database.MysqlHost|Host of the server mysql|string|
|Database.MysqlPort|Port of the server mysql|string|
|Database.MysqlUser|Host of the mysql database|string|
|Database.MysqlPassword|Password of the mysql database|string|
|Database.MysqlDatabase|Name of the mysql database|string|
|Log.Formatter|Format of the logs|ascii/json|
|Log.Output|Where the logs will be|stdout/stderr|
|Log.Level|Level of logs|debug/info/warn/crit|

# Prerequisites
Two prerequisites :
  1. Go >= 1.10
  2. Mysql >= 5.7

# Dev installation

```
git clone https://github.com/youkoulayley/api-collection
cd api-collection
go get
cp conf.example.json conf.json
cp conf.example.json conf_test.json
```

First, clone the repository, then go to the folder and fetch all dependencies. You are now ready to work.
Finally you have to copy the conf example twice, one for the app and one for the test. Fill those files with correct values.

# Credits

Based on the work of Sorastro.
