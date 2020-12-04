test-gen:
	go run main.go gen -p mibs -m IF-MIB
test-gen-dry:
	go run main.go gen -p mibs -m IF-MIB -d
gen-agentx:
	go run main.go gen -p mibs -m AGENTX-MIB
build:
	go mod tidy;\
	go build


