// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: calculator.proto

package calculator

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on BinaryOperation with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *BinaryOperation) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for FirstOperand

	// no validation rules for SecondOperand

	// no validation rules for Operation

	return nil
}

// BinaryOperationValidationError is the validation error returned by
// BinaryOperation.Validate if the designated constraints aren't met.
type BinaryOperationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BinaryOperationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BinaryOperationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BinaryOperationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BinaryOperationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BinaryOperationValidationError) ErrorName() string { return "BinaryOperationValidationError" }

// Error satisfies the builtin error interface
func (e BinaryOperationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBinaryOperation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BinaryOperationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BinaryOperationValidationError{}

// Validate checks the field values on CalculationResult with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CalculationResult) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Result

	return nil
}

// CalculationResultValidationError is the validation error returned by
// CalculationResult.Validate if the designated constraints aren't met.
type CalculationResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CalculationResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CalculationResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CalculationResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CalculationResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CalculationResultValidationError) ErrorName() string {
	return "CalculationResultValidationError"
}

// Error satisfies the builtin error interface
func (e CalculationResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCalculationResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CalculationResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CalculationResultValidationError{}