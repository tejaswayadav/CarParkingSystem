package utils

type StringList []string

func (list StringList) Contains(checkString string) bool{
	for _, v := range list{
		if v == checkString{
			return true
		}
	}
	return false
}