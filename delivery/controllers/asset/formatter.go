package asset

type UserRequestFormat struct {
	Id_category      int    `json:"id_category" form:"id_category"`
	Is_maintenence   bool   `json:"is_maintenence" form:"is_maintenence"`
	Name             string `json:"name" form:"name"`
	Description      string `json:"description" form:"description"`
	Initial_quantity int    `json:"initial_quantity" form:"initial_quantity"`
	Photo            string `json:"photo" form:"photo"`
}
