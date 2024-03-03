package models

type SupportLastOrdersResponse struct {
	OrderNumber        string `json:"order_number"`
	OrderStatus        int    `json:"order_status"`
	OrderType          int    `json:"order_type"`
	OrderCreated       string `json:"order_created"`
	OrderDate          string `json:"order_date"`
	OrderAddress       string `json:"order_address"`
	OrderTotal         string `json:"order_total"`
	OrderLink          string `json:"order_link"`
	OrderIsPreOrder    int    `json:"order_preorder"`
	OrderboxStatus     int    `json:"orderbox_status"`
	OrderboxResponded  string `json:"orderbox_responded,omitempty"`
	OrderboxPrinted    int    `json:"orderbox_printed,omitempty"`
	RestaurantId       uint   `json:"restaurant_id"`
	RestaurantName     string `json:"restaurant_name"`
	RestaurantPhone    string `json:"restaurant_phone"`
	RestaurantAddress  string `json:"restaurant_address"`
	CustomerName       string `json:"customer_name"`
	CustomerPhone      string `json:"customer_phone"`
	CustomerAddress    string `json:"customer_address"`
	CustomerEmail      string `json:"customer_email,omitempty"`
	CurrierStatus      string `json:"currier_status,omitempty"`
	CurrierReferenceId string `json:"currier_reference_id,omitempty"`
	CurrierTrackingUrl string `json:"currier_tracking_url,omitempty"`
	CurrierPickupEta   string `json:"currier_pickup_eta,omitempty"`
	CurrierDropoffEta  string `json:"currier_dropoff_eta,omitempty"`
	CurrierStatusText  string `json:"currier_status_text,omitempty"`
}
