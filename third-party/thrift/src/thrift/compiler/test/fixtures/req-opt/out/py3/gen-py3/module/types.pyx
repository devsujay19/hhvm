#
# Autogenerated by Thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#
cimport cython as __cython
from cpython.object cimport PyTypeObject, Py_LT, Py_LE, Py_EQ, Py_NE, Py_GT, Py_GE
from libcpp.memory cimport shared_ptr, make_shared, unique_ptr
from libcpp.optional cimport optional as __optional
from libcpp.string cimport string
from libcpp cimport bool as cbool
from libcpp.iterator cimport inserter as cinserter
from libcpp.utility cimport move as cmove
from cpython cimport bool as pbool
from cython.operator cimport dereference as deref, preincrement as inc, address as ptr_address
import thrift.py3.types
from thrift.py3.types import _IsSet as _fbthrift_IsSet
from thrift.py3.types cimport make_unique
cimport thrift.py3.types
cimport thrift.py3.exceptions
cimport thrift.python.exceptions
from thrift.python.std_libcpp cimport sv_to_str as __sv_to_str, string_view as __cstring_view
from thrift.py3.types cimport (
    cSetOp as __cSetOp,
    richcmp as __richcmp,
    set_op as __set_op,
    setcmp as __setcmp,
    list_index as __list_index,
    list_count as __list_count,
    list_slice as __list_slice,
    list_getitem as __list_getitem,
    set_iter as __set_iter,
    map_iter as __map_iter,
    map_contains as __map_contains,
    map_getitem as __map_getitem,
    reference_shared_ptr as __reference_shared_ptr,
    get_field_name_by_index as __get_field_name_by_index,
    reset_field as __reset_field,
    translate_cpp_enum_to_python,
    SetMetaClass as __SetMetaClass,
    const_pointer_cast,
    constant_shared_ptr,
    NOTSET as __NOTSET,
    EnumData as __EnumData,
    EnumFlagsData as __EnumFlagsData,
    UnionTypeEnumData as __UnionTypeEnumData,
    createEnumDataForUnionType as __createEnumDataForUnionType,
)
cimport thrift.py3.serializer as serializer
from thrift.python.protocol cimport Protocol as __Protocol
import folly.iobuf as _fbthrift_iobuf
from folly.optional cimport cOptional
from folly.memory cimport to_shared_ptr as __to_shared_ptr
from folly.range cimport Range as __cRange

import sys
from collections.abc import Sequence, Set, Mapping, Iterable
import weakref as __weakref
import builtins as _builtins



cdef object get_types_reflection():
    import importlib
    return importlib.import_module(
        "module.types_reflection"
    )

@__cython.auto_pickle(False)
cdef class Foo(thrift.py3.types.Struct):
    def __init__(Foo self, **kwargs):
        self._cpp_obj = make_shared[cFoo]()
        self._fields_setter = _fbthrift_types_fields.__Foo_FieldsSetter._fbthrift_create(self._cpp_obj.get())
        super().__init__(**kwargs)

    def __call__(Foo self, **kwargs):
        if not kwargs:
            return self
        cdef Foo __fbthrift_inst = Foo.__new__(Foo)
        __fbthrift_inst._cpp_obj = make_shared[cFoo](deref(self._cpp_obj))
        __fbthrift_inst._fields_setter = _fbthrift_types_fields.__Foo_FieldsSetter._fbthrift_create(__fbthrift_inst._cpp_obj.get())
        for __fbthrift_name, _fbthrift_value in kwargs.items():
            __fbthrift_inst._fbthrift_set_field(__fbthrift_name, _fbthrift_value)
        return __fbthrift_inst

    cdef void _fbthrift_set_field(self, str name, object value) except *:
        self._fields_setter.set_field(name.encode("utf-8"), value)

    cdef object _fbthrift_isset(self):
        return _fbthrift_IsSet("Foo", {
          "myInteger": deref(self._cpp_obj).myInteger_ref().has_value(),
          "myString": deref(self._cpp_obj).myString_ref().has_value(),
          "myBools": deref(self._cpp_obj).myBools_ref().has_value(),
          "myNumbers": deref(self._cpp_obj).myNumbers_ref().has_value(),
        })

    @staticmethod
    cdef _fbthrift_create(shared_ptr[cFoo] cpp_obj):
        __fbthrift_inst = <Foo>Foo.__new__(Foo)
        __fbthrift_inst._cpp_obj = cmove(cpp_obj)
        return __fbthrift_inst

    cdef inline myInteger_impl(self):

        return deref(self._cpp_obj).myInteger_ref().value()

    @property
    def myInteger(self):
        return self.myInteger_impl()

    cdef inline myString_impl(self):
        if not deref(self._cpp_obj).myString_ref().has_value():
            return None

        return (<bytes>deref(self._cpp_obj).myString_ref().value_unchecked()).decode('UTF-8')

    @property
    def myString(self):
        return self.myString_impl()

    cdef inline myBools_impl(self):

        if self.__fbthrift_cached_myBools is None:
            self.__fbthrift_cached_myBools = List__bool._fbthrift_create(__reference_shared_ptr(deref(self._cpp_obj).myBools_ref().ref(), self._cpp_obj))
        return self.__fbthrift_cached_myBools

    @property
    def myBools(self):
        return self.myBools_impl()

    cdef inline myNumbers_impl(self):

        if self.__fbthrift_cached_myNumbers is None:
            self.__fbthrift_cached_myNumbers = List__i32._fbthrift_create(__reference_shared_ptr(deref(self._cpp_obj).myNumbers_ref().ref(), self._cpp_obj))
        return self.__fbthrift_cached_myNumbers

    @property
    def myNumbers(self):
        return self.myNumbers_impl()


    def __hash__(Foo self):
        return super().__hash__()

    def __repr__(Foo self):
        return super().__repr__()

    def __str__(Foo self):
        return super().__str__()


    def __copy__(Foo self):
        cdef shared_ptr[cFoo] cpp_obj = make_shared[cFoo](
            deref(self._cpp_obj)
        )
        return Foo._fbthrift_create(cmove(cpp_obj))

    def __richcmp__(self, other, int op):
        r = self._fbthrift_cmp_sametype(other, op)
        return __richcmp[cFoo](
            self._cpp_obj,
            (<Foo>other)._cpp_obj,
            op,
        ) if r is None else r

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__Foo()

    @staticmethod
    def __get_metadata__():
        cdef __fbthrift_cThriftMetadata meta
        StructMetadata[cFoo].gen(meta)
        return __MetadataBox.box(cmove(meta))

    @staticmethod
    def __get_thrift_name__():
        return "module.Foo"

    @classmethod
    def _fbthrift_get_field_name_by_index(cls, idx):
        return __sv_to_str(__get_field_name_by_index[cFoo](idx))

    @classmethod
    def _fbthrift_get_struct_size(cls):
        return 4

    cdef _fbthrift_iobuf.IOBuf _fbthrift_serialize(Foo self, __Protocol proto):
        cdef unique_ptr[_fbthrift_iobuf.cIOBuf] data
        with nogil:
            data = cmove(serializer.cserialize[cFoo](self._cpp_obj.get(), proto))
        return _fbthrift_iobuf.from_unique_ptr(cmove(data))

    cdef cuint32_t _fbthrift_deserialize(Foo self, const _fbthrift_iobuf.cIOBuf* buf, __Protocol proto) except? 0:
        cdef cuint32_t needed
        self._cpp_obj = make_shared[cFoo]()
        with nogil:
            needed = serializer.cdeserialize[cFoo](buf, self._cpp_obj.get(), proto)
        return needed


    def _to_python(self):
        import importlib
        import thrift.python.converter
        python_types = importlib.import_module(
            "module.thrift_types"
        )
        return thrift.python.converter.to_python_struct(python_types.Foo, self)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        import importlib
        import thrift.util.converter
        py_deprecated_types = importlib.import_module("module.ttypes")
        return thrift.util.converter.to_py_struct(py_deprecated_types.Foo, self)
@__cython.auto_pickle(False)
cdef class List__bool(thrift.py3.types.List):
    def __init__(self, items=None):
        if isinstance(items, List__bool):
            self._cpp_obj = (<List__bool> items)._cpp_obj
        else:
            self._cpp_obj = List__bool._make_instance(items)

    @staticmethod
    cdef _fbthrift_create(shared_ptr[vector[cbool]] c_items):
        __fbthrift_inst = <List__bool>List__bool.__new__(List__bool)
        __fbthrift_inst._cpp_obj = cmove(c_items)
        return __fbthrift_inst

    def __copy__(List__bool self):
        cdef shared_ptr[vector[cbool]] cpp_obj = make_shared[vector[cbool]](
            deref(self._cpp_obj)
        )
        return List__bool._fbthrift_create(cmove(cpp_obj))

    def __len__(self):
        return deref(self._cpp_obj).size()

    @staticmethod
    cdef shared_ptr[vector[cbool]] _make_instance(object items) except *:
        cdef shared_ptr[vector[cbool]] c_inst = make_shared[vector[cbool]]()
        if items is not None:
            for item in items:
                if not isinstance(item, bool):
                    raise TypeError(f"{item!r} is not of type bool")
                deref(c_inst).push_back(item)
        return c_inst

    cdef _get_slice(self, slice index_obj):
        cdef int start, stop, step
        start, stop, step = index_obj.indices(deref(self._cpp_obj).size())
        return List__bool._fbthrift_create(
            __list_slice[vector[cbool]](self._cpp_obj, start, stop, step)
        )

    cdef _get_single_item(self, size_t index):
        cdef cbool citem = False
        __list_getitem(self._cpp_obj, index, citem)
        return citem

    cdef _check_item_type(self, item):
        if not self or item is None:
            return
        if isinstance(item, bool):
            return item

    def index(self, item, start=0, stop=None):
        err = ValueError(f'{item} is not in list')
        item = self._check_item_type(item)
        if item is None:
            raise err
        cdef (int, int, int) indices = slice(start, stop).indices(deref(self._cpp_obj).size())
        cdef cbool citem = item
        cdef __optional[size_t] found = __list_index[vector[cbool]](self._cpp_obj, indices[0], indices[1], citem)
        if not found.has_value():
            raise err
        return found.value()

    def count(self, item):
        item = self._check_item_type(item)
        if item is None:
            return 0
        cdef cbool citem = item
        return __list_count[vector[cbool]](self._cpp_obj, citem)

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__List__bool()


Sequence.register(List__bool)

@__cython.auto_pickle(False)
cdef class List__i32(thrift.py3.types.List):
    def __init__(self, items=None):
        if isinstance(items, List__i32):
            self._cpp_obj = (<List__i32> items)._cpp_obj
        else:
            self._cpp_obj = List__i32._make_instance(items)

    @staticmethod
    cdef _fbthrift_create(shared_ptr[vector[cint32_t]] c_items):
        __fbthrift_inst = <List__i32>List__i32.__new__(List__i32)
        __fbthrift_inst._cpp_obj = cmove(c_items)
        return __fbthrift_inst

    def __copy__(List__i32 self):
        cdef shared_ptr[vector[cint32_t]] cpp_obj = make_shared[vector[cint32_t]](
            deref(self._cpp_obj)
        )
        return List__i32._fbthrift_create(cmove(cpp_obj))

    def __len__(self):
        return deref(self._cpp_obj).size()

    @staticmethod
    cdef shared_ptr[vector[cint32_t]] _make_instance(object items) except *:
        cdef shared_ptr[vector[cint32_t]] c_inst = make_shared[vector[cint32_t]]()
        if items is not None:
            for item in items:
                if not isinstance(item, int):
                    raise TypeError(f"{item!r} is not of type int")
                item = <cint32_t> item
                deref(c_inst).push_back(item)
        return c_inst

    cdef _get_slice(self, slice index_obj):
        cdef int start, stop, step
        start, stop, step = index_obj.indices(deref(self._cpp_obj).size())
        return List__i32._fbthrift_create(
            __list_slice[vector[cint32_t]](self._cpp_obj, start, stop, step)
        )

    cdef _get_single_item(self, size_t index):
        cdef cint32_t citem = 0
        __list_getitem(self._cpp_obj, index, citem)
        return citem

    cdef _check_item_type(self, item):
        if not self or item is None:
            return
        if isinstance(item, int):
            return item

    def index(self, item, start=0, stop=None):
        err = ValueError(f'{item} is not in list')
        item = self._check_item_type(item)
        if item is None:
            raise err
        cdef (int, int, int) indices = slice(start, stop).indices(deref(self._cpp_obj).size())
        cdef cint32_t citem = item
        cdef __optional[size_t] found = __list_index[vector[cint32_t]](self._cpp_obj, indices[0], indices[1], citem)
        if not found.has_value():
            raise err
        return found.value()

    def count(self, item):
        item = self._check_item_type(item)
        if item is None:
            return 0
        cdef cint32_t citem = item
        return __list_count[vector[cint32_t]](self._cpp_obj, citem)

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__List__i32()


Sequence.register(List__i32)

