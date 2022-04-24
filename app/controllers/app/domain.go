package app

import (
	"peng-api/app/common/request"
	"peng-api/app/common/response"
	"peng-api/app/do"
	"peng-api/app/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 创建组
// Register 创建组接口
// @Summary 创建组接口
// @Tags 组管理
// @Security BearerAuth
// @Param form body request.Domain true "角色信息"
// @Success 200 {object} response.Response{data=models.Domain}
// @Router /domain [post]
func DomainCreate(c *gin.Context) {
	var form request.Domain
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(c, err))
	}
	if err, domain := services.DomainService.Create(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, domain)
	}
}

// 更新组
// Register 更新组接口
// @Summary 更新组接口
// @Tags 组管理
// @Security BearerAuth
// @Param form body request.Domain true "组信息"
// @Success 200 {object} response.Response{data=models.Domain}
// @Router /domain/{id} [put]
func DomainUpdate(c *gin.Context) {
	id := c.Param("id")
	var form request.Domain
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(c, err))
	}
	if err, domain := services.DomainService.Update(id, form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, domain)
	}
}

// 删除组
// Register 删除组接口
// @Summary 删除组接口
// @Tags 组管理
// @Security BearerAuth
// @Param id path string true "组id"
// @Success 200 {object} response.Response
// @Router /domain/{id} [delete]
func DomainDelete(c *gin.Context) {
	id := c.Param("id")
	if err := services.DomainService.Delete(id); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, "")
	}
}

// 获取组信息
// Register 获取组信息接口
// @Summary 获取组信息接口
// @Tags 组管理
// @Security BearerAuth
// @Param id path string true "组id"
// @Success 200 {object} response.Response{data=models.Domain}
// @Router /domain/{id} [get]
func DomainInfo(c *gin.Context) {
	id := c.Param("id")
	if err, domain := services.DomainService.Info(id); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, domain)
	}
}

// 获取组列表
// Register 获取组列表接口
// @Summary 获取组列表接口
// @Tags 组管理
// @Security BearerAuth
// @Param offset query string false "偏移量"
// @Param limit query string false "分页大小"
// @Success 200 {object} response.Response{data=models.Domain}
// @Router /domain/list [get]
func DomainList(c *gin.Context) {
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err, domain, count := services.DomainService.List(offset, limit); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, do.List{
			List:  domain,
			Total: count,
		})
	}
}
