package studyReflect

import (
	"errors"
	"fmt"
	"github.com/chengxuyuanermao/goTools/conv"
	"reflect"
)

type Person struct {
	Name string   `json:"name"`
	Age  int      `json:"age"`
	Like []string `json:"like"`
}

func Use() {
	p := &Person{}
	qs := make(map[string]interface{})
	qs["name"] = "lisi"
	qs["age"] = 13
	qs["like"] = []interface{}{"ball", "run"}
	err := Analyze(qs, p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}

func Analyze(qs map[string]interface{}, req interface{}) error {
	if req == nil {
		return errors.New("nils")
	}

	vp := reflect.ValueOf(req)
	v := vp.Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := t.Field(i)
		//fmt.Printf("%+v", fieldInfo)

		tag := fieldInfo.Tag
		key := tag.Get("json")

		if value := tag.Get("required"); value == "true" && qs[key] == "" {
			return fmt.Errorf("empty field: %v", key)
		}

		fieldValue := v.FieldByName(fieldInfo.Name)
		switch fieldValue.Type().Kind() {
		case reflect.Int:
			fieldValue.Set(reflect.ValueOf(conv.ToInt(qs[key])))
		case reflect.Int32:
			fieldValue.Set(reflect.ValueOf(conv.ToInt32(qs[key])))
		case reflect.Int64:
			fieldValue.Set(reflect.ValueOf(conv.ToInt64(qs[key])))
		case reflect.String:
			fieldValue.Set(reflect.ValueOf(conv.ToString(qs[key])))
		case reflect.Float64:
			fieldValue.Set(reflect.ValueOf(conv.ToFloat64(qs[key])))
		case reflect.Struct:
			Analyze(qs, fieldValue.Addr().Interface())
		case reflect.Slice:
			if _, ok := qs[key]; !ok {
				break
			}
			arr := qs[key].([]interface{})
			tt := fieldInfo.Type.Elem() // 拿到数组中每个元素的类型信息
			fmt.Printf("fieldInfo:%+v, arr:%v, tt:%v , tt.kind:%v \n", fieldInfo, arr, tt, tt.Kind())
			switch tt.Kind() {
			case reflect.Struct:
				for _, it := range arr {
					vv := reflect.New(tt)
					Analyze(it.(map[string]interface{}), vv.Interface())
					fieldValue.Set(reflect.Append(fieldValue, vv.Elem()))
				}
			case reflect.Int:
				for _, it := range arr {
					vv := reflect.ValueOf(conv.ToInt(it))
					fieldValue.Set(reflect.Append(fieldValue, vv))
				}
			case reflect.String:
				for _, it := range arr {
					vv := reflect.ValueOf(conv.ToString(it))
					fieldValue.Set(reflect.Append(fieldValue, vv))
				}
			case reflect.Ptr:
				slice := reflect.MakeSlice(fieldInfo.Type, len(arr), len(arr))
				for i, it := range arr {
					v := reflect.New(fieldInfo.Type.Elem().Elem())
					Analyze(it.(map[string]interface{}), v.Interface())
					slice.Index(i).Set(v)
				}
				fieldValue.Set(slice)
			default:

			}

		}
	}

	return nil
}
