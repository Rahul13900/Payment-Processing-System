build:
	@mkdir -p bin
	@cd src/user-service && go build -o ../../bin/user-service .

run: build
	@./bin/user-service

clean:
	@rm -f bin/user-service