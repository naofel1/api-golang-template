package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/naofel1/api-golang-template/pkg/apistatus"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

// BindData is helper function, returns false if data is not bound
func BindData(c *gin.Context, log *otelzap.Logger, req interface{}) bool {
	// Bind incoming json to struct and check for validation errors
	if err := c.ShouldBindJSON(req); err != nil {
		checkArguments(c, log, err)
		return false
	}

	return true
}

// BindURI is helper function, returns false if data is not bound
func BindURI(c *gin.Context, log *otelzap.Logger, req interface{}) bool {
	// Bind incoming json to struct and check for validation errors
	if err := c.ShouldBindUri(req); err != nil {
		checkArguments(c, log, err)
		return false
	}

	return true
}

// BindQuery is helper function, returns false if data is not bound
func BindQuery(c *gin.Context, log *otelzap.Logger, req interface{}) bool {
	// Bind incoming json to struct and check for validation errors
	if err := c.ShouldBindQuery(req); err != nil {
		checkArguments(c, log, err)
		return false
	}

	return true
}

// BindForm is helper function, returns false if data is not bound
func BindForm(c *gin.Context, log *otelzap.Logger, req interface{}) bool {
	if c.ContentType() != "multipart/form-data" {
		msg := fmt.Sprintf("%s only accepts Content-Type multipart/form-data", c.FullPath())

		err := apistatus.NewUnsupportedMediaType(msg)

		c.JSON(err.Status(), apistatus.NewErrorAPI(err))

		return false
	}
	// Bind incoming json to struct and check for validation errors
	if err := c.ShouldBind(req); err != nil {
		checkArguments(c, log, err)
		return false
	}

	return true
}

// BindHeader is helper function, returns false if data is not bound
func BindHeader(c *gin.Context, log *otelzap.Logger, req interface{}) bool {
	// bind Authorization Header to h and check for validation errors
	if err := c.ShouldBindHeader(req); err != nil {
		var valErrs validator.ValidationErrors
		if errors.As(err, &valErrs) {
			checkArguments(c, log, err)

			return false
		}

		log.Ctx(c.Request.Context()).Warn("Error binding data",
			zap.Error(err),
		)

		// otherwise error type is unknown
		err := apistatus.NewInternal()
		c.AbortWithStatusJSON(err.Status(), apistatus.NewErrorAPI(err))

		return false
	}

	return true
}

func checkArguments(c *gin.Context, log *otelzap.Logger, err error) {
	log.Ctx(c.Request.Context()).Info("Error binding data:",
		zap.Error(err),
	)
	// we used this type in bind_data to extract desired fields from errs
	// you might consider extracting it
	var collectedErrors []apistatus.InvalidArgument

	// Other, none data validation errors
	var (
		syntaxError        *json.SyntaxError
		unmarshalTypeError *json.UnmarshalTypeError
		validationErrors   validator.ValidationErrors
	)

	switch {
	// Syntax errors in json
	case errors.As(err, &syntaxError):
		log.Ctx(c.Request.Context()).Info("Error binding data",
			zap.Error(err),
		)

		collectedErrors = append(
			collectedErrors,
			apistatus.InvalidArgument{
				Field:   "general",
				Message: fmt.Sprintf("Request body contains invalid formed Json at position %d", syntaxError.Offset),
			})

	// In some circumstances Decode() may also return an
	// io.ErrUnexpectedEOF error for syntax errors in the JSON.
	case errors.Is(err, io.ErrUnexpectedEOF):
		log.Ctx(c.Request.Context()).Info("Error binding data",
			zap.Error(err),
		)

		collectedErrors = append(
			collectedErrors,
			apistatus.InvalidArgument{
				Field:   "general",
				Message: "Request body contains invalid formed Json",
			})

	// The case when trying to assign not valid type into struct.
	case errors.As(err, &unmarshalTypeError):
		log.Ctx(c.Request.Context()).Info("Error binding data",
			zap.Error(err),
		)

		collectedErrors = append(
			collectedErrors,
			apistatus.InvalidArgument{
				Field:   unmarshalTypeError.Field,
				Message: fmt.Sprintf("Invalid type specified for field %s at position %d", unmarshalTypeError.Field, unmarshalTypeError.Offset),
			})

		// The case when detected extra unexpected field in the request body.
	case strings.HasPrefix(err.Error(), "json: unknown field "):
		fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")

		collectedErrors = append(
			collectedErrors,
			apistatus.InvalidArgument{
				Field:   fieldName,
				Message: fmt.Sprintf("Request body contains unknown field %s", fieldName),
			})

	// An io.EOF error is returned by Decode() if the request body is empty.
	case errors.Is(err, io.EOF):
		log.Ctx(c.Request.Context()).Info("Empty body data", zap.Error(err))

		collectedErrors = append(
			collectedErrors,
			apistatus.InvalidArgument{
				Field:   "general",
				Message: "Request body must not be empty",
			})

	// The case the request body is too large.
	case err.Error() == "http: request body too large":
		collectedErrors = append(
			collectedErrors,
			apistatus.InvalidArgument{
				Field:   "general",
				Message: "Request body must not be larger than 1MB",
			})

	case errors.As(err, &validationErrors):
		for _, f := range validationErrors {
			collectedErrors = append(collectedErrors, apistatus.InvalidArgument{
				Field:   f.Field(),
				Value:   f.Value(),
				Tag:     f.Tag(),
				Param:   f.Param(),
				Message: checkField(f),
			},
			)
		}

	// Server Error response.
	default:
		log.Ctx(c.Request.Context()).Info("Error binding data",
			zap.Error(err),
		)
		// if we aren't able to properly extract validation errors,
		// we'll fall back and return an internal server error
		fallBack := apistatus.NewInternal()

		c.JSON(fallBack.Status(), gin.H{"error": fallBack})

		return
	}

	errAPI := apistatus.NewBadRequest("Invalid request parameters. See invalidArgs")

	c.JSON(apistatus.Status(errAPI), apistatus.NewInvalidArgsAPI(collectedErrors, errAPI))
}

func checkField(f validator.FieldError) string {
	msg := "Unknown error"

	switch f.ActualTag() {
	case "required":
		msg = fmt.Sprintf("The field %s Required", f.Field())
	case "email":
		msg = "Should be a valid email address"
	case "lte":
		msg = fmt.Sprintf("Should be less than %s" + f.Param())
	case "gte":
		msg = fmt.Sprintf("Should be greater than %s", f.Param())
	case "alpha":
		msg = "Should be alpha characters only"
	case "numeric":
		msg = "Should be numbers only"
	case "oneof":
		msg = fmt.Sprintf("Should contain one of values %s", f.Param())
	case "url":
		msg = "Should be valid web address starting with http(s)://..."
	case "min":
		msg = fmt.Sprintf("Should be minimum %s characters long", f.Param())
	case "max":
		msg = fmt.Sprintf("Should be maximum %s characters long", f.Param())
	case "datetime":
		msg = fmt.Sprintf("Should be valid Date/Time with %s format", f.Param())
	}

	return msg
}
