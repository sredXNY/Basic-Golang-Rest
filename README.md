# Basic-Golang-Rest
Basic golang restful API. Using Gorm and Gin. It's a simple service to CRUD users.

# Endpoints documentation
Inside each http handler is located a curl command as an example on how to consume that specific endpoint

# Database connection
Inside the function InitDb you will have to writte the database credentials, just replace the values inside `gorm.Open("postgres", "host=xxx.xxx.xxx.xxx user=user dbname=test sslmode=disable password=ps")`
