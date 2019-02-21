# ginserver
 ginserver is a simple framework for [gin](https://github.com/gin-gonic/gin), include web framework and restful api, and under development.
 
# Third Packages
 - Web Framework: [gin](https://github.com/gin-gonic/gin)
 - Authorization: 
    - [casbin](https://github.com/casbin/casbin) for RBAC auth
    - [session](https://github.com/gin-contrib/sessions) for web framework auth
    - [jwt]() for api auth
 - Database: [mysql](https://github.com/go-sql-driver/mysql)
 - Orm: [xorm](https://github.com/go-xorm/xorm)
 - Template: [AdminLTE](https://github.com/almasaeed2010/AdminLTE)
 - Swagger2.0: [gin-swagger](https://github.com/swaggo/gin-swagger)
 - ...

# Project Layout
 [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

# API Swagger Documents
    $ go get -u github.com/swaggo/swag/cmd/swag
    $ go get -u github.com/swaggo/gin-swagger
    $ go get -u github.com/swaggo/gin-swagger/swaggerFiles
    $ swag init -g cmd/ginserver/main.go