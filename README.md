# Golang Skeleton

### Format app
```
go fmt ./...
```
### run app
```
go run cmd/server/main.go
```

### A typical top-level directory layout


    ├── cmd                    
    │   └── server             # Contains the main entry file
    │       
    ├── internal              
    │   ├── config             # Importing of env variables
    │   ├── models             # Contains model files
    │   │  
    │   ├── server              
    │   │    ├──  handlers     # Handler functions that are called by routes
    │   │    └──  middlewares  # Middleware functions for attaching to routes
    │   └── services
    │       └──  postgres       # postgres database service
    └──  scripts                # bash scripts for aws code pipeline



#### main.go 


    ├── cmd                    
    │    └── server
    │          └── main.go      # This is the main entry to the app

- Main entry to the app.
- Logs the app logs to server.log file.
- Calls RunServer function in internal/server/server.go file.


#### server.go
    ├── internal              
    │   └── server              
    │        └──  server.go 

- Loads env variables by calling FromEnv() from internal/config/config.go
- Initializes all services to be used in the app.
- Initializes all routes from internal/server/router.go
- Adds cors configs to the app.
- Starts the http server

You can update any cors settings in this file. Here is a snippet of the default cors config.

```go
	c := cors.New(cors.Options{
		AllowedHeaders: []string{"tenant", "*"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "UPDATE", "OPTIONS", "DELETE", "PATCH"},
	})
```

#### config.go
    ├── internal              
    │   └── config              
    │        └──  config.go 

- Loads all env variables using [godotenv](github.com/joho/godotenv) package
- Processes the env using [kelseyhightower/envconfig](github.com/kelseyhightower/envconfig) package

> Checkout [kelseyhightower/envconfig](github.com/kelseyhightower/envconfig) package to get a proper understanding on env variables naming and getting.

#### router.go
    ├── internal              
    │   └── server              
    │        └──  router.go 

- This file contains all routes to the app

#### service.go
    ├── internal              
    │   └── services              
    │        └──  service.go 

- Acts as an entry point to all services
- Loads services envs
- Initializes services DAOs

#### service structure
This is how a service is structured

    ├── service-name (postgres)              
    │   ├── service-name.go     # service config files and functions         
    │   └── dao.go             # interface with method signatures


#### scripts
    ├── scripts                 # aws code pipeline bash scripts              
    │   ├── build_app.sh        # executed in the build stage         
    │   ├── install_docker.sh   # executed during before install stage
    │   ├── start_app.sh        # executed during application start stage
    │   └── stop_app.sh         # executed during application stop stage
    │
    ├── appspec.yml             # aws code pipeline appspec yml
    └── buildspec.yml           # aws code pipeline buildspec yml (Executes the above scripts)





