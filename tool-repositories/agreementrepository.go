package toolrepositories
import (
    "fmt"
    "github.com/4MattTecnologia/mtz-cellen-domain/tool-model"
)
type AbsAgreementRepository interface {
    GetAll() ([]toolmodel.Agreement, error)
    Get(id int) (toolmodel.Agreement, error)
    Insert(agreement toolmodel.Agreement) error
    Remove(id int) error
}

func NewAgreement(name string, nUsers int,
                  nMUsers int, pgLimit int,
                  repo AbsAgreementRepository) (toolmodel.Agreement, error){
    if name == "" {
        return toolmodel.Agreement{}, fmt.Errorf("Invalid empty name for agreement")
    }
    agreements, err := repo.GetAll()
    if err != nil {
        return toolmodel.Agreement{}, fmt.Errorf("Error in GetAll() query")
    }
    maxId := 0
    for _, v := range agreements {
        if maxId <= v.GetId() {
            maxId = v.GetId() + 1
        }
    }
    newAgreement, err := toolmodel.NewAgreement(maxId, name,
                                                nUsers, nMUsers,
                                                pgLimit)
    if err != nil {
        return toolmodel.Agreement{}, nil
    }
    return newAgreement, nil
}

type FakeAgreementRepository struct {
    agreements []toolmodel.Agreement
}

func (f *FakeAgreementRepository) GetAll() ([]toolmodel.Agreement, error) {
    return f.agreements, nil
}

func (f *FakeAgreementRepository) Get(id int) (toolmodel.Agreement, error) {
    for _, v := range f.agreements {
        if v.GetId() == id {
            return v, nil
        }
    }
    return toolmodel.Agreement{}, fmt.Errorf("No stakeholder for id %v", id)
}

func (f *FakeAgreementRepository) Insert(
        agreement toolmodel.Agreement) error {
    f.agreements = append(f.agreements, agreement)
    return nil
}

func (f *FakeAgreementRepository) Remove(id int) error {
    size := len(f.agreements)
    newAgr := make([]toolmodel.Agreement, 0, size)
    for _, v := range f.agreements {
        if v.GetId() != id {
            newAgr = append(newAgr, v)
        }
    }
    f.agreements = newAgr
    return nil
}
