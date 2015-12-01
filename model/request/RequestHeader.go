package request

type RequestHeader struct {
	AccountID   string `json:"accountID"`
	DigitalSign string `json:"digitalSign"`
	ReqTime     string `json:"reqTime"`
	ServiceName string `json:"serviceName"`
	Version     string `json:"version"`
}
