package services

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"

	"github.com/mikestefanello/backlite"
	"github.com/rs/zerolog"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/funcmap"
	slogzerolog "github.com/samber/slog-zerolog/v2"

	// Require by ent
	_ "github.com/mikestefanello/pagoda/ent/runtime"
)

// Container contains all services used by the application and provides an easy way to handle dependency
// injection including within tests
type Container struct {
	// Config stores the application configuration
	Config *config.Config

	// Logger stores the logger
	Logger zerolog.Logger

	// Validator stores a validator
	Validator *Validator

	// Web stores the web framework
	Web *echo.Echo

	// Cache contains the cache client
	Cache *CacheClient

	// Database stores the connection to the database
	Database *sql.DB

	// ORM stores a client to the ORM
	ORM *ent.Client

	// Clients for repository services
	Users *UsersClient

	// Auth stores an Ory authentication client. AuthClient is an interface, hence not a pointer
	Auth *OryAuthClient

	// TemplateRenderer stores a service to easily render and cache templates
	TemplateRenderer *TemplateRenderer

	// Mail stores an email sending client
	Mail *MailClient

	// Tasks stores the task client
	Tasks *backlite.Client
}

// NewContainer creates and initializes a new Container
func NewContainer() *Container {
	c := new(Container)
	c.initConfig()
	c.initLogger()
	c.initValidator()
	c.initWeb()
	c.initCache()
	c.initDatabase()
	c.initORM()
	c.initUsers()
	c.initAuth()
	c.initTemplateRenderer()
	c.initMail()
	c.initTasks()
	return c
}

// Shutdown shuts the Container down and disconnects all connections.
// If the task runner was started, cancel the context to shut it down prior to calling this.
func (c *Container) Shutdown() error {
	if err := c.ORM.Close(); err != nil {
		return err
	}
	if err := c.Database.Close(); err != nil {
		return err
	}
	c.Cache.Close()

	return nil
}

// initConfig initializes configuration
func (c *Container) initConfig() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	c.Config = &cfg
}

func (c *Container) initLogger() {
	// Use short file names for location
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}
	// Set out based on environment
	var out io.Writer
	switch c.Config.App.Environment {
	case config.EnvLocal:
		out = zerolog.ConsoleWriter{Out: os.Stdout}
	default:
		out = os.Stdout
	}
	// Create logger with configured level
	logLevel, err := zerolog.ParseLevel(c.Config.App.LogLevel)
	if err != nil {
		panic(fmt.Sprintf("invalid log level in config file: %v", c.Config.App.LogLevel))
	}
	c.Logger = zerolog.New(out).With().Timestamp().Caller().Logger().Level(logLevel)
}

// initValidator initializes the validator
func (c *Container) initValidator() {
	c.Validator = NewValidator()
}

// initWeb initializes the web framework
func (c *Container) initWeb() {
	c.Web = echo.New()
	c.Web.HideBanner = true
	c.Web.HidePort = true
	c.Web.Validator = c.Validator
}

// initCache initializes the cache
func (c *Container) initCache() {
	store, err := newInMemoryCache(c.Config.Cache.Capacity)
	if err != nil {
		c.Logger.Panic().Err(err).Msg("failed to create the cache")
	}

	c.Cache = NewCacheClient(store)
}

// initDatabase initializes the database
func (c *Container) initDatabase() {
	var err error
	c.Database, err = sql.Open(c.Config.Database.Driver, c.Config.Database.Connection)
	if err != nil {
		c.Logger.Panic().Err(err).Msg("failed to connect to database")
	}
}

// initORM initializes the ORM
func (c *Container) initORM() {
	drv := entsql.OpenDB(c.Config.Database.Driver, c.Database)
	c.ORM = ent.NewClient(ent.Driver(drv))

	// Run the auto migration tool.
	if err := c.ORM.Schema.Create(context.Background()); err != nil {
		c.Logger.Panic().Err(err).Msg("failed to create database schema")
	}
}

func (c *Container) initUsers() {
	c.Users = NewUsersClient(c.ORM)
}

// initAuth initializes the authentication client
func (c *Container) initAuth() {
	c.Auth = newOryAuthClient(c.Config, c.Users)
}

// initTemplateRenderer initializes the template renderer
func (c *Container) initTemplateRenderer() {
	c.TemplateRenderer = NewTemplateRenderer(c.Config, c.Cache, funcmap.NewFuncMap(c.Web))
}

// initMail initialize the mail client
func (c *Container) initMail() {
	var err error
	c.Mail, err = NewMailClient(c.Config, c.TemplateRenderer)
	if err != nil {
		c.Logger.Panic().Err(err).Msg("failed to create mail client")
	}
}

// initTasks initializes the task client
func (c *Container) initTasks() {
	var err error
	logger := slog.New(
		slogzerolog.Option{
			Level:  slog.LevelDebug, // TODO: convert Config.App.LogLevel
			Logger: &c.Logger,
		}.NewZerologHandler())
	// You could use a separate database for tasks, if you'd like. but using one
	// makes transaction support easier
	c.Tasks, err = backlite.NewClient(backlite.ClientConfig{
		DB:              c.Database,
		Logger:          logger,
		NumWorkers:      c.Config.Tasks.Goroutines,
		ReleaseAfter:    c.Config.Tasks.ReleaseAfter,
		CleanupInterval: c.Config.Tasks.CleanupInterval,
	})

	if err != nil {
		c.Logger.Panic().Err(err).Msg("failed to create task client")
	}

	if err = c.Tasks.Install(); err != nil {
		c.Logger.Panic().Err(err).Msg("failed to install task schema")
	}
}
