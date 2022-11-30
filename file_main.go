package filemain

import (
	"fmt"
	"github.com/lishuangquan1987/ngo/io/file"
)

func main() {
	content := "this is ngo example"
	file.WriteAllText("./test.txt", content)
	result, err := file.ReadAllText("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
