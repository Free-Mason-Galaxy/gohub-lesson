package v1

import (
    "gohub-lesson/app/models/link"
    "gohub-lesson/pkg/response"

    "github.com/gin-gonic/gin"
)

type LinksController struct {
    BaseController
}

func (class *LinksController) Index(ctx *gin.Context) {
    links := link.All()

    response.Data(ctx, links)
}

func (class *LinksController) Show(ctx *gin.Context) {
    linkModel := link.Get(ctx.Param("id"))

    if linkModel.NotExists() {
        response.Abort404(ctx)
        return
    }

    response.Data(ctx, linkModel)
}

func (class *LinksController) Store(ctx *gin.Context) {

    // 例子
    // data, errs := requests.ValidateLoginByPhone(ctx)
    // if errs.ErrsAbortWithStatusJSON(ctx) {
    //    return
    // }
    // request := requests.LinkRequest{}
    //
    // if ok := requests.Validate(ctx, &request, requests.LinkSave); !ok {
    // 	return
    // }

    linkModel := link.Link{
        // FieldName:      request.FieldName,
    }

    linkModel.Create()

    if linkModel.NotExists() {
        response.Abort500(ctx, "创建失败，请稍后尝试~")
        return
    }

    response.Created(ctx, linkModel)
}

func (class *LinksController) Update(ctx *gin.Context) {

    linkModel := link.Get(ctx.Param("id"))

    if linkModel.NotExists() {
        response.Abort404(ctx)
        return
    }

    // if ok := policies.CanModifyLink(ctx, linkModel); !ok {
    // 	response.Abort403(ctx)
    // 	return
    // }

    // 例子
    // data, errs := requests.ValidateLoginByPhone(ctx)
    // if errs.ErrsAbortWithStatusJSON(ctx) {
    //    return
    // }
    // request := requests.LinkRequest{}
    // bindOk, errs := requests.Validate(ctx, &request, requests.LinkSave)
    // if !bindOk {
    // 	return
    // }
    // if len(errs) > 0 {
    // 	response.ValidationError(ctx, errs)
    // 	return
    // }
    //
    // linkModel.FieldName = request.FieldName

    rowsAffected := linkModel.Save()

    if rowsAffected > 0 {
        response.Data(ctx, linkModel)
        return
    }

    response.Abort500(ctx, "更新失败，请稍后尝试~")
}

func (class *LinksController) Delete(ctx *gin.Context) {

    linkModel := link.Get(ctx.Param("id"))

    if linkModel.NotExists() {
        response.Abort404(ctx)
        return
    }

    // if ok := policies.CanModifyLink(ctx, linkModel); !ok {
    // 	response.Abort403(ctx)
    // 	return
    // }

    rowsAffected := linkModel.Delete()

    if rowsAffected > 0 {
        response.Success(ctx)
        return
    }

    response.Abort500(ctx, "删除失败，请稍后尝试~")
}
