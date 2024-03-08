package roleapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wstiehler/zpecreateuser-api/internal/environment"
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

	// Decodifique o corpo da resposta JSON em uma struct RoleDTO
	var roleDTO RoleDTO

	if err := json.NewDecoder(resp.Body).Decode(&roleDTO); err != nil {
		return nil, err
	}

	return &roleDTO, nil
}
