.DEFAULT_GOAL := swagger

install_swagger:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger:
	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models