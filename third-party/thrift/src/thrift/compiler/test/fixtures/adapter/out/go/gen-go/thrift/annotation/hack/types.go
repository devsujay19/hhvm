// @generated by Thrift for [[[ program path ]]]
// This file is probably not the place you want to edit!

package hack // [[[ program thrift source path ]]]

import (
    "fmt"
    "strings"

    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction)
var _ = fmt.Printf
var _ = strings.Split
var _ = thrift.ZERO


type FieldWrapper struct {
    Name string `thrift:"name,1" json:"name" db:"name"`
}
// Compile time interface enforcer
var _ thrift.Struct = &FieldWrapper{}

func NewFieldWrapper() *FieldWrapper {
    return (&FieldWrapper{}).
        SetNameNonCompat("")
}

func (x *FieldWrapper) GetNameNonCompat() string {
    return x.Name
}

func (x *FieldWrapper) GetName() string {
    return x.Name
}

func (x *FieldWrapper) SetNameNonCompat(value string) *FieldWrapper {
    x.Name = value
    return x
}

func (x *FieldWrapper) SetName(value string) *FieldWrapper {
    x.Name = value
    return x
}

func (x *FieldWrapper) writeField1(p thrift.Format) error {  // Name
    if err := p.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Name
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *FieldWrapper) readField1(p thrift.Format) error {  // Name
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.Name = result
    return nil
}

func (x *FieldWrapper) toString1() string {  // Name
    return fmt.Sprintf("%v", x.Name)
}



func (x *FieldWrapper) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("FieldWrapper"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *FieldWrapper) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        case (id == 1 && wireType == thrift.Type(thrift.STRING)):  // name
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *FieldWrapper) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("FieldWrapper({")
    sb.WriteString(fmt.Sprintf("Name:%s", x.toString1()))
    sb.WriteString("})")

    return sb.String()
}

type Wrapper struct {
    Name string `thrift:"name,1" json:"name" db:"name"`
    UnderlyingName string `thrift:"underlyingName,2" json:"underlyingName" db:"underlyingName"`
    ExtraNamespace string `thrift:"extraNamespace,3" json:"extraNamespace" db:"extraNamespace"`
}
// Compile time interface enforcer
var _ thrift.Struct = &Wrapper{}

func NewWrapper() *Wrapper {
    return (&Wrapper{}).
        SetNameNonCompat("").
        SetUnderlyingNameNonCompat("").
        SetExtraNamespaceNonCompat("thrift_adapted_types")
}

func (x *Wrapper) GetNameNonCompat() string {
    return x.Name
}

func (x *Wrapper) GetName() string {
    return x.Name
}

func (x *Wrapper) GetUnderlyingNameNonCompat() string {
    return x.UnderlyingName
}

func (x *Wrapper) GetUnderlyingName() string {
    return x.UnderlyingName
}

func (x *Wrapper) GetExtraNamespaceNonCompat() string {
    return x.ExtraNamespace
}

func (x *Wrapper) GetExtraNamespace() string {
    return x.ExtraNamespace
}

func (x *Wrapper) SetNameNonCompat(value string) *Wrapper {
    x.Name = value
    return x
}

func (x *Wrapper) SetName(value string) *Wrapper {
    x.Name = value
    return x
}

func (x *Wrapper) SetUnderlyingNameNonCompat(value string) *Wrapper {
    x.UnderlyingName = value
    return x
}

func (x *Wrapper) SetUnderlyingName(value string) *Wrapper {
    x.UnderlyingName = value
    return x
}

func (x *Wrapper) SetExtraNamespaceNonCompat(value string) *Wrapper {
    x.ExtraNamespace = value
    return x
}

func (x *Wrapper) SetExtraNamespace(value string) *Wrapper {
    x.ExtraNamespace = value
    return x
}

func (x *Wrapper) writeField1(p thrift.Format) error {  // Name
    if err := p.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Name
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Wrapper) writeField2(p thrift.Format) error {  // UnderlyingName
    if err := p.WriteFieldBegin("underlyingName", thrift.STRING, 2); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.UnderlyingName
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Wrapper) writeField3(p thrift.Format) error {  // ExtraNamespace
    if err := p.WriteFieldBegin("extraNamespace", thrift.STRING, 3); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.ExtraNamespace
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Wrapper) readField1(p thrift.Format) error {  // Name
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.Name = result
    return nil
}

func (x *Wrapper) readField2(p thrift.Format) error {  // UnderlyingName
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.UnderlyingName = result
    return nil
}

func (x *Wrapper) readField3(p thrift.Format) error {  // ExtraNamespace
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.ExtraNamespace = result
    return nil
}

func (x *Wrapper) toString1() string {  // Name
    return fmt.Sprintf("%v", x.Name)
}

func (x *Wrapper) toString2() string {  // UnderlyingName
    return fmt.Sprintf("%v", x.UnderlyingName)
}

func (x *Wrapper) toString3() string {  // ExtraNamespace
    return fmt.Sprintf("%v", x.ExtraNamespace)
}



func (x *Wrapper) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("Wrapper"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := x.writeField2(p); err != nil {
        return err
    }

    if err := x.writeField3(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Wrapper) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        case (id == 1 && wireType == thrift.Type(thrift.STRING)):  // name
            if err := x.readField1(p); err != nil {
                return err
            }
        case (id == 2 && wireType == thrift.Type(thrift.STRING)):  // underlyingName
            if err := x.readField2(p); err != nil {
                return err
            }
        case (id == 3 && wireType == thrift.Type(thrift.STRING)):  // extraNamespace
            if err := x.readField3(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *Wrapper) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("Wrapper({")
    sb.WriteString(fmt.Sprintf("Name:%s ", x.toString1()))
    sb.WriteString(fmt.Sprintf("UnderlyingName:%s ", x.toString2()))
    sb.WriteString(fmt.Sprintf("ExtraNamespace:%s", x.toString3()))
    sb.WriteString("})")

    return sb.String()
}

type Adapter struct {
    Name string `thrift:"name,1" json:"name" db:"name"`
}
// Compile time interface enforcer
var _ thrift.Struct = &Adapter{}

func NewAdapter() *Adapter {
    return (&Adapter{}).
        SetNameNonCompat("")
}

func (x *Adapter) GetNameNonCompat() string {
    return x.Name
}

func (x *Adapter) GetName() string {
    return x.Name
}

func (x *Adapter) SetNameNonCompat(value string) *Adapter {
    x.Name = value
    return x
}

func (x *Adapter) SetName(value string) *Adapter {
    x.Name = value
    return x
}

func (x *Adapter) writeField1(p thrift.Format) error {  // Name
    if err := p.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Name
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Adapter) readField1(p thrift.Format) error {  // Name
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.Name = result
    return nil
}

func (x *Adapter) toString1() string {  // Name
    return fmt.Sprintf("%v", x.Name)
}



func (x *Adapter) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("Adapter"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Adapter) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        case (id == 1 && wireType == thrift.Type(thrift.STRING)):  // name
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *Adapter) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("Adapter({")
    sb.WriteString(fmt.Sprintf("Name:%s", x.toString1()))
    sb.WriteString("})")

    return sb.String()
}

type SkipCodegen struct {
    Reason string `thrift:"reason,1" json:"reason" db:"reason"`
}
// Compile time interface enforcer
var _ thrift.Struct = &SkipCodegen{}

func NewSkipCodegen() *SkipCodegen {
    return (&SkipCodegen{}).
        SetReasonNonCompat("")
}

func (x *SkipCodegen) GetReasonNonCompat() string {
    return x.Reason
}

func (x *SkipCodegen) GetReason() string {
    return x.Reason
}

func (x *SkipCodegen) SetReasonNonCompat(value string) *SkipCodegen {
    x.Reason = value
    return x
}

func (x *SkipCodegen) SetReason(value string) *SkipCodegen {
    x.Reason = value
    return x
}

func (x *SkipCodegen) writeField1(p thrift.Format) error {  // Reason
    if err := p.WriteFieldBegin("reason", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Reason
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *SkipCodegen) readField1(p thrift.Format) error {  // Reason
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.Reason = result
    return nil
}

func (x *SkipCodegen) toString1() string {  // Reason
    return fmt.Sprintf("%v", x.Reason)
}



func (x *SkipCodegen) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("SkipCodegen"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *SkipCodegen) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        case (id == 1 && wireType == thrift.Type(thrift.STRING)):  // reason
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *SkipCodegen) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("SkipCodegen({")
    sb.WriteString(fmt.Sprintf("Reason:%s", x.toString1()))
    sb.WriteString("})")

    return sb.String()
}

type Name struct {
    Name string `thrift:"name,1" json:"name" db:"name"`
    Reason string `thrift:"reason,2" json:"reason" db:"reason"`
}
// Compile time interface enforcer
var _ thrift.Struct = &Name{}

func NewName() *Name {
    return (&Name{}).
        SetNameNonCompat("").
        SetReasonNonCompat("")
}

func (x *Name) GetNameNonCompat() string {
    return x.Name
}

func (x *Name) GetName() string {
    return x.Name
}

func (x *Name) GetReasonNonCompat() string {
    return x.Reason
}

func (x *Name) GetReason() string {
    return x.Reason
}

func (x *Name) SetNameNonCompat(value string) *Name {
    x.Name = value
    return x
}

func (x *Name) SetName(value string) *Name {
    x.Name = value
    return x
}

func (x *Name) SetReasonNonCompat(value string) *Name {
    x.Reason = value
    return x
}

func (x *Name) SetReason(value string) *Name {
    x.Reason = value
    return x
}

func (x *Name) writeField1(p thrift.Format) error {  // Name
    if err := p.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Name
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Name) writeField2(p thrift.Format) error {  // Reason
    if err := p.WriteFieldBegin("reason", thrift.STRING, 2); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Reason
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Name) readField1(p thrift.Format) error {  // Name
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.Name = result
    return nil
}

func (x *Name) readField2(p thrift.Format) error {  // Reason
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.Reason = result
    return nil
}

func (x *Name) toString1() string {  // Name
    return fmt.Sprintf("%v", x.Name)
}

func (x *Name) toString2() string {  // Reason
    return fmt.Sprintf("%v", x.Reason)
}



func (x *Name) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("Name"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := x.writeField2(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Name) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        case (id == 1 && wireType == thrift.Type(thrift.STRING)):  // name
            if err := x.readField1(p); err != nil {
                return err
            }
        case (id == 2 && wireType == thrift.Type(thrift.STRING)):  // reason
            if err := x.readField2(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *Name) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("Name({")
    sb.WriteString(fmt.Sprintf("Name:%s ", x.toString1()))
    sb.WriteString(fmt.Sprintf("Reason:%s", x.toString2()))
    sb.WriteString("})")

    return sb.String()
}

type UnionEnumAttributes struct {
    Attributes []string `thrift:"attributes,1" json:"attributes" db:"attributes"`
}
// Compile time interface enforcer
var _ thrift.Struct = &UnionEnumAttributes{}

func NewUnionEnumAttributes() *UnionEnumAttributes {
    return (&UnionEnumAttributes{}).
        SetAttributesNonCompat(make([]string, 0))
}

func (x *UnionEnumAttributes) GetAttributesNonCompat() []string {
    return x.Attributes
}

func (x *UnionEnumAttributes) GetAttributes() []string {
    if !x.IsSetAttributes() {
        return make([]string, 0)
    }

    return x.Attributes
}

func (x *UnionEnumAttributes) SetAttributesNonCompat(value []string) *UnionEnumAttributes {
    x.Attributes = value
    return x
}

func (x *UnionEnumAttributes) SetAttributes(value []string) *UnionEnumAttributes {
    x.Attributes = value
    return x
}

func (x *UnionEnumAttributes) IsSetAttributes() bool {
    return x != nil && x.Attributes != nil
}

func (x *UnionEnumAttributes) writeField1(p thrift.Format) error {  // Attributes
    if err := p.WriteFieldBegin("attributes", thrift.LIST, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Attributes
    if err := p.WriteListBegin(thrift.STRING, len(item)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
}
for _, v := range item {
    {
        item := v
        if err := p.WriteString(item); err != nil {
    return err
}
    }
}
if err := p.WriteListEnd(); err != nil {
    return thrift.PrependError("error writing list end: ", err)
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *UnionEnumAttributes) readField1(p thrift.Format) error {  // Attributes
    _ /* elemType */, size, err := p.ReadListBegin()
if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
}

listResult := make([]string, 0, size)
for i := 0; i < size; i++ {
    var elem string
    {
        result, err := p.ReadString()
if err != nil {
    return err
}
        elem = result
    }
    listResult = append(listResult, elem)
}

if err := p.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
}
result := listResult

    x.Attributes = result
    return nil
}

func (x *UnionEnumAttributes) toString1() string {  // Attributes
    return fmt.Sprintf("%v", x.Attributes)
}



func (x *UnionEnumAttributes) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("UnionEnumAttributes"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *UnionEnumAttributes) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        case (id == 1 && wireType == thrift.Type(thrift.LIST)):  // attributes
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *UnionEnumAttributes) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("UnionEnumAttributes({")
    sb.WriteString(fmt.Sprintf("Attributes:%s", x.toString1()))
    sb.WriteString("})")

    return sb.String()
}

type StructTrait struct {
    Name string `thrift:"name,1" json:"name" db:"name"`
}
// Compile time interface enforcer
var _ thrift.Struct = &StructTrait{}

func NewStructTrait() *StructTrait {
    return (&StructTrait{}).
        SetNameNonCompat("")
}

func (x *StructTrait) GetNameNonCompat() string {
    return x.Name
}

func (x *StructTrait) GetName() string {
    return x.Name
}

func (x *StructTrait) SetNameNonCompat(value string) *StructTrait {
    x.Name = value
    return x
}

func (x *StructTrait) SetName(value string) *StructTrait {
    x.Name = value
    return x
}

func (x *StructTrait) writeField1(p thrift.Format) error {  // Name
    if err := p.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Name
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *StructTrait) readField1(p thrift.Format) error {  // Name
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.Name = result
    return nil
}

func (x *StructTrait) toString1() string {  // Name
    return fmt.Sprintf("%v", x.Name)
}



func (x *StructTrait) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("StructTrait"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *StructTrait) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        case (id == 1 && wireType == thrift.Type(thrift.STRING)):  // name
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *StructTrait) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("StructTrait({")
    sb.WriteString(fmt.Sprintf("Name:%s", x.toString1()))
    sb.WriteString("})")

    return sb.String()
}

type Attributes struct {
    Attributes []string `thrift:"attributes,1" json:"attributes" db:"attributes"`
}
// Compile time interface enforcer
var _ thrift.Struct = &Attributes{}

func NewAttributes() *Attributes {
    return (&Attributes{}).
        SetAttributesNonCompat(make([]string, 0))
}

func (x *Attributes) GetAttributesNonCompat() []string {
    return x.Attributes
}

func (x *Attributes) GetAttributes() []string {
    if !x.IsSetAttributes() {
        return make([]string, 0)
    }

    return x.Attributes
}

func (x *Attributes) SetAttributesNonCompat(value []string) *Attributes {
    x.Attributes = value
    return x
}

func (x *Attributes) SetAttributes(value []string) *Attributes {
    x.Attributes = value
    return x
}

func (x *Attributes) IsSetAttributes() bool {
    return x != nil && x.Attributes != nil
}

func (x *Attributes) writeField1(p thrift.Format) error {  // Attributes
    if err := p.WriteFieldBegin("attributes", thrift.LIST, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Attributes
    if err := p.WriteListBegin(thrift.STRING, len(item)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
}
for _, v := range item {
    {
        item := v
        if err := p.WriteString(item); err != nil {
    return err
}
    }
}
if err := p.WriteListEnd(); err != nil {
    return thrift.PrependError("error writing list end: ", err)
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Attributes) readField1(p thrift.Format) error {  // Attributes
    _ /* elemType */, size, err := p.ReadListBegin()
if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
}

listResult := make([]string, 0, size)
for i := 0; i < size; i++ {
    var elem string
    {
        result, err := p.ReadString()
if err != nil {
    return err
}
        elem = result
    }
    listResult = append(listResult, elem)
}

if err := p.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
}
result := listResult

    x.Attributes = result
    return nil
}

func (x *Attributes) toString1() string {  // Attributes
    return fmt.Sprintf("%v", x.Attributes)
}



func (x *Attributes) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("Attributes"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Attributes) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        case (id == 1 && wireType == thrift.Type(thrift.LIST)):  // attributes
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *Attributes) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("Attributes({")
    sb.WriteString(fmt.Sprintf("Attributes:%s", x.toString1()))
    sb.WriteString("})")

    return sb.String()
}

type StructAsTrait struct {
}
// Compile time interface enforcer
var _ thrift.Struct = &StructAsTrait{}

func NewStructAsTrait() *StructAsTrait {
    return (&StructAsTrait{})
}



func (x *StructAsTrait) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("StructAsTrait"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *StructAsTrait) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *StructAsTrait) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("StructAsTrait({")
    sb.WriteString("})")

    return sb.String()
}

type ModuleInternal struct {
}
// Compile time interface enforcer
var _ thrift.Struct = &ModuleInternal{}

func NewModuleInternal() *ModuleInternal {
    return (&ModuleInternal{})
}



func (x *ModuleInternal) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("ModuleInternal"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *ModuleInternal) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *ModuleInternal) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("ModuleInternal({")
    sb.WriteString("})")

    return sb.String()
}

// RegisterTypes registers types found in this file that have a thrift_uri with the passed in registry.
func RegisterTypes(registry interface {
  RegisterType(name string, initializer func() any)
}) {
    registry.RegisterType("facebook.com/thrift/annotation/hack/FieldWrapper", func() any { return NewFieldWrapper() })
    registry.RegisterType("facebook.com/thrift/annotation/hack/Wrapper", func() any { return NewWrapper() })
    registry.RegisterType("facebook.com/thrift/annotation/hack/Adapter", func() any { return NewAdapter() })
    registry.RegisterType("facebook.com/thrift/annotation/hack/SkipCodegen", func() any { return NewSkipCodegen() })
    registry.RegisterType("facebook.com/thrift/annotation/hack/Name", func() any { return NewName() })
    registry.RegisterType("facebook.com/thrift/annotation/hack/UnionEnumAttributes", func() any { return NewUnionEnumAttributes() })
    registry.RegisterType("facebook.com/thrift/annotation/hack/StructTrait", func() any { return NewStructTrait() })
    registry.RegisterType("facebook.com/thrift/annotation/hack/Attributes", func() any { return NewAttributes() })
    registry.RegisterType("facebook.com/thrift/annotation/hack/StructAsTrait", func() any { return NewStructAsTrait() })
    registry.RegisterType("facebook.com/thrift/annotation/hack/ModuleInternal", func() any { return NewModuleInternal() })

}
