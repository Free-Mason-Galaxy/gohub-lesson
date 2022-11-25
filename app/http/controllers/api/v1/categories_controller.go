package v1

import (
	"gohub-lesson/app/models/category"
	"gohub-lesson/app/requests"

	"gohub-lesson/pkg/response"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	BaseController
}

func (class *CategoriesController) Index(ctx *gin.Context) {

	params, errs := requests.ValidatePagination(ctx)

	if errs.ErrsAbortWithStatusJSON(ctx) {
		return
	}

	categories, paging := category.Paginate(ctx, params.PerPage)

	response.JSON(ctx, gin.H{
		"data":  categories,
		"pager": paging,
	})
}

func (class *CategoriesController) Show(ctx *gin.Context) {
	categoryModel := category.Get(ctx.Param("id"))

	if categoryModel.NotExists() {
		response.Abort404(ctx)
		return
	}

	response.Data(ctx, categoryModel)
}

func (class *CategoriesController) Store(ctx *gin.Context) {

	data, errs := requests.ValidateCategorySave(ctx)

	if errs.ErrsAbortWithStatusJSON(ctx) {
		return
	}

	categoryModel := category.Category{
		Name:        data.Name,
		Description: data.Descr,
	}

	categoryModel.Create()

	if categoryModel.Exists() {
		response.Created(ctx, categoryModel)
		return
	}

	response.Abort500(ctx, "创建失败，请稍后尝试~")
}

func (class *CategoriesController) Update(ctx *gin.Context) {

	categoryModel := category.Get(ctx.Param("id"))

	if categoryModel.NotExists() {
		response.Abort404(ctx)
		return
	}

	// if ok := policies.CanModifyCategory(ctx, categoryModel); !ok {
	//     response.Abort403(ctx)
	//     return
	// }

	data, errs := requests.ValidateCategorySave(ctx)
	if errs.ErrsAbortWithStatusJSON(ctx) {
		return
	}
	// request := requests.CategoryRequest{}
	// bindOk, errs := requests.Validate(ctx, &request, requests.CategorySave)
	// if !bindOk {
	//     return
	// }
	// if len(errs) > 0 {
	//     response.ValidationError(ctx, errs)
	//     return
	// }
	//
	categoryModel.Name = data.Name
	categoryModel.Description = data.Descr

	rowsAffected := categoryModel.Save()

	if rowsAffected.ToBool() {
		response.Data(ctx, categoryModel)
		return
	}

	response.Abort500(ctx, "更新失败，请稍后尝试~")
}

func (class *CategoriesController) Delete(ctx *gin.Context) {

	categoryModel := category.Get(ctx.Param("id"))

	if categoryModel.NotExists() {
		response.Abort404(ctx)
		return
	}

	// if ok := policies.CanModifyCategory(ctx, categoryModel); !ok {
	//     response.Abort403(ctx)
	//     return
	// }

	rowsAffected := categoryModel.Delete()

	if rowsAffected > 0 {
		response.Success(ctx)
		return
	}

	response.Abort500(ctx, "删除失败，请稍后尝试~")
}
