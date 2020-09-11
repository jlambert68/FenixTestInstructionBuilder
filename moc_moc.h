/****************************************************************************
** Meta object code from reading C++ file 'moc.cpp'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.13.2)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include <memory>
#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#include <QtCore/QList>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'moc.cpp' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 67
#error "This file was generated using the moc from 5.13.2. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
QT_WARNING_PUSH
QT_WARNING_DISABLE_DEPRECATED
struct qt_meta_stringdata_BridgeTemplatece8f65_t {
    QByteArrayData data[12];
    char stringdata0[107];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_BridgeTemplatece8f65_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_BridgeTemplatece8f65_t qt_meta_stringdata_BridgeTemplatece8f65 = {
    {
QT_MOC_LITERAL(0, 0, 20), // "BridgeTemplatece8f65"
QT_MOC_LITERAL(1, 21, 7), // "clicked"
QT_MOC_LITERAL(2, 29, 0), // ""
QT_MOC_LITERAL(3, 30, 18), // "sendVariantListMap"
QT_MOC_LITERAL(4, 49, 2), // "v0"
QT_MOC_LITERAL(5, 52, 2), // "v1"
QT_MOC_LITERAL(6, 55, 10), // "type424d06"
QT_MOC_LITERAL(7, 66, 2), // "v2"
QT_MOC_LITERAL(8, 69, 2), // "v3"
QT_MOC_LITERAL(9, 72, 19), // "sendVariantListMap2"
QT_MOC_LITERAL(10, 92, 9), // "QObject**"
QT_MOC_LITERAL(11, 102, 4) // "core"

    },
    "BridgeTemplatece8f65\0clicked\0\0"
    "sendVariantListMap\0v0\0v1\0type424d06\0"
    "v2\0v3\0sendVariantListMap2\0QObject**\0"
    "core"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_BridgeTemplatece8f65[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       3,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       3,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    0,   29,    2, 0x06 /* Public */,
       3,    4,   30,    2, 0x06 /* Public */,
       9,    1,   39,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void,
    QMetaType::Void, QMetaType::QVariant, QMetaType::QVariantList, 0x80000000 | 6, QMetaType::QVariant,    4,    5,    7,    8,
    QMetaType::Void, 0x80000000 | 10,   11,

       0        // eod
};

void BridgeTemplatece8f65::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<BridgeTemplatece8f65 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->clicked(); break;
        case 1: _t->sendVariantListMap((*reinterpret_cast< QVariant(*)>(_a[1])),(*reinterpret_cast< QList<QVariant>(*)>(_a[2])),(*reinterpret_cast< type424d06(*)>(_a[3])),(*reinterpret_cast< QVariant(*)>(_a[4]))); break;
        case 2: _t->sendVariantListMap2((*reinterpret_cast< QObject**(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (BridgeTemplatece8f65::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&BridgeTemplatece8f65::clicked)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (BridgeTemplatece8f65::*)(QVariant , QList<QVariant> , type424d06 , QVariant );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&BridgeTemplatece8f65::sendVariantListMap)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (BridgeTemplatece8f65::*)(QObject * * );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&BridgeTemplatece8f65::sendVariantListMap2)) {
                *result = 2;
                return;
            }
        }
    }
}

QT_INIT_METAOBJECT const QMetaObject BridgeTemplatece8f65::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_BridgeTemplatece8f65.data,
    qt_meta_data_BridgeTemplatece8f65,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *BridgeTemplatece8f65::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *BridgeTemplatece8f65::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_BridgeTemplatece8f65.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int BridgeTemplatece8f65::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 3)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 3;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 3)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 3;
    }
    return _id;
}

// SIGNAL 0
void BridgeTemplatece8f65::clicked()
{
    QMetaObject::activate(this, &staticMetaObject, 0, nullptr);
}

// SIGNAL 1
void BridgeTemplatece8f65::sendVariantListMap(QVariant _t1, QList<QVariant> _t2, type424d06 _t3, QVariant _t4)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))), const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t2))), const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t3))), const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t4))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void BridgeTemplatece8f65::sendVariantListMap2(QObject * * _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}
struct qt_meta_stringdata_CustomLabelce8f65_t {
    QByteArrayData data[4];
    char stringdata0[36];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_CustomLabelce8f65_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_CustomLabelce8f65_t qt_meta_stringdata_CustomLabelce8f65 = {
    {
QT_MOC_LITERAL(0, 0, 17), // "CustomLabelce8f65"
QT_MOC_LITERAL(1, 18, 11), // "textChanged"
QT_MOC_LITERAL(2, 30, 0), // ""
QT_MOC_LITERAL(3, 31, 4) // "text"

    },
    "CustomLabelce8f65\0textChanged\0\0text"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_CustomLabelce8f65[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       1,   14, // methods
       1,   22, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       1,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   19,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::QString,    3,

 // properties: name, type, flags
       3, QMetaType::QString, 0x00495103,

 // properties: notify_signal_id
       0,

       0        // eod
};

void CustomLabelce8f65::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<CustomLabelce8f65 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->textChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (CustomLabelce8f65::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CustomLabelce8f65::textChanged)) {
                *result = 0;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<CustomLabelce8f65 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< QString*>(_v) = _t->text(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<CustomLabelce8f65 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setText(*reinterpret_cast< QString*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject CustomLabelce8f65::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_CustomLabelce8f65.data,
    qt_meta_data_CustomLabelce8f65,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *CustomLabelce8f65::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *CustomLabelce8f65::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_CustomLabelce8f65.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int CustomLabelce8f65::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 1)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 1;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 1)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 1;
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 1;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 1;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 1;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 1;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 1;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 1;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void CustomLabelce8f65::textChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
