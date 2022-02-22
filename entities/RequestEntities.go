package entities

type Request struct {
	Id               int    `json:"id" form:"id"`
	Id_user          int    `json:"id_user" form:"id_user"`
	Id_asset         int    `json:"id_asset" form:"id_asset"`
	Id_status        int    `json:"id_status" form:"id_status"`
	Request_date     string `json:"request_date" form:"request_date"`
	Return_date      string `json:"return_date" form:"return_date"`
	Description      string `json:"description" form:"description"`
	User_name        string `json:"user_name" form:"user_name"`
	Asset_name       string `json:"asset_name" form:"asset_name"`
	Category         string `json:"category" form:"category"`
	Avail_quantity   int    `json:"avail_quantity" form:"avail_quantity"`
	Initial_quantity int    `json:"initial_quantity" form:"initial_quantity"`
	Status           string `json:"status" form:"status"`
}
