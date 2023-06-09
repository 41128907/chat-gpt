package gateway

import (
	"context"

	"github.com/devfullcycle/fclx/chatservice/internal/domain/entity"
)

type ChatGateway interface {
	CreateChat(ctx context.Context, chat *entity.Chat) error
	FindChatById(ctx context.Context, chat *entity.Chat) error
	SaveChat(ctx context.Context, chat *entity.Chat) error
}
