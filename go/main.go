package main

import (
	"html/template"
	"os"
)

// type Post struct {
// 	Title    string
// 	Author   string
// 	Date     string
// 	Sections []template.HTML
// }
//
// var post Post = Post{
// 	Title:  "First Post",
// 	Author: "Joe Bonneau",
// 	Date:   "Saturday, October 21, 2023",
// 	Sections: []template.HTML{
// 		`<p>
//           This is the <strong>first</strong> paragraph! We can render <em>all kinds</em> of HTML.
//       </p>`,
// 		`<p>I generated this one using <strong>EMMETT</strong> which is real nice! Some other shit blah blah blah</p>
//       <hr />
//     `,
// 	},
// }

type Page struct {
	Title   string
	Content []template.HTML
}

var homePage Page = Page{
	Title: "Home",
	Content: []template.HTML{
		`<p>
          This is the <strong>first</strong> paragraph! We can render <em>all kinds</em> of HTML.
      </p>`,
		`<p>I generated this one using <strong>EMMETT</strong> which is real nice! Some other shit blah blah blah</p>
      <hr />
    `,
	},
}

func main() {
	// pwd, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }
	// tmplFile := path.Join(pwd, "templates/post/post.html.tmpl")
	// tmplFile := "blog.html.tmpl"
	tmplFiles := []string{
		"navbar.html.tmpl",
		"base.html.tmpl",
		"navbar.html.tmpl",
		"home.html.tmpl",
		// "post.html.tmpl",
		// "blogPost.html.tmpl",
	}
	tmpl, err := template.New("home.html.tmpl").ParseFiles(tmplFiles...)
	if err != nil {
		panic(err)
	}
	// style, err := os.ReadFile("post.css")
	// if err != nil {
	// 	panic(err)
	// }
	// tmplData := struct {
	// 	Post Post
	// }{Post: post}
	// tmplData := struct{}

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
