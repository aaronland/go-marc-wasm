package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/aaronland/go-marc/v3/fields"	
)

func main() {

	flag.Parse()

	for _, raw := range flag.Args() {
		fmt.Println(parse_marc034(raw))
	}
}

//export parse_marc034
func parse_marc034(raw string) string {

	p, err := fields.Parse034(raw)	

	if err != nil {
		return err.Error()
	} else {

		v, err := json.Marshal(p)

		if err != nil {
			return err.Error()
		}

		return string(v)
	}

}
