package main

import "fmt"

type movie struct {
    title string
	rating  string
	year int
}

func (m movie) format() string {
	fmt.Printf("The movie %s (%v) has a rating of %s", m.title, m.year, m.rating)
	return "OK"
}
