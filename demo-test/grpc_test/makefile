
# prototc 执行路径
protocExePath=/root/grpc

# 代码路径
srcPath=/root/git/go-test/demo-test/grpc_test

# protoc名字
protocName=protocol

protocRunDir=${protocExePath}/examples
protocOutDir=${srcPath}/${protocName}_out

all: build

build: clear proto

clear:
	rm -rf ${protocOutDir}
proto:
	rm -rf ${protocRunDir}/${protocName}
	cp -rf ${protocName} ${protocRunDir}
	cd ${protocRunDir} && protoc --go_out=. --go_opt=paths=source_relative \
        	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
        	${protocName}/${protocName}.proto && mv ${protocName} ${protocOutDir}
	
