package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Print(string(myoji[rand.Int()%len(myoji)]))
	fmt.Print(" ")
	fmt.Println(string(name[rand.Int()%len(name)]))
}

var myoji = []string{
	"佐藤",
	"鈴木",
	"高橋",
	"田中",
	"伊藤",
	"渡辺",
	"山本",
	"中村",
	"小林",
	"加藤",
	"吉田",
	"山田",
	"佐々木",
	"山口",
	"松本",
	"井上",
	"木村",
	"斎藤",
	"林",
	"清水",
	"山崎",
	"森",
	"阿部",
	"池田",
	"橋本",
	"山下",
	"石川",
	"中島",
	"前田",
	"藤田",
	"小川",
	"後藤",
	"岡田",
	"長谷川",
	"村上",
	"近藤",
	"石井",
	"坂本",
	"遠藤",
	"青木",
	"藤井",
	"斉藤",
	"西村",
	"太田",
	"福田",
	"三浦",
	"藤原",
	"岡本",
	"松田",
	"中川",
	"順位",
	"軒数",
	"熊谷",
	"浅野",
	"荒木",
	"大久保",
	"野田",
	"川村",
	"星野",
	"大谷",
	"堀",
	"望月",
	"黒田",
	"尾崎",
	"永田",
	"内藤",
	"広瀬",
	"松村",
	"菅野",
	"田辺",
	"西山",
	"平井",
	"大島",
	"片山",
	"岩本",
	"本間",
	"早川",
	"横田",
	"大石",
	"岡崎",
	"荒井",
	"鎌田",
	"宮田",
	"成田",
	"小田",
	"沢田",
	"石橋",
	"篠原",
	"須藤",
	"萩原",
	"高山",
	"小西",
	"南",
	"栗原",
	"松原",
	"伊東",
	"三宅",
	"福井",
	"大森",
	"奥村",
	"桑原",
	"岡",
}
var name = []string{
	"湊",
	"悠人",
	"大輝",
	"龍生",
	"陽斗",
	"陸",
	"陸斗",
	"颯真",
	"悠真",
	"颯汰",
	"蒼大",
	"悠斗",
	"陽太",
	"結人",
	"虎太郎",
	"太陽",
	"隼人",
	"遥斗",
	"陽向",
	"颯",
	"優心",
	"陽翔",
	"龍之介",
	"翔",
	"結斗",
	"春輝",
	"晴",
	"蒼",
	"蒼介",
	"智也",
	"直輝",
	"優希",
	"悠翔",
	"陽大",
	"翼",
	"琉生",
	"颯介",
	"瑛斗",
	"幹太",
	"空",
	"春翔",
	"晴琉",
	"聖",
	"奏太",
	"蒼真",
	"蒼天",
	"大智",
	"斗真",
	"楓",
	"佑真",
	"優",
	"勇斗",
	"悠",
	"雄大",
	"涼太",
	"煌",
	"煌大",
	"颯斗",
	"一輝",
	"一真",
	"瑛大",
	"詠太",
	"海音",
	"岳",
	"慶太",
	"結翔",
	"健",
	"光希",
	"航平",
	"朔也",
	"春斗",
	"瞬",
	"匠",
	"渉",
	"丈",
	"奏音",
	"蒼汰",
	"太一",
	"泰生",
	"大空",
	"大悟",
	"大晟",
	"拓海",
	"拓実",
	"暖",
	"直樹",
	"哲平",
	"碧人",
	"優斗",
	"勇翔",
	"悠雅",
	"悠介",
	"悠希",
	"悠月",
	"悠馬",
	"陽人",
	"璃空",
	"琉雅",
	"琉斗",
	"龍希",
	"龍成",
	"龍星",
	"亮太",
	"蓮斗",
	"和真",
	"翔大",
	"颯一",
	"颯人",
	"ひなた",
	"心春",
	"芽依",
	"優奈",
	"美結",
	"心咲",
	"花音",
	"心菜",
	"心音",
	"優愛",
	"陽葵",
	"芽生",
	"優花",
	"優月",
	"柚希",
	"杏奈",
	"一花",
	"真央",
	"美羽",
	"優衣",
	"凛",
	"莉子",
	"莉緒",
	"杏",
	"琴音",
	"結",
	"咲希",
	"七海",
	"心結",
	"楓",
	"優菜",
	"蘭",
	"莉愛",
	"莉央",
	"咲良",
	"志歩",
	"雫",
	"朱里",
	"朱莉",
	"心",
	"心晴",
	"美咲",
	"優芽",
	"藍",
	"梨花",
	"玲奈",
	"和奏",
	"栞",
	"茉子",
	"莉桜",
	"莉心",
	"ひまり",
	"ひより",
	"ほのか",
	"愛",
	"愛海",
	"愛桜",
	"愛奈",
	"絢音",
	"花梨",
	"結羽",
	"結音",
	"向日葵",
	"彩羽",
	"彩夏",
	"彩乃",
	"紗希",
	"心美",
	"奈々",
	"日葵",
	"寧々",
	"乃愛",
	"美空",
	"美月",
	"美織",
	"百花",
	"百華",
	"唯",
	"悠",
	"由芽",
	"陽愛",
	"陽咲",
	"すみれ",
	"ひかり",
	"愛佳",
	"愛結",
	"愛梨",
	"杏珠",
	"杏莉",
	"一華",
	"夏希",
	"夏妃",
	"希愛",
	"希実",
	"咲花",
	"咲月",
	"咲桜",
	"咲耶",
	"桜子",
	"詩織",
	"実桜",
	"心海",
	"仁菜",
	"瑞希",
	"那奈",
	"日向",
	"日菜",
	"風花",
	"萌華",
	"望愛",
	"柚奈",
	"遥",
	"遥香",
	"陽向",
	"璃音",
	"瑠花",
	"瑠奈",
	"凛音",
	"栞奈",
	"澪",
	"莉奈",
}
