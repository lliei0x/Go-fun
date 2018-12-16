package admin

import (
	"fmt"
	"net/http"

	"leeif.me/Go-fun/FIFA-Backstage/infra/crypt"
	"leeif.me/Go-fun/FIFA-Backstage/infra/init"
	"leeif.me/Go-fun/FIFA-Backstage/infra/model"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var (
	StateAdmin = "admin"
)

var (
	ErrorAdminSignInParam = errors.New("admin sign in error")
	ErrorAdminSignName    = errors.New("name is not allow")
	ErrorAdminSignUpName  = errors.New("name is exists")
)

var (
	NameInfo = fmt.Sprintf("Please Check Name length")
)

type ListSignInParam struct {
	Data struct {
		Name     string `json:"name" binding:"required" example:"18717711819"`
		Password string `json:"password" binding:"required" example:"xxxxxx"`
	} `json:"admin" binding:"required"`
}

// SignIn will find  a count
// @Summary Find a count
// @Accept json
// @Tags Admin
// @Security Bearer
// @Produce  json
// @Param admin body admin.ListSignInParam true "sign in information"
// @Resource Admins
// @Router /v1/api/admin/sign_in [post]
// @Success 200 {object} model.AdminSerializer
func SignIn(c *gin.Context) {

	var param ListSignInParam

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(400, c.AbortWithError(400, ErrorAdminSignInParam))
		return
	}

	// step 1 check name
	var ok bool
	ok = crypt.CheckSignInName(param.Data.Name)
	if ok != true {
		fmt.Println(NameInfo)
		return
	}

	// step 2 check exists or not
	var admin model.Admin
	if dbError := initiator.POSTGRES.Where("name = ?", param.Data.Name).First(&admin).Error; dbError != nil {
		c.JSON(400, c.AbortWithError(400, dbError))
		return
	}

	c.JSON(http.StatusOK, admin.Serializer())

}

type ListSignUpParam struct {
	Data struct {
		Name     string `json:"name" binding:"required" example:"18717711819"`
		Password string `json:"password" binding:"required" example:"xxxxx"`
		Phone    string `json:"phone" binding:"required" example:"18717711819"`
	} `json:"admin" binding:"required"`
}

// SignUp will create  a count
// @Summary create a count
// @Accept json
// @Tags Admin
// @Security Bearer
// @Produce  json
// @Param admin body admin.ListSignUpParam true "sign up information"
// @Resource Admins
// @Router /v1/api/admin/sign_up [post]
// @Success 200 {object} model.AdminSerializer
func SignUp(c *gin.Context) {

	var param ListSignUpParam

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(400, c.AbortWithError(400, ErrorAdminSignInParam))
		return
	}

	// step 1 check name

	var ok bool

	if ok = crypt.CheckSignInName(param.Data.Name); ok != true {
		c.JSON(400, c.AbortWithError(400, ErrorAdminSignName))
		return
	}

	// step 2 check exists or not
	var admin model.Admin
	if dbError := initiator.POSTGRES.Where("name = ?", param.Data.Name).First(&admin).Error; dbError == nil {
		c.JSON(400, c.AbortWithError(400, ErrorAdminSignUpName))
		return
	}
	// step 3 create auth token and encrypted password

	var newAdmin model.Admin

	newAdmin = model.Admin{
		Name:              param.Data.Name,
		AuthToken:         crypt.GenerateToken(),
		EncryptedPassword: crypt.PassWordEncrypted(param.Data.Password),
		Phone:             param.Data.Phone,
		State:             StateAdmin,
	}
	if dbError := initiator.POSTGRES.Create(&newAdmin).Error; dbError != nil {
		c.JSON(400, c.AbortWithError(400, dbError))
		return
	}

	c.JSON(http.StatusOK, newAdmin.Serializer())

}
