package model
import (
    "fmt"
)

type ConnectionValues map[string]interface{}
//type ExecutionSettings map[string]interface{}

type OriginInstance struct {
    id                  int
    name                string
    originId            int
    domainId            int
    connectionValues    ConnectionValues
    // Optional:
//    executionSettings   ExecutionSettings
}

func (o *OriginInstance) GetId() int {
    return o.id
}

func (o *OriginInstance) GetName() string {
    return o.name
}

func (o *OriginInstance) GetOriginId() int {
    return o.originId
}

func (o *OriginInstance) GetDomainId() int {
    return o.domainId
}

func (o *OriginInstance) GetConnectionValues() ConnectionValues {
    return o.connectionValues
}

func present(key string, list []string) bool {
    for _, v := range list {
        if v == key {
            return true
        }
    }
    return false
}

func AssertConnectionInfo(connValues ConnectionValues,
                          baseOrigin Origin) bool {
    valKeys := []string{}
    for k := range connValues {
        valKeys = append(valKeys, k)
    }
    for k := range baseOrigin.connectionInfo {
        if !present(k, valKeys) {
            return false
        }
    }
    return true
}

func NewOriginInstance(id int, name string,
                       originId int,
                       domainId int,
                       connV ConnectionValues) (OriginInstance, error) {
    if name == "" {
        return OriginInstance{}, fmt.Errorf(
            "Invalid empty name for origin instance")
    }
//    if !assertConnectionInfo(connV, origin) {
//        return OriginInstance{}, fmt.Errorf(
//            "Error creating origin instance: wrong format for connection values")
//    }
    return OriginInstance{id, name, originId, domainId, connV}, nil
}
