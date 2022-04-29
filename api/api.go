package api

import (
	"url-shortener/model"

	"github.com/gin-gonic/gin"
	"github.com/trunglen/g/x/rest"
)

type Api struct {
	rest.JsonRender
	*gin.RouterGroup
}

func NewApi(routerGroup *gin.RouterGroup) *Api {
	s := &Api{
		RouterGroup: routerGroup,
	}
	s.POST("url/shorten", s.shortenLink)
	s.POST("url/decode", s.shortenLink)
	s.GET("url/list", s.listLink)
	s.POST("url/redirect", s.shortenLink)
	//customer api

	return s
}

func (s *Api) shortenLink(c *gin.Context) {
	var shortenLink *model.Link
	c.BindJSON(&shortenLink)
	rest.AssertNil(shortenLink.Create())
	s.Success(c)
}

func (s *Api) listLink(c *gin.Context) {
	// var code = c.Query("code")
	var urls, err = new(model.Link).GetAll()
	rest.AssertNil(err)
	s.SendData(c, urls)
}

func (s *Api) decode(c *gin.Context) {
	// var code = c.Query("code")

}
