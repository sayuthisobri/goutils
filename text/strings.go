package text

import (
	"regexp"
	"strings"
)

type Strings struct {
	builder strings.Builder
}

func NewString(str string) *Strings {
	s := Strings{}

	_, _ = s.builder.WriteString(str)
	return &s
}

func (s *Strings) Append(str string) *Strings {
	s.builder.WriteString(str)
	return s
}

func (s *Strings) Set(str string) *Strings {
	s.builder.Reset()
	s.builder.WriteString(str)
	return s
}
func (s Strings) toString() string {
	return s.builder.String()
}

func (s *Strings) Split(regex string) []string {
	r := regexp.MustCompile(regex)
	return r.Split(s.builder.String(), -1)
}

func (s *Strings) Len() int {
	return s.builder.Len()
}

func (s *Strings) toUpper() *Strings {
	s.Set(strings.ToUpper(s.toString()))
	return s
}

func (s *Strings) toLower() *Strings {
	s.Set(strings.ToLower(s.toString()))
	return s
}

func (s *Strings) EqualIgnoreCase(str string) bool {
	return strings.ToUpper(s.toString()) == strings.ToUpper(str)
}

func (s *Strings) Match(regex regexp.Regexp) bool {
	return regex.MatchString(s.toString())
}

func (s *Strings) ToSnakeCase() string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(s.toString(), "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
