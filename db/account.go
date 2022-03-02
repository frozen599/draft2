package db

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/frozen599/job-interview/models"
	"github.com/jackc/pgx/v4"
)

type Account struct {
	ID      int
	PubKey  string
	PrivKey string
	Address string
	Amount  float64
	Nonce   int
}

type AccountRepo struct {
	DB *pgx.Conn
}

func NewAccountRepo(db *pgx.Conn) *AccountRepo {
	return &AccountRepo{
		DB: db,
	}
}

func (r *AccountRepo) UpdateAccountAmount(id int, amount float64) error {
	stmt := fmt.Sprintf("UPDATE accounts SET amount = amount+ %f WHERE id=$1", amount)

	_, err := r.DB.Exec(context.Background(), stmt, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepo) GetAccount(id int) (*Account, error) {
	stmt :=
		`
	SELECT id, address, amount, nonce
	FROM accounts
	WHERE id=$1
	`
	var acc Account
	err := r.DB.QueryRow(context.Background(), stmt, id).Scan(&acc.ID, &acc.Address, &acc.Amount, &acc.Nonce)
	if err != nil {
		return nil, err
	}

	return &acc, nil
}

func generateSignature(tx *models.Transaction, privateKey string) {
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	data := []byte(tx.Amount)
	hash := crypto.Keccak256Hash(data)
	signature, err := crypto.Sign(hash.Bytes(), pk)
	if err != nil {
		log.Fatal(err)
	}

	tx.Sign = hexutil.Encode(signature)
}

func (r *AccountRepo) CreateTransaction(from, to, amount string, privateKey string) *models.Transaction {
	tx := models.Transaction{}
	tx.From = from
	tx.To = to
	balance, _ := strconv.ParseInt(amount, 10, 64)
	tx.Amount = fmt.Sprintf("%s", math.U256(big.NewInt(balance)))

	generateSignature(&tx, privateKey)

	return &tx
}
