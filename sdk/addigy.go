package sdk

const (
	DevURL = "https://dev.addigy.com"
	StageURL = "https://stage.addigy.com"
	ProdURL = "https://prod.addigy.com"
)

type AddigyClient struct {
	ClientID     string
	ClientSecret string
	BaseURL		 string
}

//todo pass in realm here and remove the hardcoded urls
func NewAddigyClient(clientID string, clientSecret string) *AddigyClient {
	return &AddigyClient{
		ClientID: clientID,
		ClientSecret: clientSecret,
		BaseURL: DevURL, //todo generate the url here with the realm param
	}
}

//todo: what are these for? if they aren't being used then remove them
func (addigy AddigyClient) SwitchToDev() {
	addigy.BaseURL = DevURL
}

func (addigy AddigyClient) SwitchToStage() {
	addigy.BaseURL = StageURL
}

func (addigy AddigyClient) SwitchToProd() {
	addigy.BaseURL = ProdURL
}
//todo: create a method to make a request here that takes (request *http.Request, responseObj interface{})

//todo: create a method to generate the url that takes in (baseUrl string, params map[string]string)