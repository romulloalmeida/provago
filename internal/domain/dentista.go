package domain

type Dentista struct {
	Id          	int     `json:"id"`
	Nome        	string  `json:"nome" binding:"required"`
	Sobrenome		string  `json:"sobrenome" binding:"required"`
	Matricula  		string  `json:"matricula" binding:"required"`
}
