package store

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStorage()
}
