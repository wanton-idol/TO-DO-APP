# TO-DO-APP
It's a TO-DO App developed Using **ReactJs**, **Golang** and **MongoDB**.

## Application Requirement

### golang server requirement

1. golang https://golang.org/dl/
2. gorilla/mux package for router `go get -u github.com/gorilla/mux`
3. mongo-driver package to connect with mongoDB `go get go.mongodb.org/mongo-driver`
4. github.com/joho/godotenv to access the environment variable.

### react client

From the Application directory

`create-react-app client`

## Start the application

1. Make sure your mongoDB is started
2. Update the DB connection string in `.env` file inside the `go-server`.
3. From go-server directory, open a terminal and run
   `go run main.go`
4. From client directory,  
   a. install all the dependencies using `npm install`  
   b. start client `npm start`

## Walk through the application

Open application at http://localhost:3000

### Image

![](https://github.com/wanton-idol/TO-DO-APP/blob/main/images/Todo.png)


# License

MIT License

Copyright (c) 2023 Nishchal Gupta
