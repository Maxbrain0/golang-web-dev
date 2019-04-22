package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type record struct {
	Date time.Time
	Open float64
	High float64
	Low  float64
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	// parse csv file into data struct
	stockRecords := csvToData("table.csv")

	// parse template
	tpl, err := template.ParseGlob("*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// execute template with csv data
	err = tpl.Execute(res, stockRecords)
	if err != nil {
		log.Fatalln(err)
	}
}

func csvToData(path string) []record {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}

	csvReader := csv.NewReader(file)

	csvData, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	// skip comment line, use length-1
	records := make([]record, 0, len(csvData)-1)

	for i, r := range csvData {
		// skip index comment line
		if i == 0 {
			continue
		}

		date, _ := time.Parse("2006-01-02", r[0])
		open, _ := strconv.ParseFloat(r[1], 64)
		high, _ := strconv.ParseFloat(r[2], 64)
		low, _ := strconv.ParseFloat(r[3], 64)
		records = append(records, record{
			Date: date,
			Open: open,
			High: high,
			Low:  low,
		})
	}

	return records
}
