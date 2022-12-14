package dentistaStore

import (
	"prova/config"
	"prova/internal/domain"
	"database/sql"
	"errors"
	"fmt"
	"log"
	//"time"
)

// definindo os metodos do SGBD
type SqlStoreDentista interface {
	GetAll() []domain.Dentista
	GetByID(id int) (domain.Dentista, error)
	Create(p domain.Dentista) (domain.Dentista, error)
	Delete(id int) error
	Update(id int, p domain.Dentista) (domain.Dentista, error)
}

// relacionando com a entidade dentista
type sqlStoreDentista struct {
	db *sql.DB
}

func NewSQLStore() SqlStoreDentista {
	config.LoadConfig()
	database, err := config.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	return &sqlStoreDentista{
		db: database,
	}
}

func (s *sqlStoreDentista) GetAll() []domain.Dentista {
	var dentista domain.Dentista
	var dentistas []domain.Dentista
	rows, err := s.db.Query("SELECT * FROM dentistas")
	if err != nil {
		log.Println(err)
		return dentistas
	}

	for rows.Next() {
		if err := rows.Scan(
			&dentista.Id,
			&dentista.Nome,
			&dentista.Sobrenome,
			&dentista.Matricula,); err != nil {
			return dentistas
		}
		dentistas = append(dentistas, dentista)
	}
	return dentistas
}

func (s *sqlStoreDentista) GetByID(id int) (domain.Dentista, error) {
	var dentista domain.Dentista
	rows, err := s.db.Query("SELECT * FROM dentistas WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return dentista, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&dentista.Id,
			&dentista.Nome,
			&dentista.Sobrenome,
			&dentista.Matricula); err != nil {
			log.Println(err.Error())
			return dentista, err
		}
	}
	return dentista, nil
}

func (s *sqlStoreDentista) Create(p domain.Dentista) (domain.Dentista, error) {
	result, err := s.db.Exec("INSERT INTO dentistas (nome,sobrenome,matricula) VALUES (?,?,?)", p.Nome, p.Sobrenome, p.Matricula)
	if err != nil {
		fmt.Println(err.Error())
		return domain.Dentista{}, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return domain.Dentista{}, err
	}
	p.Id = int(lastInsertId)
	return p, nil
}

func (s *sqlStoreDentista) Update(id int, dentista domain.Dentista) (domain.Dentista, error) {
	_, err := s.db.Exec("UPDATE dentistas SET nome = ? ,sobrenome = ?, matricula = ? WHERE id = ?", dentista.Nome, dentista.Sobrenome, dentista.Matricula, id)
	if err != nil {
		log.Fatalln(err)
		return domain.Dentista{}, err
	}
	return dentista, nil
}

func (s *sqlStoreDentista) Delete(id int) error {
	result, err := s.db.Exec("DELETE FROM dentistas WHERE id=?", id)
	if err != nil {
		return err
	}
	count, err := result.LastInsertId()
	if count == 0 {
		return errors.New("dentista not found at database")
	}

	return nil
}
