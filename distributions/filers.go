package distributions

// DistributionDistributionFiler is an object representing the database table.
type DistributionDistributionFiler struct {
	ID             int                    `json:"id" yaml:"id"`
	DistributionID int                    `json:"distribution_id,omitempty" yaml:"distribution_id,omitempty"`
	FilerID        int                    `json:"filer_id,omitempty" yaml:"filer_id,omitempty"`
	Type           NullDistributionFilerT `json:"type,omitempty" yaml:"type,omitempty"`
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
	return []string{
		"ID",
		"DistributionID",
		"FilerID",
		"Type",
	}
}

func (model DistributionDistributionFiler) Table() string {
	return "distribution_distribution_filers"
}

func (model DistributionDistributionFiler) GetID() any {
	return model.ID
}
