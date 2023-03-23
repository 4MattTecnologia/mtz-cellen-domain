package toolrepositories
import (
    "fmt"
    "github.com/4MattTecnologia/mtz-cellen-domain/tool-model"
)
type AbsProfileRepository interface {
    Get(filters ...map[string]interface{}) ([]toolmodel.Profile, error)
    Insert(profile toolmodel.Profile) error
    Remove(id int) error
}

func NewProfile(name string, security map[string]bool,
                  repo AbsProfileRepository) (toolmodel.Profile, error){
    if name == "" {
        return toolmodel.Profile{}, fmt.Errorf("Invalid empty name for profile")
    }
    profiles, err := repo.Get()
    if err != nil {
        return toolmodel.Profile{}, fmt.Errorf("Error in Get() query: %v", err)
    }
    maxId := 0
    for _, v := range profiles {
        if maxId <= v.GetId() {
            maxId = v.GetId() + 1
        }
    }
    return toolmodel.NewProfile(maxId, name,
                                security)
}

type FakeProfileRepository struct {
    profiles []toolmodel.Profile
}

func (f *FakeProfileRepository) Get(
        filters ...map[string]interface{}) ([]toolmodel.Profile, error) {
    return f.profiles, nil
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
