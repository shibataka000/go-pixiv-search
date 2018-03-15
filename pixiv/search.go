package pixiv

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetNumberOfIllusts(keyword string) (int, error) {
	url := fmt.Sprintf("https://www.pixiv.net/search.php?s_mode=s_tag&word=%v", keyword)

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return 0, err
	}

	countBadge := doc.Find(".count-badge").Text()

	numOfIllustsStr := strings.TrimRight(countBadge, "件")
	numOfIllusts, err := strconv.Atoi(numOfIllustsStr)
	if err != nil {
		return 0, err
	}

	return numOfIllusts, nil
}

func Search(keyword string) (IllustList, error) {
	illustList := IllustList{}

	c := SearchAsync(keyword)

	for illust := range c {
		illustList = append(illustList, illust)
	}

	return illustList, nil
}

func SearchAsync(keyword string) chan Illust {
	c := make(chan Illust)

	go func() {
		defer close(c)

		for page := 1; page <= 10; page++ {
			url := fmt.Sprintf("https://www.pixiv.net/search.php?s_mode=s_tag&word=%v&p=%v", keyword, page)
			log.Printf("Get %v", url)

			doc, err := goquery.NewDocument(url)
			if err != nil {
				log.Println(err)
				return
			}

			searchResultBody, exists := doc.Find("#js-mount-point-search-result-list").Attr("data-items")
			if exists != true {
				log.Printf("Unexpected response in %v\n", url)
				return
			}

			searchResultBodyJSON := fmt.Sprintf("{\"illustList\": %v}", searchResultBody)
			searchResultBodyBytes := ([]byte)(searchResultBodyJSON)

			searchResult := &SearchResult{}
			err = json.Unmarshal(searchResultBodyBytes, searchResult)
			if err != nil {
				log.Println(err)
				return
			}

			if len(searchResult.IllustList) == 0 {
				return
			}

			for i, illust := range searchResult.IllustList {
				score, err := getIllustScore(illust.IllustID)
				if err != nil {
					log.Println(err)
					return
				}
				searchResult.IllustList[i].Score = score
				searchResult.IllustList[i].IllustURL = fmt.Sprintf("https://www.pixiv.net/member_illust.php?mode=medium&illust_id=%v", illust.IllustID)

				c <- searchResult.IllustList[i]
			}
		}
	}()

	return c
}

func getIllustScore(illustID string) (*IllustScore, error) {
	url := fmt.Sprintf("https://www.pixiv.net/member_illust.php?mode=medium&illust_id=%v", illustID)
	log.Printf("Get %v", url)

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	illustScore := doc.Find(".score .views")
	if len(illustScore.Nodes) != 2 {
		return nil, fmt.Errorf("Unexpected response in %v", url)
	}

	viewStr := illustScore.Eq(0).Text()
	view, err := strconv.Atoi(viewStr)
	if err != nil {
		return nil, err
	}

	niceStr := illustScore.Eq(1).Text()
	nice, err := strconv.Atoi(niceStr)
	if err != nil {
		return nil, err
	}

	return &IllustScore{
		View: view,
		Nice: nice,
	}, nil
}
