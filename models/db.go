package models

type XploitDB struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Title    string `json:"title"`
	URL      string `json:"url" gorm:"unique"`
	CVE      string `json:"cve"`
	Author   string `json:"author"`
	Platform string `json:"platform"`
	Type     string `json:"type"`
	EDBID    string `json:"edb_id"`
}
