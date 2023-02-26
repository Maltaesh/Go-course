package app2

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	_, err := IsEmail("hello")

	if err == nil {
		t.Error("hello is not an email")
	}

	_, err = IsEmail("marcin@em.pl")

	if err != nil {
		t.Error("marcin@em.pl is an email")
	}

	_, err = IsEmail("marcin@em")

	if err != nil {
		t.Error("marcin@em is not an email")
	}
}
