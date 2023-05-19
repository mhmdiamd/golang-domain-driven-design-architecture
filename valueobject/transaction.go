package valueobject

import (
	"time"

	"github.com/google/uuid"
)

// Transaction is value objet cause it has been no identity and immutable
type Transaction struct {
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
