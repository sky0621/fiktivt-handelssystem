package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const endpoint = "https://api.github.com/graphql"

func main() {
	ghToken := os.Getenv("FOR_GRAPHQL_SAMPLE_GITHUB_TOKEN")

	query := `
		query($login: String!) {
			organization(login: $login) {
				name
				email
			}
		}
	`

	b, err := json.Marshal(struct {
		Query string `json:"query"`
		Variable map[string]interface{} `json:"variables"`
	}{
		Query:query,
		Variable:map[string]interface{}{
			"login": "github",
		},
	})
	if err != nil {
		panic(err)
	}

	epURL, err := url.Parse(endpoint)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer(b)
	res, err := http.DefaultClient.Do(&http.Request{
		URL: epURL,
		Method:"POST",
		Header:http.Header{
			"Content-Type": {"application/json"},
			"Authorization": {"bearer " + ghToken},
		},
		Body:ioutil.NopCloser(buf),
	})
	if err != nil {
		panic(err)
	}

	bd, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bd))
}
