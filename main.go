// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(string(event.Source.Type))).Do(); err != nil {
					log.Print(err)
				}
		if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(string(event.Source.GroupID))).Do(); err != nil {
					log.Print(err)
				}
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				// message.ID+":"+message.Text+" (replaces)"
				var msg = ""
				log.Println(message.Text)
				if strings.Contains(message.Text, "制度") {
					msg = "欣峰控股常見問與答？\r\n1.我們合作的公司是個什麼樣的公司？\r\n答：欣峰控股集團成立於2016年在亞洲金融中心的香港，基礎設施為商業轉移企業。公司是具有多年的傳統行銷和互聯網的運作與中國資深電視媒體專業人士的經驗團隊生成的。本公司作為進入中國市場的唯一渠道，將投資者銷售量與利潤最大化。此外，公司旗下有微信平台、電視購物、超商上架渠道，滿足客戶多樣化的銷售渠道，同時確保一個穩定、安全、有保障的銷售模式回報給我們客戶。\r\n\r\n2.公司如何運作？\r\n答：欣峰控股集團是個跨境兩岸的大平台，透過蘇寧雲商與中國綠地集團實現中國唯一跨境Ｏ２Ｏ虛實整合的營銷模式。綠地控股集團為中國國有房地產企業，在２０１７年進行商圈轉型，與欣峰控股簽下戰略合同，將在中國各地打造台灣城，成為台灣區唯一招商渠道，綠地在哪欣峰就在哪。蘇寧雲商為中國最大電商平台，旗下有著數萬家實體店面，與欣峰簽下戰略合同，實現台灣區跨境Ｏ２Ｏ的唯一渠道，讓台灣商家享受到高效率營銷與低成本物流等一條龍服務，此為欣峰控股強大的造血功能。\r\n\r\n3.這是一個啥樣的投資項目？\r\n答：欣峰控股集團於２０１９年將在香港ＩＰＯ上市，２０１７年在台灣區發行分紅型電子股票，讓投資人享有上市前之優惠價格，除年底股息分紅外並於ＩＰＯ上市後股價將成長２０～３０倍達到最大收益。\r\n\r\n4.什麼是電子股票？\r\n答：電子股票是企業發行的內部股票、不同於實體股票，它屬於單邊市場。根據銷售 價格自動上漲，優勢是只漲不跌，它優於實體股票的多邊市場有漲、有跌、有風險。\r\n\r\n5.啥叫電子股票拆股？\r\n答：電子股票拆股模式是借鑒於實體股票中的配送而演變來的。\r\n例如:當價格上漲到固定的價格時，就會進行拆股了。拆股後價格減半。但會員賬戶里的持股量卻增加了一倍。隨著多次拆股，會員所持有的電子股就會成倍的增長～這就是倍增！\r\n\r\n6.電子股票是怎樣賺錢的？\r\n答：電子股經過幾次拆分後，會員賬戶里的持股量會成倍的增加，進而通過量變而達到質變，使投資者利潤增值。\r\n\r\n7.上市公司的股票是怎麼回事？\r\n答：為保障每位投資者能在這裡享受資本市場帶來的效益，公司會按照每位投資的獲利比，給予新加坡主板上市公司的股票，大家一起抱團形成創收，讓投資者在股票市場上的獲利達10~100倍。\r\n\r\n8.我們需要如何投資？\r\n答:根據個人投資能力分為四個投資額度（美金）：\r\n1000元 5000元 10000元 50000元\r\n\r\n9.收益如何？多長時間？\r\n答:電子股拆分正常投資收益約 4 - 32倍。\r\n\r\n投資:1000美元 9個月左右\r\n收益：4～8千美元。\r\n\r\n投資：5000美元12個月左右\r\n收益：2.4～4.8萬美元。\r\n\r\n投資：10000美元15個月左右\r\n收益：5.6萬美元～11.2萬美元。\r\n投資：50000美元15個月左右\r\n收益：56～102萬美元。\r\n\r\n10.這個投資安全嗎？\r\n答:任何投資都會有風險。\r\n金融投資 重點是看合作公司的市場計劃性及保障安全性。\r\n欣峰是一家於薩摩亞註冊的資本管理公司受到監管及法律保護，並有世界五百強的商業合作夥伴，有著強大的造血功能，是真實有項目運作的兩岸貿促會，並非資金盤，未來將有ＩＰＯ上市，股票保值又保本。\r\n\r\n11.做國外的金融項目會不會導致資金外流？\r\n答:貨幣的屬性是流通性。\r\n靜止的貨幣是沒有任何價值的。\r\n這種合作不叫資金外流，它是互聯網平台的合作。首先看投資是否有利潤回報；同時再看資金最終的歸口在哪裡？我們通過這種合作達到投資收益後，賺取的利潤還在本國消費，拉動本國內需的。況且，現在國與國之間都在搞投資合作了，資本市場是沒有國界的。\r\n\r\n12.我們投資資金存取自由嗎？\r\n答:當我們購股後在第一次拆股後可隨時自由交易，天天可提現，一個月發放兩次，方便靈活。\r\n\r\n13.投資這個項目每個投資人都能賺到錢嗎？\r\n答:這個項目最大的亮點是：不用找人就能賺錢 !！\r\n靜態收益足以滿足我們的投資回報需求。這是一個非常好的理財項目！有著強大的公司背書，你還等什麼？"
				} 
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msg)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
		if event.Type == linebot.EventTypeJoin {
			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("歡迎加入欣峰")).Do(); err != nil {
				log.Print(err)
			}
		}

		if event.Type == linebot.EventTypeFollow {
			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("EventTypeFollow")).Do(); err != nil {
				log.Print(err)
			}
		}

		if event.Type == linebot.EventTypeUnfollow {
			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("EventTypeUnfollow")).Do(); err != nil {
				log.Print(err)
			}
		}

		if event.Type == linebot.EventTypeLeave {
			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("EventTypeLeave")).Do(); err != nil {
				log.Print(err)
			}
		}
		if event.Type == linebot.EventTypePostback {
			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("EventTypePostback")).Do(); err != nil {
				log.Print(err)
			}
		}
		if event.Type == linebot.EventTypeBeacon {
			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("EventTypeBeacon")).Do(); err != nil {
				log.Print(err)
			}
		}
	}

}
