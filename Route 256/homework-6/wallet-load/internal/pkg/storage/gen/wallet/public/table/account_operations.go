//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var AccountOperations = newAccountOperationsTable("public", "account_operations", "")

type accountOperationsTable struct {
	postgres.Table

	// Columns
	ID            postgres.ColumnString
	AccountID     postgres.ColumnString
	Amount        postgres.ColumnInteger
	OperationID   postgres.ColumnString
	OperationType postgres.ColumnInteger
	CreatedAt     postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type AccountOperationsTable struct {
	accountOperationsTable

	EXCLUDED accountOperationsTable
}

// AS creates new AccountOperationsTable with assigned alias
func (a AccountOperationsTable) AS(alias string) *AccountOperationsTable {
	return newAccountOperationsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AccountOperationsTable with assigned schema name
func (a AccountOperationsTable) FromSchema(schemaName string) *AccountOperationsTable {
	return newAccountOperationsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AccountOperationsTable with assigned table prefix
func (a AccountOperationsTable) WithPrefix(prefix string) *AccountOperationsTable {
	return newAccountOperationsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AccountOperationsTable with assigned table suffix
func (a AccountOperationsTable) WithSuffix(suffix string) *AccountOperationsTable {
	return newAccountOperationsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAccountOperationsTable(schemaName, tableName, alias string) *AccountOperationsTable {
	return &AccountOperationsTable{
		accountOperationsTable: newAccountOperationsTableImpl(schemaName, tableName, alias),
		EXCLUDED:               newAccountOperationsTableImpl("", "excluded", ""),
	}
}

func newAccountOperationsTableImpl(schemaName, tableName, alias string) accountOperationsTable {
	var (
		IDColumn            = postgres.StringColumn("id")
		AccountIDColumn     = postgres.StringColumn("account_id")
		AmountColumn        = postgres.IntegerColumn("amount")
		OperationIDColumn   = postgres.StringColumn("operation_id")
		OperationTypeColumn = postgres.IntegerColumn("operation_type")
		CreatedAtColumn     = postgres.TimestampzColumn("created_at")
		allColumns          = postgres.ColumnList{IDColumn, AccountIDColumn, AmountColumn, OperationIDColumn, OperationTypeColumn, CreatedAtColumn}
		mutableColumns      = postgres.ColumnList{AccountIDColumn, AmountColumn, OperationIDColumn, OperationTypeColumn, CreatedAtColumn}
	)

	return accountOperationsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:            IDColumn,
		AccountID:     AccountIDColumn,
		Amount:        AmountColumn,
		OperationID:   OperationIDColumn,
		OperationType: OperationTypeColumn,
		CreatedAt:     CreatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
