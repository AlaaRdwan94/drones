## Task tools and packages

- ✨Gin✨ 
- ✨gorilla/websocket✨ 
- ✨Gorm✨
- ✨Viber✨
- [✨gocron✨](https://golangrepo.com/repo/go-co-op-gocron-go-cron-jobs)
## Task technology information

- dataBase driver : postgres
- running server machine : docker images in docker-compose

## Task architecture and model management information

- try to use clean architecture .
- drone directory will contain the Drones APIs , usecase and repo implementation
- infrastructure directory will contain the main configurations for the app (database connection , seeder ,cron job ...etc)
- entity directory will have the main blueprint for the drones table in the database
- model directory will contain the structs for the req and res Json objects
- transformer will transform between models and entities
- use `go mod` and `go vendor` for dependency management
- use `gorm migration` for creating the schema 

## Task Running

- to run the app via docker composer run `docker-compose up`
- to run the app in your machine 
   - make sure you have postgresql database server is installed and run
   - create database and user with same configurations in `config.yml` file 
    - make sure that you have redis database server is installed and run
    - run `go mod vendor` 
    - run `go mod tidy`
    - run `go run .`


