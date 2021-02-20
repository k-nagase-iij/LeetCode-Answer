package LongestCommonPrefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	assertCommonPrefix := func(parts string, partsLen int) bool {
		for _, v := range strs {
			if partsLen <= len(v) && v[0:partsLen] == parts {
				continue
			} else {
				return false
			}
		}

		return true
	}

	var ret string
	str := strs[0]
	for i:=0;i<len(str);i++ {
		if partOfStr := str[0:i+1];assertCommonPrefix(partOfStr, len(partOfStr)) {
			ret = partOfStr
		} else {
			break
		}
	}
	
	return ret
}