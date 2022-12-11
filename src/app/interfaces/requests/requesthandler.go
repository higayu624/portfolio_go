package requests

import (
)

type RequestHandler interface {
	Request()([]byte)
}