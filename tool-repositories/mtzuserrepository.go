package toolrepositories
import (
    "mattzero.com.br/domain/tool-model"
    "fmt"
    "errors"
)

type AbsMtzUserRepository interface {
    GetAll() ([]toolmodel.MtzUser, error)
    Get(id int) (toolmodel.MtzUser, error)
    Insert(toolmodel.MtzUser) error
    Remove(id int) error
}

func NewMtzUser(name string,
                password string, domainId int,
                stakeholderId int, profileId int,
                startDate string, endDate string,
                publicKey []byte, privateKey []byte,
                repo AbsMtzUserRepository) (
                    toolmodel.MtzUser, error) {
    if name == "" {
        return toolmodel.MtzUser{}, errors.New(
            "Invalid empty name for mtzUser")
    }
    mtzUsers, err := repo.GetAll()
    if err != nil {
        return toolmodel.MtzUser{}, errors.New("Error in GetAll() query")
    }
    maxId := 0
    for _, v := range mtzUsers {
        if maxId <= v.GetId() {
            maxId = v.GetId() + 1
        }
    }
    return toolmodel.NewMtzUser(maxId, name,
                                password, domainId,
                                stakeholderId, profileId,
                                startDate, endDate,
                                publicKey, privateKey)
}

type FakeMtzUserRepository struct {
    mtzUsers []toolmodel.MtzUser
}

func (f *FakeMtzUserRepository) GetAll() ([]toolmodel.MtzUser, error) {
    return f.mtzUsers, nil
}

func (f *FakeMtzUserRepository) Get(id int) (toolmodel.MtzUser, error) {
    for _, v := range f.mtzUsers {
        if v.GetId() == id {
            return v, nil
        }
    }
    return toolmodel.MtzUser{}, fmt.Errorf("No stakeholder for id %v", id)
}

func (f *FakeMtzUserRepository) Insert(
        mtzUser toolmodel.MtzUser) error {
    f.mtzUsers = append(f.mtzUsers, mtzUser)
    return nil
}

func (f *FakeMtzUserRepository) Remove(id int) error {
    size := len(f.mtzUsers)
    newAgr := make([]toolmodel.MtzUser, 0, size)
    for _, v := range f.mtzUsers {
        if v.GetId() != id {
            newAgr = append(newAgr, v)
        }
    }
    f.mtzUsers = newAgr
    return nil
}
