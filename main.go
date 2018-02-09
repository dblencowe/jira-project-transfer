package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"jira-project-export/jira"
	"time"
	"github.com/joho/godotenv"
)

func main() {
	start := time.Now()

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	completeas := jira.NewInstance(os.Getenv("SOURCE_INSTANCE_URL"), os.Getenv("SOURCE_INSTANCE_USER"), os.Getenv("SOURCE_INSTANCE_PASS"))
	response := completeas.Issues.List("CLOSEIT")
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Response Body: ", string(body))

	elapsed := time.Since(start)
	fmt.Printf("[msg] Execution took %s", elapsed)
}
