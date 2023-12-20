c=run service

build:
	@go build -o bin/frontend-link-shortener-svc

run: build
	@./bin/frontend-link-shortener-svc $(c)