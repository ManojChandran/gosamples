package tempconv

// CToF converts a Celsius to Fahernheit
func CToF(c Celsius) Fahrenhiet  {
  return Fahrenhiet(c*9/5 + 32)
}

// FToC converts a Fahrenhiet to Celsius
func FToC(f Fahrenhiet) Celsius  {
  return Celsius((f-32)* 5/9)

}
