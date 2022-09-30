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
public class MyStructFieldN28Patch implements TBase, java.io.Serializable, Cloneable {
  private static final TStruct STRUCT_DESC = new TStruct("MyStructFieldN28Patch");
  private static final TField ASSIGN_FIELD_DESC = new TField("assign", TType.LIST, (short)1);
  private static final TField CLEAR_FIELD_DESC = new TField("clear", TType.BOOL, (short)2);
  private static final TField PATCH_FIELD_DESC = new TField("patch", TType.MAP, (short)3);
  private static final TField REMOVE_FIELD_DESC = new TField("remove", TType.LIST, (short)7);
  private static final TField PREPEND_FIELD_DESC = new TField("prepend", TType.LIST, (short)8);
  private static final TField APPEND_FIELD_DESC = new TField("append", TType.LIST, (short)9);

  /**
   * Assigns a value. If set, all other operations are ignored.
   */
  public final List<Short> assign;
  /**
   * Clears a value. Applies first.
   */
  public final Boolean clear;
  /**
   * Patches list values by index. Applies second.
   */
  public final Map<Integer,I16Patch> patch;
  /**
   * Removes entries, if present. Applies third.
   */
  public final List<Short> remove;
  /**
   * Prepends to the front of a given list.
   */
  public final List<Short> prepend;
  /**
   * Appends to the back of a given list.
   */
  public final List<Short> append;
  public static final int ASSIGN = 1;
  public static final int CLEAR = 2;
  public static final int PATCH = 3;
  public static final int REMOVE = 7;
  public static final int PREPEND = 8;
  public static final int APPEND = 9;

  public MyStructFieldN28Patch(
      List<Short> assign,
      Boolean clear,
      Map<Integer,I16Patch> patch,
      List<Short> remove,
      List<Short> prepend,
      List<Short> append) {
    this.assign = assign;
    this.clear = clear;
    this.patch = patch;
    this.remove = remove;
    this.prepend = prepend;
    this.append = append;
  }

  /**
   * Performs a deep copy on <i>other</i>.
   */
  public MyStructFieldN28Patch(MyStructFieldN28Patch other) {
    if (other.isSetAssign()) {
      this.assign = TBaseHelper.deepCopy(other.assign);
    } else {
      this.assign = null;
    }
    if (other.isSetClear()) {
      this.clear = TBaseHelper.deepCopy(other.clear);
    } else {
      this.clear = null;
    }
    if (other.isSetPatch()) {
      this.patch = TBaseHelper.deepCopy(other.patch);
    } else {
      this.patch = null;
    }
    if (other.isSetRemove()) {
      this.remove = TBaseHelper.deepCopy(other.remove);
    } else {
      this.remove = null;
    }
    if (other.isSetPrepend()) {
      this.prepend = TBaseHelper.deepCopy(other.prepend);
    } else {
      this.prepend = null;
    }
    if (other.isSetAppend()) {
      this.append = TBaseHelper.deepCopy(other.append);
    } else {
      this.append = null;
    }
  }

  public MyStructFieldN28Patch deepCopy() {
    return new MyStructFieldN28Patch(this);
  }

  /**
   * Assigns a value. If set, all other operations are ignored.
   */
  public List<Short> getAssign() {
    return this.assign;
  }

  // Returns true if field assign is set (has been assigned a value) and false otherwise
  public boolean isSetAssign() {
    return this.assign != null;
  }

  /**
   * Clears a value. Applies first.
   */
  public Boolean isClear() {
    return this.clear;
  }

  // Returns true if field clear is set (has been assigned a value) and false otherwise
  public boolean isSetClear() {
    return this.clear != null;
  }

  /**
   * Patches list values by index. Applies second.
   */
  public Map<Integer,I16Patch> getPatch() {
    return this.patch;
  }

  // Returns true if field patch is set (has been assigned a value) and false otherwise
  public boolean isSetPatch() {
    return this.patch != null;
  }

  /**
   * Removes entries, if present. Applies third.
   */
  public List<Short> getRemove() {
    return this.remove;
  }

  // Returns true if field remove is set (has been assigned a value) and false otherwise
  public boolean isSetRemove() {
    return this.remove != null;
  }

  /**
   * Prepends to the front of a given list.
   */
  public List<Short> getPrepend() {
    return this.prepend;
  }

  // Returns true if field prepend is set (has been assigned a value) and false otherwise
  public boolean isSetPrepend() {
    return this.prepend != null;
  }

  /**
   * Appends to the back of a given list.
   */
  public List<Short> getAppend() {
    return this.append;
  }

  // Returns true if field append is set (has been assigned a value) and false otherwise
  public boolean isSetAppend() {
    return this.append != null;
  }

  @Override
  public boolean equals(Object _that) {
    if (_that == null)
      return false;
    if (this == _that)
      return true;
    if (!(_that instanceof MyStructFieldN28Patch))
      return false;
    MyStructFieldN28Patch that = (MyStructFieldN28Patch)_that;

    if (!TBaseHelper.equalsNobinary(this.isSetAssign(), that.isSetAssign(), this.assign, that.assign)) { return false; }

    if (!TBaseHelper.equalsNobinary(this.isSetClear(), that.isSetClear(), this.clear, that.clear)) { return false; }

    if (!TBaseHelper.equalsNobinary(this.isSetPatch(), that.isSetPatch(), this.patch, that.patch)) { return false; }

    if (!TBaseHelper.equalsNobinary(this.isSetRemove(), that.isSetRemove(), this.remove, that.remove)) { return false; }

    if (!TBaseHelper.equalsNobinary(this.isSetPrepend(), that.isSetPrepend(), this.prepend, that.prepend)) { return false; }

    if (!TBaseHelper.equalsNobinary(this.isSetAppend(), that.isSetAppend(), this.append, that.append)) { return false; }

    return true;
  }

  @Override
  public int hashCode() {
    return Arrays.deepHashCode(new Object[] {assign, clear, patch, remove, prepend, append});
  }

  // This is required to satisfy the TBase interface, but can't be implemented on immutable struture.
  public void read(TProtocol iprot) throws TException {
    throw new TException("unimplemented in android immutable structure");
  }

  public static MyStructFieldN28Patch deserialize(TProtocol iprot) throws TException {
    List<Short> tmp_assign = null;
    Boolean tmp_clear = null;
    Map<Integer,I16Patch> tmp_patch = null;
    List<Short> tmp_remove = null;
    List<Short> tmp_prepend = null;
    List<Short> tmp_append = null;
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
        case ASSIGN:
          if (__field.type == TType.LIST) {
            {
              TList _list13 = iprot.readListBegin();
              tmp_assign = new ArrayList<Short>(Math.max(0, _list13.size));
              for (int _i14 = 0; 
                   (_list13.size < 0) ? iprot.peekList() : (_i14 < _list13.size); 
                   ++_i14)
              {
                Short _elem15;
                _elem15 = iprot.readI16();
                tmp_assign.add(_elem15);
              }
              iprot.readListEnd();
            }
          } else {
            TProtocolUtil.skip(iprot, __field.type);
          }
          break;
        case CLEAR:
          if (__field.type == TType.BOOL) {
            tmp_clear = iprot.readBool();
          } else {
            TProtocolUtil.skip(iprot, __field.type);
          }
          break;
        case PATCH:
          if (__field.type == TType.MAP) {
            {
              TMap _map16 = iprot.readMapBegin();
              tmp_patch = new HashMap<Integer,I16Patch>(Math.max(0, 2*_map16.size));
              for (int _i17 = 0; 
                   (_map16.size < 0) ? iprot.peekMap() : (_i17 < _map16.size); 
                   ++_i17)
              {
                Integer _key18;
                I16Patch _val19;
                _key18 = iprot.readI32();
                _val19 = I16Patch.deserialize(iprot);
                tmp_patch.put(_key18, _val19);
              }
              iprot.readMapEnd();
            }
          } else {
            TProtocolUtil.skip(iprot, __field.type);
          }
          break;
        case REMOVE:
          if (__field.type == TType.LIST) {
            {
              TList _list20 = iprot.readListBegin();
              tmp_remove = new ArrayList<Short>(Math.max(0, _list20.size));
              for (int _i21 = 0; 
                   (_list20.size < 0) ? iprot.peekList() : (_i21 < _list20.size); 
                   ++_i21)
              {
                Short _elem22;
                _elem22 = iprot.readI16();
                tmp_remove.add(_elem22);
              }
              iprot.readListEnd();
            }
          } else {
            TProtocolUtil.skip(iprot, __field.type);
          }
          break;
        case PREPEND:
          if (__field.type == TType.LIST) {
            {
              TList _list23 = iprot.readListBegin();
              tmp_prepend = new ArrayList<Short>(Math.max(0, _list23.size));
              for (int _i24 = 0; 
                   (_list23.size < 0) ? iprot.peekList() : (_i24 < _list23.size); 
                   ++_i24)
              {
                Short _elem25;
                _elem25 = iprot.readI16();
                tmp_prepend.add(_elem25);
              }
              iprot.readListEnd();
            }
          } else {
            TProtocolUtil.skip(iprot, __field.type);
          }
          break;
        case APPEND:
          if (__field.type == TType.LIST) {
            {
              TList _list26 = iprot.readListBegin();
              tmp_append = new ArrayList<Short>(Math.max(0, _list26.size));
              for (int _i27 = 0; 
                   (_list26.size < 0) ? iprot.peekList() : (_i27 < _list26.size); 
                   ++_i27)
              {
                Short _elem28;
                _elem28 = iprot.readI16();
                tmp_append.add(_elem28);
              }
              iprot.readListEnd();
            }
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

    MyStructFieldN28Patch _that;
    _that = new MyStructFieldN28Patch(
      tmp_assign
      ,tmp_clear
      ,tmp_patch
      ,tmp_remove
      ,tmp_prepend
      ,tmp_append
    );
    _that.validate();
    return _that;
  }

  public void write(TProtocol oprot) throws TException {
    validate();

    oprot.writeStructBegin(STRUCT_DESC);
    if (this.assign != null) {
      if (isSetAssign()) {
        oprot.writeFieldBegin(ASSIGN_FIELD_DESC);
        {
          oprot.writeListBegin(new TList(TType.I16, this.assign.size()));
          for (Short _iter29 : this.assign)          {
            oprot.writeI16(_iter29);
          }
          oprot.writeListEnd();
        }
        oprot.writeFieldEnd();
      }
    }
    if (this.clear != null) {
      oprot.writeFieldBegin(CLEAR_FIELD_DESC);
      oprot.writeBool(this.clear);
      oprot.writeFieldEnd();
    }
    if (this.patch != null) {
      oprot.writeFieldBegin(PATCH_FIELD_DESC);
      {
        oprot.writeMapBegin(new TMap(TType.I32, TType.STRUCT, this.patch.size()));
        for (Map.Entry<Integer, I16Patch> _iter30 : this.patch.entrySet())        {
          oprot.writeI32(_iter30.getKey());
          _iter30.getValue().write(oprot);
        }
        oprot.writeMapEnd();
      }
      oprot.writeFieldEnd();
    }
    if (this.remove != null) {
      oprot.writeFieldBegin(REMOVE_FIELD_DESC);
      {
        oprot.writeListBegin(new TList(TType.I16, this.remove.size()));
        for (Short _iter31 : this.remove)        {
          oprot.writeI16(_iter31);
        }
        oprot.writeListEnd();
      }
      oprot.writeFieldEnd();
    }
    if (this.prepend != null) {
      oprot.writeFieldBegin(PREPEND_FIELD_DESC);
      {
        oprot.writeListBegin(new TList(TType.I16, this.prepend.size()));
        for (Short _iter32 : this.prepend)        {
          oprot.writeI16(_iter32);
        }
        oprot.writeListEnd();
      }
      oprot.writeFieldEnd();
    }
    if (this.append != null) {
      oprot.writeFieldBegin(APPEND_FIELD_DESC);
      {
        oprot.writeListBegin(new TList(TType.I16, this.append.size()));
        for (Short _iter33 : this.append)        {
          oprot.writeI16(_iter33);
        }
        oprot.writeListEnd();
      }
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

