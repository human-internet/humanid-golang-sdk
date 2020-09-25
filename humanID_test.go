package humanID

import (
	"testing"
)

func TestLogin(t *testing.T) {
	humanID := New(
		"SERVER_EN72KF8WEJ94B05K87383P",
		"NXJCuY14QDdsiHp9L6hDDs2im6-6Wuz.2p4kcvoDXg3g9tTA93lKkuALowi.fbT5",
	)
	loginResp, err := humanID.Login("MOBILE_TF3E9M4K85O8HPBY68WE6Z", "1GGUIZSQ4-Iojleyp9-w88jogckA3gsNkHU711pc~uqsW~dfyOfb-knrPYhcewsw")

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