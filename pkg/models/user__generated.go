package models

import (
	fmt "fmt"
	time "time"

	git_querycap_com_tools_datatypes "git.querycap.com/tools/datatypes"
	github_com_go_courier_sqlx_v2 "github.com/go-courier/sqlx/v2"
	github_com_go_courier_sqlx_v2_builder "github.com/go-courier/sqlx/v2/builder"
	github_com_go_courier_sqlx_v2_datatypes "github.com/go-courier/sqlx/v2/datatypes"
	srv_demo_suns_pkg_constants_types "srv-demo-suns/pkg/constants/types"
)

func (User) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (User) Indexes() github_com_go_courier_sqlx_v2_builder.Indexes {
	return github_com_go_courier_sqlx_v2_builder.Indexes{
		"I_name": []string{
			"Name",
		},
		"i_account_type": []string{
			"UserType",
		},
	}
}

func (User) UniqueIndexIPhone() string {
	return "I_phone"
}

func (User) UniqueIndexIUserID() string {
	return "I_user_id"
}

func (User) UniqueIndexes() github_com_go_courier_sqlx_v2_builder.Indexes {
	return github_com_go_courier_sqlx_v2_builder.Indexes{
		"I_phone": []string{
			"Phone",
			"DeletedAt",
		},
		"I_user_id": []string{
			"UserID",
			"DeletedAt",
		},
	}
}

var UserTable *github_com_go_courier_sqlx_v2_builder.Table

func init() {
	UserTable = DB.Register(&User{})
}

type UserIterator struct {
}

func (UserIterator) New() interface{} {
	return &User{}
}

func (UserIterator) Resolve(v interface{}) *User {
	return v.(*User)
}

func (User) TableName() string {
	return "t_user"
}

func (User) TableDescription() []string {
	return []string{
		"用户表",
	}
}

func (User) ColDescriptions() map[string][]string {
	return map[string][]string{
		"CreatedAt": []string{
			"创建时间",
		},
		"DeletedAt": []string{
			"删除时间",
		},
		"ID": []string{
			"自增 ID",
		},
		"Name": []string{
			"名称",
		},
		"Phone": []string{
			"手机号",
		},
		"UpdatedAt": []string{
			"更新时间",
		},
		"UserID": []string{
			"用户唯一 ID",
		},
		"UserType": []string{
			"用户类型",
		},
	}
}

func (User) FieldKeyID() string {
	return "ID"
}

func (m *User) FieldID() *github_com_go_courier_sqlx_v2_builder.Column {
	return UserTable.F(m.FieldKeyID())
}

func (User) FieldKeyUserID() string {
	return "UserID"
}

func (m *User) FieldUserID() *github_com_go_courier_sqlx_v2_builder.Column {
	return UserTable.F(m.FieldKeyUserID())
}

func (User) FieldKeyName() string {
	return "Name"
}

func (m *User) FieldName() *github_com_go_courier_sqlx_v2_builder.Column {
	return UserTable.F(m.FieldKeyName())
}

func (User) FieldKeyPhone() string {
	return "Phone"
}

func (m *User) FieldPhone() *github_com_go_courier_sqlx_v2_builder.Column {
	return UserTable.F(m.FieldKeyPhone())
}

func (User) FieldKeyUserType() string {
	return "UserType"
}

func (m *User) FieldUserType() *github_com_go_courier_sqlx_v2_builder.Column {
	return UserTable.F(m.FieldKeyUserType())
}

func (User) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *User) FieldCreatedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return UserTable.F(m.FieldKeyCreatedAt())
}

func (User) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *User) FieldUpdatedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return UserTable.F(m.FieldKeyUpdatedAt())
}

func (User) FieldKeyDeletedAt() string {
	return "DeletedAt"
}

func (m *User) FieldDeletedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return UserTable.F(m.FieldKeyDeletedAt())
}

func (User) ColRelations() map[string][]string {
	return map[string][]string{
		"UserID": []string{
			"User",
			"UserID",
		},
	}
}

func (m *User) IndexFieldNames() []string {
	return []string{
		"ID",
		"Name",
		"Phone",
		"UserID",
		"UserType",
	}
}

func (m *User) ConditionByStruct(db github_com_go_courier_sqlx_v2.DBExecutor) github_com_go_courier_sqlx_v2_builder.SqlCondition {
	table := db.T(m)
	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m)

	conditions := make([]github_com_go_courier_sqlx_v2_builder.SqlCondition, 0)

	for _, fieldName := range m.IndexFieldNames() {
		if v, exists := fieldValues[fieldName]; exists {
			conditions = append(conditions, table.F(fieldName).Eq(v))
			delete(fieldValues, fieldName)
		}
	}

	if len(conditions) == 0 {
		panic(fmt.Errorf("at least one of field for indexes has value"))
	}

	for fieldName, v := range fieldValues {
		conditions = append(conditions, table.F(fieldName).Eq(v))
	}

	condition := github_com_go_courier_sqlx_v2_builder.And(conditions...)

	condition = github_com_go_courier_sqlx_v2_builder.And(condition, table.F("DeletedAt").Eq(0))
	return condition
}

func (m *User) Create(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_go_courier_sqlx_v2.InsertToDB(db, m, nil))
	return err

}

func (m *User) CreateOnDuplicateWithUpdateFields(db github_com_go_courier_sqlx_v2.DBExecutor, updateFields []string) error {

	if len(updateFields) == 0 {
		panic(fmt.Errorf("must have update fields"))
	}

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, updateFields...)

	delete(fieldValues, "ID")

	table := db.T(m)

	cols, vals := table.ColumnsAndValuesByFieldValues(fieldValues)

	fields := make(map[string]bool, len(updateFields))
	for _, field := range updateFields {
		fields[field] = true
	}

	for _, fieldNames := range m.UniqueIndexes() {
		for _, field := range fieldNames {
			delete(fields, field)
		}
	}

	if len(fields) == 0 {
		panic(fmt.Errorf("no fields for updates"))
	}

	for field := range fieldValues {
		if !fields[field] {
			delete(fieldValues, field)
		}
	}

	additions := github_com_go_courier_sqlx_v2_builder.Additions{}

	switch db.Dialect().DriverName() {
	case "mysql":
		additions = append(additions, github_com_go_courier_sqlx_v2_builder.OnDuplicateKeyUpdate(table.AssignmentsByFieldValues(fieldValues)...))
	case "postgres":
		indexes := m.UniqueIndexes()
		fields := make([]string, 0)
		for _, fs := range indexes {
			fields = append(fields, fs...)
		}
		indexFields, _ := db.T(m).Fields(fields...)

		additions = append(additions,
			github_com_go_courier_sqlx_v2_builder.OnConflict(indexFields).
				DoUpdateSet(table.AssignmentsByFieldValues(fieldValues)...))
	}

	additions = append(additions, github_com_go_courier_sqlx_v2_builder.Comment("User.CreateOnDuplicateWithUpdateFields"))

	expr := github_com_go_courier_sqlx_v2_builder.Insert().Into(table, additions...).Values(cols, vals...)

	_, err := db.ExecExpr(expr)
	return err

}

func (m *User) DeleteByStruct(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(m.ConditionByStruct(db)),
				github_com_go_courier_sqlx_v2_builder.Comment("User.DeleteByStruct"),
			),
	)

	return err
}

func (m *User) FetchByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("User.FetchByID"),
			),
		m,
	)

	return err
}

func (m *User) UpdateByIDWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("User.UpdateByIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByID(db)
	}

	return nil

}

func (m *User) UpdateByIDWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *User) FetchByIDForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.ForUpdate(),
				github_com_go_courier_sqlx_v2_builder.Comment("User.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *User) DeleteByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("User.DeleteByID"),
			))

	return err
}

func (m *User) SoftDeleteByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValues{}
	if _, ok := fieldValues["DeletedAt"]; !ok {
		fieldValues["DeletedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("User.SoftDeleteByID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *User) FetchByUserID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("UserID").Eq(m.UserID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("User.FetchByUserID"),
			),
		m,
	)

	return err
}

func (m *User) UpdateByUserIDWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("UserID").Eq(m.UserID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("User.UpdateByUserIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByUserID(db)
	}

	return nil

}

func (m *User) UpdateByUserIDWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByUserIDWithMap(db, fieldValues)

}

func (m *User) FetchByUserIDForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("UserID").Eq(m.UserID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.ForUpdate(),
				github_com_go_courier_sqlx_v2_builder.Comment("User.FetchByUserIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *User) DeleteByUserID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("UserID").Eq(m.UserID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("User.DeleteByUserID"),
			))

	return err
}

func (m *User) SoftDeleteByUserID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValues{}
	if _, ok := fieldValues["DeletedAt"]; !ok {
		fieldValues["DeletedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("UserID").Eq(m.UserID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("User.SoftDeleteByUserID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *User) FetchByPhone(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("Phone").Eq(m.Phone),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("User.FetchByPhone"),
			),
		m,
	)

	return err
}

func (m *User) UpdateByPhoneWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("Phone").Eq(m.Phone),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("User.UpdateByPhoneWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByPhone(db)
	}

	return nil

}

func (m *User) UpdateByPhoneWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByPhoneWithMap(db, fieldValues)

}

func (m *User) FetchByPhoneForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("Phone").Eq(m.Phone),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.ForUpdate(),
				github_com_go_courier_sqlx_v2_builder.Comment("User.FetchByPhoneForUpdate"),
			),
		m,
	)

	return err
}

func (m *User) DeleteByPhone(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("Phone").Eq(m.Phone),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("User.DeleteByPhone"),
			))

	return err
}

func (m *User) SoftDeleteByPhone(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValues{}
	if _, ok := fieldValues["DeletedAt"]; !ok {
		fieldValues["DeletedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("Phone").Eq(m.Phone),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("User.SoftDeleteByPhone"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *User) List(db github_com_go_courier_sqlx_v2.DBExecutor, condition github_com_go_courier_sqlx_v2_builder.SqlCondition, additions ...github_com_go_courier_sqlx_v2_builder.Addition) ([]User, error) {

	list := make([]User, 0)

	table := db.T(m)
	_ = table

	condition = github_com_go_courier_sqlx_v2_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_go_courier_sqlx_v2_builder.Addition{
		github_com_go_courier_sqlx_v2_builder.Where(condition),
		github_com_go_courier_sqlx_v2_builder.Comment("User.List"),
	}

	if len(additions) > 0 {
		finalAdditions = append(finalAdditions, additions...)
	}

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(db.T(m), finalAdditions...),
		&list,
	)

	return list, err

}

func (m *User) Count(db github_com_go_courier_sqlx_v2.DBExecutor, condition github_com_go_courier_sqlx_v2_builder.SqlCondition, additions ...github_com_go_courier_sqlx_v2_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	condition = github_com_go_courier_sqlx_v2_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_go_courier_sqlx_v2_builder.Addition{
		github_com_go_courier_sqlx_v2_builder.Where(condition),
		github_com_go_courier_sqlx_v2_builder.Comment("User.Count"),
	}

	if len(additions) > 0 {
		finalAdditions = append(finalAdditions, additions...)
	}

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(
			github_com_go_courier_sqlx_v2_builder.Count(),
		).
			From(db.T(m), finalAdditions...),
		&count,
	)

	return count, err

}

func (m *User) BatchFetchByIDList(db github_com_go_courier_sqlx_v2.DBExecutor, values []uint64) ([]User, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}

func (m *User) BatchFetchByNameList(db github_com_go_courier_sqlx_v2.DBExecutor, values []string) ([]User, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("Name").In(values)

	return m.List(db, condition)

}

func (m *User) BatchFetchByPhoneList(db github_com_go_courier_sqlx_v2.DBExecutor, values []string) ([]User, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("Phone").In(values)

	return m.List(db, condition)

}

func (m *User) BatchFetchByUserIDList(db github_com_go_courier_sqlx_v2.DBExecutor, values []git_querycap_com_tools_datatypes.SFID) ([]User, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("UserID").In(values)

	return m.List(db, condition)

}

func (m *User) BatchFetchByUserTypeList(db github_com_go_courier_sqlx_v2.DBExecutor, values []srv_demo_suns_pkg_constants_types.UserType) ([]User, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("UserType").In(values)

	return m.List(db, condition)

}
