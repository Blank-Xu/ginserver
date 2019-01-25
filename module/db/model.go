package db

type Model struct{}

func (m *Model) InsertOne(model interface{}) error {
	_, err := defaultEngine.InsertOne(model)
	return err
}

func (m *Model) Insert(records interface{}) error {
	_, err := defaultEngine.Insert(records)
	return err
}

func (m *Model) Update(model interface{}) error {
	_, err := defaultEngine.Update(model)
	return err
}

func (m *Model) Delete(model interface{}) error {
	_, err := defaultEngine.Delete(model)
	return err
}

func (m *Model) SelectOne(model interface{}) error {
	_, err := defaultEngine.Get(model)
	return err
}

func (m *Model) Select(model interface{}, records interface{}) error {
	return defaultEngine.Find(records, model)
}

func (m *Model) SelectCond(model interface{}, cond interface{}, records interface{}) error {
	return defaultEngine.Where(cond).Find(records, model)
}

func (m *Model) Count(model interface{}) (int64, error) {
	return defaultEngine.Count(model)
}

func (m *Model) CountCond(model interface{}, cond interface{}) (int64, error) {
	return defaultEngine.Where(cond).Count(model)
}

func (m *Model) IsRecordExist(model interface{}, cond interface{}) (bool, error) {
	return defaultEngine.Where(cond).Exist(model)
}
