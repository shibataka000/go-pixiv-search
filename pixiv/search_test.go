package pixiv

import (
	"fmt"
	"reflect"
	"testing"
)

type testStructGetNumberOfIllusts struct {
	Keyword  string
	Expected int
}

var testDataGetNumberOfIllusts = []testStructGetNumberOfIllusts{
	// 5,411,752 illusts which has "オリジナル" tag exists at 2018/03/04
	{Keyword: "オリジナル", Expected: 5000000},
	// 0 illusts which has "06c454ea4b71c7df" tag exists at 2018/03/04
	{Keyword: "06c454ea4b71c7df", Expected: 0},
}

func TestGetNumberOfIllusts(t *testing.T) {
	for _, testData := range testDataGetNumberOfIllusts {
		actual, err := GetNumberOfIllusts(testData.Keyword)
		if err != nil {
			t.Errorf("GetNumberOfIllusts(\"%s\"): Error has occured: %v", testData.Keyword, err)
		}
		if actual < testData.Expected {
			t.Errorf("GetNumberOfIllusts(\"%s\"): Excepted is over %d but actual is %d", testData.Keyword, testData.Expected, actual)
		}
	}
}

type testStructSearch struct {
	Keyword  string
	Expected []Illust
}

var testDataSearch = []testStructSearch{
	{
		Keyword: "オリジナル100000users入り",
		Expected: []Illust{
			{
				IllustID:       "44873217",
				IllustTitle:    "鵜飼い",
				IllustType:     "0",
				URL:            "https://i.pximg.net/c/240x240/img-master/img/2014/07/23/00/02/21/44873217_p0_master1200.jpg",
				Tags:           []string{"オリジナル100000users入り"},
				UserID:         "212801",
				UserName:       "Anmi",
				UserImage:      "https://i.pximg.net/user-profile/img/2013/04/01/01/51/43/6044271_80476dbff2f6de2a161910b75caf3e1b_50.jpg",
				IsBookmarkable: "",
				IsBookmarked:   "",
				Width:          1767,
				Height:         1041,
				PageCount:      1,
				BookmarkCount:  0,
				ResponseCount:  0,
				Score:          &IllustScore{View: 1175853, Nice: 137198},
				IllustURL:      "https://www.pixiv.net/member_illust.php?mode=medium&illust_id=44873217",
			},
		},
	},
	{
		Keyword: "オリジナル50000users入り",
		Expected: []Illust{
			{
				IllustID:       "2124805",
				IllustTitle:    "おかあさん、これ買って",
				IllustType:     "0",
				URL:            "https://i.pximg.net/c/240x240/img-master/img/2008/11/14/00/26/04/2124805_p0_master1200.jpg",
				Tags:           []string{"オリジナル50000users入り"},
				UserID:         "34891",
				UserName:       "nim",
				UserImage:      "https://i.pximg.net/user-profile/img/2008/11/15/08/35/46/376021_8238ce2c6983d561b7d0ce1e7752b453_50.gif",
				IsBookmarkable: "",
				IsBookmarked:   "",
				Width:          428,
				Height:         402,
				PageCount:      1,
				BookmarkCount:  0,
				ResponseCount:  0,
				Score:          &IllustScore{View: 1098074, Nice: 60612},
				IllustURL:      "https://www.pixiv.net/member_illust.php?mode=medium&illust_id=2124805",
			},
			{
				IllustID:       "1888676",
				IllustTitle:    "断固行かぬ！",
				IllustType:     "0",
				URL:            "https://i.pximg.net/c/240x240/img-master/img/2008/10/17/21/21/30/1888676_p0_master1200.jpg",
				Tags:           []string{"オリジナル50000users入り"},
				UserID:         "107576",
				UserName:       "のじゃ",
				UserImage:      "https://i.pximg.net/user-profile/img/2009/04/21/22/41/44/704965_92aff81eafa0c49a6e5f11473e677e74_50.jpg",
				IsBookmarkable: "",
				IsBookmarked:   "",
				Width:          470,
				Height:         614,
				PageCount:      1,
				BookmarkCount:  0,
				ResponseCount:  0,
				Score:          &IllustScore{View: 946724, Nice: 42433},
				IllustURL:      "https://www.pixiv.net/member_illust.php?mode=medium&illust_id=1888676",
			},
		},
	},
}

func TestSearch(t *testing.T) {
	for _, testData := range testDataSearch {
		illustList, err := Search(testData.Keyword)
		if err != nil {
			t.Errorf("Search(\"%s\"): Error has occured: %s", testData.Keyword, err)
		}

		numOfIllusts, _ := GetNumberOfIllusts(testData.Keyword)
		if len(illustList) != numOfIllusts {
			t.Errorf("Search(\"%s\"): Expected is %d but actual is %d", testData.Keyword, numOfIllusts, len(illustList))
		}

		for _, expected := range testData.Expected {

			actual, err := findIllust(expected.IllustID, illustList)
			if err != nil {
				t.Errorf("Search(\"%s\"): Expected illust %v not found in %v", testData.Keyword, expected, illustList)
			}

			for _, tag := range expected.Tags {
				if contains(tag, actual.Tags) != true {
					t.Errorf("Search(\"%s\"): Expected is %v but actual is %v", testData.Keyword, expected, actual)
				}
			}
			if actual.Score.View < expected.Score.View || actual.Score.Nice < expected.Score.Nice {
				t.Errorf("Search(\"%s\"): Expected is %v but actual is %v", testData.Keyword, expected, actual)
			}

			expected.Tags = actual.Tags
			expected.Score = actual.Score

			if reflect.DeepEqual(&expected, actual) != true {
				t.Errorf("Search(\"%s\"): Expected is %v but actual is %v", testData.Keyword, expected, actual)
			}
		}
	}
}

func contains(x string, L []string) bool {
	for _, y := range L {
		if x == y {
			return true
		}
	}
	return false
}

func findIllust(illustID string, illustList IllustList) (*Illust, error) {
	for i, illust := range illustList {
		if illust.IllustID == illustID {
			return &(illustList[i]), nil
		}
	}
	return nil, fmt.Errorf("Illust not found")
}
