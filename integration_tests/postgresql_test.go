package integrationtests
import (
    "testing"
//    "github.com/4MattTecnologia/mtz-cellen-domain/repositories"
    "github.com/4MattTecnologia/mtz-cellen-domain/tool-repositories"
    "os"
)

var (
    dbName = os.Getenv("DB_NAME")
    dbHost = os.Getenv("DB_HOST")
    dbPort = os.Getenv("DB_PORT")
    dbUser = os.Getenv("DB_USER")
    dbPass = os.Getenv("DB_PASS")
)

//func TestModelPSQLRepositories(t *testing.T) {
//    domainPsql, _ := repositories.NewPSQLDomainRepo(dbName,
//                                          dbHost,
//                                          dbPort,
//                                          dbUser,
//                                          dbPass)
//    firstDomain, err := repositories.NewDomain("first", domainPsql)
//    if err != nil {
//        t.Fatalf("TestPSQLDomainRepository failed: %v", err)
//    }
//    err = domainPsql.Insert(firstDomain)
//    if err != nil {
//        t.Fatalf("TestPSQLDomainRepository failed: %v", err)
//    }
//    secondDomain, err := repositories.NewDomain("second", domainPsql)
//    if err != nil {
//        t.Fatalf("TestPSQLDomainRepository failed: %v", err)
//    }
//    err = domainPsql.Insert(secondDomain)
//    if err != nil {
//        t.Fatalf("TestPSQLDomainRepository failed: %v", err)
//    }
//
//    if secondDomain.GetId() <= 0 {
//        t.Fatalf("TestPSQLDomainRepository failed: " +
//            "expected id higher than 0 for second element")
//    }
//    gotDomain, _ := domainPsql.Get(1)
//    if gotDomain.GetId() != 1 {
//        t.Fatalf("TestPSQLDomainRepository failed: " +
//            "expected to fetch element with id = 1")
//    }
//    err = domainPsql.Remove(0)
//    if err != nil {
//        t.Fatalf("TestPSQLDomainRepository failed: %v", err)
//    }
//    gotAllDomains, err := domainPsql.GetAll()
//    if len(gotAllDomains) != 1 {
//        t.Fatalf("TestPSQLDomainRepository failed: " +
//            "expected array of 1 element after removal")
//    }
//
//    originPsql, _ := repositories.NewPSQLOriginRepo(dbName,
//                                              dbHost,
//                                              dbPort,
//                                              dbUser,
//                                              dbPass)
//    connInfo := map[string]string{"appID": "int"}
//    firstOrigin, err := repositories.NewOrigin(
//        "first", connInfo, originPsql)
//    if err != nil {
//        t.Fatalf("TestPSQLOriginRepository failed: %v", err)
//    }
//    err = originPsql.Insert(firstOrigin)
//    if err != nil {
//        t.Fatalf("TestPSQLOriginRepository failed: %v", err)
//    }
//    secondOrigin, err := repositories.NewOrigin(
//        "second", connInfo, originPsql)
//    if err != nil {
//        t.Fatalf("TestPSQLOriginRepository failed: %v", err)
//    }
//    err = originPsql.Insert(secondOrigin)
//    if err != nil {
//        t.Fatalf("TestPSQLOriginRepository failed: %v", err)
//    }
//
//    if secondOrigin.GetId() <= 0 {
//        t.Fatalf("TestPSQLOriginRepository failed: " +
//            "expected id higher than 0 for second element")
//    }
//    gotOrigin, _ := originPsql.Get(1)
//    if gotOrigin.GetId() != 1 {
//        t.Fatalf("TestPSQLOriginRepository failed: " +
//            "expected to fetch element with id = 1")
//    }
//    err = originPsql.Remove(0)
//    if err != nil {
//        t.Fatalf("TestPSQLOriginRepository failed: %v", err)
//    }
//    gotAllOrigins, err := originPsql.GetAll()
//    if len(gotAllOrigins) != 1 {
//        t.Fatalf("TestPSQLOriginRepository failed: " +
//            "expected array of 1 element after removal")
//    }
//
//    oInstancePsql, _ := repositories.NewPSQLOriginInstanceRepo(dbName,
//                                                               dbHost,
//                                                               dbPort,
//                                                               dbUser,
//                                                               dbPass)
//    connVals := map[string]interface{}{"appID": 0}
//    firstOInstance, err := repositories.NewOriginInstance(
//        "first", connVals, 1, 1,
//        oInstancePsql, domainPsql, originPsql)
//    if err != nil {
//        t.Fatalf("TestPSQLOriginRepository failed: %v", err)
//    }
//    err = oInstancePsql.Insert(firstOInstance)
//    if err != nil {
//        t.Fatalf("TestPSQLOriginRepository failed: %v", err)
//    }
//    secondOInstance, err := repositories.NewOriginInstance(
//        "second", connVals, 1, 1,
//        oInstancePsql, domainPsql, originPsql)
//    if err != nil {
//        t.Fatalf("TestPSQLOriginRepository failed: %v", err)
//    }
//    err = oInstancePsql.Insert(secondOInstance)
//    if err != nil {
//        t.Fatalf("TestPSQLOriginRepository failed: %v", err)
//    }
//
//    if secondOInstance.GetId() <= 0 {
//        t.Fatalf("TestPSQLOriginRepository failed: " +
//            "expected id higher than 0 for second element")
//    }
//    gotOInstance, _ := oInstancePsql.Get(1)
//    if gotOInstance.GetId() != 1 {
//        t.Fatalf("TestPSQLOriginRepository failed: " +
//            "expected to fetch element with id = 1")
//    }
//    err = oInstancePsql.Remove(0)
//    if err != nil {
//        t.Fatalf("TestPSQLOriginRepository failed: %v", err)
//    }
//    gotAllOInstances, err := oInstancePsql.GetAll()
//    if len(gotAllOInstances) != 1 {
//        t.Fatalf("TestPSQLOriginRepository failed: " +
//            "expected array of 1 element after removal")
//    }
//}

func TestToolModelPSQLRepositories(t *testing.T) {
//    agreementPsql, _ := toolrepositories.NewPSQLAgreementRepo(dbName,
//                                                    dbHost,
//                                                    dbPort,
//                                                    dbUser,
//                                                    dbPass)
//    firstAgreement, err := toolrepositories.NewAgreement(
//        "first", 1, 1, 1, agreementPsql)
//    if err != nil {
//        t.Fatalf("TestPSQLAgreementRepository failed: %v", err)
//    }
//    err = agreementPsql.Insert(firstAgreement)
//    if err != nil {
//        t.Fatalf("TestPSQLAgreementRepository failed: %v", err)
//    }
//    secondAgreement, err := toolrepositories.NewAgreement(
//        "second", 1, 1, 1, agreementPsql)
//    if err != nil {
//        t.Fatalf("TestPSQLAgreementRepository failed: %v", err)
//    }
//    err = agreementPsql.Insert(secondAgreement)
//    if err != nil {
//        t.Fatalf("TestPSQLAgreementRepository failed: %v", err)
//    }
//
//    if secondAgreement.GetId() <= 0 {
//        t.Fatalf("TestPSQLAgreementRepository failed: " +
//            "expected id higher than 0 for second element")
//    }
//    gotAgreement, _ := agreementPsql.Get(1)
//    if gotAgreement.GetId() != 1 {
//        t.Fatalf("TestPSQLAgreementRepository failed: " +
//            "expected to fetch element with id = 1")
//    }
//    err = agreementPsql.Remove(0)
//    if err != nil {
//        t.Fatalf("TestPSQLAgreementRepository failed: %v", err)
//    }
//    gotAllAgreements, err := agreementPsql.GetAll()
//    if len(gotAllAgreements) != 1 {
//        t.Fatalf("TestPSQLAgreementRepository failed: " +
//            "expected array of 1 element after removal")
//    }
//
//    mtzUserPsql, _ := toolrepositories.NewPSQLMtzUserRepo(dbName,
//                                                    dbHost,
//                                                    dbPort,
//                                                    dbUser,
//                                                    dbPass)
//    firstMtzUser, err := toolrepositories.NewMtzUser(
//        "first", "pwd",
//        1, 1, 1,
//        "2000-01-01", "2000-01-01",
//        []byte{}, []byte{},
//        mtzUserPsql)
//    if err != nil {
//        t.Fatalf("TestPSQLMtzUserRepository failed: %v", err)
//    }
//    err = mtzUserPsql.Insert(firstMtzUser)
//    if err != nil {
//        t.Fatalf("TestPSQLMtzUserRepository failed: %v", err)
//    }
//    secondMtzUser, err := toolrepositories.NewMtzUser(
//        "second", "pwd",
//        1, 1, 1,
//        "2000-01-01", "2000-01-01",
//        []byte{}, []byte{},
//        mtzUserPsql)
//    if err != nil {
//        t.Fatalf("TestPSQLMtzUserRepository failed: %v", err)
//    }
//    err = mtzUserPsql.Insert(secondMtzUser)
//    if err != nil {
//        t.Fatalf("TestPSQLMtzUserRepository failed: %v", err)
//    }
//
//    if secondMtzUser.GetId() <= 0 {
//        t.Fatalf("TestPSQLMtzUserRepository failed: " +
//            "expected id higher than 0 for second element")
//    }
//    gotMtzUser, _ := mtzUserPsql.Get(1)
//    if gotMtzUser.GetId() != 1 {
//        t.Fatalf("TestPSQLMtzUserRepository failed: " +
//            "expected to fetch element with id = 1")
//    }
//    err = mtzUserPsql.Remove(0)
//    if err != nil {
//        t.Fatalf("TestPSQLMtzUserRepository failed: %v", err)
//    }
//    gotAllMtzUsers, err := mtzUserPsql.GetAll()
//    if len(gotAllMtzUsers) != 1 {
//        t.Fatalf("TestPSQLMtzUserRepository failed: " +
//            "expected array of 1 element after removal")
//    }
//
//    stakeholderPsql, _ := toolrepositories.NewPSQLStakeholderRepo(dbName,
//                                                    dbHost,
//                                                    dbPort,
//                                                    dbUser,
//                                                    dbPass)
//    firstStakeholder, err := toolrepositories.NewStakeholder(
//        "second", []int{}, stakeholderPsql)
//    if err != nil {
//        t.Fatalf("TestPSQLStakeholderRepository failed: %v", err)
//    }
//    err = stakeholderPsql.Insert(firstStakeholder)
//    if err != nil {
//        t.Fatalf("TestPSQLStakeholderRepository failed: %v", err)
//    }
//    secondStakeholder, err := toolrepositories.NewStakeholder(
//        "second", []int{}, stakeholderPsql)
//    if err != nil {
//        t.Fatalf("TestPSQLStakeholderRepository failed: %v", err)
//    }
//    err = stakeholderPsql.Insert(secondStakeholder)
//    if err != nil {
//        t.Fatalf("TestPSQLStakeholderRepository failed: %v", err)
//    }
//
//    if secondStakeholder.GetId() <= 0 {
//        t.Fatalf("TestPSQLStakeholderRepository failed: " +
//            "expected id higher than 0 for second element")
//    }
//    gotStakeholder, _ := stakeholderPsql.Get(1)
//    if gotStakeholder.GetId() != 1 {
//        t.Fatalf("TestPSQLStakeholderRepository failed: " +
//            "expected to fetch element with id = 1")
//    }
//    err = stakeholderPsql.Remove(0)
//    if err != nil {
//        t.Fatalf("TestPSQLStakeholderRepository failed: %v", err)
//    }
//    gotAllStakeholders, err := stakeholderPsql.GetAll()
//    if len(gotAllStakeholders) != 1 {
//        t.Fatalf("TestPSQLStakeholderRepository failed: " +
//            "expected array of 1 element after removal")
//    }

    profilePsql, _ := toolrepositories.NewPSQLProfileRepo(dbName,
                                                    dbHost,
                                                    dbPort,
                                                    dbUser,
                                                    dbPass)
    firstProfile, err := toolrepositories.NewProfile(
        "first", map[string]bool{}, profilePsql)
    if err != nil {
        t.Fatalf("TestPSQLProfileRepository failed: %v", err)
    }
    err = profilePsql.Insert(firstProfile)
    if err != nil {
        t.Fatalf("TestPSQLProfileRepository failed: %v", err)
    }
    secondProfile, err := toolrepositories.NewProfile(
        "second", map[string]bool{}, profilePsql)
    if err != nil {
        t.Fatalf("TestPSQLProfileRepository failed: %v", err)
    }
    err = profilePsql.Insert(secondProfile)
    if err != nil {
        t.Fatalf("TestPSQLProfileRepository failed: %v", err)
    }

    if secondProfile.GetId() <= 0 {
        t.Fatalf("TestPSQLProfileRepository failed: " +
            "expected id higher than 0 for second element")
    }
    gotProfile, err := profilePsql.Get(1)
    if err != nil {
        t.Fatalf("TestPSQLProfileRepository failed: %v", err)
    }
    if gotProfile.GetId() != 1 {
        t.Fatalf("TestPSQLProfileRepository failed: " +
            "expected to fetch element with id = 1")
    }
    err = profilePsql.Remove(0)
    if err != nil {
        t.Fatalf("TestPSQLProfileRepository failed: %v", err)
    }
    gotAllProfiles, err := profilePsql.GetAll()
    if len(gotAllProfiles) != 1 {
        t.Fatalf("TestPSQLProfileRepository failed: " +
            "expected array of 1 element after removal")
    }
}
