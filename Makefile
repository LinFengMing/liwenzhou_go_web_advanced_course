.PHONY: all build run gotool clean help
BINARY="bluebell"
all: gotool build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY} -v
run:
	@go run ./main.go config/config.yaml
gotool:
	go fmt ./
	go vet ./
clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
help:
	@echo "make - 格式化 Go 程式碼並編譯成二進位文件"
	@echo "make build - 編譯 Go 程式碼成二進位文件"
	@echo "make run - 直接運行 Go 程式碼"
	@echo "make clean - 移除二進位文件和 vim swap 文件"