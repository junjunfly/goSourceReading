//trings.Map(mapping,s) 源码
/*
* 此方法为了 读取s中的字符串给mapping处理，处理好了返回字符串。要注意return string(b[0:nbytes]) 一定是整个字符串长度 因为当b!=nil时，长度一定会和字** 符串s相同。 而且在替换字符串的时候都使用了ascii码.
*/

// Map returns a copy of the string s with all its characters modified
// according to the mapping function. If mapping returns a negative value, the character is
// dropped from the string with no replacement.
func Map(mapping func(rune) rune, s string) string {
	// In the worst case, the string can grow when mapped, making
	// things unpleasant. But it's so rare we barge in assuming it's
	// fine. It could also shrink but that falls out naturally.
	maxbytes := len(s) // length of b
	nbytes := 0        // number of bytes encoded in b
	// The output buffer b is initialized on demand, the first
	// time a character differs.
	var b []byte

	for i, c := range s {
		r := mapping(c)
		if b == nil {
			if r == c {//相同表示非替换目标，继续执行
				continue
			}
			b = make([]byte, maxbytes)//开辟一个入参s大小的空间
			nbytes = copy(b, s[:i])//copy用法b是目标,s[:i]是源 开辟一个从s[:i]开始到b结束大小的空间
		}
		if r >= 0 {
			wid := 1
			if r >= utf8.RuneSelf {
				wid = utf8.RuneLen(r)
			}
			if nbytes+wid > maxbytes {//如果超长重新开辟空间
				// Grow the buffer.
				maxbytes = maxbytes*2 + utf8.UTFMax
				nb := make([]byte, maxbytes)
				copy(nb, b[0:nbytes])
				b = nb
			}
			nbytes += utf8.EncodeRune(b[nbytes:maxbytes], r)//替换ascii码(新->旧),
		}
	}
	if b == nil {
		return s
	}
	return string(b[0:nbytes])//返回整个字符串
}
//strings.IndexFunc(s,f)源码
//读取s中的字符串,交给f函数处理，返回第一个出现的位置eg:"abcabc" 具体还是要看函数里怎么写，如果传入函数里写 if r = 'a'{return true} 即返回a出现的第一个位置
func indexFunc(s string, f func(rune) bool, truth bool) int {
	start := 0
	for start < len(s) {
		wid := 1
		r := rune(s[start])
		if r >= utf8.RuneSelf {//超过单个字符的范围使用的处理方式
			r, wid = utf8.DecodeRuneInString(s[start:])//返回ascii码和他的size
		}
		if f(r) == truth {//只要第一个位置出现即返回位置
			return start
		}
		start += wid
	}
	return -1
}

