package models

import (
	commons_models "kiplingkelvin/golang-skeleton/internal/common/models"
	"kiplingkelvin/golang-skeleton/internal/common/utils"
	merchant_models "kiplingkelvin/golang-skeleton/internal/merchants/models"
)

type BankAccount struct {
	commons_models.Model `gorm:"embedded"`
	CountryCode          string                   `json:"country_code"`
	BankName             string                   `json:"bank_name"`
	BankBranch           string                   `json:"bank_branch"`
	BankCode             string                   `json:"bank_code"`
	AccountType          string                   `json:"account_type" gorm:"default:'Domestic'"`
	AccountName          string                   `json:"account_name"`
	AccountNumber        string                   `json:"account_number"`
	MerchantID           string                   `json:"merchant_id" gorm:"->;<-:create"` // allow read and create
	Merchant             merchant_models.Merchant `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type BankAccountCreateResponse struct {
	utils.Response
}
