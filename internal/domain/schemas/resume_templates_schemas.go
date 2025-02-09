package schemas

type CreateResumeTemplateSchemas struct {
	Title              string   `json:"title" binding:"required,max=255"`
	Description        string   `json:"description,omitempty"`
	Category           string   `json:"category,omitempty,max=100"`
	Language           string   `json:"language" binding:"required,max=20"`
	Tags               []string `json:"tags,omitempty"`
	HTMLContent        string   `json:"html_content" binding:"required"`
	CSSContent         string   `json:"css_content" binding:"required"`
	Sections           []string `json:"sections" binding:"required"`
	ThumbnailImg       string   `json:"thumbnail_img" binding:"required,url"`
	PreviewURL         string   `json:"preview_url,omitempty,url"`
	DefaultFont        string   `json:"default_font,omitempty,max=50"`
	DefaultColorScheme string   `json:"default_color_scheme,omitempty,max=50"`
	Compatibility      string   `json:"compatibility" binding:"required,oneof=A4 Letter Legal"`
	Downloads          int      `json:"downloads,omitempty"`
	IsActive           bool     `json:"is_active,omitempty"`
	IsFeatured         bool     `json:"is_featured,omitempty"`
	Rating             float32  `json:"rating,omitempty"`
	License            string   `json:"license,omitempty,max=100"`
	IsPremium          bool     `json:"is_premium,omitempty"`
}
