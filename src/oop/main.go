package main

import (
	"fmt"
	_ "mymodule/employee"
)

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {
	title   string
	content string
	author
}

func (p post) details() {
	fmt.Println("Titel: ", p.title)
	fmt.Println("content: ", p.content)
	// fmt.Println("Author: ", p.author.fullName()) // this can also be written as p.fullname
	fmt.Println("Author: ", p.fullName())
	fmt.Println("bio: ", p.author.bio)
}

type website struct {
	posts []post
}

func (w website) contents() {
	fmt.Println("contents of the website \n")
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}

func main() {
	// e := employee.New("harish", "konda", 30, 20)

	// e.LeavesRemaining()
	author1 := author{
		bio:       "born on  oct 18 1992",
		firstName: "Harish",
		lastName:  "Konda",
	}
	post1 :=
		post{
			"inheritance in go",
			"Go supports composition instead of inheritance",
			author1,
		}
	post2 := post{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	post3 := post{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}
	w := website{
		posts: []post{
			post1,
			post2,
			post3,
		},
	}
	w.contents()

}
