package user

import (
	"context"
	"fmt"
	"math"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"goyave.dev/goyave/v5/database"
	"goyave.dev/goyave/v5/util/typeutil"
	"goyave.dev/template/database/model"
	"goyave.dev/template/database/seed"
	"goyave.dev/template/dto"
)

type RepositoryMock struct {
	users []*model.User
}

func (r *RepositoryMock) First(_ context.Context, id uint) (*model.User, error) {
	u, ok := lo.Find(r.users, func(u *model.User) bool {
		return u.ID == id
	})
	if !ok {
		return nil, gorm.ErrRecordNotFound
	}
	return u, nil
}

func (r *RepositoryMock) Paginate(_ context.Context, page, pageSize int) (*database.Paginator[*model.User], error) {
	offset := (page - 1) * pageSize
	end := offset + pageSize
	records := r.users[offset:end]
	return &database.Paginator[*model.User]{
		Records:     &records,
		MaxPage:     int64(math.Ceil(float64(len(r.users)) / float64(pageSize))),
		CurrentPage: page,
		PageSize:    pageSize,
		Total:       int64(len(r.users)),
	}, nil
}

func TestUserService(t *testing.T) {

	t.Run("First", func(t *testing.T) {
		repo := &RepositoryMock{
			users: []*model.User{
				{
					ID:    4,
					Name:  "johndoe",
					Email: "johndoe@example.org",
				},
			},
		}
		s := NewService(repo)

		u, err := s.First(context.Background(), 4)
		expected := &dto.User{
			ID:    4,
			Name:  "johndoe",
			Email: "johndoe@example.org",
		}
		assert.Equal(t, expected, u)
		assert.NoError(t, err)
	})

	t.Run("Paginate", func(t *testing.T) {
		const userCount = 15
		records := database.NewFactory(seed.UserGenerator).Generate(userCount)
		repo := &RepositoryMock{
			users: records,
		}
		s := NewService(repo)

		cases := []struct {
			expect   *database.PaginatorDTO[*dto.User]
			page     int
			pageSize int
		}{
			{
				page:     1,
				pageSize: 3,
				expect: &database.PaginatorDTO[*dto.User]{
					Records:     typeutil.MustConvert[[]*dto.User](records[:3]),
					MaxPage:     5,
					CurrentPage: 1,
					PageSize:    3,
					Total:       userCount,
				},
			},
			{
				page:     2,
				pageSize: 3,
				expect: &database.PaginatorDTO[*dto.User]{
					Records:     typeutil.MustConvert[[]*dto.User](records[3:6]),
					MaxPage:     5,
					CurrentPage: 2,
					PageSize:    3,
					Total:       userCount,
				},
			},
		}

		for _, c := range cases {
			c := c
			t.Run(fmt.Sprintf("%d_%d", c.page, c.pageSize), func(t *testing.T) {
				paginator, err := s.Paginate(context.Background(), c.page, c.pageSize)
				assert.NoError(t, err)
				assert.Equal(t, c.expect, paginator)
			})
		}
	})
}
