package dataSrc

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	psqlInfo := os.Getenv("DATABASE_URL") // usa a variável Neon completa

	var db *sql.DB
	var err error

	// Retry para esperar o Postgres ficar pronto
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Printf("Tentativa %d: Erro ao abrir conexão: %v", i+1, err)
		} else if err = db.Ping(); err == nil {
			fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")
			return db, nil
		} else {
			log.Printf("Tentativa %d: Erro ao conectar ao DB: %v", i+1, err)
		}
		time.Sleep(3 * time.Second)
	}

	return nil, fmt.Errorf("não foi possível conectar ao banco de dados após várias tentativas: %v", err)
}
