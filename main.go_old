package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

func init() {
	CustomLabel_QmlRegisterType2("CustomQmlTypes", 1, 0, "CustomLabel")
	BridgeTemplate_QmlRegisterType2("CustomQmlTypes", 1, 0, "BridgeTemplate")
	//ItemTemplate_QmlRegisterType2("CustomQmlTypes", 1, 0, "ItemTemplate")

}

type CustomLabel struct {
	core.QObject

	_ func() `constructor:"init"`

	_ string `property:"text"`
}

func (l *CustomLabel) init() {
	CustomLabels = append(CustomLabels, l)
}

type BridgeTemplate struct {
	core.QObject

	_ func()                                                                            `signal:"clicked,auto"`
	_ func(*core.QVariant, []*core.QVariant, map[string]*core.QVariant, *core.QVariant) `signal:"sendVariantListMap,auto"`
	//_ func(map[string]*core.QVariant) `signal:"sendVariantListMap2,auto"`
	_ func(*core.QObject) `signal:"sendVariantListMap2,auto"`
}

func (b *BridgeTemplate) clicked() {
	for i, label := range CustomLabels {
		go func(i int, label *CustomLabel) {
			var tick int
			for range time.NewTicker(time.Duration((i+1)*25) * time.Millisecond).C {
				tick++
				label.SetText(fmt.Sprintf("%v %v", tick, time.Now().UTC().Format("15:04:05.0000")))
			}
		}(i, label)
	}
}

func (t *BridgeTemplate) sendVariantListMap(a *core.QVariant, b []*core.QVariant, c map[string]*core.QVariant, d *core.QVariant) {
	var boolean bool
	fmt.Println("sendVariantListMap:", a.ToBool(), b[0].ToDouble(&boolean), b[1].ToString(), c, d.ToBool())

	fmt.Println("map['A']:", c["A"].ToString())
	fmt.Println("map['B']:", c["B"].ToString())
	fmt.Println("map['C']:", c["C"].ToString())
	fmt.Println("map['D']:", c["D"].ToString())
	fmt.Println("map['E']:", c["E"].ToString())

}

func (t *BridgeTemplate) sendVariantListMap2(cc *core.QObject) {

	type variantListMap2 struct {
		AA bool    `json:"AA"`
		BB float64 `json:"BB"`
		CC string  `json:"CC"`
		DD struct {
			DAA bool    `json:"DAA"`
			DBB float64 `json:"DBB"`
			DCC string  `json:"DCC"`
		} `json:"DD"`
	}

	//var  myResponse variantListMap2

	fmt.Println("sendVariantListMap2:", cc)

}

/*
	fmt.Println("map['AA']:", cc["AA"].ToString())
	fmt.Println("map['BB']:", cc["BB"].ToString())
	fmt.Println("map['CC']:", cc["CC"].ToString())

	dd := cc["DD"]
	fmt.Println("map['DD']:", dd["DBB"].ToString())
	//cc.ToGoType(myResponse)
	fmt.Println("myResponse2:", cc)
	fmt.Println("myResponse2:", myResponse)
	}
*/

var CustomLabels []*CustomLabel

func initQQuickView(path string) *quick.QQuickView {

	var view = quick.NewQQuickView(nil)

	println(filepath.Dir(path))

	var watcher = core.NewQFileSystemWatcher2([]string{filepath.Dir(path)}, nil)

	var reload = func(p string) {
		println("changed:", p)
		view.SetSource(core.NewQUrl())
		view.Engine().ClearComponentCache()
		view.SetSource(core.NewQUrl3(path, 0))
	}

	//watcher.ConnectFileChanged(reload)
	watcher.ConnectDirectoryChanged(reload)

	return view
}

func main() {

	var path = filepath.Join(os.Getenv("GOPATH"), "src", "jlambert", "FenixInception3", "FenixTestInstructionBuilder", "qml", "main.qml")
	println(path)
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := widgets.NewQApplication(len(os.Args), os.Args)

	//view := quick.NewQQuickView(nil)
	var view = initQQuickView(path)
	view.SetSource(core.NewQUrl3(path, 0))

	view.SetTitle("goroutine Example")
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	//view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))
	view.Show()

	app.Exec()
}
