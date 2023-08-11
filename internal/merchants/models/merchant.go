package models

import (
	"kiplingkelvin/golang-skeleton/internal/common/models"
	"kiplingkelvin/golang-skeleton/internal/common/utils"
)

type Merchant struct {
	models.Model            `gorm:"embedded"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	CompanyName             string `json:"company_name"`
	BusinessEmail           string `json:"business_email"`
	Industry                string `json:"industry"`
	MobileNumber            string `json:"mobile_number"`
	OfficeNumber            string `json:"office_number"`
	Country                 string `json:"country"`
	City                    string `json:"city"`
	Address                 string `json:"address"`
	PricingModel            string `json:"pricing_model"`
	PersonalEmail           string `json:"personal_email"`
	Nationality             string `json:"nationality"`
	NotifyPersonalEmail     bool   `json:"notify_personal_email"`
	NotifyBusinessEmail     bool   `json:"notify_business_email"`
	NotifyOfficeNumber      bool   `json:"notify_office_number"`
	NotifyMobileNumber      bool   `json:"notify_mobile_number"`
	IsEmailVerified         bool   `json:"is_email_verified" gorm:"default:false"`
	IsAccountVerified       bool   `json:"is_account_verified" gorm:"default:false"`
	IsAccountActivated      bool   `json:"is_account_activated" gorm:"default:false"`
	IsAccountSetupComplete  bool   `json:"is_account_setup_complete" gorm:"default:false"`
	IsAcceptingCardPayments bool   `json:"is_accepting_card_payments" gorm:"default:false"`
	Password                []byte `json:"-"`
	MigrationMerchantID     string `json:"migration_merchant_id" gorm:"default:null"`
	IsShopifyActive         bool   `json:"is_shopify_active" gorm:"default:false"`
	ShopifyCode             string `json:"shopify_code" gorm:"default:null"`
	ShopifyDomain           string `json:"shopify_domain" gorm:"default:null"`
	ShopifyAccessToken      string `json:"shopify_access_token" gorm:"default:null"`
}

type MerchantRegisterRequest struct {
	Email              string `json:"email" validate:"required,email"`
	FirstName          string `json:"first_name" validate:"required"`
	LastName           string `json:"last_name" validate:"required"`
	Company            string `json:"company" validate:"required"`
	PhoneNumber        string `json:"phone_number" validate:"required"`
	Password           string `json:"password" validate:"required"`
	IsShopifyActive    bool   `json:"is_shopify_active"`
	ShopifyCode        string `json:"shopify_code"`
	ShopifyDomain      string `json:"shopify_domain"`
	ShopifyAccessToken string `json:"shopify_access_token"`
}

type MerchantRegisterResponse struct {
	utils.Response
	ConfirmEmailURL string `json:"confirm_email_url"`
	IsShopifyActive bool   `json:"is_shopify_active"`
}
