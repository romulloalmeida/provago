package domain

type Consulta struct {
	Id          	int     `json:"id"`
	Paciente       	string  `json:"paciente" binding:"required"`
	Dentista		string  `json:"dentista" binding:"required"`
	DataHora  		string  `json:"data_hora" binding:"required"`
}