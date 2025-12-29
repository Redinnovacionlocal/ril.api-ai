package sql

import (
	"log"

	"github.com/jmoiron/sqlx"
	"ril.api-ia/internal/domain/entity"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindByAiApiKey(aiApiKey string) (*entity.User, error) {
	var u entity.User
	err := r.db.Get(&u, "select id_usuario_recomendador, nombre, apellido, id_equipo, api_ai_token from sitio_usuarios_recomendador where api_ai_token=?", aiApiKey)
	log.Print(err)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) GetUserProfile(user *entity.User) (*entity.UserProfile, error) {
	var up entity.UserProfile

	query := `
        SELECT
            cms_pais.name                   AS country,
            scc.nombre                      AS charge,
            sp.provincia                    AS city,
            ss.nombre                       AS sector,
            sca.nombre                      AS area,
            scp.cargo_otro                  AS job_title
        FROM sitio_usuarios_recomendador sur
        LEFT JOIN cms_pais 
            ON cms_pais.id_pais = sur.id_pais
        LEFT JOIN sitio_conectarme_perfil scp 
            ON scp.id_usuario_recomendador = sur.id_usuario_recomendador
        LEFT JOIN sitio_conectarme_cargos scc 
            ON scc.id_cargo = scp.id_cargo
        LEFT JOIN sitio_provincias sp 
            ON sp.id = sur.id_provincia
        LEFT JOIN sitio_sectores ss 
            ON ss.id_sector = sur.id_sector
        LEFT JOIN sitio_conectarme_areasc sca 
            ON sca.id_areasc = scp.id_areasc
        WHERE sur.id_usuario_recomendador = ?
        LIMIT 1
    `
	err := r.db.Get(&up, query, user.Id)

	if err != nil {
		return nil, err
	}

	return &up, nil
}
