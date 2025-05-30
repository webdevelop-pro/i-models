package logs

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/webdevelop-pro/go-common/context/keys"
	"github.com/webdevelop-pro/go-common/db"
	"github.com/webdevelop-pro/go-common/logger"
	"github.com/webdevelop-pro/i-models/pgtype"
)

type ContextKey rune

// LogLog is an object representing the database table.
type LogLog struct {
	ID               int                `json:"id" yaml:"id"`
	ContentTypeID    int                `json:"content_type_id" yaml:"content_type_id"`
	MSGID            *string            `json:"msg_id,omitempty" yaml:"msg_id,omitempty"`
	ObjectID         string             `json:"object_id" yaml:"object_id"`
	StatusCode       int                `json:"status_code" yaml:"status_code"`
	Path             string             `json:"path" yaml:"path"`
	RequestHeaders   any                `json:"request_headers" yaml:"request_headers"`
	RequestData      string             `json:"request_data" yaml:"request_data"`
	ResponseHeaders  any                `json:"response_headers" yaml:"response_headers"`
	ResponseData     string             `json:"response_data" yaml:"response_data"`
	Service          ServicesT          `json:"service" yaml:"service"`
	Type             LogTypeT           `json:"type" yaml:"type"`
	RequestID        *string            `json:"request_id,omitempty" yaml:"request_id,omitempty"`
	MetaData         any                `json:"meta_data" yaml:"meta_data"`
	RequestCreatedAt pgtype.Timestamptz `json:"request_created_at" yaml:"request_created_at"`
	CreatedAt        pgtype.Timestamptz `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	UpdatedAt        pgtype.Timestamptz `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
}

// ToDo Use db.DB instead of pgxpool.Pool
func LogHttpRequest(
	ctx context.Context,
	db db.Repository,
	logType LogTypeT,
	objectID string,
	serviceName ServicesT,
	objectLabel string,
	objectModel ObjectType,
	req *http.Request,
	rawBody []byte,
) (*LogLog, error) {
	var (
		reqID, _ = ctx.Value(keys.RequestID).(string)
		msgID, _ = ctx.Value(keys.MSGID).(string)
		log      = logger.FromCtx(ctx, "logs")
		model    = LogLog{}
	)

	sql := fmt.Sprintf(`INSERT INTO %s(
				service, "type", content_type_id, object_id, path,
				request_created_at, request_id, request_headers, request_data, msg_id
			) VALUES (
				$1, $2, $3   , $4, $5, $6, $7, $8, $9, $10
			) RETURNING id`, model.Table(),
	)

	contentID, err := GetContentID(ctx, db, objectLabel, objectModel)
	if err != nil {
		log.Error().Err(err).Msg("can't get log from database")
		return nil, err
	}

	// no tabs or new lines
	rawBody = bytes.ReplaceAll(rawBody, []byte("\t"), []byte{})
	rawBody = bytes.ReplaceAll(rawBody, []byte("\n"), []byte{})

	header := req.Header
	if header == nil {
		header = map[string][]string{}
	}

	path := req.RequestURI
	if path == "" && req.URL != nil {
		path = req.URL.EscapedPath()
	}

	err = db.QueryRow(
		ctx,
		sql,
		serviceName,
		logType,
		contentID,
		objectID,
		path,
		time.Now(),
		reqID,
		header,
		rawBody,
		msgID,
	).Scan(&model.ID)
	if err != nil {
		log.Error().Err(err).Msg("can't save log in database")
		return &model, err
	}
	return &model, nil
}

// ToDo Use db.DB instead of pgxpool.Pool
func GetContentID(ctx context.Context, db db.Repository, serviceLabel string, serviceModel ObjectType) (int, error) {
	contentID := 1

	sql, args, err := sq.Select("id").
		From("django_content_type").Where(sq.And{sq.Eq{"app_label": serviceLabel}, sq.Eq{"model": serviceModel}}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return contentID, err
	}

	err = db.QueryRow(ctx, sql, args...).Scan(&contentID)
	if err != nil {
		return contentID, err
	}

	return contentID, nil
}

func intToUUID(i string) string {
	res := i
	for x := len(i); x < 12; x++ {
		res = fmt.Sprintf("0%s", res)
	}
	//00000000-0000-0000-0000-000000000001
	return fmt.Sprintf("00000000-0000-0000-0000-%s", res)
}

// ToDo Use db.DB instead of pgxpool.Pool
func (model *LogLog) LogRequest(log logger.Logger, db db.Repository, objectID string, serviceName ServicesT, appLabel string, objectModel ObjectType) func(req *http.Request) {
	return func(req *http.Request) {
		var (
			err error
		)

		rawBody := []byte("{}")
		if req.Body != nil {
			rawBody, _ = io.ReadAll(req.Body)
			req.Body = io.NopCloser(bytes.NewReader(rawBody))
		}

		if val, ok := req.Context().Value(keys.RequestID).(string); ok && val != "" {
			req.Header.Add("X-Request-Id", val)
		} else {
			req.Header.Add("X-Request-Id", intToUUID(objectID))
		}

		// ToDo: move to adapters/dwolla
		// https://developers.dwolla.com/docs/balance/api-reference/api-fundamentals/idempotency-key
		if serviceName == ServicesTDwolla {
			if val, ok := req.Context().Value(keys.RequestID).(string); ok && val != "" {
				req.Header.Add("Idempotency-Key", val)
			} else {
				req.Header.Add("Idempotency-Key", intToUUID(objectID))
			}
		}

		obj, err := LogHttpRequest(
			req.Context(),
			db,
			LogTypeTOutcoming,
			objectID,
			serviceName,
			appLabel,
			objectModel,
			// ToDo
			// get url from webhook header
			req,
			rawBody,
		)
		model.ID = obj.ID

		// TODO: Use the same format for incoming logs
		log.Trace().Str("path", req.RequestURI).Str("service", serviceName.String()).Msg(MsgSendRequest)

		if err != nil {
			log.Error().Err(err).Msg("can't save log in database")
		}
	}
}

func (model *LogLog) LogResponse(log logger.Logger, db db.Repository, serviceName ServicesT) func(req *http.Response) {
	return func(resp *http.Response) {
		model.UpdateLog(resp.Request.Context(), log, db, serviceName, resp)
	}
}

func (model *LogLog) UpdateLog(ctx context.Context, log logger.Logger, db db.Repository, serviceName ServicesT, resp *http.Response) {
	logID, _ := resp.Request.Context().Value(keys.RequestLogID).(int)

	sql := fmt.Sprintf(`UPDATE %s
			SET status_code=$2, response_headers=$3, response_data=$4
			WHERE id=$1`, model.Table(),
	)

	rawBody := []byte("{}")
	if resp.Body != nil {
		rawBody, _ = io.ReadAll(resp.Body)
		resp.Body = io.NopCloser(bytes.NewReader(rawBody))
	}

	if len(string(rawBody)) == 0 {
		rawBody = []byte("{}")
	}

	// TODO: Use the same format for incoming logs
	log.Trace().Str("path", resp.Request.RequestURI).
		Int("logID", model.ID).Str("service", serviceName.String()).Msg(MsgRequestProcessed)

	res, err := db.Exec(
		resp.Request.Context(),
		sql,
		model.ID,
		resp.StatusCode,
		resp.Header,
		rawBody,
	)
	if err != nil {
		log.Error().Err(err).Int("log_id", logID).Msg("can't update log in database")
	}
	if res.String() == "UPDATE 0" {
		log.Error().Err(err).Int("log_id", logID).Msg("can't update log in database")
	}
}

func (model LogLog) ToJSON() map[string]any {
	return map[string]any{
		"id":                 model.ID,
		"content_type_id":    model.ContentTypeID,
		"msg_id":             model.MSGID,
		"object_id":          model.ObjectID,
		"status_code":        model.StatusCode,
		"path":               model.Path,
		"request_headers":    model.RequestHeaders,
		"request_data":       model.RequestData,
		"response_headers":   model.ResponseHeaders,
		"response_data":      model.ResponseData,
		"service":            model.Service,
		"type":               model.Type,
		"request_id":         model.RequestID,
		"meta_data":          model.MetaData,
		"request_created_at": model.RequestCreatedAt,
		"created_at":         model.CreatedAt,
		"updated_at":         model.UpdatedAt,
	}
}

func (model LogLog) Fields() []string {
	return []string{
		"ID",
		"ContentTypeID",
		"MSGID",
		"ObjectID",
		"StatusCode",
		"Path",
		"RequestHeaders",
		"RequestData",
		"ResponseHeaders",
		"ResponseData",
		"Service",
		"Type",
		"RequestID",
		"MetaData",
		"RequestCreatedAt",
		"CreatedAt",
		"UpdatedAt",
	}
}

func (model LogLog) Table() string {
	return "log_logs"
}

func (model LogLog) GetID() any {
	return model.ID
}
