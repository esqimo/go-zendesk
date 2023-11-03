package mock

import (
	"github.com/esqimo/go-zendesk/zendesk"
)

var _ zendesk.API = (*Client)(nil)
