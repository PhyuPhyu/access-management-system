# Access Management System

Access management system developed with [Gin](https://gin-gonic.com) web framework and MySQL. [GORM](https://gorm.io) library used for database operations.

## Project Setup 

* Clone the project into your $GOPATH
```
git clone git@git.irrasoft.com:team-nhtp/nhtp-api.git
cd projectFolder
```
* Create database schema in mysql 
```
CREATE SCHEMA `access_management_system` ;
```

* Modify app.env file
```
DB_USERNAME=dbusername
DB_PASSWORD=dbpassword
```

* Install Dependencies

From project directory, run:
```
go mod tidy
```

* Run the app
```
go run main.go
```
The AMS will run on the default port 8080