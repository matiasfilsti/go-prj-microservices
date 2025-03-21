package contracts

import "context"

type MessageService interface {
	Send(ctx context.Context) error
}
