package cache

var APP_CONFIG = APPCONFIG_CACHE{}

func Init() {
	APP_CONFIG.Refresh()
}
