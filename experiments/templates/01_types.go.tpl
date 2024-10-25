{{if .Table.IsJoinTable -}}
{{else -}}
{{- $alias := .Aliases.Table .Table.Name -}}

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	{{if .Table.IsView -}}
	// These are used in some views
	_ = fmt.Sprintln("")
	_ = reflect.Int
	_ = strings.Builder{}
	_ = sync.Mutex{}
	_ = strmangle.Plural("")
	_ = strconv.IntSize
	{{- end}}
)
{{end -}}
