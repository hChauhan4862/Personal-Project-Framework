package MODEL_LOG

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

type CrateResponse struct {
	Cols     []interface{} `json:"cols"`
	Duration float64       `json:"duration"`
	Rowcount int           `json:"rowcount"`
	Rows     []interface{} `json:"rows"`
}

type Conn struct {
	HOST     string
	PORT     uint16
	USERNAME string
	PASSWORD string
	STATUS   bool
}

var myConn Conn

func (c *Conn) Connect() error {
	_, err := c.Query("SELECT 1")
	if err == nil {
		c.STATUS = true
		myConn = *c
	}
	return err
}

func (c *Conn) Query(QUERY string, args ...interface{}) (CrateResponse, error) {
	return queryCrateDB(c.URL(), c.USERNAME, c.PASSWORD, QUERY, args...)
}

func (c *Conn) URL() string {
	return "https://" + c.HOST + ":" + fmt.Sprint(c.PORT) + "/_sql"
}

func queryCrateDB(URL, USERNAME, PASSWORD, QUERY string, args ...interface{}) (CrateResponse, error) {
	type crateReq struct {
		Stmt string        `json:"stmt"`
		Args []interface{} `json:"args"`
	}

	type crateReq2 struct {
		Stmt string `json:"stmt"`
	}

	var response CrateResponse
	var req crateReq
	var req2 crateReq2

	req.Stmt = QUERY
	req.Args = args

	req2.Stmt = QUERY

	var buf bytes.Buffer
	// if Args is empty, then it will be removed from the buffer
	var err error
	if len(args) == 0 {
		err = json.NewEncoder(&buf).Encode(req2)
	} else {
		err = json.NewEncoder(&buf).Encode(req)
	}
	if err != nil {
		return response, err
	}

	client := http.Client{}
	r, err := http.NewRequest("POST", URL, &buf)
	if err != nil {
		return response, err
	}
	r.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(USERNAME+":"+PASSWORD)))
	r.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(r)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}
