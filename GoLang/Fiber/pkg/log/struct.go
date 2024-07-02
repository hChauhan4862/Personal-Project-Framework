package MODEL_LOG

import "time"

type LOG struct {
	REQUESTID      string  `json:"REQUEST_ID"`
	TIMESTAMP      string  `json:"TIMESTAMP"`
	USERID         string  `json:"USER_ID"`
	USERROLE       string  `json:"USER_ROLE"`
	REQUEST_METHOD string  `json:"REQUEST_METHOD"`
	REQUEST_URI    string  `json:"REQUEST_URI"`
	REQUEST_BODY   string  `json:"REQUEST_BODY"`
	RESPONSE_BODY  string  `json:"RESPONSE_BODY"`
	STATUS_CODE    int     `json:"STATUS_CODE"`
	EXECUTION_TIME float64 `json:"EXECUTION_TIME"`
	IP_ADDRESS     string  `json:"IP_ADDRESS"`
	DEVICE_ID      string  `json:"DEVICE_ID"`
	USER_AGENT     string  `json:"USER_AGENT"`
}

func (l *LOG) Insert() error {
	if !myConn.STATUS {
		return nil
	}
	_, err := myConn.Query(`
		INSERT INTO smart_office_logs
			(REQUEST_ID, TIMESTAMP, USER_ID, USER_ROLE, REQUEST_METHOD, REQUEST_URI, REQUEST_BODY, RESPONSE_BODY, STATUS_CODE, EXECUTION_TIME, IP_ADDRESS, DEVICE_ID, USER_AGENT)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`, l.REQUESTID, time.Now(), l.USERID, l.USERROLE, l.REQUEST_METHOD, l.REQUEST_URI, l.REQUEST_BODY, l.RESPONSE_BODY, l.STATUS_CODE, l.EXECUTION_TIME, l.IP_ADDRESS, l.DEVICE_ID, l.USER_AGENT)

	return err
}
