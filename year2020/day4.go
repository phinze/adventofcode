package year2020

import (
	"bufio"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Height string

type Passport struct {
	BirthYear      int    `passport:"byr"`
	IssueYear      int    `passport:"iyr"`
	ExpirationYear int    `passport:"eyr"`
	Height         string `passport:"hgt"`
	HairColor      string `passport:"hcl"`
	EyeColor       string `passport:"ecl"`
	PassportID     string `passport:"pid"`
	CountryID      string `passport:"cid"`
}

func (p Passport) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.BirthYear, validation.Required,
			validation.Min(1920), validation.Max(2002)),
		validation.Field(&p.IssueYear, validation.Required,
			validation.Min(2010), validation.Max(2020)),
		validation.Field(&p.ExpirationYear, validation.Required,
			validation.Min(2020), validation.Max(2030)),
		validation.Field(&p.Height, validation.Required,
			validation.Match(regexp.MustCompile("^[0-9]$"))),
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
	NumValid int
}

func parseDayFourLine(p *Passport, line string) {
	v := reflect.ValueOf(p).Elem()

	for _, field := range strings.Split(line, " ") {
		keyValue := strings.Split(field, ":")
		key := keyValue[0]
		value := keyValue[1]
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldType := v.Type().Field(i)
			tag := fieldType.Tag.Get("passport")
			if key == tag {
				switch fieldType.Type.Kind() {
				case reflect.Int:
					num, err := strconv.Atoi(value)
					if err != nil {
						panic(fmt.Sprintf("cannot convert to number: %#v", err))
					}
					field.Set(reflect.ValueOf(num))
				case reflect.String:
					field.Set(reflect.ValueOf(value))
				default:
					panic(fmt.Sprintf("unexpected field type: %#v", fieldType))
				}
			}
		}
	}
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
		if err := p.Validate(); err != nil {
			log.Printf("invalid: %s", err)
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
