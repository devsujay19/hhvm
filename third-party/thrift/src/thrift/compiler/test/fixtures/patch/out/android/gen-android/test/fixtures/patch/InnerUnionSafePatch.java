/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
package test.fixtures.patch;

import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.util.HashMap;
import java.util.Set;
import java.util.HashSet;
import java.util.Collections;
import java.util.BitSet;
import java.util.Arrays;
import com.facebook.thrift.*;
import com.facebook.thrift.annotations.*;
import com.facebook.thrift.async.*;
import com.facebook.thrift.meta_data.*;
import com.facebook.thrift.server.*;
import com.facebook.thrift.transport.*;
import com.facebook.thrift.protocol.*;

@SuppressWarnings({ "unused", "serial" })
public class InnerUnionSafePatch implements TBase, java.io.Serializable, Cloneable {
  private static final TStruct STRUCT_DESC = new TStruct("InnerUnionSafePatch");
  private static final TField VERSION_FIELD_DESC = new TField("version", TType.I32, (short)1);
  private static final TField DATA_FIELD_DESC = new TField("data", TType.STRING, (short)2);

  public final Integer version;
  public final byte[] data;
  public static final int VERSION = 1;
  public static final int DATA = 2;

  public InnerUnionSafePatch(
      Integer version,
      byte[] data) {
    this.version = version;
    this.data = data;
  }

  /**
   * Performs a deep copy on <i>other</i>.
   */
  public InnerUnionSafePatch(InnerUnionSafePatch other) {
    if (other.isSetVersion()) {
      this.version = TBaseHelper.deepCopy(other.version);
    } else {
      this.version = null;
    }
    if (other.isSetData()) {
      this.data = TBaseHelper.deepCopy(other.data);
    } else {
      this.data = null;
    }
  }

  public InnerUnionSafePatch deepCopy() {
    return new InnerUnionSafePatch(this);
  }

  public Integer getVersion() {
    return this.version;
  }

  // Returns true if field version is set (has been assigned a value) and false otherwise
  public boolean isSetVersion() {
    return this.version != null;
  }

  public byte[] getData() {
    return this.data;
  }

  // Returns true if field data is set (has been assigned a value) and false otherwise
  public boolean isSetData() {
    return this.data != null;
  }

  @Override
  public boolean equals(Object _that) {
    if (_that == null)
      return false;
    if (this == _that)
      return true;
    if (!(_that instanceof InnerUnionSafePatch))
      return false;
    InnerUnionSafePatch that = (InnerUnionSafePatch)_that;

    if (!TBaseHelper.equalsNobinary(this.isSetVersion(), that.isSetVersion(), this.version, that.version)) { return false; }

    if (!TBaseHelper.equalsSlow(this.isSetData(), that.isSetData(), this.data, that.data)) { return false; }

    return true;
  }

  @Override
  public int hashCode() {
    return Arrays.deepHashCode(new Object[] {version, data});
  }

  // This is required to satisfy the TBase interface, but can't be implemented on immutable struture.
  public void read(TProtocol iprot) throws TException {
    throw new TException("unimplemented in android immutable structure");
  }

  public static InnerUnionSafePatch deserialize(TProtocol iprot) throws TException {
    Integer tmp_version = null;
    byte[] tmp_data = null;
    TField __field;
    iprot.readStructBegin();
    while (true)
    {
      __field = iprot.readFieldBegin();
      if (__field.type == TType.STOP) {
        break;
      }
      switch (__field.id)
      {
        case VERSION:
          if (__field.type == TType.I32) {
            tmp_version = iprot.readI32();
          } else {
            TProtocolUtil.skip(iprot, __field.type);
          }
          break;
        case DATA:
          if (__field.type == TType.STRING) {
            tmp_data = iprot.readBinary();
          } else {
            TProtocolUtil.skip(iprot, __field.type);
          }
          break;
        default:
          TProtocolUtil.skip(iprot, __field.type);
          break;
      }
      iprot.readFieldEnd();
    }
    iprot.readStructEnd();

    InnerUnionSafePatch _that;
    _that = new InnerUnionSafePatch(
      tmp_version
      ,tmp_data
    );
    _that.validate();
    return _that;
  }

  public void write(TProtocol oprot) throws TException {
    validate();

    oprot.writeStructBegin(STRUCT_DESC);
    if (this.version != null) {
      oprot.writeFieldBegin(VERSION_FIELD_DESC);
      oprot.writeI32(this.version);
      oprot.writeFieldEnd();
    }
    if (this.data != null) {
      oprot.writeFieldBegin(DATA_FIELD_DESC);
      oprot.writeBinary(this.data);
      oprot.writeFieldEnd();
    }
    oprot.writeFieldStop();
    oprot.writeStructEnd();
  }

  @Override
  public String toString() {
    return toString(1, true);
  }

  @Override
  public String toString(int indent, boolean prettyPrint) {
    return TBaseHelper.toStringHelper(this, indent, prettyPrint);
  }

  public void validate() throws TException {
    // check for required fields
  }

}

