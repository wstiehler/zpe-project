package roleapi

import (
	"encoding/json"
	"io"
)

func decodeRoleDTO(body io.Reader) (*RoleDTO, error) {
	var roleDTO RoleDTO
	if err := json.NewDecoder(body).Decode(&roleDTO); err != nil {
		return nil, err
	}
	return &roleDTO, nil
}
