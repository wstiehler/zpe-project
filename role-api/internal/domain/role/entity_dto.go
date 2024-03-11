package role

type RoleDTO struct {
	Role string `json:"role"`
	Id   uint   `json:"id"`
}

type PermissionDTO struct {
	Name string `json:"name"`
}
