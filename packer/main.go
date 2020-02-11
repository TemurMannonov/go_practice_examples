package main

import (
	"fmt"
	str "strings" // Package Alias
	"rsc.io/quote"
	"github.com/TemurMannonov/packer/numbers"
	"github.com/TemurMannonov/packer/strings"
	"github.com/TemurMannonov/packer/strings/greeting" // Importing a nested package
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println(numbers.IsPrime(19))

	fmt.Println(greeting.WelcomeText)

	fmt.Println(strings.Reverse("callicoder"))

	fmt.Println(str.Count("Go is Awesome. I love Go", "Go"))
}
