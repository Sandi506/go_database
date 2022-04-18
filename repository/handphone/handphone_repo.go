package handphone

import (
	"context"
	"go_database/entity"
)

type HandphoneRepository interface {
	Insert(ctx context.Context, handphone entity.Handphone) (entity.Handphone, error)
	FindById(ctx context.Context, id int32) (entity.Handphone, error)
	FindAll(ctx context.Context) ([]entity.Handphone, error)
	Update(ctx context.Context, id int32, handphone entity.Handphone) (entity.Handphone, error)
	Delete(ctx context.Context, id int32) (string, error)
}
