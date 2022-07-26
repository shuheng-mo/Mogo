package model

type ProductSeo struct {
	ID             int64  `gorm:"primary_key;not_null;auto_increment" json:"id"`
	SeoTitle       string `json:"seo_title"`
	SeoKeyWords    string `json:"seo_key_words"`
	SeoDescription string `json:"seo_description"`
	SeoCode        string `json:"seo_code"`
	SeoProductID   int64  `json:"seo_product_id"`
}
