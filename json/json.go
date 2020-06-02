package jsonutils

import "encoding/json"

type J map[string]interface{}

func (j *J) Keys() []string {
	keys := make([]string, 0, len(*j))
	for k := range *j {
		keys = append(keys, k)
	}
	return keys
}

func (j *J) Values() []interface{} {
	values := make([]interface{}, 0, len(*j))
	for _, v := range *j {
		values = append(values, v)
	}
	return values
}

func (j *J) Get(key string) interface{} {
	return (*j)[key]
}

func (j *J) Set(key string, val interface{}) {
	(*j)[key] = val
}

func (j *J) Remove(key string) {
	_, ok := (*j)[key]
	if ok {
		delete(*j, key)
	}
}

func (j *J) FromString(str string) error {
	return json.Unmarshal([]byte(str), j)
}

func (j *J) From(obj interface{}) error {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, j)
}

func (j *J) Serialize() string {
	bytes, _ := json.Marshal(j)
	return string(bytes)
}
