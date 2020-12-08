// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type actionResultTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *actionResultTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("action_results").
func (v *actionResultTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *actionResultTableType) Columns() []string {
	return []string{
		"id",
		"pmm_agent_id",
		"done",
		"error",
		"output",
		"created_at",
		"updated_at",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *actionResultTableType) NewStruct() reform.Struct {
	return new(ActionResult)
}

// NewRecord makes a new record for that table.
func (v *actionResultTableType) NewRecord() reform.Record {
	return new(ActionResult)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *actionResultTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// ActionResultTable represents action_results view or table in SQL database.
var ActionResultTable = &actionResultTableType{
	s: parse.StructInfo{
		Type:    "ActionResult",
		SQLName: "action_results",
		Fields: []parse.FieldInfo{
			{Name: "ID", Type: "string", Column: "id"},
			{Name: "PMMAgentID", Type: "string", Column: "pmm_agent_id"},
			{Name: "Done", Type: "bool", Column: "done"},
			{Name: "Error", Type: "string", Column: "error"},
			{Name: "Output", Type: "string", Column: "output"},
			{Name: "CreatedAt", Type: "time.Time", Column: "created_at"},
			{Name: "UpdatedAt", Type: "time.Time", Column: "updated_at"},
		},
		PKFieldIndex: 0,
	},
	z: new(ActionResult).Values(),
}

// String returns a string representation of this struct or record.
func (s ActionResult) String() string {
	res := make([]string, 7)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "PMMAgentID: " + reform.Inspect(s.PMMAgentID, true)
	res[2] = "Done: " + reform.Inspect(s.Done, true)
	res[3] = "Error: " + reform.Inspect(s.Error, true)
	res[4] = "Output: " + reform.Inspect(s.Output, true)
	res[5] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[6] = "UpdatedAt: " + reform.Inspect(s.UpdatedAt, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *ActionResult) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.PMMAgentID,
		s.Done,
		s.Error,
		s.Output,
		s.CreatedAt,
		s.UpdatedAt,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *ActionResult) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.PMMAgentID,
		&s.Done,
		&s.Error,
		&s.Output,
		&s.CreatedAt,
		&s.UpdatedAt,
	}
}

// View returns View object for that struct.
func (s *ActionResult) View() reform.View {
	return ActionResultTable
}

// Table returns Table object for that record.
func (s *ActionResult) Table() reform.Table {
	return ActionResultTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *ActionResult) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *ActionResult) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *ActionResult) HasPK() bool {
	return s.ID != ActionResultTable.z[ActionResultTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.ID = pk.
func (s *ActionResult) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = ActionResultTable
	_ reform.Struct = (*ActionResult)(nil)
	_ reform.Table  = ActionResultTable
	_ reform.Record = (*ActionResult)(nil)
	_ fmt.Stringer  = (*ActionResult)(nil)
)

func init() {
	parse.AssertUpToDate(&ActionResultTable.s, new(ActionResult))
}
