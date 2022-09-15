package gooremote

// external types

// Use this struct to decode data pushed to webhook
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
	}
}

type verifyResponse struct {
	Ok               bool   `json:"ok"`
	ErrorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`
	Result           struct {
		Paid bool `json:"paid"`
	}
}

type payment struct {
	Amount      int    `bson:"amount" json:"amount"`
	Next        string `bson:"next" json:"next"`
	Webhook     string `bson:"webhook" json:"webhook"`
	Description string `bson:"description" json:"description"`
}
