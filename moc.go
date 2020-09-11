

package main

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
std_core "github.com/therecipe/qt/core"
"github.com/therecipe/qt"
"strings"
"unsafe"
)
func cGoFreePacked(ptr unsafe.Pointer) { std_core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_Moc_PackedString) string { defer cGoFreePacked(s.ptr)
 if int(s.len) == -1 {
 return C.GoString(s.data)
 }
 return C.GoStringN(s.data, C.int(s.len)) }
func cGoUnpackBytes(s C.struct_Moc_PackedString) []byte { defer cGoFreePacked(s.ptr)
 if int(s.len) == -1 {
 gs := C.GoString(s.data)
 return []byte(gs)
 }
 return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len)) }
func unpackStringList(s string) []string {
if len(s) == 0 {
return make([]string, 0)
}
return strings.Split(s, "¡¦!")
}
type BridgeTemplate_ITF interface {
std_core.QObject_ITF
BridgeTemplatece8f65_PTR() *BridgeTemplate

}

func (ptr *BridgeTemplate) BridgeTemplate_PTR() *BridgeTemplate {
return ptr
}

func (ptr *BridgeTemplate) Pointer() unsafe.Pointer {
if ptr != nil {
return ptr.QObject_PTR().Pointer()
}
return nil
}

func (ptr *BridgeTemplate) SetPointer(p unsafe.Pointer) {
if ptr != nil{
ptr.QObject_PTR().SetPointer(p)
}
}

func PointerFromBridgeTemplate(ptr BridgeTemplate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.BridgeTemplate_PTR().Pointer()
	}
	return nil
}

func NewBridgeTemplateFromPointer(ptr unsafe.Pointer) (n *BridgeTemplate) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(BridgeTemplate)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
			case *BridgeTemplate:
				n = deduced

			case *std_core.QObject:
				n = &BridgeTemplate{QObject: *deduced }

			default:
				n = new(BridgeTemplate)
				n.SetPointer(ptr)
		}
	}
	return
}
//export callbackBridgeTemplatece8f65_Constructor
func callbackBridgeTemplatece8f65_Constructor(ptr unsafe.Pointer) {this := NewBridgeTemplateFromPointer(ptr)
qt.Register(ptr, this)
this.ConnectClicked(this.clicked)
this.ConnectSendVariantListMap(this.sendVariantListMap)
this.ConnectSendVariantListMap2(this.sendVariantListMap2)
}

//export callbackBridgeTemplatece8f65_Clicked
func  callbackBridgeTemplatece8f65_Clicked(ptr unsafe.Pointer) {
if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
(*(*func () )(signal))()
}

}

func (ptr *BridgeTemplate) ConnectClicked(f func () ) {
if ptr.Pointer() != nil {

if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
C.BridgeTemplatece8f65_ConnectClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "clicked")))
}

if signal := qt.LendSignal(ptr.Pointer(), "clicked"); signal != nil {
	f := func ()  {
(*(*func ())(signal))()
f()
}
	qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))} else {
	qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))}
}
}

func (ptr *BridgeTemplate) DisconnectClicked() {
if ptr.Pointer() != nil {
C.BridgeTemplatece8f65_DisconnectClicked(ptr.Pointer())
qt.DisconnectSignal(ptr.Pointer(), "clicked")
}
}

func (ptr *BridgeTemplate) Clicked() {
if ptr.Pointer() != nil {
C.BridgeTemplatece8f65_Clicked(ptr.Pointer())
}
}

//export callbackBridgeTemplatece8f65_SendVariantListMap
func  callbackBridgeTemplatece8f65_SendVariantListMap(ptr unsafe.Pointer, v0 unsafe.Pointer, v1 C.struct_Moc_PackedList, v2 C.struct_Moc_PackedList, v3 unsafe.Pointer) {
if signal := qt.GetSignal(ptr, "sendVariantListMap"); signal != nil {
(*(*func (*std_core.QVariant, []*std_core.QVariant, map[string]*std_core.QVariant, *std_core.QVariant) )(signal))(func()*std_core.QVariant{ tmpValue:=std_core.NewQVariantFromPointer(v0); qt.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant); return tmpValue; }(), func(l C.struct_Moc_PackedList)[]*std_core.QVariant{out := make([]*std_core.QVariant, int(l.len))
tmpList := NewBridgeTemplateFromPointer(l.data)
for i:=0;i<len(out);i++{ out[i] = tmpList.__sendVariantListMap_v1_atList(i) }
return out}(v1), func(l C.struct_Moc_PackedList)map[string]*std_core.QVariant{out := make(map[string]*std_core.QVariant, int(l.len))
tmpList := NewBridgeTemplateFromPointer(l.data)
for i,v:=range tmpList.__sendVariantListMap_v2_keyList(){ out[v] = tmpList.__sendVariantListMap_v2_atList(v, i) }
return out}(v2), func()*std_core.QVariant{ tmpValue:=std_core.NewQVariantFromPointer(v3); qt.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant); return tmpValue; }())
}

}

func (ptr *BridgeTemplate) ConnectSendVariantListMap(f func (v0 *std_core.QVariant, v1 []*std_core.QVariant, v2 map[string]*std_core.QVariant, v3 *std_core.QVariant) ) {
if ptr.Pointer() != nil {

if !qt.ExistsSignal(ptr.Pointer(), "sendVariantListMap") {
C.BridgeTemplatece8f65_ConnectSendVariantListMap(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "sendVariantListMap")))
}

if signal := qt.LendSignal(ptr.Pointer(), "sendVariantListMap"); signal != nil {
	f := func (v0 *std_core.QVariant, v1 []*std_core.QVariant, v2 map[string]*std_core.QVariant, v3 *std_core.QVariant)  {
(*(*func (*std_core.QVariant, []*std_core.QVariant, map[string]*std_core.QVariant, *std_core.QVariant))(signal))(v0, v1, v2, v3)
f(v0, v1, v2, v3)
}
	qt.ConnectSignal(ptr.Pointer(), "sendVariantListMap", unsafe.Pointer(&f))} else {
	qt.ConnectSignal(ptr.Pointer(), "sendVariantListMap", unsafe.Pointer(&f))}
}
}

func (ptr *BridgeTemplate) DisconnectSendVariantListMap() {
if ptr.Pointer() != nil {
C.BridgeTemplatece8f65_DisconnectSendVariantListMap(ptr.Pointer())
qt.DisconnectSignal(ptr.Pointer(), "sendVariantListMap")
}
}

func (ptr *BridgeTemplate) SendVariantListMap(v0 std_core.QVariant_ITF, v1 []*std_core.QVariant, v2 map[string]*std_core.QVariant, v3 std_core.QVariant_ITF) {
if ptr.Pointer() != nil {
C.BridgeTemplatece8f65_SendVariantListMap(ptr.Pointer(), std_core.PointerFromQVariant(v0), func() unsafe.Pointer {
tmpList := NewBridgeTemplateFromPointer(NewBridgeTemplateFromPointer(nil).__sendVariantListMap_v1_newList())
for _,v := range v1{
tmpList.__sendVariantListMap_v1_setList(v)
}
return tmpList.Pointer()
}(), func() unsafe.Pointer {
tmpList := NewBridgeTemplateFromPointer(NewBridgeTemplateFromPointer(nil).__sendVariantListMap_v2_newList())
for k,v := range v2{
tmpList.__sendVariantListMap_v2_setList(k, std_core.NewQVariant1(v))
}
return tmpList.Pointer()
}(), std_core.PointerFromQVariant(v3))
}
}

//export callbackBridgeTemplatece8f65_SendVariantListMap2
func  callbackBridgeTemplatece8f65_SendVariantListMap2(ptr unsafe.Pointer, *core. unsafe.Pointer) {
if signal := qt.GetSignal(ptr, "sendVariantListMap2"); signal != nil {
(*(*func (*std_core.QObject) )(signal))(std_core.NewQObjectFromPointer(*core.))
}

}

func (ptr *BridgeTemplate) ConnectSendVariantListMap2(f func (*core. *std_core.QObject) ) {
if ptr.Pointer() != nil {

if !qt.ExistsSignal(ptr.Pointer(), "sendVariantListMap2") {
C.BridgeTemplatece8f65_ConnectSendVariantListMap2(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "sendVariantListMap2")))
}

if signal := qt.LendSignal(ptr.Pointer(), "sendVariantListMap2"); signal != nil {
	f := func (*core. *std_core.QObject)  {
(*(*func (*std_core.QObject))(signal))(*core.)
f(*core.)
}
	qt.ConnectSignal(ptr.Pointer(), "sendVariantListMap2", unsafe.Pointer(&f))} else {
	qt.ConnectSignal(ptr.Pointer(), "sendVariantListMap2", unsafe.Pointer(&f))}
}
}

func (ptr *BridgeTemplate) DisconnectSendVariantListMap2() {
if ptr.Pointer() != nil {
C.BridgeTemplatece8f65_DisconnectSendVariantListMap2(ptr.Pointer())
qt.DisconnectSignal(ptr.Pointer(), "sendVariantListMap2")
}
}

func (ptr *BridgeTemplate) SendVariantListMap2(*core. std_core.QObject_ITF) {
if ptr.Pointer() != nil {
C.BridgeTemplatece8f65_SendVariantListMap2(ptr.Pointer(), std_core.PointerFromQObject(*core.))
}
}

func  BridgeTemplate_QRegisterMetaType() int{
return int(int32(C.BridgeTemplatece8f65_BridgeTemplatece8f65_QRegisterMetaType()))
}

func (ptr *BridgeTemplate) QRegisterMetaType() int{
return int(int32(C.BridgeTemplatece8f65_BridgeTemplatece8f65_QRegisterMetaType()))
}

func  BridgeTemplate_QRegisterMetaType2(typeName string) int{
var typeNameC *C.char
if typeName != "" {
typeNameC = C.CString(typeName)
defer C.free(unsafe.Pointer(typeNameC))
}
return int(int32(C.BridgeTemplatece8f65_BridgeTemplatece8f65_QRegisterMetaType2(typeNameC)))
}

func (ptr *BridgeTemplate) QRegisterMetaType2(typeName string) int{
var typeNameC *C.char
if typeName != "" {
typeNameC = C.CString(typeName)
defer C.free(unsafe.Pointer(typeNameC))
}
return int(int32(C.BridgeTemplatece8f65_BridgeTemplatece8f65_QRegisterMetaType2(typeNameC)))
}

func  BridgeTemplate_QmlRegisterType() int{
return int(int32(C.BridgeTemplatece8f65_BridgeTemplatece8f65_QmlRegisterType()))
}

func (ptr *BridgeTemplate) QmlRegisterType() int{
return int(int32(C.BridgeTemplatece8f65_BridgeTemplatece8f65_QmlRegisterType()))
}

func  BridgeTemplate_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int{
var uriC *C.char
if uri != "" {
uriC = C.CString(uri)
defer C.free(unsafe.Pointer(uriC))
}
var qmlNameC *C.char
if qmlName != "" {
qmlNameC = C.CString(qmlName)
defer C.free(unsafe.Pointer(qmlNameC))
}
return int(int32(C.BridgeTemplatece8f65_BridgeTemplatece8f65_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *BridgeTemplate) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int{
var uriC *C.char
if uri != "" {
uriC = C.CString(uri)
defer C.free(unsafe.Pointer(uriC))
}
var qmlNameC *C.char
if qmlName != "" {
qmlNameC = C.CString(qmlName)
defer C.free(unsafe.Pointer(qmlNameC))
}
return int(int32(C.BridgeTemplatece8f65_BridgeTemplatece8f65_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *BridgeTemplate) __children_atList(i int) *std_core.QObject{
if ptr.Pointer() != nil {
tmpValue := std_core.NewQObjectFromPointer(C.BridgeTemplatece8f65___children_atList(ptr.Pointer(), C.int(int32(i))))
if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
tmpValue.ConnectDestroyed(func(*std_core.QObject){ tmpValue.SetPointer(nil) })
}
return tmpValue
}
return nil
}

func (ptr *BridgeTemplate) __children_setList(i std_core.QObject_ITF) {
if ptr.Pointer() != nil {
C.BridgeTemplatece8f65___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
}
}

func (ptr *BridgeTemplate) __children_newList() unsafe.Pointer{
return C.BridgeTemplatece8f65___children_newList(ptr.Pointer())
}

func (ptr *BridgeTemplate) __dynamicPropertyNames_atList(i int) *std_core.QByteArray{
if ptr.Pointer() != nil {
tmpValue := std_core.NewQByteArrayFromPointer(C.BridgeTemplatece8f65___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
qt.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
return tmpValue
}
return nil
}

func (ptr *BridgeTemplate) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
if ptr.Pointer() != nil {
C.BridgeTemplatece8f65___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
}
}

func (ptr *BridgeTemplate) __dynamicPropertyNames_newList() unsafe.Pointer{
return C.BridgeTemplatece8f65___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *BridgeTemplate) __findChildren_atList(i int) *std_core.QObject{
if ptr.Pointer() != nil {
tmpValue := std_core.NewQObjectFromPointer(C.BridgeTemplatece8f65___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
tmpValue.ConnectDestroyed(func(*std_core.QObject){ tmpValue.SetPointer(nil) })
}
return tmpValue
}
return nil
}

func (ptr *BridgeTemplate) __findChildren_setList(i std_core.QObject_ITF) {
if ptr.Pointer() != nil {
C.BridgeTemplatece8f65___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
}
}

func (ptr *BridgeTemplate) __findChildren_newList() unsafe.Pointer{
return C.BridgeTemplatece8f65___findChildren_newList(ptr.Pointer())
}

func (ptr *BridgeTemplate) __findChildren_atList3(i int) *std_core.QObject{
if ptr.Pointer() != nil {
tmpValue := std_core.NewQObjectFromPointer(C.BridgeTemplatece8f65___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
tmpValue.ConnectDestroyed(func(*std_core.QObject){ tmpValue.SetPointer(nil) })
}
return tmpValue
}
return nil
}

func (ptr *BridgeTemplate) __findChildren_setList3(i std_core.QObject_ITF) {
if ptr.Pointer() != nil {
C.BridgeTemplatece8f65___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
}
}

func (ptr *BridgeTemplate) __findChildren_newList3() unsafe.Pointer{
return C.BridgeTemplatece8f65___findChildren_newList3(ptr.Pointer())
}

func (ptr *BridgeTemplate) __sendVariantListMap_v1_atList(i int) *std_core.QVariant{
if ptr.Pointer() != nil {
tmpValue := std_core.NewQVariantFromPointer(C.BridgeTemplatece8f65___sendVariantListMap_v1_atList(ptr.Pointer(), C.int(int32(i))))
qt.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
return tmpValue
}
return nil
}

func (ptr *BridgeTemplate) __sendVariantListMap_v1_setList(i std_core.QVariant_ITF) {
if ptr.Pointer() != nil {
C.BridgeTemplatece8f65___sendVariantListMap_v1_setList(ptr.Pointer(), std_core.PointerFromQVariant(i))
}
}

func (ptr *BridgeTemplate) __sendVariantListMap_v1_newList() unsafe.Pointer{
return C.BridgeTemplatece8f65___sendVariantListMap_v1_newList(ptr.Pointer())
}

func (ptr *BridgeTemplate) __sendVariantListMap_v2_atList(v string, i int) *std_core.QVariant{
if ptr.Pointer() != nil {
var vC *C.char
if v != "" {
vC = C.CString(v)
defer C.free(unsafe.Pointer(vC))
}
tmpValue := std_core.NewQVariantFromPointer(C.BridgeTemplatece8f65___sendVariantListMap_v2_atList(ptr.Pointer(), C.struct_Moc_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
qt.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
return tmpValue
}
return nil
}

func (ptr *BridgeTemplate) __sendVariantListMap_v2_setList(key string, i std_core.QVariant_ITF) {
if ptr.Pointer() != nil {
var keyC *C.char
if key != "" {
keyC = C.CString(key)
defer C.free(unsafe.Pointer(keyC))
}
C.BridgeTemplatece8f65___sendVariantListMap_v2_setList(ptr.Pointer(), C.struct_Moc_PackedString{data: keyC, len: C.longlong(len(key))}, std_core.PointerFromQVariant(i))
}
}

func (ptr *BridgeTemplate) __sendVariantListMap_v2_newList() unsafe.Pointer{
return C.BridgeTemplatece8f65___sendVariantListMap_v2_newList(ptr.Pointer())
}

func (ptr *BridgeTemplate) __sendVariantListMap_v2_keyList() []string{
if ptr.Pointer() != nil {
return func(l C.struct_Moc_PackedList)[]string{out := make([]string, int(l.len))
tmpList := NewBridgeTemplateFromPointer(l.data)
for i:=0;i<len(out);i++{ out[i] = tmpList.____sendVariantListMap_v2_keyList_atList(i) }
return out}(C.BridgeTemplatece8f65___sendVariantListMap_v2_keyList(ptr.Pointer()))
}
return make([]string, 0)
}

func (ptr *BridgeTemplate) ____sendVariantListMap_v2_keyList_atList(i int) string{
if ptr.Pointer() != nil {
return cGoUnpackString(C.BridgeTemplatece8f65_____sendVariantListMap_v2_keyList_atList(ptr.Pointer(), C.int(int32(i))))
}
return ""
}

func (ptr *BridgeTemplate) ____sendVariantListMap_v2_keyList_setList(i string) {
if ptr.Pointer() != nil {
var iC *C.char
if i != "" {
iC = C.CString(i)
defer C.free(unsafe.Pointer(iC))
}
C.BridgeTemplatece8f65_____sendVariantListMap_v2_keyList_setList(ptr.Pointer(), C.struct_Moc_PackedString{data: iC, len: C.longlong(len(i))})
}
}

func (ptr *BridgeTemplate) ____sendVariantListMap_v2_keyList_newList() unsafe.Pointer{
return C.BridgeTemplatece8f65_____sendVariantListMap_v2_keyList_newList(ptr.Pointer())
}

func  NewBridgeTemplate(parent std_core.QObject_ITF) *BridgeTemplate{
BridgeTemplatece8f65_QRegisterMetaType()
tmpValue := NewBridgeTemplateFromPointer(C.BridgeTemplatece8f65_NewBridgeTemplate(std_core.PointerFromQObject(parent)))
if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
tmpValue.ConnectDestroyed(func(*std_core.QObject){ tmpValue.SetPointer(nil) })
}
return tmpValue
}

//export callbackBridgeTemplatece8f65_DestroyBridgeTemplate
func  callbackBridgeTemplatece8f65_DestroyBridgeTemplate(ptr unsafe.Pointer) {
if signal := qt.GetSignal(ptr, "~BridgeTemplate"); signal != nil {
(*(*func () )(signal))()
}else{
NewBridgeTemplateFromPointer(ptr).DestroyBridgeTemplateDefault()
}
}

func (ptr *BridgeTemplate) ConnectDestroyBridgeTemplate(f func () ) {
if ptr.Pointer() != nil {

if signal := qt.LendSignal(ptr.Pointer(), "~BridgeTemplate"); signal != nil {
	f := func ()  {
(*(*func ())(signal))()
f()
}
	qt.ConnectSignal(ptr.Pointer(), "~BridgeTemplate", unsafe.Pointer(&f))} else {
	qt.ConnectSignal(ptr.Pointer(), "~BridgeTemplate", unsafe.Pointer(&f))}
}
}

func (ptr *BridgeTemplate) DisconnectDestroyBridgeTemplate() {
if ptr.Pointer() != nil {

qt.DisconnectSignal(ptr.Pointer(), "~BridgeTemplate")
}
}

func (ptr *BridgeTemplate) DestroyBridgeTemplate() {
if ptr.Pointer() != nil {

qt.SetFinalizer(ptr, nil)
C.BridgeTemplatece8f65_DestroyBridgeTemplate(ptr.Pointer())
ptr.SetPointer(nil)
}
}

func (ptr *BridgeTemplate) DestroyBridgeTemplateDefault() {
if ptr.Pointer() != nil {

qt.SetFinalizer(ptr, nil)
C.BridgeTemplatece8f65_DestroyBridgeTemplateDefault(ptr.Pointer())
ptr.SetPointer(nil)
}
}

//export callbackBridgeTemplatece8f65_ChildEvent
func  callbackBridgeTemplatece8f65_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
(*(*func (*std_core.QChildEvent) )(signal))(std_core.NewQChildEventFromPointer(event))
}else{
NewBridgeTemplateFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
}
}

func (ptr *BridgeTemplate) ChildEventDefault(event std_core.QChildEvent_ITF) {
if ptr.Pointer() != nil {
C.BridgeTemplatece8f65_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
}
}

//export callbackBridgeTemplatece8f65_ConnectNotify
func  callbackBridgeTemplatece8f65_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
(*(*func (*std_core.QMetaMethod) )(signal))(std_core.NewQMetaMethodFromPointer(sign))
}else{
NewBridgeTemplateFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
}
}

func (ptr *BridgeTemplate) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
if ptr.Pointer() != nil {
C.BridgeTemplatece8f65_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
}
}

//export callbackBridgeTemplatece8f65_CustomEvent
func  callbackBridgeTemplatece8f65_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
(*(*func (*std_core.QEvent) )(signal))(std_core.NewQEventFromPointer(event))
}else{
NewBridgeTemplateFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
}
}

func (ptr *BridgeTemplate) CustomEventDefault(event std_core.QEvent_ITF) {
if ptr.Pointer() != nil {
C.BridgeTemplatece8f65_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
}
}

//export callbackBridgeTemplatece8f65_DeleteLater
func  callbackBridgeTemplatece8f65_DeleteLater(ptr unsafe.Pointer) {
if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
(*(*func () )(signal))()
}else{
NewBridgeTemplateFromPointer(ptr).DeleteLaterDefault()
}
}

func (ptr *BridgeTemplate) DeleteLaterDefault() {
if ptr.Pointer() != nil {

qt.SetFinalizer(ptr, nil)
C.BridgeTemplatece8f65_DeleteLaterDefault(ptr.Pointer())
}
}

//export callbackBridgeTemplatece8f65_Destroyed
func  callbackBridgeTemplatece8f65_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
(*(*func (*std_core.QObject) )(signal))(std_core.NewQObjectFromPointer(obj))
}
qt.Unregister(ptr)

}

//export callbackBridgeTemplatece8f65_DisconnectNotify
func  callbackBridgeTemplatece8f65_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
(*(*func (*std_core.QMetaMethod) )(signal))(std_core.NewQMetaMethodFromPointer(sign))
}else{
NewBridgeTemplateFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
}
}

func (ptr *BridgeTemplate) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
if ptr.Pointer() != nil {
C.BridgeTemplatece8f65_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
}
}

//export callbackBridgeTemplatece8f65_Event
func  callbackBridgeTemplatece8f65_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char{
if signal := qt.GetSignal(ptr, "event"); signal != nil {
return C.char(int8(qt.GoBoolToInt((*(*func (*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
}

return C.char(int8(qt.GoBoolToInt(NewBridgeTemplateFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *BridgeTemplate) EventDefault(e std_core.QEvent_ITF) bool{
if ptr.Pointer() != nil {
return int8(C.BridgeTemplatece8f65_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
}
return false
}

//export callbackBridgeTemplatece8f65_EventFilter
func  callbackBridgeTemplatece8f65_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char{
if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
return C.char(int8(qt.GoBoolToInt((*(*func (*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

return C.char(int8(qt.GoBoolToInt(NewBridgeTemplateFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *BridgeTemplate) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool{
if ptr.Pointer() != nil {
return int8(C.BridgeTemplatece8f65_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
}
return false
}



//export callbackBridgeTemplatece8f65_ObjectNameChanged
func  callbackBridgeTemplatece8f65_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
(*(*func (string) )(signal))(cGoUnpackString(objectName))
}

}

//export callbackBridgeTemplatece8f65_TimerEvent
func  callbackBridgeTemplatece8f65_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
(*(*func (*std_core.QTimerEvent) )(signal))(std_core.NewQTimerEventFromPointer(event))
}else{
NewBridgeTemplateFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
}
}

func (ptr *BridgeTemplate) TimerEventDefault(event std_core.QTimerEvent_ITF) {
if ptr.Pointer() != nil {
C.BridgeTemplatece8f65_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
}
}

type CustomLabel_ITF interface {
std_core.QObject_ITF
CustomLabelce8f65_PTR() *CustomLabel

}

func (ptr *CustomLabel) CustomLabel_PTR() *CustomLabel {
return ptr
}

func (ptr *CustomLabel) Pointer() unsafe.Pointer {
if ptr != nil {
return ptr.QObject_PTR().Pointer()
}
return nil
}

func (ptr *CustomLabel) SetPointer(p unsafe.Pointer) {
if ptr != nil{
ptr.QObject_PTR().SetPointer(p)
}
}

func PointerFromCustomLabel(ptr CustomLabel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.CustomLabel_PTR().Pointer()
	}
	return nil
}

func NewCustomLabelFromPointer(ptr unsafe.Pointer) (n *CustomLabel) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(CustomLabel)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
			case *CustomLabel:
				n = deduced

			case *std_core.QObject:
				n = &CustomLabel{QObject: *deduced }

			default:
				n = new(CustomLabel)
				n.SetPointer(ptr)
		}
	}
	return
}
func (this *CustomLabel) Init() { this.init() }
//export callbackCustomLabelce8f65_Constructor
func callbackCustomLabelce8f65_Constructor(ptr unsafe.Pointer) {this := NewCustomLabelFromPointer(ptr)
qt.Register(ptr, this)
this.init()
}

//export callbackCustomLabelce8f65_Text
func  callbackCustomLabelce8f65_Text(ptr unsafe.Pointer) C.struct_Moc_PackedString{
if signal := qt.GetSignal(ptr, "text"); signal != nil {
tempVal := (*(*func () string)(signal))()
return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}
tempVal := NewCustomLabelFromPointer(ptr).TextDefault()
return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *CustomLabel) ConnectText(f func () string) {
if ptr.Pointer() != nil {

if signal := qt.LendSignal(ptr.Pointer(), "text"); signal != nil {
	f := func () string {
(*(*func ()string)(signal))()
return f()
}
	qt.ConnectSignal(ptr.Pointer(), "text", unsafe.Pointer(&f))} else {
	qt.ConnectSignal(ptr.Pointer(), "text", unsafe.Pointer(&f))}
}
}

func (ptr *CustomLabel) DisconnectText() {
if ptr.Pointer() != nil {

qt.DisconnectSignal(ptr.Pointer(), "text")
}
}

func (ptr *CustomLabel) Text() string{
if ptr.Pointer() != nil {
return cGoUnpackString(C.CustomLabelce8f65_Text(ptr.Pointer()))
}
return ""
}

func (ptr *CustomLabel) TextDefault() string{
if ptr.Pointer() != nil {
return cGoUnpackString(C.CustomLabelce8f65_TextDefault(ptr.Pointer()))
}
return ""
}

//export callbackCustomLabelce8f65_SetText
func  callbackCustomLabelce8f65_SetText(ptr unsafe.Pointer, text C.struct_Moc_PackedString) {
if signal := qt.GetSignal(ptr, "setText"); signal != nil {
(*(*func (string) )(signal))(cGoUnpackString(text))
}else{
NewCustomLabelFromPointer(ptr).SetTextDefault(cGoUnpackString(text))
}
}

func (ptr *CustomLabel) ConnectSetText(f func (text string) ) {
if ptr.Pointer() != nil {

if signal := qt.LendSignal(ptr.Pointer(), "setText"); signal != nil {
	f := func (text string)  {
(*(*func (string))(signal))(text)
f(text)
}
	qt.ConnectSignal(ptr.Pointer(), "setText", unsafe.Pointer(&f))} else {
	qt.ConnectSignal(ptr.Pointer(), "setText", unsafe.Pointer(&f))}
}
}

func (ptr *CustomLabel) DisconnectSetText() {
if ptr.Pointer() != nil {

qt.DisconnectSignal(ptr.Pointer(), "setText")
}
}

func (ptr *CustomLabel) SetText(text string) {
if ptr.Pointer() != nil {
var textC *C.char
if text != "" {
textC = C.CString(text)
defer C.free(unsafe.Pointer(textC))
}
C.CustomLabelce8f65_SetText(ptr.Pointer(), C.struct_Moc_PackedString{data: textC, len: C.longlong(len(text))})
}
}

func (ptr *CustomLabel) SetTextDefault(text string) {
if ptr.Pointer() != nil {
var textC *C.char
if text != "" {
textC = C.CString(text)
defer C.free(unsafe.Pointer(textC))
}
C.CustomLabelce8f65_SetTextDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: textC, len: C.longlong(len(text))})
}
}

//export callbackCustomLabelce8f65_TextChanged
func  callbackCustomLabelce8f65_TextChanged(ptr unsafe.Pointer, text C.struct_Moc_PackedString) {
if signal := qt.GetSignal(ptr, "textChanged"); signal != nil {
(*(*func (string) )(signal))(cGoUnpackString(text))
}

}

func (ptr *CustomLabel) ConnectTextChanged(f func (text string) ) {
if ptr.Pointer() != nil {

if !qt.ExistsSignal(ptr.Pointer(), "textChanged") {
C.CustomLabelce8f65_ConnectTextChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "textChanged")))
}

if signal := qt.LendSignal(ptr.Pointer(), "textChanged"); signal != nil {
	f := func (text string)  {
(*(*func (string))(signal))(text)
f(text)
}
	qt.ConnectSignal(ptr.Pointer(), "textChanged", unsafe.Pointer(&f))} else {
	qt.ConnectSignal(ptr.Pointer(), "textChanged", unsafe.Pointer(&f))}
}
}

func (ptr *CustomLabel) DisconnectTextChanged() {
if ptr.Pointer() != nil {
C.CustomLabelce8f65_DisconnectTextChanged(ptr.Pointer())
qt.DisconnectSignal(ptr.Pointer(), "textChanged")
}
}

func (ptr *CustomLabel) TextChanged(text string) {
if ptr.Pointer() != nil {
var textC *C.char
if text != "" {
textC = C.CString(text)
defer C.free(unsafe.Pointer(textC))
}
C.CustomLabelce8f65_TextChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: textC, len: C.longlong(len(text))})
}
}

func  CustomLabel_QRegisterMetaType() int{
return int(int32(C.CustomLabelce8f65_CustomLabelce8f65_QRegisterMetaType()))
}

func (ptr *CustomLabel) QRegisterMetaType() int{
return int(int32(C.CustomLabelce8f65_CustomLabelce8f65_QRegisterMetaType()))
}

func  CustomLabel_QRegisterMetaType2(typeName string) int{
var typeNameC *C.char
if typeName != "" {
typeNameC = C.CString(typeName)
defer C.free(unsafe.Pointer(typeNameC))
}
return int(int32(C.CustomLabelce8f65_CustomLabelce8f65_QRegisterMetaType2(typeNameC)))
}

func (ptr *CustomLabel) QRegisterMetaType2(typeName string) int{
var typeNameC *C.char
if typeName != "" {
typeNameC = C.CString(typeName)
defer C.free(unsafe.Pointer(typeNameC))
}
return int(int32(C.CustomLabelce8f65_CustomLabelce8f65_QRegisterMetaType2(typeNameC)))
}

func  CustomLabel_QmlRegisterType() int{
return int(int32(C.CustomLabelce8f65_CustomLabelce8f65_QmlRegisterType()))
}

func (ptr *CustomLabel) QmlRegisterType() int{
return int(int32(C.CustomLabelce8f65_CustomLabelce8f65_QmlRegisterType()))
}

func  CustomLabel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int{
var uriC *C.char
if uri != "" {
uriC = C.CString(uri)
defer C.free(unsafe.Pointer(uriC))
}
var qmlNameC *C.char
if qmlName != "" {
qmlNameC = C.CString(qmlName)
defer C.free(unsafe.Pointer(qmlNameC))
}
return int(int32(C.CustomLabelce8f65_CustomLabelce8f65_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *CustomLabel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int{
var uriC *C.char
if uri != "" {
uriC = C.CString(uri)
defer C.free(unsafe.Pointer(uriC))
}
var qmlNameC *C.char
if qmlName != "" {
qmlNameC = C.CString(qmlName)
defer C.free(unsafe.Pointer(qmlNameC))
}
return int(int32(C.CustomLabelce8f65_CustomLabelce8f65_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *CustomLabel) __children_atList(i int) *std_core.QObject{
if ptr.Pointer() != nil {
tmpValue := std_core.NewQObjectFromPointer(C.CustomLabelce8f65___children_atList(ptr.Pointer(), C.int(int32(i))))
if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
tmpValue.ConnectDestroyed(func(*std_core.QObject){ tmpValue.SetPointer(nil) })
}
return tmpValue
}
return nil
}

func (ptr *CustomLabel) __children_setList(i std_core.QObject_ITF) {
if ptr.Pointer() != nil {
C.CustomLabelce8f65___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
}
}

func (ptr *CustomLabel) __children_newList() unsafe.Pointer{
return C.CustomLabelce8f65___children_newList(ptr.Pointer())
}

func (ptr *CustomLabel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray{
if ptr.Pointer() != nil {
tmpValue := std_core.NewQByteArrayFromPointer(C.CustomLabelce8f65___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
qt.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
return tmpValue
}
return nil
}

func (ptr *CustomLabel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
if ptr.Pointer() != nil {
C.CustomLabelce8f65___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
}
}

func (ptr *CustomLabel) __dynamicPropertyNames_newList() unsafe.Pointer{
return C.CustomLabelce8f65___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *CustomLabel) __findChildren_atList(i int) *std_core.QObject{
if ptr.Pointer() != nil {
tmpValue := std_core.NewQObjectFromPointer(C.CustomLabelce8f65___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
tmpValue.ConnectDestroyed(func(*std_core.QObject){ tmpValue.SetPointer(nil) })
}
return tmpValue
}
return nil
}

func (ptr *CustomLabel) __findChildren_setList(i std_core.QObject_ITF) {
if ptr.Pointer() != nil {
C.CustomLabelce8f65___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
}
}

func (ptr *CustomLabel) __findChildren_newList() unsafe.Pointer{
return C.CustomLabelce8f65___findChildren_newList(ptr.Pointer())
}

func (ptr *CustomLabel) __findChildren_atList3(i int) *std_core.QObject{
if ptr.Pointer() != nil {
tmpValue := std_core.NewQObjectFromPointer(C.CustomLabelce8f65___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
tmpValue.ConnectDestroyed(func(*std_core.QObject){ tmpValue.SetPointer(nil) })
}
return tmpValue
}
return nil
}

func (ptr *CustomLabel) __findChildren_setList3(i std_core.QObject_ITF) {
if ptr.Pointer() != nil {
C.CustomLabelce8f65___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
}
}

func (ptr *CustomLabel) __findChildren_newList3() unsafe.Pointer{
return C.CustomLabelce8f65___findChildren_newList3(ptr.Pointer())
}

func  NewCustomLabel(parent std_core.QObject_ITF) *CustomLabel{
CustomLabelce8f65_QRegisterMetaType()
tmpValue := NewCustomLabelFromPointer(C.CustomLabelce8f65_NewCustomLabel(std_core.PointerFromQObject(parent)))
if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
tmpValue.ConnectDestroyed(func(*std_core.QObject){ tmpValue.SetPointer(nil) })
}
return tmpValue
}

//export callbackCustomLabelce8f65_DestroyCustomLabel
func  callbackCustomLabelce8f65_DestroyCustomLabel(ptr unsafe.Pointer) {
if signal := qt.GetSignal(ptr, "~CustomLabel"); signal != nil {
(*(*func () )(signal))()
}else{
NewCustomLabelFromPointer(ptr).DestroyCustomLabelDefault()
}
}

func (ptr *CustomLabel) ConnectDestroyCustomLabel(f func () ) {
if ptr.Pointer() != nil {

if signal := qt.LendSignal(ptr.Pointer(), "~CustomLabel"); signal != nil {
	f := func ()  {
(*(*func ())(signal))()
f()
}
	qt.ConnectSignal(ptr.Pointer(), "~CustomLabel", unsafe.Pointer(&f))} else {
	qt.ConnectSignal(ptr.Pointer(), "~CustomLabel", unsafe.Pointer(&f))}
}
}

func (ptr *CustomLabel) DisconnectDestroyCustomLabel() {
if ptr.Pointer() != nil {

qt.DisconnectSignal(ptr.Pointer(), "~CustomLabel")
}
}

func (ptr *CustomLabel) DestroyCustomLabel() {
if ptr.Pointer() != nil {

qt.SetFinalizer(ptr, nil)
C.CustomLabelce8f65_DestroyCustomLabel(ptr.Pointer())
ptr.SetPointer(nil)
}
}

func (ptr *CustomLabel) DestroyCustomLabelDefault() {
if ptr.Pointer() != nil {

qt.SetFinalizer(ptr, nil)
C.CustomLabelce8f65_DestroyCustomLabelDefault(ptr.Pointer())
ptr.SetPointer(nil)
}
}

//export callbackCustomLabelce8f65_ChildEvent
func  callbackCustomLabelce8f65_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
(*(*func (*std_core.QChildEvent) )(signal))(std_core.NewQChildEventFromPointer(event))
}else{
NewCustomLabelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
}
}

func (ptr *CustomLabel) ChildEventDefault(event std_core.QChildEvent_ITF) {
if ptr.Pointer() != nil {
C.CustomLabelce8f65_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
}
}

//export callbackCustomLabelce8f65_ConnectNotify
func  callbackCustomLabelce8f65_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
(*(*func (*std_core.QMetaMethod) )(signal))(std_core.NewQMetaMethodFromPointer(sign))
}else{
NewCustomLabelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
}
}

func (ptr *CustomLabel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
if ptr.Pointer() != nil {
C.CustomLabelce8f65_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
}
}

//export callbackCustomLabelce8f65_CustomEvent
func  callbackCustomLabelce8f65_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
(*(*func (*std_core.QEvent) )(signal))(std_core.NewQEventFromPointer(event))
}else{
NewCustomLabelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
}
}

func (ptr *CustomLabel) CustomEventDefault(event std_core.QEvent_ITF) {
if ptr.Pointer() != nil {
C.CustomLabelce8f65_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
}
}

//export callbackCustomLabelce8f65_DeleteLater
func  callbackCustomLabelce8f65_DeleteLater(ptr unsafe.Pointer) {
if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
(*(*func () )(signal))()
}else{
NewCustomLabelFromPointer(ptr).DeleteLaterDefault()
}
}

func (ptr *CustomLabel) DeleteLaterDefault() {
if ptr.Pointer() != nil {

qt.SetFinalizer(ptr, nil)
C.CustomLabelce8f65_DeleteLaterDefault(ptr.Pointer())
}
}

//export callbackCustomLabelce8f65_Destroyed
func  callbackCustomLabelce8f65_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
(*(*func (*std_core.QObject) )(signal))(std_core.NewQObjectFromPointer(obj))
}
qt.Unregister(ptr)

}

//export callbackCustomLabelce8f65_DisconnectNotify
func  callbackCustomLabelce8f65_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
(*(*func (*std_core.QMetaMethod) )(signal))(std_core.NewQMetaMethodFromPointer(sign))
}else{
NewCustomLabelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
}
}

func (ptr *CustomLabel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
if ptr.Pointer() != nil {
C.CustomLabelce8f65_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
}
}

//export callbackCustomLabelce8f65_Event
func  callbackCustomLabelce8f65_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char{
if signal := qt.GetSignal(ptr, "event"); signal != nil {
return C.char(int8(qt.GoBoolToInt((*(*func (*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
}

return C.char(int8(qt.GoBoolToInt(NewCustomLabelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *CustomLabel) EventDefault(e std_core.QEvent_ITF) bool{
if ptr.Pointer() != nil {
return int8(C.CustomLabelce8f65_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
}
return false
}

//export callbackCustomLabelce8f65_EventFilter
func  callbackCustomLabelce8f65_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char{
if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
return C.char(int8(qt.GoBoolToInt((*(*func (*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

return C.char(int8(qt.GoBoolToInt(NewCustomLabelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *CustomLabel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool{
if ptr.Pointer() != nil {
return int8(C.CustomLabelce8f65_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
}
return false
}



//export callbackCustomLabelce8f65_ObjectNameChanged
func  callbackCustomLabelce8f65_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
(*(*func (string) )(signal))(cGoUnpackString(objectName))
}

}

//export callbackCustomLabelce8f65_TimerEvent
func  callbackCustomLabelce8f65_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
(*(*func (*std_core.QTimerEvent) )(signal))(std_core.NewQTimerEventFromPointer(event))
}else{
NewCustomLabelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
}
}

func (ptr *CustomLabel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
if ptr.Pointer() != nil {
C.CustomLabelce8f65_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
}
}

func init() {
qt.ItfMap["main.BridgeTemplate_ITF"] = BridgeTemplate{}
qt.FuncMap["main.NewBridgeTemplate"] = NewBridgeTemplate
qt.ItfMap["main.CustomLabel_ITF"] = CustomLabel{}
qt.FuncMap["main.NewCustomLabel"] = NewCustomLabel
}
