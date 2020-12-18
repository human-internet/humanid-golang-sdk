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

// av9WfTOwZtnc6ibLaLffsSmA%2FJJ5av7pNZJergIE%2BtOW0SqBRdQhiuIakPMMphvv%2BgoX%2FQQ8gWnq9OKwZjYlQdts3fGEzrSCX48foKw7j5dO1QRhXNpdENrKS%2B8C3ziASiHzox8Wg3cJ162KJlzPjr7wz2P%2BCw3Akg40PymguJCTUWQ%3D%3D
func TestVerifyToken(t *testing.T) {
	humanID := New(
		"SERVER_EN72KF8WEJ94B05K87383P",
		"NXJCuY14QDdsiHp9L6hDDs2im6-6Wuz.2p4kcvoDXg3g9tTA93lKkuALowi.fbT5",
	)
	verifyTokenResp, err := humanID.VerifyToken("zGlKvSdg5oelDP5gYklAElD1%2FxaG0%2BpeKuSWDRRmYI3TWsrxo9ROvKYGW6YZTn9itu4Oo%2FVpmCMYRP05ng5KJAcagGjdxOBJd5WKtcxLBWmDgAxVYm0dcIGxI5xV8sJYFC%2BD7V6GDLm0JwRy7PLu9h37HzU7XIxADPBm%2B85b0YoJS%2Bw%3D%3D")

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