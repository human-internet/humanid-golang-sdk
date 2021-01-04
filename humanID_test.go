package humanID

import (
	"testing"
	"os"
)

func TestLogin(t *testing.T) {
	humanID := New(
		os.Getenv("SERVER_ID"),
		os.Getenv("SERVER_SECRET"),
	)
	loginResp, err := humanID.Login(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))

	if err != nil {
		t.Error(err.Error())
		return
	}

	if loginResp.Code == "401" {
		t.Error("Request unauthorized:", loginResp)
		return
	}

	if !loginResp.Success {
		t.Error("Request failed:", loginResp)
		return
	}

	t.Log("Response:", loginResp)
}

func TestVerifyToken(t *testing.T) {
	humanID := New(
		os.Getenv("SERVER_ID"),
		os.Getenv("SERVER_SECRET"),
	)
	verifyTokenResp, err := humanID.VerifyToken("INPUT_HERE")

	if err != nil {
		t.Error(err.Error())
		return
	}

	if verifyTokenResp.Code == "401" {
		t.Error("Request unauthorized:", verifyTokenResp)
		return
	}

	if !verifyTokenResp.Success {
		t.Error("Request failed:", verifyTokenResp)
		return
	}

	t.Log("Verify token response:", verifyTokenResp)
}