package db

type Model struct{}

func (m *Model) InsertOne(model interface{}) error {
	_, err := defaultEngine.InsertOne(&model)
	return err
}

func (m *Model) Insert(records []interface{}) error {
	return nil
}

func (m *Model) Update(model interface{}) error {
	_, err := defaultEngine.Update(&model)
	return err
}

func (m *Model) Delete(model interface{}) error {
	_, err := defaultEngine.Delete(&model)
	return err
}

func (m *Model) SelectOne(model interface{}) error {
	_, err := defaultEngine.Get(&model)
	return err
}

func (m *Model) Select(model interface{}) (records []interface{}, err error) {
	err = defaultEngine.Find(&records, model)
	return
}

func (m *Model) SelectCond(model interface{}, cond interface{}) (records []interface{}, err error) {
	err = defaultEngine.Where(cond).Find(&records, model)
	return
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
