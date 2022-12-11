package usecase

import (
)

type PlaylistRepository interface {
	GetOne() ([]byte)
}