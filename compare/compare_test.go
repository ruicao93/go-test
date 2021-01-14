package compare

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"
)

type MapWrapper map[string]string

func getReturnValue() (res bool) {
	return
}

func TestCompareStringSlice(t *testing.T) {
	sa := []string{"1", "2"}
	sb := []string{"1", "2"}
	sc := []string{"2", "1"}
	t.Log(reflect.DeepEqual(sa, sb))
	t.Log(reflect.DeepEqual(sa, sc))

	ma := map[string]string{"1": "1", "2": "2"}
	mb := map[string]string{"2": "2", "1": "1"}
	t.Log(reflect.DeepEqual(ma, mb))

	e1 := fmt.Errorf("error 1")
	errs := []error{e1}
	e1 = fmt.Errorf("error2")
	t.Log(errs)

	mp := MapWrapper(ma)
	t.Logf("%p, %p\n", ma, mp)
	mp["1"] = "11"
	t.Log(ma)
	t.Log(mp)

	mc := map[string]interface{}{"1": "1"}
	str := "1"
	t.Log(mc["1"] == str)

	str2 := fmt.Sprintf(`os=$(uname);if [[ "$os" == "Windows_NT" ]];then for i in $(seq 1 20); do date && nc -vz -w 2 %s.%s %d && exit 0 || sleep 2; done; exit 1; else for i in $(seq 1 5     ); do date && nc -vz -w 8 %s.%s %d && exit 0 || sleep 1; done; exit 1; fi`, "a", "b", 80, "8.8", "8.8", 80)
	t.Log(str2)

	t.Log(getReturnValue())
}

type Box struct {
	Number int
}

func TestReflect(t *testing.T) {
	b1 := Box{Number: 1}
	boxType := reflect.TypeOf(b1)
	b2 := reflect.New(boxType).Interface().(*Box)
	t.Log(b1.Number)
	t.Log(b2.Number)
	lock := sync.Mutex{}
	lock.Lock()
	t.Log("Locked")
	lock.Unlock()
	//lock.Unlock()
	t.Log("Unlock twice is not allowed")
}

func TestSlice(t *testing.T) {
	f := func(list []int) {
		list[0] = 2
	}
	list := []int{1}
	t.Log("Before pass list as param", list)
	f(list)
	t.Log("After pass list as param", list)

	m := make(map[string][]string)
	v, ok := m["1"]
	t.Logf("Value exists: %v, value: %v", ok, v)
	t.Log(m["1"])
	l2 := append(m["1"], "2")
	t.Log(l2)
}

func TestWaitGroup(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time.Sleep(1)
		wg.Done()
		t.Log("Task Done")
	}()
	wg.Wait()
	t.Log("Task Finish")
}
