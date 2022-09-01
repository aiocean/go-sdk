// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: wibuwallpaper/v1/wallpaper.proto

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

	// no validation rules for TotalViews

	// no validation rules for TotalLikes

	// no validation rules for TotalDownloads

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

	for idx, item := range m.GetTopics() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, WallpaperValidationError{
						field:  fmt.Sprintf("Topics[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, WallpaperValidationError{
						field:  fmt.Sprintf("Topics[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WallpaperValidationError{
					field:  fmt.Sprintf("Topics[%v]", idx),
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
