package cache

import (
	DB "SMART_OFFICE_APP/pkg/db"
	mdl "SMART_OFFICE_APP/pkg/db/model"
	"encoding/json"
)

type APPCONFIG_CACHE struct {
}

// Init function is used to initialize the cache
func (m *APPCONFIG_CACHE) Name() string {
	return "APP_CONFIG"
}

func (m *APPCONFIG_CACHE) Refresh() {
	DATA := []mdl.APPCONFIG{}
	if err := DB.Q.APPCONFIG.DO.Scan(&DATA); err == nil {
		if buf, err := json.Marshal(DATA); err == nil {
			_ = DB.DBSTORAGE.Set("DBCACHE"+"APPCONFIG", buf, 0) == nil
		}
	}
}

func (m *APPCONFIG_CACHE) Get() []mdl.APPCONFIG {
	DATA := []mdl.APPCONFIG{}
	if buf, err := DB.DBSTORAGE.Get("DBCACHE" + "APPCONFIG"); err == nil {
		if err := json.Unmarshal(buf, &DATA); err == nil {
			return DATA
		}
	} else {
		DB.Q.APPCONFIG.DO.Scan(&DATA)
	}
	return DATA
}

func (m *APPCONFIG_CACHE) Filter(f func(mdl.APPCONFIG) bool) []mdl.APPCONFIG {
	DATA := m.Get()
	var result []mdl.APPCONFIG
	for i := range DATA {
		if f(DATA[i]) {
			result = append(result, DATA[i])
		}
	}
	return result
}

func (m *APPCONFIG_CACHE) Find(f func(mdl.APPCONFIG) bool) *mdl.APPCONFIG {
	DATA := m.Get()
	for i := range DATA {
		if f(DATA[i]) {
			return &DATA[i]
		}
	}
	return nil
}
