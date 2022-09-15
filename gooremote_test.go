package gooremote_test

import (
	"testing"

	gooremote "github.com/blackestwhite/go-oremote"
)

func TestNewPayment(t *testing.T) {
	i := gooremote.New("1234567890")
	_, err := i.NewPayment(10000, "https://oremote.org", "", "test")
	if err == nil {
		t.Error("error should not be nil, oops: ", err)
	}
}

func TestVerify(t *testing.T) {
	i := gooremote.New("1234567890")
	paid, err := i.Verify("fewgwrwgreg")
	if err == nil {
		t.Error("error should not be nil, oops: ", err)
	}

	if paid {
		t.Error("paid should be false")
	}
}
