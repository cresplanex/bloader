// Package container provides the dependencies for the application.
package container

import (
	"context"
	"fmt"
	"time"

	"github.com/cresplanex/bloader/internal/auth"
	"github.com/cresplanex/bloader/internal/clock"
	"github.com/cresplanex/bloader/internal/clock/fakeclock"
	"github.com/cresplanex/bloader/internal/config"
	"github.com/cresplanex/bloader/internal/encrypt"
	"github.com/cresplanex/bloader/internal/i18n"
	"github.com/cresplanex/bloader/internal/logger"
	"github.com/cresplanex/bloader/internal/store"
	"github.com/cresplanex/bloader/internal/target"
)

// Container holds the dependencies for the application
type Container struct {
	Ctx                    context.Context
	Clocker                clock.Clock
	Translator             i18n.Translation
	Config                 config.ValidConfig
	Logger                 logger.Logger
	Store                  store.Store
	EncypterContainer      encrypt.Container
	AuthenticatorContainer auth.AuthenticatorContainer
	TargetContainer        target.Container
}

// NewContainer creates a new Container
func NewContainer() *Container {
	return &Container{}
}

// Init initializes the Container
func (c *Container) Init(cfg config.ValidConfig) error {
	c.Ctx = context.Background()
	var err error

	// ----------------------------------------
	// Set Config
	// ----------------------------------------
	c.Config = cfg

	// ----------------------------------------
	// Set Default Language
	// ----------------------------------------
	switch c.Config.Language.Default {
	case config.LanguageTypeEnglish:
		i18n.Default = i18n.English
	case config.LanguageTypeJapanese:
		i18n.Default = i18n.Japanese
	default:
		c.Config.Language.Default = config.LanguageTypeEnglish
		i18n.Default = i18n.English
	}

	// ----------------------------------------
	// Set Clock
	// ----------------------------------------
	if _, err = time.Parse(c.Config.Clock.Format, c.Config.Clock.Format); err != nil {
		c.Config.Clock.Format = "2006-01-02 15:04:05"
	}

	clk := clock.New()
	if cfg.Clock.Fake.Enabled {
		fakeClk := fakeclock.New(cfg.Clock.Fake.Time)
		clk = fakeClk
	}
	c.Clocker = clk

	//----------------------------------------
	// Set Translator
	//----------------------------------------
	c.Translator, err = i18n.NewTranslator()
	if err != nil {
		return fmt.Errorf("failed to create translator: %w", err)
	}

	// ----------------------------------------
	// Set Logger
	// ----------------------------------------
	c.Logger, err = logger.NewLoggerFromConfig(cfg.Env, cfg.Logging)
	if err != nil {
		return fmt.Errorf("failed to create logger: %w", err)
	}
	if err := c.Logger.SetupLogger(cfg.Env, cfg.Logging); err != nil {
		return fmt.Errorf("failed to setup logger: %w", err)
	}

	if cfg.Type != config.ConfigTypeSlave {
		// ----------------------------------------
		// Set Store
		// ----------------------------------------
		c.Store, err = store.NewStoreFromConfig(cfg.Store)
		if err != nil {
			return fmt.Errorf("failed to create store: %w", err)
		}
		if err := c.Store.SetupStore(cfg.Env, cfg.Store); err != nil {
			return fmt.Errorf("failed to setup store: %w", err)
		}
	}

	// ----------------------------------------
	// Set Encrypter
	// ----------------------------------------
	c.EncypterContainer, err = encrypt.NewContainerFromConfig(c.Store, cfg.Encrypts)
	if err != nil {
		return fmt.Errorf("failed to create encrypter: %w", err)
	}

	if cfg.Type == config.ConfigTypeSlave {
		return nil
	}

	// ----------------------------------------
	// Set AuthToken
	// ----------------------------------------
	c.AuthenticatorContainer, err = auth.NewAuthenticatorContainerFromConfig(
		c.Store,
		c.Config,
	)
	if err != nil {
		return fmt.Errorf("failed to create authenticator container: %w", err)
	}

	// ----------------------------------------
	// Set Target
	// ----------------------------------------
	c.TargetContainer = target.NewContainer(cfg.Env, cfg.Targets)

	return nil
}

// Close closes the Container
func (c *Container) Close() error {
	if c.Logger != nil {
		if err := c.Logger.Close(); err != nil {
			return fmt.Errorf("failed to close logger: %w", err)
		}
	}
	if c.Store != nil {
		if err := c.Store.Close(); err != nil {
			return fmt.Errorf("failed to close store: %w", err)
		}
	}
	return nil
}
