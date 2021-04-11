package main

import "fmt"

type Person struct {
	Name  string
	Likes []string
}
type People struct {
	pep []*Person
}

func (p People) get_likes() map[string][]*Person {
	likes := make(map[string][]*Person)
	for _, per := range p.pep {
		for _, l := range per.Likes {
			likes[l] = append(likes[l], per)
		}

	}
	return likes
}

func main() {
	aron := &Person{Name: "Aaron", Likes: []string{"Cheese", "Bacon", "Cookies"}}
	baron := &Person{Name: "Baron", Likes: []string{"Chocolate", "Bacon", "Coke"}}
	caron := &Person{Name: "Caron", Likes: []string{"Cake", "Bacon", "Cookies"}}
	daron := &Person{Name: "Daron", Likes: []string{"Chocolate", "Bane", "Canes"}}
	karon := &Person{Name: "Karon", Likes: []string{"Cheese", "Acorn", "Coke"}}
	saron := &Person{Name: "Saron", Likes: []string{"Cake", "Corn", "Wookies"}}

	Ny := &People{}
	var people []*Person
	people = append(people, aron)
	people = append(people, baron)
	people = append(people, caron)
	people = append(people, daron)
	people = append(people, karon)
	people = append(people, saron)
	Ny.pep = people

	liked := Ny.get_likes()

	for k, p := range liked {
		for _, person := range p {
			fmt.Println(person.Name, " likes ", k, ".")
		}
	}

}
