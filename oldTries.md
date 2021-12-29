```go
	file, err := os.Create("test.csv") // instantiate csv
	if err != nil {
		log.Fatalf("cannot create file due to %s", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file) // create a writer to modify csv
	defer writer.Flush()
	writer.Write([]string{"first", "last", "wins", "losses"})

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
	indivC := colly.NewCollector(
		colly.AllowedDomains(
			"www.bjjheroes.com/",
			"bjjheroes.com/",
			"https://bjjheroes.com/",
			"www.bjjheroes.com",
			"bjjheroes.com",
			"https://bjjheroes.com",
		),
	)
	c.OnHTML("tr", func(e *colly.HTMLElement) {
		fmt.Printf("row found:\n %q \n", e.Text)

		first := e.ChildText("td.column-1")
		last := e.ChildText("td.column-2")
		baseUrl := "https://www.bjjheroes.com/bjj-fighters/"
		fighterPage := baseUrl + first + "-" + last
		wins := ""
		lose := ""
		indivC.OnHTML("div.Win_title", func(t *colly.HTMLElement) {
			wins = t.ChildText("em")
		})
		indivC.OnHTML("div.Win_title_lose", func(t *colly.HTMLElement) {
			lose = t.ChildText("em")
		})
		if wins != "" || lose != "" {
			writer.Write([]string{
				first,
				last,
				/* fighterPage ,*/
				wins,
				lose,
			})
		}
		indivC.Visit(fighterPage)
	})
	c.Visit("https://www.bjjheroes.com/a-z-bjj-fighters-list")
	fmt.Printf("!!!Done!!!")
```
