package pgdb

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

// JSONValue is used to implement Value method for structs which will be represented in database as raw json
func JSONValue(data interface{}) (driver.Value, error) {
	data, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal data", logan.F{
			"data_type": fmt.Sprintf("%T", data),
		})
	}

	return data, nil
}

// JSONScan is used to implement Scan method for structs which will be represented in database as raw json
func JSONScan(src, dest interface{}) error {
	var data []byte
	switch rawData := src.(type) {
	case []byte:
		data = rawData
	case string:
		data = []byte(rawData)
	case nil:
		data = []byte("null")
	default:
		return errors.From(errors.New("unexpected raw type"), logan.F{
			"src_type":  fmt.Sprintf("%T", src),
			"dest_type": fmt.Sprintf("%T", dest),
		})
	}

	err := json.Unmarshal(data, dest)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal data", logan.F{
			"dest_type": fmt.Sprintf("%T", dest),
			"raw":       string(data),
		})
	}

	return nil
}
