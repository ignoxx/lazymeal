package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"lazymeal/app/db"
	"os"
	"time"
)

const SITE_BASE_URL = "https://lazy-meal.com/"
const SITE_BASE_URL_SLUG_FORMAT = SITE_BASE_URL + "%s"

type ChangeFreqType string

const (
	ChangeFreqAlways  ChangeFreqType = "always"
	ChangeFreqHourly  ChangeFreqType = "hourly"
	ChangeFreqDaily   ChangeFreqType = "daily"
	ChangeFreqWeekly  ChangeFreqType = "weekly"
	ChangeFreqMonthly ChangeFreqType = "monthly"
	ChangeFreqYearly  ChangeFreqType = "yearly"
	ChangeFreqNever   ChangeFreqType = "never"
)

// URL represents a single URL entry in the sitemap.
type URL struct {
	Loc        string         `xml:"loc"`
	LastMod    string         `xml:"lastmod,omitempty"`
	ChangeFreq ChangeFreqType `xml:"changefreq,omitempty"`
	Priority   float64        `xml:"priority,omitempty"`
}

// URLSet represents the root element of the sitemap XML.
type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}

func main() {
	ctx := context.Background()
	d := db.Get()

	meals, err := d.GetAllMeals(ctx)
	if err != nil {
		panic(err)
	}

	urls := []URL{
		{
			Loc:        SITE_BASE_URL,
			LastMod:    time.Now().Format(time.RFC3339), // Include time
			ChangeFreq: "daily",
			Priority:   1.0,
		},
		{
			Loc:        SITE_BASE_URL + "meal-guide",
			LastMod:    time.Now().Format(time.RFC3339), // Include time
			ChangeFreq: "monthly",
			Priority:   0.9,
		},
	}

	for _, meal := range meals {
		url := URL{
			Loc:        SITE_BASE_URL + meal.Slug.String,
			LastMod:    meal.UpdatedAt.Format(time.RFC3339),
			ChangeFreq: ChangeFreqMonthly,
			Priority:   0.8,
		}

		urls = append(urls, url)
	}

	urlSet := URLSet{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  urls,
	}

	// Marshal the URLSet into XML format.
	xmlData, err := xml.MarshalIndent(urlSet, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling XML: %v\n", err)
		return
	}

	// Add XML header.
	xmlData = []byte(xml.Header + string(xmlData))

	// Write the XML data to a file.
	err = os.WriteFile("sitemap.xml", xmlData, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Println("Sitemap generated successfully: sitemap.xml")
}
