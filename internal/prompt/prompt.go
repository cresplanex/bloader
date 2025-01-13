// Package prompt provides a set of functions to prompt the user to enter a value.
package prompt

import (
	"fmt"
	"strconv"
	"time"

	"github.com/manifoldco/promptui"
)

var (
	// ErrorMaxAttemptExceeded is returned when the maximum number of attempts is exceeded.
	ErrorMaxAttemptExceeded = fmt.Errorf("max attempts exceeded")
	// ErrorPasswordMismatch is returned when the password and the confirmation password do not match.
	ErrorPasswordMismatch = fmt.Errorf("password mismatch")
	// ErrorTimeDurationFormat is returned when the time duration format is invalid.
	ErrorTimeDurationFormat = fmt.Errorf("invalid time duration format")
	// ErrorTimeFormat is returned when the time format is invalid.
	ErrorTimeFormat = fmt.Errorf("invalid time format")
	// ErrorNumberFormat is returned when the number format is invalid.
	ErrorNumberFormat = fmt.Errorf("invalid number format")
)

// AgainCallback is a callback function that is called when the user
// needs to be prompted again.
type AgainCallback func(error) bool

type timeUnit string

const (
	// Nanosecond timeUnit = "ns"
	Nanosecond timeUnit = "ns"
	// Microsecond timeUnit = "us"
	Microsecond timeUnit = "us"
	// Millisecond timeUnit = "ms"
	Millisecond timeUnit = "ms"
	// Second timeUnit = "s"
	Second timeUnit = "s"
	// Minute timeUnit = "m"
	Minute timeUnit = "m"
	// Hour timeUnit = "h"
	Hour timeUnit = "h"
	// Day timeUnit = "d"
	Day timeUnit = "d"
	// Week timeUnit = "w"
	Week timeUnit = "w"
	// Month timeUnit = "M"
	Month timeUnit = "M"
	// Year timeUnit = "y"
	Year timeUnit = "y"
)

// Number prompts the user to enter a number.
func Number(
	label string,
	maxAttempts int,
	promptAgainCallback AgainCallback,
	isConfirm bool,
) (int, error) {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: isConfirm,
	}
	var value int
	var success bool
	for i := 0; i < maxAttempts; i++ {
		result, err := prompt.Run()
		if err != nil {
			return 0, fmt.Errorf("prompt number: %w", err)
		}
		value, err = strconv.Atoi(result)
		if err != nil {
			if pass := promptAgainCallback(ErrorNumberFormat); !pass {
				continue
			}
		}
		success = true
		break
	}
	if !success {
		return 0, ErrorMaxAttemptExceeded
	}
	return value, nil
}

// Select prompts the user to select an item from a list.
func Select(
	label string,
	items []string,
) (string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}
	index, _, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt select: %w", err)
	}
	return items[index], nil
}

// Duration prompts the user to enter a duration.
func Duration(
	label string,
	timeUnit timeUnit,
	maxAttempts int,
	promptAgainCallback AgainCallback,
	isConfirm bool,
) (time.Duration, error) {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: isConfirm,
	}
	var value time.Duration
	var success bool
	for i := 0; i < maxAttempts; i++ {
		result, err := prompt.Run()
		if err != nil {
			return 0, fmt.Errorf("prompt duration: %w", err)
		}
		value, err = time.ParseDuration(result + string(timeUnit))
		if err != nil {
			if pass := promptAgainCallback(ErrorTimeDurationFormat); !pass {
				continue
			}
		}
		success = true
		break
	}
	if !success {
		return 0, ErrorMaxAttemptExceeded
	}
	return value, nil
}

// Bool prompts the user to enter a boolean value.
func Bool(label string) (bool, error) {
	prompt := promptui.Select{
		Label: label,
		Items: []string{"yes", "no"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		return false, fmt.Errorf("prompt bool: %w", err)
	}
	return result == "Yes", nil
}

// Password prompts the user to enter a password.
func Password(
	label string,
	confirmLabel string,
	maxAttempts int,
	promptAgainCallback AgainCallback,
	isConfirm bool,
) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
		Mask:  '*',
	}
	confirmPrompt := promptui.Prompt{
		Label: confirmLabel,
		Mask:  '*',
	}
	var value string
	var success bool
	for i := 0; i < maxAttempts; i++ {
		result, err := prompt.Run()
		if err != nil {
			return "", fmt.Errorf("prompt password: %w", err)
		}
		value = result
		if isConfirm {
			confirmValue, err := confirmPrompt.Run()
			if err != nil {
				return "", fmt.Errorf("prompt password: %w", err)
			}
			if value != confirmValue {
				if pass := promptAgainCallback(ErrorPasswordMismatch); !pass {
					continue
				}
			}
		}
		success = true
		break
	}
	if !success {
		return "", ErrorMaxAttemptExceeded
	}
	return value, nil
}

// Text prompts the user to enter a text.
func Text(
	label string,
	isConfirm bool,
) (string, error) {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: isConfirm,
	}
	result, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt text: %w", err)
	}
	return result, nil
}

// Time prompts the user to enter a time.
func Time(
	label string,
	layout string,
	maxAttempts int,
	promptAgainCallback AgainCallback,
	isConfirm bool,
) (time.Time, error) {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: isConfirm,
	}
	var value time.Time
	var success bool
	for i := 0; i < maxAttempts; i++ {
		result, err := prompt.Run()
		if err != nil {
			return time.Time{}, fmt.Errorf("prompt time: %w", err)
		}
		value, err = time.Parse(layout, result)
		if err != nil {
			if pass := promptAgainCallback(ErrorTimeFormat); !pass {
				continue
			}
		}
		success = true
		break
	}

	if !success {
		return time.Time{}, ErrorMaxAttemptExceeded
	}
	return value, nil
}
