package tokens

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math"
	"math/big"
	"strings"
	"ubiq-explorer/models"
	"ubiq-explorer/node"
)

type TokenUtils struct {
	//Client      *ethclient.Client
	Tokens      map[string]*models.TokenInfo
	Ignored     map[string]bool
	Context     context.Context
	BlockNumber *big.Int
}

func NewTokenUtils(block *big.Int) *TokenUtils {
	return &TokenUtils{
		//node.Client(),
		make(map[string]*models.TokenInfo),
		make(map[string]bool),
		context.TODO(), // TODO - obviously
		block,
	}
}

/* Just a check to see if there is code at an address to avoid doing other operations unnecessarily.
if we already know there is no code at this address it just returns false without checking again.
*/
func (t *TokenUtils) HasCode(address common.Address) bool {
	if t.GetCachedToken(address) != nil {
		return true
	}
	code, err := node.Client().CodeAt(t.Context, address, t.BlockNumber)
	if err != nil || len(code) == 0 {
		return false
	}
	return true
}

/* Check if we are already ignoring this address because there isn't an ERC20 compliant token there */
func (t *TokenUtils) GetIgnoredAddress(address common.Address) bool {
	if t.Ignored[address.String()] {
		return true
	}
	return false
}

/* Add ignored address to the cache */
func (t *TokenUtils) AddIgnoredAddress(address common.Address) {
	t.Ignored[address.String()] = true
}

/* Cache this token info */
func (t *TokenUtils) AddCachedToken(address common.Address, tokenInfo *models.TokenInfo) {
	t.Tokens[address.String()] = tokenInfo
}

/* Get a cached token */
func (t *TokenUtils) GetCachedToken(address common.Address) *models.TokenInfo {
	if t.Tokens[address.String()] != nil {
		return t.Tokens[address.String()]
	}
	return nil
}

/* I think I planned on grabbing the token from the DB here, but I haven't done it or felt the need to */
func (t *TokenUtils) GetSavedToken(address common.Address) *models.TokenInfo {
	return &models.TokenInfo{}
}

/* Same as above */
func (t *TokenUtils) SaveToken(address common.Address, token *models.TokenInfo) (bool, error) {
	return true, nil
}

/*
	Grab token info using standard ABI - see tokens.go
	It's generated via abigen, but the code it generates when only running an ubiq node is not correct, so it's statically committed here
*/
func (t *TokenUtils) GetTokenInfo(address common.Address) (*models.TokenInfo, bool, error) {
	cached := t.GetCachedToken(address)
	if cached != nil {
		return cached, false, nil
	}
	if t.GetIgnoredAddress(address) {
		return nil, false, errors.New("This address is ingored because no token contract was found")
	}
	// Don't try to call the ABI on an address w/ no code
	if !t.HasCode(address) {
		t.AddIgnoredAddress(address)
		return nil, false, errors.New("There is no code at this address")
	}
	abi, err := NewToken(address, node.Client())
	if err != nil {
		t.AddIgnoredAddress(address)
		return nil, false, errors.New("This address has a contract that does not have a token or does not follow the standard ABI token specifications")
	}
	name, err := abi.Name(nil)
	if err != nil {
		return nil, false, err
	}
	symbol, err := abi.Symbol(nil)
	if err != nil {
		return nil, false, err
	}
	decimals, err := abi.Decimals(nil)
	if err != nil {
		return nil, false, err
	}

	i := &models.TokenInfo{
		name,
		address,
		symbol,
		decimals,
	}
	t.AddCachedToken(address, i)
	return i, true, nil
}

/* Get token transactions including transferred value */
func (t *TokenUtils) GetTransactionInfo(txn *models.Transaction) (map[int]*models.TokenTransaction, error) {
	rcpt, err := node.Client().TransactionReceipt(t.Context, txn.Hash)
	if err != nil {
		return nil, err
	}
	tokenTransactionInfo := make(map[int]*models.TokenTransaction)
	for i, log := range rcpt.Logs {
		tokenInfo, _, err := t.GetTokenInfo(txn.To)
		if err != nil {
			return nil, err
		}

		xferAmount, err := t.GetTokenTransferValue(common.Bytes2Hex(log.Data), tokenInfo.Decimals)
		if err != nil {
			fmt.Println("Failed to parse the transfer amount")
			return nil, err
		}

		// Not sure what this is exactly, because I am not an expert, but sometimes there is no destination, maybe burn?
		// There are only 3 or 4 transactions without a destination address, but it causes the daemon to crash if we don't handle it
		to := common.Address{}
		if len(log.Topics) > 2 {
			to = common.HexToAddress(log.Topics[2].String())
		}
		tokenTxn := &models.TokenTransaction{
			Address:   txn.To,
			Hash:      txn.Hash,
			Timestamp: txn.Timestamp,
			From:      common.HexToAddress(log.Topics[1].String()),
			To:        to,
			Value:     xferAmount,
			TokenInfo: tokenInfo,
		}
		tokenTransactionInfo[i] = tokenTxn
	}
	return tokenTransactionInfo, nil
}

/*
   Convert data to token transfer value including decimal placement
   I would wager that there is a better way to do this, but this is how I'm doing it
*/
func (t *TokenUtils) GetTokenTransferValue(hash string, decimals uint8) (*big.Float, error) {
	hash = strings.Replace(hash, "0x", "", -1)
	amount, _, err := new(big.Float).Parse(hash, 16)
	if err != nil {
		return nil, err
	}
	d := new(big.Float).SetFloat64(math.Pow10(int(decimals)))
	z := new(big.Float).Quo(amount, d)
	return z, nil
}

func (t *TokenUtils) GetTokenBalance(tokenAddress common.Address, address common.Address) (*big.Float, error) {
	tokenInfo, _, err := t.GetTokenInfo(tokenAddress)
	if err != nil {
		return nil, err
	}

	abi, err := NewToken(tokenAddress, node.Client())
	if err != nil {
		return nil, err
	}

	balance, err := abi.BalanceOf(nil, address)
	if err != nil {
		return nil, err
	}

	bFloat, _, err := new(big.Float).Parse(balance.String(), 10)

	if err != nil {
		return nil, err
	}
	d := new(big.Float).SetFloat64(math.Pow10(int(tokenInfo.Decimals)))
	z := new(big.Float).Quo(bFloat, d)

	return z, nil
}
