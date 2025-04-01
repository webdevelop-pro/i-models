package models

import (
	"reflect"
)

type GenericModel struct {
	ID int `db:"id" json:"id"`
}

func (gm GenericModel) Table() string {
	return "generic_table"
}

func (gm GenericModel) PrimaryFieldKey() string {
	return "id"
}

func (gm GenericModel) PrimaryFieldValue() any {
	return gm.ID
}

func (gm GenericModel) Fields() []string {
	return DefaultFields(&gm)
}

func (gm GenericModel) ToJSON() map[string]any {
	return map[string]any{
		"ID": gm.ID,
	}
}

func (gm GenericModel) GetID() any {
	return gm.ID
}

func (gm *GenericModel) SetID(id any) {
	gm.ID = id.(int)
}

func DefaultFields(obj any) []string {
	var res = []string{}
	val := reflect.ValueOf(obj).Elem()
	for i := 0; i < val.NumField(); i++ {
		dbtag := string(val.Type().Field(i).Tag.Get("db"))
		if dbtag == "" {
			dbtag = string(val.Type().Field(i).Tag.Get("json"))
		}
		if dbtag != "" && dbtag != "-" {
			res = append(res, dbtag)
		}
	}
	return res
}
