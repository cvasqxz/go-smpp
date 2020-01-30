package variables

var TypeOfNumber = map[string][]byte{
	"unknown":           []byte{0},
	"international":     []byte{1},
	"national":          []byte{2},
	"network_specific":  []byte{3},
	"subscriber_number": []byte{4},
	"alphanumeric":      []byte{5},
	"abbreviated":       []byte{6},
}

var NumericPlanIndicator = map[string][]byte{
	"unknown":     []byte{0},
	"isdn":        []byte{1},
	"data":        []byte{3},
	"telex":       []byte{4},
	"land_mobile": []byte{5},
	"national":    []byte{8},
	"private":     []byte{9},
	"ermes":       []byte{10},
	"ip":          []byte{14},
	"wap":         []byte{18},
}
