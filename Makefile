RUN := go run -ldflags=-linkmode=internal

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

typography:
	$(RUN) ./typography

composition:
	$(RUN) ./composition

theme:
	$(RUN) ./theme

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

fmt:
	go run ./cmd/goifmt -w .

fmt-list:
	go run ./cmd/goifmt -l .

.PHONY: search spike counter gallery hello scroll typography composition theme titlebar datalist bridge svg window format fmt fmt-list
