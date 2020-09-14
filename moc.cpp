

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLayout>
#include <QMap>
#include <QMetaMethod>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QString>
#include <QTimerEvent>
#include <QVariant>
#include <QWidget>
#include <QWindow>

#ifdef QT_QML_LIB
	#include <QQmlEngine>
#endif


typedef QMap<QString, QVariant> type424d06;
class BridgeTemplatece8f65: public QObject
{
Q_OBJECT
public:
	BridgeTemplatece8f65(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");BridgeTemplatece8f65_BridgeTemplatece8f65_QRegisterMetaType();BridgeTemplatece8f65_BridgeTemplatece8f65_QRegisterMetaTypes();callbackBridgeTemplatece8f65_Constructor(this);};
	void Signal_Clicked() { callbackBridgeTemplatece8f65_Clicked(this); };
	void Signal_SendVariantListMap(QVariant v0, QList<QVariant> v1, type424d06 v2) { callbackBridgeTemplatece8f65_SendVariantListMap(this, new QVariant(v0), ({ QList<QVariant>* tmpValue5a6df7 = new QList<QVariant>(v1); Moc_PackedList { tmpValue5a6df7, tmpValue5a6df7->size() }; }), ({ QMap<QString, QVariant>* tmpValuea1047e = new QMap<QString, QVariant>(v2); Moc_PackedList { tmpValuea1047e, tmpValuea1047e->size() }; })); };
	 ~BridgeTemplatece8f65() { callbackBridgeTemplatece8f65_DestroyBridgeTemplate(this); };
	void childEvent(QChildEvent * event) { callbackBridgeTemplatece8f65_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackBridgeTemplatece8f65_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackBridgeTemplatece8f65_CustomEvent(this, event); };
	void deleteLater() { callbackBridgeTemplatece8f65_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackBridgeTemplatece8f65_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackBridgeTemplatece8f65_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackBridgeTemplatece8f65_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackBridgeTemplatece8f65_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackBridgeTemplatece8f65_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackBridgeTemplatece8f65_TimerEvent(this, event); };
signals:
	void clicked();
	void sendVariantListMap(QVariant v0, QList<QVariant> v1, type424d06 v2);
public slots:
private:
};

Q_DECLARE_METATYPE(BridgeTemplatece8f65*)


void BridgeTemplatece8f65_BridgeTemplatece8f65_QRegisterMetaTypes() {
	qRegisterMetaType<type424d06>("type424d06");
}

class CustomLabelce8f65: public QObject
{
Q_OBJECT
Q_PROPERTY(QString text READ text WRITE setText NOTIFY textChanged)
public:
	CustomLabelce8f65(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");CustomLabelce8f65_CustomLabelce8f65_QRegisterMetaType();CustomLabelce8f65_CustomLabelce8f65_QRegisterMetaTypes();callbackCustomLabelce8f65_Constructor(this);};
	QString text() { return ({ Moc_PackedString tempVal = callbackCustomLabelce8f65_Text(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setText(QString text) { QByteArray* t372ea0 = new QByteArray(text.toUtf8()); Moc_PackedString textPacked = { const_cast<char*>(t372ea0->prepend("WHITESPACE").constData()+10), t372ea0->size()-10, t372ea0 };callbackCustomLabelce8f65_SetText(this, textPacked); };
	void Signal_TextChanged(QString text) { QByteArray* t372ea0 = new QByteArray(text.toUtf8()); Moc_PackedString textPacked = { const_cast<char*>(t372ea0->prepend("WHITESPACE").constData()+10), t372ea0->size()-10, t372ea0 };callbackCustomLabelce8f65_TextChanged(this, textPacked); };
	 ~CustomLabelce8f65() { callbackCustomLabelce8f65_DestroyCustomLabel(this); };
	void childEvent(QChildEvent * event) { callbackCustomLabelce8f65_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackCustomLabelce8f65_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackCustomLabelce8f65_CustomEvent(this, event); };
	void deleteLater() { callbackCustomLabelce8f65_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackCustomLabelce8f65_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackCustomLabelce8f65_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackCustomLabelce8f65_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackCustomLabelce8f65_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackCustomLabelce8f65_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackCustomLabelce8f65_TimerEvent(this, event); };
	QString textDefault() { return _text; };
	void setTextDefault(QString p) { if (p != _text) { _text = p; textChanged(_text); } };
signals:
	void textChanged(QString text);
public slots:
private:
	QString _text;
};

Q_DECLARE_METATYPE(CustomLabelce8f65*)


void CustomLabelce8f65_CustomLabelce8f65_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

void BridgeTemplatece8f65_ConnectClicked(void* ptr, long long t)
{
	QObject::connect(static_cast<BridgeTemplatece8f65*>(ptr), static_cast<void (BridgeTemplatece8f65::*)()>(&BridgeTemplatece8f65::clicked), static_cast<BridgeTemplatece8f65*>(ptr), static_cast<void (BridgeTemplatece8f65::*)()>(&BridgeTemplatece8f65::Signal_Clicked), static_cast<Qt::ConnectionType>(t));
}

void BridgeTemplatece8f65_DisconnectClicked(void* ptr)
{
	QObject::disconnect(static_cast<BridgeTemplatece8f65*>(ptr), static_cast<void (BridgeTemplatece8f65::*)()>(&BridgeTemplatece8f65::clicked), static_cast<BridgeTemplatece8f65*>(ptr), static_cast<void (BridgeTemplatece8f65::*)()>(&BridgeTemplatece8f65::Signal_Clicked));
}

void BridgeTemplatece8f65_Clicked(void* ptr)
{
	static_cast<BridgeTemplatece8f65*>(ptr)->clicked();
}

void BridgeTemplatece8f65_ConnectSendVariantListMap(void* ptr, long long t)
{
	QObject::connect(static_cast<BridgeTemplatece8f65*>(ptr), static_cast<void (BridgeTemplatece8f65::*)(QVariant, QList<QVariant>, QMap<QString, QVariant>)>(&BridgeTemplatece8f65::sendVariantListMap), static_cast<BridgeTemplatece8f65*>(ptr), static_cast<void (BridgeTemplatece8f65::*)(QVariant, QList<QVariant>, QMap<QString, QVariant>)>(&BridgeTemplatece8f65::Signal_SendVariantListMap), static_cast<Qt::ConnectionType>(t));
}

void BridgeTemplatece8f65_DisconnectSendVariantListMap(void* ptr)
{
	QObject::disconnect(static_cast<BridgeTemplatece8f65*>(ptr), static_cast<void (BridgeTemplatece8f65::*)(QVariant, QList<QVariant>, QMap<QString, QVariant>)>(&BridgeTemplatece8f65::sendVariantListMap), static_cast<BridgeTemplatece8f65*>(ptr), static_cast<void (BridgeTemplatece8f65::*)(QVariant, QList<QVariant>, QMap<QString, QVariant>)>(&BridgeTemplatece8f65::Signal_SendVariantListMap));
}

void BridgeTemplatece8f65_SendVariantListMap(void* ptr, void* v0, void* v1, void* v2)
{
	static_cast<BridgeTemplatece8f65*>(ptr)->sendVariantListMap(*static_cast<QVariant*>(v0), ({ QList<QVariant>* tmpP = static_cast<QList<QVariant>*>(v1); QList<QVariant> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }), ({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(v2); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

int BridgeTemplatece8f65_BridgeTemplatece8f65_QRegisterMetaType()
{
	return qRegisterMetaType<BridgeTemplatece8f65*>();
}

int BridgeTemplatece8f65_BridgeTemplatece8f65_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<BridgeTemplatece8f65*>(const_cast<const char*>(typeName));
}

int BridgeTemplatece8f65_BridgeTemplatece8f65_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<BridgeTemplatece8f65>();
#else
	return 0;
#endif
}

int BridgeTemplatece8f65_BridgeTemplatece8f65_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<BridgeTemplatece8f65>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int BridgeTemplatece8f65_BridgeTemplatece8f65_QmlRegisterUncreatableType(char* uri, int versionMajor, int versionMinor, char* qmlName, struct Moc_PackedString message)
{
#ifdef QT_QML_LIB
	return qmlRegisterUncreatableType<BridgeTemplatece8f65>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName), QString::fromUtf8(message.data, message.len));
#else
	return 0;
#endif
}

void* BridgeTemplatece8f65___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void BridgeTemplatece8f65___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* BridgeTemplatece8f65___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* BridgeTemplatece8f65___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void BridgeTemplatece8f65___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* BridgeTemplatece8f65___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* BridgeTemplatece8f65___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void BridgeTemplatece8f65___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* BridgeTemplatece8f65___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* BridgeTemplatece8f65___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void BridgeTemplatece8f65___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* BridgeTemplatece8f65___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* BridgeTemplatece8f65___sendVariantListMap_v1_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void BridgeTemplatece8f65___sendVariantListMap_v1_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* BridgeTemplatece8f65___sendVariantListMap_v1_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* BridgeTemplatece8f65___sendVariantListMap_v2_atList(void* ptr, struct Moc_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void BridgeTemplatece8f65___sendVariantListMap_v2_setList(void* ptr, struct Moc_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* BridgeTemplatece8f65___sendVariantListMap_v2_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct Moc_PackedList BridgeTemplatece8f65___sendVariantListMap_v2_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

struct Moc_PackedString BridgeTemplatece8f65_____sendVariantListMap_v2_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); Moc_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void BridgeTemplatece8f65_____sendVariantListMap_v2_keyList_setList(void* ptr, struct Moc_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* BridgeTemplatece8f65_____sendVariantListMap_v2_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

void* BridgeTemplatece8f65_NewBridgeTemplate(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new BridgeTemplatece8f65(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new BridgeTemplatece8f65(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new BridgeTemplatece8f65(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new BridgeTemplatece8f65(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new BridgeTemplatece8f65(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new BridgeTemplatece8f65(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new BridgeTemplatece8f65(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new BridgeTemplatece8f65(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new BridgeTemplatece8f65(static_cast<QWindow*>(parent));
	} else {
		return new BridgeTemplatece8f65(static_cast<QObject*>(parent));
	}
}

void BridgeTemplatece8f65_DestroyBridgeTemplate(void* ptr)
{
	static_cast<BridgeTemplatece8f65*>(ptr)->~BridgeTemplatece8f65();
}

void BridgeTemplatece8f65_DestroyBridgeTemplateDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void BridgeTemplatece8f65_ChildEventDefault(void* ptr, void* event)
{
	static_cast<BridgeTemplatece8f65*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void BridgeTemplatece8f65_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<BridgeTemplatece8f65*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void BridgeTemplatece8f65_CustomEventDefault(void* ptr, void* event)
{
	static_cast<BridgeTemplatece8f65*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void BridgeTemplatece8f65_DeleteLaterDefault(void* ptr)
{
	static_cast<BridgeTemplatece8f65*>(ptr)->QObject::deleteLater();
}

void BridgeTemplatece8f65_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<BridgeTemplatece8f65*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char BridgeTemplatece8f65_EventDefault(void* ptr, void* e)
{
	return static_cast<BridgeTemplatece8f65*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char BridgeTemplatece8f65_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<BridgeTemplatece8f65*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}



void BridgeTemplatece8f65_TimerEventDefault(void* ptr, void* event)
{
	static_cast<BridgeTemplatece8f65*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

struct Moc_PackedString CustomLabelce8f65_Text(void* ptr)
{
	return ({ QByteArray* t671b6e = new QByteArray(static_cast<CustomLabelce8f65*>(ptr)->text().toUtf8()); Moc_PackedString { const_cast<char*>(t671b6e->prepend("WHITESPACE").constData()+10), t671b6e->size()-10, t671b6e }; });
}

struct Moc_PackedString CustomLabelce8f65_TextDefault(void* ptr)
{
	return ({ QByteArray* t99e093 = new QByteArray(static_cast<CustomLabelce8f65*>(ptr)->textDefault().toUtf8()); Moc_PackedString { const_cast<char*>(t99e093->prepend("WHITESPACE").constData()+10), t99e093->size()-10, t99e093 }; });
}

void CustomLabelce8f65_SetText(void* ptr, struct Moc_PackedString text)
{
	static_cast<CustomLabelce8f65*>(ptr)->setText(QString::fromUtf8(text.data, text.len));
}

void CustomLabelce8f65_SetTextDefault(void* ptr, struct Moc_PackedString text)
{
	static_cast<CustomLabelce8f65*>(ptr)->setTextDefault(QString::fromUtf8(text.data, text.len));
}

void CustomLabelce8f65_ConnectTextChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<CustomLabelce8f65*>(ptr), static_cast<void (CustomLabelce8f65::*)(QString)>(&CustomLabelce8f65::textChanged), static_cast<CustomLabelce8f65*>(ptr), static_cast<void (CustomLabelce8f65::*)(QString)>(&CustomLabelce8f65::Signal_TextChanged), static_cast<Qt::ConnectionType>(t));
}

void CustomLabelce8f65_DisconnectTextChanged(void* ptr)
{
	QObject::disconnect(static_cast<CustomLabelce8f65*>(ptr), static_cast<void (CustomLabelce8f65::*)(QString)>(&CustomLabelce8f65::textChanged), static_cast<CustomLabelce8f65*>(ptr), static_cast<void (CustomLabelce8f65::*)(QString)>(&CustomLabelce8f65::Signal_TextChanged));
}

void CustomLabelce8f65_TextChanged(void* ptr, struct Moc_PackedString text)
{
	static_cast<CustomLabelce8f65*>(ptr)->textChanged(QString::fromUtf8(text.data, text.len));
}

int CustomLabelce8f65_CustomLabelce8f65_QRegisterMetaType()
{
	return qRegisterMetaType<CustomLabelce8f65*>();
}

int CustomLabelce8f65_CustomLabelce8f65_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<CustomLabelce8f65*>(const_cast<const char*>(typeName));
}

int CustomLabelce8f65_CustomLabelce8f65_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CustomLabelce8f65>();
#else
	return 0;
#endif
}

int CustomLabelce8f65_CustomLabelce8f65_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CustomLabelce8f65>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int CustomLabelce8f65_CustomLabelce8f65_QmlRegisterUncreatableType(char* uri, int versionMajor, int versionMinor, char* qmlName, struct Moc_PackedString message)
{
#ifdef QT_QML_LIB
	return qmlRegisterUncreatableType<CustomLabelce8f65>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName), QString::fromUtf8(message.data, message.len));
#else
	return 0;
#endif
}

void* CustomLabelce8f65___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomLabelce8f65___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomLabelce8f65___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* CustomLabelce8f65___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomLabelce8f65___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* CustomLabelce8f65___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* CustomLabelce8f65___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomLabelce8f65___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomLabelce8f65___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomLabelce8f65___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomLabelce8f65___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomLabelce8f65___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomLabelce8f65_NewCustomLabel(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new CustomLabelce8f65(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new CustomLabelce8f65(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new CustomLabelce8f65(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new CustomLabelce8f65(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new CustomLabelce8f65(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new CustomLabelce8f65(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new CustomLabelce8f65(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new CustomLabelce8f65(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new CustomLabelce8f65(static_cast<QWindow*>(parent));
	} else {
		return new CustomLabelce8f65(static_cast<QObject*>(parent));
	}
}

void CustomLabelce8f65_DestroyCustomLabel(void* ptr)
{
	static_cast<CustomLabelce8f65*>(ptr)->~CustomLabelce8f65();
}

void CustomLabelce8f65_DestroyCustomLabelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void CustomLabelce8f65_ChildEventDefault(void* ptr, void* event)
{
	static_cast<CustomLabelce8f65*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void CustomLabelce8f65_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CustomLabelce8f65*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void CustomLabelce8f65_CustomEventDefault(void* ptr, void* event)
{
	static_cast<CustomLabelce8f65*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void CustomLabelce8f65_DeleteLaterDefault(void* ptr)
{
	static_cast<CustomLabelce8f65*>(ptr)->QObject::deleteLater();
}

void CustomLabelce8f65_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CustomLabelce8f65*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char CustomLabelce8f65_EventDefault(void* ptr, void* e)
{
	return static_cast<CustomLabelce8f65*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char CustomLabelce8f65_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<CustomLabelce8f65*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}



void CustomLabelce8f65_TimerEventDefault(void* ptr, void* event)
{
	static_cast<CustomLabelce8f65*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
