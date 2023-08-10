package models

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type Merchant struct {
	Model                   `gorm:"embedded"`
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

// Create a custom MerchantModel type which wraps the gorm.DB connection pool.
type MerchantModel struct {
	DB *gorm.DB
}

func NewMerchantModel(db *gorm.DB) *MerchantModel {
	return &MerchantModel{
		DB: db,
	}
}

func (dao *MerchantModel) Create(ctx context.Context, merchant Merchant) (*uint, error) {

	tx := dao.DB.Where("email = ?", merchant.BusinessEmail).FirstOrCreate(&merchant)

	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected != 1 {
		return nil, errors.New("exists")
	}

	return &merchant.ID, nil
}

func (dao *MerchantModel) Update(ctx context.Context, merchant Merchant) error {
	tx := dao.DB.Model(&Merchant{}).Where("id = ?", merchant.ID).Updates(merchant)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (dao *MerchantModel) Get(ctx context.Context, model Merchant) (*Merchant, error) {
	var merchant *Merchant
	tx := dao.DB.Model(model).First(&merchant)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return merchant, nil
}

func (dao *MerchantModel) GetAll(ctx context.Context) (*[]Merchant, error) {
	var merchants []Merchant
	tx := dao.DB.Find(&merchants)

	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected != 1 {
		return nil, errors.New("not found")
	}

	return &merchants, nil
}
