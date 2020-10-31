SHELL=/usr/bin/env bash

all: build
.PHONY: all

unexport GOFLAGS

GOVERSION:=$(shell go version | cut -d' ' -f 3 | cut -d. -f 2)
ifeq ($(shell expr $(GOVERSION) \< 14), 1)
$(warning Your Golang version is go 1.$(GOVERSION))
$(error Update Golang to version $(shell grep '^go' go.mod))
endif

# git modules that need to be loaded
MODULES:=

CLEAN:=
BINS:=


#获取完整commit id（如：bb4f92a7d4cbafb67d259edea5a1fa2dd6b4cxxx）
#git rev-parse HEAD
#获取short commit id（如：bb4f92a）
#git rev-parse --short HEAD

ldflags=-X=github.com/shepf/star-tools/build.CurrentCommit=+git.$(subst -,.,$(shell git describe --always --match=NeVeRmAtCh --dirty 2>/dev/null || git rev-parse --short HEAD 2>/dev/null))
ifneq ($(strip $(LDFLAGS)),)
	ldflags+=-extldflags=$(LDFLAGS)
endif

GOFLAGS+=-ldflags="$(ldflags)"


## MAIN BINARIES
deps: $(BUILD_DEPS)
.PHONY: deps

debug: GOFLAGS+=-tags=debug
debug: star


star:
	rm -f star
	go build $(GOFLAGS) -o star ./cmd/star
#	go run github.com/GeertJohan/go.rice/rice append --exec star -i ./build
.PHONY: star
BINS+=star


monitor:
	rm -f star-monitor
	go build $(GOFLAGS) -o star-monitor ./cmd/star-monitor
.PHONY: star-monitor
BINS+=star-monitor


build: star monitor
	@[[ $$(type -P "star") ]] && echo "Caution: you have \
an existing star binary in your PATH. This may cause problems if you don't run 'sudo make install'" || true
.PHONY: build

install:
	install -C ./star /usr/local/bin/star
	install -C ./star-monitor /usr/local/bin/star-monitor


# MISC
clean:
	rm -rf $(CLEAN) $(BINS)
.PHONY: clean