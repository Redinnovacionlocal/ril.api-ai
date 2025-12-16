package entity

type User struct {
	Id         int64   `db:"id_usuario_recomendador"`
	FirstName  string  `db:"nombre"`
	LastName   string  `db:"apellido"`
	IdTeam     int64   `db:"id_equipo"`
	ApiAiToken *string `db:"api_token"`
	UserProfile
}

type UserProfile struct {
	Country  *string `db:"country"`
	City     *string `db:"city"`
	Charge   *string `db:"charge"`
	Sector   *string `db:"sector"`
	Area     *string `db:"area"`
	JobTitle *string `db:"job_title"`
}
