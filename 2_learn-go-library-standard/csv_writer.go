package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"budi", "hari", "yaya"})
	_ = writer.Write([]string{"jaja", "kaka", "tata"})
	_ = writer.Write([]string{"lala", "wawa", "nana"})

	writer.Flush()
}
