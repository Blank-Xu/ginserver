package db

type Model struct{}

// InsertOne insert table one record
// param: modelPtr is a pointer struct like *Model
func (p *Model) InsertOne(modelPtr interface{}, cols ...string) (int64, error) {
	return defaultEngine.Cols(cols...).InsertOne(modelPtr)
}

// Insert insert table one or more records
// param: modelsPtr is a pointer slice struct like *[]*Model
func (p *Model) Insert(modelsPtr interface{}, cols ...string) (int64, error) {
	return defaultEngine.Cols(cols...).Insert(modelsPtr)
}

// Update is update table records
// param: modelPtr is a pointer struct like *Model
func (p *Model) Update(modelPtr interface{}, cols ...string) (int64, error) {
	return defaultEngine.Cols(cols...).Update(modelPtr)
}

// UpdateCond is update table records with conditions
// param: modelPtr is a pointer struct like *Model
// param: cond is xorm builder condition
func (p *Model) UpdateCond(modelPtr, cond interface{}, cols ...string) (int64, error) {
	return defaultEngine.Cols(cols...).Where(cond).Update(modelPtr)
}

// Delete delete table records
// param: modelPtr is a pointer struct like *Model
func (p *Model) Delete(modelPtr interface{}) (int64, error) {
	return defaultEngine.Delete(modelPtr)
}

// DeleteCond delete table records
// param: modelPtr is a pointer struct like *Model
// param: cond is xorm builder condition
func (p *Model) DeleteCond(modelPtr, cond interface{}) (int64, error) {
	return defaultEngine.Where(cond).Delete(modelPtr)
}

// SelectOne select one record and reflect to struct
// param: modelPtr is a pointer struct like *Model
func (p *Model) SelectOne(modelPtr interface{}, cols ...string) (bool, error) {
	return defaultEngine.Cols(cols...).Get(modelPtr)
}

// SelectAll select table records
// param: modelPtr is a pointer struct like *Model
// param: modelsPtr is a pointer slice struct like *[]*Model
// param: cols is table's column
func (p *Model) SelectAll(modelPtr, modelsPtr interface{}, cols ...string) error {
	return defaultEngine.Cols(cols...).Find(modelsPtr, modelPtr)
}

// SelectCond select table records with condition
// param: modelPtr is a pointer struct like *Model
// param: modelsPtr is a pointer slice struct like *[]*Model
// param: cond is xorm builder condition
func (p *Model) SelectCond(modelPtr, modelsPtr, cond interface{}, orderBy string, paging *Paging, cols ...string) error {
	return defaultEngine.Cols(cols...).Where(cond).OrderBy(orderBy).Limit(paging.LimitOffset()).Find(modelsPtr, modelPtr)
}

// Count select table count
// param: modelPtr is a pointer struct like *Model
func (p *Model) Count(modelPtr interface{}) (int64, error) {
	return defaultEngine.Count(modelPtr)
}

// CountCond select table count with condition
// param: modelPtr is a pointer struct like *Model
// param: cond is xorm builder condition
func (p *Model) CountCond(modelPtr, cond interface{}) (int64, error) {
	return defaultEngine.Where(cond).Count(modelPtr)
}

// IsExists select table record with exists condition
// param: modelPtr is a pointer struct like *Model
func (p *Model) IsExists(modelPtr interface{}) (bool, error) {
	return defaultEngine.Exist(modelPtr)
}

// IsExistsCond select table record with exists condition
// param: modelPtr is a pointer struct like *Model
// param: cond is xorm builder condition
func (p *Model) IsExistsCond(modelPtr, cond interface{}, col ...string) (bool, error) {
	return defaultEngine.Cols(col...).Where(cond).Exist(modelPtr)
}
