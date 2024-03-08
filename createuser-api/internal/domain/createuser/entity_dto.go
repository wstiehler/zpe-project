package createuser

type DTOEntity struct {
	UserView     *UserDTO
	UserListView *UserListDTO
}

type UserDTO struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Email string `json:"email"`
}

type UserListDTO struct {
	User []UserDTO `json:"users"`
}
