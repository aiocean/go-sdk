// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: wibuzone/v1/collection.proto

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

// Validate checks the field values on Collection with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Collection) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Collection with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CollectionMultiError, or
// nil if none found.
func (m *Collection) ValidateAll() error {
	return m.validate(true)
}

func (m *Collection) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetCreateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CollectionValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CollectionValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CollectionValidationError{
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
				errors = append(errors, CollectionValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CollectionValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CollectionValidationError{
				field:  "UpdateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for CoverUrl

	if all {
		switch v := interface{}(m.GetPublisher()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CollectionValidationError{
					field:  "Publisher",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CollectionValidationError{
					field:  "Publisher",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPublisher()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CollectionValidationError{
				field:  "Publisher",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Visibility

	for idx, item := range m.GetTags() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CollectionValidationError{
						field:  fmt.Sprintf("Tags[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CollectionValidationError{
						field:  fmt.Sprintf("Tags[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CollectionValidationError{
					field:  fmt.Sprintf("Tags[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CollectionMultiError(errors)
	}

	return nil
}

// CollectionMultiError is an error wrapping multiple validation errors
// returned by Collection.ValidateAll() if the designated constraints aren't met.
type CollectionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CollectionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CollectionMultiError) AllErrors() []error { return m }

// CollectionValidationError is the validation error returned by
// Collection.Validate if the designated constraints aren't met.
type CollectionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CollectionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CollectionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CollectionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CollectionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CollectionValidationError) ErrorName() string { return "CollectionValidationError" }

// Error satisfies the builtin error interface
func (e CollectionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCollection.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CollectionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CollectionValidationError{}

// Validate checks the field values on CollectionStats with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CollectionStats) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CollectionStats with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CollectionStatsMultiError, or nil if none found.
func (m *CollectionStats) ValidateAll() error {
	return m.validate(true)
}

func (m *CollectionStats) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TotalPhotos

	if len(errors) > 0 {
		return CollectionStatsMultiError(errors)
	}

	return nil
}

// CollectionStatsMultiError is an error wrapping multiple validation errors
// returned by CollectionStats.ValidateAll() if the designated constraints
// aren't met.
type CollectionStatsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CollectionStatsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CollectionStatsMultiError) AllErrors() []error { return m }

// CollectionStatsValidationError is the validation error returned by
// CollectionStats.Validate if the designated constraints aren't met.
type CollectionStatsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CollectionStatsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CollectionStatsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CollectionStatsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CollectionStatsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CollectionStatsValidationError) ErrorName() string { return "CollectionStatsValidationError" }

// Error satisfies the builtin error interface
func (e CollectionStatsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCollectionStats.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CollectionStatsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CollectionStatsValidationError{}