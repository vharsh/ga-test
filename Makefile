UUID := $(shell uuidgen -t)
ga-test: clean
	go build -ldflags "-X github.com/vharsh/ga-test/main.uuid=$(UUID)"

clean:
	rm -f ga-test

.PHONY: clean
