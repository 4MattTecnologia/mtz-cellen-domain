package model
import (
    "fmt"
)

type Domain struct {
    id          int
    name        string
}

func (d *Domain) GetId() int {
    return d.id
}

func (d *Domain) GetName() string {
    return d.name
}

func NewDomain(id int, name string) (Domain, error) {
    if name == "" {
        return Domain{}, fmt.Errorf("Invalid empty name for domain")
    }
    return Domain{id, name}, nil
}
