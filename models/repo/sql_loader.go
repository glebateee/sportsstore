package repo

import (
	"database/sql"
	"os"
	"platform/config"
	"platform/logging"
	"reflect"

	_ "modernc.org/sqlite"
)

func openDB(cfg config.Configuration, logger logging.Logger) (db *sql.DB, commands *SqlCommands, needInit bool) {
	driver := cfg.GetStringDefault("sql:driver_name", "sqlite")
	dbName, found := cfg.GetString("sql:connection_str")
	if !found {
		logger.Panic("Cannot read SQL connection string from config")
		return
	}
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		needInit = true
	}
	var err error
	if db, err = sql.Open(driver, dbName); err == nil {
		commands = loadCommands(db, cfg, logger)
	} else {
		logger.Panic(err.Error())
	}
	return
}

func loadCommands(db *sql.DB, cfg config.Configuration, logger logging.Logger) (commands *SqlCommands) {
	commands = &SqlCommands{}
	commandVal := reflect.ValueOf(commands).Elem()
	commandType := reflect.TypeOf(commands).Elem()
	for i := range commandType.NumField() {
		commandName := commandType.Field(i).Name
		logger.Debugf("Loading SQL command: %v", commandName)
		stmt := prepareCommand(db, commandName, cfg, logger)
		commandVal.Field(i).Set(reflect.ValueOf(stmt))
	}
	return commands
}

func prepareCommand(db *sql.DB, commandName string, cfg config.Configuration, logger logging.Logger) *sql.Stmt {
	commandFile, found := cfg.GetString("sql:commands:" + commandName)
	if !found {
		logger.Panicf("Config does not contain location for SQL command: %v", commandName)
	}
	data, err := os.ReadFile(commandFile)
	if err != nil {
		logger.Panicf("Cannot read SQL command file: %v", commandFile)
	}
	stmt, err := db.Prepare(string(data))
	if err != nil {
		logger.Panicf(err.Error())
	}
	return stmt
}
