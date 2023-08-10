package models

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Branding struct {
	Model          `gorm:"embedded"`
	BrandAlias     string   `json:"brand_alias"`
	Tagline        string   `json:"tagline" gorm:"default:''"`
	Description    string   `json:"description" gorm:"default:''"`
	Logo           string   `json:"logo" gorm:"default:''"`
	Industry       string   `json:"industry" gorm:"default:''"`
	WebsiteURL     string   `json:"website_url" gorm:"default:''"`
	ReceiptMessage string   `json:"receipt_message"`
	InstagramURL   string   `json:"instagram_url" gorm:"default:''"`
	TwitterURL     string   `json:"twitter_url" gorm:"default:''"`
	FacebookURL    string   `json:"facebook_url" gorm:"default:''"`
	TiktokURL      string   `json:"tiktok_url" gorm:"default:''"`
	BrandColour    string   `json:"brand_colour" gorm:"default:''"`
	MerchantID     string   `json:"merchant_id" gorm:"->;<-:create"` // allow read and create
	Merchant       Merchant `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Create a custom BrandingModel type which wraps the gorm.DB connection pool.
type BrandingModel struct {
	DB *gorm.DB
}

func NewBrandingModel(db *gorm.DB) *BrandingModel {
	return &BrandingModel{
		DB: db,
	}
}

func (dao *BrandingModel) Create(ctx context.Context, branding Branding) (*uint, error) {

	tx := dao.DB.Where("merchant_id = ?", branding.MerchantID).FirstOrCreate(&branding)

	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected != 1 {
		return nil, errors.New("exists")
	}

	return &branding.ID, nil
}

func (dao *BrandingModel) Update(ctx context.Context, branding Branding) error {
	tx := dao.DB.Model(&Branding{}).Where("id = ?", branding.ID).Updates(branding)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (dao *BrandingModel) Get(ctx context.Context, model Branding) (*Branding, error) {
	var branding *Branding
	tx := dao.DB.Model(model).First(&branding)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return branding, nil
}

func (dao *BrandingModel) GetAll(ctx context.Context) (*[]Branding, error) {
	var brandings []Branding
	tx := dao.DB.Find(&brandings)

	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected != 1 {
		return nil, errors.New("not found")
	}

	return &brandings, nil
}
