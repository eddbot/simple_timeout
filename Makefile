run: build execute

build:
	@echo building...
	@go build -o ./bin/out
execute:
	@echo running...
	@./bin/out
