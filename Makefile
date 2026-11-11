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

FMT := go run github.com/goi-gui/goi/cmd/goifmt

fmt:
	$(FMT) -w .

fmt-list:
	$(FMT) -l .

.PHONY: search spike counter gallery hello scroll composition theme seed titlebar datalist bridge svg window format fmt fmt-list i18n cursor table league charts media
