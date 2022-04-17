package services

import (
	"errors"
	"peng-api/app/common/request"
	"peng-api/app/models"
	"peng-api/global"
)

type domainService struct{}

var DomainService = new(domainService)

func (domainService *domainService) Create(parmas request.Domain) (err error, domain models.Domain) {
	result := global.App.DB.Where("domainName = ?", parmas.DomainName).Select("id").First(&models.Domain{})
	if result.RowsAffected != 0 {
		err = errors.New("组已存在")
		return
	}
	domain = models.Domain{DomainName: parmas.DomainName}
	err = global.App.DB.Create(&domain).Error
	return
}

func (domainService *domainService) Update(id string, parmas request.Domain) (err error, domain models.Domain) {
	err, domain = domainService.Info(id)
	if err != nil {
		return
	}
	domain.DomainName = parmas.DomainName
	err = global.App.DB.Save(&domain).Error
	if err != nil {
		err = errors.New("数据更新失败")
	}
	return
}

func (domainService *domainService) Info(id string) (err error, domain models.Domain) {
	err = global.App.DB.First(&domain, "id = ?", id).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}

func (domainService *domainService) List(offset int, limit int) (err error, domains []models.Domain, count int64) {
	err = global.App.DB.Model(&models.Domain{}).Offset(offset).Limit(limit).Find(&domains).Count(&count).Error
	if err != nil {
		err = errors.New("查询失败")
	}
	return
}

func (domainService *domainService) Delete(id string) (err error) {
	err, domain := domainService.Info(id)
	if err != nil {
		return
	}
	err = global.App.DB.Delete(&domain).Error
	if err != nil {
		err = errors.New("删除失败")
	}
	return
}
