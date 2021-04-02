package main

import (
	"fmt"
	"log"

	"sort"
	"strings"
)

var States = map[string]string{
	"AZ": "Arizona",
	"CA": "California",
	"ID": "Idaho",
	"IN": "Indiana",
	"MA": "Massachusetts",
	"OK": "Oklahoma",
	"PA": "Pennsylvania",
	"VA": "Virginia",
}

const preFix = "....."

type Person struct {
	Name    string
	Address string
	City    string
	ST      string
}

func (p Person) ToString() string {
	return strings.Join([]string{preFix, p.Name, p.Address, p.City, States[p.ST]}, " ")
}

type AddressBook []Person

type NamedAddressList map[string][]Person

func ByState(str string) string {
	// your code
	PersonNames := NamedAddressList{}

	AddBook := ToPersons(str)

	//get a sorted list of states from states map keys
	var St []string
	for k := range States {
		St = append(St, k)
	}
	sort.Strings(St)
	//loop through sorted array and for strings for every customer with that state

	for _, p := range AddBook {

		//output  += States[s]+ "\n"+ preFix + p.Name +p.Address+p.City  +States[p.ST] + "\n"
		PersonNames[p.ST] = append(PersonNames[p.ST], p)

	}
	var finaladdress []string
	for _, s := range St {
		if persons, ok := PersonNames[s]; ok {
			var output []string
			//output  += States[s]+ "\n"+ preFix + p.Name +p.Address+p.City  +States[p.ST] + "\n"
			sort.SliceStable(persons, func(i, j int) bool {
				return persons[i].Name < persons[j].Name
			})
			output = append(output, States[s])

			for _, p := range persons {
				output = append(output, p.ToString())
				//output += preFix + p.Name +p.Address+p.City  +States[p.ST] + "\n"
			}
			finaladdress = append(finaladdress, strings.Join(output, "\n"))
		}
		//log.Println(output)

	}

	return strings.Join(finaladdress, "\n ")

}

func ToPersons(str string) []Person {
	AddBook := AddressBook{}

	input := strings.Split(strings.TrimRight(str, ""), "\n")
	for _, a := range input {
		fullAdd := strings.Split(a, ",")
		citystate := fullAdd[2]
		AddBook = append(AddBook, Person{
			Name:    strings.Trim(fullAdd[0], " "),
			Address: strings.Trim(fullAdd[1], " "),
			City:    strings.Trim(citystate[:len(citystate)-2], " "),
			ST:      strings.Trim(citystate[len(citystate)-2:], " "),
		})
	}
	return AddBook
}

//
//func StatesPeopleMap(AddBook []Person) []NamedAddressList {
//	PersonNames := make([]NamedAddressList,0)
//	for _,p := range AddBook{
//		peopleInState := PersonNames[p.ST]
//		updatedPeopleState := make([]Person,0)
//		updatedPeopleState = append(peopleInState,p)
//		//output  += States[s]+ "\n"+ preFix + p.Name +p.Address+p.City  +States[p.ST] + "\n"
//		PersonNames[p.ST] = append(PersonNames[p.ST],)
//
//	}
//	return PersonNames
//}
//
//func FormatAddresses(St []string, PersonNames []NamedAddressList) string {
//	var finaladdress []string
//	for _,s := range St{
//		if persons:= PersonNames[s];ok {
//			var output []string
//			//output  += States[s]+ "\n"+ preFix + p.Name +p.Address+p.City  +States[p.ST] + "\n"
//			sort.SliceStable(persons, func(i, j int) bool {
//				return persons[i].Name < persons[j].Name
//			})
//			output = append(output,States[s])
//
//			for _, p := range persons{
//				output =append(output,p.ToString())
//				//output += preFix + p.Name +p.Address+p.City  +States[p.ST] + "\n"
//			}
//			finaladdress = append(finaladdress,strings.Join(output,"\n"))
//		}
//		//log.Println(output)
//
//
//	}
//
//	return strings.Join(finaladdress,"\n ")
//}
func main() {
	var ad0 = `John Daggett, 341 King Road, Plymouth MA
Alice Ford, 22 East Broadway, Richmond VA
Orville Thomas, 11345 Oak Bridge Road, Tulsa OK
Terry Kalkas, 402 Lans Road, Beaver Falls PA
Eric Adams, 20 Post Road, Sudbury MA
Hubert Sims, 328A Brook Road, Roanoke VA
Amy Wilde, 334 Bayshore Pkwy, Mountain View CA
Sal Carpenter, 73 6th Street, Boston MA`
	fmt.Println(ByState(ad0))
	ad0sol := "California\n..... Amy Wilde 334 Bayshore Pkwy Mountain View California\n Massachusetts\n..... Eric Adams 20 Post Road Sudbury Massachusetts\n..... John Daggett 341 King Road Plymouth Massachusetts\n..... Sal Carpenter 73 6th Street Boston Massachusetts\n Oklahoma\n..... Orville Thomas 11345 Oak Bridge Road Tulsa Oklahoma\n Pennsylvania\n..... Terry Kalkas 402 Lans Road Beaver Falls Pennsylvania\n Virginia\n..... Alice Ford 22 East Broadway Richmond Virginia\n..... Hubert Sims 328A Brook Road Roanoke Virginia"
	log.Print(ByState(ad0) == ad0sol)

}
