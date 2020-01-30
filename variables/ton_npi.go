package variables

var TypeOfNumber = map[string]uint8{
	"unknown":           0,
	"international":     1,
	"national":          2,
	"network_specific":  3,
	"subscriber_number": 4,
	"alphanumeric":      5,
	"abbreviated":       6,
}

var NumericPlanIndicator = map[string]uint8{
	"unknown":     0,
	"isdn":        1,
	"data":        3,
	"telex":       4,
	"land_mobile": 5,
	"national":    8,
	"private":     9,
	"ermes":       10,
	"ip":          14,
	"wap":         18,
}
