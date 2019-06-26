package access

import (
	"testing"
)

func TestToByte(t *testing.T) {

	type Car struct {
		Brand string
		Price float64
	}

	type People struct {
		Name     string
		Age      int
		Money1   int32
		Money2   int64
		Money3   float32
		Money4   float64
		Websites []string
		Account  map[string]string
		MyCar1   *Car
		MyCar2   *Car
	}

	c1 := new(Car)
	c1.Brand = "AAAA"
	c1.Price = 930382.33

	c2 := Car{
		Brand: "BBBBB",
		Price: 1828881.28,
	}

	p := new(People)
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
	p.MyCar1 = c1
	p.MyCar2 = &c2

	data := ToByte(p)

	expect := `{"Name":"liyuliang","Age":"25","Money1":"31","Money2":"25","Money3":"3333.33","Money4":"3333.33","Websites":["http://1","http://2","http://3","http://4"],"Account":{"CCBC":"300","ICBC":"100","IKBC":"20"},"MyCar1":{"Brand":"AAAA","Price":"930382.33"},"MyCar2":{"Brand":"BBBBB","Price":"1828881.28"}}`
	if string(data) != string(expect) {
		t.Error("Object to byte json wrong, expect is ", string(expect), "\n but get :", string(data))
	}
}

func TestPairsToByte(t *testing.T) {
	type Obj struct {
		P []Pairs
	}
	var pairs []Pairs
	pair1 := Pairs{
		"c",
		"val1",
	}
	pair2 := Pairs{
		"a",
		"val2",
	}
	pair3 := Pairs{
		"g",
		"val3",
	}
	pair4 := Pairs{
		"d",
		"val4",
	}
	pairs = append(pairs, pair1)
	pairs = append(pairs, pair2)
	pairs = append(pairs, pair3)
	pairs = append(pairs, pair4)

	o := new(Obj)
	o.P = pairs
	data := ToByte(o)
	expect := `{"P":{"c":"val1","a":"val2","g":"val3","d":"val4"}}`
	if string(data) != string(expect) {
		t.Error("Object with Pairs to byte json wrong, expect is ", string(expect), "\n but get :", string(data))
	}
}
