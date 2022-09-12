// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: wibuzone/v1/profile_api.proto

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

// Validate checks the field values on SaveProfileRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SaveProfileRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SaveProfileRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SaveProfileRequestMultiError, or nil if none found.
func (m *SaveProfileRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SaveProfileRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for AvatarUrl

	// no validation rules for Description

	if len(errors) > 0 {
		return SaveProfileRequestMultiError(errors)
	}

	return nil
}

// SaveProfileRequestMultiError is an error wrapping multiple validation errors
// returned by SaveProfileRequest.ValidateAll() if the designated constraints
// aren't met.
type SaveProfileRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SaveProfileRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SaveProfileRequestMultiError) AllErrors() []error { return m }

// SaveProfileRequestValidationError is the validation error returned by
// SaveProfileRequest.Validate if the designated constraints aren't met.
type SaveProfileRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SaveProfileRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SaveProfileRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SaveProfileRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SaveProfileRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SaveProfileRequestValidationError) ErrorName() string {
	return "SaveProfileRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SaveProfileRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveProfileRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SaveProfileRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SaveProfileRequestValidationError{}

// Validate checks the field values on SaveProfileResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SaveProfileResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SaveProfileResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SaveProfileResponseMultiError, or nil if none found.
func (m *SaveProfileResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SaveProfileResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetProfile()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SaveProfileResponseValidationError{
					field:  "Profile",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SaveProfileResponseValidationError{
					field:  "Profile",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProfile()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SaveProfileResponseValidationError{
				field:  "Profile",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SaveProfileResponseMultiError(errors)
	}

	return nil
}

// SaveProfileResponseMultiError is an error wrapping multiple validation
// errors returned by SaveProfileResponse.ValidateAll() if the designated
// constraints aren't met.
type SaveProfileResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SaveProfileResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SaveProfileResponseMultiError) AllErrors() []error { return m }

// SaveProfileResponseValidationError is the validation error returned by
// SaveProfileResponse.Validate if the designated constraints aren't met.
type SaveProfileResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SaveProfileResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SaveProfileResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SaveProfileResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SaveProfileResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SaveProfileResponseValidationError) ErrorName() string {
	return "SaveProfileResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SaveProfileResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveProfileResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SaveProfileResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SaveProfileResponseValidationError{}

// Validate checks the field values on GetProfileRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetProfileRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProfileRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProfileRequestMultiError, or nil if none found.
func (m *GetProfileRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProfileRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetProfileRequestMultiError(errors)
	}

	return nil
}

// GetProfileRequestMultiError is an error wrapping multiple validation errors
// returned by GetProfileRequest.ValidateAll() if the designated constraints
// aren't met.
type GetProfileRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProfileRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProfileRequestMultiError) AllErrors() []error { return m }

// GetProfileRequestValidationError is the validation error returned by
// GetProfileRequest.Validate if the designated constraints aren't met.
type GetProfileRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProfileRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProfileRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProfileRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProfileRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProfileRequestValidationError) ErrorName() string {
	return "GetProfileRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProfileRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProfileRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProfileRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProfileRequestValidationError{}

// Validate checks the field values on GetProfileResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetProfileResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProfileResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProfileResponseMultiError, or nil if none found.
func (m *GetProfileResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProfileResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetProfile()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetProfileResponseValidationError{
					field:  "Profile",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetProfileResponseValidationError{
					field:  "Profile",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProfile()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetProfileResponseValidationError{
				field:  "Profile",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetProfileResponseMultiError(errors)
	}

	return nil
}

// GetProfileResponseMultiError is an error wrapping multiple validation errors
// returned by GetProfileResponse.ValidateAll() if the designated constraints
// aren't met.
type GetProfileResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProfileResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProfileResponseMultiError) AllErrors() []error { return m }

// GetProfileResponseValidationError is the validation error returned by
// GetProfileResponse.Validate if the designated constraints aren't met.
type GetProfileResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProfileResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProfileResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProfileResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProfileResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProfileResponseValidationError) ErrorName() string {
	return "GetProfileResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetProfileResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProfileResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProfileResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProfileResponseValidationError{}

// Validate checks the field values on GetProfileStatsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetProfileStatsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProfileStatsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProfileStatsRequestMultiError, or nil if none found.
func (m *GetProfileStatsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProfileStatsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetProfileStatsRequestMultiError(errors)
	}

	return nil
}

// GetProfileStatsRequestMultiError is an error wrapping multiple validation
// errors returned by GetProfileStatsRequest.ValidateAll() if the designated
// constraints aren't met.
type GetProfileStatsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProfileStatsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProfileStatsRequestMultiError) AllErrors() []error { return m }

// GetProfileStatsRequestValidationError is the validation error returned by
// GetProfileStatsRequest.Validate if the designated constraints aren't met.
type GetProfileStatsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProfileStatsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProfileStatsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProfileStatsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProfileStatsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProfileStatsRequestValidationError) ErrorName() string {
	return "GetProfileStatsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProfileStatsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProfileStatsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProfileStatsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProfileStatsRequestValidationError{}

// Validate checks the field values on GetProfileStatsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetProfileStatsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProfileStatsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProfileStatsResponseMultiError, or nil if none found.
func (m *GetProfileStatsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProfileStatsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStats()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetProfileStatsResponseValidationError{
					field:  "Stats",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetProfileStatsResponseValidationError{
					field:  "Stats",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStats()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetProfileStatsResponseValidationError{
				field:  "Stats",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetProfileStatsResponseMultiError(errors)
	}

	return nil
}

// GetProfileStatsResponseMultiError is an error wrapping multiple validation
// errors returned by GetProfileStatsResponse.ValidateAll() if the designated
// constraints aren't met.
type GetProfileStatsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProfileStatsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProfileStatsResponseMultiError) AllErrors() []error { return m }

// GetProfileStatsResponseValidationError is the validation error returned by
// GetProfileStatsResponse.Validate if the designated constraints aren't met.
type GetProfileStatsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProfileStatsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProfileStatsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProfileStatsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProfileStatsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProfileStatsResponseValidationError) ErrorName() string {
	return "GetProfileStatsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetProfileStatsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProfileStatsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProfileStatsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProfileStatsResponseValidationError{}
