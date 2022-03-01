package db

import "database/sql"

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
	UPDATE accounts a
	SET a.amount = a.amount + ?
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
