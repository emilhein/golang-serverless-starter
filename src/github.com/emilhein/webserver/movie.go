package main

import (
	"fmt"
)

type Movie struct {
	Title  string `json:"title"`
	Rating string `json:"rating"`
	Year   int    `json:"year"`
}

func (m Movie) format() string {
	fmt.Printf("The movie %s (%v) has a rating of %s", m.Title, m.Year, m.Rating)
	return "OK"
}
