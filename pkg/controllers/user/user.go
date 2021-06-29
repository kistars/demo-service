package user

import (
	"context"
	"git.querycap.com/practice/srv-demo-suns/pkg/constants/errors"
	"git.querycap.com/practice/srv-demo-suns/pkg/constants/types"
	"git.querycap.com/practice/srv-demo-suns/pkg/models"
	"git.querycap.com/practice/srv-demo-suns/pkg/utils/db"
	"git.querycap.com/practice/srv-demo-suns/pkg/utils/idgen"
	"github.com/go-courier/sqlx-pg/pgbuilder"
	"github.com/go-courier/sqlx/v2"
)

// 创建
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

// 获取
func RetrieveUser(ctx context.Context, userID types.SFID) (*models.User, error) {
	d := db.DBExecutorFromContext(ctx)
	user := &models.User{}
	err := pgbuilder.Use(d).
		Select(nil).
		From(user).
		Where(user.FieldUserID().Eq(userID)).Scan(user)

	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return nil, errors.UserNotFoundError
		}
		return nil, err
	}

	return user, nil
}

// 删除
func DeleteUser(ctx context.Context, userID types.SFID) error {
	d := db.DBExecutorFromContext(ctx)
	user := &models.User{}

	expr := pgbuilder.Use(d).
		Delete(user).
		Where(user.FieldUserID().Eq(userID))

	return expr.Do()
}

// 更新
func UpdateUser(ctx context.Context, base models.UserBase, userID types.SFID) error {
	d := db.DBExecutorFromContext(ctx)
	user := &models.User{}
	user.UserBase = base
	user.MarkUpdatedAt()

	err := pgbuilder.Use(d).
		Update(user).
		SetFrom(user).
		Where(user.FieldUserID().Eq(userID)).
		Do()

	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return errors.UserNotFoundError
		}
		return errors.InternalServerError.StatusErr().WithDesc(err.Error())
	}

	return nil
}
