package app

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

/**
-- Calling Features:
	var pl Player
	pl.id = 23
	pl.name = "Hans"
	id1, id2 := pl.save()
	log.Println(id1)
	log.Println(id2)
	pl.get()
	log.Println("Player retrieved:", pl.id, " Player name:", pl.name)
*/

func checkErr(err error) {
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

type Player struct {
	Id   int
	Name string
}

func (p *Player) Save() (int64, int64) {
	db, err := sql.Open("mysql", "dota2go:hansberg@/dota2go")
	checkErr(err)
	defer db.Close()
	stmt, err := db.Prepare("INSERT players SET id=?, name=?")
	checkErr(err)
	res, err := stmt.Exec(p.Id, p.Name)
	checkErr(err)
	lastRowInserted, _ := res.LastInsertId()
	rowsAffected, _ := res.RowsAffected()
	return lastRowInserted, rowsAffected
}

func (p *Player) GetPlayer(name string) {
	db, err := sql.Open("mysql", "dota2go:hansberg@/dota2go")
	checkErr(err)
	defer db.Close()
	rows := db.QueryRow("SELECT id, name FROM players WHERE name=?", name)
	rows.Scan(&p.Id, &p.Name)
	log.Println("Id Found:", p.Id)
}
