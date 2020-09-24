package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Date time.Time

func (j *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = Date(t)
	return nil
}

func (j Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(j)
}
func (j Date) Format() string {
	t := time.Time(j)
	return t.Format("2006-01-02")
}

type Payload struct {
	fields  []string
	data    map[string]string
	field   string
	errors  map[string]string
	regexps map[string]*regexp.Regexp
}

func ParamID(c *gin.Context, name string) uint64 {
	id, _ := strconv.ParseUint(c.Params.ByName(name), 10, 64)
	return id
}

func StringExists(str string, array []string) bool {
	for _, v := range array {
		if v == str {
			return true
		}
	}
	return false
}

func NewPayload(c *gin.Context, fields ...string) *Payload {
	p := &Payload{}
	p.data = make(map[string]string)
	p.errors = make(map[string]string)
	p.regexps = make(map[string]*regexp.Regexp)
	c.BindJSON(&p.data)
	fmt.Printf("%v \n", p.data)
	return p
}

func (p *Payload) addError(format string, a ...interface{}) {
	p.errors[p.field] = fmt.Sprintf(format, a...)
}

func (p *Payload) Field(field string, isOptional bool) *Payload {
	p.fields = append(p.fields, field)
	p.field = field

	return p
}

func (p *Payload) Length(minLength int, maxLength int) *Payload {
	p.MinLength(minLength)
	p.MaxLength(maxLength)
	return p
}

func (p *Payload) MinLength(minLength int) *Payload {
	if len(p.data[p.field]) < minLength {
		p.addError("%s should contain a minimum of %d chars", p.field, minLength)
	}
	return p
}
func (p *Payload) MaxLength(maxLength int) *Payload {
	if len(p.data[p.field]) > maxLength {
		p.addError("%s should be no longer than %d chars", p.field, maxLength)
	}
	return p
}

func (p *Payload) RegExp(re string, message string) *Payload {
	if _, ok := p.regexps[re]; !ok {
		p.regexps[re] = regexp.MustCompile(re)
	}
	if !p.regexps[re].MatchString(p.data[p.field]) {
		p.addError(message)
	}
	return p
}

func (p *Payload) IsValid(ignoreExtraFields bool) bool {
	if !ignoreExtraFields {
		for fk, _ := range p.data {
			for _, f := range p.fields {
				if f != fk {
					continue
				}
				p.field = fk
				p.addError("unknown field")
			}

		}
	}
	return len(p.errors) < 1
}
func (p *Payload) Errors() *map[string]string {
	if len(p.errors) < 1 {
		return nil
	}
	return &p.errors
}

func parseJson(c *gin.Context, validKeys ...string) (error, map[string]string) {
	var p map[string]string
	if err := c.BindJSON(&p); err != nil {
		return errors.New("invalid request body; " + err.Error()), p
	}

	if len(p) < 1 {
		return errors.New("cannot find any valid json fields"), p
	}

	for f, _ := range p {
		if !StringExists(f, validKeys) {
			return errors.New(fmt.Sprintf("invalid field %s", f)), p
		}
	}
	return nil, p
}
