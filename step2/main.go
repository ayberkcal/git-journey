package main

import "git-journey/step2/security"
import "git-journey/step2/recipe-manager/recipe"
import "fmt"

const username = "Ayberk"

func main() {
	if !security.IsAuthenticated(username) {
		fmt.Printf("%s is not authenticated, logging \n", username)
		security.Login(username)
	}

	recipe.Add("Pizza")
	recipe.Add("Cake")
	recipe.Add("Sushi")
	recipe.Add("Cake")

	security.Logout()

	if security.IsAuthenticated(username) {
		fmt.Println("ERROR - User strill authenticated")
	}
}
