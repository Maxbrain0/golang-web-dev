package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     int
}

type region struct {
	Region string
	Hotels []hotel
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	regions := []region{
		region{
			Region: "Southern",
			Hotels: []hotel{
				hotel{
					Name:    "Cooltel",
					Address: "12345 Coolton",
					City:    "Apply Valley",
					Zip:     12345,
				},
				hotel{
					Name:    "Lametel",
					Address: "54321 Lamelane",
					City:    "San Diego",
					Zip:     54321,
				},
			},
		},
		region{
			Region: "Central",
			Hotels: []hotel{
				hotel{
					Name:    "Fresno Inn",
					Address: "1 Fresno Ave",
					City:    "Fresno",
					Zip:     91111,
				},
				hotel{
					Name:    "Merced Suites",
					Address: "111 Cow Lane",
					City:    "Merced",
					Zip:     91211,
				},
			},
		},
		region{
			Region: "Northern",
			Hotels: []hotel{
				hotel{
					Name:    "Seaside Dream",
					Address: "1 Trinidad Coast Drive",
					City:    "Trinidad",
					Zip:     94768,
				},
				hotel{
					Name:    "Y Not?",
					Address: "512 Yreka Ct",
					City:    "Yreka",
					Zip:     95123,
				},
			},
		},
	}
	err := tpl.Execute(os.Stdout, regions)
	if err != nil {
		log.Fatalln(err)
	}
}
