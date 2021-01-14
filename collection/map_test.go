package collection

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestSyncMap(t *testing.T) {
	var m sync.Map
	var m2 map[string]int
	m.Store("1", 1)
	// m2 is nil, will panic
	// m2["1"] = 1
	t.Log(m2 == nil)
}

type Value struct {
	data string
}

func TestMapValue(t *testing.T) {
	v := Value{data: "data"}
	m := map[string]Value{"1": v}
	v.data = "updatedData"
	t.Log(m["1"])

	s1 := []string{"1", "2", "3"}
	s2 := []string{"1", "2", "3"}
	s3 := []string{"3", "2", "1"}
	t.Log(reflect.DeepEqual(s1, s2))
	t.Log(reflect.DeepEqual(s1, s3))
}

func parse(ii interface{}) (string, error) {
	switch ii.(type) {
	case map[string]string:
		if _, ok := ii.(map[string]string)["ovs_version"]; ok {
			return ii.(map[string]string)["ovs_version"], nil
		}
	case map[string]interface{}:
		if _, ok := ii.(map[string]interface{})["ovs_version"]; ok {
			return ii.(map[string]interface{})["ovs_version"].(string), nil
		}
	default:
		fmt.Print("default")
	}
	return "", fmt.Errorf("unknown")
}

func TestParseOVSVersion(t *testing.T) {
	m1 := map[string]string{"ovs_version": "string"}
	m2 := map[string]interface{}{"ovs_version": "interface"}
	m3 := "heheda"
	s, err := parse(m1)
	fmt.Printf("%s: %v\n", s, err)
	s, err = parse(m2)
	fmt.Printf("%s: %v\n", s, err)
	s, err = parse(m3)
	fmt.Printf("%s: %v\n", s, err)
}
