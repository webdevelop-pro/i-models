{{$once := onceNew}}
{{$onceNull := onceNew}}
{{- $alias := .Aliases.Table .Table.Name -}}
{{- $orig_tbl_name := .Table.Name -}}
{{- $table := .Table -}}

	{{- range $col := $table.Columns | filterColumnsByEnum -}}
		{{- $name := parseEnumName $col.DBType -}}
		{{- $vals := parseEnumVals $col.DBType -}}
		{{- $isNamed := ne (len $name) 0}}
		{{- $enumName := "" -}}
		{{- if not (and
			$isNamed
			(and
				($once.Has $name)
				($onceNull.Has $name)
			)
		) -}}
			{{- if gt (len $vals) 0}}
				{{- if $isNamed -}}
					{{ $enumName = titleCase $name}}
				{{- else -}}
					{{ $enumName = printf "%s%s" (titleCase $table.Name) (titleCase $col.Name)}}
				{{- end -}}
				{{/* First iteration for enum type $name (nullable or not) */}}
				{{- $enumFirstIter := and
					(not ($once.Has $name))
					(not ($onceNull.Has $name))
				-}}

				{{- if $enumFirstIter -}}
					{{$enumType := "string" }}
					{{$allvals := "\n"}}

					{{if $.AddEnumTypes}}
						{{- $enumType = $enumName -}}
						type {{$enumName}} string
					{{end}}

					// Enum values for {{$enumName}}
					const (
					{{range $val := $vals -}}
						{{- $enumValue := titleCase $val -}}
						{{$enumName}}{{$enumValue}} {{$enumType}} = {{printf "%q" $val}}
						{{$allvals = printf "%s%s%s,\n" $allvals $enumName $enumValue -}}
					{{end -}}
					)

					func All{{$enumName}}() []{{$enumType}} {
						return []{{$enumType}}{ {{$allvals}} }
					}
				{{- end -}}

				{{if $.AddEnumTypes}}
					{{ if $enumFirstIter }}
						func (e {{$enumName}}) IsValid() error {
							{{- /* $first is being used to add a comma to all enumValues, but the first one.*/ -}}
							{{- $first := true -}}
							{{- /* $enumValues will contain a comma separated string holding all enum consts */ -}}
							{{- $enumValues := "" -}}
							{{ range $val := $vals -}}
								{{- if $first -}}
									{{- $first = false -}}
								{{- else -}}
									{{- $enumValues = printf "%s%s" $enumValues ", " -}}
								{{- end -}}

								{{- $enumValue := titleCase $val -}}
								{{- $enumValues = printf "%s%s%s" $enumValues $enumName $enumValue -}}
							{{- end}}
							switch e {
							case {{$enumValues}}:
								return nil
							default:
								return errors.New("enum is not valid")
							}
						}

						func (e {{$enumName}}) String() string {
							return string(e)
						}
					{{- end -}}

					{{ if and
						$col.Nullable
						(not ($onceNull.Has $name))
					}}
						{{$enumType := ""}}
						{{- if $isNamed -}}
							{{- $enumType = (print $.EnumNullPrefix $enumName) }}
						{{- else -}}
							{{- $enumType = printf "%s%s" (titleCase $table.Name) (print $.EnumNullPrefix (titleCase $col.Name)) -}}
						{{- end -}}
						// {{$enumType}} is a nullable {{$enumName}} enum type. It supports SQL and JSON serialization.
						type {{$enumType}} struct {
							Val		{{$enumName}}
							Valid	bool
						}

						// {{$enumType}}From creates a new {{$enumName}} that will never be blank.
						func {{$enumType}}From(v {{$enumName}}) {{$enumType}} {
							return New{{$enumType}}(v, true)
						}

						// {{$enumType}}FromPtr creates a new {{$enumType}} that be null if s is nil.
						func {{$enumType}}FromPtr(v *{{$enumName}}) {{$enumType}} {
							if v == nil {
								return New{{$enumType}}("", false)
							}
							return New{{$enumType}}(*v, true)
						}

						// New{{$enumType}} creates a new {{$enumType}}
						func New{{$enumType}}(v {{$enumName}}, valid bool) {{$enumType}} {
							return {{$enumType}}{
								Val:	v,
								Valid:  valid,
							}
						}

						// UnmarshalJSON implements json.Unmarshaler.
						func (e *{{$enumType}}) UnmarshalJSON(data []byte) error {
							if err := json.Unmarshal(data, &e.Val); err != nil {
								return err
							}

							e.Valid = true
							return nil
						}

						// MarshalJSON implements json.Marshaler.
						func (e {{$enumType}}) MarshalJSON() ([]byte, error) {
							return json.Marshal(e.Val)
						}

						// MarshalText implements encoding.TextMarshaler.
						func (e {{$enumType}}) MarshalText() ([]byte, error) {
							if !e.Valid {
								return []byte{}, nil
							}
							return []byte(e.Val), nil
						}

						// UnmarshalText implements encoding.TextUnmarshaler.
						func (e *{{$enumType}}) UnmarshalText(text []byte) error {
							if text == nil || len(text) == 0 {
								e.Valid = false
								return nil
							}

							e.Val = {{$enumName}}(text)
							e.Valid = true
							return nil
						}

						// SetValid changes this {{$enumType}} value and also sets it to be non-null.
						func (e *{{$enumType}}) SetValid(v {{$enumName}}) {
							e.Val = v
							e.Valid = true
						}

						// Ptr returns a pointer to this {{$enumType}} value, or a nil pointer if this {{$enumType}} is null.
						func (e {{$enumType}}) Ptr() *{{$enumName}} {
							if !e.Valid {
								return nil
							}
							return &e.Val
						}

						// IsZero returns true for null types.
						func (e {{$enumType}}) IsZero() bool {
							return !e.Valid
						}

						// Value implements the driver Valuer interface.
						func (e {{$enumType}}) Value() (any, error) {
							if !e.Valid {
								return nil, nil
							}
							return string(e.Val), nil
						}
					{{end -}}
				{{end -}}
			{{else}}
				// Enum values for {{$table.Name}} {{$col.Name}} are not proper Go identifiers, cannot emit constants
			{{- end -}}
			{{/* Save column type name after generation.
			 Needs to be at the bottom because we check for the first iteration
			 inside the $table.Columns loop. */}}
			{{- if $isNamed -}}
				{{- if $col.Nullable -}}
					{{$_ := $onceNull.Put $name}}
				{{- else -}}
					{{$_ := $once.Put $name}}
				{{- end -}}
			{{- end -}}
		{{- end -}}
	{{- end -}}


// {{$alias.UpSingular}} is an object representing the database table.
type {{$alias.UpSingular}} struct {
	{{- range $index, $column := .Table.Columns -}}
	{{- $colAlias := $alias.Column $column.Name -}}
	{{- $orig_col_name := $column.Name -}}
	{{- range $column.Comment | splitLines -}} 
	{{- if eq $index 0 -}}
	{{ "\n" }}
	{{- end -}}
	// {{ . }}
	{{ end -}}

	{{if ignore $orig_tbl_name $orig_col_name $.TagIgnore -}}
	{{$colAlias}} {{$column.Type}} `{{generateIgnoreTags $.Tags}}boil:"{{$column.Name}}" json:"-" toml:"-" yaml:"-"`
	{{ else -}}

	{{- /* render column alias and column type */ -}}
	{{ $colAlias }} {{ $column.Type -}}

	{{- /*
	  handle struct tags
	  StructTagCasing will be replaced with $.StructTagCases
	  however we need to keep this backward compatible
	  $.StructTagCasing will only be used when it's set to "alias"
    */ -}}
	`
	{{- if eq $.StructTagCasing "alias" -}}
	    {{- generateTags $.Tags $colAlias -}}
	    {{- generateTagWithCase "json" $column.Name $colAlias "alias" $column.Nullable -}}
	    {{- trim (generateTagWithCase "yaml" $column.Name $colAlias "alias" $column.Nullable) -}}
	{{- else -}}
	    {{- generateTags $.Tags $column.Name }}
	    {{- generateTagWithCase "json" $column.Name $colAlias $.StructTagCases.Json $column.Nullable -}}
	    {{- trim (generateTagWithCase "yaml" $column.Name $colAlias $.StructTagCases.Yaml $column.Nullable) -}}
	{{- end -}}
	`
	{{ end -}}
	{{ end -}}

}


func (model {{$alias.UpSingular}}) ToJSON() map[string]any {
	return map[string]any{
		{{range $column := .Table.Columns -}}
		{{- $colAlias := $alias.Column $column.Name -}}
		"{{$column.Name}}": model.{{$colAlias}},
		{{end -}}
	}
}

func (model {{$alias.UpSingular}}) Fields() []string {
	return models.DefaultFields(user)
}

func (model {{$alias.UpSingular}}) Table() string {
	return "{{$orig_tbl_name}}"
}

func (model {{$alias.UpSingular}}) GetID() any {
	return model.ID
}


func (model {{$alias.UpSingular}}) SetID() any {
	return model.ID
}

func (model {{$alias.UpSingular}}) ToJSON() map[string]any {
	return map[string]any{
		{{range $column := .Table.Columns -}}
		{{- $colAlias := $alias.Column $column.Name -}}
		"{{$column.Name}}": model.{{$colAlias}},
		{{end -}}
	}
}

func (model {{$alias.UpSingular}}) GetField(name string) any {
	switch name {
	{{range $column := .Table.Columns -}}
	case "{{- $alias.Column $column.Name -}}":
		return model.{{- $alias.Column $column.Name -}}
	{{end -}}
	}
	return nil
}
