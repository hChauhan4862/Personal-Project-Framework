package DB

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/storage/redis/v3"
)

func getStorageDBCredential() (string, error) {
	var dsn string
	if err := Q.APPCONFIG.DO.Where(Q.APPCONFIG.NAME.Eq("SERVER_CONFIG_REDIS")).Select(Q.APPCONFIG.VALUE).Scan(&dsn); err != nil {
		return "", err
	}
	return dsn, nil
}

func connectToStorageDB(dsn string) (*redis.Storage, error) {
	store := redis.New(redis.Config{
		URL:   dsn,
		Reset: true,
	})

	if _, err := store.Conn().Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	if err := store.Set("hc_test", []byte("test"), 1*time.Minute); err != nil {
		return nil, err
	}

	return store, nil
}

func closeStorageDBConnection() {
	if DBSTORAGE == nil {
		return
	}

	if err := DBSTORAGE.Close(); err != nil {
		log.Println("Error closing Redis connection:", err)
	}
	DBSTORAGE = nil
}
