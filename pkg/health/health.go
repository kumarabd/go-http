package health

import (
	"github.com/kumarabd/gokit/models"
	"github.com/kumarabd/gokit/utils"
)

func GetHealth() *models.Health {
	version, _ := utils.Git()
	return &models.Health{
		Version: version,
		Status:  "Healthy",
		Error:   "None",
	}
}
