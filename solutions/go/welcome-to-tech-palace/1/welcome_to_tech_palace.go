package techpalace

import "strings"
import "fmt"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    removeStarts := strings.Trim(oldMsg, "*")
    fmt.Println(removeStarts)
    removeStarts = strings.TrimSpace(removeStarts)
    fmt.Println(removeStarts)
    removeStarts = strings.TrimLeft(removeStarts, "* ")
    fmt.Println(removeStarts)
    // removeStarts = strings.TrimRight(removeStarts, " ")
    removeStarts = strings.TrimRight(removeStarts, " *")
    // removeStarts = removeStarts + "aaaa"
    fmt.Println(removeStarts)
	return removeStarts
}
