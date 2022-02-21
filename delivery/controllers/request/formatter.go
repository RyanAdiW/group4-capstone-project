package request

type RequestFormat struct {
	Id_user      int    `json:"id_user" form:"id_user"`
	Id_asset     int    `json:"id_asset" form:"id_asset"`
	Id_status    int    `json:"id_status" form:"id_status"`
	Request_date string `json:"request_date" form:"request_date"`
	Return_date  string `json:"return_date" form:"return_date"`
	Descrition   string `json:"description" form:"description"`
}
