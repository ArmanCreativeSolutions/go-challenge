# Arman Go Challenge
This project using Domain Driven Design handles request from user segment service using: 
* [**Gin Gonic**](https://github.com/gin-gonic/gin)
* [**Gorm**](https://github.com/go-gorm/gorm)
* [**Sqlite**](https://sqlite.org)

I use gin framework in order to make application communicate with
REST api, and Sqlite because our data is structured, and also for its simplicity.
I use gorm since I prefer to use ORMs in terms of safety 
and more control over data programmatically. Also, I have chosen REST api
because it allows great variety of data formats, which make services
more flexible, and I have prior proficiency in working with REST.

## Terminology
- __User Segment__ &mdash; User segment is the main entity of this project. in each record I store user id which is considers as a UUID and the segment related to that user which has string format.

## Structure of the Code

```
 ┌───┐
 │ / │
 └─┬─┘
   │
   ├───────▶ application ───▶ user-segments ─ ─ ─ ─ related service
   │                                         │ 
   │                                         └──▶ dto ─ ─ ─ ─ data transfer object to recieve and map requests to models and vice versa
   │
   ├───────▶ domain ─ ─ ─ ─ models are defined here
   │                   │ 
   │                   └──▶ usersegments ─ ─ ─ ─ User Segment Model
   │
   ├───────▶ infrastructure  ─ ─ ─ ─ ─ Configs, general models (validations, response, errors, etc) and everything that the whole application uses
   │                             │
   │                             └──▶ dbconfig ─ ─ ─ ─ ─  Database configurations
   │                             │
   │                             └──▶ exception ─ ─ ─ ─ ─ General models for error handling
   │                             │
   │                             └──▶ repository ─ ─ ─ ─ ─ Repositories to communicate with database
   │                             │
   │                             └──▶ response ─ ─ ─ ─ ─ General models for application responses
   │                             │
   │                             └──▶ routes ─ ─ ─ ─ ─ Api routing setup
   │                             │
   │                             └──▶ validation ─ ─ ─ ─ ─ General validation for requests
   │
   ├───────▶ presentation ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ Controller layer 
   │              │
   │              └──▶ cronjobs ─ ─ ─ ─ ─ ─ ─ ─ ─ Define cron jobs                          
   │              │
   │              └──▶ usersegmentcontroller ─ ─ ─ Expose apis
```

## Setting up the environment
The system needs to have sqlite database. database name is hardcoded in sqlite-config.go as a constant variable. it is removed and created with each run.
also, application runs on localhost:8080 which is also hardcoded in main.go.

## Current issues
- It is better to use a separate environment file to avoid hardcode database name and port address.
- I tried to dockerized this project with a proper makefile, but the whole application package structure needed to be changed.
- I prefer to put repository next to the related model in domain package (I have done this practice in Java), however, since I had to do dependency injection without an IoC container (like Spring boot), I faced a circular dependency, so I had to move the repository from domain to infrastructure. I think the best solution is to have a repository interface and implement it in each repository.
- Use a middleware for logging could be a better practice.