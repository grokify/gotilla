package urlutil

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// ToSlug creates a slug byte array from an input byte array.
// Slugs have words separated by a hyphen with no punctuation
// or spaces.
func ToSlug(slug []byte) []byte {
	// Convert punctuation and spaces to hyphens: string([]byte{45}) = "-"
	slug = regexp.MustCompile(`[\*\s]+`).ReplaceAll(slug, []byte{45})
	return regexp.MustCompile(`(^-+|-+$)`).ReplaceAll(slug, []byte{})
}

// ToSlugLowerString creates a lower-cased slug string
func ToSlugLowerString(s string) string {
	return string(ToSlug([]byte(strings.ToLower(s))))
}

func BuildURLFromMap(baseUrl string, queryParams map[string]string) string {
	if len(queryParams) < 1 {
		return baseUrl
	}
	queryValues := url.Values{}
	for key, val := range queryParams {
		queryValues.Set(key, val)
	}
	return BuildURL(baseUrl, queryValues)
}

func BuildURL(baseUrl string, queryValues url.Values) string {
	qryString := queryValues.Encode()
	if len(qryString) > 0 {
		return baseUrl + "?" + qryString
	}
	return baseUrl
}

func GetURLBody(absoluteUrl string) ([]byte, error) {
	req, err := http.NewRequest("GET", absoluteUrl, nil)
	cli := &http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func GetURLPostBody(absoluteUrl string, bodyType string, reqBody io.Reader) ([]byte, error) {
	client := &http.Client{}
	res, err := client.Post(absoluteUrl, bodyType, reqBody)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
