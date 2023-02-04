.PHONY: all run gotool  help

BINARY="bluebell"

all: gotool run

run:
	@air ./conf/config.yaml

gotool:
	@go fmt ./
	@go vet ./

help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make run - 使用air热加载 Go 代码"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"
	@echo "测试一下"