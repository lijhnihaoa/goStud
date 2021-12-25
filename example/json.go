package example

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}

func marshaling() []byte {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed : %s", err)
	}
	fmt.Printf("%s\n", data)

	data, err = json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("json marshling failed : %s", err)
	}
	fmt.Printf("%s\n", data)

	return data
}

func unmarshal(data []byte) {
	var titles []struct{ Title string }

	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("json unmarshaling failed : %s", err)
	}

	fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
}

const IssuesURL = "https://api.github.com/search/issues"

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Number   int
	HTMLURL  string `json:"html_url"`
	Title    string
	State    string
	User     *User
	CreateAt time.Time `json:"created_at"`
	Body     string
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q) //repo:golang/go is:open json decoder
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func StudJSON() {
	// data := marshaling()
	// unmarshal(data)

	result, err := SearchIssues(os.Args[1:]) //repo:golang/go is:open json decoder
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	beforeMouth := time.Now().Unix() - 30*24*60*60
	beforeOneYear := time.Now().Unix() - 365*24*60*60
	for _, item := range result.Items {
		if item.CreateAt.Unix() > beforeMouth {
			fmt.Printf("A: #%-5d %9.9s %.55s \t\t%v\n", item.Number, item.User, item.Title, item.CreateAt)
		} else if item.CreateAt.Unix() > beforeOneYear {

			fmt.Printf("B: #%-5d %9.9s %.55s \t\t%v\n", item.Number, item.User, item.Title, item.CreateAt)
		} else {

			fmt.Printf("C: #%-5d %9.9s %.55s \t\t%v\n", item.Number, item.User, item.Title, item.CreateAt)
		}
	}
}
