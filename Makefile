RUN := go run -ldflags=-linkmode=internal

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

fmt:
	$(FMT) -w .

fmt-list:
	$(FMT) -l .

.PHONY: search spike counter gallery hello scroll composition theme seed titlebar datalist bridge svg window format fmt fmt-list i18n cursor table league charts media web markdown