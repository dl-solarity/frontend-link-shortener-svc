# Solarity Link Shortener Microservice

## Install

```bash
git clone github.com/dl-solarity/frontend-link-shortener-svc
cd frontend-link-shortener-svc
go build main.go
export KV_VIPER_FILE=./config.yaml
./main run service
```

## Documentation

We do use openapi:json standard for API. We use swagger for documenting our API.

To open online documentation, go to [swagger editor](http://localhost:8080/swagger-editor/) here is how you can start it

```bash
cd docs
npm install
npm start
```

To build documentation use `npm run build` command,
that will create open-api documentation in `web_deploy` folder.

## Running from docker 
  
Make sure that docker is installed.

use `docker run ` with `-p 8080:80` to expose port 80 to 8080

```bash
docker build -t github.com/dl-solarity/frontend-link-shortener-svc .
docker run -e KV_VIPER_FILE=/config.yaml github.com/dl-solarity/frontend-link-shortener-svc
```

## Running from Source

* Set up environment value with config file path `KV_VIPER_FILE=./config.yaml`
* Provide valid config file

### Database

This service utilizes ***Redis*** database. 
The easiest way to set it up is to use [docker image](https://hub.docker.com/_/redis).
