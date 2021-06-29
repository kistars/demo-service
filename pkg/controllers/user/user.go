package user

import (
	"context"
	"github.com/go-courier/sqlx-pg/pgbuilder"
	"github.com/go-courier/sqlx/v2"
	"srv-demo-suns/pkg/constants/errors"
	"srv-demo-suns/pkg/models"
	"srv-demo-suns/pkg/utils/db"
	"srv-demo-suns/pkg/utils/idgen"
)

func CreateUser(ctx context.Context, userBase models.UserBase) (*models.User, error) {
	d := db.DBExecutorFromContext(ctx)
	userID, err := idgen.IDGenFromContext(ctx).GenerateID()
	if err != nil {
		return nil, err
	}

	user := &models.User{}
	user.UserID = userID
	user.UserBase = userBase
	user.MarkCreatedAt()

	err = pgbuilder.Use(d).
		Insert().Into(user).
		ValuesFrom(user).Do()

	if err != nil {
		if sqlx.DBErr(err).IsConflict() {
			return nil, errors.ConflictError
		}
		return nil, errors.InternalServerError.StatusErr().WithDesc(err.Error())
	}

	return user, nil
}
