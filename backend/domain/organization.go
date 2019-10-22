package domain

import "github.com/sky0621/fiktivt-handelssystem/domain/repository"

func NewOrganization(organizationRepository repository.Organization, userRepository repository.User) Organization {
	return &organization{
		organizationRepository: organizationRepository,
		userRepository:         userRepository,
	}
}

type Organization interface {
}

type organization struct {
	organizationRepository repository.Organization
	userRepository         repository.User
}

type OrganizationModel struct {
	ID                  string
	Name                string
	UpperOrganizationID *string
	Users               []*UserModel
}

type UserModel struct {
	ID            string
	Name          string
	Organizations []*OrganizationModel
}
