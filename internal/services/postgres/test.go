package postgres

import "context"

func (dao *Postgres) PostgresPing(ctx context.Context) error {
	err := dao.Ping()
	if err != nil {
		return err
	}

	return  nil
}

