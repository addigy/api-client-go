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

func NewAddigyClient(clientID string, clientSecret string) *AddigyClient {
	return &AddigyClient{
		ClientID: clientID,
		ClientSecret: clientSecret,
		BaseURL: DevURL,
	}
}

func (addigy AddigyClient) SwitchToDev() {
	addigy.BaseURL = DevURL
}

func (addigy AddigyClient) SwitchToStage() {
	addigy.BaseURL = StageURL
}

func (addigy AddigyClient) SwitchToProd() {
	addigy.BaseURL = ProdURL
}