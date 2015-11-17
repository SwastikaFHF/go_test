package response

type ResponseHeader struct {
	RspType string `json:"rspType"`
	RspCode string `json:"rspCode"`
	RspMsg  string `json:"rspMsg"`
}
