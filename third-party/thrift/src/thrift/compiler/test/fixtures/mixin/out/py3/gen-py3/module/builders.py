#
# Autogenerated by Thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#
import typing as _typing

import folly.iobuf as _fbthrift_iobuf
import thrift.py3.builder


import module.types as _module_types


_fbthrift_struct_type__Mixin1 = _module_types.Mixin1
class Mixin1_Builder(thrift.py3.builder.StructBuilder):
    _struct_type = _fbthrift_struct_type__Mixin1

    def __init__(self):
        self.field1: _typing.Optional[str] = None

    def __iter__(self):
        yield "field1", self.field1

_fbthrift_struct_type__Mixin2 = _module_types.Mixin2
class Mixin2_Builder(thrift.py3.builder.StructBuilder):
    _struct_type = _fbthrift_struct_type__Mixin2

    def __init__(self):
        self.m1: _typing.Any = None
        self.field2: _typing.Optional[str] = None

    def __iter__(self):
        yield "m1", self.m1
        yield "field2", self.field2

_fbthrift_struct_type__Mixin3Base = _module_types.Mixin3Base
class Mixin3Base_Builder(thrift.py3.builder.StructBuilder):
    _struct_type = _fbthrift_struct_type__Mixin3Base

    def __init__(self):
        self.field3: _typing.Optional[str] = None

    def __iter__(self):
        yield "field3", self.field3

_fbthrift_struct_type__Foo = _module_types.Foo
class Foo_Builder(thrift.py3.builder.StructBuilder):
    _struct_type = _fbthrift_struct_type__Foo

    def __init__(self):
        self.field4: _typing.Optional[str] = None
        self.m2: _typing.Any = None
        self.m3: _typing.Any = None

    def __iter__(self):
        yield "field4", self.field4
        yield "m2", self.m2
        yield "m3", self.m3

