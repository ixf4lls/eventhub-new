package repository

type OrganizationRepository interface {
	Create(name string, founderID uint) error
}
