package services

import (
	"github.com/adityasuryadi/ewallet/internal/interfaces"
)

type HealtcheckServices struct {
	HealtCheckRepository interfaces.IHealtcheckRepository
}

func (s *HealtcheckServices) HealthcheckServices() (string, error) {
	return "service healty", nil
}
