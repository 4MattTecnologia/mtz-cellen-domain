package toolmodel
import (
    "fmt"
)

type MtzUser struct {
    id              int
    name            string
    password        string
    domainId        int
    stakeholderId   int
    profileId       int
    startDate       string
    endDate         string
    publicKey       []byte
    privateKey      []byte
}

func NewMtzUser(id int, name string,
                password string, domainId int,
                stakeholderId int, profileId int,
                startDate string, endDate string,
                publicKey []byte, privateKey []byte) (
                    MtzUser, error) {
    if name == "" {
        return MtzUser{}, fmt.Errorf("Invalid empty name for MtzUser")
    }
//    if len(publicKey) == 0 {
//        return MtzUser{}, fmt.Errorf("Invalid empty public key")
//    }
//    if len(privateKey) == 0 {
//        return MtzUser{}, fmt.Errorf("Invalid empty private key")
//    }
    return MtzUser{
        id, name,
        password, domainId,
        stakeholderId, profileId,
        startDate, endDate,
        publicKey, privateKey,
    }, nil
}

func (m *MtzUser) GetId() int {
    return m.id
}
func (m *MtzUser) GetName() string {
    return m.name
}
func (m *MtzUser) GetPassword() string {
    return m.password
}
func (m *MtzUser) GetDomainId() int {
    return m.domainId
}
func (m *MtzUser) GetStakeholderId() int {
    return m.stakeholderId
}
func (m *MtzUser) GetProfileId() int {
    return m.profileId
}
func (m *MtzUser) GetStartDate() string {
    return m.startDate
}
func (m *MtzUser) GetEndDate() string {
    return m.endDate
}
func (m *MtzUser) GetPublicKey() []byte {
    return m.publicKey
}