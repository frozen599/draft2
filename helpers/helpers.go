package helpers

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"

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

func HandleConnSrv3(conn net.Conn, r *db.AccountRepo) {
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

		// See which server is sending request to server 3
		var acc *db.Account
		if strings.Compare(tx.From, "server1") == 0 {
			acc, err = r.GetAccount(1)
			if err != nil {
				io.WriteString(conn, err.Error())
				return
			}
		} else {
			acc, err = r.GetAccount(2)
			if err != nil {
				log.Println(err)
				return
			}
		}

		// verify the amount that was sent
		hashData := crypto.Keccak256Hash([]byte(tx.Amount))
		verified := crypto.VerifySignature([]byte(acc.PubKey), hashData.Bytes(), []byte(tx.Sign))
		if !verified {
			io.WriteString(conn, errors.New("error: transaction changed or malformed").Error())
			return
		}

		amount, _ := strconv.ParseFloat(tx.GetAmount(), 64)

		// update amount of account 3
		id, err := r.UpdateAccountAmount(3, amount)
		if err != nil {
			io.WriteString(conn, err.Error())
			return
		} else {
			io.WriteString(conn, fmt.Sprintf("Successfully updated amount of account %d", id))
		}
	}
}
