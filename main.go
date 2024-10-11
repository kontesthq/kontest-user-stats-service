package main

import (
	"fmt"
	"kontest-user-stats-service/service"
)

func main() {
	fmt.Println("Hello, World!")

	codeChefService := service.NewCodeChefService()

	userData, err := codeChefService.GetUserData("ayushs_2k4")

	fmt.Println()
	fmt.Println()
	fmt.Println()

	if err != nil {
		fmt.Println("Error in fetching codechef user data: ", err)
		return
	}

	fmt.Println(userData)
}
