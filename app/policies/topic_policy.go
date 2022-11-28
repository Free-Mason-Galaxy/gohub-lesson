package policies

import (
    "gohub-lesson/app/models/topic"
    "gohub-lesson/pkg/auth"

    "github.com/gin-gonic/gin"
)

func CanModifyTopic(c *gin.Context, topicModel topic.Topic) bool {
	return auth.CurrentUID(c) == topicModel.UserID
}

// func CanViewTopic(c *gin.Context, topicModel topic.Topic) bool {}
// func CanCreateTopic(c *gin.Context, topicModel topic.Topic) bool {}
// func CanUpdateTopic(c *gin.Context, topicModel topic.Topic) bool {}
// func CanDeleteTopic(c *gin.Context, topicModel topic.Topic) bool {}
