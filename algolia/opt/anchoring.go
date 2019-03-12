// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type AnchoringOption struct {
	value string
}

func Anchoring(v string) *AnchoringOption {
	return &AnchoringOption{v}
}

func (o AnchoringOption) Get() string {
	return o.value
}

func (o AnchoringOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *AnchoringOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = ""
		return nil
	}
	return json.Unmarshal(data, &o.value)
}