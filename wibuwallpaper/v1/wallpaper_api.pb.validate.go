// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: wibuwallpaper/v1/wallpaper_api.proto

package wibuwallpaperv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on ListWallpapersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListWallpapersRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListWallpapersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListWallpapersRequestMultiError, or nil if none found.
func (m *ListWallpapersRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListWallpapersRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PageToken

	if len(errors) > 0 {
		return ListWallpapersRequestMultiError(errors)
	}

	return nil
}

// ListWallpapersRequestMultiError is an error wrapping multiple validation
// errors returned by ListWallpapersRequest.ValidateAll() if the designated
// constraints aren't met.
type ListWallpapersRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListWallpapersRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListWallpapersRequestMultiError) AllErrors() []error { return m }

// ListWallpapersRequestValidationError is the validation error returned by
// ListWallpapersRequest.Validate if the designated constraints aren't met.
type ListWallpapersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListWallpapersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListWallpapersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListWallpapersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListWallpapersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListWallpapersRequestValidationError) ErrorName() string {
	return "ListWallpapersRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListWallpapersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListWallpapersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListWallpapersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListWallpapersRequestValidationError{}

// Validate checks the field values on ListWallpapersResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListWallpapersResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListWallpapersResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListWallpapersResponseMultiError, or nil if none found.
func (m *ListWallpapersResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListWallpapersResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetWallpapers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListWallpapersResponseValidationError{
						field:  fmt.Sprintf("Wallpapers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListWallpapersResponseValidationError{
						field:  fmt.Sprintf("Wallpapers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListWallpapersResponseValidationError{
					field:  fmt.Sprintf("Wallpapers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for NextPageToken

	if len(errors) > 0 {
		return ListWallpapersResponseMultiError(errors)
	}

	return nil
}

// ListWallpapersResponseMultiError is an error wrapping multiple validation
// errors returned by ListWallpapersResponse.ValidateAll() if the designated
// constraints aren't met.
type ListWallpapersResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListWallpapersResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListWallpapersResponseMultiError) AllErrors() []error { return m }

// ListWallpapersResponseValidationError is the validation error returned by
// ListWallpapersResponse.Validate if the designated constraints aren't met.
type ListWallpapersResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListWallpapersResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListWallpapersResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListWallpapersResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListWallpapersResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListWallpapersResponseValidationError) ErrorName() string {
	return "ListWallpapersResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListWallpapersResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListWallpapersResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListWallpapersResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListWallpapersResponseValidationError{}

// Validate checks the field values on CreateWallpaperRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateWallpaperRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateWallpaperRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateWallpaperRequestMultiError, or nil if none found.
func (m *CreateWallpaperRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateWallpaperRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetWallpaper()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateWallpaperRequestValidationError{
					field:  "Wallpaper",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateWallpaperRequestValidationError{
					field:  "Wallpaper",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWallpaper()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateWallpaperRequestValidationError{
				field:  "Wallpaper",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateWallpaperRequestMultiError(errors)
	}

	return nil
}

// CreateWallpaperRequestMultiError is an error wrapping multiple validation
// errors returned by CreateWallpaperRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateWallpaperRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateWallpaperRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateWallpaperRequestMultiError) AllErrors() []error { return m }

// CreateWallpaperRequestValidationError is the validation error returned by
// CreateWallpaperRequest.Validate if the designated constraints aren't met.
type CreateWallpaperRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateWallpaperRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateWallpaperRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateWallpaperRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateWallpaperRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateWallpaperRequestValidationError) ErrorName() string {
	return "CreateWallpaperRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateWallpaperRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateWallpaperRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateWallpaperRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateWallpaperRequestValidationError{}

// Validate checks the field values on CreateWallpaperResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateWallpaperResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateWallpaperResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateWallpaperResponseMultiError, or nil if none found.
func (m *CreateWallpaperResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateWallpaperResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetWallpaper()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateWallpaperResponseValidationError{
					field:  "Wallpaper",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateWallpaperResponseValidationError{
					field:  "Wallpaper",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWallpaper()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateWallpaperResponseValidationError{
				field:  "Wallpaper",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateWallpaperResponseMultiError(errors)
	}

	return nil
}

// CreateWallpaperResponseMultiError is an error wrapping multiple validation
// errors returned by CreateWallpaperResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateWallpaperResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateWallpaperResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateWallpaperResponseMultiError) AllErrors() []error { return m }

// CreateWallpaperResponseValidationError is the validation error returned by
// CreateWallpaperResponse.Validate if the designated constraints aren't met.
type CreateWallpaperResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateWallpaperResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateWallpaperResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateWallpaperResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateWallpaperResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateWallpaperResponseValidationError) ErrorName() string {
	return "CreateWallpaperResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateWallpaperResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateWallpaperResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateWallpaperResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateWallpaperResponseValidationError{}
