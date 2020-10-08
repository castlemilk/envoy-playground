

proto.generate:
	docker build -t proto-builder -f protos/Dockerfile .
	docker run -v ${PWD}/protos:/build proto-builder


local.run:
	go run cmd/main.go

local.test:
	curl -vvv http://localhost:8080/api/v1/calculator -d '{"operation": "ADD", "first_operand": "1.1", "second_operand": "2.2"}'
	grpcurl -plaintext -d '{ "operation": "ADD", "first_operand": 1.1, "second_operand": 2.2 }' localhost:8080 Calculator/Calculate