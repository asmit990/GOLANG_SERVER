package main

import (
	"fmt"
	
	"net/url"
)

const myurl string = "https://lco.dev:3000?learn=true&coursename=reactjs"

func main() {
	fmt.Println("Welcome to URL Handling")
	fmt.Println(myurl)

	// Parsing
	result, err := url.Parse(myurl) // ✅ Add error handling
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("Scheme:", result.Scheme)   // ✅ Fixed capitalization
	fmt.Println("Host:", result.Host)
	fmt.Println("Port:", result.Port())
	fmt.Println("Raw Query:", result.RawQuery)

	qparams := result.Query() // ✅ Extract query parameters
	fmt.Printf("The type of query params is: %T\n", qparams)

	// Access query parameters
	fmt.Printf("Coursename:", qparams["coursename"])
}
