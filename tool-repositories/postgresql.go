package toolrepositories
import (
    "encoding/json"
    "strconv"
    "log"
    helperDb "github.com/4MattTecnologia/mtz-cellen-helpers/database"
    "github.com/4MattTecnologia/mtz-cellen-domain/tool-model"
    "github.com/lib/pq"
    "time"
)

// PSQL repositores for all entities in the tool modelling.
// In order to navigate, utilize the tags:
//  - AGREEMENT
//  - MODULE
//  - MTZUSER
//  - PROFILE
//  - STAKEHOLDER

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

// AGREEMENT -------------------------------------------------------------------
type PSQLAgreementRepo struct {
    helperDb.PostgreSQLDatabase
}

func NewPSQLAgreementRepo(dbName string,
                          dbHost string,
                          dbPort string,
                          dbUser string,
                          dbPwd string) (
                          *PSQLAgreementRepo, error) {
    baseDb := helperDb.PostgreSQLDatabase{}
    repo := PSQLAgreementRepo{
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

func NewCloudPSQLAgreementRepo(
        baseDb helperDb.PostgreSQLDatabase) *PSQLAgreementRepo {
    repo := PSQLAgreementRepo{
        PostgreSQLDatabase: baseDb,
    }
    repoPtr := &repo
    return repoPtr
}

func (p *PSQLAgreementRepo) Get(
        filters ...map[string]interface{}) ([]toolmodel.Agreement, error) {
    var (
        id                  int
        name                string
        numMtzUsers         int
        numMonitoredUsers   int
        pageLimit           int
        agreement           toolmodel.Agreement

        query               string = "agreement_id, agreement_name, "+
                                     "num_mtz_users, num_monitored_users, "+
                                     "page_limit FROM agreements "
        whereClause         string = ""
        params              []interface{}
    )
    if len(filters) > 0 {
        whereClause, params = parseFilters(filters[0])
    }
    data := make([]toolmodel.Agreement, 0)
    rows, err := p.DBConn.Query(query + whereClause, params...)

    if err != nil {
        log.Printf("Error in PSQLAgreementRepo Get(): %v", err)
        return []toolmodel.Agreement{}, err
    }

    for rows.Next() {
        if err := rows.Scan(&id,
                            &name,
                            &numMtzUsers,
                            &numMonitoredUsers,
                            &pageLimit);
        err != nil {
            log.Printf("Error in PSQLAgreementRepo Get(): %v", err)
            return []toolmodel.Agreement{}, err
        }
        agreement, _ = toolmodel.NewAgreement(id,
                                              name,
                                              numMtzUsers,
                                              numMonitoredUsers,
                                              pageLimit)
        data = append(data, agreement)
    }
    return data, nil
}

func (p *PSQLAgreementRepo) Insert(agreement toolmodel.Agreement) error {
    var (
        id                  int
        name                string
        numMtzUsers         int
        numMonitoredUsers   int
        pageLimit           int
    )
    id = agreement.GetId()
    name = agreement.GetName()
    numMtzUsers = agreement.GetNumMtzUsers()
    numMonitoredUsers = agreement.GetNumMonitoredUsers()
    pageLimit = agreement.GetPageLimit()
    _, err := p.DBConn.Exec(
        "INSERT INTO agreements "+
        "(agreement_id, agreement_name, "+
        "num_mtz_users, num_monitored_users, page_limit) "+
        "VALUES ($1, $2, $3, $4, $5)",
        id, name,
        numMtzUsers, numMonitoredUsers, pageLimit)
    return err
}

func (p *PSQLAgreementRepo) Remove(id int) error {
    _, err := p.DBConn.Exec(
        "DELETE FROM agreements "+
        "WHERE agreement_id = $1", id)
    return err
}

// MODULE ----------------------------------------------------------------------
type PSQLModuleRepo struct {
    helperDb.PostgreSQLDatabase
}

func NewPSQLModuleRepo(dbName string,
                       dbHost string,
                       dbPort string,
                       dbUser string,
                       dbPwd string) (
                       *PSQLModuleRepo, error) {
    baseDb := helperDb.PostgreSQLDatabase{}
    repo := PSQLModuleRepo{
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

func NewCloudPSQLModuleRepo(
        baseDb helperDb.PostgreSQLDatabase) *PSQLModuleRepo {
    repo := PSQLModuleRepo{
        PostgreSQLDatabase: baseDb,
    }
    repoPtr := &repo
    return repoPtr
}

func (p *PSQLModuleRepo) Get(
        filters ...map[string]interface{}) ([]toolmodel.Module, error) {
    var (
        id          int
        name        string
        module      toolmodel.Module
        query       string = "SELECT module_id, module_name " +
                             "FROM modules "
        whereClause string = ""
        params      []interface{}
    )
    if len(filters) > 0 {
        whereClause, params = parseFilters(filters[0])
    }
    data := make([]toolmodel.Module, 0)
    rows, err := p.DBConn.Query(query + whereClause, params...)

    if err != nil {
        log.Printf("Error in PSQLModuleRepo Get(): %v", err)
        return []toolmodel.Module{}, err
    }

    for rows.Next() {
        if err := rows.Scan(&id,
                            &name);
        err != nil {
            log.Printf("Error in PSQLModuleRepo Get(): %v", err)
            return []toolmodel.Module{}, err
        }
        module, _ = toolmodel.NewModule(id, name);
        data = append(data, module)
    }
    return data, nil
}

func (p *PSQLModuleRepo) Insert(module toolmodel.Module) error {
    var (
        id      int
        name    string
    )
    id = module.GetId()
    name = module.GetName()
    _, err := p.DBConn.Exec(
        "INSERT INTO modules "+
        "(module_id, module_name) "+
        "VALUES ($1, $2)",
        id, name)
    return err
}

func (p *PSQLModuleRepo) Remove(id int) error {
    _, err := p.DBConn.Exec(
        "DELETE FROM modules "+
        "WHERE module_id = $1", id)
    return err
}

// MTZUSER ---------------------------------------------------------------------
type PSQLMtzUserRepo struct {
    helperDb.PostgreSQLDatabase
}

func NewPSQLMtzUserRepo(dbName string,
                        dbHost string,
                        dbPort string,
                        dbUser string,
                        dbPwd string) (
                        *PSQLMtzUserRepo, error) {
    baseDb := helperDb.PostgreSQLDatabase{}
    repo := PSQLMtzUserRepo{
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

func NewCloudPSQLMtzUserRepo(
        baseDb helperDb.PostgreSQLDatabase) *PSQLMtzUserRepo {
    repo := PSQLMtzUserRepo{
        PostgreSQLDatabase: baseDb,
    }
    repoPtr := &repo
    return repoPtr
}

func (p *PSQLMtzUserRepo) Get(
        filters ...map[string]interface{}) ([]toolmodel.MtzUser, error) {
    var (
        id              int
        name            string
        password        string
        domainId        int
        stakeholderId   int
        profileId       int
        startDateRaw    time.Time
        startDate       string
        endDateRaw      time.Time
        endDate         string
        publicKey       []byte
        privateKey      []byte
        mtzUser         toolmodel.MtzUser

        query           string = "SELECT user_id, user_name, " +
                                 "password, " +
                                 "domain_id, stakeholder_id, "+
                                 "profile_id, "+
                                 "start_date, end_date, "+
                                 "public_key, private_key "+
                                 "FROM mtz_users "
        whereClause     string = ""
        params          []interface{}
    )
    if len(filters) > 0 {
        whereClause, params = parseFilters(filters[0])
    }
    data := make([]toolmodel.MtzUser, 0)
    rows, err := p.DBConn.Query(query + whereClause, params...)

    if err != nil {
        log.Printf("Error in PSQLMtzUserRepo Get(): %v", err)
        return []toolmodel.MtzUser{}, err
    }

    for rows.Next() {
        if err := rows.Scan(&id,
                            &name,
                            &password,
                            &domainId,
                            &stakeholderId,
                            &profileId,
                            &startDateRaw,
                            &endDateRaw,
                            &publicKey,
                            &privateKey);
        err != nil {
            log.Printf("Error in PSQLMtzUserRepo Get(): %v", err)
            return []toolmodel.MtzUser{}, err
        }
        startDate = startDateRaw.Format("2006-01-02")
        endDate = endDateRaw.Format("2006-01-02")
        mtzUser, _ = toolmodel.NewMtzUser(id,
                                          name,
                                          password,
                                          domainId,
                                          stakeholderId,
                                          profileId,
                                          startDate,
                                          endDate,
                                          publicKey,
                                          privateKey);
        data = append(data, mtzUser)
    }
    return data, nil
}

func (p *PSQLMtzUserRepo) Insert(mtzUser toolmodel.MtzUser) error {
    var (
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
    )
    id              = mtzUser.GetId()
    name            = mtzUser.GetName()
    password        = mtzUser.GetPassword()
    domainId        = mtzUser.GetDomainId()
    stakeholderId   = mtzUser.GetStakeholderId()
    profileId       = mtzUser.GetProfileId()
    startDate       = mtzUser.GetStartDate()
    endDate         = mtzUser.GetEndDate()
    publicKey       = mtzUser.GetPublicKey()
    privateKey      = mtzUser.GetPrivateKey()
    startDateRaw, err := time.Parse("2006-01-02", startDate)
    if err != nil {
        return err
    }
    endDateRaw, err := time.Parse("2006-01-02", endDate)
    if err != nil {
        return err
    }
    _, err = p.DBConn.Exec(
        "INSERT INTO mtz_users " +
        "(user_id, user_name, " +
        "password, " +
        "domain_id, stakeholder_id, " +
        "profile_id, " +
        "start_date, end_date, " +
        "public_key, private_key) " +
        "VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
        id, name,
        password,
        domainId, stakeholderId,
        profileId,
        startDateRaw, endDateRaw,
        publicKey, privateKey)
    return err
}

func (p *PSQLMtzUserRepo) Remove(id int) error {
    _, err := p.DBConn.Exec(
        "DELETE FROM mtz_users "+
        "WHERE user_id = $1", id)
    return err
}

// PROFILE ---------------------------------------------------------------------
type PSQLProfileRepo struct {
    helperDb.PostgreSQLDatabase
}

func NewPSQLProfileRepo(dbName string,
                        dbHost string,
                        dbPort string,
                        dbUser string,
                        dbPwd string) (
                        *PSQLProfileRepo, error) {
    baseDb := helperDb.PostgreSQLDatabase{}
    repo := PSQLProfileRepo{
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

func NewCloudPSQLProfileRepo(
        baseDb helperDb.PostgreSQLDatabase) *PSQLProfileRepo {
    repo := PSQLProfileRepo{
        PostgreSQLDatabase: baseDb,
    }
    repoPtr := &repo
    return repoPtr
}

func (p *PSQLProfileRepo) Get(
        filters ...map[string]interface{}) ([]toolmodel.Profile, error) {
    var (
        id          int
        name        string
        rawSecurity []byte
        security    map[string]bool
        profile     toolmodel.Profile
        query       string = "SELECT profile_id, profile_name, security " +
                             "FROM profiles"
        whereClause string = ""
        params      []interface{}
    )
    if len(filters) > 0 {
        whereClause, params = parseFilters(filters[0])
    }
    data := make([]toolmodel.Profile, 0)
    rows, err := p.DBConn.Query(query + whereClause, params...)

    if err != nil {
        log.Printf("Error in PSQLProfileRepo Get(): %v", err)
        return []toolmodel.Profile{}, err
    }

    for rows.Next() {
        if err := rows.Scan(&id,
                            &name,
                            &rawSecurity);
        err != nil {
            log.Printf("Error iterating through PSQLProfileRepo "+
                "Get(): %v", err)
            return []toolmodel.Profile{}, err
        }
        json.Unmarshal(rawSecurity, &security)
        profile, _ = toolmodel.NewProfile(id, name, security);
        data = append(data, profile)
    }
    return data, nil
}

func (p *PSQLProfileRepo) Insert(profile toolmodel.Profile) error {
    var (
        id      int
        name    string
        rawSecurity []byte
        security    map[string]bool
    )
    id          = profile.GetId()
    name        = profile.GetName()
    security    = profile.GetSecurity()
    rawSecurity, _ = json.Marshal(security)
    _, err := p.DBConn.Exec(
        "INSERT INTO profiles "+
        "(profile_id, profile_name, security) "+
        "VALUES ($1, $2, $3)",
        id, name, rawSecurity)
    return err
}

func (p *PSQLProfileRepo) Remove(id int) error {
    _, err := p.DBConn.Exec(
        "DELETE FROM profiles "+
        "WHERE profile_id = $1", id)
    return err
}

// STAKEHOLDER -----------------------------------------------------------------
type PSQLStakeholderRepo struct {
    helperDb.PostgreSQLDatabase
}

func NewPSQLStakeholderRepo(dbName string,
                            dbHost string,
                            dbPort string,
                            dbUser string,
                            dbPwd string) (
                            *PSQLStakeholderRepo, error) {
    baseDb := helperDb.PostgreSQLDatabase{}
    repo := PSQLStakeholderRepo{
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

func NewCloudPSQLStakeholderRepo(
        baseDb helperDb.PostgreSQLDatabase) *PSQLStakeholderRepo {
    repo := PSQLStakeholderRepo{
        PostgreSQLDatabase: baseDb,
    }
    repoPtr := &repo
    return repoPtr
}

func (p *PSQLStakeholderRepo) Get(
        filters ...map[string]interface{}) ([]toolmodel.Stakeholder, error) {
    var (
        id          int
        name        string
        domainIds   []int
        stakeholder toolmodel.Stakeholder
        query       string = "SELECT stakeholder_id, "+
                             "stakeholder_name, domain_ids " +
                             "FROM stakeholders"
        whereClause string = ""
        params      []interface{}
    )
    if len(filters) > 0 {
        whereClause, params = parseFilters(filters[0])
    }
    data := make([]toolmodel.Stakeholder, 0)
    rows, err := p.DBConn.Query(query + whereClause, params...)

    if err != nil {
        log.Printf("Error in PSQLStakeholderRepo Get(): %v", err)
        return []toolmodel.Stakeholder{}, err
    }

    for rows.Next() {
        if err := rows.Scan(&id,
                            &name,
                            pq.Array(&domainIds));
        err != nil {
            log.Printf("Error in PSQLStakeholderRepo Get(): %v", err)
            return []toolmodel.Stakeholder{}, err
        }
        stakeholder, _ = toolmodel.NewStakeholder(id, name, domainIds);
        data = append(data, stakeholder)
    }
    return data, nil
}

func (p *PSQLStakeholderRepo) Insert(stakeholder toolmodel.Stakeholder) error {
    var (
        id          int
        name        string
        domainIds   []int
    )
    id          = stakeholder.GetId()
    name        = stakeholder.GetName()
    domainIds   = stakeholder.GetDomainIds()
    _, err := p.DBConn.Exec(
        "INSERT INTO stakeholders "+
        "(stakeholder_id, stakeholder_name, domain_ids) "+
        "VALUES ($1, $2, $3)",
        id, name, pq.Array(domainIds))
    return err
}

func (p *PSQLStakeholderRepo) Remove(id int) error {
    _, err := p.DBConn.Exec(
        "DELETE FROM stakeholders "+
        "WHERE stakeholder_id = $1", id)
    return err
}
