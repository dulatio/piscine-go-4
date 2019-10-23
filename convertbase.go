package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return ItoaBase(AtoiBase(nbr, baseFrom), baseTo)
}
