package toolmodel
import (
    "fmt"
)

type Agreement struct {
    id                  int
    name                string
    numMtzUsers         int
    numMonitoredUsers   int
    pageLimit           int
}

func NewAgreement(id int, name string, numMtzUsers int,
                  numMonitoredUsers int, pageLimit int) (
                    Agreement, error) {
    if name == "" {
        return Agreement{}, fmt.Errorf("Invalid empty name for Agreement")
    }
    if numMtzUsers <= 0 {
        return Agreement{}, fmt.Errorf("Invalid number of users in agreement: smaller than one")
    }
    return Agreement{
        id, name, numMtzUsers,
        numMonitoredUsers, pageLimit,
    }, nil
}

func (a *Agreement) GetId() int {
    return a.id
}

func (a *Agreement) GetName() string {
    return a.name
}

func (a *Agreement) GetNumMtzUsers() int {
    return a.numMtzUsers
}

func (a *Agreement) GetNumMonitoredUsers() int {
    return a.numMonitoredUsers
}

func (a *Agreement) GetPageLimit() int {
    return a.pageLimit
}
