# go-marc-wasm

WebAssembly bindings (wasmjs and wasip1) for the [aaronland/go-marc](https://github.com/aaronland/go-marc) package.

## Work in progress

This package, including documentation, is incomplete.

## Building

The easiest thing to do is use the handy `wasm` Makefile target. Note that you will need to have both [Go](https://go.dev/) and [TinyGo](https://tinygo.org/) (version 0.36 or higher) installed to build the wasm(js) and wasi(p1) modules, respectively.

```
$> make wasm
GOOS=js GOARCH=wasm \
		go build -mod vendor -ldflags="-s -w" \
		-o www/wasm/parse_marc034.wasm \
		cmd/parse-marc034-wasmjs/main.go
tinygo build \
		-target wasi \
		-no-debug \
		-o www/wasi/parse_marc034.wasm \
		./cmd/parse-marc034-wasip1/main.go
```

There is currently no support for the Go 1.24 [//go:wasmexport](https://go.dev/blog/wasmexport) directive since it does not support stings by default yet.

## wasm(js)

_TBW_

## wasi(p1)

For example:

```
$> wasmer ./www/wasi/parse_marc034.wasm '1#$aa$b80000$dW0825500$eW0822000$fN0273000$gN0265000' | jq
{
  "Scale": {
    "Code": "1"
  },
  "Ring": {
    "Code": "#"
  },
  "Subfields": {
    "$a": {
      "Code": "$a",
      "Value": "a"
    },
    "$b": {
      "Code": "$b",
      "Value": "80000"
    },
    "$d": {
      "Code": "$d",
      "Value": "W0825500"
    },
    "$e": {
      "Code": "$e",
      "Value": "W0822000"
    },
    "$f": {
      "Code": "$f",
      "Value": "N0273000"
    },
    "$g": {
      "Code": "$g",
      "Value": "N0265000"
    }
  }
}
```

## See also

* https://github.com/aaronland/go-marc
* https://go.dev/wiki/WebAssembly