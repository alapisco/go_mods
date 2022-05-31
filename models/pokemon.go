package models

type Pokemon struct {
	ID			int		`json:"id"`
	NAME 		string	`json:"name"`
	TYPE1 		string	`json:"type1"`
	TYPE2 		string	`json:"type2"`
	TOTAL		int		`json:"total"`
	HP			int		`json:"hp"`
	ATTACK		int		`json:"attack"`
	DEFENSE		int		`json:"defendse"`
	SPATK		int		`json:"sp_atk"`
	SPDEF		int		`json:"sp_def"`
	SPEED		int		`json:"speed"`
	GENERATION	int		`json:"generation"`
	LEGENDARY	bool	`json:"legendary"`
}