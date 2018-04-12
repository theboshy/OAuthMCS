package custom

import (
	"database/sql"
	"strings"
	"encoding/json"
)

type JsonNullInt64 struct {
	sql.NullInt64
}

func DefineSystemValue(value int64) JsonNullInt64 {
return 	JsonNullInt64{sql.NullInt64{value,true}}
}

func (v JsonNullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	} else {
		return json.Marshal(nil)
	}
}


func (v *JsonNullInt64) UnmarshalJSON(data []byte) error {
	var intData *int64
	var dataConst = strings.Replace(string(data), "\"", "", -1)
	//fmt.Println(dataConst)
	if err := json.Unmarshal([]byte(dataConst), &intData); err != nil {
		return err
	}
	if intData != nil {
		v.Valid = true
		v.Int64 = *intData
	} else {
		v.Valid = false
	}
	return nil
}
