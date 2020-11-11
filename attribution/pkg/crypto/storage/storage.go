package storage

type storage interface {
	Storage(groupId string, encryptKey string) error
	Fetch(groupId string) (string, error)
}
