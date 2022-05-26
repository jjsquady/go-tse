package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"go-tse/extractors"
	"go-tse/models"
	"golang.org/x/text/encoding/charmap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type customReader struct{ r io.Reader }

func (r *customReader) Read(b []byte) (n int, err error) {
	x := make([]byte, len(b))
	if n, err = r.r.Read(x); err != nil {
		return n, err
	}
	copy(b, bytes.Replace(x, []byte("\""), []byte(" "), -1))
	return n, nil
}

func main() {
	// dsn := "root:secret@tcp(6.tcp.ngrok.io:15699)/eleicao2020?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(
		&models.Candidato{},
		&models.Partido{},
	)
	if err != nil {
		return
	}

	go func() {
		f, err := os.Open("bweb_1t_SP_181120201549.csv")

		if err != nil {
			log.Fatal(err)
		}

		custom := &customReader{f}
		r := csv.NewReader(charmap.ISO8859_1.NewDecoder().Reader(custom))
		r.Comma = ';'

		for {

			record, err := r.Read()
			defer f.Close()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal(err)
			}

			candidatoExtractor := new(extractors.CandidatoExtractor)
			candidato := candidatoExtractor.Extract(record)
			if candidato != nil {
				fmt.Println(candidato)
				db.FirstOrCreate(&candidato)
			}

			partidoEx := new(extractors.PartidoExtractor)
			partido := partidoEx.Extract(record)
			if partido != nil {
				fmt.Println(partido)
				db.FirstOrCreate(&partido)
			}

			fmt.Println("Lendo do arquivo A")
		}
	}()

	// go func() {
	// 	f, err := os.Open("bweb_1t_SP_181120201549.csv")
	//
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	custom := &customReader{f}
	// 	r := csv.NewReader(charmap.ISO8859_1.NewDecoder().Reader(custom))
	// 	r.Comma = ';'
	//
	// 	for {
	//
	// 		record, err := r.Read()
	// 		defer f.Close()
	//
	// 		if err == io.EOF {
	// 			break
	// 		}
	//
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	//
	// 		fmt.Println(record)
	// 		fmt.Println("Lendo arquivo B...")
	// 	}
	// }()

	fmt.Println("Processando...")
	for {
	}
}
