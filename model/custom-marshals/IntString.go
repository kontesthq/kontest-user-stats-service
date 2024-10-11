package custom_marshals

import (
	"encoding/json"
	"fmt"
)

type IntString int

// UnmarshalJSON for custom unmarshaling of IntString
func (is *IntString) UnmarshalJSON(data []byte) error {
	var strValue string
	if err := json.Unmarshal(data, &strValue); err != nil {
		return err
	}
	var intValue int
	if _, err := fmt.Sscanf(strValue, "%d", &intValue); err != nil {
		return err
	}
	*is = IntString(intValue)
	return nil
}
