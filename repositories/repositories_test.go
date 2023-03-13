package repositories
import (
    "testing"
    "mattzero.com.br/domain/model"
)

func TestNewDomain(t *testing.T) {
    repo := &(FakeDomainRepository{})
    first, _ := NewDomain("first", repo)
    repo.Insert(first)
    second, _ := NewDomain("second", repo)
    repo.Insert(second)
    if second.GetId() != 1 {
        t.Fatalf("TestNewDomain failed: second.GetId() value was set to %v",
                 second.GetId())
    }
    query, _ := repo.GetAll()
    if len(query) > 2 {
        t.Fatalf("TestNewDomain failed: expected 2 items in repository")
    }
}

func TestNewOrigin(t *testing.T) {
    repo := &(FakeOriginRepository{})
    connInfo := map[string]string{
        "key": "value",
    }
    first, _ := NewOrigin("first", connInfo, repo)
    repo.Insert(first)
    second, _ := NewOrigin("second", connInfo, repo)
    repo.Insert(second)
    if second.GetId() != 1 {
        t.Fatalf("TestNewOrigin failed: second.GetId() value was set to %v",
                 second.GetId())
    }
    query, _ := repo.GetAll()
    if len(query) > 2 {
        t.Fatalf("TestNewOrigin failed: expected 2 items in repository")
    }
}

func TestNewOriginInstance(t *testing.T) {
    domainRepo := &(FakeDomainRepository{})
    baseDomain, _ := NewDomain("first", domainRepo)
    domainRepo.Insert(baseDomain)

    originRepo := &(FakeOriginRepository{})
    connInfo := map[string]string{
        "key": "string",
    }
    baseOrigin, _ := NewOrigin("first", connInfo, originRepo)
    originRepo.Insert(baseOrigin)

    repo := &(FakeOriginInstanceRepository{})
    connValues := model.ConnectionValues{"key": "somekey"}
    first, _ := NewOriginInstance("first", connValues,
                                  0, 0, repo,
                                  domainRepo,
                                  originRepo)
    repo.Insert(first)
    second, _ := NewOriginInstance("second", connValues,
                                   0, 0, repo,
                                   domainRepo,
                                   originRepo)
    repo.Insert(second)
    if second.GetId() != 1 {
        t.Fatalf("TestNewOrigin failed: second.GetId() value was set to %v",
                 second.GetId())
    }
    query, _ := repo.GetAll()
    if len(query) > 2 {
        t.Fatalf("TestNewOrigin failed: expected 2 items in repository")
    }
}
