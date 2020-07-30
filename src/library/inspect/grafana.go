package inspect

import (
	"strings"
	"library/bot"
	"library/common"
	"models"
	"context"
	_ "flag"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	// "github.com/chromedp/chromedp/client"
	// "io/ioutil"
	"math"
	"time"
)



func Grafana(grafana *models.InspectGrafana) {
	parent, cancel := chromedp.NewExecAllocator(context.Background(), 
		chromedp.Flag("window-size", "1280,1024"),
		chromedp.Flag("headless", "true"),
	)
	defer cancel()
	// 创建 context
	// ctx, cancel := chromedp.NewContext(context.Background())
	ctx, cancel := chromedp.NewContext(parent)
	defer cancel()

	// 截取图片by 元素
	var buf []byte

	err := chromedp.Run(ctx, chromedp.Tasks{
		// 设置cookie
		chromedp.ActionFunc(func(ctx context.Context) error {
			// 创建 cookie 过期时间
			expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
			for _, session := range strings.Split(grafana.Sessions, ",") {
				// 添加cookie 到chromedp
				k := strings.Split(session, "=")[0]
				v := strings.Split(session, "=")[1]
				success, err := network.SetCookie(k, v).
					WithExpires(&expr).
					WithDomain(grafana.Domain).
					WithHTTPOnly(true).
					Do(ctx)
				if err != nil {
					return err
				}
				if !success {
					return fmt.Errorf("could not set cookie %q to %q", "grafana_session", "af0f0b4f7a3d4b6f44f92bbe5fcbc883")
				}
			}
			return nil
		}),
		// 访问网页
		chromedp.Navigate(grafana.Url),
		// 连接再次等待网页内容
		chromedp.WaitVisible(grafana.WaitVisible, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			common.Sleep(1)
			return nil
		}),
		chromedp.WaitVisible(grafana.WaitVisible, chromedp.ByQuery),
		// 截取图片
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			buf, err = page.CaptureScreenshot().
				WithQuality(100).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	})

	if err != nil {
		beego.Error("requese grafana err: ", err)
	}

	// if err := ioutil.WriteFile("mojotv_full_screen_shot.png", buf, 0644); err != nil {
	// 	beego.Error("save grafana photo err: ", err)
	// }

	var telegram bot.Telegram
	telegram.AkReptileInit()
	// if err := telegram.SendText("mojotv_full_screen_shot.png"); err != nil {
	// 	beego.Error("send text err: ", err)
	// }
	// if err := telegram.SendPhotoFile("mojotv_full_screen_shot.png"); err != nil {
	// 	beego.Error("send photo err: ", err)
	// }
	timeStr:=time.Now().Format("2006-01-02 15:04:05")
	if err := telegram.SendPhoto(buf, fmt.Sprintf("%s %s", timeStr, grafana.Name)); err != nil {
		beego.Error("send photo err: ", err)
	}

}
