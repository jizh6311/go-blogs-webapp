# go-blogs-webapp
This is a web app written in Go echo framework for backend and VueJS for frontend.

## Install GO
   - https://golang.org/dl/ version:1.10
   - Set Environment on Windows:
     - GOROOT: Go installation directory. C:/Go. MAC: "/usr/local/go"
     - GOPATH: Your preferred directory for multiple Go projects. The correct path for Glide tool looks 
     like: GOPATH/src/<project_directory>/vendor/... (It is tricky ...)
     - GOBIN: <YOUR_GOROOT>/bin

## Install MongoDB
   - https://www.mongodb.com/download-center#community. On Mac, you can do ```brew install mongodb``` to install
   - MongoDB Compass is very helpful: https://www.mongodb.com/products/compass
   - Create a folder for local MongoDB data. The default directory is C:/data. On Mac: /data/db
   
## Install Glide
   - Glide is a package manager for Go language. The version must be 0.12.3 on windows. The new version is OK on Mac
   - On Mac, use ```curl https://glide.sh/get | sh``` to install the latest glide
   - On Windows, download the zip file from https://github.com/Masterminds/glide/releases. Copy and paste the application file under <Your_GOBIN>
   - Run ```glide -v``` to check if the correct version is installed
   - For more information, please refer https://glide.readthedocs.io/en/latest/

## Project Setup
1. Clone this project under <Your_GOPATH>/src/go-blogs-webapp. Run ```glide install``` to install all packages in the vendor folder. To install new packages, run ```glide get <package>```
2. Start local MongoDB. The default .exe file is at ../MongoDB/Server/3.4/bin/mongod.exe
3. Under <project_directory>, run ```go run app.go``` to start the web app backend on localhost:8081
4. Under <project_directory>/public, run ```npm run dev``` to start the npm project on port 8080. The web services on 8081 will be available on 8080 as well because of the proxy
5. Go to <project_directory>/test and run ```go test ./...``` can execute all unit tests. Run ```go test``` under each test folder can execute every single test class

## Go Echo Framework Document
https://echo.labstack.com/

## VueJS Framework Document
https://vuejs.org

## Other Suggestions
1. Editor: Go-plus in Atom is recommended. GoLand in Intellij could also be a good option.
2. Backend API call: curl command or Postman software
3. Webpack is being used in the npm project. The changes made can be seen on the browser immediately 
