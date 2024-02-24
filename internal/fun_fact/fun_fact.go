package fun_fact

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Fact struct {
	Id        string `json:"id"`
	Text      string `json:"text"`
	Source    string `json:"source"`
	SourceUrl string `json:"source_url"`
	Language  string `json:"language"`
	Permalink string `json:"permalink"`
}

func FunFact() (string, error) {

	url := "https://uselessfacts.jsph.pl/api/v2/facts/today"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var fact Fact
	err := json.Unmarshal(body, &fact)
	if err != nil {
		fmt.Println("error:", err)
	}

	return fact.Text, nil
}
