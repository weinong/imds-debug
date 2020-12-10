GOVER ?= 1.13.15
GOOS ?= linux

all: build copy-bin

clean:
	rm -f imds-debug*

build:
	docker build --build-arg GOVER=$(GOVER) --build-arg GOOS=$(GOOS) -t build .

copy-bin:
	docker run --rm -v `pwd`:/target --name build build $(GOVER)
