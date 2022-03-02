package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func prettyJsonTo(w io.Writer, v interface{}) error {
	enc := json.NewEncoder(w)

	enc.SetIndent("", "\t")

	return enc.Encode(v)
}

func die(err interface{}) {
	fmt.Fprintf(os.Stderr, "%v\n", err)

	os.Exit(1)
}

func dief(f string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, f, v...)

	os.Exit(1)
}
