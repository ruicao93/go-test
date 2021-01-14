SHELL              := /bin/bash
GO                 ?= go
BINDIR             ?= $(CURDIR)/bin


.PHONY:bin
bin:
	@mkdir -p $(BINDIR)
	GOOS=windows $(GO) build -o $(BINDIR) github.com/ruicao93/go-test/cmd/...
