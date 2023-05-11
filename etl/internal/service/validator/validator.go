package validator

import (
	"context"
	"errors"
	"etl/internal/config"
)

type validateFunc func(ctx context.Context, rows []map[string]interface{}) ([]map[string]interface{}, error)

type Validator struct {
	//repo *repository.Repository
}

func (v *Validator) Run(ctx context.Context, rows []map[string]interface{}) ([]map[string]interface{}, error) {
	var checkList []validateFunc

	switch stage := config.GetStage(); stage {
	case config.STG:
		// TODO: добавить проверок для STG
		checkList = v.validateSTG()
	case config.ODS:
		// TODO: добавить проверок для ODS
		checkList = v.validateODS()
	case config.DDS:
		// TODO: добавить проверок для DDS
		checkList = v.validateDDS()
	case config.DM:
		// TODO: добавить проверок для DM
		checkList = v.validateDM()
	default:
		return nil, errors.New(stage)
	}

	var err error
	for _, check := range checkList {
		rows, err = check(ctx, rows)
		if err != nil {
			return nil, err
		}
	}
	return rows, nil
}

func (v *Validator) validateSTG() []validateFunc {
	return []validateFunc{
		// TODO: добавить проверок для STG
	}
}

func (v *Validator) validateODS() []validateFunc {
	return []validateFunc{
		// TODO: добавить проверок для ODS
	}
}

func (v *Validator) validateDDS() []validateFunc {
	return []validateFunc{
		// TODO: добавить проверок для DDS
	}
}

func (v *Validator) validateDM() []validateFunc {
	return []validateFunc{
		// TODO: добавить проверок для DM
	}
}
