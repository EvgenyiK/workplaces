package models

type User struct {
	Id        int64  `json:"id"`
	PCName    string `json:"pc_name"`
	MacAdress string `json:"mac_adress"`
	UserName  string `json:"user_name"`
}
