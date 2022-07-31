package continuum

import (
	"math/big"

	"github.com/realForbis/continuum-go-sdk/pkg/types"
	"github.com/realForbis/continuum-go-sdk/pkg/util"
)

type UtilApi struct {
	client *Client
}

// NewUtilAPI creates unit module for client
func NewUtilAPI(c *Client) *UtilApi {
	return &UtilApi{client: c}
}

// Decrypt decrypts cryptograph to raw by passphrase
func (u *UtilApi) Decrypt(cryptograph string, passphrase string) (string, error) {
	return util.Decrypt(cryptograph, passphrase)
}

// Encrypt encrypts raw to cryptograph by passphrase
func (u *UtilApi) Encrypt(raw string, passphrase string) (string, error) {
	return util.Encrypt(raw, passphrase)
}

type APIBalance struct {
	*big.Float
}

func (b *APIBalance) MarshalText() ([]byte, error) {
	return []byte(b.String()), nil
}

func (b *APIBalance) String() string {
	return b.Text('f', -1)
}

// RawToBalance transforms Q amount from raw to unit
func (u *UtilApi) RawToBalance(balance types.Balance, unit string) (APIBalance, error) {
	var b APIBalance
	if err := u.client.getClient().Call(&b, "util_rawToBalance", balance, unit); err != nil {
		return APIBalance{big.NewFloat(0)}, err
	}
	return b, nil
}

// RawToBalance transforms token (not Q) amount from raw
func (u *UtilApi) RawToBalanceForToken(balance types.Balance, tokenName string) (APIBalance, error) {
	var b APIBalance
	if err := u.client.getClient().Call(&b, "util_rawToBalance", balance, "", tokenName); err != nil {
		return APIBalance{big.NewFloat(0)}, err
	}
	return b, nil
}

// RawToBalance transforms Q amount from unit to raw
func (u *UtilApi) BalanceToRaw(balance types.Balance, unit string) (types.Balance, error) {
	var b types.Balance
	if err := u.client.getClient().Call(&b, "util_balanceToRaw", balance, unit); err != nil {
		return types.ZeroBalance, err
	}
	return b, nil
}

// RawToBalance transforms token (not Q) amount to raw
func (u *UtilApi) BalanceToRawForToken(balance types.Balance, tokenName string) (types.Balance, error) {
	var b types.Balance
	if err := u.client.getClient().Call(&b, "util_balanceToRaw", balance, "", tokenName); err != nil {
		return types.ZeroBalance, err
	}
	return b, nil
}
