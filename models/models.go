package models

import (
	"reflect"
)

type GeneralModel struct {
	table string
}

func New(table string) *GeneralModel {
	return &GeneralModel{
		table: table,
	}
}

func (gm GeneralModel) Fields() []string {
	obj := &gm
	var res = []string{}
	val := reflect.ValueOf(obj).Elem()
	for i := 0; i < val.NumField(); i++ {
		dbtag := string(val.Type().Field(i).Tag.Get("db"))
		if dbtag != "" {
			res = append(res, dbtag)
		}
	}
	return res
}

func (gm GeneralModel) Table() string {
	return gm.table
}

func (pfc GeneralModel) GetID() any {
	return ""
}

func (pfc GeneralModel) ToJSON() map[string]any {
	return map[string]any{}
}
