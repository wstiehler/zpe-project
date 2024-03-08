package createuser

type DTOEntity struct {
	UserView     *UserDTO
	UserListView *UserListDTO
}

type UserDTO struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Role  RoleDTO `json:"role"`
	Email string  `json:"email"`
}

type RoleDTO struct {
	Role string `json:"role"`
}

type PermissionDTO struct {
	Name string `json:"name"`
}

type UserListDTO struct {
	Users []UserDTO `json:"users"`
}
