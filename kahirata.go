/**********************************************************************

Copyright (C) 2020  Ashwin Godbole

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

**********************************************************************/

package main

import (
	"fmt"
	"io/ioutil"
	"github.com/pkg/term"
)

func main() {
	script := []map[string]string {
	hira,
	kata,
	}
	index := 0
	seq := ""
	out := ""
	for {
		ch, _ := getChar()
		seq += ch

		if seq == "nn" {
			out += script[index]["nn"]
			fmt.Print(script[index]["nn"])
			seq = ""
			continue
		}

		if ch == "q" || ch == string('c'&0x1f) {
			ioutil.WriteFile(("kahirata_temp"), []byte(out), 0644)
			break
		} else if ch == "a" || ch == "e" || ch == "i" || ch == "o" || ch == "u" {
			fmt.Print(script[index][seq])
			out += script[index][seq]
			seq = ""
			continue
		} else if ch == " " || ch == string(0x0D) {
			seq = ""
			out += ch
			fmt.Print(ch)
		} else if ch == string('s'&0x1f) {
			seq = ""
			if index == 0 {
				index = 1
			} else {
				index = 0
			}
			continue
		}
	}
}

var hira = map[string]string {
	"a": "あ",
	"e": "え",
	"i": "い",
	"o": "お",
	"u": "う",
	"ka":"か",
	"ke":"け",
	"ki":"き",
	"ko":"こ",
	"ku":"く",
	"kya":"きゃ",
	"kyu":"きゅ",
	"kyo":"きょ",
	"sa":"さ",
	"se":"せ",
	"shi":"し",
	"so":"そ",
	"su":"す",
	"sha":"しゃ",
	"shu":"しゅ",
	"sho":"しょ",
	"ta":"た",
	"chi":"ち",
	"tsu":"つ",
	"te":"て",
	"to":"と",
	"nn":"ん",
	"na":"な",
        "ni":"に",
        "nu":"ぬ",
        "ne":"ね",
        "no":"の",
	"ha":"は",
	"hi":"ひ",
	"fu":"ふ",
	"he":"へ",
	"ho":"ほ",
	"ma":"ま",
	"mi":"み",
	"mu":"む",
	"me":"め",
	"mo":"も",
	"ya":"や",
	"yu":"ゆ",
	"yo":"よ",
	"ra":"ら",
	"ri":"り",
	"ru":"る",
	"re":"れ",
	"ro":"ろ",
	"wa":"わ",
	"wo":"を",
	"cha":"ちゃ",
	"chu":"ちゅ",
	"cho":"ちょ",
	"nya":"にゃ",
	"nyu":"にゅ",
	"nyo":"にょ",
	"hya":"ひゃ",
	"hyu":"ひゅ",
	"hyo":"ひょ",
	"mya":"みゃ",
	"myu":"みゅ",
	"myo":"みょ",
	"rya":"りゃ",
	"ryu":"りゅ",
	"ryo":"りょ",
	"ga":"が",
	"gi":"ぎ",
	"gu":"ぐ",
	"ge":"げ",
	"go":"ご",
	"gya":"ぎゃ",
	"gyu":"ぎゅ",
	"gyo":"ぎょ",
	"za":"ざ",
	"ji":"じ",
	"zu":"ず",
	"ze":"ぜ",
	"zo":"ぞ",
	"ja":"じゃ",
	"ju":"じゅ",
	"jo":"じょ",
	"da":"だ",
	//"ji":"ぢ",
	//"zu":"づ",
	"de":"で",
	"do":"ど",
	//"ja":"ぢゃ",
	//"ju":"ぢゅ",
	//"jo":"ぢょ",
	"ba":"ば",
	"bi":"び",
	"bu":"ぶ",
	"be":"べ",
	"bo":"ぼ",
	"bya":"びゃ",
	"byu":"びゅ",
	"byo":"びょ",
	"pa":"ぱ",
	"pi":"ぴ",
	"pu":"ぷ",
	"pe":"ぺ",
	"po":"ぽ",
	"pya":"ぴゃ",
	"pyu":"ぴゅ",
	"pyo":"ぴょ",

}

var kata = map[string]string {
	"ka":"カ",
	"ki":"キ",
	"ku":"ク",
	"ke":"ケ",
	"ko":"コ",
	"kya":"キャ",
	"kyu":"キュ",
	"kyo":"キョ",
	"sa":"サ",
	"shi":"シ",
	"su":"ス",
	"se":"セ",
	"so":"ソ",
	"sha":"シャ",
	"shu":"シュ",
	"sho":"ショ",
	"ta":"タ",
	"chi":"チ",
	"tsu":"ツ",
	"te":"テ",
	"to":"ト",
	"cha":"チャ",
	"chu":"チュ",
	"cho":"チョ",
	"na":"ナ",
	"ni":"ニ",
	"nu":"ヌ",
	"ne":"ネ",
	"no":"ノ",
	"nya":"ニャ",
	"nyu":"ニュ",
	"nyo":"ニョ",
	"ha":"ハ",
	"hi":"ヒ",
	"fu":"フ",
	"he":"ヘ",
	"ho":"ホ",
	"hya":"ヒャ",
	"hyu":"ヒュ",
	"hyo":"ヒョ",
	"ma":"マ",
	"mi":"ミ",
	"mu":"ム",
	"me":"メ",
	"mo":"モ",
	"mya":"ミャ",
	"myu":"ミュ",
	"myo":"ミョ",
	"ya":"ヤ",
	"yu":"ユ",
	"yo":"ヨ",
	"ra":"ラ",
	"ri":"リ",
	"ru":"ル",
	"re":"レ",
	"ro":"ロ",
	"rya":"リャ",
	"ryu":"リュ",
	"ryo":"リョ",
	"wa":"ワ ",
	"wi":"ヰ",
	"we":"ヱ",
	"wo":"ヲ",
	"nn":"ン",
	"ga":"ガ",
	"gi":"ギ",
	"gu":"グ",
	"ge":"ゲ",
	"go":"ゴ",
	"gya":"ギャ",
	"gyu":"ギュ",
	"gyo":"ギョ",
	"za":"ザ",
	"ji":"ジ",
	"zu":"ズ",
	"ze":"ゼ",
	"zo":"ゾ",
	"ja":"ジャ",
	"ju":"ジュ",
	"jo":"ジョ",
	"da":"ダ",
	//"ji":"ヂ"
	//"zu":"ヅ"
	"de":"デ",
	"do":"ド",
	//"ja":"ヂャ"
	//"ju":"ヂュ"
	//"jo":"ヂョ"
	"ba":"バ",
	"bi":"ビ",
	"bu":"ブ",
	"be":"ベ",
	"bo":"ボ",
	"bya":"ビャ",
	"byu":"ビュ",
	"byo":"ビョ",
	"pa":"パ",
	"pi":"ピ",
	"pu":"プ",
	"pe":"ペ",
	"po":"ポ",
	"pya":"ピャ",
	"pyu":"ピュ",
	"pyo":"ピョ",
	"a":"ア",
	"i":"イ",
	"u":"ウ",
	"e":"エ",
	"o":"オ",
	//"ya":"ャ",
	//"yu":"ュ",
	//"yo":"ョ",
}

func getChar() (char string, err error) {
	t, _ := term.Open("/dev/tty")
	term.RawMode(t)
	bytes := make([]byte, 3)

	var numRead int
	numRead, err = t.Read(bytes)
	if err != nil {
		return
	} else if numRead == 1 {
		char = string(bytes[0])
	}
	t.Restore()
	t.Close()
	return
}
