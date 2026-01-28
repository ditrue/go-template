package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/ditrue/go-template/global"
	"github.com/ditrue/go-template/model/app"
)

func importGroupPhonics() error {
	var list = `
/gri/ http://rs.fengkuangapp.com/spe/Sk77t.mp3
/draɪ/ http://rs.fengkuangapp.com/spe/oPr6d.mp3
/vər/ http://rs.fengkuangapp.com/spe/SC5fd.mp3
/əd/ http://rs.fengkuangapp.com/spe/uDvw5.mp3
/vænst/ http://rs.fengkuangapp.com/spe/LEM3m.mp3
/ni/ http://rs.fengkuangapp.com/spe/XYD8m.mp3
/mʌ/ http://rs.fengkuangapp.com/spe/t5CZM.mp3
/bju:/ http://rs.fengkuangapp.com/spe/MGzyJ.mp3
/fl/ http://rs.fengkuangapp.com/spe/vHawh.mp3
/tɪ/ http://rs.fengkuangapp.com/spe/LKfJS.mp3
/swe/ http://rs.fengkuangapp.com/spe/Yo3uG.mp3
/bɑːr/ http://rs.fengkuangapp.com/spe/YjWcB.mp3
/kju:/ http://rs.fengkuangapp.com/spe/aMLf8.mp3
/twi:n/ http://rs.fengkuangapp.com/spe/qCFqw.mp3
/æn/ http://rs.fengkuangapp.com/spe/CWn6C.mp3
/sər/ http://rs.fengkuangapp.com/spe/6Ysn9.mp3
/fæk/ http://rs.fengkuangapp.com/spe/jFvvj.mp3
/blæŋ/ http://rs.fengkuangapp.com/spe/mZhrt.mp3
/kɪt/ http://rs.fengkuangapp.com/spe/8rfRo.mp3
/gən/ http://rs.fengkuangapp.com/spe/CaYdt.mp3
/məs/ http://rs.fengkuangapp.com/spe/zw3yX.mp3
/saɪd/ http://rs.fengkuangapp.com/spe/ShXM6.mp3
/məs/ http://rs.fengkuangapp.com/spe/zw3yX.mp3
/saɪd/ http://rs.fengkuangapp.com/spe/ShXM6.mp3
/ɪ/ http://rs.fengkuangapp.com/spe/veG2U.mp3
/lek/ http://rs.fengkuangapp.com/spe/K7eDp.mp3
/nɪk/ http://rs.fengkuangapp.com/spe/vwEEr.mp3
/trɑː/ http://rs.fengkuangapp.com/spe/LKLHy.mp3
/trɑː/ http://rs.fengkuangapp.com/spe/LKLHy.mp3
/æ/ http://rs.fengkuangapp.com/spe/8TXmg.mp3
/dres/ http://rs.fengkuangapp.com/spe/zxLWB.mp3
/ɪn/ http://rs.fengkuangapp.com/spe/pK45a.mp3
/tre/ http://rs.fengkuangapp.com/spe/P4VRD.mp3
/stɪŋ/ http://rs.fengkuangapp.com/spe/VKsEP.mp3
/ɪn/ http://rs.fengkuangapp.com/spe/pK45a.mp3
/tre/ http://rs.fengkuangapp.com/spe/P4VRD.mp3
/stɪŋ/ http://rs.fengkuangapp.com/spe/VKsEP.mp3
/li:t/ http://rs.fengkuangapp.com/spe/tFm3a.mp3
/æθ/ http://rs.fengkuangapp.com/spe/YchrH.mp3
/stər/ http://rs.fengkuangapp.com/spe/TBBuL.mp3
/dɪk/ http://rs.fengkuangapp.com/spe/x27Qz.mp3
/ʃə/ http://rs.fengkuangapp.com/spe/d9Sps.mp3
/ne/ http://rs.fengkuangapp.com/spe/SeYB7.mp3
/ri/ http://rs.fengkuangapp.com/spe/RuBE4.mp3
/ri/ http://rs.fengkuangapp.com/spe/RuBE4.mp3
/deɪ/ http://rs.fengkuangapp.com/spe/2qRix.mp3
/bɜːθ/ http://rs.fengkuangapp.com/spe/SVPjC.mp3
/tʃɪn/ http://rs.fengkuangapp.com/spe/mMo3C.mp3
/kɪ/ http://rs.fengkuangapp.com/spe/Hj6Jk.mp3
/ke/ http://rs.fengkuangapp.com/spe/hdgVz.mp3
/tl/ http://rs.fengkuangapp.com/spe/k22PE.mp3
/stər/ http://rs.fengkuangapp.com/spe/TBBuL.mp3
/hæm/ http://rs.fengkuangapp.com/spe/kVh9X.mp3
/dʒe/ http://rs.fengkuangapp.com/spe/L3KPU.mp3
/nə/ http://rs.fengkuangapp.com/spe/PXEYt.mp3
/rəs/ http://rs.fengkuangapp.com/spe/y7Bat.mp3
/mɪks/ http://rs.fengkuangapp.com/spe/AkZdX.mp3
/kə/ http://rs.fengkuangapp.com/spe/tW3tS.mp3
/nɑː/ http://rs.fengkuangapp.com/spe/f2jgB.mp3
/fæn/ http://rs.fengkuangapp.com/spe/nYFYE.mp3
/dʒæ/ http://rs.fengkuangapp.com/spe/eYdFf.mp3
/ɪn/ http://rs.fengkuangapp.com/spe/5hTqQ.mp3
/strʌk/ http://rs.fengkuangapp.com/spe/GpAZh.mp3
/t:k/ http://rs.fengkuangapp.com/spe/zSRHe.mp3
/dæ/ http://rs.fengkuangapp.com/spe/B9WPh.mp3
/mɪdʒ/ http://rs.fengkuangapp.com/spe/f9ejj.mp3


	`
	lines := strings.Split(list, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			fmt.Println("格式错误", line)
			continue
		}
		symbol := parts[0]
		audioPath := parts[1]
		// 先去空格和tab，再去掉符号两边的/
		symbol = strings.TrimSpace(symbol)
		symbol = strings.Trim(symbol, `/`)
		fmt.Println(symbol)
		// 如果存在不处理
		var groupPhonic app.GroupPhonic
		if err := global.GVA_DB.Where("symbol = ?", symbol).First(&groupPhonic).Error; err == nil && groupPhonic.ID > 0 {
			fmt.Println("已存在", symbol)
			continue
		}
		// 下载音频
		resp, err := http.Get(audioPath)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		audioPath = fmt.Sprintf("audio/phonics/%s.mp3", symbol)
		os.WriteFile(audioPath, body, 0644)
		groupPhonic = app.GroupPhonic{
			Symbol:    symbol,
			AudioPath: audioPath,
		}
		global.GVA_DB.Create(&groupPhonic)
	}
	return nil
}
