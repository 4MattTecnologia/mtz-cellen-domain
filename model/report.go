package model

type Report struct {
    metadata        map[string]interface{}
    data            interface{}
}

func NewReport(metadata map[string]interface{},
               data interface{}) (Report, error) {
    return Report{metadata, data}, nil
}
