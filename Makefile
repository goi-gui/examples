# Every desktop example is its own goi project: cd into it and goi run reads the
# goi.yaml there, which carries the buildFlags and the assets directory. Install the
# CLI first: go install github.com/goi-gui/goi/cmd/goi@latest
GOI ?= goi

# The browser build: the same design compiled to wasm with goi build.
web:
	cd web && $(GOI) build --os js --arch wasm && \
		cp "$$(go env GOROOT)/lib/wasm/wasm_exec.js" dist/js-wasm/ && \
		cp index.html dist/js-wasm/ && \
		echo "serve it with: python3 -m http.server -d web/dist/js-wasm"

search:
	cd search && $(GOI) run

counter:
	cd counter && $(GOI) run

popover:
	cd popover && $(GOI) run

gallery:
	cd gallery && $(GOI) run

hello:
	cd hello && $(GOI) run

scroll:
	cd scroll && $(GOI) run

composition:
	cd composition && $(GOI) run

theme:
	cd theme && $(GOI) run

seed:
	cd seed && $(GOI) run

i18n:
	cd i18n && $(GOI) run

cursor:
	cd cursor && $(GOI) run

table:
	cd table && $(GOI) run

titlebar:
	cd titlebar && $(GOI) run

datalist:
	cd datalist && $(GOI) run

bridge:
	cd bridge && $(GOI) run

svg:
	cd svg && $(GOI) run

window:
	cd window && $(GOI) run

format:
	cd format && $(GOI) run

league:
	cd league && $(GOI) run

charts:
	cd charts && $(GOI) run

media:
	cd media && $(GOI) run

markdown:
	cd markdown && $(GOI) run

map:
	cd map && $(GOI) run

tv:
	cd tv && $(GOI) run

shopping:
	cd shopping && $(GOI) run

# goi setup installs the formatter next to the CLI.
FMT ?= goifmt

fmt:
	$(FMT) -w .

fmt-list:
	$(FMT) -l .

.PHONY: search spike counter popover gallery hello scroll composition theme seed titlebar datalist bridge svg window format fmt fmt-list i18n cursor table league charts media web markdown map tv shopping
