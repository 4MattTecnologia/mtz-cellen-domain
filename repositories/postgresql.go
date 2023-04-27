package repositories
import (
    "encoding/json"
    "strconv"
    "log"
    helperDb "github.com/4MattTecnologia/mtz-cellen-helpers/database"
    "github.com/4MattTecnologia/mtz-cellen-domain/model"
    _ "github.com/lib/pq"
    "fmt"
)

// PSQL repositores for all entities in the domain modelling.
// In order to navigate, utilize the tags:
//  - DOMAIN
//  - ORIGIN
//  - OINSTANCE
//  - REPORT
//  - ETL_DATA

func parseFilters(filters map[string]interface{}) (string, []interface{}) {
    counter := 1
    whereClause := "WHERE "
    params := make([]interface{}, 0)
    for k, v := range(filters) {
        if counter > 1 {
            whereClause += "AND "
        }
        whereClause = whereClause + k + "= $" +
                      strconv.Itoa(counter) + " "
        counter += 1
        params = append(params, v)
    }
    return whereClause, params
}


// DOMAIN ----------------------------------------------------------------------
type PSQLDomainRepo struct {
    helperDb.PostgreSQLDatabase
}

func NewPSQLDomainRepo(dbName string,
                       dbHost string,
                       dbPort string,
                       dbUser string,
                       dbPwd string) (*PSQLDomainRepo, error) {
    baseDb := helperDb.PostgreSQLDatabase{}
    repo := PSQLDomainRepo{
        PostgreSQLDatabase: baseDb,
    }
    repoPtr := &repo
    err := repoPtr.Connect(dbName,
                           dbHost,
                           dbPort,
                           dbUser,
                           dbPwd)
    return repoPtr, err
}

func NewCloudPSQLDomainRepo(
        baseDb helperDb.PostgreSQLDatabase) *PSQLDomainRepo {
    repo := PSQLDomainRepo{
        PostgreSQLDatabase: baseDb,
    }
    repoPtr := &repo
    return repoPtr
}

func (p *PSQLDomainRepo) Get(
        filters ...map[string]interface{}) ([]model.Domain, error) {
    var (
        id      int
        name    string
        domain  model.Domain

        query           string = "SELECT domain_id, domain_name "+
                                 "FROM domains "
        whereClause     string = ""
        params          []interface{}
    )

    if len(filters) > 0 {
        whereClause, params = parseFilters(filters[0])
    }

    data := make([]model.Domain, 0)
    rows, err := p.DBConn.Query(query + whereClause, params...)

    if err != nil {
        log.Printf("Error in PSQLDomainRepo Get(): %v", err)
        return []model.Domain{}, err
    }

    for rows.Next() {
        if err := rows.Scan(&id, &name);
        err != nil {
            log.Printf("Error in PSQLDomainRepo Get(): %v", err)
            return []model.Domain{}, err
        }
        domain, _ = model.NewDomain(id, name)
        data = append(data, domain)
    }
    return data, nil
}

//func (p *PSQLDomainRepo) Get(id int) (model.Domain, error) {
//    var (
//        auxId   int
//        name    string
//        domain  model.Domain
//    )
//    err := p.DBConn.QueryRow(
//        "SELECT domain_id, domain_name "+
//        "FROM domains "+
//        "WHERE domain_id = $1", id).Scan(&auxId, &name)
//    if err != nil {
//        log.Printf("Error in PSQLDomainRepo Get(): %v", err)
//        return model.Domain{}, err
//    }
//    domain, _ = model.NewDomain(auxId, name)
//    return domain, nil
//}

func (p *PSQLDomainRepo) Insert(domain model.Domain) error {
    var (
        id      int
        name    string
    )
    id = domain.GetId()
    name = domain.GetName()
    _, err := p.DBConn.Exec(
        "INSERT INTO domains "+
        "(domain_id, domain_name) "+
        "VALUES ($1, $2)", id, name)
    return err
}

func (p *PSQLDomainRepo) Remove(id int) error {
    _, err := p.DBConn.Exec(
        "DELETE FROM domains "+
        "WHERE domain_id = $1", id)
    return err
}

// ORIGIN ----------------------------------------------------------------------
type PSQLOriginRepo struct {
    helperDb.PostgreSQLDatabase
}

func NewPSQLOriginRepo(dbName string,
                       dbHost string,
                       dbPort string,
                       dbUser string,
                       dbPwd string) (*PSQLOriginRepo, error) {
    baseDb := helperDb.PostgreSQLDatabase{}
    repo := PSQLOriginRepo{
        PostgreSQLDatabase: baseDb,
    }
    repoPtr := &repo
    err := repoPtr.Connect(dbName,
                           dbHost,
                           dbPort,
                           dbUser,
                           dbPwd)
    return repoPtr, err
}

func NewCloudPSQLOriginRepo(
        baseDb helperDb.PostgreSQLDatabase) *PSQLOriginRepo {
    repo := PSQLOriginRepo{
        PostgreSQLDatabase: baseDb,
    }
    repoPtr := &repo
    return repoPtr
}

func (p *PSQLOriginRepo) Get(
        filters ...map[string]interface{}) ([]model.Origin, error) {
    var (
        id              int
        name            string
        cInfoRaw        []byte
        connectionInfo  map[string]string
        origin  model.Origin

        query           string = "SELECT origin_id, origin_name, "+
                                 "connection_info "+
                                 "FROM origins "
        whereClause     string = ""
        params          []interface{}
    )

    if len(filters) > 0 {
        whereClause, params = parseFilters(filters[0])
    }

    data := make([]model.Origin, 0)
    rows, err := p.DBConn.Query(query + whereClause, params...)

    if err != nil {
        log.Printf("Error in PSQLOriginRepo Get(): %v", err)
        return []model.Origin{}, err
    }
    fmt.Println(rows)
    for rows.Next() {
        if err := rows.Scan(&id, &name, &cInfoRaw);
        err != nil {
            log.Printf("Error in PSQLOriginRepo Get(): %v", err)
            return []model.Origin{}, err
        }
        err = json.Unmarshal(cInfoRaw, &connectionInfo)
        if err != nil {
            log.Printf("Error in PSQLOriginRepo Get(): %v", err)
            return []model.Origin{}, err
        }
        fmt.Println(connectionInfo)
        origin, _ = model.NewOrigin(id, name, connectionInfo)
        data = append(data, origin)
        fmt.Println(origin)
    }
    return data, nil
}
//func (p *PSQLOriginRepo) Get(id int) (model.Origin, error) {
//    var (
//        auxId   int
//        name    string
//        cInfoRaw        []byte
//        connectionInfo  map[string]string
//        origin  model.Origin
//    )
//    err := p.DBConn.QueryRow(
//        "SELECT origin_id, origin_name, connection_info "+
//        "FROM origins "+
//        "WHERE origin_id = $1", id).Scan(&auxId, &name, &cInfoRaw)
//    if err != nil {
//        log.Printf("Error in PSQLOriginRepo Get(): %v", err)
//        return model.Origin{}, err
//    }
//    err = json.Unmarshal(cInfoRaw, &connectionInfo)
//    if err != nil {
//        log.Printf("Error in PSQLOriginRepo Get(): %v", err)
//        return model.Origin{}, err
//    }
//    origin, _ = model.NewOrigin(auxId, name, connectionInfo)
//    return origin, nil
//}

func (p *PSQLOriginRepo) Insert(origin model.Origin) error {
    var (
        id              int
        name            string
        cInfoRaw        []byte
        connectionInfo  map[string]string
    )
    id = origin.GetId()
    name = origin.GetName()
    connectionInfo = origin.GetConnectionInfo()
    cInfoRaw, err := json.Marshal(connectionInfo)
    if err != nil {
        log.Printf("Error in PSQLOriginRepo Insert(): %v", err)
        return err
    }
    _, err = p.DBConn.Exec(
        "INSERT INTO origins "+
        "(origin_id, origin_name, connection_info) "+
        "VALUES ($1, $2, $3)", id, name, cInfoRaw)
    return err
}

func (p *PSQLOriginRepo) Remove(id int) error {
    _, err := p.DBConn.Exec(
        "DELETE FROM origins "+
        "WHERE origin_id = $1", id)
    return err
}

// OINSTANCE -------------------------------------------------------------------
type PSQLOriginInstanceRepo struct {
    helperDb.PostgreSQLDatabase
}

func NewPSQLOriginInstanceRepo(dbName string,
                               dbHost string,
                               dbPort string,
                               dbUser string,
                               dbPwd string) (
                               *PSQLOriginInstanceRepo, error) {
    baseDb := helperDb.PostgreSQLDatabase{}
    repo := PSQLOriginInstanceRepo{
        PostgreSQLDatabase: baseDb,
    }
    repoPtr := &repo
    err := repoPtr.Connect(dbName,
                           dbHost,
                           dbPort,
                           dbUser,
                           dbPwd)
    return repoPtr, err
}

func NewCloudPSQLOriginInstanceRepo(
        baseDb helperDb.PostgreSQLDatabase) *PSQLOriginInstanceRepo {
    repo := PSQLOriginInstanceRepo{
        PostgreSQLDatabase: baseDb,
    }
    repoPtr := &repo
    return repoPtr
}

func (p *PSQLOriginInstanceRepo) Get(
        filters ...map[string]interface{}) ([]model.OriginInstance, error) {
    var (
        id              int
        name            string
        originId        int
        domainId        int
        status          bool
        cValsRaw        []byte
        connectionVals  model.ConnectionValues
        oInstance       model.OriginInstance

        query           string = "SELECT origin_instance_id, "+
                                 "origin_instance_name, "+
                                 "origin_id, domain_id, life_status, " +
                                 "connection_values " +
                                 "FROM origin_instances "
        whereClause     string = ""
        params          []interface{}
    )

    if len(filters) > 0 {
        whereClause, params = parseFilters(filters[0])
    }

    data := make([]model.OriginInstance, 0)
    rows, err := p.DBConn.Query(query + whereClause, params...)

    if err != nil {
        log.Printf("Error in PSQLOriginRepo Get(): %v", err)
        return []model.OriginInstance{}, err
    }

    for rows.Next() {
        if err := rows.Scan(&id, &name, &originId, &domainId, &status, &cValsRaw);
        err != nil {
            log.Printf("Error in PSQLOriginInstanceRepo Get(): %v", err)
            return []model.OriginInstance{}, err
        }
        err = json.Unmarshal(cValsRaw, &connectionVals)
        if err != nil {
            log.Printf("Error in PSQLOriginInstanceRepo Get(): %v", err)
            return []model.OriginInstance{}, err
        }
        oInstance, _ = model.NewOriginInstance(id, name, originId,
                                            domainId, connectionVals, status)
        data = append(data, oInstance)
    }
    return data, nil
}

//func (p *PSQLOriginInstanceRepo) Get(id int) (model.OriginInstance, error) {
//    var (
//        auxId           int
//        name            string
//        originId        int
//        domainId        int
//        cValsRaw        []byte
//        connectionVals  model.ConnectionValues
//        oInstance       model.OriginInstance
//    )
//    err := p.DBConn.QueryRow(
//        "SELECT origin_instance_id, "+
//        "origin_instance_name, "+
//        "origin_id, domain_id, " +
//        "connection_values " +
//        "FROM origin_instances "+
//        "WHERE origin_instance_id = $1", id).Scan(
//        &auxId, &name, &originId, &domainId, &cValsRaw)
//    if err != nil {
//        log.Printf("Error in PSQLOriginRepo Get(): %v", err)
//        return model.OriginInstance{}, err
//    }
//    err = json.Unmarshal(cValsRaw, &connectionVals)
//    if err != nil {
//        log.Printf("Error in PSQLOriginRepo Get(): %v", err)
//        return model.OriginInstance{}, err
//    }
//    oInstance, _ = model.NewOriginInstance(auxId, name, originId,
//                                           domainId, connectionVals)
//    return oInstance, nil
//}

func (p *PSQLOriginInstanceRepo) Insert(oInstance model.OriginInstance) error {
    var (
        id          int
        name        string
        originId    int
        domainId    int
        status      bool
        cValsRaw    []byte
        cVals       model.ConnectionValues
    )
    id          = oInstance.GetId()
    name        = oInstance.GetName()
    originId    = oInstance.GetOriginId()
    domainId    = oInstance.GetDomainId()
    status      = oInstance.GetStatus()
    cVals       = oInstance.GetConnectionValues()
    cValsRaw, err := json.Marshal(cVals)
    if err != nil {
        log.Printf("Error in PSQLOriginInstanceRepo Insert(): %v", err)
        return err
    }
    _, err = p.DBConn.Exec(
        "INSERT INTO origin_instances "+
        "(origin_instance_id, origin_instance_name, " +
        "origin_id, domain_id, life_status, connection_values) "+
        "VALUES ($1, $2, $3, $4, $5, $6)",
        id, name, originId, domainId, status, cValsRaw)
    return err
}

func (p *PSQLOriginInstanceRepo) Remove(id int) error {
    _, err := p.DBConn.Exec(
        "DELETE FROM origin_instances "+
        "WHERE origin_instance_id = $1", id)
    return err
}
