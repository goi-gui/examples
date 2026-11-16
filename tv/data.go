package main

type Channel struct {
	Name string
	URL  string
}

var channels = []Channel{
	{
		Name: "TRT1",
		URL:  "https://tv-trt1.medya.trt.com.tr/master.m3u8",
	},
	{
		Name: "TRT Haber",
		URL:  "https://tv-trthaber.medya.trt.com.tr/master.m3u8",
	},
	{
		Name: "TRT Belgesel",
		URL:  "https://tv-trtbelgesel.medya.trt.com.tr/master.m3u8",
	},
	{
		Name: "TRT Müzik",
		URL:  "https://tv-trtmuzik.medya.trt.com.tr/master.m3u8",
	},
	{
		Name: "TRT Avaz",
		URL:  "https://tv-trtavaz.medya.trt.com.tr/master.m3u8",
	},
	{
		Name: "TRT Arabi",
		URL:  "https://tv-trtarabi.medya.trt.com.tr/master.m3u8",
	},
	{
		Name: "TRT World",
		URL:  "https://tv-trtworld.medya.trt.com.tr/master.m3u8",
	},
	{
		Name: "TRT2",
		URL:  "https://tv-trt2.medya.trt.com.tr/master.m3u8",
	},
	{
		Name: "TBMM TV",
		URL:  "https://meclistv-live.ercdn.net/meclistv/meclistv.m3u8",
	},
	{
		Name: "ATV",
		URL:  "https://rnttwmjcin.turknet.ercdn.net/lcpmvefbyo/atv/atv_1080p.m3u8",
	},
	{
		Name: "A Haber",
		URL:  "https://rnttwmjcin.turknet.ercdn.net/lcpmvefbyo/ahaber/ahaber.m3u8",
	},
	{
		Name: "A Spor",
		URL:  "https://rnttwmjcin.turknet.ercdn.net/lcpmvefbyo/aspor/aspor.m3u8",
	},
	{
		Name: "Kanal D",
		URL:  "https://demiroren.daioncdn.net/kanald/kanald.m3u8?app=kanald_web&ce=3",
	},
	{
		Name: "Star TV",
		URL:  "https://dogus.daioncdn.net/startv/startv_720p.m3u8?app=a20ac41e-bdc3-4aa1-934d-26b484480ac9&ce=3&sid=8l4w3lst4co5",
	},
	{
		Name: "NTV",
		URL:  "https://dogus.daioncdn.net/ntv/ntv.m3u8?app=ntv_web",
	},
	{
		Name: "NOW",
		URL:  "https://uycyyuuzyh.turknet.ercdn.net/nphindgytw/nowtv/nowtv_720p.m3u8",
	},
	{
		Name: "TV8",
		URL:  "https://tv8.daioncdn.net/tv8/tv8.m3u8?app=7ddc255a-ef47-4e81-ab14-c0e5f2949788&ce=3",
	},
	{
		Name: "Kanal 7",
		URL:  "https://kanal7.blutv.com/blutv_kanal7_live/live.m3u8",
	},
	{
		Name: "Halk TV",
		URL:  "https://halktv-live.daioncdn.net/halktv/halktv.m3u8",
	},
	{
		Name: "Habertürk",
		URL:  "https://tv.ensonhaber.com/haberturk/haberturk.m3u8",
	},
	{
		Name: "Haber Global",
		URL:  "https://tv.ensonhaber.com/haberglobal/haberglobal.m3u8",
	},
	{
		Name: "TGRT Haber",
		URL:  "https://canli.tgrthaber.com/tgrt.m3u8",
	},
}
