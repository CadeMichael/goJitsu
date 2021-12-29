package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

// main function  
func main() {
	createList()
        visitAll()
        indiv("https://www.bjjheroes.com/?p=7556")
}

// visitAll function  
func visitAll() {
	/* instatiate colly */
	c := colly.NewCollector(
		colly.AllowedDomains(
			"www.bjjheroes.com/",
			"bjjheroes.com/",
			"https://bjjheroes.com/",
			"www.bjjheroes.com",
			"bjjheroes.com",
			"https://bjjheroes.com",
		),
	)
	c.OnHTML("td.column-1 a[href]", func(e *colly.HTMLElement) {
                url := e.Request.AbsoluteURL(e.Attr("href"))
                indiv(url)

	})
	c.Visit("https://www.bjjheroes.com/a-z-bjj-fighters-list")
}

// createList function  
func createList() {
	file, err := os.OpenFile("list.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // instantiate csv
	if err != nil {
		log.Fatalf("cannot create file due to %s", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file) // create a writer to modify csv
	writer.Write([]string{
		"Name",
		"wins",
		"losses",
		"Win Points",
		"Win Advantages",
		"Win Submissions",
		"Win Decision",
		"Win Penalties",
		"Win EBI / OT",
		"Win DQ",
		"Losses Points",
		"Losses Advantages",
		"Losses Submissions",
		"Losses Decision",
		"Losses Penalties",
		"Losses EBI / OT",
		"Losses DQ",
	})
	writer.Flush()
}

// work function  
func indiv(url string) {
	file, err := os.OpenFile("list.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // instantiate csv
	if err != nil {
		log.Fatalf("cannot create file due to %s", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file) // create a writer to modify csv
	defer writer.Flush()
	/* instatiate colly */
	c := colly.NewCollector(
		colly.AllowedDomains(
			"www.bjjheroes.com/",
			"bjjheroes.com/",
			"https://bjjheroes.com/",
			"www.bjjheroes.com",
			"bjjheroes.com",
			"https://bjjheroes.com",
		),
	)
	fName := "0"
	wins := "0"
	lose := "0"
	wPoints := "0"
	wAdvantage := "0"
	wSubs := "0"
	wDec := "0"
	wPen := "0"
	wEBI := "0"
	wDQ := "0"
	lPoints := "0"
	lAdvantage := "0"
	lSubs := "0"
	lDec := "0"
	lPen := "0"
	lEBI := "0"
	lDQ := "0"
	/* <h1 itemprop="name">Aaron Johnson</h1> */
	c.OnHTML("h1", func(a *colly.HTMLElement) {
		fName = a.Text
	})
	c.OnHTML("div.Win_title", func(b *colly.HTMLElement) {
		wins = b.ChildText("em")
	})
	c.OnHTML("div.wrapper_canvas li", func(d *colly.HTMLElement) {
		switch d.ChildText("span.by_points") {
		case "BY POINTS":
			wPoints = d.ChildText("span.per_no")
		case "BY ADVANTAGES":
			wAdvantage = d.ChildText("span.per_no")
		case "BY SUBMISSION":
			wSubs = d.ChildText("span.per_no")
		case "BY DECISION":
			wDec = d.ChildText("span.per_no")
		case "BY PENALTIES":
			wPen = d.ChildText("span.per_no")
		case "BY EBI/OT":
			wEBI = d.ChildText("span.per_no")
		case "BY DQ":
			wDQ = d.ChildText("span.per_no")
		}
	})
	c.OnHTML("div.wrapper_canvas_lose li", func(d *colly.HTMLElement) {
		switch d.ChildText("span.by_points") {
		case "BY POINTS":
			lPoints = d.ChildText("span.per_no_lose")
		case "BY ADVANTAGES":
			lAdvantage = d.ChildText("span.per_no_lose")
		case "BY SUBMISSION":
			lSubs = d.ChildText("span.per_no_lose")
		case "BY DECISION":
			lDec = d.ChildText("span.per_no_lose")
		case "BY PENALTIES":
			lPen = d.ChildText("span.per_no_lose")
		case "BY EBI/OT":
			lEBI = d.ChildText("span.per_no_lose")
		case "BY DQ":
			lDQ = d.ChildText("span.per_no_lose")
		}
	})
	c.OnHTML("div.Win_title_lose", func(t *colly.HTMLElement) {
		lose = t.ChildText("em")
	})

	c.Visit(url)

	if wins != "0" || lose != "0" {
		writer.Write([]string{
			fName,
			wins,
			lose,
			wPoints,
			wAdvantage,
			wSubs,
			wDec,
			wPen,
			wEBI,
			wDQ,
			lPoints,
			lAdvantage,
			lSubs,
			lDec,
			lPen,
			lEBI,
			lDQ,
		})
	}
	fmt.Printf("!!!Done!!!")
}
