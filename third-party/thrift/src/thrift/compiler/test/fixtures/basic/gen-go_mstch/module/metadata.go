// @generated by Thrift for [[[ program path ]]]
// This file is probably not the place you want to edit!

package module // [[[ program thrift source path ]]]

import (
    "maps"

    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift"
    metadata "github.com/facebook/fbthrift/thrift/lib/thrift/metadata"
)

// (needed to ensure safety because of naive import list construction)
var _ = thrift.ZERO
var _ = maps.Copy[map[int]int, map[int]int]
var _ = metadata.GoUnusedProtection__

// Premade Thrift types
var (
    premadeThriftType_i64 = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_I64_TYPE.Ptr(),
            )
    premadeThriftType_string = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_STRING_TYPE.Ptr(),
            )
    premadeThriftType_module_MyDataItem = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("module.MyDataItem"),
            )
    premadeThriftType_module_MyEnum = metadata.NewThriftType().SetTEnum(
        metadata.NewThriftEnumType().
            SetName("module.MyEnum"),
            )
    premadeThriftType_bool = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_BOOL_TYPE.Ptr(),
            )
    premadeThriftType_float = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_FLOAT_TYPE.Ptr(),
            )
    premadeThriftType_set_float = metadata.NewThriftType().SetTSet(
        metadata.NewThriftSetType().
            SetValueType(premadeThriftType_float),
            )
    premadeThriftType_module_MyStruct = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("module.MyStruct"),
            )
    premadeThriftType_i32 = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_I32_TYPE.Ptr(),
            )
    premadeThriftType_void = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_VOID_TYPE.Ptr(),
            )
    premadeThriftType_module_ReservedKeyword = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("module.ReservedKeyword"),
            )
    premadeThriftType_binary = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_BINARY_TYPE.Ptr(),
            )
)

var structMetadatas = []*metadata.ThriftStruct{
    metadata.NewThriftStruct().
    SetName("module.MyStruct").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("MyIntField").
    SetIsOptional(false).
    SetType(premadeThriftType_i64),
            metadata.NewThriftField().
    SetId(2).
    SetName("MyStringField").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
            metadata.NewThriftField().
    SetId(3).
    SetName("MyDataField").
    SetIsOptional(false).
    SetType(premadeThriftType_module_MyDataItem),
            metadata.NewThriftField().
    SetId(4).
    SetName("myEnum").
    SetIsOptional(false).
    SetType(premadeThriftType_module_MyEnum),
            metadata.NewThriftField().
    SetId(5).
    SetName("oneway").
    SetIsOptional(false).
    SetType(premadeThriftType_bool),
            metadata.NewThriftField().
    SetId(6).
    SetName("readonly").
    SetIsOptional(false).
    SetType(premadeThriftType_bool),
            metadata.NewThriftField().
    SetId(7).
    SetName("idempotent").
    SetIsOptional(false).
    SetType(premadeThriftType_bool),
            metadata.NewThriftField().
    SetId(8).
    SetName("floatSet").
    SetIsOptional(false).
    SetType(premadeThriftType_set_float),
            metadata.NewThriftField().
    SetId(9).
    SetName("no_hack_codegen_field").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.MyDataItem").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("module.MyUnion").
    SetIsUnion(true).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("myEnum").
    SetIsOptional(false).
    SetType(premadeThriftType_module_MyEnum),
            metadata.NewThriftField().
    SetId(2).
    SetName("myStruct").
    SetIsOptional(false).
    SetType(premadeThriftType_module_MyStruct),
            metadata.NewThriftField().
    SetId(3).
    SetName("myDataItem").
    SetIsOptional(false).
    SetType(premadeThriftType_module_MyDataItem),
            metadata.NewThriftField().
    SetId(4).
    SetName("floatSet").
    SetIsOptional(false).
    SetType(premadeThriftType_set_float),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.ReservedKeyword").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("reserved_field").
    SetIsOptional(false).
    SetType(premadeThriftType_i32),
        },
    ),
    metadata.NewThriftStruct().
    SetName("module.UnionToBeRenamed").
    SetIsUnion(true).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("reserved_field").
    SetIsOptional(false).
    SetType(premadeThriftType_i32),
        },
    ),
}

var exceptionMetadatas = []*metadata.ThriftException{
}

var enumMetadatas = []*metadata.ThriftEnum{
    metadata.NewThriftEnum().
    SetName("module.MyEnum").
    SetElements(
        map[int32]string{
            0: "MyValue1",
            1: "MyValue2",
        },
    ),
    metadata.NewThriftEnum().
    SetName("module.HackEnum").
    SetElements(
        map[int32]string{
            0: "Value1",
            1: "Value2",
        },
    ),
}

var serviceMetadatas = []*metadata.ThriftService{
    metadata.NewThriftService().
    SetName("module.FooService").
    SetFunctions(
        []*metadata.ThriftFunction{
            metadata.NewThriftFunction().
    SetName("simple_rpc").
    SetIsOneway(false).
    SetReturnType(premadeThriftType_void),
        },
    ),
    metadata.NewThriftService().
    SetName("module.FB303Service").
    SetFunctions(
        []*metadata.ThriftFunction{
            metadata.NewThriftFunction().
    SetName("simple_rpc").
    SetIsOneway(false).
    SetReturnType(premadeThriftType_module_ReservedKeyword).
    SetArguments(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("int_parameter").
    SetIsOptional(false).
    SetType(premadeThriftType_i32),
        },
    ),
        },
    ),
    metadata.NewThriftService().
    SetName("module.MyService").
    SetFunctions(
        []*metadata.ThriftFunction{
            metadata.NewThriftFunction().
    SetName("ping").
    SetIsOneway(false).
    SetReturnType(premadeThriftType_void),
            metadata.NewThriftFunction().
    SetName("getRandomData").
    SetIsOneway(false).
    SetReturnType(premadeThriftType_string),
            metadata.NewThriftFunction().
    SetName("sink").
    SetIsOneway(false).
    SetReturnType(premadeThriftType_void).
    SetArguments(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("sink").
    SetIsOptional(false).
    SetType(premadeThriftType_i64),
        },
    ),
            metadata.NewThriftFunction().
    SetName("putDataById").
    SetIsOneway(false).
    SetReturnType(premadeThriftType_void).
    SetArguments(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("id").
    SetIsOptional(false).
    SetType(premadeThriftType_i64),
            metadata.NewThriftField().
    SetId(2).
    SetName("data").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
        },
    ),
            metadata.NewThriftFunction().
    SetName("hasDataById").
    SetIsOneway(false).
    SetReturnType(premadeThriftType_bool).
    SetArguments(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("id").
    SetIsOptional(false).
    SetType(premadeThriftType_i64),
        },
    ),
            metadata.NewThriftFunction().
    SetName("getDataById").
    SetIsOneway(false).
    SetReturnType(premadeThriftType_string).
    SetArguments(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("id").
    SetIsOptional(false).
    SetType(premadeThriftType_i64),
        },
    ),
            metadata.NewThriftFunction().
    SetName("deleteDataById").
    SetIsOneway(false).
    SetReturnType(premadeThriftType_void).
    SetArguments(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("id").
    SetIsOptional(false).
    SetType(premadeThriftType_i64),
        },
    ),
            metadata.NewThriftFunction().
    SetName("lobDataById").
    SetIsOneway(true).
    SetReturnType(premadeThriftType_void).
    SetArguments(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("id").
    SetIsOptional(false).
    SetType(premadeThriftType_i64),
            metadata.NewThriftField().
    SetId(2).
    SetName("data").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
        },
    ),
            metadata.NewThriftFunction().
    SetName("invalid_return_for_hack").
    SetIsOneway(false).
    SetReturnType(premadeThriftType_set_float),
            metadata.NewThriftFunction().
    SetName("rpc_skipped_codegen").
    SetIsOneway(false).
    SetReturnType(premadeThriftType_void),
        },
    ),
    metadata.NewThriftService().
    SetName("module.DbMixedStackArguments").
    SetFunctions(
        []*metadata.ThriftFunction{
            metadata.NewThriftFunction().
    SetName("getDataByKey0").
    SetIsOneway(false).
    SetReturnType(premadeThriftType_binary).
    SetArguments(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("key").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
        },
    ),
            metadata.NewThriftFunction().
    SetName("getDataByKey1").
    SetIsOneway(false).
    SetReturnType(premadeThriftType_binary).
    SetArguments(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("key").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
        },
    ),
        },
    ),
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
