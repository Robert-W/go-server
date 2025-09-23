package requestutil

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

type RequestUtils struct {
	SchemaDecoder *schema.Decoder
	Validate      *validator.Validate
}

type Option = func(*RequestUtils)

func New(options ...Option) *RequestUtils {
	utils := &RequestUtils{}
	for _, opt := range options {
		opt(utils)
	}

	return utils
}

func WithSchemaDecoder() Option {
	return func(utils *RequestUtils) {
		schemaDecoder := schema.NewDecoder()
		schemaDecoder.IgnoreUnknownKeys(true)
		utils.SchemaDecoder = schemaDecoder
	}
}

func WithValidator() Option {
	return func(utils *RequestUtils) {
		utils.Validate = validator.New(validator.WithRequiredStructEnabled())
	}
}
