package continuum

import "github.com/realForbis/continuum-go-sdk/pkg/types"

type DebugApi struct {
	client *Client
}

func NewDebugAPI(c *Client) *DebugApi {
	return &DebugApi{client: c}
}

func (l *DebugApi) BlockCacheCount() (map[string]uint64, error) {
	var r map[string]uint64
	err := l.client.getClient().Call(&r, "debug_blockCacheCount")
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (l *DebugApi) BlockLink(hash types.Hash) (map[string]types.Hash, error) {
	var r map[string]types.Hash
	err := l.client.getClient().Call(&r, "debug_blockLink", hash)
	if err != nil {
		return nil, err
	}
	return r, nil
}
