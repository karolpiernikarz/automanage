package altapay

import (
	"fmt"
	"github.com/beevik/etree"
	"io"
	"net/http"
)

var settings struct {
	username string
	password string
	baseUrl  string
}

func SetCredentials(username, password, baseUrl string) {
	settings.username = username
	settings.password = password
	settings.baseUrl = baseUrl
}

func GetTerminals() (terminals []string, err error) {
	// create a new request without running it
	req, err := http.NewRequest("GET", settings.baseUrl+"/merchant/API/getTerminals", nil)
	if err != nil {
		fmt.Println(err)
	}
	// set the basic auth header
	req.SetBasicAuth(settings.username, settings.password)
	// send the request using http client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	// read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	// convert the response body to string
	response := string(body)
	doc := etree.NewDocument()
	if err = doc.ReadFromString(response); err != nil {
		panic(err)
	}

	if doc.FindElement("//Header/ErrorCode").Text() != "0" {
		return nil, fmt.Errorf("error code: %s", doc.FindElement("//Header/ErrorCode").Text())
	}

	terminalsElement := doc.FindElement("//Body/Terminals")
	for _, v2 := range terminalsElement.SelectElements("Terminal") {
		// append terminal title
		terminals = append(terminals, v2.SelectElement("Title").Text())
	}
	return
}
