package mailer

import (
	"Repo-watcher/src/test/factory"
	"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	c := testFactory.NewConfig()

	m := New(c)
	fmt.Printf("m: %+v, c: %+v\n", *m, *c)
	if c.Mail.From != m.From {
		t.Errorf("two from is not same")
		t.Errorf("Expect: %v, got: %v", c.Mail.From, m.From)
	}
	if !reflect.DeepEqual(c.Mail.To, m.To) {
		t.Errorf("two to is not same")
		t.Errorf("Expect: %v, got: %v", c.Mail.To, m.To)
	}
	if c.Mail.Username != m.Username {
		t.Errorf("two username is not same")
		t.Errorf("Expect: %s, got: %s", c.Mail.Username, m.Username)
	}
	if c.Mail.Password != m.Password {
		t.Errorf("two password is not same")
		t.Errorf("Expect: %s, got: %s", c.Mail.Password, m.Password)
	}
}
