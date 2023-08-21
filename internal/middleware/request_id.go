package middleware

import (
	"context"
	"product-mall/internal/constants"

	util "product-mall/internal/tools"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func WithRequsetId() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuID := c.GetHeader(constants.HeaderXRequestID)
		if len(uuID) == 0 {
			uuID = uuid.New().String()
		}

		ctx := context.WithValue(c.Request.Context(), constants.HeaderXRequestID, uuID)
		c.Request = c.Request.WithContext(ctx)
		c.Header(constants.HeaderXRequestID, uuID)

		fields := logrus.Fields{}
		if c.Request.Method == "GET" {
			fields["param"] = c.Request.URL.Query()
			util.LogrusObj.WithContext(c).WithFields(fields).Info("request param")
		}
		// if c.Request.Method == "POST" {
		// 	// 使用GetRawData方法获取请求的body数据
		// 	body, err := c.GetRawData()
		// 	fmt.Printf("body %s error %s", string(body), err)
		// 	if err == nil {
		// 		// 打印请求的body数据
		// 		fields["param"] = string(body)
		// 		util.LogrusObj.WithContext(c).WithFields(fields).Info("request param")
		// 	}
		// }

		c.Next()
	}
}
