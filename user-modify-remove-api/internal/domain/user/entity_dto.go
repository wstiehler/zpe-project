package user

type DTOEntity struct {
	UserView      *UserDTO
	UserListView  *UserListDTO
	UserLoginView *UserLoginDTO
}

type UserDTO struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Role  RoleDTO `json:"role"`
	Email string  `json:"email"`
}

type UserLoginDTO struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Token string `json:"token"`
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
