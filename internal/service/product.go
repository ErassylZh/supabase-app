package service

import "context"

type Product interface {
	CreateFromAirtable(ctx context.Context)
}
