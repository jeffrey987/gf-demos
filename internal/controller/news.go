package controller

import (
	"context"

	"github.com/gogf/gf-demos/v2/apiv1"
	"github.com/gogf/gf-demos/v2/internal/model"
	"github.com/gogf/gf-demos/v2/internal/service"
)

var News = cNews{}

type cNews struct{}

func (c *cNews) Create(ctx context.Context, req *apiv1.NewsCreateReq) (res *apiv1.NewsCreateRes, err error) {
	err = service.News().Create(ctx, model.NewsCreateInput{
		Title:   req.Title,
		Content: req.Content,
	})
	return
}
