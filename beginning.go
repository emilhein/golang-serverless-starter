package main

import (
	"fmt"
	"math/rand"
	"time"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"encoding/json"


)
const S3_REGION = "eu-west-1"
const S3_BUCKET = "ehe-development"

func main() {

	// someMap := map[string]int{"Food": 1, "music": 2}
	// printKeysAndValues(someMap)
	
	// if result, message, err := computeTotal(5,10); err != nil {
	// 	fmt.Println("We got a big problem")
	// } else {
	// 	fmt.Println("We are perfect")
	// 	fmt.Printf("Sum is %v and message is %s \n", result, message)

	// }
	
	// myMovie := movie{title: "Avengers", year: 2018, rating: "7.9"}
	// myMovie.format()

	sess, err := session.NewSession(&aws.Config{Region: aws.String(S3_REGION)})
    if err != nil {
        // Handle error
    }

    handler := S3Handler{
        Session: sess,
        Bucket:  S3_BUCKET,
    }
   contents, err := handler.ReadFile("movies.json")
    if err != nil {
		fmt.Println(err) 
        // Handle error
    }

    fmt.Println(contents) // "This is a test file"

}


func printKeysAndValues(themap map[string]int) string {
    for index, num := range themap {
		fmt.Printf("Index is %s, value is %v \n", index, num)
    }
	return "ok"
}
func computeTotal(a int, b int) (int, string, error) {
	rand.Seed(time.Now().UnixNano())
	min := 0
    max := 10
	randomN := rand.Intn(max - min) + min
	fmt.Printf("The random number is :- %v \n", randomN)
	if randomN >= 5 {
		return a+b, "Everything ok we have a high number ", nil
	} else {
		return a+b, "Everything bad ", errors.New("We dont work with small numbers")
	}
}