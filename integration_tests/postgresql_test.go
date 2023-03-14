package integrationtests
import (
    "testing"
    "github.com/4MattTecnologia/mtz-cellen-domain/repositories"
    "os"
)

var (
    dbName = os.Getenv("DB_NAME")
    dbHost = os.Getenv("DB_HOST")
    dbPort = os.Getenv("DB_PORT")
    dbUser = os.Getenv("DB_USER")
    dbPass = os.Getenv("DB_PASS")
)
func TestPSQLRepositories(t *testing.T) {
    domainPsql, _ := repositories.NewPSQLDomainRepo(dbName,
                                          dbHost,
                                          dbPort,
                                          dbUser,
                                          dbPass)
    firstDomain, err := repositories.NewDomain("first", domainPsql)
    if err != nil {
        t.Fatalf("TestPSQLDomainRepository failed: %v", err)
    }
    err = domainPsql.Insert(firstDomain)
    if err != nil {
        t.Fatalf("TestPSQLDomainRepository failed: %v", err)
    }
    secondDomain, err := repositories.NewDomain("second", domainPsql)
    if err != nil {
        t.Fatalf("TestPSQLDomainRepository failed: %v", err)
    }
    err = domainPsql.Insert(secondDomain)
    if err != nil {
        t.Fatalf("TestPSQLDomainRepository failed: %v", err)
    }

    if secondDomain.GetId() <= 0 {
        t.Fatalf("TestPSQLDomainRepository failed: " +
            "expected id higher than 0 for second element")
    }
    gotDomain, _ := domainPsql.Get(1)
    if gotDomain.GetId() != 1 {
        t.Fatalf("TestPSQLDomainRepository failed: " +
            "expected to fetch element with id = 1")
    }
    err = domainPsql.Remove(0)
    if err != nil {
        t.Fatalf("TestPSQLDomainRepository failed: %v", err)
    }
    gotAllDomains, err := domainPsql.GetAll()
    if len(gotAllDomains) != 1 {
        t.Fatalf("TestPSQLDomainRepository failed: " +
            "expected array of 1 element after removal")
    }

    originPsql, _ := repositories.NewPSQLOriginRepo(dbName,
                                              dbHost,
                                              dbPort,
                                              dbUser,
                                              dbPass)
    connInfo := map[string]string{"appID": "int"}
    firstOrigin, err := repositories.NewOrigin(
        "first", connInfo, originPsql)
    if err != nil {
        t.Fatalf("TestPSQLOriginRepository failed: %v", err)
    }
    err = originPsql.Insert(firstOrigin)
    if err != nil {
        t.Fatalf("TestPSQLOriginRepository failed: %v", err)
    }
    secondOrigin, err := repositories.NewOrigin(
        "second", connInfo, originPsql)
    if err != nil {
        t.Fatalf("TestPSQLOriginRepository failed: %v", err)
    }
    err = originPsql.Insert(secondOrigin)
    if err != nil {
        t.Fatalf("TestPSQLOriginRepository failed: %v", err)
    }

    if secondOrigin.GetId() <= 0 {
        t.Fatalf("TestPSQLOriginRepository failed: " +
            "expected id higher than 0 for second element")
    }
    gotOrigin, _ := originPsql.Get(1)
    if gotOrigin.GetId() != 1 {
        t.Fatalf("TestPSQLOriginRepository failed: " +
            "expected to fetch element with id = 1")
    }
    err = originPsql.Remove(0)
    if err != nil {
        t.Fatalf("TestPSQLOriginRepository failed: %v", err)
    }
    gotAllOrigins, err := originPsql.GetAll()
    if len(gotAllOrigins) != 1 {
        t.Fatalf("TestPSQLOriginRepository failed: " +
            "expected array of 1 element after removal")
    }

    oInstancePsql, _ := repositories.NewPSQLOriginInstanceRepo(dbName,
                                                               dbHost,
                                                               dbPort,
                                                               dbUser,
                                                               dbPass)
    connVals := map[string]interface{}{"appID": 0}
    firstOInstance, err := repositories.NewOriginInstance(
        "first", connVals, 1, 1,
        oInstancePsql, domainPsql, originPsql)
    if err != nil {
        t.Fatalf("TestPSQLOriginRepository failed: %v", err)
    }
    err = oInstancePsql.Insert(firstOInstance)
    if err != nil {
        t.Fatalf("TestPSQLOriginRepository failed: %v", err)
    }
    secondOInstance, err := repositories.NewOriginInstance(
        "second", connVals, 1, 1,
        oInstancePsql, domainPsql, originPsql)
    if err != nil {
        t.Fatalf("TestPSQLOriginRepository failed: %v", err)
    }
    err = oInstancePsql.Insert(secondOInstance)
    if err != nil {
        t.Fatalf("TestPSQLOriginRepository failed: %v", err)
    }

    if secondOInstance.GetId() <= 0 {
        t.Fatalf("TestPSQLOriginRepository failed: " +
            "expected id higher than 0 for second element")
    }
    gotOInstance, _ := oInstancePsql.Get(1)
    if gotOInstance.GetId() != 1 {
        t.Fatalf("TestPSQLOriginRepository failed: " +
            "expected to fetch element with id = 1")
    }
    err = oInstancePsql.Remove(0)
    if err != nil {
        t.Fatalf("TestPSQLOriginRepository failed: %v", err)
    }
    gotAllOInstances, err := oInstancePsql.GetAll()
    if len(gotAllOInstances) != 1 {
        t.Fatalf("TestPSQLOriginRepository failed: " +
            "expected array of 1 element after removal")
    }
}
