package gooremote

// external types

// Use this struct to decode data sent to webhook
type WebhookPayload struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
}

// internal types

type instance struct {
	Key string
}

type newPaymentResponse struct {
	Ok               bool   `json:"ok"`
	ErrorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`
	Result           struct {
		ID string `json:"id"`
	} `json:"result"`
}

type verifyResponse struct {
	Ok               bool   `json:"ok"`
	ErrorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`
	Result           struct {
		Paid bool `json:"paid"`
	} `json:"result"`
}

type payment struct {
	Amount      int    `bson:"amount" json:"amount"`
	Next        string `bson:"next" json:"next"`
	Webhook     string `bson:"webhook" json:"webhook"`
	Description string `bson:"description" json:"description"`
}

type getRawGatewayURLResponse struct {
	Ok               bool   `json:"ok"`
	ErrorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`
	Result           struct {
		URL string `json:"url"`
	} `json:"result"`
}
