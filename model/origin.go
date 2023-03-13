package model
import (
    "fmt"
)

type Origin struct {
    id              int
    name            string
//    version         int
//    releaseDate     string
    connectionInfo  map[string]string
}

func (o *Origin) GetId() int {
    return o.id
}
func (o *Origin) GetName() string {
    return o.name
}
func (o *Origin) GetConnectionInfo() map[string]string {
    return o.connectionInfo
}

func NewOrigin(id int, name string,
               connInfo map[string]string) (Origin, error) {
    if name == "" {
        return Origin{}, fmt.Errorf("Invalid empty name for origin")
    }
    return Origin{id, name, connInfo}, nil
}
