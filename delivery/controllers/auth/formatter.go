package auth

type LoginEmailRequestFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginResponseFormat struct {
	Token   string `json:"token" form:"token"`
	Id_user int    `json:"id_user" form:"id_user"`
	Id_role int    `json:"id_role" form:"id_role"`
	Name    string `json:"name" form:"name"`
}
