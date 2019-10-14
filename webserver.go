package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

const S3_REGION = "eu-west-1"

const S3_BUCKET = "ehe-development"

var wg sync.WaitGroup

func finder(mine [6]string, oreChannel chan string) {
	defer wg.Done()
	for _, value := range mine {
		if value == "ore" {
			oreChannel <- value //send item on oreChannel
		}
	}
	close(oreChannel)

}
func breaker(oreChannel <-chan string, minedOreChan chan<- string) {
	defer wg.Done()
	for elem := range oreChannel {
		fmt.Println("From Finder: ", elem)
		minedOreChan <- "minedOre" //send to minedOreChan
	}
	close(minedOreChan)

}
func smelter(minedOreChan <-chan string, name string, maxFound int) {
	defer wg.Done()
	for minedOre := range minedOreChan { //read from minedOreChan by ranging
		fmt.Println("From Miner in FUNCTION: ", minedOre)
		fmt.Printf("From Smelter (%s): Ore is smelted \n", name)
	}
}
func StartMining(w http.ResponseWriter, r *http.Request) {
	wg.Add(3)
	theMine := [6]string{"rock", "ore", "ore", "rock", "ore", "ore"}
	oreChannel := make(chan string)

	minedOreChan := make(chan string)
	// Finder
	go finder(theMine, oreChannel)
	// Ore Breaker
	go breaker(oreChannel, minedOreChan)

	// Smelters
	go smelter(minedOreChan, "Bob", len(theMine))

	wg.Wait()
	fmt.Println("Main: Completed")
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])

}
func GetS3File(w http.ResponseWriter, r *http.Request) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(S3_REGION)})
	if err != nil {
		// Handle error
	}

	handler := S3Handler{
		Session: sess,
		Bucket:  S3_BUCKET,
	}
	fileString := "movies.json"
	contents, err := handler.ReadFile(fileString)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(contents) // "This is a test file"
	toJson, error := extractMovieData(contents)
	if err != nil {
		fmt.Println(error)
	}
	for _, p := range toJson {
		log.Printf("Name: %s , adsense: %s \n", p.Title, p.Rating)
	}

	fmt.Fprintf(w, "Check the log. you")

}
func Simple(w http.ResponseWriter, r *http.Request) {
	someMap := map[string]int{"Food": 1, "music": 2}
	printKeysAndValues(someMap)

	if result, message, err := computeTotal(5, 10); err != nil {
		fmt.Println("We got a big problem")
	} else {
		fmt.Println("We are perfect")
		fmt.Printf("Sum is %v and message is %s \n", result, message)

	}

	myMovie := Movie{Title: "Avengers", Year: 2018, Rating: "7.9"}
	myMovie.format()

	// fmt.Println(toJson) // "This is a test file"

	fmt.Fprintf(w, "Check the log. Simple things has been written")

}
func main() {
	http.HandleFunc("/startmining", StartMining)
	http.HandleFunc("/getS3file", GetS3File)
	http.HandleFunc("/simple", Simple)
	http.ListenAndServe(":8080", nil)

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
	randomN := rand.Intn(max-min) + min
	fmt.Printf("The random number is :- %v \n", randomN)
	if randomN >= 5 {
		return a + b, "Everything ok we have a high number ", nil
	} else {
		return a + b, "Everything bad ", errors.New("We dont work with small numbers")
	}
}
