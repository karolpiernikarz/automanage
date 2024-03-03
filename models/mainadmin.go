package models

import (
	"gorm.io/datatypes"
)

type Restaurant struct {
	ID                uint                                `json:"id" db:"id"`
	CompanyId         int                                 `json:"company_id" db:"company_id"`
	Name              string                              `json:"name,omitempty" db:"name,omitempty"`
	Logo              string                              `json:"logo,omitempty" db:"logo,omitempty"`
	Address           string                              `json:"address,omitempty" db:"address"`
	IsActive          int                                 `json:"is_active,omitempty" db:"is_active"`
	Email             string                              `json:"email,omitempty" db:"email"`
	Phone             string                              `json:"phone,omitempty" db:"phone"`
	Gsm               string                              `json:"gsm,omitempty" db:"gsm"`
	Website           string                              `json:"website,omitempty" db:"website"`
	Token             string                              `json:"token,omitempty" db:"token"`
	Password          string                              `json:"password,omitempty" db:"password"`
	CommissionDetails string                              `json:"commission_details,omitempty" db:"commission_details"`
	PaymentDetails    string                              `json:"payment_details,omitempty" db:"payment_details"`
	Info              datatypes.JSONType[*RestaurantInfo] `json:"info,omitempty" db:"info,omitempty"`
	CreatedAt         string                              `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt         string                              `json:"updated_at,omitempty" db:"updated_at"`
}

type RestaurantInfo struct {
	CurrierType      string      `json:"currier_type" db:"currier_type,omitempty"`
	VenueId          string      `json:"venue_id" db:"venue_id,omitempty"`
	Currier1         string      `json:"currier_1" db:"currier_1,omitempty"`
	Currier2         string      `json:"currier_2" db:"currier_2,omitempty"`
	Currier3         string      `json:"currier_3" db:"currier_3,omitempty"`
	Currier4         string      `json:"currier_4" db:"currier_4,omitempty"`
	Currier5         string      `json:"currier_5" db:"currier_5,omitempty"`
	Currier6         string      `json:"currier_6" db:"currier_6,omitempty"`
	Currier7         string      `json:"currier_7" db:"currier_7,omitempty"`
	Currier8         string      `json:"currier_8" db:"currier_8,omitempty"`
	Currier9         string      `json:"currier_9" db:"currier_9,omitempty"`
	Currier10        string      `json:"currier_10" db:"currier_10,omitempty"`
	Currier11        string      `json:"currier_11" db:"currier_11,omitempty"`
	Currier12        string      `json:"currier_12" db:"currier_12,omitempty"`
	Currier13        string      `json:"currier_13" db:"currier_13,omitempty"`
	Currier14        string      `json:"currier_14" db:"currier_14,omitempty"`
	TerminalType     string      `json:"terminal_type" db:"terminal_type,omitempty"`
	TerminalId       interface{} `json:"terminal_id" db:"terminal_id,omitempty"`
	TerminalUsername string      `json:"terminal_username" db:"terminal_username,omitempty"`
	TerminalPassword string      `json:"terminal_password" db:"terminal_password,omitempty"`
}

type OrderboxInfo struct {
	TerminalType     string `json:"terminal_type" db:"terminal_type,"`
	TerminalUsername string `json:"terminal_username" db:"terminal_username,"`
	TerminalPassword string `json:"terminal_password" db:"terminal_password,"`
	ID               uint   `json:"id," db:"id,"`
}

type Companies struct {
	Id              int                                  `json:"id" db:"id"`
	Name            string                               `json:"name" db:"name"`
	ChainName       string                               `json:"chain_name" db:"chain_name"`
	Email           string                               `json:"email" db:"email"`
	Password        string                               `json:"password" db:"password"`
	Description     string                               `json:"description" db:"description"`
	Phone           string                               `json:"phone" db:"phone"`
	TaxNumber       string                               `json:"tax_number" db:"tax_number"`
	Address         string                               `json:"address" db:"address"`
	Bank            datatypes.JSONType[CompaniesBank]    `json:"bank" db:"bank"`
	Contact         datatypes.JSONType[CompaniesContact] `json:"contact" db:"contact"`
	Domain          string                               `json:"domain" db:"domain"`
	Social          datatypes.JSONType[CompaniesSocial]  `json:"social" db:"social"`
	BackgroundImage string                               `json:"background_image" db:"background_image"`
	IsActive        int                                  `json:"is_active" db:"is_active"`
	CreatedAt       string                               `json:"created_at" db:"created_at"`
	UpdatedAt       string                               `json:"updated_at" db:"updated_at"`
}

type CompaniesBank struct {
	Navn        datatypes.JSON `json:"navn" db:"navn"`
	RegNumber   datatypes.JSON `json:"reg_number" db:"reg_number"`
	KontoNumber datatypes.JSON `json:"konto_number" db:"konto_number"`
	Swift       datatypes.JSON `json:"swift" db:"swift"`
	Iban        datatypes.JSON `json:"iban" db:"iban"`
}

type CompaniesContact struct {
	Title datatypes.JSON `json:"title" db:"title"`
	Name  datatypes.JSON `json:"name" db:"name"`
	Email datatypes.JSON `json:"email" db:"email"`
	Phone datatypes.JSON `json:"phone" db:"phone"`
	Notes datatypes.JSON `json:"notes" db:"notes"`
}

type CompaniesSocial struct {
	Facebook  datatypes.JSON `json:"facebook" db:"facebook"`
	Instagram datatypes.JSON `json:"instagram" db:"instagram"`
}

type MainAdminRestaurant struct {
	AdminId      int    `json:"admin_id" db:"admin_id"`
	RestaurantId int    `json:"restaurant_id" db:"restaurant_id"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	UpdatedAt    string `json:"updated_at" db:"updated_at"`
}

type CompanyHtmlResponse struct {
	Logo          string `json:"logo" html:"logo" binding:"required"`
	CompanyName   string `json:"company_name" html:"company_name" binding:"required"`
	CompanyDesc   string `json:"company_desc" html:"company_desc" binding:"required"`
	BackgroundUrl string `json:"background_url" html:"background_url" binding:"required"`
	Restaurants   []CompanyHtmlResponseRestaurants
}

type CompanyHtmlResponseRestaurants struct {
	Name    string `json:"name" html:"name" binding:"required"`
	Address string `json:"address" html:"address" binding:"required"`
	Url     string `json:"url" html:"url" binding:"required"`
}
