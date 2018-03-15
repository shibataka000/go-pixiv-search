package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"

	"os"

	"github.com/cheggaaa/pb"
	"github.com/shibataka000/pixiv-search/pixiv"
)

func main() {
	log.SetOutput(ioutil.Discard)

	if len(os.Args) != 2 {
		fmt.Println("Usage: pixiv-search KEYWORD")
		os.Exit(1)
	}

	keyword := os.Args[1]

	numOfIllusts, err := pixiv.GetNumberOfIllusts(keyword)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bar := pb.StartNew(numOfIllusts)

	c := pixiv.SearchAsync(keyword)

	illustList := pixiv.IllustList{}

	for illust := range c {
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
}
