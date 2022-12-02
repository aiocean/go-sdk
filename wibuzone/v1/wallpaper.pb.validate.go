// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: wibuzone/v1/wallpaper.proto

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

// Validate checks the field values on Wallpaper with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Wallpaper) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Wallpaper with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in WallpaperMultiError, or nil
// if none found.
func (m *Wallpaper) ValidateAll() error {
	return m.validate(true)
}

func (m *Wallpaper) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Description

	// no validation rules for SizeMb

	if all {
		switch v := interface{}(m.GetPublishedTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WallpaperValidationError{
					field:  "PublishedTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WallpaperValidationError{
					field:  "PublishedTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPublishedTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WallpaperValidationError{
				field:  "PublishedTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WallpaperValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WallpaperValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WallpaperValidationError{
				field:  "CreateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WallpaperValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WallpaperValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WallpaperValidationError{
				field:  "UpdateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ThumbnailUrl

	// no validation rules for HdUrl

	if all {
		switch v := interface{}(m.GetPublisher()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WallpaperValidationError{
					field:  "Publisher",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WallpaperValidationError{
					field:  "Publisher",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPublisher()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WallpaperValidationError{
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
					errors = append(errors, WallpaperValidationError{
						field:  fmt.Sprintf("Tags[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, WallpaperValidationError{
						field:  fmt.Sprintf("Tags[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WallpaperValidationError{
					field:  fmt.Sprintf("Tags[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Visibility

	// no validation rules for SourceUrl

	// no validation rules for Height

	// no validation rules for Width

	if all {
		switch v := interface{}(m.GetStats()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WallpaperValidationError{
					field:  "Stats",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WallpaperValidationError{
					field:  "Stats",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStats()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WallpaperValidationError{
				field:  "Stats",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for AgeRating

	if len(errors) > 0 {
		return WallpaperMultiError(errors)
	}

	return nil
}

// WallpaperMultiError is an error wrapping multiple validation errors returned
// by Wallpaper.ValidateAll() if the designated constraints aren't met.
type WallpaperMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WallpaperMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WallpaperMultiError) AllErrors() []error { return m }

// WallpaperValidationError is the validation error returned by
// Wallpaper.Validate if the designated constraints aren't met.
type WallpaperValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WallpaperValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WallpaperValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WallpaperValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WallpaperValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WallpaperValidationError) ErrorName() string { return "WallpaperValidationError" }

// Error satisfies the builtin error interface
func (e WallpaperValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWallpaper.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WallpaperValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WallpaperValidationError{}

// Validate checks the field values on WallpaperStats with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *WallpaperStats) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WallpaperStats with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in WallpaperStatsMultiError,
// or nil if none found.
func (m *WallpaperStats) ValidateAll() error {
	return m.validate(true)
}

func (m *WallpaperStats) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for TotalViews

	// no validation rules for TotalLikes

	// no validation rules for TotalDownloads

	if len(errors) > 0 {
		return WallpaperStatsMultiError(errors)
	}

	return nil
}

// WallpaperStatsMultiError is an error wrapping multiple validation errors
// returned by WallpaperStats.ValidateAll() if the designated constraints
// aren't met.
type WallpaperStatsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WallpaperStatsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WallpaperStatsMultiError) AllErrors() []error { return m }

// WallpaperStatsValidationError is the validation error returned by
// WallpaperStats.Validate if the designated constraints aren't met.
type WallpaperStatsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WallpaperStatsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WallpaperStatsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WallpaperStatsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WallpaperStatsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WallpaperStatsValidationError) ErrorName() string { return "WallpaperStatsValidationError" }

// Error satisfies the builtin error interface
func (e WallpaperStatsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWallpaperStats.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WallpaperStatsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WallpaperStatsValidationError{}

// Validate checks the field values on UnlikeWallpaperRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UnlikeWallpaperRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UnlikeWallpaperRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UnlikeWallpaperRequestMultiError, or nil if none found.
func (m *UnlikeWallpaperRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UnlikeWallpaperRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for WallpaperId

	if len(errors) > 0 {
		return UnlikeWallpaperRequestMultiError(errors)
	}

	return nil
}

// UnlikeWallpaperRequestMultiError is an error wrapping multiple validation
// errors returned by UnlikeWallpaperRequest.ValidateAll() if the designated
// constraints aren't met.
type UnlikeWallpaperRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UnlikeWallpaperRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UnlikeWallpaperRequestMultiError) AllErrors() []error { return m }

// UnlikeWallpaperRequestValidationError is the validation error returned by
// UnlikeWallpaperRequest.Validate if the designated constraints aren't met.
type UnlikeWallpaperRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnlikeWallpaperRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnlikeWallpaperRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnlikeWallpaperRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnlikeWallpaperRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnlikeWallpaperRequestValidationError) ErrorName() string {
	return "UnlikeWallpaperRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UnlikeWallpaperRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnlikeWallpaperRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnlikeWallpaperRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnlikeWallpaperRequestValidationError{}

// Validate checks the field values on UnlikeWallpaperResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UnlikeWallpaperResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UnlikeWallpaperResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UnlikeWallpaperResponseMultiError, or nil if none found.
func (m *UnlikeWallpaperResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UnlikeWallpaperResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UnlikeWallpaperResponseMultiError(errors)
	}

	return nil
}

// UnlikeWallpaperResponseMultiError is an error wrapping multiple validation
// errors returned by UnlikeWallpaperResponse.ValidateAll() if the designated
// constraints aren't met.
type UnlikeWallpaperResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UnlikeWallpaperResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UnlikeWallpaperResponseMultiError) AllErrors() []error { return m }

// UnlikeWallpaperResponseValidationError is the validation error returned by
// UnlikeWallpaperResponse.Validate if the designated constraints aren't met.
type UnlikeWallpaperResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnlikeWallpaperResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnlikeWallpaperResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnlikeWallpaperResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnlikeWallpaperResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnlikeWallpaperResponseValidationError) ErrorName() string {
	return "UnlikeWallpaperResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UnlikeWallpaperResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnlikeWallpaperResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnlikeWallpaperResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnlikeWallpaperResponseValidationError{}

// Validate checks the field values on LikeWallpaperRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *LikeWallpaperRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LikeWallpaperRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LikeWallpaperRequestMultiError, or nil if none found.
func (m *LikeWallpaperRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LikeWallpaperRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for WallpaperId

	if len(errors) > 0 {
		return LikeWallpaperRequestMultiError(errors)
	}

	return nil
}

// LikeWallpaperRequestMultiError is an error wrapping multiple validation
// errors returned by LikeWallpaperRequest.ValidateAll() if the designated
// constraints aren't met.
type LikeWallpaperRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LikeWallpaperRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LikeWallpaperRequestMultiError) AllErrors() []error { return m }

// LikeWallpaperRequestValidationError is the validation error returned by
// LikeWallpaperRequest.Validate if the designated constraints aren't met.
type LikeWallpaperRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LikeWallpaperRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LikeWallpaperRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LikeWallpaperRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LikeWallpaperRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LikeWallpaperRequestValidationError) ErrorName() string {
	return "LikeWallpaperRequestValidationError"
}

// Error satisfies the builtin error interface
func (e LikeWallpaperRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLikeWallpaperRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LikeWallpaperRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LikeWallpaperRequestValidationError{}

// Validate checks the field values on LikeWallpaperResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *LikeWallpaperResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LikeWallpaperResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LikeWallpaperResponseMultiError, or nil if none found.
func (m *LikeWallpaperResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *LikeWallpaperResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return LikeWallpaperResponseMultiError(errors)
	}

	return nil
}

// LikeWallpaperResponseMultiError is an error wrapping multiple validation
// errors returned by LikeWallpaperResponse.ValidateAll() if the designated
// constraints aren't met.
type LikeWallpaperResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LikeWallpaperResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LikeWallpaperResponseMultiError) AllErrors() []error { return m }

// LikeWallpaperResponseValidationError is the validation error returned by
// LikeWallpaperResponse.Validate if the designated constraints aren't met.
type LikeWallpaperResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LikeWallpaperResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LikeWallpaperResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LikeWallpaperResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LikeWallpaperResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LikeWallpaperResponseValidationError) ErrorName() string {
	return "LikeWallpaperResponseValidationError"
}

// Error satisfies the builtin error interface
func (e LikeWallpaperResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLikeWallpaperResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LikeWallpaperResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LikeWallpaperResponseValidationError{}