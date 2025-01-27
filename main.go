package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

// 拉取 GitHub 上的模板仓库
func cloneTemplateRepo(repoURL, targetDir string) error {
	// 使用 go-git 拉取远程仓库
	_, err := git.PlainClone(targetDir, false, &git.CloneOptions{
		URL: repoURL,
	})
	if err != nil {
		return fmt.Errorf("failed to clone the repository: %v", err)
	}
	return nil
}

// 修改文件夹中的内容，将模板中的项目名替换为用户输入的项目名
func replaceProjectNameInFiles(projectName, targetDir string) error {
	// 遍历文件夹，替换所有文件中的项目名
	err := filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 只处理文件，目录跳过
		if !info.IsDir() {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			// 替换模板中的项目名
			newContent := strings.Replace(string(content), "{{.ProjectName}}", projectName, -1)

			// 将替换后的内容写回文件
			err = os.WriteFile(path, []byte(newContent), 0644)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

// 创建新项目
func createProjectFromTemplate(projectName, repoURL, targetDir string) error {
	// 拉取 GitHub 上的仓库模板
	fmt.Printf("Cloning repository %s...\n", repoURL)
	err := cloneTemplateRepo(repoURL, targetDir)
	if err != nil {
		return err
	}

	// 替换模板中的项目名
	fmt.Printf("Replacing project name with: %s...\n", projectName)
	err = replaceProjectNameInFiles(projectName, targetDir)
	if err != nil {
		return err
	}

	fmt.Printf("Project %s created successfully!\n", projectName)
	return nil
}

func main() {
	var repoURL string
	var rootCmd = &cobra.Command{
		Use:   "selgo",
		Short: "Selgo is a project scaffolding tool.",
	}

	// `selgo new <projectname>` 命令
	var newCmd = &cobra.Command{
		Use:   "new <projectname>",
		Short: "Create a new project from a template",
		Args:  cobra.ExactArgs(1), // Expect exactly one argument (the project name)
		Run: func(cmd *cobra.Command, args []string) {
			projectName := args[0]
			// 设置 GitHub 仓库模板 URL
			repoURL = "https://github.com/Selteve/selgo.git"

			// 创建目标目录
			targetDir := projectName
			if _, err := os.Stat(targetDir); !os.IsNotExist(err) {
				fmt.Printf("Error: Project directory '%s' already exists.\n", targetDir)
				return
			}

			// 创建新项目
			err := createProjectFromTemplate(projectName, repoURL, targetDir)
			if err != nil {
				fmt.Println("Error:", err)
			}
		},
	}

	// 将 `new` 命令添加到 root 命令
	rootCmd.AddCommand(newCmd)

	// 执行命令
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
