package app

import (
	"peng-api/app/common/request"
	"peng-api/app/common/response"
	"peng-api/app/do"
	"peng-api/app/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 创建角色
// Register 创建角色接口
// @Summary 创建角色接口
// @Tags 角色管理
// @Security BearerAuth
// @Param form body request.Role true "角色信息"
// @Success 200 {object} response.Response{data=models.Role}
// @Router /role [post]
func RoleCreate(c *gin.Context) {
	var form request.Role
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(c, err))
	}
	if err, role := services.RoleService.Create(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, role)
	}
}

// 更新角色
// Register 更新角色接口
// @Summary 更新角色接口
// @Tags 角色管理
// @Security BearerAuth
// @Param form body request.Role true "角色信息"
// @Success 200 {object} response.Response{data=models.Role}
// @Router /role/{id} [put]
func RoleUpdate(c *gin.Context) {
	id := c.Param("id")
	var form request.UpdateRole
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(c, err))
	}
	if err, role := services.RoleService.Update(id, form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, role)
	}
}

// 删除角色
// Register 删除角色接口
// @Summary 删除角色接口
// @Tags 角色管理
// @Security BearerAuth
// @Param id path string true "角色id"
// @Success 200 {object} response.Response
// @Router /role/{id} [delete]
func RoleDelete(c *gin.Context) {
	id := c.Param("id")
	if err := services.RoleService.Delete(id); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, "")
	}
}

// 获取角色信息
// Register 获取角色信息接口
// @Summary 获取角色信息接口
// @Tags 角色管理
// @Security BearerAuth
// @Param id path string true "角色id"
// @Success 200 {object} response.Response{data=models.Role}
// @Router /role/{id} [get]
func RoleInfo(c *gin.Context) {
	id := c.Param("id")
	if err, role := services.RoleService.Info(id); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, role)
	}
}

// 获取角色列表
// Register 获取角色列表接口
// @Summary 获取角色列表接口
// @Tags 角色管理
// @Security BearerAuth
// @Param id path string true "角色id"
// @Success 200 {object} response.Response{data=models.Role}
// @Router /role/list [get]
func RoleList(c *gin.Context) {
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err, roles, count := services.RoleService.List(offset, limit); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, do.List{
			List:  roles,
			Total: count,
		})
	}
}
