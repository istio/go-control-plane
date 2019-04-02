// Code generated by protoc-gen-validate
// source: envoy/config/filter/fault/v2/fault.proto
// DO NOT EDIT!!!

package v2

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on FaultDelay with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *FaultDelay) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Type

	if v, ok := interface{}(m.GetPercentage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FaultDelayValidationError{
				Field:  "Percentage",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	switch m.FaultDelaySecifier.(type) {

	case *FaultDelay_FixedDelay:

		if d := m.GetFixedDelay(); d != nil {
			dur := *d

			gt := time.Duration(0*time.Second + 0*time.Nanosecond)

			if dur <= gt {
				return FaultDelayValidationError{
					Field:  "FixedDelay",
					Reason: "value must be greater than 0s",
				}
			}

		}

	case *FaultDelay_HeaderDelay_:

		if v, ok := interface{}(m.GetHeaderDelay()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FaultDelayValidationError{
					Field:  "HeaderDelay",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	default:
		return FaultDelayValidationError{
			Field:  "FaultDelaySecifier",
			Reason: "value is required",
		}

	}

	return nil
}

// FaultDelayValidationError is the validation error returned by
// FaultDelay.Validate if the designated constraints aren't met.
type FaultDelayValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e FaultDelayValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultDelay.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = FaultDelayValidationError{}

// Validate checks the field values on FaultRateLimit with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *FaultRateLimit) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPercentage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FaultRateLimitValidationError{
				Field:  "Percentage",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	switch m.LimitType.(type) {

	case *FaultRateLimit_FixedLimit_:

		if v, ok := interface{}(m.GetFixedLimit()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FaultRateLimitValidationError{
					Field:  "FixedLimit",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *FaultRateLimit_HeaderLimit_:

		if v, ok := interface{}(m.GetHeaderLimit()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FaultRateLimitValidationError{
					Field:  "HeaderLimit",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	default:
		return FaultRateLimitValidationError{
			Field:  "LimitType",
			Reason: "value is required",
		}

	}

	return nil
}

// FaultRateLimitValidationError is the validation error returned by
// FaultRateLimit.Validate if the designated constraints aren't met.
type FaultRateLimitValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e FaultRateLimitValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultRateLimit.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = FaultRateLimitValidationError{}

// Validate checks the field values on FaultDelay_HeaderDelay with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FaultDelay_HeaderDelay) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// FaultDelay_HeaderDelayValidationError is the validation error returned by
// FaultDelay_HeaderDelay.Validate if the designated constraints aren't met.
type FaultDelay_HeaderDelayValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e FaultDelay_HeaderDelayValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultDelay_HeaderDelay.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = FaultDelay_HeaderDelayValidationError{}

// Validate checks the field values on FaultRateLimit_FixedLimit with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FaultRateLimit_FixedLimit) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetLimitKbps() < 1 {
		return FaultRateLimit_FixedLimitValidationError{
			Field:  "LimitKbps",
			Reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// FaultRateLimit_FixedLimitValidationError is the validation error returned by
// FaultRateLimit_FixedLimit.Validate if the designated constraints aren't met.
type FaultRateLimit_FixedLimitValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e FaultRateLimit_FixedLimitValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultRateLimit_FixedLimit.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = FaultRateLimit_FixedLimitValidationError{}

// Validate checks the field values on FaultRateLimit_HeaderLimit with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FaultRateLimit_HeaderLimit) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// FaultRateLimit_HeaderLimitValidationError is the validation error returned
// by FaultRateLimit_HeaderLimit.Validate if the designated constraints aren't met.
type FaultRateLimit_HeaderLimitValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e FaultRateLimit_HeaderLimitValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultRateLimit_HeaderLimit.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = FaultRateLimit_HeaderLimitValidationError{}