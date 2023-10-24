package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// MovieDb represents a movie entry in the database.
type MovieDb struct {
	Title    string
	Id       int
	Year     int
	Director string
}

var movieSlice []MovieDb

// create is used to add a new movie to the database.
func create() {
	fmt.Println("Please Enter the Name, Year, and Director of the Movie you want to add, or enter 'quit' to quit")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		var err error

		// Input movie title
		fmt.Print("Enter movie title: ")
		scanner.Scan()
		input := scanner.Text()

		// Input movie year
		fmt.Print("Enter movie year: ")
		scanner.Scan()
		inyearStr := scanner.Text()

		inyear, err := strconv.Atoi(inyearStr)
		if err != nil {
			fmt.Println("Invalid year input. Please enter a valid integer.")
			continue
		}

		// Input movie director
		fmt.Print("Enter movie director: ")
		scanner.Scan()
		indirect := scanner.Text()

		movie := MovieDb{
			Title:    input,
			Id:       len(movieSlice) + 1,
			Year:     inyear,
			Director: indirect,
		}
		movieSlice = append(movieSlice, movie)
		fmt.Println("Movie created successfully")

		jsonBytes, err := json.Marshal(&movieSlice)
		if err != nil {
			log.Println(err)
		}

		err = ioutil.WriteFile("data/movie.json", jsonBytes, 0644)
		if err != nil {
			log.Println(err)
		}

		fmt.Print("Do you want to quit or enter another movie?\n")
		fmt.Print("Enter 'quit' to quit\n")
		fmt.Print("Press Enter to add another movie\n")
		scanner.Scan()
		exit := scanner.Text()
		if exit == "quit" {
			break
		} else {
			continue
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
}

// movielist lists all movies in the database.
func movielist() (err error) {
	jsonBytes, err := ioutil.ReadFile("data/movie.json")
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(jsonBytes, &movieSlice)

	if len(movieSlice) == 0 {
		return errors.New("The database is empty")
	}

	for _, movies := range movieSlice {
		fmt.Printf("ID: %v, Title: %v, Year: %v, Director: %v\n", movies.Id, movies.Title, movies.Year, movies.Director)
	}
	return nil
}

// moviesearch searches for a movie in the database by title.
func moviesearch() {
	var param string
	fmt.Print("Enter the movie title to search: ")
	fmt.Scanln(&param)
	found := false

	for _, movies := range movieSlice {
		if param == movies.Title {
			fmt.Printf("Search result: ID: %v, Title: %v, Year: %v, Director: %v\n", movies.Id, movies.Title, movies.Year, movies.Director)
			found = true
		}
	}

	if !found {
		fmt.Println("No search results found")
	}
}

func main() {
	var input string
	var err error

myLoop:
	for {
		fmt.Println("\nWelcome to your Movie Database")
		fmt.Println("Please enter a command to continue")
		fmt.Println("List, Create, Search, Quit")
		fmt.Scanln(&input)

		switch input {
		case "list":
			err = movielist()
			if err != nil {
				fmt.Println(err)
			}
		case "create":
			create()
		case "search":
			moviesearch()
		case "quit":
			break myLoop
		default:
			fmt.Println("Not a valid command, please try again")
		}
	}
}
