package repositories
//import(
//    "mattzero.com.br/domain/model"
//)
//
//type AbsReportRepository interface {
//    GetAll() ([]model.Report, error)
//    Get(id int) (model.Report, error)
//    Insert(domain model.Report) error
//    Remove(id int) error
//}
//
//type FakeReportRepository struct {
//    reports []model.Report
//}
//
//func (f *FakeReportRepository) GetAll() ([]model.Report, error) {
//    return f.reports, nil
//}
//func (f *FakeReportRepository) Get(id int) (model.Report, error) {
//    for _, v := range f.reports {
//        if v.GetId() == id {
//            return v, nil
//        }
//    }
//    return model.Report{}, fmt.Errorf("No report found for id %v", id)
//}
//func (f *FakeReportRepository) Insert(d model.Report) error {
//    f.reports = append(f.reports, d)
//    return nil
//}
//func (f *FakeReportRepository) Remove(id int) error {
//    size := len(f.reports)
//    newReports := make([]model.Report, 0, size)
//    for _, v := range f.reports {
//        if v.GetId() != id {
//            newReports = append(newReports, v)
//        }
//    }
//    f.reports = newReports
//    return nil
//}
