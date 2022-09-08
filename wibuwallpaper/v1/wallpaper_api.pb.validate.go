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

// Validate checks the field values on GetWallpaperRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetWallpaperRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetWallpaperRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetWallpaperRequestMultiError, or nil if none found.
func (m *GetWallpaperRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetWallpaperRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for WallpaperId

	if len(errors) > 0 {
		return GetWallpaperRequestMultiError(errors)
	}

	return nil
}

// GetWallpaperRequestMultiError is an error wrapping multiple validation
// errors returned by GetWallpaperRequest.ValidateAll() if the designated
// constraints aren't met.
type GetWallpaperRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetWallpaperRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetWallpaperRequestMultiError) AllErrors() []error { return m }

// GetWallpaperRequestValidationError is the validation error returned by
// GetWallpaperRequest.Validate if the designated constraints aren't met.
type GetWallpaperRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetWallpaperRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetWallpaperRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetWallpaperRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetWallpaperRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetWallpaperRequestValidationError) ErrorName() string {
	return "GetWallpaperRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetWallpaperRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetWallpaperRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetWallpaperRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetWallpaperRequestValidationError{}

// Validate checks the field values on GetWallpaperResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetWallpaperResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetWallpaperResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetWallpaperResponseMultiError, or nil if none found.
func (m *GetWallpaperResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetWallpaperResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetWallpaper()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetWallpaperResponseValidationError{
					field:  "Wallpaper",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetWallpaperResponseValidationError{
					field:  "Wallpaper",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWallpaper()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetWallpaperResponseValidationError{
				field:  "Wallpaper",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetWallpaperResponseMultiError(errors)
	}

	return nil
}

// GetWallpaperResponseMultiError is an error wrapping multiple validation
// errors returned by GetWallpaperResponse.ValidateAll() if the designated
// constraints aren't met.
type GetWallpaperResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetWallpaperResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetWallpaperResponseMultiError) AllErrors() []error { return m }

// GetWallpaperResponseValidationError is the validation error returned by
// GetWallpaperResponse.Validate if the designated constraints aren't met.
type GetWallpaperResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetWallpaperResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetWallpaperResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetWallpaperResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetWallpaperResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetWallpaperResponseValidationError) ErrorName() string {
	return "GetWallpaperResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetWallpaperResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetWallpaperResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetWallpaperResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetWallpaperResponseValidationError{}

// Validate checks the field values on ListWallpapersFilter with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListWallpapersFilter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListWallpapersFilter with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListWallpapersFilterMultiError, or nil if none found.
func (m *ListWallpapersFilter) ValidateAll() error {
	return m.validate(true)
}

func (m *ListWallpapersFilter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Term

	if len(errors) > 0 {
		return ListWallpapersFilterMultiError(errors)
	}

	return nil
}

// ListWallpapersFilterMultiError is an error wrapping multiple validation
// errors returned by ListWallpapersFilter.ValidateAll() if the designated
// constraints aren't met.
type ListWallpapersFilterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListWallpapersFilterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListWallpapersFilterMultiError) AllErrors() []error { return m }

// ListWallpapersFilterValidationError is the validation error returned by
// ListWallpapersFilter.Validate if the designated constraints aren't met.
type ListWallpapersFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListWallpapersFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListWallpapersFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListWallpapersFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListWallpapersFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListWallpapersFilterValidationError) ErrorName() string {
	return "ListWallpapersFilterValidationError"
}

// Error satisfies the builtin error interface
func (e ListWallpapersFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListWallpapersFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListWallpapersFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListWallpapersFilterValidationError{}

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

	if all {
		switch v := interface{}(m.GetFilter()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListWallpapersRequestValidationError{
					field:  "Filter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListWallpapersRequestValidationError{
					field:  "Filter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFilter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListWallpapersRequestValidationError{
				field:  "Filter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

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

// Validate checks the field values on SaveWallpaperRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SaveWallpaperRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SaveWallpaperRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SaveWallpaperRequestMultiError, or nil if none found.
func (m *SaveWallpaperRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SaveWallpaperRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Description

	// no validation rules for SizeMb

	// no validation rules for ThumbnailUrl

	// no validation rules for HdUrl

	if all {
		switch v := interface{}(m.GetPublisher()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SaveWallpaperRequestValidationError{
					field:  "Publisher",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SaveWallpaperRequestValidationError{
					field:  "Publisher",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPublisher()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SaveWallpaperRequestValidationError{
				field:  "Publisher",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetTags() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SaveWallpaperRequestValidationError{
						field:  fmt.Sprintf("Tags[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SaveWallpaperRequestValidationError{
						field:  fmt.Sprintf("Tags[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SaveWallpaperRequestValidationError{
					field:  fmt.Sprintf("Tags[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for SourceUrl

	// no validation rules for Height

	// no validation rules for Width

	// no validation rules for AgeRating

	if len(errors) > 0 {
		return SaveWallpaperRequestMultiError(errors)
	}

	return nil
}

// SaveWallpaperRequestMultiError is an error wrapping multiple validation
// errors returned by SaveWallpaperRequest.ValidateAll() if the designated
// constraints aren't met.
type SaveWallpaperRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SaveWallpaperRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SaveWallpaperRequestMultiError) AllErrors() []error { return m }

// SaveWallpaperRequestValidationError is the validation error returned by
// SaveWallpaperRequest.Validate if the designated constraints aren't met.
type SaveWallpaperRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SaveWallpaperRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SaveWallpaperRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SaveWallpaperRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SaveWallpaperRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SaveWallpaperRequestValidationError) ErrorName() string {
	return "SaveWallpaperRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SaveWallpaperRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveWallpaperRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SaveWallpaperRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SaveWallpaperRequestValidationError{}

// Validate checks the field values on SaveWallpaperResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SaveWallpaperResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SaveWallpaperResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SaveWallpaperResponseMultiError, or nil if none found.
func (m *SaveWallpaperResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SaveWallpaperResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetWallpaper()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SaveWallpaperResponseValidationError{
					field:  "Wallpaper",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SaveWallpaperResponseValidationError{
					field:  "Wallpaper",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWallpaper()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SaveWallpaperResponseValidationError{
				field:  "Wallpaper",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SaveWallpaperResponseMultiError(errors)
	}

	return nil
}

// SaveWallpaperResponseMultiError is an error wrapping multiple validation
// errors returned by SaveWallpaperResponse.ValidateAll() if the designated
// constraints aren't met.
type SaveWallpaperResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SaveWallpaperResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SaveWallpaperResponseMultiError) AllErrors() []error { return m }

// SaveWallpaperResponseValidationError is the validation error returned by
// SaveWallpaperResponse.Validate if the designated constraints aren't met.
type SaveWallpaperResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SaveWallpaperResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SaveWallpaperResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SaveWallpaperResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SaveWallpaperResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SaveWallpaperResponseValidationError) ErrorName() string {
	return "SaveWallpaperResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SaveWallpaperResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveWallpaperResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SaveWallpaperResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SaveWallpaperResponseValidationError{}

// Validate checks the field values on ImportWallpaperRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ImportWallpaperRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ImportWallpaperRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ImportWallpaperRequestMultiError, or nil if none found.
func (m *ImportWallpaperRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ImportWallpaperRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Url

	if len(errors) > 0 {
		return ImportWallpaperRequestMultiError(errors)
	}

	return nil
}

// ImportWallpaperRequestMultiError is an error wrapping multiple validation
// errors returned by ImportWallpaperRequest.ValidateAll() if the designated
// constraints aren't met.
type ImportWallpaperRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ImportWallpaperRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ImportWallpaperRequestMultiError) AllErrors() []error { return m }

// ImportWallpaperRequestValidationError is the validation error returned by
// ImportWallpaperRequest.Validate if the designated constraints aren't met.
type ImportWallpaperRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImportWallpaperRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImportWallpaperRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImportWallpaperRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImportWallpaperRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImportWallpaperRequestValidationError) ErrorName() string {
	return "ImportWallpaperRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ImportWallpaperRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImportWallpaperRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImportWallpaperRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImportWallpaperRequestValidationError{}

// Validate checks the field values on ImportWallpaperResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ImportWallpaperResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ImportWallpaperResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ImportWallpaperResponseMultiError, or nil if none found.
func (m *ImportWallpaperResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ImportWallpaperResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetWallpaper()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ImportWallpaperResponseValidationError{
					field:  "Wallpaper",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ImportWallpaperResponseValidationError{
					field:  "Wallpaper",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWallpaper()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ImportWallpaperResponseValidationError{
				field:  "Wallpaper",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ImportWallpaperResponseMultiError(errors)
	}

	return nil
}

// ImportWallpaperResponseMultiError is an error wrapping multiple validation
// errors returned by ImportWallpaperResponse.ValidateAll() if the designated
// constraints aren't met.
type ImportWallpaperResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ImportWallpaperResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ImportWallpaperResponseMultiError) AllErrors() []error { return m }

// ImportWallpaperResponseValidationError is the validation error returned by
// ImportWallpaperResponse.Validate if the designated constraints aren't met.
type ImportWallpaperResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImportWallpaperResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImportWallpaperResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImportWallpaperResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImportWallpaperResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImportWallpaperResponseValidationError) ErrorName() string {
	return "ImportWallpaperResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ImportWallpaperResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImportWallpaperResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImportWallpaperResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImportWallpaperResponseValidationError{}

// Validate checks the field values on ImportBookmarkRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ImportBookmarkRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ImportBookmarkRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ImportBookmarkRequestMultiError, or nil if none found.
func (m *ImportBookmarkRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ImportBookmarkRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Url

	if len(errors) > 0 {
		return ImportBookmarkRequestMultiError(errors)
	}

	return nil
}

// ImportBookmarkRequestMultiError is an error wrapping multiple validation
// errors returned by ImportBookmarkRequest.ValidateAll() if the designated
// constraints aren't met.
type ImportBookmarkRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ImportBookmarkRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ImportBookmarkRequestMultiError) AllErrors() []error { return m }

// ImportBookmarkRequestValidationError is the validation error returned by
// ImportBookmarkRequest.Validate if the designated constraints aren't met.
type ImportBookmarkRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImportBookmarkRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImportBookmarkRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImportBookmarkRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImportBookmarkRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImportBookmarkRequestValidationError) ErrorName() string {
	return "ImportBookmarkRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ImportBookmarkRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImportBookmarkRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImportBookmarkRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImportBookmarkRequestValidationError{}

// Validate checks the field values on ImportBookmarkResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ImportBookmarkResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ImportBookmarkResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ImportBookmarkResponseMultiError, or nil if none found.
func (m *ImportBookmarkResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ImportBookmarkResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ImportBookmarkResponseMultiError(errors)
	}

	return nil
}

// ImportBookmarkResponseMultiError is an error wrapping multiple validation
// errors returned by ImportBookmarkResponse.ValidateAll() if the designated
// constraints aren't met.
type ImportBookmarkResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ImportBookmarkResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ImportBookmarkResponseMultiError) AllErrors() []error { return m }

// ImportBookmarkResponseValidationError is the validation error returned by
// ImportBookmarkResponse.Validate if the designated constraints aren't met.
type ImportBookmarkResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImportBookmarkResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImportBookmarkResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImportBookmarkResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImportBookmarkResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImportBookmarkResponseValidationError) ErrorName() string {
	return "ImportBookmarkResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ImportBookmarkResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImportBookmarkResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImportBookmarkResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImportBookmarkResponseValidationError{}
