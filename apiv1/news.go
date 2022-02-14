package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type NewsCreateReq struct {
	g.Meta  `path:"/news/create" method:"post" tags:"News" summary:"create a new News "`
	Title   string `v:"required|length:6,16"`
	Content string `v:"required|length:6,16"`
}
type NewsCreateRes struct{}
