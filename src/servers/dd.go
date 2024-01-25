package servers

import (
	"os/exec"
	"regexp"
	"strings"
)

func convertShortToLiveUrl(url string) (string, error) {
	if !strings.Contains(url, "v.douyin.com") {
		return url, nil
	}

	url = extractLink(url)
	// 构造 Node.js 脚本的调用命令
	cmd := exec.Command("node", "dd.js", url)

	// 运行命令并获取输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func extractLink(input string) string {
	// 正则表达式匹配链接
	// 这个正则表达式假设链接以 https 开头，以空格或结束符结束
	re := regexp.MustCompile(`https://\S+`)

	// 查找匹配的字符串
	match := re.FindString(input)

	return match
}
