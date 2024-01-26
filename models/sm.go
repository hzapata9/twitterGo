package models

type Secret struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	JWTSign  string `json:"jswsign"`
	Database string `json:"database"`
}
