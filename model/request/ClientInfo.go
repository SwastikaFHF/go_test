package request

type ClientInfo struct {
	ClientIp      string `json:"clientIp"`
	DeviceId      string `json:"deviceId"`
	Extend        string `json:"extend"`
	MacAddress    string `json:"macAddress"`
	NetworkType   string `json:"networkType"`
	PlatId        string `json:"platId"`
	PushInfo      string `json:"pushInfo"`
	RefId         string `json:"refId"`
	VersionNumber string `json:"versionNumber"`
	VersionType   string `json:"versionType"`
}
