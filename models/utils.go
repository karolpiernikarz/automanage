package models

// SlackWebhookMessage is the struct for Slack webhook message
type SlackWebhookMessage struct {
	Text   string                      `json:"text"`
	Blocks []SlackWebhookMessageBlocks `json:"blocks,omitempty"`
}

type SlackWebhookMessageBlocks struct {
	Type    string `json:"type"`
	BlockId string `json:"block_id,omitempty"`
	Text    struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"text,omitempty"`
	//Accessory SlackWebhookMessageBlocksAccessory `json:"accessory,omitempty"`
	Fields []SlackWebhookMessageBlocksField `json:"fields,omitempty"`
}

type SlackWebhookMessageBlocksField struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

type SlackWebhookMessageBlocksAccessory struct {
	Type     string `json:"type,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
	AltText  string `json:"alt_text,omitempty"`
}

type OrderBoxPrinter struct {
	FirstValue      string `json:"first_value,omitempty"`
	OrderType       string `json:"order_type,omitempty"`
	DeliveryType    string `json:"delivery_type,omitempty"`
	DeliveryTime    string `json:"delivery_time,omitempty"`
	CustomerName    string `json:"customer_name,omitempty"`
	PaymentType     string `json:"payment_type,omitempty"`
	OrderNumber     string `json:"order_number,omitempty"`
	Order           string `json:"order,omitempty"`
	SubTotal        string `json:"sub_total,omitempty"`
	BagFee          string `json:"bag_fee,omitempty"`
	DeliveryFee     string `json:"delivery_fee,omitempty"`
	ServiceFee      string `json:"service_fee,omitempty"`
	Total           string `json:"total,omitempty"`
	Notes           string `json:"notes,omitempty"`
	KundeInfo       string `json:"kunde_info,omitempty"`
	CustomerDetails string `json:"customer_details,omitempty"`
	CustomerPhone   string `json:"customer_phone,omitempty"`
	CustomerCard    string `json:"customer_card,omitempty"`
}
