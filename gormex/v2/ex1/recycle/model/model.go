package model

type Category struct {
	BaseModel
	Name             string      `gorm:"type:varchar(16);not null" json:"name"`
	ParentCategoryID int32       `json:"parent_id"`
	ParentCategory   *Category   `json:"-"`
	SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID;references:ID" json:"sub_category"`
	Level            int32       `gorm:"type:int;not null;default:1" json:"level"`
	IsTab            bool        `gorm:"default:false;not null" json:"is_tab"`
}
