package service

type ServiceProvider struct {
	GuestService GuestService
}

func NewServiceProvider(guestService GuestService) (sp *ServiceProvider) {
	return &ServiceProvider{GuestService: guestService}
}
