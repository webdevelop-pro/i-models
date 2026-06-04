package models_test

import (
	"reflect"
	"testing"

	"github.com/webdevelop-pro/i-models/distributions"
	"github.com/webdevelop-pro/i-models/filers"
	"github.com/webdevelop-pro/i-models/historylogs"
	logmodels "github.com/webdevelop-pro/i-models/logs"
	"github.com/webdevelop-pro/i-models/offers"
	"github.com/webdevelop-pro/i-models/pubsublogs"
)

func TestFieldsUseDatabaseColumns(t *testing.T) {
	tests := []struct {
		name string
		got  []string
		want []string
	}{
		{
			name: "filers",
			got:  filers.FilerFiler{}.Fields(),
			want: []string{
				"id", "user_id", "group_id", "filename", "url", "mime", "name",
				"description", "bucket_name", "bucket_path", "meta_data", "created_at", "updated_at",
			},
		},
		{
			name: "historylogs",
			got:  historylogs.HistoryLog{}.Fields(),
			want: []string{
				"id", "content_type_id", "object_id", "action_flag", "object_repr",
				"change_message", "user_id", "action_time",
			},
		},
		{
			name: "pubsublogs",
			got:  pubsublogs.PubsubLog{}.Fields(),
			want: []string{"id", "topic", "msg", "attr", "msg_id", "executed", "created_at", "updated_at"},
		},
		{
			name: "logs",
			got:  logmodels.LogLog{}.Fields(),
			want: []string{
				"id", "content_type_id", "msg_id", "object_id", "status_code", "path",
				"request_headers", "request_data", "response_headers", "response_data",
				"service", "type", "request_id", "meta_data", "request_created_at", "created_at", "updated_at",
			},
		},
		{
			name: "offer comments",
			got:  offers.OfferComment{}.Fields(),
			want: []string{"id", "user_id", "offer_id", "comment", "related", "status", "approved_at", "created_at", "updated_at"},
		},
		{
			name: "offer filers",
			got:  offers.OfferOfferFiler{}.Fields(),
			want: []string{"id", "offer_id", "filer_id", "type"},
		},
		{
			name: "distribution filers",
			got:  distributions.DistributionDistributionFiler{}.Fields(),
			want: []string{"id", "distribution_id", "filer_id", "type"},
		},
		{
			name: "distributions",
			got:  distributions.DistributionDistribution{}.Fields(),
			want: []string{
				"id", "user_id", "profile_id", "investment_id", "report_id", "amount",
				"status", "data", "submited_at", "created_at", "updated_at",
			},
		},
		{
			name: "distribution reports",
			got:  distributions.DistributionDistributionReport{}.Fields(),
			want: []string{
				"id", "user_id", "offer_id", "amount", "submited_at", "data",
				"created_at", "updated_at", "image_link_id",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Fatalf("unexpected fields:\n got: %#v\nwant: %#v", tt.got, tt.want)
			}
		})
	}
}
