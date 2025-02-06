package entities

type PaperFormat string

//const (
//	A4     PaperFormat = "A4"
//	Letter PaperFormat = "Letter"
//	Legal  PaperFormat = "Legal"
//)

// ResumeTemplateEntity for resume templates all entities
type ResumeTemplateEntity struct {
	BaseEntity
	ID      uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	UUID    string `gorm:"type:varchar(36);unique;not null" json:"uuid"`
	Version string `gorm:"type:varchar(10);not null" json:"version"`

	Title       string `gorm:"type:varchar(255);not null;index" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Category    string `gorm:"type:varchar(100);index" json:"category"`

	HTMLContent  string `gorm:"type:text;not null" json:"html_content"`
	CSSContent   string `gorm:"type:text;not null" json:"css_content"`
	CSSVariables string `gorm:"type:jsonb" json:"css_variables"`
	Sections     string `gorm:"type:jsonb;not null" json:"sections"`

	ThumbnailImg string `gorm:"type:varchar(512);not null" json:"thumbnail_img"`
	PreviewURL   string `gorm:"type:varchar(512)" json:"preview_url"`

	Downloads int  `gorm:"default:0" json:"downloads"`
	IsActive  bool `gorm:"default:true;index" json:"is_active"`

	DefaultFont        string      `gorm:"type:varchar(50)" json:"default_font"`
	DefaultColorScheme string      `gorm:"type:varchar(50)" json:"default_color_scheme"`
	Compatibility      PaperFormat `gorm:"type:varchar(20);check:compatibility IN ('A4','Letter','Legal');not null;default:'A4'" json:"compatibility"`
}
