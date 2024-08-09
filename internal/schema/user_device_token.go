package schema

type UserDeviceTokenCreateRequest struct {
	UserId string `json:"-"`
	Token  string `json:"token"`
}
