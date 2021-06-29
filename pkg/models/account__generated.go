package models

import (
	fmt "fmt"
	time "time"

	git_querycap_com_tools_datatypes "git.querycap.com/tools/datatypes"
	github_com_go_courier_sqlx_v2 "github.com/go-courier/sqlx/v2"
	github_com_go_courier_sqlx_v2_builder "github.com/go-courier/sqlx/v2/builder"
	github_com_go_courier_sqlx_v2_datatypes "github.com/go-courier/sqlx/v2/datatypes"
)

func (Account) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (Account) Indexes() github_com_go_courier_sqlx_v2_builder.Indexes {
	return github_com_go_courier_sqlx_v2_builder.Indexes{
		"I_created_at": []string{
			"CreatedAt",
		},
		"I_nick_name": []string{
			"NickName",
		},
	}
}

func (Account) UniqueIndexIAccountID() string {
	return "I_account_id"
}

func (Account) UniqueIndexes() github_com_go_courier_sqlx_v2_builder.Indexes {
	return github_com_go_courier_sqlx_v2_builder.Indexes{
		"I_account_id": []string{
			"AccountID",
		},
	}
}

var AccountTable *github_com_go_courier_sqlx_v2_builder.Table

func init() {
	AccountTable = DB.Register(&Account{})
}

type AccountIterator struct {
}

func (AccountIterator) New() interface{} {
	return &Account{}
}

func (AccountIterator) Resolve(v interface{}) *Account {
	return v.(*Account)
}

func (Account) TableName() string {
	return "t_account"
}

func (Account) TableDescription() []string {
	return []string{
		"账户 --表注释",
	}
}

func (Account) ColDescriptions() map[string][]string {
	return map[string][]string{
		"AccountID": []string{
			"账户唯一 ID",
		},
		"AccountType": []string{
			"账户类型",
		},
		"Avatar": []string{
			"头像",
		},
		"CreatedAt": []string{
			"创建时间",
		},
		"ID": []string{
			"自增 ID",
		},
		"NickName": []string{
			"昵称",
		},
		"State": []string{
			"状态 --字段注释",
		},
		"UpdatedAt": []string{
			"更新时间",
		},
	}
}

func (Account) FieldKeyID() string {
	return "ID"
}

func (m *Account) FieldID() *github_com_go_courier_sqlx_v2_builder.Column {
	return AccountTable.F(m.FieldKeyID())
}

func (Account) FieldKeyAccountID() string {
	return "AccountID"
}

func (m *Account) FieldAccountID() *github_com_go_courier_sqlx_v2_builder.Column {
	return AccountTable.F(m.FieldKeyAccountID())
}

func (Account) FieldKeyNickName() string {
	return "NickName"
}

func (m *Account) FieldNickName() *github_com_go_courier_sqlx_v2_builder.Column {
	return AccountTable.F(m.FieldKeyNickName())
}

func (Account) FieldKeyAccountType() string {
	return "AccountType"
}

func (m *Account) FieldAccountType() *github_com_go_courier_sqlx_v2_builder.Column {
	return AccountTable.F(m.FieldKeyAccountType())
}

func (Account) FieldKeyAvatar() string {
	return "Avatar"
}

func (m *Account) FieldAvatar() *github_com_go_courier_sqlx_v2_builder.Column {
	return AccountTable.F(m.FieldKeyAvatar())
}

func (Account) FieldKeyState() string {
	return "State"
}

func (m *Account) FieldState() *github_com_go_courier_sqlx_v2_builder.Column {
	return AccountTable.F(m.FieldKeyState())
}

func (Account) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *Account) FieldCreatedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return AccountTable.F(m.FieldKeyCreatedAt())
}

func (Account) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *Account) FieldUpdatedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return AccountTable.F(m.FieldKeyUpdatedAt())
}

func (Account) ColRelations() map[string][]string {
	return map[string][]string{
		"AccountID": []string{
			"Account",
			"AccountID",
		},
	}
}

func (m *Account) IndexFieldNames() []string {
	return []string{
		"AccountID",
		"CreatedAt",
		"ID",
		"NickName",
	}
}

func (m *Account) ConditionByStruct(db github_com_go_courier_sqlx_v2.DBExecutor) github_com_go_courier_sqlx_v2_builder.SqlCondition {
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

	return condition
}

func (m *Account) Create(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_go_courier_sqlx_v2.InsertToDB(db, m, nil))
	return err

}

func (m *Account) CreateOnDuplicateWithUpdateFields(db github_com_go_courier_sqlx_v2.DBExecutor, updateFields []string) error {

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

func (m *Account) DeleteByStruct(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(m.ConditionByStruct(db)),
				github_com_go_courier_sqlx_v2_builder.Comment("Account.DeleteByStruct"),
			),
	)

	return err
}

func (m *Account) FetchByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Account.FetchByID"),
			),
		m,
	)

	return err
}

func (m *Account) UpdateByIDWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("Account.UpdateByIDWithMap"),
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

func (m *Account) UpdateByIDWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *Account) FetchByIDForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
				)),
				github_com_go_courier_sqlx_v2_builder.ForUpdate(),
				github_com_go_courier_sqlx_v2_builder.Comment("Account.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *Account) DeleteByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Account.DeleteByID"),
			))

	return err
}

func (m *Account) FetchByAccountID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("AccountID").Eq(m.AccountID),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Account.FetchByAccountID"),
			),
		m,
	)

	return err
}

func (m *Account) UpdateByAccountIDWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("AccountID").Eq(m.AccountID),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("Account.UpdateByAccountIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByAccountID(db)
	}

	return nil

}

func (m *Account) UpdateByAccountIDWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByAccountIDWithMap(db, fieldValues)

}

func (m *Account) FetchByAccountIDForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("AccountID").Eq(m.AccountID),
				)),
				github_com_go_courier_sqlx_v2_builder.ForUpdate(),
				github_com_go_courier_sqlx_v2_builder.Comment("Account.FetchByAccountIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *Account) DeleteByAccountID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("AccountID").Eq(m.AccountID),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Account.DeleteByAccountID"),
			))

	return err
}

func (m *Account) List(db github_com_go_courier_sqlx_v2.DBExecutor, condition github_com_go_courier_sqlx_v2_builder.SqlCondition, additions ...github_com_go_courier_sqlx_v2_builder.Addition) ([]Account, error) {

	list := make([]Account, 0)

	table := db.T(m)
	_ = table

	finalAdditions := []github_com_go_courier_sqlx_v2_builder.Addition{
		github_com_go_courier_sqlx_v2_builder.Where(condition),
		github_com_go_courier_sqlx_v2_builder.Comment("Account.List"),
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

func (m *Account) Count(db github_com_go_courier_sqlx_v2.DBExecutor, condition github_com_go_courier_sqlx_v2_builder.SqlCondition, additions ...github_com_go_courier_sqlx_v2_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	finalAdditions := []github_com_go_courier_sqlx_v2_builder.Addition{
		github_com_go_courier_sqlx_v2_builder.Where(condition),
		github_com_go_courier_sqlx_v2_builder.Comment("Account.Count"),
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

func (m *Account) BatchFetchByAccountIDList(db github_com_go_courier_sqlx_v2.DBExecutor, values []git_querycap_com_tools_datatypes.SFID) ([]Account, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("AccountID").In(values)

	return m.List(db, condition)

}

func (m *Account) BatchFetchByCreatedAtList(db github_com_go_courier_sqlx_v2.DBExecutor, values []github_com_go_courier_sqlx_v2_datatypes.Timestamp) ([]Account, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("CreatedAt").In(values)

	return m.List(db, condition)

}

func (m *Account) BatchFetchByIDList(db github_com_go_courier_sqlx_v2.DBExecutor, values []uint64) ([]Account, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}

func (m *Account) BatchFetchByNickNameList(db github_com_go_courier_sqlx_v2.DBExecutor, values []string) ([]Account, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("NickName").In(values)

	return m.List(db, condition)

}
