package access

import "testing"

func TestToByte(t *testing.T) {

	type people struct {
		Name     string
		Age      int
		Money1    int32
		Money2    int64
		Money3    float32
		Money4    float64
		Websites []string
		Account  map[string]string
	}
	p := new(people)
	p.Name = "liyuliang"
	p.Age = 25
	p.Money1 = 31
	p.Money2 = 25
	p.Money3 = 3333.33
	p.Money4 = 3333.33
	p.Websites = []string{
		"http://1",
		"http://2",
		"http://3",
		"http://4",
	}

	p.Account = make(map[string]string)
	p.Account["ICBC"] = "100"
	p.Account["IKBC"] = "20"
	p.Account["CCBC"] = "300"
	data := ToByte(p)

	expect := `{"Name":"liyuliang","Age":"25","Money1":"31","Money2":"25","Money3":"3333.33","Money4":"3333.33","Websites":["http://1","http://2","http://3","http://4"],"Account":[{"ICBC":"100"},{"IKBC":"20"},{"CCBC":"300"}]}`
	if string(data) != string(expect) {
		t.Error("Object to byte json wrong, expect is ", string(expect), "\n but get :", string(data))
	}

}
