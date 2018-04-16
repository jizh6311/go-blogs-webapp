# go-blogs-webapp
This is a web app written in Go echo framework

## Install GO
   - https://golang.org/dl/ version:1.10
   - Set Environment on Windows:
     - GOROOT: Go installation directory. C:\Go
     - GOPATH: Your preferred directory for multiple Go projects
     - GOBIN: <YOUR_GOROOT>\bin

## Install MongoDB
   - https://www.mongodb.com/download-center#community
   - MongoDB Compass is very helpful: https://www.mongodb.com/products/compass
   - Create a folder for local MongoDB data. The default directory is C:\data
   
## Install Glide
   - Glide is a package manager for Go language. The version being used is 0.12.3
   - Download the zip file from https://github.com/Masterminds/glide/releases. Copy and paste the application file under <Your_GOBIN>
   - Run ```glide -v``` to check if the correct version is installed
   - For more information, please refer https://glide.readthedocs.io/en/latest/

## Project Setup
1. Clone this project under <Your_GOPATH>/src/go-blogs-webapp. Run ```glide install``` to install all packages in the vendor folder. To install new packages, run ```glide get <package>```
2. Start local MongoDB. The default .exe file is at ../MongoDB/Server/3.4/bin/mongod.exe
3. Under <project_directory>, run ```go run app.go``` to start the web app on localhost:8000
4. Go to <project_directory>/test and run ```go test ./...``` can execute all unit tests. Run ```go test``` under each test folder can 
execute every single test class

## Go Echo Framework Document
https://echo.labstack.com/

## Other Suggestions
1. Editor: Go-plus in Atom is recommended. GoLand in Intellij could also be a good option.
2. Backend API call: Postman
