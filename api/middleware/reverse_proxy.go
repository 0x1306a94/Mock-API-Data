package middleware

import (
	"Mock-API-Data/api/controller"
	"Mock-API-Data/constant"
	"Mock-API-Data/model"
	"Mock-API-Data/proxy"
	"net/http"
	"net/url"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	proxyMap *sync.Map
)

func init() {
	proxyMap = &sync.Map{}
}

func ReverseProxyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		storageHelper, ok := controller.ExtractStorageHelper(c)
		if !ok {
			c.Writer.WriteHeader(http.StatusServiceUnavailable)
			c.Abort()
			return
		}
		projectKey := c.Request.Header.Get(constant.HTTPHeaderMockProjectKey)
		if projectKey == "" {
			c.Next()
			return
		}

		var project model.Project
		err := storageHelper.DB().Where("key = ?", projectKey).Find(&project).Error
		if err != nil {
			c.AbortWithError(http.StatusBadGateway, err)
			return
		}

		var targetProxy *proxy.ReverseProxy
		if v, ok := proxyMap.Load(project.Host); ok {
			if vv, ok := v.(*proxy.ReverseProxy); ok {
				targetProxy = vv
			}
		}

		if targetProxy == nil {
			targetURL, err := url.Parse(project.Host)
			if err != nil {
				c.AbortWithError(http.StatusBadGateway, err)
				return
			}

			targetProxy = proxy.NewSingleHostReverseProxy(targetURL)
			if targetProxy.Transport != nil {
				if v, ok := targetProxy.Transport.(*http.Transport); ok {
					v.TLSClientConfig.InsecureSkipVerify = project.InsecureSkipVerify
				}
			}
			proxyMap.Store(project.Host, targetProxy)
		}
		targetProxy.ServeHTTP(c.Writer, c.Request)
		// 中断后续 中间件
		c.Abort()
	}
}
