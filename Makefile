.PHONY: build
build:
	@docker build -t short . 


.PHONY: in-memory
in-memory: build
	@docker run --name=shortener -it -p 8080:8080 -p 9090:9090 short in-memory


.PHONY: postgres
postgres: build
	@docker run --name=shortener -it -p 8080:8080 -p 9090:9090 short postgres


.PHONY: clean
clean:
	@docker rm shortener
	@docker rmi short


.PHONY: test
test:
	go test -v ./internal/grcpserver
	go test -v ./internal/shorting
	go test -v ./internal/storage
