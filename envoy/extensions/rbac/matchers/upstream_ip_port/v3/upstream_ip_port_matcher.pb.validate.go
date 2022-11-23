// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/rbac/matchers/upstream_ip_port/v3/upstream_ip_port_matcher.proto

package envoy_extensions_rbac_matchers_upstream_ip_port_v3

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

// Validate checks the field values on UpstreamIpPortMatcher with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpstreamIpPortMatcher) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpstreamIpPortMatcher with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpstreamIpPortMatcherMultiError, or nil if none found.
func (m *UpstreamIpPortMatcher) ValidateAll() error {
	return m.validate(true)
}

func (m *UpstreamIpPortMatcher) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUpstreamIp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpstreamIpPortMatcherValidationError{
					field:  "UpstreamIp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpstreamIpPortMatcherValidationError{
					field:  "UpstreamIp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpstreamIp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpstreamIpPortMatcherValidationError{
				field:  "UpstreamIp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpstreamPortRange()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpstreamIpPortMatcherValidationError{
					field:  "UpstreamPortRange",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpstreamIpPortMatcherValidationError{
					field:  "UpstreamPortRange",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpstreamPortRange()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpstreamIpPortMatcherValidationError{
				field:  "UpstreamPortRange",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpstreamIpPortMatcherMultiError(errors)
	}
	return nil
}

// UpstreamIpPortMatcherMultiError is an error wrapping multiple validation
// errors returned by UpstreamIpPortMatcher.ValidateAll() if the designated
// constraints aren't met.
type UpstreamIpPortMatcherMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpstreamIpPortMatcherMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpstreamIpPortMatcherMultiError) AllErrors() []error { return m }

// UpstreamIpPortMatcherValidationError is the validation error returned by
// UpstreamIpPortMatcher.Validate if the designated constraints aren't met.
type UpstreamIpPortMatcherValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpstreamIpPortMatcherValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpstreamIpPortMatcherValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpstreamIpPortMatcherValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpstreamIpPortMatcherValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpstreamIpPortMatcherValidationError) ErrorName() string {
	return "UpstreamIpPortMatcherValidationError"
}

// Error satisfies the builtin error interface
func (e UpstreamIpPortMatcherValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpstreamIpPortMatcher.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpstreamIpPortMatcherValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpstreamIpPortMatcherValidationError{}
