package config

var (
	App AppConfig
	DB  dbConfig
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
