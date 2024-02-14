package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	fullMsg := ""
	for i := 0; i < numStarsPerLine; i++ {
		fullMsg += "*"
	}

	fullMsg += "\n" + welcomeMsg + "\n"

	for i := 0; i < numStarsPerLine; i++ {
		fullMsg += "*"
	}

	return fullMsg
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	cleanMsg := ""
	for _, char := range oldMsg {
		if char != '*' {
			cleanMsg += string(char)
		}
	}

	// Remove leading and trailing whitespace
	return strings.TrimSpace(cleanMsg)
}
