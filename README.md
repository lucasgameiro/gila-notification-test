# GILA CODING CHALLENGE - NOTIFICATION TEST

## BACK-END / API

Back-end part of Gila Software coding challenge is written in Go. 

There is a Dockerfile inside backend dir to build and run the project

### REQUIREMENTS

Go 1.19+

### BUILD

Building project:

`$ go mod download`
`$ go build -o ./gila-test.bin /app/interface/http/main.go`

To build with docker you'll only need to run:

`$ docker build --tag backend-gila-test .`

### RUN

Running project:

`$ ./gila-test.bin`

To run project with docker:

`$ docker run -d -p 8080:8080 backend-gila-test`

## FRONT-END

Front-end part of Gila Software coding challenge is written in Javascript and uses React and Next.js

### REQUIREMENTS

Node.js 14.6.0+
Yarn 1.22
Back-end server running up

### BUILD

If needed, open the .env.local file to change back-end address (if localhost no changes is required)

Building project:

`$ yarn`
`$ yarn build`

### RUN

Running project:

`$ yarn start`
