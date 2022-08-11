package profiles

// Profile -
type Profile struct {
	Name               string `json:"Name"`
	AWSAccessKeyId     string `json:"AccessKeyId"`
	AWSSecretAccessKey string `json:"SecretAccessKeyId"`
	AWSSessionToken    string `json:"SessionToken"`
	Region             string `json:"Region"`
	Output             string `json:"Output"`
}
