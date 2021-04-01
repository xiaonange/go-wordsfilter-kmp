package go_wordsfilter_kmp

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
)

var DefaultPlaceholder = "****"
var DefaultStripSpace = false

type WordsFilter struct {
	Placeholder string
	StripSpace  bool
	Text        []string
}

// New creates a words filter.
func (wf *WordsFilter)Create() *WordsFilter {
	return &WordsFilter{
		Placeholder: DefaultPlaceholder,
		StripSpace:  DefaultStripSpace}
}

//设置是否过滤文本空格
func (wf *WordsFilter) SetStripSpace(StripSpace bool)*WordsFilter {
	wf.StripSpace = StripSpace
	return wf
}

//设置是否过滤文本空格
func (wf *WordsFilter) SetPlaceholder(Placeholder string)*WordsFilter {
	wf.Placeholder = Placeholder
	return wf
}

// 设置敏感字文本
func (wf *WordsFilter) SetText(texts []string)*WordsFilter {
	for _, text := range texts {
		wf.Add(text)
	}
	return wf
}

// 读取txt文本词条
func (wf *WordsFilter) ReadWithFile(path string) (*WordsFilter, error) {
	fd, err := os.Open(path)
	if err != nil {
		return wf, err
	}
	defer fd.Close()
	buf := bufio.NewReader(fd)
	var texts []string
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return wf, err
			}
		}
		text := strings.TrimSpace(string(line))
		if text == "" {
			continue
		}
		texts = append(texts, text)
	}
	wf.SetText(texts)
	return wf, nil
}

// 添加单个敏感词文本
func (wf *WordsFilter) Add(text string) *WordsFilter {
	if wf.StripSpace {
		text = stripSpace(text)
	}
	wf.Text = append(wf.Text, text)
	return wf
}

// 替换函数
func (wf *WordsFilter) Replace(text string) string {
	if wf.StripSpace {
		text = stripSpace(text)
	}
	for _, v := range wf.Text {
		if SearchWords(text, v) != -1 {
			text = strings.Replace(text, v, wf.Placeholder, -1)
		}
	}
	return text
}

// 是否存在
func (wf *WordsFilter) Contains(text string) bool {
	if wf.StripSpace {
		text = stripSpace(text)
	}
	for _, v := range wf.Text {
		if SearchWords(text, v) != -1 {
			return true
		}
	}
	return false
}

// 删除函数
func (wf *WordsFilter) Remove(text string) string {
	if wf.StripSpace {
		text = stripSpace(text)
	}
	for _, v := range wf.Text {
		if SearchWords(text, v) != -1 {
			text = strings.Replace(text, v, "", -1)
		}
	}
	return text
}

// 过滤字符
func stripSpace(str string) string {
	fields := strings.Fields(str)
	var bf bytes.Buffer
	for _, field := range fields {
		bf.WriteString(field)
	}
	return bf.String()
}
