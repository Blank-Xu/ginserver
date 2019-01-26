package db

type Model struct{}

// InsertOne insert table one record
// param: modelPtr is a pointer struct like *Model
func (p *Model) InsertOne(modelPtr interface{}) (int64, error) {
	return defaultEngine.InsertOne(modelPtr)
}

// Insert insert table one or more records
// param: modelsPtr is a pointer slice struct like *[]*Model
func (p *Model) Insert(modelsPtr interface{}) (int64, error) {
	return defaultEngine.Insert(modelsPtr)
}

// Update is update table records
// param: modelPtr is a pointer struct like *Model
func (p *Model) Update(modelPtr interface{}) (int64, error) {
	return defaultEngine.Update(modelPtr)
}

// Delete delete table records
// param: modelPtr is a pointer struct like *Model
func (p *Model) Delete(modelPtr interface{}) (int64, error) {
	return defaultEngine.Delete(modelPtr)
}

// SelectOne select one table record and reflect to struct
// param: modelPtr is a pointer struct like *Model
func (p *Model) SelectOne(modelPtr interface{}) (bool, error) {
	return defaultEngine.Get(modelPtr)
}

// Select select table records
// param: modelPtr is a pointer struct like *Model
// param: modelsPtr is a pointer slice struct like *[]*Model
func (p *Model) Select(modelPtr interface{}, modelsPtr interface{}) error {
	return defaultEngine.Find(modelsPtr, modelPtr)
}

// SelectCond select table records with condition
// param: modelPtr is a pointer struct like *Model
// param: modelsPtr is a pointer slice struct like *[]*Model
func (p *Model) SelectCond(modelPtr interface{}, cond interface{}, modelsPtr interface{}) error {
	return defaultEngine.Where(cond).Find(modelsPtr, modelPtr)
}

// Count select table count
// param: modelPtr is a pointer struct like *Model
func (p *Model) Count(modelPtr interface{}) (int64, error) {
	return defaultEngine.Count(modelPtr)
}

// CountCond select table count with condition
// param: modelPtr is a pointer struct like *Model
func (p *Model) CountCond(modelPtr interface{}, cond interface{}) (int64, error) {
	return defaultEngine.Where(cond).Count(modelPtr)
}

// IsRecordExist select table record with exists condition
// param: modelPtr is a pointer struct like *Model
func (p *Model) IsRecordExists(modelPtr interface{}, cond interface{}) (bool, error) {
	return defaultEngine.Where(cond).Exist(modelPtr)
}
