package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/common_config"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/grpc_api/mother_server_grpc_api"
	"net"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"time"

	"github.com/google/uuid"
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

type QmlBridge struct {
	core.QObject

	_ func(data bool)          `signal:"sendToQml"`
	_ func(data string) string `slot:"sendToGo"`
	_ func() string            `slot:"generateGuid"`
}

func (l *CustomLabel) init() {
	CustomLabels = append(CustomLabels, l)
}

type BridgeTemplate struct {
	core.QObject

	_ func()                                                            `signal:"clicked,auto"`
	_ func(*core.QVariant, []*core.QVariant, map[string]*core.QVariant) `signal:"sendVariantListMap,auto"`
	_ func(itf *core.QVariant)                                          `signal:"sendglobalForJs,auto"`
	_ func(*core.QVariant)                                              `signal:"sendJsonObject,auto"`
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

func (t *BridgeTemplate) sendJsonObject(a *core.QVariant) {

	fmt.Println("sendJsonObject IN:", a.ToString())

	type response2 struct {
		Page   string `json:"a"`
		Fruits string `json:"b"`
	}
	str := a.ToString()
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println("res: ", res)
	fmt.Println("Page: ", res.Page)
	fmt.Println("Fruits: ", res.Fruits)
	//	fmt.Println("sendJsonObject:", a.ToString())
}

func (t *BridgeTemplate) sendglobalForJs(a *core.QVariant) {
	fmt.Println("sendglobalForJs:", a.ToString())
}

func (t *BridgeTemplate) sendVariantListMap(a *core.QVariant, b []*core.QVariant, c map[string]*core.QVariant) {

	var boolean bool
	fmt.Println("sendVariantListMap:", a.ToBool(), b, b[0], b[1], c)
	fmt.Println("sendVariantListMap:", a.ToBool(), b[0].ToDouble(&boolean), b[1].ToString(), c)

	varD := c["json"] //["D"]

	fmt.Println("map['A']:", c["A"].ToString())
	fmt.Println("map['B']:", c["B"].ToString())
	fmt.Println("map['C']:", c["C"].ToString())
	fmt.Println("varD.ToJsonObject():", varD.ToJsonObject())
	fmt.Println("varD.ToJsonObject().varD.ToString():", varD.ToString())
	fmt.Println("varD.ToJsonObject().Keys():", len(varD.ToJsonObject().Keys()))
	fmt.Println("varD.ToJsonObject().Keys():", varD.ToJsonObject().Keys())
	// Keys for "D"
	fmt.Println("varD.ToJsonObject().Keys():", len(varD.ToJsonObject().ToVariantMap()["D"].ToJsonObject().Keys()))
	fmt.Println("varD.ToJsonObject().Keys():", varD.ToJsonObject().ToVariantMap()["D"].ToJsonObject().Keys())

	// Next level down
	fmt.Println("varD.ToJsonObject().Keys():", len(varD.ToJsonObject().ToVariantMap()["D"].ToJsonObject().ToVariantMap()["DA"].ToJsonObject().Keys()))
	fmt.Println("varD.ToJsonObject().Keys():", varD.ToJsonObject().ToVariantMap()["D"].ToJsonObject().ToVariantMap()["DA"].ToJsonObject().Keys())
	fmt.Println("varD.ToJsonObject().Keys():", len(varD.ToJsonObject().ToVariantMap()["D"].ToJsonObject().ToVariantMap()["DC"].ToJsonObject().Keys()))
	fmt.Println("varD.ToJsonObject().Keys():", varD.ToJsonObject().ToVariantMap()["D"].ToJsonObject().ToVariantMap()["DC"].ToJsonObject().Keys())

	fmt.Println("varD.ToJsonObject().ToVariantMap():", varD.ToJsonObject().ToVariantMap())
	fmt.Println("varD.ToJsonObject().ToVariantMap()[\"DA\"].ToString()]:", varD.ToJsonObject().ToVariantMap()["DA"].ToString())

	// List all Keys with Values
	allKeys := returnAllKeys(varD.ToJsonObject())
	fmt.Println("All KeyValues for varD.ToJsonObject():", allKeys)
	// Print KeyValues as Map
	allKeysAsMap := returnKeysValueMap(allKeys)
	fmt.Println(allKeysAsMap)

	//var myJson AutoGenerated
	//qjsonData := varD.ToJsonDocument().ToJson()
	//var jsonData []byte
	jsonData := varD.ToJsonDocument().ToJson() //.. ToByteArray().Length() //.ToJsonDocument().ToJson().Length()// .ConstData()
	fmt.Println("jsonData:", jsonData)
	fmt.Println("jsonData.ConstData:", jsonData.ConstData())
	fmt.Println("*jsonData:", jsonData.Length())

	var ba = AutoGenerated2{
		DA_struct: 4.5,
		DB_struct: "Hej",
	}

	ba.PrintFields()

	var bb = AutoGenerated2{
		DA_struct: 4.5,
		DB_struct: "Hejdå",
	}

	fmt.Println("bb before:", bb)
	bb.MapKeyValueMapIntoStruct_test(allKeysAsMap)
	fmt.Println("bb after:", bb)
	bb.testChangeStruct()
	fmt.Println("bb after last call:", bb)
	bb.testChangeStruct2()
	fmt.Println("bb after last call:", bb)

	/*	err := json.Unmarshal(jsonData.ConstData(), &myJson)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("map['D']:", c["D"].ToString())
	*/
	//	fmt.Println("map['D']:", c["D"].ToString())
	//	fmt.Println("map['E']:", c["E"].ToString())

}

func (bb *AutoGenerated2) testChangeStruct() {
	bb.DA_struct = 0.001
	bb.DB_struct = "Nu är det nya värden"

}

func (dd *AutoGenerated2) testChangeStruct2() {
	dd.DA_struct = 17.001
	dd.DB_struct = "OJ OJ, Nu är det nya värden"

}

type AutoGenerated2 struct {
	DA_struct float64 `json:"DA"`
	DB_struct string  `json:"DB"`
}

func (bb AutoGenerated2) PrintFields() {
	var myValue interface{}
	var myValueToCast string
	var myFloat64 float64
	//var myValueType interface{}

	val := reflect.ValueOf(bb)
	for i := 0; i < val.Type().NumField(); i++ {
		fmt.Println(val.Type().Field(i).Name)
		fmt.Println(val.Type().Field(i).Tag.Get("json"))
		fmt.Println(val.Type().Field(i).Type)

		myValue = val.Type().Field(i)
		myValue = 3.14
		myValueToCast = "3.1415"
		//myFloat64, _ = strconv.ParseFloat(myValueToCast,64)
		//myValueType = val.Type().Field(i).Type

		//value.(type)
		switch myValue.(type) {
		case string:
			fmt.Println("Detta är en: String", myValue, val.Type().Field(i).Type)
		case int32, int64:
			fmt.Println("Detta är en:int32, int64", myValue, val.Type().Field(i).Type)
		case float32, float64:
			fmt.Println("Detta är en:float32, float64", myValue, val.Type().Field(i).Type)
			myFloat64, _ = strconv.ParseFloat(myValueToCast, 64)
			fmt.Println("Cast string into float64:", myFloat64)
		default:
			fmt.Println("Detta är en:unknown", myValue, val.Type().Field(i).Type.Kind())
		}

	}
}

func (cc *AutoGenerated2) MapKeyValueMapIntoStruct(keyValueMap map[string]keyValuePar) {

	var fieldStructName string
	var jsonKey string
	var jsonValue interface{}
	var myValue interface{}
	var myString string
	var myFloat64 float64
	//var myValueToCast string
	bb := *cc

	fmt.Println("cc:", cc)
	fmt.Println("bb:", bb)

	val := reflect.ValueOf(bb)

	for i := 0; i < val.Type().NumField(); i++ {
		fieldStructName = val.Type().Field(i).Name
		fieldStructType := val.Type().Field(i).Type
		jsonKey = val.Type().Field(i).Tag.Get("json")
		jsonValue = keyValueMap[jsonKey].qvalue.ToInterface()

		//fmt.Println(val.Type())
		fmt.Println("START")
		fmt.Println("fieldStructName", fieldStructName)
		fmt.Println("fieldStructType", fieldStructType)
		fmt.Println("jsonKey", jsonKey)
		fmt.Println("jsonValue", jsonValue)
		fmt.Println(val.Type().Field(i).Type)
		fmt.Println("END")
		//val.Set(3.14)
		//val.Set(keyValueMap[""].qvalue.ToGoType(val.Type().Field(i).Type))
		//keyValueMap[""].qvalue.ToGoType(val.Type().Field(i).Type)
		myValue = val.Type().Field(i)

		fmt.Println("myValue = val.Type().Field(i)", myValue)
		//myValueToCast = keyValueMap[jsonKey].value

		//v := reflect.ValueOf(myValue)
		//v = reflect.Indirect(v)

		switch reflect.ValueOf(&bb).Elem().FieldByName(fieldStructName).Kind() { //myValue.(type) {//myValue.(type) {

		case reflect.String:
			fmt.Println("**Detta är en: String", myValue, val.Type().Field(i).Type)
			myString, _ = getString(jsonValue)
			fmt.Println("myString", myString, "jsonValue", jsonValue)
			fmt.Println("Old value in struct for String:", bb.DB_struct)
			reflect.ValueOf(&bb).Elem().FieldByName(fieldStructName).SetString(myString)

			// pointer to struct - addressable
			ps := reflect.ValueOf(&bb)
			// struct
			s := ps.Elem()
			if s.Kind() == reflect.Struct {
				// exported field
				f := s.FieldByName(fieldStructName)
				if f.IsValid() {
					// A Value can be changed only if it is
					// addressable and was not obtained by
					// the use of unexported struct fields.
					if f.CanSet() {
						// change value of N
						if f.Kind() == reflect.String {
							//		x := string(7)
							//		if !f.OverflowInt(x) {
							f.SetString(myString)
							//		}
						}
					}
				}
			}
			fmt.Println("New value in struct for String:", bb.DB_struct)

		case reflect.Int32, reflect.Int64:
			fmt.Println("Detta är en:int32, int64", myValue, val.Type().Field(i).Type)

		case reflect.Float64:
			fmt.Println("***Detta är en:float32, float64", myValue, val.Type().Field(i).Type)
			myFloat64, _ = getFloat(jsonValue)
			reflect.ValueOf(&bb).Elem().FieldByName(fieldStructName).SetFloat(myFloat64)

			fmt.Println("New value in struct for Float64:", bb.DA_struct)
		default:
			fmt.Println("Detta är en:unknown", myValue, val.Type().Field(i).Type.Kind())
			//reflect.ValueOf(&bb).Elem().FieldByName(fieldStructName).SetFloat(jsonValue)
		}
	}

	fmt.Println("Before leaving function, bb, &bb, cc:", bb, &bb, cc)
	cc = &bb
	fmt.Println("Before leaving function, bb, &bb, cc:", bb, &bb, cc)
}

func (bb *AutoGenerated2) MapKeyValueMapIntoStruct_test(keyValueMap map[string]keyValuePar) {

	var fieldStructName string
	var jsonKey string
	var jsonValue interface{}
	var myValue interface{}
	var myString string
	var myFloat64 float64

	fmt.Println("bb:", bb)

	val := reflect.ValueOf(*bb)

	for i := 0; i < val.Type().NumField(); i++ {
		fieldStructName = val.Type().Field(i).Name
		fieldStructType := val.Type().Field(i).Type
		jsonKey = val.Type().Field(i).Tag.Get("json")
		jsonValue = keyValueMap[jsonKey].qvalue.ToInterface()

		//fmt.Println(val.Type())
		fmt.Println("START")
		fmt.Println("fieldStructName", fieldStructName)
		fmt.Println("fieldStructType", fieldStructType)
		fmt.Println("jsonKey", jsonKey)
		fmt.Println("jsonValue", jsonValue)
		fmt.Println(val.Type().Field(i).Type)
		fmt.Println("END")
		//val.Set(3.14)
		//val.Set(keyValueMap[""].qvalue.ToGoType(val.Type().Field(i).Type))
		//keyValueMap[""].qvalue.ToGoType(val.Type().Field(i).Type)
		myValue = val.Type().Field(i)

		fmt.Println("myValue = val.Type().Field(i)", myValue)
		//myValueToCast = keyValueMap[jsonKey].value

		//v := reflect.ValueOf(myValue)
		//v = reflect.Indirect(v)

		switch reflect.ValueOf(bb).Elem().FieldByName(fieldStructName).Kind() { //myValue.(type) {//myValue.(type) {

		case reflect.String:
			fmt.Println("**Detta är en: String", myValue, val.Type().Field(i).Type)
			myString, _ = getString(jsonValue)
			fmt.Println("myString", myString, "jsonValue", jsonValue)
			fmt.Println("Old value in struct for String:", &bb.DB_struct)
			reflect.ValueOf(bb).Elem().FieldByName(fieldStructName).SetString(myString)

			// pointer to struct - addressable
			ps := reflect.ValueOf(bb)
			// struct
			s := ps.Elem()
			if s.Kind() == reflect.Struct {
				// exported field
				f := s.FieldByName(fieldStructName)
				if f.IsValid() {
					// A Value can be changed only if it is
					// addressable and was not obtained by
					// the use of unexported struct fields.
					if f.CanSet() {
						// change value of N
						if f.Kind() == reflect.String {
							//		x := string(7)
							//		if !f.OverflowInt(x) {
							f.SetString(myString)
							//		}
						}
					}
				}
			}
			fmt.Println("New value in struct for String:", &bb.DB_struct)

		case reflect.Int32, reflect.Int64:
			fmt.Println("Detta är en:int32, int64", myValue, val.Type().Field(i).Type)

		case reflect.Float64:
			fmt.Println("***Detta är en:float32, float64", myValue, val.Type().Field(i).Type)
			myFloat64, _ = getFloat(jsonValue)
			reflect.ValueOf(bb).Elem().FieldByName(fieldStructName).SetFloat(myFloat64)

			fmt.Println("New value in struct for Float64:", &bb.DA_struct)
		default:
			fmt.Println("Detta är en:unknown", myValue, val.Type().Field(i).Type.Kind())
			//reflect.ValueOf(&bb).Elem().FieldByName(fieldStructName).SetFloat(jsonValue)
		}
	}

	fmt.Println("Before leaving function, bb, &bb,:", bb, &bb)

}

var floatType = reflect.TypeOf(float64(0))
var int64Type = reflect.TypeOf(int64(0))
var stringType = reflect.TypeOf(string(0))

func getFloat(unk interface{}) (float64, error) {
	v := reflect.ValueOf(unk)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(floatType) {
		return 0, fmt.Errorf("cannot convert %v to float64", v.Type())
	}
	fv := v.Convert(floatType)
	return fv.Float(), nil
}

func getInt64(unk interface{}) (int64, error) {
	v := reflect.ValueOf(unk)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(int64Type) {
		return 0, fmt.Errorf("cannot convert %v to int64", v.Type())
	}
	fv := v.Convert(int64Type)
	return fv.Int(), nil

}

func getString(unk interface{}) (string, error) {
	fmt.Println("incoming value for geString: ", unk)
	v := reflect.ValueOf(unk)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(stringType) {
		return "0", fmt.Errorf("cannot convert %v to string", v.Type())
	}
	fv := v.Convert(stringType)
	return fv.String(), nil

}

// *********************************************************************
// Creats an array of all keys with their corresponding values
//
type keyValuePar struct {
	key    string
	qvalue *core.QVariant
	value  string
	//myKind reflect.Kind
}

func returnAllKeys(jsobObject *core.QJsonObject) []keyValuePar {
	var keyArray []keyValuePar
	keyArray = nil

	var jsonObjectVariantMap map[string]*core.QVariant
	jsonObjectVariantMap = jsobObject.ToVariantMap()

	// Get Keys and number of Keys for this level
	currentKeys := jsobObject.Keys()
	numberOfCurrentKeys := len(currentKeys)

	// Loop of the the Keys if there are any
	if numberOfCurrentKeys != 0 {
		for _, keyValue := range currentKeys {
			// Check if Key has children
			if len(jsobObject.ToVariantMap()[keyValue].ToJsonObject().Keys()) == 0 {
				// WHen key doesn't have children then add value to return array
				keyArray = append(keyArray, keyValuePar{
					key:    keyValue,
					qvalue: jsonObjectVariantMap[keyValue],
					value:  jsonObjectVariantMap[keyValue].ToString(),
				})
			} else {
				// WHen Key has children then get child values and add to return Array
				childrensKeyValues := returnAllKeys(jsobObject.ToVariantMap()[keyValue].ToJsonObject())
				if childrensKeyValues != nil {
					keyArray = append(keyArray, childrensKeyValues...)
				}
			}

		}
	} else {
		// Return nil if there are no keys, this should not happened
		return nil
	}

	return keyArray
}

// *********************************************************************

// *********************************************************************
// Convert Key-Value array into an array
//

type keyValueMapType map[string]keyValuePar

func returnKeysValueMap(keyValueParArray []keyValuePar) keyValueMapType {

	var returnVKeysValueMap keyValueMapType
	//returnVKeysValueMap = nil
	returnVKeysValueMap = make(keyValueMapType)
	// If no incoming elements in array then return nil
	if len(keyValueParArray) == 0 {
		return nil
	}

	// Loop through array and create a map of the values
	for _, keyValue := range keyValueParArray {
		fmt.Println(keyValue)
		returnVKeysValueMap[keyValue.key] = keyValue
	}

	return returnVKeysValueMap
}

// *********************************************************************
func generateGuid() string {
	var returnGuid string
	newGuid, _ := uuid.NewUUID()

	returnGuid = newGuid.String()

	return returnGuid
}

// *********************************************************************
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

// Used for only process cleanup once
var cleanupProcessed bool = false

func cleanup() {

	if cleanupProcessed == false {

		cleanupProcessed = true

		// Cleanup before close down application
		motherObject.logger.WithFields(logrus.Fields{
			"ID": "ab1abcdb-d786-4bcc-b274-b52bd931f43d",
		}).Info("Clean up and shut down servers")

		motherObject.logger.WithFields(logrus.Fields{
			"ID": "a22e8de2-c38c-4483-95e0-e9dc21dd7ca8",
		}).Info("Gracefull stop for: registerTaxiHardwareStreamServer")

		registerMotherServer.GracefulStop()

		motherObject.logger.WithFields(logrus.Fields{
			"common_config.MotherServer_port: ": common_config.MotherServer_port,
			"ID":                                "19593393-3ac9-45ac-96e9-cdf911b167c7",
		}).Info("Close net.Listing")
		lis.Close()

		//log.Println("Close DB_session: %v", DB_session)
		//DB_session.Close()
	}
}

func main() {

	// ******************************************************
	var err error
	defer cleanup()

	// Set up WorkerObject
	motherObject = &MotherObject_struct{}

	// Init logger
	motherObject.InitLogger("")

	// Find first non allocated port from defined start port
	motherObject.logger.WithFields(logrus.Fields{
		"ID": "56b6419f-d714-4ab0-be62-f3c7f08b9558",
	}).Info("Mother Server tries to start")

	motherObject.logger.WithFields(logrus.Fields{
		"common_config.MotherServer_port): ": common_config.MotherServer_port,
		"ID":                                 "8f904ace-9d24-452b-891a-5b8e5c247ba2",
	}).Info("Start listening on:")
	lis, err = net.Listen("tcp", string(common_config.MotherServer_port))

	if err != nil {
		motherObject.logger.WithFields(logrus.Fields{
			"err: ": err,
			"ID":    "0fbf0f08-6114-4cd8-9992-6a387955fb5f",
		}).Fatal("failed to listen:")

	} else {
		motherObject.logger.WithFields(logrus.Fields{
			"common_config.MotherServer_port): ": common_config.MotherServer_port,
			"ID":                                 "93496c07-2b6c-4edc-a1f9-3fd555fa1201",
		}).Info("Success in listening on port:")

	}

	// Creates a new RegisterMotherServer gRPC server
	go func() {
		motherObject.logger.WithFields(logrus.Fields{
			"ID": "ebc26735-9d13-4b13-91b8-20999cd5e254",
		}).Info("Starting Mother Server")

		registerMotherServer = grpc.NewServer()
		mother_server_grpc_api.RegisterMotherServerServer(registerMotherServer, &MotherServer{})

		motherObject.logger.WithFields(logrus.Fields{
			"common_config.MotherServer_port): ": common_config.MotherServer_port,
			"ID":                                 "cfed87c4-55aa-4cd1-980a-a15eb75ab6fb",
		}).Info("registerMotherServer for Mother Server started")

		registerMotherServer.Serve(lis)
	}()
	// ******************************************************

	var path = filepath.Join(os.Getenv("GOPATH"), "src", "jlambert", "FenixInception3", "FenixTestInstructionBuilder", "qml", "qtProject", "FenixTestInstructionBuilder", "main.qml")
	println(path)
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := widgets.NewQApplication(len(os.Args), os.Args)

	//view := quick.NewQQuickView(nil)
	var view = initQQuickView(path)
	var qmlBridge = NewQmlBridge(nil)
	qmlBridge.ConnectSendToGo(func(data string) string {
		fmt.Println("go:", data)
		return "hello from go"
	})
	qmlBridge.ConnectGenerateGuid(generateGuid)

	view.RootContext().SetContextProperty("QmlBridge", qmlBridge)
	view.SetSource(core.NewQUrl3(path, 0))

	view.SetTitle("goroutine Example")
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	//view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))
	view.Show()

	// Tell QML that it should use Go-server for instead of local mocks for certain functions
	qmlBridge.SendToQml(true)

	app.Exec()
}
