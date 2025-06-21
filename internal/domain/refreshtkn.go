package domain

import "github.com/google/uuid"

type RefreshToken struct {
	Value   uuid.UUID
	ExpTime int64
}

func (rt *RefreshToken) String() string {
	return rt.Value.String()
}
