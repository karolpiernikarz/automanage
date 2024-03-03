package models

import "time"

type WoltWebHookBody struct {
	Token string `json:"token"`
}

type WoltWebHook struct {
	DispatchedAt time.Time `json:"dispatched_at"`
	Type         string    `json:"type"`
	Details      struct {
		Id                       string `json:"id"`
		VenueId                  string `json:"venue_id"`
		WoltOrderReferenceId     string `json:"wolt_order_reference_id"`
		TrackingReference        string `json:"tracking_reference"`
		MerchantOrderReferenceId string `json:"merchant_order_reference_id"`
		Price                    struct {
			Amount   int    `json:"amount"`
			Currency string `json:"currency"`
		} `json:"price"`
		Pickup struct {
			Eta *time.Time `json:"eta,omitempty"`
		} `json:"pickup"`
		Dropoff *struct {
			Eta         *time.Time `json:"eta,omitempty"`
			CompletedAt *time.Time `json:"completed_at,omitempty"`
		} `json:"dropoff"`
		Courier *interface{} `json:"courier"`
	} `json:"details"`
}
