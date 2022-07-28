package models

type XploitDB struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	URL    string `json:"url"`
	CVE    string `json:"cve" gorm:"unique"`
	Author string `json:"author"`
}
