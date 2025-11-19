//有暗金,纪念品的武器切片28line index调为2只有普通则为1

package get_wear

import (
	"BUFF/creatclient"
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Data struct {
	Name    string
	Quality string
	Max     float64
	Min     float64
	Len     float64
	Type    string
}

var reg1 = regexp.MustCompile(`paintwear_choices: (.*),`)
var reg2 = regexp.MustCompile(`\d.\d+`)
var reg3 = regexp.MustCompile(`(\d+\.\d+)"\]\]`)
var reg5 = regexp.MustCompile(`data-goodsid="\d+"`)
var reg6 = regexp.MustCompile(`\d+`)

func Gwear(ul string, gid string, tpe string, index int) {
	var data Data
	var strslice = make([]string, 0, 7)
	strslice = append(strslice, gid)
	rsp := creatclient.Cclient(ul)
	doc, _ := ioutil.ReadAll(rsp.Body)
	str1 := reg1.FindString(string(doc))
	mini := reg2.FindString(str1)

	ggid := reg5.FindAllString(string(doc), -1)
	for _, g := range ggid {
		wgid := reg6.FindString(g)
		if wgid != gid {
			strslice = append(strslice, wgid)
		}
	}

	strdoc := strings.NewReader(string(doc))

	doc1, _ := goquery.NewDocumentFromReader(strdoc)
	doc1.Find("body > div.market-list > div > div.detail-header.black > div.detail-cont").Each(func(i int, s *goquery.Selection) {
		name := strings.TrimSpace(s.Find("div:nth-child(1) > h1").Text())
		quality := s.Find("p > span:nth-child(1)").Text()
		fmt.Println(name, quality, tpe)
		data.Name = name
		data.Quality = quality
		data.Type = tpe
	})

	rsp.Body.Close()
	if len(strslice) == 1 {
		time.Sleep(3 * time.Second)
		Gwear(ul, gid, tpe, index)
		return
	}
	ul2 := "https://buff.163.com/goods/" + strslice[len(strslice)-index]
	time.Sleep(1200 * time.Millisecond)
	err := errors.New("first")
	for err != nil {
		rsp2 := creatclient.Cclient(ul2)
		defer rsp2.Body.Close()
		doc2, _ := ioutil.ReadAll(rsp2.Body)
		str2 := reg1.FindString(string(doc2))
		mx := reg2.FindString(reg3.FindString(str2))
		data.Min, _ = strconv.ParseFloat(mini, 64)
		data.Max, err = strconv.ParseFloat(mx, 64)
		if err != nil {
			fmt.Println(err)
			time.Sleep(3 * time.Second)
		}
	}
	data.Len = data.Max - data.Min
	fmt.Println(data.Min, data.Max, data.Len)
	//db.DB.Exec("insert into paintwear (type,quality,name,min,max,len) values (?,?,?,?,?,?)", data.Type, data.Quality, data.Name, data.Min, data.Max, data.Len)
}
