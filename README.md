# go-blogs-webapp
This is a web app written in Go echo framework

## Install GO
   - https://golang.org/dl/ version:1.10
   - Set Environment on Windows:
     - GOROOT: Go installation directory. C:\Go
     - GOPATH: Your project directory

## Install MongoDB
   - https://www.mongodb.com/download-center#community
   - MongoDB Compass is very helpful: https://www.mongodb.com/products/compass
   - Create a folder for local MongoDB data. The default directory is C:\data

## Project Setup
1. Install all go packages from imports
```
go get "github.com/labstack/echo"
go get "gopkg.in/mgo.v2/bson"
etc
```
2. Start local MongoDB. The default .exe file is at ../MongoDB/Server/3.4/bin/mongod.exe
3. Go to <project_directory>/main and run ```go run app.go``` to start the web app on localhost:8000
4. Go to <project_directory>/test and run ```go test ./...``` can execute all unit tests. Run ```go test``` under each test folder can 
every single test class

## Go Echo Frameword Document
https://echo.labstack.com/

## Other Suggestions
1. Editor: Go-plus in Atom is recommended. GoLand in Intellij could also be a good option.
2. Backend API call: Postman
