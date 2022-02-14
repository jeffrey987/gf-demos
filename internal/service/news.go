package service

import (
	"context"

	"github.com/gogf/gf-demos/v2/internal/model"
	"github.com/gogf/gf-demos/v2/internal/service/internal/dao"
	"github.com/gogf/gf-demos/v2/internal/service/internal/do"
	"github.com/gogf/gf/v2/database/gdb"
)

type (
	// sNews is service struct of module User.
	sNews struct{}
)

var (
	// insNews is the instance of service User.
	insNews = sNews{}
)

// News returns the interface of News service.
func News() *sNews {
	return &insNews
}

// Create creates user account.
func (s *sNews) Create(ctx context.Context, in model.NewsCreateInput) (err error) {
	// If Nickname is not specified, it then uses Passport as its default Nickname.
	if in.Title == "" {
		in.Title = "title"
	}

	return dao.News.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.News.Ctx(ctx).Data(do.News{
			Title:   in.Title,
			Content: in.Content,
		}).Insert()
		return err
	})
}
func (s *sNews) List(ctx context.Context, in string) (gdb.Result, error) {
	list, err := dao.News.Ctx(ctx).All()
	if err != nil {
		return nil, err
	}
	return list, nil
}
