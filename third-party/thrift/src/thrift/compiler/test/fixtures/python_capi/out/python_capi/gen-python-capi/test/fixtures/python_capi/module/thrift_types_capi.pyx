
#
# Autogenerated by Thrift
#
# DO NOT EDIT
#  @generated
#

from cpython.ref cimport PyObject
from libcpp.utility cimport move as __move
from folly.iobuf cimport (
    from_unique_ptr as __IOBuf_from_unique_ptr,
    IOBuf as __IOBuf,
)
from thrift.python.serializer import (
    deserialize as __deserialize,
    Protocol as __Protocol,
    serialize_iobuf as __serialize_iobuf,
)
import test.fixtures.python_capi.module.thrift_types as __thrift_types

cdef api int can_extract__test__fixtures__python_capi__module__MyStruct(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyStruct) else 0


cdef api object init__test__fixtures__python_capi__module__MyStruct(object data):
    return __thrift_types.MyStruct._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__MyDataItem(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyDataItem) else 0


cdef api object init__test__fixtures__python_capi__module__MyDataItem(object data):
    return __thrift_types.MyDataItem._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__TransitiveDoubler(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.TransitiveDoubler) else 0


cdef api object init__test__fixtures__python_capi__module__TransitiveDoubler(object data):
    return __thrift_types.TransitiveDoubler._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__DoubledPair(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.DoubledPair) else 0


cdef api object init__test__fixtures__python_capi__module__DoubledPair(object data):
    return __thrift_types.DoubledPair._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__StringPair(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.StringPair) else 0


cdef api object init__test__fixtures__python_capi__module__StringPair(object data):
    return __thrift_types.StringPair._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__EmptyStruct(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.EmptyStruct) else 0


cdef api object init__test__fixtures__python_capi__module__EmptyStruct(object data):
    return __thrift_types.EmptyStruct._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__PrimitiveStruct(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.PrimitiveStruct) else 0


cdef api object init__test__fixtures__python_capi__module__PrimitiveStruct(object data):
    return __thrift_types.PrimitiveStruct._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__AdaptedFields(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.AdaptedFields) else 0


cdef api object init__test__fixtures__python_capi__module__AdaptedFields(object data):
    return __thrift_types.AdaptedFields._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__ListStruct(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.ListStruct) else 0


cdef api object init__test__fixtures__python_capi__module__ListStruct(object data):
    return __thrift_types.ListStruct._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__SetStruct(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.SetStruct) else 0


cdef api object init__test__fixtures__python_capi__module__SetStruct(object data):
    return __thrift_types.SetStruct._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__MapStruct(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MapStruct) else 0


cdef api object init__test__fixtures__python_capi__module__MapStruct(object data):
    return __thrift_types.MapStruct._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__ComposeStruct(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.ComposeStruct) else 0


cdef api object init__test__fixtures__python_capi__module__ComposeStruct(object data):
    return __thrift_types.ComposeStruct._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__Onion(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.Onion) else 0


cdef api object init__test__fixtures__python_capi__module__Onion(object data):
    return __thrift_types.Onion._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__MyStructPatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyStructPatch) else 0

cdef api __cIOBuf* extract__test__fixtures__python_capi__module__MyStructPatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__python_capi__module__MyStructPatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyStructPatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__python_capi__module__MyStructPatch(object data):
    return __thrift_types.MyStructPatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__MyStructField4Patch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyStructField4Patch) else 0

cdef api __cIOBuf* extract__test__fixtures__python_capi__module__MyStructField4Patch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__python_capi__module__MyStructField4Patch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyStructField4Patch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__python_capi__module__MyStructField4Patch(object data):
    return __thrift_types.MyStructField4Patch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__MyStructField6Patch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyStructField6Patch) else 0

cdef api __cIOBuf* extract__test__fixtures__python_capi__module__MyStructField6Patch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__python_capi__module__MyStructField6Patch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyStructField6Patch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__python_capi__module__MyStructField6Patch(object data):
    return __thrift_types.MyStructField6Patch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__MyStructField7Patch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyStructField7Patch) else 0

cdef api __cIOBuf* extract__test__fixtures__python_capi__module__MyStructField7Patch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__python_capi__module__MyStructField7Patch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyStructField7Patch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__python_capi__module__MyStructField7Patch(object data):
    return __thrift_types.MyStructField7Patch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__MyStructField8Patch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyStructField8Patch) else 0

cdef api __cIOBuf* extract__test__fixtures__python_capi__module__MyStructField8Patch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__python_capi__module__MyStructField8Patch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyStructField8Patch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__python_capi__module__MyStructField8Patch(object data):
    return __thrift_types.MyStructField8Patch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__MyStructFieldPatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyStructFieldPatch) else 0

cdef api __cIOBuf* extract__test__fixtures__python_capi__module__MyStructFieldPatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__python_capi__module__MyStructFieldPatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyStructFieldPatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__python_capi__module__MyStructFieldPatch(object data):
    return __thrift_types.MyStructFieldPatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__MyStructEnsureStruct(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyStructEnsureStruct) else 0

cdef api __cIOBuf* extract__test__fixtures__python_capi__module__MyStructEnsureStruct(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__python_capi__module__MyStructEnsureStruct(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyStructEnsureStruct,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__python_capi__module__MyStructEnsureStruct(object data):
    return __thrift_types.MyStructEnsureStruct._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__MyStructSafePatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyStructSafePatch) else 0

cdef api __cIOBuf* extract__test__fixtures__python_capi__module__MyStructSafePatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__python_capi__module__MyStructSafePatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyStructSafePatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__python_capi__module__MyStructSafePatch(object data):
    return __thrift_types.MyStructSafePatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__MyDataItemPatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyDataItemPatch) else 0

cdef api __cIOBuf* extract__test__fixtures__python_capi__module__MyDataItemPatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__python_capi__module__MyDataItemPatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyDataItemPatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__python_capi__module__MyDataItemPatch(object data):
    return __thrift_types.MyDataItemPatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__MyDataItemFieldPatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyDataItemFieldPatch) else 0

cdef api __cIOBuf* extract__test__fixtures__python_capi__module__MyDataItemFieldPatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__python_capi__module__MyDataItemFieldPatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyDataItemFieldPatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__python_capi__module__MyDataItemFieldPatch(object data):
    return __thrift_types.MyDataItemFieldPatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__MyDataItemEnsureStruct(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyDataItemEnsureStruct) else 0

cdef api __cIOBuf* extract__test__fixtures__python_capi__module__MyDataItemEnsureStruct(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__python_capi__module__MyDataItemEnsureStruct(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyDataItemEnsureStruct,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__python_capi__module__MyDataItemEnsureStruct(object data):
    return __thrift_types.MyDataItemEnsureStruct._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__MyDataItemSafePatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyDataItemSafePatch) else 0

cdef api __cIOBuf* extract__test__fixtures__python_capi__module__MyDataItemSafePatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__python_capi__module__MyDataItemSafePatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyDataItemSafePatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__python_capi__module__MyDataItemSafePatch(object data):
    return __thrift_types.MyDataItemSafePatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__python_capi__module__MyEnum(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyEnum) else 0

cdef api object construct__test__fixtures__python_capi__module__MyEnum(int64_t __val):
    return __thrift_types.MyEnum(__val)

cdef api int can_extract__test__fixtures__python_capi__module__AnnoyingEnum(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.AnnoyingEnum) else 0

cdef api object construct__test__fixtures__python_capi__module__AnnoyingEnum(int64_t __val):
    return __thrift_types.AnnoyingEnum(__val)

