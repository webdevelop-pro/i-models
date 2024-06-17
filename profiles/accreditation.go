package profiles

type Accreditation struct {
	Status AccType `json:"status"`
	//nolint:godox
	// ToDo
	// if you want to use pgtype.Timestamptz you will get an error
	// can't scan into dest[7]: parsing time \"2023-09-06\" as
	// \"2006-01-02T15:04:05.999999999Z07:00\": cannot parse \"\" as \"T\""
	// question, where f*ck T come from?
	CreatedAt   string            `json:"created_at"`
	CompletedAt string            `json:"completed_at,omitempty"`
	Data        map[string]string `json:"data"`
}
