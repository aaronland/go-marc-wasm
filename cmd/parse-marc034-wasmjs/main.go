package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"syscall/js"

	"github.com/aaronland/go-marc/v2/fields"
)

func ParseFunc() js.Func {

	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		marc034_str := args[0].String()

		logger := slog.Default()
		logger = logger.With("raw", marc034_str)

		logger.Info("Parse MARC 034")

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			resolve := args[0]
			reject := args[1]

			p, err := fields.Parse034(marc034_str)

			if err != nil {
				logger.Error("Failed to parse MARC 034", "error", err)
				reject.Invoke(fmt.Printf("Failed to parse '%s', %v\n", marc034_str, err))
				return nil
			}
			
			enc, err := json.Marshal(p)

			if err != nil {
				logger.Error("Failed to marshal response", "error", err)
				reject.Invoke(fmt.Printf("Failed to marshal result for '%s', %v\n", marc034_str, err))
				return nil
			}

			resolve.Invoke(string(enc))
			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}

func main() {

	parse_func := ParseFunc()
	defer parse_func.Release()

	js.Global().Set("parse_marc034", parse_func)

	c := make(chan struct{}, 0)

	slog.Info("WASM parse_marc034 function initialized")
	<-c

}
