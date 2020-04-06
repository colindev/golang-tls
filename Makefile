
DIR := .
SUBJ := $(shell cat subj.txt)

keypair:
	mkdir -p $(DIR)
	openssl ecparam -genkey -name secp384r1 -out $(DIR)/server.key
	openssl req -new -x509 -sha256 -key $(DIR)/server.key -out $(DIR)/server.crt -days 3650 -subj $(SUBJ)

run-server:
	go run src/server/main.go --addr :10443 -crt $(DIR)/server.crt -key $(DIR)/server.key

run-client:
	curl -i --cacert $(DIR)/server.crt https://xxx.io:10443 --resolve xxx.io:10443:127.0.0.1
