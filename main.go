package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"

	"os"

	"github.com/cheggaaa/pb"
	"github.com/shibataka000/pixiv-search/pixiv"
	"github.com/urfave/cli"
)

func main() {
	log.SetOutput(ioutil.Discard)

	app := cli.NewApp()
	app.Name = "pixiv-seawrch"
	app.Usage = "Show pixiv illusts ranking."
	app.UsageText = "pixiv-search keyword"
	app.Version = "v0.0.1"
	app.Action = func(c *cli.Context) error {
		if c.NArg() != 1 {
			cli.ShowAppHelpAndExit(c, 1)
		}

		keyword := c.Args().Get(0)

		numOfIllusts, err := pixiv.GetNumberOfIllusts(keyword)
		if err != nil {
			return err
		}

		bar := pb.StartNew(numOfIllusts)

		ch := pixiv.SearchAsync(keyword)

		illustList := pixiv.IllustList{}

		for illust := range ch {
			illustList = append(illustList, illust)
			bar.Increment()
		}
		bar.Finish()

		sort.Slice(illustList, func(i, j int) bool {
			return illustList[i].Score.Nice > illustList[j].Score.Nice
		})
		fmt.Printf("View\tNice\tURL\t\t\t\t\t\t\t\t\tTitle\n")
		for _, illust := range illustList {
			fmt.Printf("%d\t%d\t%s\t%s\n",
				illust.Score.View,
				illust.Score.Nice,
				illust.IllustURL,
				illust.IllustTitle)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
