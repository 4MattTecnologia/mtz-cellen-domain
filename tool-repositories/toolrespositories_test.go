package toolrepositories
import (
    "testing"
)

func TestNewAgreement(t *testing.T) {
    repo := &(FakeAgreementRepository{})
    first, _ := NewAgreement("first", 1, 1, 1, repo)
    repo.Insert(first)
    second, _ := NewAgreement("second", 1, 1, 1, repo)
    repo.Insert(second)
    if second.GetId() != 1 {
        t.Fatalf("TestNewAgreement failed: second.GetId() value was set to %v",
                 second.GetId())
    }
    query, _ := repo.GetAll()
    if len(query) != 2 {
        t.Fatalf("TestNewAgreement failed: expected 2 items in repository")
    }
}

func TestNewStakeholder(t *testing.T) {
    repo := &(FakeStakeholderRepository{})
    first, _ := NewStakeholder("first", repo)
    repo.Insert(first)
    second, _ := NewStakeholder("second", repo)
    repo.Insert(second)
    if second.GetId() != 1 {
        t.Fatalf("TestNewStakeholder failed: second.GetId() value was set to %v",
                 second.GetId())
    }
    query, _ := repo.GetAll()
    if len(query) != 2 {
        t.Fatalf("TestNewStakeholder failed: expected 2 items in repository")
    }
}

func TestNewMtzUser(t *testing.T) {
    repo := &(FakeMtzUserRepository{})
    first, _ := NewMtzUser("first",
                           "pwd", 0,
                           0, 0,
                           "0/0/0", "0/0/0",
                           []byte{}, []byte{},
                           repo)
    repo.Insert(first)
    second, _ := NewMtzUser("second",
                            "pwd", 0,
                            0, 0,
                            "0/0/0", "0/0/0",
                            []byte{}, []byte{},
                            repo)
    repo.Insert(second)
    if second.GetId() != 1 {
        t.Fatalf("TestNewMtzUser failed: second.GetId() value was set to %v",
                 second.GetId())
    }
    query, _ := repo.GetAll()
    if len(query) != 2 {
        t.Fatalf("TestNewMtzUser failed: expected 2 items in repository")
    }
}

func TestNewProfile(t *testing.T) {
    repo := &(FakeProfileRepository{})
    first, _ := NewProfile("first", map[string]bool{"admin": true},
                           repo)
    repo.Insert(first)
    second, _ := NewProfile("second", map[string]bool{"admin": true},
                            repo)
    repo.Insert(second)
    if second.GetId() != 1 {
        t.Fatalf("TestNewProfile failed: second.GetId() value was set to %v",
                 second.GetId())
    }
    query, _ := repo.GetAll()
    if len(query) != 2 {
        t.Fatalf("TestNewProfile failed: expected 2 items in repository")
    }
}
