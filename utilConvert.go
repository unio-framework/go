package unio

import (
	"encoding/json"
)

func (u *Util) InterfaceToStruct(m interface{}, val interface{}) error {
	tmp, err := json.Marshal(m); if err != nil {
		u.TraceError(err)
		return err
	}
	err = json.Unmarshal(tmp, val); if err != nil {
		u.TraceError(err)
		return err
	}
	return nil
}
