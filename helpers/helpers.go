package helpers

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/frozen599/job-interview/db"
	"github.com/frozen599/job-interview/models"
	"github.com/golang/protobuf/proto"
)

func HandleConnSrv1(conn net.Conn, r *db.AccountRepo) {
	defer conn.Close()
}

func HandleConnSrv2(conn net.Conn) {
	defer conn.Close()
}

func HandleConnSrv3(conn net.Conn, r *db.AccountRepo, fromAccID int) {
	defer conn.Close()

	for {
		var (
			buffer bytes.Buffer
			tx     models.Transaction
		)

		_, err := buffer.ReadFrom(conn)
		if err != nil {
			log.Fatal(err)
			return
		}

		err = proto.Unmarshal(buffer.Bytes(), &tx)
		if err != nil {
			log.Fatal(err)
			return
		}

		fromAccount, _ := r.GetAccount(fromAccID)
		hashData := crypto.Keccak256Hash([]byte(tx.Amount))
		verified := crypto.VerifySignature([]byte(fromAccount.PubKey), hashData.Bytes(), []byte(tx.Sign))
		if !verified {
			io.WriteString(conn, errors.New("error: transaction changed or malformed").Error())
			return
		}

		amount, _ := strconv.ParseFloat(tx.GetAmount(), 64)

		id, err := r.UpdateAccountAmount(3, amount)
		if err != nil {
			io.WriteString(conn, err.Error())
		} else {
			io.WriteString(conn, fmt.Sprintf("Successfully updated amount of account %d", id))
		}
	}
}
