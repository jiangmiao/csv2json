CSV2JSON

Convert CSV to JSON
Work with [jq](https://stedolan.github.io/jq/) a powerful JSON querier.

Installation:
  go get github.com/jiangmiao/csv2json

Usage:
  csv2json < STDIN

Options:
  -comma string
        comma (default ",")
  -no-header
        no header, the row will be Array instead of Object
  -no-trim
        disable trim leading space

Example:
  $ cat example.csv
  Name,Age
  Alice,21
  Bob,22
  Eve,21

  $ csv2json < example.csv
  [{"Age":"21","Name":"Alice"},{"Age":"22","Name":"Bob"},{"Age":"21","Name":"Eve"}]

  $ csv2json < example.csv | jq '.[]|select(.Name=="Eve")'
  {
    "Age": "21",
    "Name": "Eve"
  }
