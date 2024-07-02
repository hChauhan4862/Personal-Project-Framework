package DB

import (
	MDL "SMART_OFFICE_APP/pkg/db/model"
	"encoding/json"
	"net/http"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func connectToMainDB(dsn string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	if db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{}); err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(500)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := migrateMainDBModels(db); err != nil {
		_ = " "
		// return nil, err
	}

	return db, nil
}

func getMainDBCredential() (string, error) {
	type DatabaseSecrets struct {
		Raw                string `json:"raw"`
		Computed           string `json:"computed"`
		Note               string `json:"note"`
		RawVisibility      string `json:"rawVisibility"`
		ComputedVisibility string `json:"computedVisibility"`
	}

	type SecretsResponse struct {
		Secrets map[string]DatabaseSecrets `json:"secrets"`
		Success bool                       `json:"success"`
	}

	SecretToken := "_"
	URL := "https://api.doppler.com/v3/configs/config/secrets?project=smart_office&config=prd&include_dynamic_secrets=false&secrets=DATABASE&include_managed_secrets=false"

	// Get the database credentials from Doppler
	client := http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "Bearer "+SecretToken)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Parse the response
	var secretsResponse SecretsResponse
	if err := json.NewDecoder(resp.Body).Decode(&secretsResponse); err != nil {
		return "", err
	}

	// Extract the database credentials
	dsn := secretsResponse.Secrets["DATABASE"].Raw

	return dsn, nil
}

func migrateMainDBModels(db *gorm.DB) error {
	return db.AutoMigrate(
		&MDL.APPCONFIG{},
	)
}

func closeMainDBConnection() {
	if DBMAIN == nil {
		return
	}
	oldDB, err := DBMAIN.DB()
	if err == nil {
		defer oldDB.Close()
	}
	DBMAIN = nil
}
