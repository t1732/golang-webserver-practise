package config

var (
	DB dbConfig
)

func Init(env string) error {
	if err := initDatabaseConfig(); err != nil {
		return err
	}

	return nil
}
