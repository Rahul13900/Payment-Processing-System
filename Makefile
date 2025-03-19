build:
	@mkdir -p bin
	@cd src/user-service && go build -o ../../bin/user-service .
	@cd src/payment-service && go build -o ../../bin/payment-service .

build-user:
	@mkdir -p bin
	@cd src/user-service && go build -o ../../bin/user-service .

build-payment:
	@mkdir -p bin
	@cd src/payment-service && go build -o ../../bin/payment-service .

run: build
	@./bin/user-service &
	@./bin/payment-service &

run-user: build-user
	@./bin/user-service

run-payment: build-payment
	@./bin/payment-service

clean:
	@rm -f bin/user-service bin/payment-service
