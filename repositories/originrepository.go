package repositories
import (
    "fmt"
    "github.com/4MattTecnologia/mtz-cellen-domain/model"
)

type AbsOriginRepository interface {
    Get(filters ...map[string]interface{}) ([]model.Origin, error)
    Insert(origin model.Origin) error
    Remove(id int) error
}

func NewOrigin(name string, connInfo map[string]string,
               repo AbsOriginRepository) (model.Origin, error) {
    origins, err := repo.GetAll()
    if err != nil {
        return model.Origin{}, fmt.Errorf("Error attempting to query origins in NewOrigin function: %v", err)
    }
    maxId := 0
    for _, v := range origins {
        if maxId <= v.GetId() {
            maxId = v.GetId() + 1
        }
    }

    return model.NewOrigin(maxId, name, connInfo)
}

type FakeOriginRepository struct {
    origins []model.Origin
}

func (f *FakeOriginRepository) Get(
        filters ...map[string]interface{}) ([]model.Origin, error) {
    return f.origins, nil
}

func (f *FakeOriginRepository) Insert(o model.Origin) error {
    f.origins = append(f.origins, o)
    return nil
}
func (f *FakeOriginRepository) Remove(id int) error {
    size := len(f.origins)
    newOrigins := make([]model.Origin, 0, size)
    for _, v := range f.origins {
        if v.GetId() != id {
            newOrigins = append(newOrigins, v)
        }
    }
    f.origins = newOrigins
    return nil
}
