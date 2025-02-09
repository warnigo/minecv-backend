package entities

import (
	"gorm.io/datatypes"
)

type PaperFormat string

// ResumeTemplateEntity for resume templates all entities
type ResumeTemplateEntity struct {
	BaseEntity
	ID          string         `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"type:varchar(255);not null;index" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	Category    string         `gorm:"type:varchar(100);index" json:"category"`
	Language    string         `gorm:"type:varchar(20);not null" json:"language"`
	Tags        datatypes.JSON `gorm:"type:jsonb" json:"tags"`

	HTMLContent string         `gorm:"type:text;not null" json:"html_content"`
	CSSContent  string         `gorm:"type:text;not null" json:"css_content"`
	Sections    datatypes.JSON `gorm:"type:jsonb;not null" json:"sections"`

	ThumbnailImg string `gorm:"type:varchar(255);not null" json:"thumbnail_img"`
	PreviewURL   string `gorm:"type:varchar(255)" json:"preview_url"`

	Downloads  int  `gorm:"default:0" json:"downloads"`
	IsActive   bool `gorm:"default:true;index" json:"is_active"`
	IsFeatured bool `gorm:"default:false" json:"is_featured"`

	DefaultFont        string      `gorm:"type:varchar(50)" json:"default_font"`
	DefaultColorScheme string      `gorm:"type:varchar(50)" json:"default_color_scheme"`
	Compatibility      PaperFormat `gorm:"type:varchar(20);check:compatibility IN ('A4','Letter','Legal');not null;default:'A4'" json:"compatibility"`
	Rating             float32     `gorm:"type:decimal(3,2);default:0.0" json:"rating"`
	License            string      `gorm:"type:varchar(100)" json:"license"`
	IsPremium          bool        `gorm:"default:false" json:"is_premium"`
}
