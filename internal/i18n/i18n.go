// Package i18n provides internationalization support for the application.
package i18n

import (
	"embed"
	"fmt"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"

	"github.com/cresplanex/bloader/internal/config"
)

// Message is a string that can be localized.
type Message struct {
	// ID uniquely identifies the message.
	ID string

	// Hash uniquely identifies the content of the message
	// that this message was translated from.
	Hash string

	// Description describes the message to give additional
	// context to translators that may be relevant for translation.
	Description string

	// LeftDelim is the left Go template delimiter.
	LeftDelim string

	// RightDelim is the right Go template delimiter.
	RightDelim string

	// Zero is the content of the message for the CLDR plural form "zero".
	Zero string

	// One is the content of the message for the CLDR plural form "one".
	One string

	// Two is the content of the message for the CLDR plural form "two".
	Two string

	// Few is the content of the message for the CLDR plural form "few".
	Few string

	// Many is the content of the message for the CLDR plural form "many".
	Many string

	// Other is the content of the message for the CLDR plural form "other".
	Other string
}

// Locale is a locale for a language.
type Locale struct {
	call config.LanguageType
	tag  language.Tag
}

var (
	// English is a locale for English.
	English = Locale{
		call: config.LanguageTypeEnglish,
		tag:  language.English,
	}
	// Japanese is a locale for Japanese.
	Japanese = Locale{
		call: config.LanguageTypeJapanese,
		tag:  language.Japanese,
	}

	// Default is the default locale.
	Default = English
)

var langs = []Locale{English, Japanese}

//go:embed *
var files embed.FS

// Translation is an interface for translating messages.
type Translation interface {
	Translate(id string) string
	TranslateWithOpts(id string, opts Options) string
	TranslateWithLocale(locale Locale, id string) string
	TranslateWithLocaleAndOpts(locale Locale, id string, opts Options) string
}

// Options is used to configure a translation.
type Options struct {
	// TemplateData is the data passed when executing the message's template.
	// If TemplateData is nil and PluralCount is not nil, then the message template
	// will be executed with data that contains the plural count.
	TemplateData any

	// PluralCount determines which plural form of the message is used.
	PluralCount any

	// DefaultMessage is used if the message is not found in any message files.
	DefaultMessage *Message
}

// Translator is a translation service.
type Translator struct {
	mu     sync.RWMutex
	bundle *i18n.Bundle
}

// NewTranslator creates a new Translator.
func NewTranslator() (*Translator, error) {
	t := &Translator{}
	bundle := i18n.NewBundle(Default.tag)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	for _, l := range langs {
		if _, err := bundle.LoadMessageFileFS(files, "lang/"+string(l.call)+".toml"); err != nil {
			return nil, fmt.Errorf("failed to load message file: %w", err)
		}
	}
	t.bundle = bundle
	return t, nil
}

// Translate translates a message with the given ID.
func (t *Translator) Translate(id string) string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	localizer := i18n.NewLocalizer(t.bundle, Default.tag.String())
	return localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: id,
	})
}

// TranslateWithOpts translates a message with the given ID.
func (t *Translator) TranslateWithOpts(id string, opts Options) string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	localizer := i18n.NewLocalizer(t.bundle, Default.tag.String())
	defMsg := &i18n.Message{
		ID:          opts.DefaultMessage.ID,
		Hash:        opts.DefaultMessage.Hash,
		Description: opts.DefaultMessage.Description,
		LeftDelim:   opts.DefaultMessage.LeftDelim,
		RightDelim:  opts.DefaultMessage.RightDelim,
		Zero:        opts.DefaultMessage.Zero,
		One:         opts.DefaultMessage.One,
		Two:         opts.DefaultMessage.Two,
		Few:         opts.DefaultMessage.Few,
		Many:        opts.DefaultMessage.Many,
		Other:       opts.DefaultMessage.Other,
	}
	return localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:      id,
		TemplateData:   opts.TemplateData,
		PluralCount:    opts.PluralCount,
		DefaultMessage: defMsg,
	})
}

// TranslateWithLocale translates a message with the given ID to the given locale.
func (t *Translator) TranslateWithLocale(locale Locale, id string) string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	localizer := i18n.NewLocalizer(t.bundle, locale.tag.String())
	return localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: id,
	})
}

// TranslateWithLocaleAndOpts translates a message with the given ID to the given locale.
func (t *Translator) TranslateWithLocaleAndOpts(locale Locale, id string, opts Options) string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	localizer := i18n.NewLocalizer(t.bundle, locale.tag.String())
	defMsg := &i18n.Message{
		ID:          opts.DefaultMessage.ID,
		Hash:        opts.DefaultMessage.Hash,
		Description: opts.DefaultMessage.Description,
		LeftDelim:   opts.DefaultMessage.LeftDelim,
		RightDelim:  opts.DefaultMessage.RightDelim,
		Zero:        opts.DefaultMessage.Zero,
		One:         opts.DefaultMessage.One,
		Two:         opts.DefaultMessage.Two,
		Few:         opts.DefaultMessage.Few,
		Many:        opts.DefaultMessage.Many,
		Other:       opts.DefaultMessage.Other,
	}
	return localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:      id,
		TemplateData:   opts.TemplateData,
		PluralCount:    opts.PluralCount,
		DefaultMessage: defMsg,
	})
}

var _ Translation = &Translator{}
