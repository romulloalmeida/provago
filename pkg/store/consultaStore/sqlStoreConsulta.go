package consultaStore

import (
	"prova/config"
	"prova/internal/domain"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

// definindo os metodos do SGBD
type SqlStoreConsulta interface {
	GetAll() []domain.Consulta
	GetByID(id int) (domain.Consulta, error)
	Create(p domain.Consulta) (domain.Consulta, error)
	Delete(id int) error
	Update(id int, p domain.Consulta) (domain.Consulta, error)
}

// relacionando com a entidade consulta
type sqlStoreConsulta struct {
	db *sql.DB
}

func NewSQLStore() SqlStoreConsulta {
	config.LoadConfig()
	database, err := config.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	return &sqlStoreConsulta{
		db: database,
	}
}

func (s *sqlStoreConsulta) GetAll() []domain.Consulta {
	var consulta domain.Consulta
	var consultas []domain.Consulta
	rows, err := s.db.Query("SELECT * FROM consultas")
	if err != nil {
		log.Println(err)
		return consultas
	}

	for rows.Next() {
		if err := rows.Scan(
			&consulta.Id,
			&consulta.Paciente,
			&consulta.Dentista,
			&consulta.DataHora); err != nil {
			return consultas
		}
		consultas = append(consultas, consulta)
	}
	return consultas
}

func (s *sqlStoreConsulta) GetByID(id int) (domain.Consulta, error) {
	var consulta domain.Consulta
	rows, err := s.db.Query("SELECT * FROM consultas WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return consulta, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&consulta.Id,
			&consulta.Paciente,
			&consulta.Dentista,
			&consulta.DataHora); err != nil {
			log.Println(err.Error())
			return consulta, err
		}
	}
	return consulta, nil
}

func (s *sqlStoreConsulta) Create(p domain.Consulta) (domain.Consulta, error) {
	timeParsed, err := time.ParseInLocation("02/01/2006", p.DataHora, time.Local)
	if err != nil {
		return domain.Consulta{}, errors.New("error while trying to save DataHora data")
	}
	result, err := s.db.Exec("INSERT INTO consultas (paciente,dentista,data_hora) VALUES (?,?,?)", p.Paciente, p.Dentista, timeParsed)
	if err != nil {
		fmt.Println(err.Error())
		return domain.Consulta{}, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return domain.Consulta{}, err
	}
	p.Id = int(lastInsertId)
	return p, nil
}

func (s *sqlStoreConsulta) Update(id int, consulta domain.Consulta) (domain.Consulta, error) {
	timeParsed, err := time.ParseInLocation("02/01/2006", consulta.DataHora, time.Local)
	if err != nil {
		return domain.Consulta{}, errors.New("error while trying to save DataHora data")
	}
	_, err = s.db.Exec("UPDATE consultas SET paciente = ? ,dentista = ?, data_hora = ? WHERE id = ?", consulta.Paciente, consulta.Dentista, timeParsed, id)
	if err != nil {
		log.Fatalln(err)
		return domain.Consulta{}, err
	}
	return consulta, nil
}

func (s *sqlStoreConsulta) Delete(id int) error {

	result, err := s.db.Exec("DELETE FROM consultas WHERE id=?", id)
	if err != nil {
		return err
	}
	count, err := result.LastInsertId()
	if count == 0 {
		return errors.New("consulta not found at database")
	}

	return nil
}
