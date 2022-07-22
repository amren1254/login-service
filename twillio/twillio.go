package twillio

import (
	"fmt"
	"io/ioutil"
	"login-service/entity"
	"net/http"
	"strings"
)

func InitTwillio(t entity.SendOTP) {

	url := "https://verify.twilio.com/v2/Services/VAfe3d572648bfa354e9fa4d0276c3768b/Verifications"
	method := "POST"

	payload := strings.NewReader("To=%2B918896726484&Channel=sms")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Basic QUMwNTFkNjdiYjYyYThiOTUwOTgzMzEyNjNiOWJlZjRmZTo3YWE1YTIxZjViMDg2MDk3ZjYyNzg2OWJhMzBjOTA2ZA==")
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
