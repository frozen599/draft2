package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/frozen599/job-interview/db"
	"github.com/frozen599/job-interview/helpers"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "postgres"
)

func main() {
	// srv1, err := net.Listen("tcp", "localhost:5001")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// srv2, err := net.Listen("tcp", "localhost:5002")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// srv3, err := net.Listen("tcp", "localhost:5003")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for {
	// 	conn1, _ := srv1.Accept()
	// 	helpers.HandleConnSrv1(conn1)
	// 	conn2, _ := srv2.Accept()
	// 	helpers.HandleConnSrv2(conn2)
	// 	conn3, _ := srv3.Accept()
	// 	helpers.HandleConnSrv3(conn3)
	// }

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	sqlDB, err := db.InitDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close(context.Background())

	accountRepo := db.NewAccountRepo(sqlDB)
	srv3, err := net.Listen("tcp", "localhost:5003")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := srv3.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		helpers.HandleConnSrv3(conn, accountRepo)
	}
}
