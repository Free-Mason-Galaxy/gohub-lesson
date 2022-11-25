package make

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CmdMakeSeeder = &cobra.Command{
	Use:   "seeder",
	Short: "Create seeder file, example: make seeder user",
	Run:   runMakeSeeder,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeSeeder(cmd *cobra.Command, args []string) {

	// 格式化模型名称，返回一个 Model 对象
	model := makeModelFromString(args[0])

	// os.MkdirAll 会确保父目录和子目录都会创建，第二个参数是目录权限，使用 0777
	os.MkdirAll("database/seeders", os.ModePerm)

	// 拼接目标文件路径
	filePath := fmt.Sprintf("database/seeders/%s_seeder.go", model.VariableNamePlural)

	// 基于模板创建文件（做好变量替换）
	createFileFromStub(filePath, "seeder", model)
}
