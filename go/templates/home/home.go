package home

import (
	"fmt"
	"html/template"
	"os"
	"path"
)

type Page struct {
	Title   string
	Content []template.HTML
}

var homePage Page = Page{
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
}

func GenerateHomePage() {
	pwd, err := os.Getwd()
	fmt.Println(pwd)
	if err != nil {
		panic(err)
	}
	pathToHome := path.Join(pwd, "templates/home/home.html.tmpl")
	fmt.Println(pathToHome)
	tmplFiles := []string{
		path.Join(pwd, "templates/base/base.html.tmpl"),
		pathToHome,
	}
	fmt.Println(path.Join(pwd, "templates/base/base.html.tmpl"))
	tmpl, err := template.New(pathToHome).ParseFiles(tmplFiles...)
	if err != nil {
		panic(err)
	}

	var f *os.File
	f, err = os.Create("home.html")
	if err != nil {
		panic(err)
	}
	for _, template := range tmpl.Templates() {
		fmt.Println(template.Name())
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
