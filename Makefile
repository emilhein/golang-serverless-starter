.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/echo echo/main.go
	sls deploy

# deploy:
# 	sls deploy

# clean:
# 	rm -rf ./bin ./vendor Gopkg.lock

# deploy: clean build
# 	sls deploy --verbose