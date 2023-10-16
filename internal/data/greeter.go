package data

import (
	"context"

	"github.com/nextmicro/logger"
	"github.com/nextmicro/next-layout/internal/biz"
)

type greeterRepo struct {
	data *Data
	log  logger.Logger
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger logger.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  logger,
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
	return nil, nil
}
