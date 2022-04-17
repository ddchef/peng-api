package models

type Domain struct {
	ID
	DomainName string `json:"string" gorm:"not null;comment:组名称"`
	Timestamps
	SoftDeletes
}
