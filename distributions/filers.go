package distributions

// DistributionDistributionFiler is an object representing the database table.
type DistributionDistributionFiler struct {
	ID             int                    `db:"id" json:"id" yaml:"id"`
	DistributionID *int                   `db:"distribution_id" json:"distribution_id,omitempty" yaml:"distribution_id,omitempty"`
	FilerID        *int                   `db:"filer_id" json:"filer_id,omitempty" yaml:"filer_id,omitempty"`
	Type           NullDistributionFilerT `db:"type" json:"type,omitempty" yaml:"type,omitempty"`
}

func (model DistributionDistributionFiler) ToJSON() map[string]any {
	return map[string]any{
		"id":              model.ID,
		"distribution_id": model.DistributionID,
		"filer_id":        model.FilerID,
		"type":            model.Type,
	}
}

func (model DistributionDistributionFiler) Fields() []string {
	return []string{"id", "distribution_id", "filer_id", "type"}
}

func (model DistributionDistributionFiler) Table() string {
	return "distribution_distribution_filers"
}

func (model DistributionDistributionFiler) GetID() any {
	return model.ID
}

func (model *DistributionDistributionFiler) SetID(id any) {
	model.ID = id.(int)
}
