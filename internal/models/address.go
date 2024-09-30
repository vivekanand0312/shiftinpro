package models

type Address struct {
	ID        int     `json:"id"`
	Pincode   float64 `json:"pincode"`
	City      string  `json:"city"`
	District  string  `json:"district"`
	State     string  `json:"state"`
	CountryID int     `json:"country_id"`
	Country   Country `json:"country"`
}

type Country struct {
	ID       int    `json:"id"`
	Sortname string `json:"sortname"`
	Name     string `json:"name"`
}
