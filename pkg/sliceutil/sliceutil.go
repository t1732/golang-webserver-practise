// slice オブジェクトを扱う処理
package sliceutil

// chars で指定した文字配列内に target の文字列が存在するか?
func ContainsChar(chars []string, target string) bool {
	for _, char := range chars {
		if char == target {
			return true
		}
	}
	return false
}
