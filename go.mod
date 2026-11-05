module github.com/goi-gui/examples

go 1.27.3

require github.com/goi-gui/goi v0.2.0

require (
	git.sr.ht/~jackmordaunt/go-toast v1.1.2 // indirect
	github.com/akavel/rsrc v0.10.2 // indirect
	github.com/alecthomas/chroma/v2 v2.27.0 // indirect
	github.com/caarlos0/env/v11 v11.4.2 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dchest/jsmin v1.0.0 // indirect
	github.com/dlclark/regexp2/v2 v2.12.0 // indirect
	github.com/esiqveland/notify v0.14.0 // indirect
	github.com/fsnotify/fsnotify v1.10.1 // indirect
	github.com/gabriel-vasile/mimetype v1.4.16 // indirect
	github.com/gen2brain/beeep v0.11.3 // indirect
	github.com/go-logr/logr v1.4.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-playground/locales v0.14.2 // indirect
	github.com/go-playground/universal-translator v0.18.2 // indirect
	github.com/go-playground/validator/v10 v10.30.6 // indirect
	github.com/go-text/typesetting v0.3.5 // indirect
	github.com/go-webgpu/goffi v0.6.8 // indirect
	github.com/go-webgpu/webgpu v0.5.5 // indirect
	github.com/godbus/dbus/v5 v5.2.4 // indirect
	github.com/gogpu/gpucontext v0.31.3 // indirect
	github.com/gogpu/gputypes v0.8.2 // indirect
	github.com/gogpu/naga v0.21.2 // indirect
	github.com/gogpu/wgpu v0.34.7 // indirect
	github.com/jackmordaunt/icns/v3 v3.0.1 // indirect
	github.com/josephspurrier/goversioninfo v1.7.0 // indirect
	github.com/ksckaan1/logger v0.4.0 // indirect
	github.com/leodido/go-urn v1.5.0 // indirect
	github.com/mattn/go-colorable v0.1.16 // indirect
	github.com/mattn/go-isatty v0.0.24 // indirect
	github.com/ncruces/zenity v0.10.15 // indirect
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646 // indirect
	github.com/randall77/makefat v0.0.0-20260406194835-1b91746796b7 // indirect
	github.com/rs/zerolog v1.35.1 // indirect
	github.com/sergeymakinen/go-bmp v1.0.0 // indirect
	github.com/sergeymakinen/go-ico v1.0.0 // indirect
	github.com/srwiley/oksvg v0.0.0-20221011165216-be6e8873101c // indirect
	github.com/srwiley/rasterx v0.0.0-20220730225603-2ab79fcdd4ef // indirect
	github.com/tadvi/systray v0.0.0-20190226123456-11a2b8fa57af // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/otel v1.47.0 // indirect
	go.opentelemetry.io/otel/log v1.47.0 // indirect
	go.opentelemetry.io/otel/metric v1.47.0 // indirect
	go.opentelemetry.io/otel/trace v1.47.0 // indirect
	golang.org/x/crypto v0.58.0 // indirect
	golang.org/x/image v0.47.0 // indirect
	golang.org/x/net v0.62.0 // indirect
	golang.org/x/sys v0.49.0 // indirect
	golang.org/x/text v0.43.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
)

// The examples live next to a checkout of the framework, so the local tree is
// what they build against. Drop this to build against the module proxy.
replace github.com/goi-gui/goi => ../goi
