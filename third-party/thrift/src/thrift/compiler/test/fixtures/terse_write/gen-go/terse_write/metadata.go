// @generated by Thrift for [[[ program path ]]]
// This file is probably not the place you want to edit!

package terse_write // [[[ program thrift source path ]]]

import (
    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift"
    metadata "github.com/facebook/fbthrift/thrift/lib/thrift/metadata"
)

// mapsCopy is a copy of maps.Copy from Go 1.21
// TODO: remove mapsCopy once we can safely upgrade to Go 1.21 without requiring any rollback.
func mapsCopy[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2) {
	for k, v := range src {
		dst[k] = v
	}
}

// (needed to ensure safety because of naive import list construction)
var _ = thrift.ZERO
// TODO: uncomment when can safely upgrade to Go 1.21 without requiring any rollback.
// var _ = maps.Copy[map[int]int, map[int]int]
var _ = metadata.GoUnusedProtection__

// Premade Thrift types
var (
    premadeThriftType_bool = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_BOOL_TYPE.Ptr(),
            )
    premadeThriftType_byte = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_BYTE_TYPE.Ptr(),
            )
    premadeThriftType_i16 = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_I16_TYPE.Ptr(),
            )
    premadeThriftType_i32 = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_I32_TYPE.Ptr(),
            )
    premadeThriftType_i64 = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_I64_TYPE.Ptr(),
            )
    premadeThriftType_float = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_FLOAT_TYPE.Ptr(),
            )
    premadeThriftType_double = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_DOUBLE_TYPE.Ptr(),
            )
    premadeThriftType_string = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_STRING_TYPE.Ptr(),
            )
    premadeThriftType_binary = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_BINARY_TYPE.Ptr(),
            )
    premadeThriftType_terse_write_MyEnum = metadata.NewThriftType().SetTEnum(
        metadata.NewThriftEnumType().
            SetName("terse_write.MyEnum"),
            )
    premadeThriftType_list_i16 = metadata.NewThriftType().SetTList(
        metadata.NewThriftListType().
            SetValueType(premadeThriftType_i16),
            )
    premadeThriftType_set_i16 = metadata.NewThriftType().SetTSet(
        metadata.NewThriftSetType().
            SetValueType(premadeThriftType_i16),
            )
    premadeThriftType_map_i16_i16 = metadata.NewThriftType().SetTMap(
        metadata.NewThriftMapType().
            SetKeyType(premadeThriftType_i16).
            SetValueType(premadeThriftType_i16),
            )
    premadeThriftType_terse_write_MyStruct = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("terse_write.MyStruct"),
            )
    premadeThriftType_terse_write_MyUnion = metadata.NewThriftType().SetTUnion(
        metadata.NewThriftUnionType().
            SetName("terse_write.MyUnion"),
            )
    premadeThriftType_terse_write_MyStructWithCustomDefault = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("terse_write.MyStructWithCustomDefault"),
            )
    premadeThriftType_terse_write_MyInteger = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("terse_write.MyInteger").
            SetUnderlyingType(premadeThriftType_i32),
            )
)

var structMetadatas = []*metadata.ThriftStruct{
    metadata.NewThriftStruct().
    SetName("terse_write.MyStruct").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("terse_write.MyUnion").
    SetIsUnion(true).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("bool_field").
    SetIsOptional(false).
    SetType(premadeThriftType_bool),
            metadata.NewThriftField().
    SetId(2).
    SetName("byte_field").
    SetIsOptional(false).
    SetType(premadeThriftType_byte),
            metadata.NewThriftField().
    SetId(3).
    SetName("short_field").
    SetIsOptional(false).
    SetType(premadeThriftType_i16),
            metadata.NewThriftField().
    SetId(4).
    SetName("int_field").
    SetIsOptional(false).
    SetType(premadeThriftType_i32),
            metadata.NewThriftField().
    SetId(5).
    SetName("long_field").
    SetIsOptional(false).
    SetType(premadeThriftType_i64),
            metadata.NewThriftField().
    SetId(6).
    SetName("float_field").
    SetIsOptional(false).
    SetType(premadeThriftType_float),
            metadata.NewThriftField().
    SetId(7).
    SetName("double_field").
    SetIsOptional(false).
    SetType(premadeThriftType_double),
            metadata.NewThriftField().
    SetId(8).
    SetName("string_field").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
            metadata.NewThriftField().
    SetId(9).
    SetName("binary_field").
    SetIsOptional(false).
    SetType(premadeThriftType_binary),
            metadata.NewThriftField().
    SetId(10).
    SetName("enum_field").
    SetIsOptional(false).
    SetType(premadeThriftType_terse_write_MyEnum),
            metadata.NewThriftField().
    SetId(11).
    SetName("list_field").
    SetIsOptional(false).
    SetType(premadeThriftType_list_i16),
            metadata.NewThriftField().
    SetId(12).
    SetName("set_field").
    SetIsOptional(false).
    SetType(premadeThriftType_set_i16),
            metadata.NewThriftField().
    SetId(13).
    SetName("map_field").
    SetIsOptional(false).
    SetType(premadeThriftType_map_i16_i16),
            metadata.NewThriftField().
    SetId(14).
    SetName("struct_field").
    SetIsOptional(false).
    SetType(premadeThriftType_terse_write_MyStruct),
        },
    ),
    metadata.NewThriftStruct().
    SetName("terse_write.MyStructWithCustomDefault").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("field1").
    SetIsOptional(false).
    SetType(premadeThriftType_i64),
        },
    ),
    metadata.NewThriftStruct().
    SetName("terse_write.StructLevelTerseStruct").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("bool_field").
    SetIsOptional(false).
    SetType(premadeThriftType_bool),
            metadata.NewThriftField().
    SetId(2).
    SetName("byte_field").
    SetIsOptional(false).
    SetType(premadeThriftType_byte),
            metadata.NewThriftField().
    SetId(3).
    SetName("short_field").
    SetIsOptional(false).
    SetType(premadeThriftType_i16),
            metadata.NewThriftField().
    SetId(4).
    SetName("int_field").
    SetIsOptional(false).
    SetType(premadeThriftType_i32),
            metadata.NewThriftField().
    SetId(5).
    SetName("long_field").
    SetIsOptional(false).
    SetType(premadeThriftType_i64),
            metadata.NewThriftField().
    SetId(6).
    SetName("float_field").
    SetIsOptional(false).
    SetType(premadeThriftType_float),
            metadata.NewThriftField().
    SetId(7).
    SetName("double_field").
    SetIsOptional(false).
    SetType(premadeThriftType_double),
            metadata.NewThriftField().
    SetId(8).
    SetName("string_field").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
            metadata.NewThriftField().
    SetId(9).
    SetName("binary_field").
    SetIsOptional(false).
    SetType(premadeThriftType_binary),
            metadata.NewThriftField().
    SetId(10).
    SetName("enum_field").
    SetIsOptional(false).
    SetType(premadeThriftType_terse_write_MyEnum),
            metadata.NewThriftField().
    SetId(11).
    SetName("list_field").
    SetIsOptional(false).
    SetType(premadeThriftType_list_i16),
            metadata.NewThriftField().
    SetId(12).
    SetName("set_field").
    SetIsOptional(false).
    SetType(premadeThriftType_set_i16),
            metadata.NewThriftField().
    SetId(13).
    SetName("map_field").
    SetIsOptional(false).
    SetType(premadeThriftType_map_i16_i16),
            metadata.NewThriftField().
    SetId(14).
    SetName("struct_field").
    SetIsOptional(false).
    SetType(premadeThriftType_terse_write_MyStruct),
            metadata.NewThriftField().
    SetId(15).
    SetName("union_field").
    SetIsOptional(false).
    SetType(premadeThriftType_terse_write_MyUnion),
        },
    ),
    metadata.NewThriftStruct().
    SetName("terse_write.FieldLevelTerseStruct").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("terse_bool_field").
    SetIsOptional(false).
    SetType(premadeThriftType_bool),
            metadata.NewThriftField().
    SetId(2).
    SetName("terse_byte_field").
    SetIsOptional(false).
    SetType(premadeThriftType_byte),
            metadata.NewThriftField().
    SetId(3).
    SetName("terse_short_field").
    SetIsOptional(false).
    SetType(premadeThriftType_i16),
            metadata.NewThriftField().
    SetId(4).
    SetName("terse_int_field").
    SetIsOptional(false).
    SetType(premadeThriftType_i32),
            metadata.NewThriftField().
    SetId(5).
    SetName("terse_long_field").
    SetIsOptional(false).
    SetType(premadeThriftType_i64),
            metadata.NewThriftField().
    SetId(6).
    SetName("terse_float_field").
    SetIsOptional(false).
    SetType(premadeThriftType_float),
            metadata.NewThriftField().
    SetId(7).
    SetName("terse_double_field").
    SetIsOptional(false).
    SetType(premadeThriftType_double),
            metadata.NewThriftField().
    SetId(8).
    SetName("terse_string_field").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
            metadata.NewThriftField().
    SetId(9).
    SetName("terse_binary_field").
    SetIsOptional(false).
    SetType(premadeThriftType_binary),
            metadata.NewThriftField().
    SetId(10).
    SetName("terse_enum_field").
    SetIsOptional(false).
    SetType(premadeThriftType_terse_write_MyEnum),
            metadata.NewThriftField().
    SetId(11).
    SetName("terse_list_field").
    SetIsOptional(false).
    SetType(premadeThriftType_list_i16),
            metadata.NewThriftField().
    SetId(12).
    SetName("terse_set_field").
    SetIsOptional(false).
    SetType(premadeThriftType_set_i16),
            metadata.NewThriftField().
    SetId(13).
    SetName("terse_map_field").
    SetIsOptional(false).
    SetType(premadeThriftType_map_i16_i16),
            metadata.NewThriftField().
    SetId(14).
    SetName("terse_struct_field").
    SetIsOptional(false).
    SetType(premadeThriftType_terse_write_MyStruct),
            metadata.NewThriftField().
    SetId(15).
    SetName("bool_field").
    SetIsOptional(false).
    SetType(premadeThriftType_bool),
            metadata.NewThriftField().
    SetId(16).
    SetName("byte_field").
    SetIsOptional(false).
    SetType(premadeThriftType_byte),
            metadata.NewThriftField().
    SetId(17).
    SetName("short_field").
    SetIsOptional(false).
    SetType(premadeThriftType_i16),
            metadata.NewThriftField().
    SetId(18).
    SetName("int_field").
    SetIsOptional(false).
    SetType(premadeThriftType_i32),
            metadata.NewThriftField().
    SetId(19).
    SetName("long_field").
    SetIsOptional(false).
    SetType(premadeThriftType_i64),
            metadata.NewThriftField().
    SetId(20).
    SetName("float_field").
    SetIsOptional(false).
    SetType(premadeThriftType_float),
            metadata.NewThriftField().
    SetId(21).
    SetName("double_field").
    SetIsOptional(false).
    SetType(premadeThriftType_double),
            metadata.NewThriftField().
    SetId(22).
    SetName("string_field").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
            metadata.NewThriftField().
    SetId(23).
    SetName("binary_field").
    SetIsOptional(false).
    SetType(premadeThriftType_binary),
            metadata.NewThriftField().
    SetId(24).
    SetName("enum_field").
    SetIsOptional(false).
    SetType(premadeThriftType_terse_write_MyEnum),
            metadata.NewThriftField().
    SetId(25).
    SetName("list_field").
    SetIsOptional(false).
    SetType(premadeThriftType_list_i16),
            metadata.NewThriftField().
    SetId(26).
    SetName("set_field").
    SetIsOptional(false).
    SetType(premadeThriftType_set_i16),
            metadata.NewThriftField().
    SetId(27).
    SetName("map_field").
    SetIsOptional(false).
    SetType(premadeThriftType_map_i16_i16),
            metadata.NewThriftField().
    SetId(28).
    SetName("struct_field").
    SetIsOptional(false).
    SetType(premadeThriftType_terse_write_MyStruct),
            metadata.NewThriftField().
    SetId(29).
    SetName("terse_union_field").
    SetIsOptional(false).
    SetType(premadeThriftType_terse_write_MyUnion),
            metadata.NewThriftField().
    SetId(30).
    SetName("union_field").
    SetIsOptional(false).
    SetType(premadeThriftType_terse_write_MyUnion),
        },
    ),
    metadata.NewThriftStruct().
    SetName("terse_write.TerseStructWithCustomDefault").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("bool_field").
    SetIsOptional(false).
    SetType(premadeThriftType_bool),
            metadata.NewThriftField().
    SetId(2).
    SetName("byte_field").
    SetIsOptional(false).
    SetType(premadeThriftType_byte),
            metadata.NewThriftField().
    SetId(3).
    SetName("short_field").
    SetIsOptional(false).
    SetType(premadeThriftType_i16),
            metadata.NewThriftField().
    SetId(4).
    SetName("int_field").
    SetIsOptional(false).
    SetType(premadeThriftType_i32),
            metadata.NewThriftField().
    SetId(5).
    SetName("long_field").
    SetIsOptional(false).
    SetType(premadeThriftType_i64),
            metadata.NewThriftField().
    SetId(6).
    SetName("float_field").
    SetIsOptional(false).
    SetType(premadeThriftType_float),
            metadata.NewThriftField().
    SetId(7).
    SetName("double_field").
    SetIsOptional(false).
    SetType(premadeThriftType_double),
            metadata.NewThriftField().
    SetId(8).
    SetName("string_field").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
            metadata.NewThriftField().
    SetId(9).
    SetName("binary_field").
    SetIsOptional(false).
    SetType(premadeThriftType_binary),
            metadata.NewThriftField().
    SetId(10).
    SetName("enum_field").
    SetIsOptional(false).
    SetType(premadeThriftType_terse_write_MyEnum),
            metadata.NewThriftField().
    SetId(11).
    SetName("list_field").
    SetIsOptional(false).
    SetType(premadeThriftType_list_i16),
            metadata.NewThriftField().
    SetId(12).
    SetName("set_field").
    SetIsOptional(false).
    SetType(premadeThriftType_set_i16),
            metadata.NewThriftField().
    SetId(13).
    SetName("map_field").
    SetIsOptional(false).
    SetType(premadeThriftType_map_i16_i16),
            metadata.NewThriftField().
    SetId(14).
    SetName("struct_field").
    SetIsOptional(false).
    SetType(premadeThriftType_terse_write_MyStructWithCustomDefault),
        },
    ),
    metadata.NewThriftStruct().
    SetName("terse_write.AdaptedFields").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("field1").
    SetIsOptional(false).
    SetType(premadeThriftType_terse_write_MyInteger),
            metadata.NewThriftField().
    SetId(2).
    SetName("field2").
    SetIsOptional(false).
    SetType(premadeThriftType_i32),
            metadata.NewThriftField().
    SetId(3).
    SetName("field3").
    SetIsOptional(false).
    SetType(premadeThriftType_terse_write_MyInteger),
        },
    ),
    metadata.NewThriftStruct().
    SetName("terse_write.WrappedFields").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("field1").
    SetIsOptional(false).
    SetType(premadeThriftType_i32),
        },
    ),
}

var exceptionMetadatas = []*metadata.ThriftException{
    metadata.NewThriftException().
    SetName("terse_write.TerseException").
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("msg").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
        },
    ),
}

var enumMetadatas = []*metadata.ThriftEnum{
    metadata.NewThriftEnum().
    SetName("terse_write.MyEnum").
    SetElements(
        map[int32]string{
            0: "ME0",
            1: "ME1",
        },
    ),
}

var serviceMetadatas = []*metadata.ThriftService{
}

// GetThriftMetadata returns complete Thrift metadata for current and imported packages.
func GetThriftMetadata() *metadata.ThriftMetadata {
    allEnums := GetEnumsMetadata()
    allStructs := GetStructsMetadata()
    allExceptions := GetExceptionsMetadata()
    allServices := GetServicesMetadata()

    return metadata.NewThriftMetadata().
        SetEnums(allEnums).
        SetStructs(allStructs).
        SetExceptions(allExceptions).
        SetServices(allServices)
}

// GetEnumsMetadata returns Thrift metadata for enums in the current and recursively included packages.
func GetEnumsMetadata() map[string]*metadata.ThriftEnum {
    allEnumsMap := make(map[string]*metadata.ThriftEnum)

    // Add enum metadatas from the current program...
    for _, enumMetadata := range enumMetadatas {
        allEnumsMap[enumMetadata.GetName()] = enumMetadata
    }

    // ...now add enum metadatas from recursively included programs.

    return allEnumsMap
}

// GetStructsMetadata returns Thrift metadata for structs in the current and recursively included packages.
func GetStructsMetadata() map[string]*metadata.ThriftStruct {
    allStructsMap := make(map[string]*metadata.ThriftStruct)

    // Add struct metadatas from the current program...
    for _, structMetadata := range structMetadatas {
        allStructsMap[structMetadata.GetName()] = structMetadata
    }

    // ...now add struct metadatas from recursively included programs.

    return allStructsMap
}

// GetExceptionsMetadata returns Thrift metadata for exceptions in the current and recursively included packages.
func GetExceptionsMetadata() map[string]*metadata.ThriftException {
    allExceptionsMap := make(map[string]*metadata.ThriftException)

    // Add exception metadatas from the current program...
    for _, exceptionMetadata := range exceptionMetadatas {
        allExceptionsMap[exceptionMetadata.GetName()] = exceptionMetadata
    }

    // ...now add exception metadatas from recursively included programs.

    return allExceptionsMap
}

// GetServicesMetadata returns Thrift metadata for services in the current and recursively included packages.
func GetServicesMetadata() map[string]*metadata.ThriftService {
    allServicesMap := make(map[string]*metadata.ThriftService)

    // Add service metadatas from the current program...
    for _, serviceMetadata := range serviceMetadatas {
        allServicesMap[serviceMetadata.GetName()] = serviceMetadata
    }

    // ...now add service metadatas from recursively included programs.

    return allServicesMap
}

// GetThriftMetadataForService returns Thrift metadata for the given service.
func GetThriftMetadataForService(scopedServiceName string) *metadata.ThriftMetadata {
    thriftMetadata := GetThriftMetadata()

    allServicesMap := thriftMetadata.GetServices()
    relevantServicesMap := make(map[string]*metadata.ThriftService)

    serviceMetadata := allServicesMap[scopedServiceName]
    // Visit and record all recursive parents of the target service.
    for serviceMetadata != nil {
        relevantServicesMap[serviceMetadata.GetName()] = serviceMetadata
        if serviceMetadata.IsSetParent() {
            serviceMetadata = allServicesMap[serviceMetadata.GetParent()]
        } else {
            serviceMetadata = nil
        }
    }

    thriftMetadata.SetServices(relevantServicesMap)

    return thriftMetadata
}
