package access

import "testing"

func TestToByte(t *testing.T) {

	type people struct {
		Name     string
		Age      int
		Money    float64
		Websites []string
		Account  map[string]string
	}
	p := new(people)
	p.Name = "liyuliang"
	p.Age = 25
	p.Money = 3333.33
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
	ToByte(p)
	//b := ToByte(p)
	//println(string(b))

}
