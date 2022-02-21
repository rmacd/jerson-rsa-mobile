// automatically generated by the FlatBuffers compiler, do not modify

import * as flatbuffers from 'flatbuffers';

export class PrivateKeyInfo {
  bb: flatbuffers.ByteBuffer|null = null;
  bb_pos = 0;
__init(i:number, bb:flatbuffers.ByteBuffer):PrivateKeyInfo {
  this.bb_pos = i;
  this.bb = bb;
  return this;
}

static getRootAsPrivateKeyInfo(bb:flatbuffers.ByteBuffer, obj?:PrivateKeyInfo):PrivateKeyInfo {
  return (obj || new PrivateKeyInfo()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
}

static getSizePrefixedRootAsPrivateKeyInfo(bb:flatbuffers.ByteBuffer, obj?:PrivateKeyInfo):PrivateKeyInfo {
  bb.setPosition(bb.position() + flatbuffers.SIZE_PREFIX_LENGTH);
  return (obj || new PrivateKeyInfo()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
}

bitLen():flatbuffers.Long {
  const offset = this.bb!.__offset(this.bb_pos, 4);
  return offset ? this.bb!.readInt64(this.bb_pos + offset) : this.bb!.createLong(0, 0);
}

mutate_bitLen(value:flatbuffers.Long):boolean {
  const offset = this.bb!.__offset(this.bb_pos, 4);

  if (offset === 0) {
    return false;
  }

  this.bb!.writeInt64(this.bb_pos + offset, value);
  return true;
}

size():flatbuffers.Long {
  const offset = this.bb!.__offset(this.bb_pos, 6);
  return offset ? this.bb!.readInt64(this.bb_pos + offset) : this.bb!.createLong(0, 0);
}

mutate_size(value:flatbuffers.Long):boolean {
  const offset = this.bb!.__offset(this.bb_pos, 6);

  if (offset === 0) {
    return false;
  }

  this.bb!.writeInt64(this.bb_pos + offset, value);
  return true;
}

error():string|null
error(optionalEncoding:flatbuffers.Encoding):string|Uint8Array|null
error(optionalEncoding?:any):string|Uint8Array|null {
  const offset = this.bb!.__offset(this.bb_pos, 8);
  return offset ? this.bb!.__string(this.bb_pos + offset, optionalEncoding) : null;
}

static startPrivateKeyInfo(builder:flatbuffers.Builder) {
  builder.startObject(3);
}

static addBitLen(builder:flatbuffers.Builder, bitLen:flatbuffers.Long) {
  builder.addFieldInt64(0, bitLen, builder.createLong(0, 0));
}

static addSize(builder:flatbuffers.Builder, size:flatbuffers.Long) {
  builder.addFieldInt64(1, size, builder.createLong(0, 0));
}

static addError(builder:flatbuffers.Builder, errorOffset:flatbuffers.Offset) {
  builder.addFieldOffset(2, errorOffset, 0);
}

static endPrivateKeyInfo(builder:flatbuffers.Builder):flatbuffers.Offset {
  const offset = builder.endObject();
  return offset;
}

static createPrivateKeyInfo(builder:flatbuffers.Builder, bitLen:flatbuffers.Long, size:flatbuffers.Long, errorOffset:flatbuffers.Offset):flatbuffers.Offset {
  PrivateKeyInfo.startPrivateKeyInfo(builder);
  PrivateKeyInfo.addBitLen(builder, bitLen);
  PrivateKeyInfo.addSize(builder, size);
  PrivateKeyInfo.addError(builder, errorOffset);
  return PrivateKeyInfo.endPrivateKeyInfo(builder);
}
}