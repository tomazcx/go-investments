package accountio

import (
	"time"

	"github.com/tomazcx/go-investments/internal/entities"
)

type FindAccountOutput struct {
	Account *entities.Account
	CreatedAt time.Time
}
