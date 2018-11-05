UUID := $(shell uuidgen -t)
ga-test: clean
	go build -ldflags "-X github.com/vharsh/ga-test/pkg/ga.UUID=$(UUID)"

clean:
	rm -f ga-test

.PHONY: clean
