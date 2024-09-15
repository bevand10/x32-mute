
app=x32-mute
go=$(shell which go)
fmt=$(shell which gofmt)
static=CGO_ENABLED=0 GOOS=linux GOARCH=amd64

usage:
	@echo usage
	@echo
	@echo make build - compile $(app)
	@echo make static - compile $(app)
	@echo make dist
	@echo
	@echo make on - toggle muted
	@echo make off - toggle muted

static: tidy fmt
	$(static) $(go) build -o $(app) cmd/main.go

build: tidy fmt
	$(go) build -o $(app) cmd/main.go

tidy:
	$(go) mod tidy

fmt:
	$(fmt) -w .

on:
	./$(app) on

off:
	./$(app) off

dist:
	scp -C $(app) pcs:/usr/local/bin
