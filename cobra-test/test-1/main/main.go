package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {

	// 测试父子命令内部 PersistentPreRun PersistentPostRun 运行关系，关系如下：
	// 子命令优先运行，如果子命令没有则运行父命令的，都没有则不运行。PersistentPreRun PersistentPostRun 都符合。
	// PreRun 和 PostRun 不存在父子复杂关系，只是各自自己的。

	var rootCmd = &cobra.Command{
		Use:   "root [sub]",
		Short: "My root command",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd PersistentPreRun with args: %v\n", args)
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd PreRun with args: %v\n", args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd Run with args: %v\n", args)
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd PostRun with args: %v\n", args)
		},
		//PersistentPostRun: func(cmd *cobra.Command, args []string) {
		//	fmt.Printf("Inside rootCmd PersistentPostRun with args: %v\n", args)
		//},
	}

	var subCmd = &cobra.Command{
		Use:   "sub [no options!]",
		Short: "My subcommand",
		//PersistentPreRun: func(cmd *cobra.Command, args []string) {
		//	fmt.Printf("Inside subCmd PersistentPreRun with args: %v\n", args)
		//},
		//PreRun: func(cmd *cobra.Command, args []string) {
		//	fmt.Printf("Inside subCmd PreRun with args: %v\n", args)
		//},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside subCmd Run with args: %v\n", args)
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside subCmd PostRun with args: %v\n", args)
		},
		//PersistentPostRun: func(cmd *cobra.Command, args []string) {
		//	fmt.Printf("Inside subCmd PersistentPostRun with args: %v\n", args)
		//},
	}

	rootCmd.AddCommand(subCmd)

	rootCmd.SetArgs([]string{""})
	rootCmd.Execute()
	fmt.Println()
	rootCmd.SetArgs([]string{"sub", "arg1", "arg2"})
	rootCmd.Execute()
}
