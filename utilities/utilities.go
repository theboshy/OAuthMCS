package utilities

import (
	"encoding/json"
	"../models"
	"os"
	"database/sql"
	"strconv"
	//"os/exec"
	"github.com/google/uuid"
	"time"
)

func GetConfiguration() (models.Configuration, error) {
	config := models.Configuration{}
	file, err := os.Open("./configuration.json")
	if err != nil {
		return config, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}


func ToNullInt64(s string) sql.NullInt64 {
	if s == "0" {
	    s ="null"
	}
	i, err := strconv.Atoi(s)

	return sql.NullInt64{Int64 : int64(i), Valid : err == nil}
}

func GetUUID() uuid.UUID {
	uuid,err :=uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return uuid
}

func UnixNow() int64 {
	now := time.Now()
	return now.Unix()
}

