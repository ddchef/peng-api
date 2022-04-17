package request

type Domain struct {
	DomainName string `form:"domainName" json:"name" binding:"max=10,min=3,required"`
}

func (domain Domain) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"domainName.max":      "组名称长度必须大于10",
		"domainName.min":      "组名称长度必须大于3",
		"domainName.required": "组名称为必填项",
	}
}
