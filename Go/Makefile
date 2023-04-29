# vim: noet:ts=2:sw=2

.PHONY: build
binary = "../Docker/server"


build: $(binary)

$(binary): main.go
	CGO_ENABLED=0 go build -o $(binary)

clean:
	rm -f $(binary)

