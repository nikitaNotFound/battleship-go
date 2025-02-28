package storage

type StorageConfig struct {
	ConnectionString string
}

func NewStorageConfig(connectionString string) (*StorageConfig, error) {
	return &StorageConfig{
		ConnectionString: connectionString,
	}, nil
}
