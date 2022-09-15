package gooremote_test

import (
	"os"
	"testing"

	gooremote "github.com/blackestwhite/go-oremote"
	"github.com/joho/godotenv"
)

func TestNewPayment(t *testing.T) {
	godotenv.Load()
	accessKey := os.Getenv("ACCESS_KEY")

	i := gooremote.New("1234567890")
	_, err := i.NewPayment(10000, "https://oremote.org", "", "test")
	if err == nil {
		t.Error("error should not be nil, oops: ", err)
	}

	i = gooremote.New(accessKey)
	_, err = i.NewPayment(10000, "https://oremote.org", "", "test")
	if err != nil {
		t.Error("error should be nil, err: ", err)
	}
}

func TestVerify(t *testing.T) {
	godotenv.Load()
	accessKey := os.Getenv("ACCESS_KEY")

	i := gooremote.New("1234567890")
	paid, err := i.Verify("fewgwrwgreg")
	if err == nil {
		t.Error("error should not be nil, oops: ", err)
	}
	if paid {
		t.Error("paid should be false")
	}

	i = gooremote.New(accessKey)
	id, _ := i.NewPayment(10000, "https://oremote.org", "", "test")
	paid, err = i.Verify(id)
	if err != nil {
		t.Error(id, err.Error())
	}
	if paid {
		t.Error("paid should be false")
	}
}
