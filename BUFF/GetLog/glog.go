package GetLog

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

// 直接从json读 气
type JData struct {
	Code string `json:"code"`
	Data struct {
		Container struct {
			IconURL string `json:"icon_url"`
			Name    string `json:"name"`
		} `json:"container"`
		HasUnusual bool `json:"has_unusual"`
		Items      []struct {
			Goods struct {
				Appid              int         `json:"appid"`
				Description        interface{} `json:"description"`
				Game               string      `json:"game"`
				GoodsID            int         `json:"goods_id"`
				IconURL            string      `json:"icon_url"`
				IsCharm            bool        `json:"is_charm"`
				ItemID             interface{} `json:"item_id"`
				KeychainColorImg   interface{} `json:"keychain_color_img"`
				MarketHashName     string      `json:"market_hash_name"`
				MarketMinPrice     string      `json:"market_min_price"`
				Name               string      `json:"name"`
				OriginalIconURL    string      `json:"original_icon_url"`
				SellMinPrice       string      `json:"sell_min_price"`
				SellReferencePrice string      `json:"sell_reference_price"`
				ShortName          string      `json:"short_name"`
				SteamPrice         string      `json:"steam_price"`
				SteamPriceCny      string      `json:"steam_price_cny"`
				Tags               struct {
					Category struct {
						Category      string `json:"category"`
						ID            int    `json:"id"`
						InternalName  string `json:"internal_name"`
						LocalizedName string `json:"localized_name"`
					} `json:"category"`
					CategoryGroup struct {
						Category      string `json:"category"`
						ID            int    `json:"id"`
						InternalName  string `json:"internal_name"`
						LocalizedName string `json:"localized_name"`
					} `json:"category_group"`
					Custom struct {
						Category      string `json:"category"`
						ID            int    `json:"id"`
						InternalName  string `json:"internal_name"`
						LocalizedName string `json:"localized_name"`
					} `json:"custom"`
					Exterior struct {
						Category      string `json:"category"`
						ID            int    `json:"id"`
						InternalName  string `json:"internal_name"`
						LocalizedName string `json:"localized_name"`
					} `json:"exterior"`
					Itemset struct {
						Category      string `json:"category"`
						ID            int    `json:"id"`
						InternalName  string `json:"internal_name"`
						LocalizedName string `json:"localized_name"`
					} `json:"itemset"`
					ModelVersion struct {
						Category      string `json:"category"`
						ID            int    `json:"id"`
						InternalName  string `json:"internal_name"`
						LocalizedName string `json:"localized_name"`
					} `json:"model_version"`
					Quality struct {
						Category      string `json:"category"`
						ID            int    `json:"id"`
						InternalName  string `json:"internal_name"`
						LocalizedName string `json:"localized_name"`
					} `json:"quality"`
					Rarity struct {
						Category      string `json:"category"`
						ID            int    `json:"id"`
						InternalName  string `json:"internal_name"`
						LocalizedName string `json:"localized_name"`
					} `json:"rarity"`
					Type struct {
						Category      string `json:"category"`
						ID            int    `json:"id"`
						InternalName  string `json:"internal_name"`
						LocalizedName string `json:"localized_name"`
					} `json:"type"`
					Weapon struct {
						Category      string `json:"category"`
						ID            int    `json:"id"`
						InternalName  string `json:"internal_name"`
						LocalizedName string `json:"localized_name"`
					} `json:"weapon"`
				} `json:"tags"`
			} `json:"goods"`
			GoodsID       int    `json:"goods_id"`
			LocalizedName string `json:"localized_name"`
			MaxPrice      string `json:"max_price"`
			MinPrice      string `json:"min_price"`
		} `json:"items"`
	} `json:"data"`
	Msg interface{} `json:"msg"`
}

var Log [50]string
var Gid [50]string
var Logsize int

func Glog(name string) {
	file := "data/" + name + ".json"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		println(err.Error())
	}
	var id JData
	err = json.Unmarshal(content, &id)
	if err != nil {
		println(err)
	}
	for _, item := range id.Data.Items {
		gid := strconv.Itoa(item.Goods.GoodsID)
		ul := "https://buff.163.com/goods/" + gid
		Log[Logsize] = ul
		Gid[Logsize] = gid
		Logsize++
	}
}
