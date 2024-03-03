package newAdmin

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
)

func decodeBody(r *http.Response, out interface{}) error {
	return json.NewDecoder(r.Body).Decode(out)
}

func CreateSlug(input string) string {
	// Remove special characters
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		panic(err)
	}
	processedString := reg.ReplaceAllString(input, " ")

	// Remove leading and trailing spaces
	processedString = strings.TrimSpace(processedString)

	// Replace spaces with dashes
	slug := strings.ReplaceAll(processedString, " ", "-")

	// Convert to lowercase
	slug = strings.ToLower(slug)

	return slug
}
