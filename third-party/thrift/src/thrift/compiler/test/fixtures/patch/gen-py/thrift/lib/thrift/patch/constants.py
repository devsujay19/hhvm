#
# Autogenerated by Thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

from __future__ import absolute_import
import sys
from thrift.util.Recursive import fix_spec
from thrift.Thrift import TType, TMessageType, TPriority, TRequestContext, TProcessorEventHandler, TServerInterface, TProcessor, TException, TApplicationException, UnimplementedTypedef
from thrift.protocol.TProtocol import TProtocolException

from json import loads
import sys
if sys.version_info[0] >= 3:
  long = int

import thrift.lib.thrift.standard.ttypes
import thrift.lib.thrift.id.ttypes


from .ttypes import UTF8STRINGS, PatchOp, GeneratePatch, AssignOnlyPatch, BoolPatch, BytePatch, I16Patch, I32Patch, I64Patch, FloatPatch, DoublePatch, StringPatch, BinaryPatch, FieldId, FieldIdList

