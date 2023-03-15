package toolmodel
import (
    "errors"
)

type Stakeholder struct {
    id          int
    name        string
    domainIds   []int
}

func NewStakeholder (id int, name string,
                     domainIds []int) (Stakeholder, error) {
    if name == "" {
        return Stakeholder{}, errors.New("Invalid empty name for Stakeholder")
    }
    return Stakeholder{id, name, domainIds}, nil
}

func RemoveDomain(stakeholder Stakeholder, domainId int) (
        Stakeholder, error) {
    domainIds := make([]int, 0)
    found := false
    for _, v := range stakeholder.domainIds {
        if v != domainId {
            domainIds = append(domainIds, v)
        }
        if v == domainId {
            found = true
        }
    }
    if !found {
        return stakeholder, errors.New(
            "Error in RemoveDomain: could not find domainId")
    }
    newStakeholder := stakeholder
    newStakeholder.domainIds = domainIds
    return newStakeholder, nil
}

func AddDomain(stakeholder Stakeholder,
               domainId int) (Stakeholder, error) {
    for _, v := range stakeholder.domainIds {
        if v == domainId {
            return stakeholder, errors.New(
                "Error in AddDomain: domain already exists")
        }
    }
    newStakeholder := stakeholder
    newStakeholder.domainIds = append(
        newStakeholder.domainIds, domainId)
    return newStakeholder, nil
}

func (s *Stakeholder) GetId() int {
    return s.id
}

func (s *Stakeholder) GetName() string {
    return s.name
}

func (s *Stakeholder) GetDomainIds() []int {
    return s.domainIds
}
