package v1

import (
	"gohub-lesson/app/models/user"
	"gohub-lesson/app/requests"
	"gohub-lesson/pkg/auth"
	"gohub-lesson/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseController
}

// CurrentUser 当前登录用户信息
func (class *UsersController) CurrentUser(ctx *gin.Context) {
	users := auth.CurrentUser(ctx)
	response.Data(ctx, users)
}

func (class *UsersController) UpdatePassword(ctx *gin.Context) {

	data, errs := requests.ValidateUserUpdatePassword(ctx)
	if errs.ErrsAbortWithStatusJSON(ctx) {
		return
	}

	currentUser := auth.CurrentUser(ctx)

	// 验证原始密码是否正确
	_, err := auth.Attempt(currentUser.Name, data.Password)

	if err != nil {
		// 失败，显示错误提示
		response.Unauthorized(ctx, "原密码不正确")
		return
	}

	// 更新密码为新密码
	currentUser.Password = data.NewPassword
	currentUser.Save()

	response.Success(ctx)
}

func (class *UsersController) UpdatePhone(ctx *gin.Context) {

	data, errs := requests.ValidateUserUpdatePhone(ctx)
	if errs.ErrsAbortWithStatusJSON(ctx) {
		return
	}

	currentUser := auth.CurrentUser(ctx)
	currentUser.Phone = data.Phone
	rowsAffected := currentUser.Save()

	if rowsAffected.ToBool() {
		response.Success(ctx)
		return
	}

	response.Abort500(ctx, "更新失败，请稍后尝试~")
}

func (class *UsersController) UpdateEmail(ctx *gin.Context) {

	data, errs := requests.ValidateUserUpdateEmail(ctx)
	if errs.ErrsAbortWithStatusJSON(ctx) {
		return
	}

	currentUser := auth.CurrentUser(ctx)
	currentUser.Email = data.Email
	rowsAffected := currentUser.Save()

	if rowsAffected.ToBool() {
		response.Success(ctx)
		return
	}

	// 失败，显示错误提示
	response.Abort500(ctx, "更新失败，请稍后尝试~")
}

func (class *UsersController) UpdateProfile(ctx *gin.Context) {
	data, errs := requests.ValidateUserUpdateProfile(ctx)
	if errs.ErrsAbortWithStatusJSON(ctx) {
		return
	}

	currentUser := auth.CurrentUser(ctx)
	currentUser.Name = data.Name
	currentUser.City = data.City
	currentUser.Introduction = data.Introduction
	rowsAffected := currentUser.Save()

	if rowsAffected.ToBool() {
		response.Data(ctx, currentUser)
		return
	}

	response.Abort500(ctx, "更新失败，请稍后尝试~")
}

func (class *UsersController) Index(ctx *gin.Context) {

	params, errs := requests.ValidatePagination(ctx)

	if errs.ErrsAbortWithStatusJSON(ctx) {
		return
	}

	data, pager := user.Paginate(ctx, params.PerPage)

	response.JSON(ctx, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (class *UsersController) Show(ctx *gin.Context) {
	userModel := user.Get(ctx.Param("id"))

	if userModel.NotExists() {
		response.Abort404(ctx)
		return
	}

	response.Data(ctx, userModel)
}

func (class *UsersController) Store(ctx *gin.Context) {

	// 例子
	// data, errs := requests.ValidateLoginByPhone(ctx)
	// if errs.ErrsAbortWithStatusJSON(ctx) {
	//    return
	// }
	// request := requests.UserRequest{}
	//
	// if ok := requests.Validate(ctx, &request, requests.UserSave); !ok {
	//     return
	// }
	//
	userModel := user.User{
		// FieldName: request.FieldName,
	}

	userModel.Create()

	if userModel.NotExists() {
		response.Created(ctx, userModel)
		return
	}

	response.Abort500(ctx, "创建失败，请稍后尝试~")
}

func (class *UsersController) Update(ctx *gin.Context) {

	userModel := user.Get(ctx.Param("id"))

	if userModel.NotExists() {
		response.Abort404(ctx)
		return
	}

	// if ok := policies.CanModifyUser(ctx, userModel); !ok {
	//     response.Abort403(c)
	//     return
	// }

	// 例子
	// data, errs := requests.ValidateLoginByPhone(ctx)
	// if errs.ErrsAbortWithStatusJSON(ctx) {
	//    return
	// }
	// request := requests.UserRequest{}
	// bindOk, errs := requests.Validate(ctx, &request, requests.UserSave)
	// if !bindOk {
	//     return
	// }
	// if len(errs) > 0 {
	//     response.ValidationError(ctx, errs)
	//     return
	// }
	//
	// userModel.FieldName = request.FieldName

	rowsAffected := userModel.Save()

	if rowsAffected > 0 {
		response.Data(ctx, userModel)
		return
	}

	response.Abort500(ctx, "更新失败，请稍后尝试~")
}

func (class *UsersController) Delete(ctx *gin.Context) {

	userModel := user.Get(ctx.Param("id"))

	if userModel.NotExists() {
		response.Abort404(ctx)
		return
	}

	// if ok := policies.CanModifyUser(ctx, userModel); !ok {
	//     response.Abort403(c)
	//     return
	// }

	rowsAffected := userModel.Delete()

	if rowsAffected > 0 {
		response.Success(ctx)
		return
	}

	response.Abort500(ctx, "删除失败，请稍后尝试~")
}
