test-gen:
	go run main.go gen -p mibs -m IF-MIB
test-gen-dry:
	go run main.go gen -p mibs -m IF-MIB -d

