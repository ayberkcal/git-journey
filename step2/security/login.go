package security

import (
	"fmt"
)

// Login is a function to auth a user with 'username'
func Login(username string) {
	currentUser = username
	fmt.Printf("login successfully '%s'", username)
}
