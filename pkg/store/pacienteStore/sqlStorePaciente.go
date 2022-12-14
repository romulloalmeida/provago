package pacienteStore

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
type SqlStorePaciente interface {
	GetAll() []domain.Paciente
	GetByID(id int) (domain.Paciente, error)
	Create(p domain.Paciente) (domain.Paciente, error)
	Delete(id int) error
	Update(id int, p domain.Paciente) (domain.Paciente, error)
}

// relacionando com a entidade paciente
type sqlStorePaciente struct {
	db *sql.DB
}

func NewSQLStore() SqlStorePaciente {
	config.LoadConfig()
	database, err := config.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	return &sqlStorePaciente{
		db: database,
	}
}

func (s *sqlStorePaciente) GetAll() []domain.Paciente {
	var paciente domain.Paciente
	var pacientes []domain.Paciente
	rows, err := s.db.Query("SELECT * FROM pacientes")
	if err != nil {
		log.Println(err)
		return pacientes
	}

	for rows.Next() {
		if err := rows.Scan(
			&paciente.Id,
			&paciente.Nome,
			&paciente.Sobrenome,
			&paciente.RG,
			&paciente.DataCadastro); err != nil {
			return pacientes
		}
		pacientes = append(pacientes, paciente)
	}
	return pacientes
}

func (s *sqlStorePaciente) GetByID(id int) (domain.Paciente, error) {
	var paciente domain.Paciente
	rows, err := s.db.Query("SELECT * FROM pacientes WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return paciente, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&paciente.Id,
			&paciente.Nome,
			&paciente.Sobrenome,
			&paciente.RG,
			&paciente.DataCadastro); err != nil {
			log.Println(err.Error())
			return paciente, err
		}
	}
	return paciente, nil
}

func (s *sqlStorePaciente) Create(p domain.Paciente) (domain.Paciente, error) {
	timeParsed, err := time.ParseInLocation("02/01/2006", p.DataCadastro, time.Local)
	if err != nil {
		return domain.Paciente{}, errors.New("error while trying to save DataCadastro data")
	}
	result, err := s.db.Exec("INSERT INTO pacientes (nome,sobrenome,rg,data_cadastro) VALUES (?,?,?,?)", p.Nome, p.Sobrenome, p.RG, timeParsed)
	if err != nil {
		fmt.Println(err.Error())
		return domain.Paciente{}, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return domain.Paciente{}, err
	}
	p.Id = int(lastInsertId)
	return p, nil
}

func (s *sqlStorePaciente) Update(id int, paciente domain.Paciente) (domain.Paciente, error) {
	timeParsed, err := time.ParseInLocation("02/01/2006", paciente.DataCadastro, time.Local)
	if err != nil {
		return domain.Paciente{}, errors.New("error while trying to save DataCadastro data")
	}
	_, err = s.db.Exec("UPDATE pacientes SET nome = ? ,sobrenome = ?, rg = ?,data_cadastro = ? WHERE id = ?", paciente.Nome, paciente.Sobrenome, paciente.RG, timeParsed, id)
	if err != nil {
		log.Fatalln(err)
		return domain.Paciente{}, err
	}
	return paciente, nil
}

func (s *sqlStorePaciente) Delete(id int) error {

	result, err := s.db.Exec("DELETE FROM pacientes WHERE id=?", id)
	if err != nil {
		return err
	}
	count, err := result.LastInsertId()
	if count == 0 {
		return errors.New("paciente not found at database")
	}

	return nil
}
