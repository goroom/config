package config

func FileUnmarshal(file_path string, object interface{}) error {
	cf := NewConfig()
	err := cf.LoadFile(file_path)
	if err != nil {
		return err
	}

	return cf.Unmarshal(object)
}
