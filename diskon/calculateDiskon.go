package diskon

func CalculateDiskon(price int, percentase int) int {
	diskon := (price * percentase) / 100
	return price - diskon
}