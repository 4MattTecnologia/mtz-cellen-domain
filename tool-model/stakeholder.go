package toolmodel
import (
    "errors"
)

type Stakeholder struct {
    id      int
    name    string
    dAgreements map[int]int // maps domainIds to agreementIds
}

func NewStakeholder (id int, name string) (Stakeholder, error) {
    if name == "" {
        return Stakeholder{}, errors.New("Invalid empty name for Stakeholder")
    }
    return Stakeholder{id, name, map[int]int{}}, nil
}

func  SetAgreementToDomain(stakeholder Stakeholder,
                         domainId int,
                         agreementId int) (Stakeholder, error) {
    _, ok := stakeholder.dAgreements[domainId]
    if !ok {
        return stakeholder, errors.New(
            "Error in AddDomainAgreement: could not find domainId")
    }
    newStakeholder := stakeholder
    newStakeholder.dAgreements[domainId] = agreementId
    return newStakeholder, nil
}

func RemoveDomain(stakeholder Stakeholder, domainId int) (
        Stakeholder, error) {
    _, ok := stakeholder.dAgreements[domainId]
    if !ok {
        return stakeholder, errors.New(
            "Error in RemoveDomain: could not find domainId")
    }
    newStakeholder := stakeholder
    delete(newStakeholder.dAgreements, domainId)
    return newStakeholder, nil
}

func AddDomain(stakeholder Stakeholder,
               domainId int) Stakeholder {
    _, ok := stakeholder.dAgreements[domainId]
    if ok {
        return stakeholder
    }
    newStakeholder := stakeholder
    newStakeholder.dAgreements[domainId] = -1
    return newStakeholder
}

func (s *Stakeholder) GetId() int {
    return s.id
}

func (s *Stakeholder) GetName() string {
    return s.name
}

func (s *Stakeholder) GetAgreementFromDomain(domainId int) (int, error) {
    val, ok := s.dAgreements[domainId]
    if !ok {
        return 0, errors.New("Error in GetAgreementFromDomain: "+
                             "domainId not found")
    }
    return val, nil
}
