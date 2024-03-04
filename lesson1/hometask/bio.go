package main

import "fmt"

func main() {
    // Bio of Bill Gates
    name := "Bill Gates"
    birthYear := 1955
    occupation := "Co-founder of Microsoft"
    netWorth := 130 * 1000000000 // in USD

    // Print bio using different format verbs
    fmt.Println("Biography of", name)
    fmt.Printf("Name: %s\n", name)
    fmt.Printf("Birth Year: %d\n", birthYear)
    fmt.Printf("Occupation: %s\n", occupation)
    fmt.Printf("Net Worth: $%d\n", netWorth)
}
