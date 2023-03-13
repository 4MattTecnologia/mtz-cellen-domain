package toolmodel
import (
    "fmt"
)

type Module struct {
    id      int
    name    string
}

func NewModule(id int, name string) (Module, error) {
    if name == "" {
        return Module{}, fmt.Errorf("Invalid empty name for Module")
    }
    return Module{id, name}, nil
}

func (m *Module) GetId() int {
    return m.id
}

func (m *Module) GetName() string {
    return m.name
}
