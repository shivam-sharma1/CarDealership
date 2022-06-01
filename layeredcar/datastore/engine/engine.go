package engine

import (
	"assignments/layeredcar/model"
	"database/sql"
	"github.com/google/uuid"
)

type enginestorer struct {
	db *sql.DB
}

func New(Db *sql.DB) enginestorer {
	return enginestorer{db: Db}
}

func (e enginestorer) GetById(id uuid.UUID) (model.Engine, error) {
	var engine2 model.Engine
	res := e.db.QueryRow(SelectByIdQuery, id.String())
	err := res.Scan(&engine2.EngineId, &engine2.Displacement, &engine2.Noc, &engine2.Rng)
	if err != nil {
		return model.Engine{}, err
	}
	return engine2, nil
}

func (e enginestorer) Create(engine2 model.Engine) (model.Engine, error) {
	id := uuid.New()
	engine2.EngineId = id
	_, err := e.db.Exec(InsertQuery, engine2.EngineId.String(), engine2.Displacement, engine2.Noc, engine2.Rng)
	if err != nil {
		return model.Engine{}, err
	}
	return engine2, nil
}

func (e enginestorer) Update(id uuid.UUID, engine2 model.Engine) (model.Engine, error) {
	_, err := e.db.Exec(UpdateQuery, engine2.Displacement, engine2.Noc, engine2.Rng, id.String())
	if err != nil {
		return model.Engine{}, err
	}
	return engine2, nil
}

func (e enginestorer) Delete(id uuid.UUID) error {
	_, err := e.db.Exec(DeleteQuery, id.String())
	if err != nil {
		//log.Println("Error while deleting!", err)
		return err
	}
	return nil
}
