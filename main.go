package main

import (
	"fmt"
	"log"
	"strconv"
)

// One example of an interface type frome the standard library is the fmt.Stringer interface
// type Stringer intercace {
// String() string
// }

type Book struct {
	Title  string
	Author string
}

// Declare a Book type which satisfies the fmt.Stringer interface.
func (b Book) String() string {
	return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

// Declare a Count type which satisfies the fmt.Stringer interface.
type Count int

func (c Count) String() string {
	return strconv.Itoa(int(c))
}

// Describe a WriteLog() function which takes any object that satisfies the fmt.Stringer interface as a parameter

func WriteLog(s fmt.Stringer) {
	log.Println(s.String())
}

func main() {
	// Initialize a Book object and pass it to WriteLog()
	book := Book{"Think and Grow Rich", "Napolean Hill"}
	WriteLog(book)

	// Initialize a Count object and pass it to WriteLog()
	count := Count(3)
	WriteLog(count)

}
