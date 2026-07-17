package logs

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/global-torque/go-common/db/v2"
)

// RawEntry is a normalized third-party request/response log entry.
type RawEntry struct {
	Service          ServicesT
	Type             LogTypeT
	AppLabel         string
	ObjectModel      ObjectType
	ObjectID         string
	Path             string
	RequestCreatedAt time.Time
	RequestID        string
	RequestHeaders   any
	RequestData      any
	ResponseHeaders  any
	ResponseData     any
	StatusCode       int
}

// CreateRawEntry inserts a completed third-party request/response log entry.
func CreateRawEntry(ctx context.Context, repo db.Repository, entry RawEntry) (*LogLog, error) {
	if entry.RequestCreatedAt.IsZero() {
		entry.RequestCreatedAt = time.Now()
	}

	contentID, err := GetContentID(ctx, repo, entry.AppLabel, entry.ObjectModel)
	if err != nil {
		return nil, fmt.Errorf("get content type for %s.%s: %w", entry.AppLabel, entry.ObjectModel, err)
	}

	requestHeaders, err := json.Marshal(normalizeJSONValue(entry.RequestHeaders, map[string][]string{}))
	if err != nil {
		return nil, fmt.Errorf("marshal request headers: %w", err)
	}
	requestData, err := json.Marshal(normalizeJSONValue(entry.RequestData, map[string]any{}))
	if err != nil {
		return nil, fmt.Errorf("marshal request data: %w", err)
	}
	responseHeaders, err := json.Marshal(normalizeJSONValue(entry.ResponseHeaders, map[string][]string{}))
	if err != nil {
		return nil, fmt.Errorf("marshal response headers: %w", err)
	}
	responseData, err := json.Marshal(normalizeJSONValue(entry.ResponseData, map[string]any{}))
	if err != nil {
		return nil, fmt.Errorf("marshal response data: %w", err)
	}

	const sql = `INSERT INTO log_logs(
		service, "type", content_type_id, object_id, path,
		request_created_at, request_id, request_headers, request_data,
		response_headers, response_data, status_code
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
	) RETURNING id`

	model := &LogLog{}
	if err := repo.QueryRow(
		ctx,
		sql,
		entry.Service,
		entry.Type,
		contentID,
		entry.ObjectID,
		entry.Path,
		entry.RequestCreatedAt,
		entry.RequestID,
		requestHeaders,
		requestData,
		responseHeaders,
		responseData,
		entry.StatusCode,
	).Scan(&model.ID); err != nil {
		return model, fmt.Errorf("insert raw log entry: %w", err)
	}

	return model, nil
}

func normalizeJSONValue(value any, fallback any) any {
	if value == nil {
		return fallback
	}
	return value
}
