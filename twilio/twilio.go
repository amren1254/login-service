package twilio

import (
	"fmt"
	"io/ioutil"
	"login-service/entity"
	"net/http"
	"os"
	"strings"
)

var (
	//phonenum      = os.Getenv("PHONE_NUM")
	authorization = os.Getenv("AUTH_ID")

//	code          = os.Getenv("CODE")
)

func InitTwilio(t entity.SendOTP) {
	//	phonenum := os.Getenv("PHONE_NUM")
	authorization := os.Getenv("AUTH_ID")
	url := "https://verify.twilio.com/v2/Services/VAf5b4688f3feb22783cc12af5e0391d3a/Verifications"
	method := "POST"

	payload := strings.NewReader(t.PhoneNumber)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", authorization)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	fmt.Println("otp send to phone number")
}
func VerifiyOtp(t entity.SendOTP) {

	url := "https://verify.twilio.com/v2/Services/VAf5b4688f3feb22783cc12af5e0391d3a/VerificationCheck"
	method := "POST"

	payload := strings.NewReader(t.Code)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", authorization)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
