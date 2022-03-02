package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/frozen599/job-interview/db"
	"github.com/golang/protobuf/proto"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "postgres"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	sqlDB, err := db.InitDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close(context.Background())

	// listener, err := net.Listen("tcp", "localhost:5001")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	accountRepo := db.NewAccountRepo(sqlDB)
	//clientConn, err := listener.Accept()
	if err != nil {
		log.Println(err)
		return
	}

	acc, _ := accountRepo.GetAccount(2)
	fmt.Println(acc)
	tx := accountRepo.CreateTransaction("account2", "account3", "300", acc.PrivKey)
	fmt.Println(tx)
	data, _ := proto.Marshal(tx)
	fmt.Println(data)

	srvConn, err := net.Dial("tcp", "localhost:5003")
	if err != nil {
		log.Panicln(err)
		return
	}

	_, err = io.Copy(srvConn, bytes.NewBuffer(data))
	if err != nil {
		log.Println(err)
		return
	}

	//io.WriteString(clientConn, "Send transaction to server3 successfully")

}
