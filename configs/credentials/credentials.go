package credentials

var AUTH_USERNAME string
var AUTH_PASSWORD string
var AUTH_KEY string = "UNAUTORIZED"

// TODO: MAke AUTH_KEY private and create access and validation functions

func ValidLoginCredentials(user string, pass string) bool {
	return user == AUTH_USERNAME && pass == AUTH_PASSWORD
}

// If returns true means that the request is authenticated
func ValidAuthKey(key string) bool {
	return key == AUTH_KEY
}
