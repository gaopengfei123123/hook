package hook

import (
	"github.com/gin-gonic/gin"
	// "github.com/sirupsen/logrus"
)

// GithubSecret 中间件, 没啥用
func GithubSecret() gin.HandlerFunc {
	return func(c *gin.Context) {
		event := c.GetHeader("X-GitHub-Event")
		log.Infof("X-GitHub-Event: %v \n", event)
		c.Next()
	}
}
