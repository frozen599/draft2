package db

import (
	"database/sql"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/frozen599/job-interview/models"
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
	DB *sql.DB
}

func NewAccountRepo(db *sql.DB) *AccountRepo {
	return &AccountRepo{
		DB: db,
	}
}

func (r *AccountRepo) UpdateAccountAmount(id int, amount float64) (int, error) {
	stmt := `
	UPDATE accounts
	SET amount = amount + ?
	WHERE id = ?
	`

	result, err := r.DB.Exec(stmt)
	if err != nil {
		return 0, err
	}

	accountID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(accountID), nil
}

func (r *AccountRepo) GetAccount(id int) (*Account, error) {
	stmt :=
		`
	SELECT id, address, amount, nonce
	FROM accounts
	WHERE id = ?
	`
	var acc Account
	err := r.DB.QueryRow(stmt, id).Scan(&acc.ID, &acc.Address, &acc.Amount, &acc.Nonce)
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

func (r *AccountRepo) CreateTransaction(from, to, amount string, fromAccountID int) *models.Transaction {
	tx := models.Transaction{}
	tx.From = from
	tx.To = to
	balance, _ := strconv.ParseInt(amount, 10, 64)
	tx.Amount = fmt.Sprintf("%s", math.U256(big.NewInt(balance)))

	acc, err := r.GetAccount(fromAccountID)
	if err != nil {
		log.Fatal(err)
	}

	generateSignature(&tx, acc.PrivKey)

	return &tx
}
