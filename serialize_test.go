package gogotest_test

import (
	"encoding/json"
	"fmt"
	"testing"

	gogo "github.com/buckhx/gogo-test/gen/gogo"
	"github.com/gogo/protobuf/proto"
	//golang "github.com/buckhx/gogo-test/gen/golang"
)

var subjects = []struct {
	name    string
	email   string
	surname string
	url     string
}{
	{"ash", "ash@pallet.gov", "ketchum", ""},
	{"silph", "contact@silph.co", "", "https://silph.co"},
}

func TestGogoRoundTrip(t *testing.T) {
	for _, test := range subjects {
		in := &gogo.Subject{
			Label:   test.name,
			Contact: &gogo.Subject_Contact{Email: test.email},
		}
		if test.surname != "" {
			name := &gogo.Person_Name{First: test.name, Last: test.surname}
			in.Entity = &gogo.Subject_Person{&gogo.Person{Name: name}}
		} else {
			in.Entity = &gogo.Subject_Application{&gogo.Application{Name: test.name, Url: test.url}}
		}
		str, _ := json.Marshal(in)
		fmt.Println(string(str))
		buf, err := proto.Marshal(in)
		if err != nil {
			t.Error(err)
		}
		out := &gogo.Subject{}
		if err := proto.Unmarshal(buf, out); err != nil {
			t.Error(err)
		}
		str, _ = json.Marshal(out)
		fmt.Println(string(str))
		switch {
		case test.name != out.Label:
			t.Errorf("Subject.Label didn't match")
		case test.email != out.Contact.Email:
			t.Errorf("Subject.Contact.Email didn't match")
		case test.surname != "" && test.surname != out.GetPerson().Name.Last:
			t.Errorf("Subject.Person.Name.Last didn't match")
		case test.url != "" && test.url != out.GetApplication().Url:
			t.Errorf("Subject.Application.Url didn't match")
		}
	}
}
