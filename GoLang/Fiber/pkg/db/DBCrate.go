package DB

import (
	"github.com/jackc/pgx/v5"

	MODEL_LOG "SMART_OFFICE_APP/pkg/log"
)

func connectToLogDB(dsn string) (MODEL_LOG.Conn, error) {

	conn := MODEL_LOG.Conn{}
	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		return conn, err
	}

	conn.HOST = config.Host
	conn.PORT = config.Port
	conn.USERNAME = config.User
	conn.PASSWORD = config.Password
	if err := conn.Connect(); err != nil {
		return conn, err
	}

	if _, err := conn.Query(`
		CREATE TABLE IF NOT EXISTS SMART_OFFICE_LOGS (
			REQUEST_ID VARCHAR(255),
			TIMESTAMP TIMESTAMP WITH TIME ZONE,
			USER_ID VARCHAR(255),
			USER_ROLE VARCHAR(255),
			REQUEST_METHOD VARCHAR(255),
			REQUEST_URI VARCHAR(255),
			REQUEST_BODY TEXT,
			RESPONSE_BODY TEXT,
			STATUS_CODE INT,
			EXECUTION_TIME FLOAT,
			IP_ADDRESS VARCHAR(255),
			DEVICE_ID VARCHAR(255),
			USER_AGENT VARCHAR(255)
		)
	`); err != nil {
		return conn, err
	}

	return conn, nil
}

func getLogDBCredential() (string, error) {
	var dsn string
	if err := Q.APPCONFIG.DO.Where(Q.APPCONFIG.NAME.Eq("SERVER_CONFIG_CRATEDB")).Select(Q.APPCONFIG.VALUE).Scan(&dsn); err != nil {
		return "", err
	}
	return dsn, nil
}

func UpdateCrateDBConn(dsn string) error {
	if Q == nil {
		return nil
	}

	// Connect to CrateDB for logging
	if _, err := connectToLogDB(dsn); err != nil {
		return err
	}

	// Update CrateDB connection details in the main database
	if _, err := Q.APPCONFIG.DO.Where(Q.APPCONFIG.NAME.Eq("SERVER_CONFIG_CRATEDB")).Update(Q.APPCONFIG.VALUE, dsn); err != nil {
		return err
	}

	return nil
}
