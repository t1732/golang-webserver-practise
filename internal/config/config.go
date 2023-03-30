package config

var (
	_appCnf AppConfig
	_dbCnf  dbConfig
)

func Init(env string) error {
	if err := initApplicationConfig(env); err != nil {
		return err
	}

	if err := initDatabaseConfig(); err != nil {
		return err
	}

	return nil
}

func App() *AppConfig {
	return &_appCnf
}

func DB() *dbConfig {
	return &_dbCnf
}
