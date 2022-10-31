/**
 * Autogenerated by Thrift for thrift/compiler/test/fixtures/inject_metadata_fields/src/module.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated @nocommit
 */

#include "thrift/compiler/test/fixtures/inject_metadata_fields/gen-cpp2/module_data.h"

#include <thrift/lib/cpp2/gen/module_data_cpp.h>

FOLLY_CLANG_DISABLE_WARNING("-Wunused-macros")

#if defined(__GNUC__) && defined(__linux__) && !FOLLY_MOBILE
// These attributes are applied to the static data members to ensure that they
// are not stripped from the compiled binary, in order to keep them available
// for use by debuggers at runtime.
//
// The "used" attribute is required to ensure the compiler always emits unused
// data.
//
// The "section" attribute is required to stop the linker from stripping used
// data. It works by forcing all of the data members (both used and unused ones)
// into the same section. As the linker strips data on a per-section basis, it
// is then unable to remove unused data without also removing used data.
// This has a similar effect to the "retain" attribute, but works with older
// toolchains.
#define THRIFT_DATA_MEMBER [[gnu::used]] [[gnu::section(".rodata.thrift.data")]]
#else
#define THRIFT_DATA_MEMBER
#endif

namespace apache {
namespace thrift {

THRIFT_DATA_MEMBER const std::array<folly::StringPiece, 1> TStructDataStorage<::cpp2::Fields>::fields_names = {{
  "injected_field",
}};
THRIFT_DATA_MEMBER const std::array<int16_t, 1> TStructDataStorage<::cpp2::Fields>::fields_ids = {{
  100,
}};
THRIFT_DATA_MEMBER const std::array<protocol::TType, 1> TStructDataStorage<::cpp2::Fields>::fields_types = {{
  TType::T_STRING,
}};
THRIFT_DATA_MEMBER const std::array<folly::StringPiece, 1> TStructDataStorage<::cpp2::Fields>::storage_names = {{
  "__fbthrift_field_injected_field",
}};
THRIFT_DATA_MEMBER const std::array<int, 1> TStructDataStorage<::cpp2::Fields>::isset_indexes = {{
  0,
}};

THRIFT_DATA_MEMBER const std::array<folly::StringPiece, 1> TStructDataStorage<::cpp2::FieldsInjectedToEmptyStruct>::fields_names = {{
  "injected_field",
}};
THRIFT_DATA_MEMBER const std::array<int16_t, 1> TStructDataStorage<::cpp2::FieldsInjectedToEmptyStruct>::fields_ids = {{
  -1100,
}};
THRIFT_DATA_MEMBER const std::array<protocol::TType, 1> TStructDataStorage<::cpp2::FieldsInjectedToEmptyStruct>::fields_types = {{
  TType::T_STRING,
}};
THRIFT_DATA_MEMBER const std::array<folly::StringPiece, 1> TStructDataStorage<::cpp2::FieldsInjectedToEmptyStruct>::storage_names = {{
  "__fbthrift_field_injected_field",
}};
THRIFT_DATA_MEMBER const std::array<int, 1> TStructDataStorage<::cpp2::FieldsInjectedToEmptyStruct>::isset_indexes = {{
  0,
}};

THRIFT_DATA_MEMBER const std::array<folly::StringPiece, 2> TStructDataStorage<::cpp2::FieldsInjectedToStruct>::fields_names = {{
  "string_field",
  "injected_field",
}};
THRIFT_DATA_MEMBER const std::array<int16_t, 2> TStructDataStorage<::cpp2::FieldsInjectedToStruct>::fields_ids = {{
  1,
  -1100,
}};
THRIFT_DATA_MEMBER const std::array<protocol::TType, 2> TStructDataStorage<::cpp2::FieldsInjectedToStruct>::fields_types = {{
  TType::T_STRING,
  TType::T_STRING,
}};
THRIFT_DATA_MEMBER const std::array<folly::StringPiece, 2> TStructDataStorage<::cpp2::FieldsInjectedToStruct>::storage_names = {{
  "__fbthrift_field_string_field",
  "__fbthrift_field_injected_field",
}};
THRIFT_DATA_MEMBER const std::array<int, 2> TStructDataStorage<::cpp2::FieldsInjectedToStruct>::isset_indexes = {{
  0,
  1,
}};

THRIFT_DATA_MEMBER const std::array<folly::StringPiece, 4> TStructDataStorage<::cpp2::FieldsInjectedWithIncludedStruct>::fields_names = {{
  "string_field",
  "injected_field",
  "injected_structured_annotation_field",
  "injected_unstructured_annotation_field",
}};
THRIFT_DATA_MEMBER const std::array<int16_t, 4> TStructDataStorage<::cpp2::FieldsInjectedWithIncludedStruct>::fields_ids = {{
  1,
  -1100,
  -1101,
  -1102,
}};
THRIFT_DATA_MEMBER const std::array<protocol::TType, 4> TStructDataStorage<::cpp2::FieldsInjectedWithIncludedStruct>::fields_types = {{
  TType::T_STRING,
  TType::T_STRING,
  TType::T_STRING,
  TType::T_STRING,
}};
THRIFT_DATA_MEMBER const std::array<folly::StringPiece, 4> TStructDataStorage<::cpp2::FieldsInjectedWithIncludedStruct>::storage_names = {{
  "__fbthrift_field_string_field",
  "__fbthrift_field_injected_field",
  "__fbthrift_field_injected_structured_annotation_field",
  "__fbthrift_field_injected_unstructured_annotation_field",
}};
THRIFT_DATA_MEMBER const std::array<int, 4> TStructDataStorage<::cpp2::FieldsInjectedWithIncludedStruct>::isset_indexes = {{
  0,
  1,
  -1,
  -1,
}};

} // namespace thrift
} // namespace apache
