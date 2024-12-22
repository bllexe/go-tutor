package main

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/go-sql-driver/mysql"
)

type Album struct {
	ID int64
	Title string
	Artist string
	Price float32
}

var db *sql.DB

func main() {
	//use environment variable
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recording",
	}
	var err error

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	// check if does exist or not
	albums,err := getAlbumsByArtist("John Coltrane")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums By Artist John Coltrane: %v\n", albums)

	// check album by id
	alb,err :=getAlbumsById(1)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("Album By Id: %v\n", alb)

	// ading new album
	id, err := addNewAlbum(Album{Title: "The Modern Sound of Betty Carter", Artist: "Betty Carter", Price: 49.99})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", id)
}

func getAlbumsById(id int64) (Album, error) {
	var alb Album
	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

func getAlbumsByArtist(name string) ([]Album, error) {
	var albums []Album
	row, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer row.Close()

	for row.Next() {
		var alb Album
		if err != row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price) {
			return nil, fmt.Errorf("scan album: %v", err)
		}
		albums = append(albums, alb)
	}
	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return albums, nil
}

func addNewAlbum(alb Album) (int64, error) {
	res, err := db.Exec("INSERT INTO album (title,artist,price) VALUES (?,?,?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("insert album: %v", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("insert album: %v", err)
	}
	return id, nil
}

func updateAlbumById(int int64, alb Album) (Album, error) {

	res,err :=db.Exec("UPDATE album SET title = ?, artist = ?, price = ? WHERE id = ?", alb.Title, alb.Artist, alb.Price, int)
	if err != nil {
		return alb, fmt.Errorf("update album: %v", err)
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return alb, fmt.Errorf("update album: %v", err)
	}
	if ra == 0 {
		return alb, fmt.Errorf("update album: %v", err)
	}
	return alb, nil
}



