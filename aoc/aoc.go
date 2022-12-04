package aoc

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	req "github.com/imroc/req/v3"
)

const cacheDir = "../.aoc-input-cache"
const inputUrl = "https://adventofcode.com/%d/day/%d/input"
const userAgent = "phinze-aoc-fetcher/dev https://github.com/phinze/adventofcode/blob/main/aoc/aoc.go"

func CachePath(year, day int) string {
	return filepath.Join(cacheDir, strconv.Itoa(year), strconv.Itoa(day))
}

func IsCached(year, day int) bool {
	_, err := os.Stat(CachePath(year, day))
	return !errors.Is(err, os.ErrNotExist)
}

func CacheWrite(year, day int, result string) error {
	path := CachePath(year, day)

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	return os.WriteFile(path, []byte(result), 0644)
}

func CacheRead(year, day int) (string, error) {
	data, err := os.ReadFile(CachePath(year, day))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ReadThroughCache(year, day int, fetch func() (string, error)) (string, error) {
	result, err := fetch()
	if err != nil {
		return "", err
	}
	if err := CacheWrite(year, day, result); err != nil {
		return "", err
	}
	return result, nil
}

func FetchInput(year int, day int) (string, error) {
	if IsCached(year, day) {
		return CacheRead(year, day)
	} else {
		return ReadThroughCache(year, day, func() (string, error) {
			return GetInputFromServer(year, day)
		})
	}
}

func GetInputFromServer(year int, day int) (string, error) {
	// Setting user agent as requested here:
	//   https://reddit.com/r/adventofcode/comments/z9dhtd/please_include_your_contact_info_in_the_useragent/
	client := req.C().SetUserAgent(userAgent)
	url := fmt.Sprintf(inputUrl, year, day)

	// Need manually arranged session cookie grabbed from browser
	session, err := ioutil.ReadFile("../session.cookie")
	if err != nil {
		return "", err
	}
	cookie := &http.Cookie{
		Name:  "session",
		Value: strings.TrimSpace(string(session)),
	}

	resp, err := client.R().SetCookies(cookie).Get(url)
	if err != nil {
		return "", err
	}

	return resp.ToString()
}
