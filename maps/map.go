package maps

func Search(mp map[string]string, key string) string {

	return mp[key]

}

type CustomMap map[string]string

func (cm CustomMap) Search(key string) string {
	return cm[key]
}
