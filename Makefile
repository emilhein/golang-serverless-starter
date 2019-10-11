build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/echo echo/main.go
	sls deploy

.PHONY: clean
clean:
	rm -rf ./bin ./vendor Gopkg.lock

.PHONY: deploy
deploy: clean build
	sls deploy --verbose