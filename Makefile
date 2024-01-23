.PHONY: build
build:
	@docker build -t shortener . 


.PHONY: in-memory
in-memory: build
	@docker run -it -p 8080:8080 -p 9090:9090 shortener in-memory


.PHONY: postgres
postgres: build
	@docker run -it -p 8080:8080 -p 9090:9090 shortener postgres


.PHONY: test
test:
	go test -v ./internal/grcpserver
	go test -v ./internal/shorting
	go test -v ./internal/storage
