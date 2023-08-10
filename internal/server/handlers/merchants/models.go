package merchants

import "kiplingkelvin/golang-skeleton/internal/server/handlers/common/utils"

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
