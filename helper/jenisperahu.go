package helper

func JenisPerahuDic() func(int) string {
	innerMap := map[int]string{
		1: "perahu motor",
		2: "perahu pesiar",
		3: "perahu layar",
	}

	return func(key int) string {
		return innerMap[key]
	}
}
