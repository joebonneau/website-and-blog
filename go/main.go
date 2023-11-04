package main

import (
	"html/template"
	"os"
)

type RecentPost struct {
	Title       string
	Description string
	Date        string
	Tag         string
}

type HomePage struct {
	Title       string
	Content     []template.HTML
	RecentPosts []RecentPost
}

var homePage HomePage = HomePage{
	Title: "Home",
	Content: []template.HTML{
		`
    <h1>Hi, my name is</h1>
    <h2>Joe Bonneau.</h2>
    <h3>I'm a fullstack engineer.</h3>
    `,
		`
    <p>
      I primarily build web applications including writing APIs and building complex, beautiful interfaces.
      I love optimizing my development workflows and honing in on how to automate the developer experience.
      Outside of work, you'll find me doing more workflow optimization, discovering new music, or planning
      my next meal.
    </p>
    `,
	},
	RecentPosts: []RecentPost{
		{
			Title:       "My Path to Software Engineering",
			Description: "How I got into software engineering and what I've learned along the way.",
			Date:        "2023-10-31",
			Tag:         "Career",
		},
		{
			Title:       "Avoiding Decision Paralysis",
			Description: "Tech is exciting; explore without getting overwhelmed.",
			Date:        "2023-10-32",
			Tag:         "Career",
		},
	},
}

func GenerateHomePage() {
	tmplFiles := []string{
		"base.html.tmpl",
		"home.html.tmpl",
		"recentPost.html.tmpl",
	}
	tmpl, err := template.New("home.html.tmpl").ParseFiles(tmplFiles...)
	if err != nil {
		panic(err)
	}

	var f *os.File
	f, err = os.Create("home.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(f, homePage)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
}

func main() {
	GenerateHomePage()
}
