package toolrepositories
import (
    "github.com/4MattTecnologia/mtz-cellen-domain/tool-model"
    "fmt"
)

type AbsStakeholderRepository interface {
    GetAll() ([]toolmodel.Stakeholder, error)
    Get(id int) (toolmodel.Stakeholder, error)
    Insert(stakeholder toolmodel.Stakeholder) error
    Remove(id int) error
}


func NewStakeholder(name string, domainIds []int,
                    repo AbsStakeholderRepository) (
                    toolmodel.Stakeholder, error) {
    if name == "" {
        return toolmodel.Stakeholder{}, fmt.Errorf("Invalid empty name for stakeholder")
    }
    stakeholders, err := repo.GetAll()
    if err != nil {
        return toolmodel.Stakeholder{}, fmt.Errorf("Error in GetAll() query")
    }
    maxId := 0
    for _, v := range stakeholders {
        if maxId <= v.GetId() {
            maxId = v.GetId() + 1
        }
    }
    return toolmodel.NewStakeholder(maxId, name, domainIds)
}

type FakeStakeholderRepository struct {
    stakeholders []toolmodel.Stakeholder
}

func (f *FakeStakeholderRepository) GetAll() ([]toolmodel.Stakeholder, error) {
    return f.stakeholders, nil
}

func (f *FakeStakeholderRepository) Get(id int) (toolmodel.Stakeholder, error) {
    for _, v := range f.stakeholders {
        if v.GetId() == id {
            return v, nil
        }
    }
    return toolmodel.Stakeholder{}, fmt.Errorf("No stakeholder for id %v", id)
}
func (f *FakeStakeholderRepository) Update(
        id int, stakeholder toolmodel.Stakeholder) error {
    size := len(f.stakeholders)
    newStholders := make([]toolmodel.Stakeholder, 0, size)
    for _, v := range f.stakeholders {
        if v.GetId() != id {
            newStholders = append(newStholders, v)
        } else {
            newStholders = append(newStholders, stakeholder)
        }
    }
    f.stakeholders = newStholders
    return nil
}
func (f *FakeStakeholderRepository) Insert(
        stakeholder toolmodel.Stakeholder) error {
    f.stakeholders = append(f.stakeholders, stakeholder)
    return nil
}
func (f *FakeStakeholderRepository) Remove(id int) error {
    size := len(f.stakeholders)
    newStholders := make([]toolmodel.Stakeholder, 0, size)
    for _, v := range f.stakeholders {
        if v.GetId() != id {
            newStholders = append(newStholders, v)
        }
    }
    f.stakeholders = newStholders
    return nil
}
