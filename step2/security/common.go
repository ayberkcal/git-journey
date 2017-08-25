package security

// package variable, private
var currentUser string

func IsAuthenticated(username string) bool {
	return username == currentUser
}
