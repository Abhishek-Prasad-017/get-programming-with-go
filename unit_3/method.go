package unit_3

/*
Write a program with celsius, fahrenheit, and kelvin types and methods to convert from any temperature type to any other temperature type.
*/

type celsius float64
type fahrenheit float64

//type kelvin float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}
func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}
func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}
