package main

import "fmt"

func main() {
    // Bio of Bill Gates
    name := "Bill Gates"
    birthYear := 1955
    birthPlace := "Seattle, Washington, United States"
    education := "Harvard University (dropped out)"
    spouse := "Melinda Gates"
    children := 3
    occupation := "Co-founder of Microsoft, philanthropist"
    netWorth := 130 * 1000000000 // in USD
    company := "Microsoft Corporation"
    residence := "Medina, Washington, United States"
    nationality := "American"
    awards := []string{
        "Presidential Medal of Freedom",
        "Padma Bhushan",
        "Honorary Knight Commander of the Order of the British Empire",
    }
    philanthropyInitiatives := []string{
        "Bill & Melinda Gates Foundation",
        "Gates Foundation's Global Health Program",
        "Gates Foundation's Education Program",
    }
    notableWorks := []string{
        "Windows operating system",
        "Microsoft Office suite",
        "Co-author of 'The Road Ahead'",
    }

    // Print bio using different format verbs
    fmt.Println("Biography of", name)
    fmt.Printf("Name: %s\n", name)
    fmt.Printf("Birth Year: %d\n", birthYear)
    fmt.Printf("Place of Birth: %s\n", birthPlace)
    fmt.Printf("Education: %s\n", education)
    fmt.Printf("Spouse: %s\n", spouse)
    fmt.Printf("Children: %d\n", children)
    fmt.Printf("Occupation: %s\n", occupation)
    fmt.Printf("Net Worth: $%d\n", netWorth)
    fmt.Printf("Co-founded Company: %s\n", company)
    fmt.Printf("Residence: %s\n", residence)
    fmt.Printf("Nationality: %s\n", nationality)
    fmt.Println("Awards:")
    for _, award := range awards {
        fmt.Printf("- %s\n", award)
    }
    fmt.Println("Philanthropy Initiatives:")
    for _, initiative := range philanthropyInitiatives {
        fmt.Printf("- %s\n", initiative)
    }
    fmt.Println("Notable Works:")
    for _, work := range notableWorks {
        fmt.Printf("- %s\n", work)
    }
}
