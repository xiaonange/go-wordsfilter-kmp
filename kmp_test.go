package go_wordsfilter_kmp

import "testing"

func TestKmp(t *testing.T) {
	texts := []string{
		"子程序参数",
		"子程序指针",
		"自动控制模式",
		"自动运行启动",
		"自动运行停止",
	}
	kmp := new(WordsFilter).Create().SetText(texts)
	c1 := kmp.Contains("什么作用域")
	if c1 != false {
		t.Errorf("Test Contains expect false, get %T, %v", c1, c1)
	}
	c2 := kmp.Contains("什么子程序参数")
	if c2 != true {
		t.Errorf("Test Contains expect true, get %T, %v", c2, c2)
	}
	r1 := kmp.Replace("禁止什么作用域")
	if r1 != "*******什么*******" {
		t.Errorf("Test Replace expect *******什么******* %T,%v", r1, r1)
	}
	// Test read with file.
	kmp, _ = kmp.ReadWithFile("./word.txt")
	if kmp.Contains("字符串") != true {
		t.Errorf("Test Contains expect true, get %T, %v", c2, c2)
	}
}
