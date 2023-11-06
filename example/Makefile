build:
	$(RM) example
	CGO_ENABLED=0 go build -ldflags="-w -s" -gcflags=-trimpath=$(GOPATH)/src -asmflags=-trimpath=$(GOPATH)/src -o example example.go
