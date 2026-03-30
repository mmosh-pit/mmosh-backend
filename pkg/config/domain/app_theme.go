package domain

type AppTheme struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	CodeName        string `json:"code_name"`
	BackgroundColor string `json:"background_color"`
	PrimaryColor    string `json:"primary_color"`
	SecondaryColor  string `json:"secondary_color"`
	Logo            string `json:"logo"`
}
