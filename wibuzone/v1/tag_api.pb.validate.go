// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: wibuzone/v1/tag_api.proto

package wibuzonev1

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

// Validate checks the field values on ListTagsFilter with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListTagsFilter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListTagsFilter with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListTagsFilterMultiError,
// or nil if none found.
func (m *ListTagsFilter) ValidateAll() error {
	return m.validate(true)
}

func (m *ListTagsFilter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Term

	if len(errors) > 0 {
		return ListTagsFilterMultiError(errors)
	}

	return nil
}

// ListTagsFilterMultiError is an error wrapping multiple validation errors
// returned by ListTagsFilter.ValidateAll() if the designated constraints
// aren't met.
type ListTagsFilterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListTagsFilterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListTagsFilterMultiError) AllErrors() []error { return m }

// ListTagsFilterValidationError is the validation error returned by
// ListTagsFilter.Validate if the designated constraints aren't met.
type ListTagsFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTagsFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTagsFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTagsFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTagsFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTagsFilterValidationError) ErrorName() string { return "ListTagsFilterValidationError" }

// Error satisfies the builtin error interface
func (e ListTagsFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTagsFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTagsFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTagsFilterValidationError{}

// Validate checks the field values on ListTagsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListTagsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListTagsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListTagsRequestMultiError, or nil if none found.
func (m *ListTagsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListTagsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetFilter()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListTagsRequestValidationError{
					field:  "Filter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListTagsRequestValidationError{
					field:  "Filter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFilter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListTagsRequestValidationError{
				field:  "Filter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for PageToken

	// no validation rules for PageSize

	if len(errors) > 0 {
		return ListTagsRequestMultiError(errors)
	}

	return nil
}

// ListTagsRequestMultiError is an error wrapping multiple validation errors
// returned by ListTagsRequest.ValidateAll() if the designated constraints
// aren't met.
type ListTagsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListTagsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListTagsRequestMultiError) AllErrors() []error { return m }

// ListTagsRequestValidationError is the validation error returned by
// ListTagsRequest.Validate if the designated constraints aren't met.
type ListTagsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTagsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTagsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTagsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTagsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTagsRequestValidationError) ErrorName() string { return "ListTagsRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListTagsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTagsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTagsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTagsRequestValidationError{}

// Validate checks the field values on ListTagsResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListTagsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListTagsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListTagsResponseMultiError, or nil if none found.
func (m *ListTagsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListTagsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetTags() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListTagsResponseValidationError{
						field:  fmt.Sprintf("Tags[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListTagsResponseValidationError{
						field:  fmt.Sprintf("Tags[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListTagsResponseValidationError{
					field:  fmt.Sprintf("Tags[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for NextPageToken

	if len(errors) > 0 {
		return ListTagsResponseMultiError(errors)
	}

	return nil
}

// ListTagsResponseMultiError is an error wrapping multiple validation errors
// returned by ListTagsResponse.ValidateAll() if the designated constraints
// aren't met.
type ListTagsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListTagsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListTagsResponseMultiError) AllErrors() []error { return m }

// ListTagsResponseValidationError is the validation error returned by
// ListTagsResponse.Validate if the designated constraints aren't met.
type ListTagsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTagsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTagsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTagsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTagsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTagsResponseValidationError) ErrorName() string { return "ListTagsResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListTagsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTagsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTagsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTagsResponseValidationError{}

// Validate checks the field values on GetTagRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetTagRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTagRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetTagRequestMultiError, or
// nil if none found.
func (m *GetTagRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTagRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetTagRequestMultiError(errors)
	}

	return nil
}

// GetTagRequestMultiError is an error wrapping multiple validation errors
// returned by GetTagRequest.ValidateAll() if the designated constraints
// aren't met.
type GetTagRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTagRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTagRequestMultiError) AllErrors() []error { return m }

// GetTagRequestValidationError is the validation error returned by
// GetTagRequest.Validate if the designated constraints aren't met.
type GetTagRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTagRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTagRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTagRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTagRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTagRequestValidationError) ErrorName() string { return "GetTagRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetTagRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTagRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTagRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTagRequestValidationError{}

// Validate checks the field values on GetTagResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetTagResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTagResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetTagResponseMultiError,
// or nil if none found.
func (m *GetTagResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTagResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTag()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetTagResponseValidationError{
					field:  "Tag",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetTagResponseValidationError{
					field:  "Tag",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTag()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetTagResponseValidationError{
				field:  "Tag",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetTagResponseMultiError(errors)
	}

	return nil
}

// GetTagResponseMultiError is an error wrapping multiple validation errors
// returned by GetTagResponse.ValidateAll() if the designated constraints
// aren't met.
type GetTagResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTagResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTagResponseMultiError) AllErrors() []error { return m }

// GetTagResponseValidationError is the validation error returned by
// GetTagResponse.Validate if the designated constraints aren't met.
type GetTagResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTagResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTagResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTagResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTagResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTagResponseValidationError) ErrorName() string { return "GetTagResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetTagResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTagResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTagResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTagResponseValidationError{}