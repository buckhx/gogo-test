all: go

clean:
	rm -rf gen/*

go:
	@mkdir -p ./gen/gogo 
	@mkdir -p ./gen/golang 
	protoc -I ./proto -I /usr/local/include \
		-I ${GOPATH}/src \
		--go_out=./gen/golang \
		./proto/*.proto
	protoc -I ./proto -I /usr/local/include \
		-I ${GOPATH}/src \
		--gogo_out=./gen/gogo \
		./proto/*.proto
