package diskon

func calculateDiskon(price int, percentase int) int {
	diskon := (price * percentase) / 100
	return price - diskon
}