package roleapi

import (
	"fmt"
	"net/http"

	"github.com/wstiehler/zpeupdateuser-api/internal/environment"
)

func GetRoleByID(roleID uint) (*RoleDTO, error) {
	env := environment.GetInstance()
	url := fmt.Sprintf("%s/v1/role/%v", env.APPLICATION_URL_ROLE_API, roleID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get role details: %s", resp.Status)
	}

	roleDTO, err := decodeRoleDTO(resp.Body)
	if err != nil {
		return nil, err
	}

	return roleDTO, nil
}

func GetPermissionByRoleID(roleID uint) (*[]PermissionDTO, error) {
	env := environment.GetInstance()
	url := fmt.Sprintf("%s/v1/permission/%v", env.APPLICATION_URL_ROLE_API, roleID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get role details: %s", resp.Status)
	}

	permissionDTO, err := decodePermissionDTO(resp.Body)
	if err != nil {
		return nil, err
	}

	return permissionDTO, nil
}
