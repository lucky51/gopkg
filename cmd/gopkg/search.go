package main

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/lucky51/gopkg/internal/crawler"
	"github.com/spf13/cobra"
)

var keyWord string
var top int
var greenFun, YellowFunc = color.New(color.FgGreen).SprintFunc(), color.New(color.FgYellow).SprintFunc()
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search go.dev package",
	Run: func(cmd *cobra.Command, args []string) {
		if keyWord == "" {
			cmd.Println("please input package keyword")
			return
		}
		lis, err := crawler.Crawl(keyWord, top)
		if err != nil {
			log.Fatal(err.Error())
		}
		for _, item := range lis {
			fmt.Printf("%s %s \n", greenFun(item.PkgName), YellowFunc(item.PkgPath))
			if item.PkgDescription != "" {
				color.Blue(item.PkgDescription)
			}
			fmt.Println("")
		}
	},
}

func init() {
	searchCmd.Flags().StringVarP(&keyWord, "keyword", "k", "", "package keyword")
	searchCmd.Flags().IntVarP(&top, "top", "t", 5, "limits the rows returned in a query result.")
}
