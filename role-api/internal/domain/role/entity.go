package role

type ModelsEntity struct {
	RoleEntity *RoleEntity
	Permission *PermissionEntity
}

type RoleEntity struct {
	ID          uint               `gorm:"primary_key"`
	Role        string             `json:"role" gorm:"unique"`
	Permissions []PermissionEntity `json:"permissions" gorm:"foreignKey:RoleId"`
}

type PermissionEntity struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	RoleId uint   `json:"role_id"`
	Name   string `json:"name"`
}

func (r *RoleEntity) TableName() string {
	return "roles"
}

func (p *PermissionEntity) TableName() string {
	return "permissions"
}
