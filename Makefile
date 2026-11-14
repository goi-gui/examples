# The desktop examples use the purego SDL3 backend by default: no C toolchain and
# no SDL development package, only the SDL3 runtime library. `make ... PUREGO=0`
# selects the cgo backend instead (Linux needs the internal linker).
PUREGO ?= 1
UNAME := $(shell uname -s)
ifeq ($(PUREGO),1)
GO_ENV := CGO_ENABLED=0
GO_FLAGS := -tags nofakecgo
LD :=
else
GO_ENV := CGO_ENABLED=1
GO_FLAGS :=
ifeq ($(UNAME),Linux)
LD := -ldflags=-linkmode=internal
else
LD :=
endif
endif
RUN := $(GO_ENV) go run $(GO_FLAGS) $(LD)

# The browser build: the same design compiled to wasm, served from web/dist.
web:
	cd web && GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o dist/goi.wasm . && \
		cp "$$(go env GOROOT)/lib/wasm/wasm_exec.js" dist/ && cp index.html dist/ && \
		echo "serve it with: python3 -m http.server -d web/dist"

search:
	$(RUN) ./search

counter:
	$(RUN) ./counter

gallery:
	$(RUN) ./gallery

hello:
	$(RUN) ./hello

scroll:
	$(RUN) ./scroll

composition:
	$(RUN) ./composition

theme:
	$(RUN) ./theme

seed:
	$(RUN) ./seed

i18n:
	$(RUN) ./i18n

cursor:
	$(RUN) ./cursor

table:
	$(RUN) ./table

titlebar:
	$(RUN) ./titlebar

datalist:
	$(RUN) ./datalist

bridge:
	$(RUN) ./bridge

svg:
	$(RUN) ./svg

window:
	$(RUN) ./window

format:
	$(RUN) ./format

league:
	$(RUN) ./league

charts:
	$(RUN) ./charts

media:
	$(RUN) ./media

markdown:
	$(RUN) ./markdown

FMT := go run github.com/goi-gui/goi/cmd/goifmt

# Show which window backend the examples would use.
purego:
	@echo "$(RUN)"

fmt:
	$(FMT) -w .

fmt-list:
	$(FMT) -l .

.PHONY: search spike counter gallery hello scroll composition theme seed titlebar datalist bridge svg window format fmt fmt-list i18n cursor table league charts media web markdown purego