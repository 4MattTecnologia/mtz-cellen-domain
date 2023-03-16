package repositories
import (
    "fmt"
    "github.com/4MattTecnologia/mtz-cellen-domain/model"
)

type AbsOriginInstanceRepository interface {
    Get(filters ...map[string]interface{}) ([]model.OriginInstance, error)
    Insert(oInstance model.OriginInstance) error
    Remove(id int) error
}

func NewOriginInstance(name string, connValues model.ConnectionValues,
                       domainId int, originId int,
                       oInstanceRepo AbsOriginInstanceRepository,
                       domainRepo AbsDomainRepository,
                       originRepo AbsOriginRepository) (model.OriginInstance, error) {

//    domain, err := domainRepo.Get(domainId)
//    if err != nil {
//        return model.OriginInstance{}, fmt.Errorf(
//            "Error querying for domain while creating origin instance")
//    }
//
//    origin, err := originRepo.Get(originId)
//    if err != nil {
//        return model.OriginInstance{}, fmt.Errorf(
//            "Error querying for origin type while creating origin instance")
//    }

    oInstances, err := oInstanceRepo.GetAll()
    if err != nil {
        return model.OriginInstance{}, fmt.Errorf(
            "Error attempting to query origins in NewOriginInstance function: %v", err)
    }
    // verificar se par nome+nome domínio configuram chave primária
    maxId := 0
    for _, v := range oInstances {
        if maxId <= v.GetId() {
            maxId = v.GetId() + 1
        }
    }

    return model.NewOriginInstance(maxId, name, originId,
                                   domainId, connValues)
}

type FakeOriginInstanceRepository struct {
    oInstances []model.OriginInstance
}

func (f *FakeOriginInstanceRepository) GetAll(filters ...map[string]interface{}) ([]model.OriginInstance, error) {
    return f.oInstances, nil
}

func (f *FakeOriginInstanceRepository) Insert(o model.OriginInstance) error {
    f.oInstances = append(f.oInstances, o)
    return nil
}
func (f *FakeOriginInstanceRepository) Remove(id int) error {
    size := len(f.oInstances)
    newOriginInstances := make([]model.OriginInstance, 0, size)
    for _, v := range f.oInstances {
        if v.GetId() != id {
            newOriginInstances = append(newOriginInstances, v)
        }
    }
    f.oInstances = newOriginInstances
    return nil
}
