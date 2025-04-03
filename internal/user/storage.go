package user

import "context"

//БД может быть любая, потому создаю интерфейс

type Storage interface {
	Create(ctx context.Context, user User) (string, error)
	FindOne(ctx context.Context, id string) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
}
