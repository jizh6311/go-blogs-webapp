# go-blogs-webapp
This is a web app written in Go echo framework for backend and VueJS for frontend.

## Install GO
   - https://golang.org/doc/install
   - Set Environment:
     - GOPATH: Your preferred directory for multiple Go projects.
     - Highly suggest using direnv to manage your environment variables: `brew install direnv`

## Install MongoDB
   - Install MongoDB Community Server by:
   ```
    brew tap mongodb/brew
   ```
   ```
    brew install mongodb-community
   ```
   - MongoDB Compass is very helpful: https://www.mongodb.com/products/compass

## Project Setup
1. Clone this project under <Your_GOPATH>/src/go-blogs-webapp. Run `go install` to install all package dependencies
2. Run `./start-mongod.sh` to start local MongoDB server
3. Under <project_directory>, run ```go run app.go``` to start the web app backend on localhost:8081
4. Under <project_directory>/public, run ```npm run dev``` to start the npm project on port 8080. The web services on 8081 will be available on 8080 as well because of the proxy
5. Go to <project_directory>/test and run ```go test ./...``` can execute all unit tests. Run ```go test``` under each test folder can execute every single test class

## Go Echo Framework Document
https://echo.labstack.com/

## VueJS Framework Document
https://vuejs.org

## Let us pay more attentions to our work-life balance:
[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu)
[![LICENSE](https://img.shields.io/badge/license-NPL%20(The%20996%20Prohibited%20License)-blue.svg)](https://github.com/996icu/996.ICU/blob/master/LICENSE)
