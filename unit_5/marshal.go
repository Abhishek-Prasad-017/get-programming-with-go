package unit_5


/*
Write a program that outputs coordinates in JSON format, expanding on work done for the preceding quick check.
The JSON output should provide each coordinate in decimal degrees (DD) as well as the degrees, minutes, seconds format:
{
"decimal": 135.9,
"dms": "135o54'0.0\" E", "degrees": 135, "minutes": 54, "seconds": 0, "hemisphere": "E"
}
This can be achieved without modifying the coordinate structure by satisfying the json.
Marshaler interface to customize the JSON. The MarshalJSON method you write may make use of json.Marshal.
 */

import (
	"encoding/json"
	"fmt"
	"os"
)

type coordinate struct {
	d,m,s float64
	h rune
}

func CallJsonMasrhal() {
	ely := location2{
		coordinate{4, 30, 0.0, 'N'},
		coordinate{135, 54, 0.0, 'E'},
		"Elysium Planitia",
	}
	bytes, err := json.MarshalIndent(ely, "", " ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(bytes))
}

type location2 struct {
	Lat coordinate `json:"latitude"`
	long coordinate `json:"longitude"`
	Name string `json:"name"`
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1 }
	return sign * (c.d + c.m/60 + c.s/3600)
}

func (c coordinate) String() string {
	return fmt.Sprint("%v %v %f %c", c.d , c.m,c.s ,c.h)
}

func (c coordinate) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		DD  float64 `json:"decimal"`
		DMS string  `json:"dms"`
		D   float64 `json:"degrees"`
		M   float64 `json:"minutes"`
		S   float64 `json:"seconds"`
		H   string  `json:"hemisphere"`
	}{
		DD: c.decimal(),
		DMS: c.String(),
		D: c.d,
		M: c.m,
		S: c.s,
		H: string(c.h),
	})
}




