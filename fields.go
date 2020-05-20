package over

var _globalFields = make([]string, 0)

func SetGlobalFields(fields []string) {
	_globalFields = fields
}

func GetGlobalFields() []string {
	return _globalFields
}

func AddGlobalFields(field string) {
	_globalFields = append(_globalFields, field)
}

func ClearGlobalFields() {
	_globalFields = make([]string, 0)
}