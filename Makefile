GO = go
SCDOC = scdoc
LDFLAGS = "-s -w"
all:
	$(GO) build
fmt:
	$(GO) fmt
