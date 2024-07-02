package DB

import (
	"log"

	"github.com/gofiber/storage/redis/v3"
	"gorm.io/gorm"

	QU "SMART_OFFICE_APP/pkg/db/query"
	MODEL_LOG "SMART_OFFICE_APP/pkg/log"
)

var (
	Q         *QU.Query
	DBMAIN    *gorm.DB
	DBLOG     MODEL_LOG.Conn
	DBSTORAGE *redis.Storage
)

func Init() {
	if err := ConnectToDatabase(); err != nil {
		log.Println("ERROR IN CONNECTING DB :: ", err)
	} else {
		log.Println("Connected to the database")
	}
}

func ConnectToDatabase() error {
	var (
		DSN_MAIN  string
		DSN_LOG   string
		DSN_REDIS string
		err       error
	)
	// MAIN DATABASE
	if DSN_MAIN, err = getMainDBCredential(); err != nil {
		return err
	}
	if TEMPDB, err := connectToMainDB(DSN_MAIN); err != nil {
		return err
	} else {
		closeMainDBConnection()
		DBMAIN = TEMPDB
		Q = QU.Use(DBMAIN)
	}

	// LOG DATABASE
	if DSN_LOG, err = getLogDBCredential(); err != nil {
		return err
	}

	if TEMPDB, err := connectToLogDB(DSN_LOG); err != nil {
		return err
	} else {
		DBLOG = TEMPDB
	}

	// REDIS DATABASE
	if DSN_REDIS, err = getStorageDBCredential(); err != nil {
		return err
	}
	if TEMPDB, err := connectToStorageDB(DSN_REDIS); err != nil {
		return err
	} else {
		closeStorageDBConnection()
		DBSTORAGE = TEMPDB
	}

	return nil
}
