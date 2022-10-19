package main

import (
	"fmt"
	"time"
)

type Name string
type Title string


type LendAudit struct {
	checkOut time.Time
	checkIn time.Time
}


type Member struct {
	Name string
	books map[Title]LendAudit
}

type BookEntry struct {
	total int
	lended int
}

type Library struct {
	members map[Name]Member
	books map[Title]BookEntry
}

func printMemberAudit(member *Member){
	for title, audit := range member.books {
	   var returnTime string
	   if audit.checkIn.IsZero() {
		returnTime = "[not return yet]"
	   }else{
		returnTime = audit.checkIn.String()
	   }

	   fmt.Println(member.Name, ":", title, ":", audit.checkOut.String(), ":", returnTime)
	}
}

func printMemberAudits(library *Library) {
	for _, member:= range library.members {
		printMemberAudit(&member)
	}
}

func printLibraryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/ lended:", book.lended)
	}
	fmt.Println()
}

func checkOutBook(library *Library, title Title, member *Member)bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Library not found")
		return false
	}
	
	if book.lended == book.total {
		fmt.Println("No more books available to lend")
		return false
	}

	book.lended += 1
	library.books[title] = book
	member.books[title] = LendAudit{checkOut: time.Now()}
	return true
}


func returnBook(library *Library, title Title, member *Member) bool {

	book, found := library.books[title]
	if !found {
		fmt.Println("Book not found ...")
		return false
	}

	audit, found := member.books[title]

	if !found {
		fmt.Println("Member did not checkout this book")
		return false
	}

	book.lended -= 1
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit
	return true;
}

func main(){
	library := Library{
		members: make(map[Name]Member),
		books: make(map[Title]BookEntry),
	}
	library.books["web app's in go"] = BookEntry{
		total: 5,
		lended: 0,
	}
	library.books["Let's learn go"] = BookEntry{
		total: 1,
		lended: 0,
	}
	library.books["go bootcamp"] = BookEntry{
		total: 3,
		lended: 0,
	}

	library.members["Dreamer"] = Member{
		Name:"Dreamer",
		books: make(map[Title]LendAudit),
	}


	library.members["Lost"] = Member{
		Name:"Lost",
		books: make(map[Title]LendAudit),
	}

	fmt.Println("\nInitial")
	printLibraryBooks(&library)
	printMemberAudits(&library)

	member := library.members["Dreamer"]
	checkedOut := checkOutBook(&library, "go bootcamp", &member)
	fmt.Println("\nCheck out book")

	if checkedOut {
	 printLibraryBooks(&library)
	 printMemberAudits(&library)
	}

	returnBook := returnBook(&library, "go bootcamp", &member)
	fmt.Println("\nreturn  book")

	if returnBook {
	 printLibraryBooks(&library)
	 printMemberAudits(&library)
	}
	
}