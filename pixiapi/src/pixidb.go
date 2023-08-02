package src

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

var dbhandle *sql.DB

func init() {
	dbhandle = getdbhandle()
	pingErr := dbhandle.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
func getdbhandle() *sql.DB {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "123456",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "mydb",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db
}
func InsertSpritetoDB(sp SpriteEnt) (int64, error) {
	result, err := dbhandle.Exec("insert into spriteinfo (name,url,width,height,x,y,visable) VALUES (?, ?, ?,?,?,?,?) ", sp.Name, sp.Url, sp.Width, sp.Height, sp.X, sp.Y, sp.Visable)
	if err != nil {
		return 0, fmt.Errorf("InsertSpritetoDB %q: %v", sp.Name, err)

	}
	id, err := result.LastInsertId()
	return id, nil
}
func GetSpritebyName(name string) (SpriteEnt, error) {
	var alb SpriteEnt
	var ID int
	results := dbhandle.QueryRow("SELECT * FROM spriteinfo WHERE name = ?", name)
	if err := results.Scan(&ID, &alb.Name, &alb.Url, &alb.Width, &alb.Height, &alb.X, &alb.Y, &alb.Visable); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("GetSpritebyName %d: no such sprite", name)
		}
		return alb, fmt.Errorf("GetSpritebyName %d: %v", name, err)
	}
	return alb, nil
}
func UpdataSprite(sp SpriteEnt) (int64, error) {
	_, err := dbhandle.Exec("update spriteinfo set name=?,url=?,width=?,height=?,x=?,y=?,visable=? where name=?", sp.Name, sp.Url, sp.Width, sp.Height, sp.X, sp.Y, sp.Visable, sp.Name)
	if err != nil {
		return 0, fmt.Errorf("GetALLSpritefromDB: %v", err)
	}
	return 1, nil
}
func GetALLSpritefromDB() ([]SpriteEnt, error) {
	var sprites []SpriteEnt
	results, err := dbhandle.Query("SELECT * FROM spriteinfo WHERE True")
	if err != nil {
		return nil, fmt.Errorf("GetALLSpritefromDB: %v", err)
	}
	defer results.Close()
	for results.Next() {
		var alb SpriteEnt
		var id int
		if err := results.Scan(&id, &alb.Name, &alb.Url, &alb.Width, &alb.Height, &alb.X, &alb.Y, &alb.Visable); err != nil {
			return nil, fmt.Errorf("GetALLSpritefromDB %q: %v", alb.Name, err)
		}
		sprites = append(sprites, alb)

	}
	if err := results.Err(); err != nil {
		return nil, fmt.Errorf("GetALLSpritefromDB: %v", err)
	}
	return sprites, nil
}
