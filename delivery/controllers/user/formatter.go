package user

type UserRequestFormat struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Divisi   string `json:"divisi" form:"divisi"`
	Id_role  int    `json:"id_role" form:"id_role"`
}
