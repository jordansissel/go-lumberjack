lumberjack:
	GOPATH=$(PWD) go build -o $@

# Use make because GOPATH doesn't default to $PWD.
.DEFAULT:
	GOPATH=$(PWD) go $@
