package repositories
import (
    "fmt"
    "github.com/4MattTecnologia/mtz-cellen-domain/model"
)

type AbsDomainRepository interface {
    GetAll() ([]model.Domain, error)
    Get(id int) (model.Domain, error)
    Insert(domain model.Domain) error
    Remove(id int) error
}

func NewDomain(name string, repo AbsDomainRepository) (model.Domain, error) {
    domains, err := repo.GetAll()
    if err != nil {
        return model.Domain{}, fmt.Errorf("Error in GetAll() query")
    }
    maxId := 0
    for _, v := range domains {
        if maxId <= v.GetId() {
            maxId = v.GetId() + 1
        }
    }
    return model.NewDomain(maxId, name)
}

type FakeDomainRepository struct {
    domains []model.Domain
}

func (f *FakeDomainRepository) GetAll() ([]model.Domain, error) {
    return f.domains, nil
}
func (f *FakeDomainRepository) Get(id int) (model.Domain, error) {
    for _, v := range f.domains {
        if v.GetId() == id {
            return v, nil
        }
    }
    return model.Domain{}, fmt.Errorf("No domain found for id %v", id)
}
func (f *FakeDomainRepository) Insert(d model.Domain) error {
    f.domains = append(f.domains, d)
    return nil
}
func (f *FakeDomainRepository) Remove(id int) error {
    size := len(f.domains)
    newDomains := make([]model.Domain, 0, size)
    for _, v := range f.domains {
        if v.GetId() != id {
            newDomains = append(newDomains, v)
        }
    }
    f.domains = newDomains
    return nil
}
