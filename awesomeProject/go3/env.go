package go3

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func Env() {
	path, ok := os.LookupEnv("PATH")
	if !ok {
		fmt.Println("PATH env variable not set")
	} else {
		fmt.Println("PATH: " + path)
	}

	_ = os.Setenv("USERUSER", "MyEnvUser")
	user := os.Getenv("USERUSER")
	fmt.Println("USERUSER: " + user)

	/*	_ = os.Unsetenv("USERUSER")

		githubUsername, ok := os.LookupEnv("GITHUB_USERNAME")
		if ok {
			fmt.Println("githubUsername: ", githubUsername)
		}*/
}
