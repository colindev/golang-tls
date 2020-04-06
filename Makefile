
DIR := ./keypair_1
SUBJ := $(shell cat subj.txt)
HOST := xxx.io

keypair:
	mkdir -p $(DIR)
	openssl ecparam -genkey -name secp384r1 -out $(DIR)/server.key
	openssl req -new -x509 -sha256 -key $(DIR)/server.key -out $(DIR)/server.crt -days 3650 -subj $(SUBJ)

run-server:
	go run src/server/main.go --addr :10443 -crt $(DIR)/server.crt -key $(DIR)/server.key

run-client-curl:
	curl -i --cacert $(DIR)/server.crt https://$(HOST):10443 --resolve $(HOST):10443:127.0.0.1

run-client-go:
	go run src/client/main.go -cacert $(DIR)/server.crt -i -resolve $(HOST):10443:127.0.0.1 https://$(HOST):10443
