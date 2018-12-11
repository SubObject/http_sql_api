package outputformat

type JsonOut struct {
	Code		int
	Msg			string
	ErrMsg		error
	Data		map[string]interface{}
}