package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	flagNoTrim := flag.Bool("no-trim", false, "disable trim leading space")
	flagComma := flag.String("comma", ",", "comma")
	flagNoHeader := flag.Bool("no-header", false, "no header, the row will be Array instead of Object")
	flag.Parse()

	stat, err := os.Stdin.Stat()
	if stat.Mode()&os.ModeCharDevice != 0 {
		// no data in stdin
		fmt.Println(`convert csv to json from stdin

Usage:
  csv2json < STDIN

Options:`)
		flag.PrintDefaults()
		os.Exit(1)
	}

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bytes.NewBuffer(input))
	r.TrimLeadingSpace = !*flagNoTrim
	r.Comma = []rune(*flagComma)[0]
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var res interface{}
	if !*flagNoHeader {
		if len(rows) > 1 {
			header := rows[0]
			rows = rows[1:]
			objs := make([]map[string]string, len(rows))
			for y, row := range rows {
				obj := map[string]string{}
				for x, cell := range row {
					obj[header[x]] = cell
				}
				objs[y] = obj
			}
			res = objs
		} else {
			res = []map[string]string{}
		}
	} else {
		res = rows
	}
	output, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(output)
}
