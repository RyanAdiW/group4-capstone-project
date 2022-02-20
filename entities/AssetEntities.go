package entities

type Asset struct {
	Id               int    `json:"id" form:"id"`
	Id_category      int    `json:"id_category" form:"id_category"`
	Is_maintenance   bool   `json:"is_maintenance" form:"is_maintenance"`
	Name             string `json:"name" form:"name"`
	Description      string `json:"description" form:"description"`
	Initial_quantity int    `json:"initial_quantity" form:"initial_quantity"`
	Avail_quantity   int    `json:"avail_quantity" form:"avail_quantity"`
	Photo            string `json:"photo" form:"photo"`
	Category         string `json:"category" form:"category"`
}

type SummaryAsset struct {
	Total_asset int `json:"total_asset" form:"total_asset"`
	Use         int `json:"use" form:"use"`
	Maintenance int `json:"maintenance" form:"maintenance"`
	Available   int `json:"available" form:"available"`
}

type HistoryUsage struct {
	Id           int    `json:"id" form:"id"`
	Id_category  int    `json:"id_category" form:"id_category"`
	Name         string `json:"name" form:"name"`
	Description  string `json:"description" form:"description"`
	Photo        string `json:"photo" form:"photo"`
	Category     string `json:"category" form:"category"`
	List_history []History
}

type History struct {
	Id           int    `json:"id" form:"id"`
	Id_user      int    `json:"id_user" form:"id_user"`
	Id_asset     int    `json:"id_asset" form:"id_asset"`
	Name         string `json:"name" form:"name"`
	Request_date string `json:"request_date" form:"request_date"`
	Status       string `json:"status" form:"status"`
}
