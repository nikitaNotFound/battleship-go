package storage

type BattleshipStorage struct {
	Game *GameStorage
	User *UserStorage
}

func InitializeStorage(cfg *StorageConfig) (*BattleshipStorage, error) {
	gameStorage := &GameStorage{}
	userStorage := &UserStorage{}

	return &BattleshipStorage{
		Game: gameStorage,
		User: userStorage,
	}, nil
}
