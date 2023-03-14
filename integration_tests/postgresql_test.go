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
func TestPSQLDomainRepository(t *testing.T) {
    psql, _ := repositories.NewPSQLDomainRepo(dbName,
                                          dbHost,
                                          dbPort,
                                          dbUser,
                                          dbPass)
    first, err := repositories.NewDomain("first", psql)
    if err != nil {
        t.Fatalf("TestPSQLDomainRepository failed: %v", err)
    }
    err = psql.Insert(first)
    if err != nil {
        t.Fatalf("TestPSQLDomainRepository failed: %v", err)
    }
    second, err := repositories.NewDomain("second", psql)
    if err != nil {
        t.Fatalf("TestPSQLDomainRepository failed: %v", err)
    }
    err = psql.Insert(second)
    if err != nil {
        t.Fatalf("TestPSQLDomainRepository failed: %v", err)
    }

    if second.GetId() <= 0 {
        t.Fatalf("TestPSQLDomainRepository failed: " +
            "expected id higher than 0 for second element")
    }
    got, _ := psql.Get(1)
    if got.GetId() != 1 {
        t.Fatalf("TestPSQLDomainRepository failed: " +
            "expected to fetch element with id = 1")
    }
    err = psql.Remove(0)
    if err != nil {
        t.Fatalf("TestPSQLDomainRepository failed: %v", err)
    }
    gotAll, err := psql.GetAll()
    if len(gotAll) != 1 {
        t.Fatalf("TestPSQLDomainRepository failed: " +
            "expected array of 1 element after removal")
    }
}

//func TestPSQLOriginInstanceRepository(t *testing.T) {
//}
//
//func TestPSQLOriginRepository(t *testing.T) {
//}
