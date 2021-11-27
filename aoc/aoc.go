package aoc

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/imroc/req"
)

const inputUrl = "https://adventofcode.com/%d/day/%d/input"

func FetchInput(year int, day int) (string, error) {
	url := fmt.Sprintf(inputUrl, year, day)

	session, err := ioutil.ReadFile("../session.cookie")
	if err != nil {
		return "", err
	}
	cookie := &http.Cookie{Name: "session", Value: strings.TrimSpace(string(session))}

	resp, err := req.Get(url, cookie)
	if err != nil {
		return "", err
	}

	return resp.ToString()
}
