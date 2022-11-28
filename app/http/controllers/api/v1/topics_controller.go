package v1

import (
	"gohub-lesson/app/models/topic"
	"gohub-lesson/app/requests"
	"gohub-lesson/pkg/auth"
	"gohub-lesson/pkg/response"

	"github.com/gin-gonic/gin"
)

type TopicsController struct {
	BaseController
}

func (class *TopicsController) Index(ctx *gin.Context) {
	topics := topic.All()

	response.Data(ctx, topics)
}

func (class *TopicsController) Show(ctx *gin.Context) {
	topicModel := topic.Get(ctx.Param("id"))

	if topicModel.NotExists() {
		response.Abort404(ctx)
		return
	}

	response.Data(ctx, topicModel)
}

func (class *TopicsController) Store(ctx *gin.Context) {

	data, errs := requests.ValidateTopicSave(ctx)
	if errs.ErrsAbortWithStatusJSON(ctx) {
		return
	}
	// request := requests.TopicRequest{}
	//
	// if ok := requests.Validate(ctx, &request, requests.TopicSave); !ok {
	//     return
	// }

	topicModel := topic.Topic{
		Title:      data.Title,
		Body:       data.Body,
		UserID:     auth.CurrentUID(ctx),
		CategoryID: data.CategoryID,
	}

	topicModel.Create()

	if topicModel.NotExists() {
		response.Abort500(ctx, "创建失败，请稍后尝试~")
		return
	}

	response.Created(ctx, topicModel)
}

func (class *TopicsController) Update(ctx *gin.Context) {

	topicModel := topic.Get(ctx.Param("id"))

	if topicModel.NotExists() {
		response.Abort404(ctx)
		return
	}

	// if ok := policies.CanModifyTopic(ctx, topicModel); !ok {
	//     response.Abort403(ctx)
	//     return
	// }

	data, errs := requests.ValidateTopicSave(ctx)
	if errs.ErrsAbortWithStatusJSON(ctx) {
		return
	}
	// request := requests.TopicRequest{}
	// bindOk, errs := requests.Validate(ctx, &request, requests.TopicSave)
	// if !bindOk {
	//     return
	// }
	// if len(errs) > 0 {
	//     response.ValidationError(ctx, errs)
	//     return
	// }

	topicModel.Title = data.Title

	rowsAffected := topicModel.Save()

	if rowsAffected.ToBool() {
		response.Data(ctx, topicModel)
		return
	}

	response.Abort500(ctx, "更新失败，请稍后尝试~")
}

func (class *TopicsController) Delete(ctx *gin.Context) {

	topicModel := topic.Get(ctx.Param("id"))

	if topicModel.NotExists() {
		response.Abort404(ctx)
		return
	}

	// if ok := policies.CanModifyTopic(ctx, topicModel); !ok {
	//     response.Abort403(ctx)
	//     return
	// }

	rowsAffected := topicModel.Delete()

	if rowsAffected.ToBool() {
		response.Success(ctx)
		return
	}

	response.Abort500(ctx, "删除失败，请稍后尝试~")
}
