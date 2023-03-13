package toolmodel
import (
    "fmt"
)

type Profile struct {
    id          int
    name        string
    security    map[string]bool
}

func NewProfile(id int, name string,
                security map[string]bool) (
                    Profile, error) {
    if name == "" {
        return Profile{}, fmt.Errorf("Invalid empty name for Profile")
    }
    if len(security) == 0 {
        return Profile{}, fmt.Errorf(
            "Invalid empty security options for Profile")
    }
    return Profile{id, name, security}, nil
}

func (p *Profile) GetId() int {
    return p.id
}

func (p *Profile) GetName() string {
    return p.name
}

func (p *Profile) GetSecurity() map[string]bool {
    return p.security
}
