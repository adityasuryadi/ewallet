package services

type HealtcheckServices struct {
}

func (s *HealtcheckServices) HealthcheckServices() (string, error) {
	return "service healty", nil
}
