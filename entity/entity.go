package entity

type UserProfile struct {
	FullName    string `json:"fullname"`
	EmailId     string `json:"emailid"`
	PhoneNumber string `json:"phonenumber"`
}

type OTP struct {
	PhoneNumber string `json:"phonenumber"`
	Otp         string `json:"otp"`
}
type SendOTP struct {
	PhoneNumber string `json:"phonenumber"`
}
