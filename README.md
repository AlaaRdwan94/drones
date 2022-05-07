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

## Task NOTEs:
I decide to make this simulation as simple as possible by doing the following:
  - assume that the fleet should registered with only 10 drones 
  - the create of the drone serial and the validation is from the server side , so no validation function needed in the API request side
  - the server pre data needed is seeded in the seed service 
  - I prevent to make the service for creating a medications
  - the validations in the medications also applied in the server side
  - the whole simulation for the drone status (delivering , delivered , returned) is done in the server side throw crons
  - I assume that the drone with idle status must have recharged for next use
  - I assume that the battery is only in use when the status is delivering or returned
  - after the drone is delivered its loadings are deleted from the table droneMedications and its weight is being reseted to 0 .
  - I prevent to add an init value for the drone weight during the registeration
  - I add a JWT middleware service but not use it for future enhancments

please do not hesitate to reply to me if there is a new requirements for this tasks or if there a need to change one or more of the previous notes

## Task Running

- to run the app via docker composer run `docker-compose up`
- to run the app in your machine 
   - make sure you have postgresql database server is installed and run
   - create database and user with same configurations in `config.yml` file 
    - make sure that you have redis database server is installed and run
    - run `go mod vendor` 
    - run `go mod tidy`
    - run `go run .`


