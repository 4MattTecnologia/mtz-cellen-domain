package toolrepositories
import (
    "fmt"
    "github.com/4MattTecnologia/mtz-cellen-domain/tool-model"
)
type AbsProfileRepository interface {
    GetAll() ([]toolmodel.Profile, error)
    Get(id int) (toolmodel.Profile, error)
    Insert(profile toolmodel.Profile) error
    Remove(id int) error
}

func NewProfile(name string, security map[string]bool,
                  repo AbsProfileRepository) (toolmodel.Profile, error){
    if name == "" {
        return toolmodel.Profile{}, fmt.Errorf("Invalid empty name for profile")
    }
    profiles, err := repo.GetAll()
    if err != nil {
        return toolmodel.Profile{}, fmt.Errorf("Error in GetAll() query")
    }
    maxId := 0
    for _, v := range profiles {
        if maxId <= v.GetId() {
            maxId = v.GetId() + 1
        }
    }
    newProfile, err := toolmodel.NewProfile(maxId, name,
                                            security)
    if err != nil {
        return toolmodel.Profile{}, nil
    }
    return newProfile, nil
}

type FakeProfileRepository struct {
    profiles []toolmodel.Profile
}

func (f *FakeProfileRepository) GetAll() ([]toolmodel.Profile, error) {
    return f.profiles, nil
}

func (f *FakeProfileRepository) Get(id int) (toolmodel.Profile, error) {
    for _, v := range f.profiles {
        if v.GetId() == id {
            return v, nil
        }
    }
    return toolmodel.Profile{}, fmt.Errorf("No stakeholder for id %v", id)
}

func (f *FakeProfileRepository) Insert(
        profile toolmodel.Profile) error {
    f.profiles = append(f.profiles, profile)
    return nil
}

func (f *FakeProfileRepository) Remove(id int) error {
    size := len(f.profiles)
    newProfile := make([]toolmodel.Profile, 0, size)
    for _, v := range f.profiles {
        if v.GetId() != id {
            newProfile = append(newProfile, v)
        }
    }
    f.profiles = newProfile
    return nil
}
