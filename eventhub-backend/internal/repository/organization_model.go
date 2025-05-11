package repository

type OrganizationModel struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"unique" json:"name"`
	FounderID  uint   `json:"founder_id"`
	InviteCode string `gorm:"unique" json:"invite_code"`
}

func (OrganizationModel) TableName() string {
	return "organizations"
}

type OrganizationMemberModel struct {
	UserID         uint `json:"user_id"`
	OrganizationID uint `json:"organization_id"`
}

func (OrganizationMemberModel) TableName() string {
	return "organization_members"
}
