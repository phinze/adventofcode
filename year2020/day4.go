package year2020

import (
	"bufio"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Height string

var heightRegexp = regexp.MustCompile("^([0-9]+)(in|cm)$")

func (h Height) Validate() error {
	err := validation.Validate(string(h),
		validation.Required,
		validation.Match(heightRegexp))
	if err != nil {
		return err
	}

	matches := heightRegexp.FindStringSubmatch(string(h))
	num, err := strconv.Atoi(matches[1])
	if err != nil {
		return err
	}
	var min, max int
	switch matches[2] {
	case "in":
		min = 59
		max = 76
	case "cm":
		min = 150
		max = 193
	default:
		return fmt.Errorf("unexpected height unit: %s", matches[2])
	}
	return validation.Validate(num,
		validation.Min(min),
		validation.Max(max),
	)
}

type Passport struct {
	BirthYear      int    `passport:"byr"`
	IssueYear      int    `passport:"iyr"`
	ExpirationYear int    `passport:"eyr"`
	Height         Height `passport:"hgt"`
	HairColor      string `passport:"hcl"`
	EyeColor       string `passport:"ecl"`
	PassportID     string `passport:"pid"`
	CountryID      string `passport:"cid"`
}

var validEyeColors = []interface{}{
	"amb",
	"blu",
	"brn",
	"gry",
	"grn",
	"hzl",
	"oth",
}

func (p Passport) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.BirthYear, validation.Required,
			validation.Min(1920), validation.Max(2002)),
		validation.Field(&p.IssueYear, validation.Required,
			validation.Min(2010), validation.Max(2020)),
		validation.Field(&p.ExpirationYear, validation.Required,
			validation.Min(2020), validation.Max(2030)),
		validation.Field(&p.Height),
		validation.Field(&p.HairColor, validation.Required,
			validation.Match(regexp.MustCompile("^#[0-9a-f]{6}$"))),
		validation.Field(&p.EyeColor, validation.Required,
			validation.In(validEyeColors...)),
		validation.Field(&p.PassportID, validation.Required,
			validation.Match(regexp.MustCompile("^[0-9]{9}$"))),
	)
}

func (p *Passport) IsValid() bool {
	v := reflect.ValueOf(p).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)
		if fieldType.Tag.Get("passport") == "cid" {
			continue
		}
		if field.String() == "" {
			return false
		}
	}
	return true
}

type DayFourInput struct {
	Passports []*Passport
}

type DayFourOutput struct {
	NumValid        int
	NumValidPartTwo int
}

func parseDayFourLine(p *Passport, line string) error {
	for _, field := range strings.Split(line, " ") {
		keyValue := strings.Split(field, ":")
		key := keyValue[0]
		value := keyValue[1]
		switch key {
		case "byr":
			num, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			p.BirthYear = num
		case "iyr":
			num, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			p.IssueYear = num
		case "eyr":
			num, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			p.ExpirationYear = num
		case "hgt":
			p.Height = Height(value)
		case "hcl":
			p.HairColor = value
		case "ecl":
			p.EyeColor = value
		case "pid":
			p.PassportID = value
		case "cid":
			p.CountryID = value
		default:
			return fmt.Errorf("unexpected key: %s", key)
		}
	}
	return nil
}

func parseDayFour(rawInput string) (*DayFourInput, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))
	in := &DayFourInput{}

	curPassport := &Passport{}
	for scanner.Scan() {
		line := scanner.Text()
		// empty lines mean next record
		if line == "" {
			in.Passports = append(in.Passports, curPassport)
			curPassport = &Passport{}
		} else {
			parseDayFourLine(curPassport, line)
		}
	}
	// append last record
	in.Passports = append(in.Passports, curPassport)

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

func solveDayFour(in *DayFourInput) (*DayFourOutput, error) {
	out := &DayFourOutput{}

	for _, p := range in.Passports {
		if p.IsValid() {
			out.NumValid++
		}
		if err := p.Validate(); err == nil {
			out.NumValidPartTwo++
		}
	}

	return out, nil
}

func DayFour(rawInput string) (*DayFourOutput, error) {
	in, err := parseDayFour(rawInput)
	if err != nil {
		return nil, err
	}

	out, err := solveDayFour(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}
