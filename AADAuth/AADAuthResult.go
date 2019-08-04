package AADAuth

const (
	AuthType_Office365 = iota
	AuthType_IFD       = iota
	AuthType_AD        = iota
)

const (
	AuthMode_Password = iota
	AuthMode_Secret   = iota
)

type AADAuthResult struct {
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    string `json:"expires_in"`
	ExtExpiresIn string `json:"ext_expires_in"`
	ExpiresOn    string `json:"expires_on"`
	NotBefore    string `json:"not_before"`
	Resource     string `json:"resource"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`

	AuthType int32
	AuthMode int32

	AdUserName string
	AdPassword string
	AdDomain   string
}
