# Examples

Each directory here is a complete goi app: a `main.go` that loads one or more
`.goi` designs, finds the nodes it needs by id and attaches behaviour. The
screenshots below are those apps, captured in the light scheme and in the dark
scheme, at the size the app opens with.

Run one from this directory — a design is loaded by its relative path, so the
working directory has to be `examples/`:

```sh
CGO_ENABLED=0 go run -tags "nofakecgo all" ./counter
```

`make counter` does the same, and every example has a target of that name. The
examples enable every optional feature with `-tags all`; a plain build keeps only
the core and turns on a feature with its own tag (`emoji`, `svg`, `mermaid`,
`markdown`, `code`, `net`, `dialog`). The build is cgo-free: purego loads the
SDL3 runtime library at run time, so there is no C toolchain to install. A few examples need more than the framework —
`media` needs `ffmpeg` on `PATH` (or `GOI_FFMPEG`), while `map` and `tv` need
network access.

## Start here

### hello

Nested boxes: radius, border, padding and alignment, with real text. It is the
visual check for the renderer, and the smallest design in the repository.

| Light | Dark |
| --- | --- |
| ![hello in the light scheme](images/hello-light.png) | ![hello in the dark scheme](images/hello-dark.png) |

### counter

The smoke test for the whole path: a window loaded from a `.goi` file, buttons
whose handlers update a label through the typed property API, and a counter that
lives in Go, outside `Bind`.

| Light | Dark |
| --- | --- |
| ![counter in the light scheme](images/counter-light.png) | ![counter in the dark scheme](images/counter-dark.png) |

### gallery

Every Material 3 Expressive component in one window, one category page at a
time. It is the visual reference for the component set and the manual test bed
for click, drag, scroll and type.

| Light | Dark |
| --- | --- |
| ![gallery in the light scheme](images/gallery-light.png) | ![gallery in the dark scheme](images/gallery-dark.png) |

## Layout and painting

### composition

One design split across several `.goi` files: a root that imports fragments, and
fragments that take props.

| Light | Dark |
| --- | --- |
| ![composition in the light scheme](images/composition-light.png) | ![composition in the dark scheme](images/composition-dark.png) |

### shapes

The primitive-quality test bed: circles, rounded boxes with radii and borders,
thick curved strokes and a stroked bezier path, all on one page.

| Light | Dark |
| --- | --- |
| ![shapes in the light scheme](images/shapes-light.png) | ![shapes in the dark scheme](images/shapes-dark.png) |

### scroll

The `scroll` property and the Go offset API: a chip row that scrolls sideways
and a card page that scrolls down.

| Light | Dark |
| --- | --- |
| ![scroll in the light scheme](images/scroll-light.png) | ![scroll in the dark scheme](images/scroll-dark.png) |

### popover

The Popover component: three alignments, a panel that flips above the bottom edge
instead of leaving the window, and one panel kept open through a bridge value.

| Light | Dark |
| --- | --- |
| ![popover in the light scheme](images/popover-light.png) | ![popover in the dark scheme](images/popover-dark.png) |

## Data and language

### bridge

The design calls Go handlers through a `bridge` block, and a fragment receives
one of those values and one of those handlers as props. Every event is appended
to the log.

| Light | Dark |
| --- | --- |
| ![bridge in the light scheme](images/bridge-light.png) | ![bridge in the dark scheme](images/bridge-dark.png) |

### datalist

A list the design draws from data the app holds: the `.goi` file declares the
values and handlers, Go fills them with `ui.SetData` and `ui.SetHandler`, and the
`for … range` loop builds the rows.

| Light | Dark |
| --- | --- |
| ![datalist in the light scheme](images/datalist-light.png) | ![datalist in the dark scheme](images/datalist-dark.png) |

### table

A sortable data table filled from Go. A sortable header sorts, clicking a row
selects it, and Go can sort or select too.

| Light | Dark |
| --- | --- |
| ![table in the light scheme](images/table-light.png) | ![table in the dark scheme](images/table-dark.png) |

### format

The `formatf` builtin: every text is interpolated from the bridge, the button
writes a new count from Go, and each line follows it on the next frame.

| Light | Dark |
| --- | --- |
| ![format in the light scheme](images/format-light.png) | ![format in the dark scheme](images/format-dark.png) |

### i18n

`t(...)` reads the catalogs under `i18n/`, the buttons switch the language from
Go, and editing a JSON catalog updates the window without a restart.

| Light | Dark |
| --- | --- |
| ![i18n in the light scheme](images/i18n-light.png) | ![i18n in the dark scheme](images/i18n-dark.png) |

### markdown

One document rendered by the Markdown component: headings, emphasis, lists,
quotes and code, in a scrolling column.

| Light | Dark |
| --- | --- |
| ![markdown in the light scheme](images/markdown-light.png) | ![markdown in the dark scheme](images/markdown-dark.png) |

## Theme and platform

### theme

The built-in light and dark schemes plus a brand scheme the app defines after the
window exists, which a `.goi` file selects with `theme: brand`.

| Light | Dark |
| --- | --- |
| ![theme in the light scheme](images/theme-light.png) | ![theme in the dark scheme](images/theme-dark.png) |

### seed

One colour, a whole scheme: `ui.NewThemeFromSeed` builds a light and a dark theme
from the same seed, both defined by name, so the design picks one with
`theme: lightFromSeed` or `theme: darkFromSeed`. A button opens the colour
picker; OK applies the picked colour as the new seed.

| Light | Dark |
| --- | --- |
| ![seed in the light scheme](images/seed-light.png) | ![seed in the dark scheme](images/seed-dark.png) |

### cursor

Hover any tile to see the pointer shape its `cursor` names. A child inherits its
parent's cursor until it declares its own, and `auto` returns to the framework's
shape.

| Light | Dark |
| --- | --- |
| ![cursor in the light scheme](images/cursor-light.png) | ![cursor in the dark scheme](images/cursor-dark.png) |

### window

The Window element's own properties: size limits, the resize, minimize and
maximize switches, fullscreen, and the state the platform reports back.

| Light | Dark |
| --- | --- |
| ![window in the light scheme](images/window-light.png) | ![window in the dark scheme](images/window-dark.png) |

### titlebar

A title bar drawn by the design. The `.goi` file picks the Linux or the macOS
layout with `app.os`; Go only wires the buttons to the Window control methods.

| Light | Dark |
| --- | --- |
| ![titlebar in the light scheme](images/titlebar-light.png) | ![titlebar in the dark scheme](images/titlebar-dark.png) |

### svg

One SVG icon file recoloured three ways: by the design, by a theme role, and from
Go.

| Light | Dark |
| --- | --- |
| ![svg in the light scheme](images/svg-light.png) | ![svg in the dark scheme](images/svg-dark.png) |

## Whole apps

### charts

A line, an area, a bar, a pie and a scatter chart plus a gauge in one window. The
design styles them with properties, Go feeds the series that change, and hovering
a point shows a tooltip.

| Light | Dark |
| --- | --- |
| ![charts in the light scheme](images/charts-light.png) | ![charts in the dark scheme](images/charts-dark.png) |

### map

Searches PiriAPI for a place and draws its boundary on an OpenStreetMap map. The
design owns the search field, the result list and the handlers; Go owns the two
HTTP calls.

| Light | Dark |
| --- | --- |
| ![map in the light scheme](images/map-light.png) | ![map in the dark scheme](images/map-dark.png) |

### media

Audio and video players. The video plays a local clip with its SRT subtitle
track drawn over the picture, and the two audio bars show the default bar and the
`wavy` variant with its `surfaceColor` set from a theme role. The buttons drive
the same methods a design calls by id.

| Light | Dark |
| --- | --- |
| ![media in the light scheme](images/media-light.png) | ![media in the dark scheme](images/media-dark.png) |

### tv

A live television player. The channel list is bridge data filled from Go, and
picking a row sets the URL the `Video` plays; the list itself slides in and out
with a tween on its width, driven from a Go handler. The channels are public HLS
streams, so it needs network access.

| Light | Dark |
| --- | --- |
| ![tv in the light scheme](images/tv-light.png) | ![tv in the dark scheme](images/tv-dark.png) |

### search

The smallest bridge: a text field and a filled button, both sending the field's
text to a Go handler.

| Light | Dark |
| --- | --- |
| ![search in the light scheme](images/search-light.png) | ![search in the dark scheme](images/search-dark.png) |

### league

A whole app: the Turkish Süper Lig table with zone colours, team badges and each
team's last five results, built from a JSON file and a `for … range` over the
rows. The switch in the corner changes the scheme through a bridge handler.

| Light | Dark |
| --- | --- |
| ![league in the light scheme](images/league-light.png) | ![league in the dark scheme](images/league-dark.png) |

### shopping

A product page: a rail of categories, a picture column, a size and colour
picker, and a review card. One rose seed builds both schemes with
`NewThemeFromSeed`, and the design reflows from a desktop rail to a phone
bottom bar with `responsive`, a `Flex` that turns from a row into a column, and a
scroll column.

| Desktop, light | Desktop, dark |
| --- | --- |
| ![shopping at 1000x800 in the light scheme](images/shopping-light.png) | ![shopping at 1000x800 in the dark scheme](images/shopping-dark.png) |

| Phone, light | Phone, dark |
| --- | --- |
| ![shopping at 500x1000 in the light scheme](images/shopping-mobile-light.png) | ![shopping at 500x1000 in the dark scheme](images/shopping-mobile-dark.png) |

## In the browser

`web` is the counter design compiled to WebAssembly: the same `.goi` file, tree,
layout and painter as the desktop build, drawn with WebGPU into a canvas.

```sh
make web && python3 -m http.server -d web/dist
```

It has no screenshot here; a browser window is not the app's own.
