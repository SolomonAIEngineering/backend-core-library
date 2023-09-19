/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "gorm";

export interface GormFileOptions {
}

export interface GormMessageOptions {
  ormable: boolean;
  include: ExtraField[];
  table: string;
  multiAccount: boolean;
}

export interface ExtraField {
  type: string;
  name: string;
  tag: GormTag | undefined;
  package: string;
}

export interface GormFieldOptions {
  tag: GormTag | undefined;
  drop: boolean;
  hasOne?: HasOneOptions | undefined;
  belongsTo?: BelongsToOptions | undefined;
  hasMany?: HasManyOptions | undefined;
  manyToMany?: ManyToManyOptions | undefined;
  referenceOf: string;
}

export interface GormTag {
  column: string;
  type: string;
  size: number;
  precision: number;
  primaryKey: boolean;
  unique: boolean;
  default: string;
  notNull: boolean;
  autoIncrement: boolean;
  index: string;
  uniqueIndex: string;
  embedded: boolean;
  embeddedPrefix: string;
  ignore: boolean;
  foreignkey: string;
  associationForeignkey: string;
  manyToMany: string;
  jointableForeignkey: string;
  associationJointableForeignkey: string;
  disableAssociationAutoupdate: boolean;
  disableAssociationAutocreate: boolean;
  associationSaveReference: boolean;
  preload: boolean;
  serializer: string;
}

export interface HasOneOptions {
  foreignkey: string;
  foreignkeyTag: GormTag | undefined;
  associationForeignkey: string;
  disableAssociationAutoupdate: boolean;
  disableAssociationAutocreate: boolean;
  associationSaveReference: boolean;
  preload: boolean;
  replace: boolean;
  append: boolean;
  clear: boolean;
}

export interface BelongsToOptions {
  foreignkey: string;
  foreignkeyTag: GormTag | undefined;
  associationForeignkey: string;
  disableAssociationAutoupdate: boolean;
  disableAssociationAutocreate: boolean;
  associationSaveReference: boolean;
  preload: boolean;
}

export interface HasManyOptions {
  foreignkey: string;
  foreignkeyTag: GormTag | undefined;
  associationForeignkey: string;
  positionField: string;
  positionFieldTag: GormTag | undefined;
  disableAssociationAutoupdate: boolean;
  disableAssociationAutocreate: boolean;
  associationSaveReference: boolean;
  preload: boolean;
  replace: boolean;
  append: boolean;
  clear: boolean;
}

export interface ManyToManyOptions {
  jointable: string;
  foreignkey: string;
  jointableForeignkey: string;
  associationForeignkey: string;
  associationJointableForeignkey: string;
  disableAssociationAutoupdate: boolean;
  disableAssociationAutocreate: boolean;
  associationSaveReference: boolean;
  preload: boolean;
  replace: boolean;
  append: boolean;
  clear: boolean;
}

export interface AutoServerOptions {
  autogen: boolean;
  txnMiddleware: boolean;
  withTracing: boolean;
}

export interface MethodOptions {
  objectType: string;
}

function createBaseGormFileOptions(): GormFileOptions {
  return {};
}

export const GormFileOptions = {
  encode(_: GormFileOptions, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GormFileOptions {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGormFileOptions();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): GormFileOptions {
    return {};
  },

  toJSON(_: GormFileOptions): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<GormFileOptions>, I>>(base?: I): GormFileOptions {
    return GormFileOptions.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GormFileOptions>, I>>(_: I): GormFileOptions {
    const message = createBaseGormFileOptions();
    return message;
  },
};

function createBaseGormMessageOptions(): GormMessageOptions {
  return { ormable: false, include: [], table: "", multiAccount: false };
}

export const GormMessageOptions = {
  encode(message: GormMessageOptions, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ormable === true) {
      writer.uint32(8).bool(message.ormable);
    }
    for (const v of message.include) {
      ExtraField.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.table !== "") {
      writer.uint32(26).string(message.table);
    }
    if (message.multiAccount === true) {
      writer.uint32(32).bool(message.multiAccount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GormMessageOptions {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGormMessageOptions();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.ormable = reader.bool();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.include.push(ExtraField.decode(reader, reader.uint32()));
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.table = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.multiAccount = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GormMessageOptions {
    return {
      ormable: isSet(object.ormable) ? Boolean(object.ormable) : false,
      include: Array.isArray(object?.include) ? object.include.map((e: any) => ExtraField.fromJSON(e)) : [],
      table: isSet(object.table) ? String(object.table) : "",
      multiAccount: isSet(object.multiAccount) ? Boolean(object.multiAccount) : false,
    };
  },

  toJSON(message: GormMessageOptions): unknown {
    const obj: any = {};
    if (message.ormable === true) {
      obj.ormable = message.ormable;
    }
    if (message.include?.length) {
      obj.include = message.include.map((e) => ExtraField.toJSON(e));
    }
    if (message.table !== "") {
      obj.table = message.table;
    }
    if (message.multiAccount === true) {
      obj.multiAccount = message.multiAccount;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GormMessageOptions>, I>>(base?: I): GormMessageOptions {
    return GormMessageOptions.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GormMessageOptions>, I>>(object: I): GormMessageOptions {
    const message = createBaseGormMessageOptions();
    message.ormable = object.ormable ?? false;
    message.include = object.include?.map((e) => ExtraField.fromPartial(e)) || [];
    message.table = object.table ?? "";
    message.multiAccount = object.multiAccount ?? false;
    return message;
  },
};

function createBaseExtraField(): ExtraField {
  return { type: "", name: "", tag: undefined, package: "" };
}

export const ExtraField = {
  encode(message: ExtraField, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.type !== "") {
      writer.uint32(10).string(message.type);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.tag !== undefined) {
      GormTag.encode(message.tag, writer.uint32(26).fork()).ldelim();
    }
    if (message.package !== "") {
      writer.uint32(34).string(message.package);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ExtraField {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseExtraField();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.type = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.name = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.tag = GormTag.decode(reader, reader.uint32());
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.package = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ExtraField {
    return {
      type: isSet(object.type) ? String(object.type) : "",
      name: isSet(object.name) ? String(object.name) : "",
      tag: isSet(object.tag) ? GormTag.fromJSON(object.tag) : undefined,
      package: isSet(object.package) ? String(object.package) : "",
    };
  },

  toJSON(message: ExtraField): unknown {
    const obj: any = {};
    if (message.type !== "") {
      obj.type = message.type;
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.tag !== undefined) {
      obj.tag = GormTag.toJSON(message.tag);
    }
    if (message.package !== "") {
      obj.package = message.package;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ExtraField>, I>>(base?: I): ExtraField {
    return ExtraField.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ExtraField>, I>>(object: I): ExtraField {
    const message = createBaseExtraField();
    message.type = object.type ?? "";
    message.name = object.name ?? "";
    message.tag = (object.tag !== undefined && object.tag !== null) ? GormTag.fromPartial(object.tag) : undefined;
    message.package = object.package ?? "";
    return message;
  },
};

function createBaseGormFieldOptions(): GormFieldOptions {
  return {
    tag: undefined,
    drop: false,
    hasOne: undefined,
    belongsTo: undefined,
    hasMany: undefined,
    manyToMany: undefined,
    referenceOf: "",
  };
}

export const GormFieldOptions = {
  encode(message: GormFieldOptions, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.tag !== undefined) {
      GormTag.encode(message.tag, writer.uint32(10).fork()).ldelim();
    }
    if (message.drop === true) {
      writer.uint32(16).bool(message.drop);
    }
    if (message.hasOne !== undefined) {
      HasOneOptions.encode(message.hasOne, writer.uint32(26).fork()).ldelim();
    }
    if (message.belongsTo !== undefined) {
      BelongsToOptions.encode(message.belongsTo, writer.uint32(34).fork()).ldelim();
    }
    if (message.hasMany !== undefined) {
      HasManyOptions.encode(message.hasMany, writer.uint32(42).fork()).ldelim();
    }
    if (message.manyToMany !== undefined) {
      ManyToManyOptions.encode(message.manyToMany, writer.uint32(50).fork()).ldelim();
    }
    if (message.referenceOf !== "") {
      writer.uint32(58).string(message.referenceOf);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GormFieldOptions {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGormFieldOptions();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.tag = GormTag.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.drop = reader.bool();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.hasOne = HasOneOptions.decode(reader, reader.uint32());
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.belongsTo = BelongsToOptions.decode(reader, reader.uint32());
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.hasMany = HasManyOptions.decode(reader, reader.uint32());
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.manyToMany = ManyToManyOptions.decode(reader, reader.uint32());
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.referenceOf = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GormFieldOptions {
    return {
      tag: isSet(object.tag) ? GormTag.fromJSON(object.tag) : undefined,
      drop: isSet(object.drop) ? Boolean(object.drop) : false,
      hasOne: isSet(object.hasOne) ? HasOneOptions.fromJSON(object.hasOne) : undefined,
      belongsTo: isSet(object.belongsTo) ? BelongsToOptions.fromJSON(object.belongsTo) : undefined,
      hasMany: isSet(object.hasMany) ? HasManyOptions.fromJSON(object.hasMany) : undefined,
      manyToMany: isSet(object.manyToMany) ? ManyToManyOptions.fromJSON(object.manyToMany) : undefined,
      referenceOf: isSet(object.referenceOf) ? String(object.referenceOf) : "",
    };
  },

  toJSON(message: GormFieldOptions): unknown {
    const obj: any = {};
    if (message.tag !== undefined) {
      obj.tag = GormTag.toJSON(message.tag);
    }
    if (message.drop === true) {
      obj.drop = message.drop;
    }
    if (message.hasOne !== undefined) {
      obj.hasOne = HasOneOptions.toJSON(message.hasOne);
    }
    if (message.belongsTo !== undefined) {
      obj.belongsTo = BelongsToOptions.toJSON(message.belongsTo);
    }
    if (message.hasMany !== undefined) {
      obj.hasMany = HasManyOptions.toJSON(message.hasMany);
    }
    if (message.manyToMany !== undefined) {
      obj.manyToMany = ManyToManyOptions.toJSON(message.manyToMany);
    }
    if (message.referenceOf !== "") {
      obj.referenceOf = message.referenceOf;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GormFieldOptions>, I>>(base?: I): GormFieldOptions {
    return GormFieldOptions.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GormFieldOptions>, I>>(object: I): GormFieldOptions {
    const message = createBaseGormFieldOptions();
    message.tag = (object.tag !== undefined && object.tag !== null) ? GormTag.fromPartial(object.tag) : undefined;
    message.drop = object.drop ?? false;
    message.hasOne = (object.hasOne !== undefined && object.hasOne !== null)
      ? HasOneOptions.fromPartial(object.hasOne)
      : undefined;
    message.belongsTo = (object.belongsTo !== undefined && object.belongsTo !== null)
      ? BelongsToOptions.fromPartial(object.belongsTo)
      : undefined;
    message.hasMany = (object.hasMany !== undefined && object.hasMany !== null)
      ? HasManyOptions.fromPartial(object.hasMany)
      : undefined;
    message.manyToMany = (object.manyToMany !== undefined && object.manyToMany !== null)
      ? ManyToManyOptions.fromPartial(object.manyToMany)
      : undefined;
    message.referenceOf = object.referenceOf ?? "";
    return message;
  },
};

function createBaseGormTag(): GormTag {
  return {
    column: "",
    type: "",
    size: 0,
    precision: 0,
    primaryKey: false,
    unique: false,
    default: "",
    notNull: false,
    autoIncrement: false,
    index: "",
    uniqueIndex: "",
    embedded: false,
    embeddedPrefix: "",
    ignore: false,
    foreignkey: "",
    associationForeignkey: "",
    manyToMany: "",
    jointableForeignkey: "",
    associationJointableForeignkey: "",
    disableAssociationAutoupdate: false,
    disableAssociationAutocreate: false,
    associationSaveReference: false,
    preload: false,
    serializer: "",
  };
}

export const GormTag = {
  encode(message: GormTag, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.column !== "") {
      writer.uint32(10).string(message.column);
    }
    if (message.type !== "") {
      writer.uint32(18).string(message.type);
    }
    if (message.size !== 0) {
      writer.uint32(24).int32(message.size);
    }
    if (message.precision !== 0) {
      writer.uint32(32).int32(message.precision);
    }
    if (message.primaryKey === true) {
      writer.uint32(40).bool(message.primaryKey);
    }
    if (message.unique === true) {
      writer.uint32(48).bool(message.unique);
    }
    if (message.default !== "") {
      writer.uint32(58).string(message.default);
    }
    if (message.notNull === true) {
      writer.uint32(64).bool(message.notNull);
    }
    if (message.autoIncrement === true) {
      writer.uint32(72).bool(message.autoIncrement);
    }
    if (message.index !== "") {
      writer.uint32(82).string(message.index);
    }
    if (message.uniqueIndex !== "") {
      writer.uint32(90).string(message.uniqueIndex);
    }
    if (message.embedded === true) {
      writer.uint32(96).bool(message.embedded);
    }
    if (message.embeddedPrefix !== "") {
      writer.uint32(106).string(message.embeddedPrefix);
    }
    if (message.ignore === true) {
      writer.uint32(112).bool(message.ignore);
    }
    if (message.foreignkey !== "") {
      writer.uint32(122).string(message.foreignkey);
    }
    if (message.associationForeignkey !== "") {
      writer.uint32(130).string(message.associationForeignkey);
    }
    if (message.manyToMany !== "") {
      writer.uint32(138).string(message.manyToMany);
    }
    if (message.jointableForeignkey !== "") {
      writer.uint32(146).string(message.jointableForeignkey);
    }
    if (message.associationJointableForeignkey !== "") {
      writer.uint32(154).string(message.associationJointableForeignkey);
    }
    if (message.disableAssociationAutoupdate === true) {
      writer.uint32(160).bool(message.disableAssociationAutoupdate);
    }
    if (message.disableAssociationAutocreate === true) {
      writer.uint32(168).bool(message.disableAssociationAutocreate);
    }
    if (message.associationSaveReference === true) {
      writer.uint32(176).bool(message.associationSaveReference);
    }
    if (message.preload === true) {
      writer.uint32(184).bool(message.preload);
    }
    if (message.serializer !== "") {
      writer.uint32(194).string(message.serializer);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GormTag {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGormTag();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.column = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.type = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.size = reader.int32();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.precision = reader.int32();
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.primaryKey = reader.bool();
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.unique = reader.bool();
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.default = reader.string();
          continue;
        case 8:
          if (tag !== 64) {
            break;
          }

          message.notNull = reader.bool();
          continue;
        case 9:
          if (tag !== 72) {
            break;
          }

          message.autoIncrement = reader.bool();
          continue;
        case 10:
          if (tag !== 82) {
            break;
          }

          message.index = reader.string();
          continue;
        case 11:
          if (tag !== 90) {
            break;
          }

          message.uniqueIndex = reader.string();
          continue;
        case 12:
          if (tag !== 96) {
            break;
          }

          message.embedded = reader.bool();
          continue;
        case 13:
          if (tag !== 106) {
            break;
          }

          message.embeddedPrefix = reader.string();
          continue;
        case 14:
          if (tag !== 112) {
            break;
          }

          message.ignore = reader.bool();
          continue;
        case 15:
          if (tag !== 122) {
            break;
          }

          message.foreignkey = reader.string();
          continue;
        case 16:
          if (tag !== 130) {
            break;
          }

          message.associationForeignkey = reader.string();
          continue;
        case 17:
          if (tag !== 138) {
            break;
          }

          message.manyToMany = reader.string();
          continue;
        case 18:
          if (tag !== 146) {
            break;
          }

          message.jointableForeignkey = reader.string();
          continue;
        case 19:
          if (tag !== 154) {
            break;
          }

          message.associationJointableForeignkey = reader.string();
          continue;
        case 20:
          if (tag !== 160) {
            break;
          }

          message.disableAssociationAutoupdate = reader.bool();
          continue;
        case 21:
          if (tag !== 168) {
            break;
          }

          message.disableAssociationAutocreate = reader.bool();
          continue;
        case 22:
          if (tag !== 176) {
            break;
          }

          message.associationSaveReference = reader.bool();
          continue;
        case 23:
          if (tag !== 184) {
            break;
          }

          message.preload = reader.bool();
          continue;
        case 24:
          if (tag !== 194) {
            break;
          }

          message.serializer = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GormTag {
    return {
      column: isSet(object.column) ? String(object.column) : "",
      type: isSet(object.type) ? String(object.type) : "",
      size: isSet(object.size) ? Number(object.size) : 0,
      precision: isSet(object.precision) ? Number(object.precision) : 0,
      primaryKey: isSet(object.primaryKey) ? Boolean(object.primaryKey) : false,
      unique: isSet(object.unique) ? Boolean(object.unique) : false,
      default: isSet(object.default) ? String(object.default) : "",
      notNull: isSet(object.notNull) ? Boolean(object.notNull) : false,
      autoIncrement: isSet(object.autoIncrement) ? Boolean(object.autoIncrement) : false,
      index: isSet(object.index) ? String(object.index) : "",
      uniqueIndex: isSet(object.uniqueIndex) ? String(object.uniqueIndex) : "",
      embedded: isSet(object.embedded) ? Boolean(object.embedded) : false,
      embeddedPrefix: isSet(object.embeddedPrefix) ? String(object.embeddedPrefix) : "",
      ignore: isSet(object.ignore) ? Boolean(object.ignore) : false,
      foreignkey: isSet(object.foreignkey) ? String(object.foreignkey) : "",
      associationForeignkey: isSet(object.associationForeignkey) ? String(object.associationForeignkey) : "",
      manyToMany: isSet(object.manyToMany) ? String(object.manyToMany) : "",
      jointableForeignkey: isSet(object.jointableForeignkey) ? String(object.jointableForeignkey) : "",
      associationJointableForeignkey: isSet(object.associationJointableForeignkey)
        ? String(object.associationJointableForeignkey)
        : "",
      disableAssociationAutoupdate: isSet(object.disableAssociationAutoupdate)
        ? Boolean(object.disableAssociationAutoupdate)
        : false,
      disableAssociationAutocreate: isSet(object.disableAssociationAutocreate)
        ? Boolean(object.disableAssociationAutocreate)
        : false,
      associationSaveReference: isSet(object.associationSaveReference)
        ? Boolean(object.associationSaveReference)
        : false,
      preload: isSet(object.preload) ? Boolean(object.preload) : false,
      serializer: isSet(object.serializer) ? String(object.serializer) : "",
    };
  },

  toJSON(message: GormTag): unknown {
    const obj: any = {};
    if (message.column !== "") {
      obj.column = message.column;
    }
    if (message.type !== "") {
      obj.type = message.type;
    }
    if (message.size !== 0) {
      obj.size = Math.round(message.size);
    }
    if (message.precision !== 0) {
      obj.precision = Math.round(message.precision);
    }
    if (message.primaryKey === true) {
      obj.primaryKey = message.primaryKey;
    }
    if (message.unique === true) {
      obj.unique = message.unique;
    }
    if (message.default !== "") {
      obj.default = message.default;
    }
    if (message.notNull === true) {
      obj.notNull = message.notNull;
    }
    if (message.autoIncrement === true) {
      obj.autoIncrement = message.autoIncrement;
    }
    if (message.index !== "") {
      obj.index = message.index;
    }
    if (message.uniqueIndex !== "") {
      obj.uniqueIndex = message.uniqueIndex;
    }
    if (message.embedded === true) {
      obj.embedded = message.embedded;
    }
    if (message.embeddedPrefix !== "") {
      obj.embeddedPrefix = message.embeddedPrefix;
    }
    if (message.ignore === true) {
      obj.ignore = message.ignore;
    }
    if (message.foreignkey !== "") {
      obj.foreignkey = message.foreignkey;
    }
    if (message.associationForeignkey !== "") {
      obj.associationForeignkey = message.associationForeignkey;
    }
    if (message.manyToMany !== "") {
      obj.manyToMany = message.manyToMany;
    }
    if (message.jointableForeignkey !== "") {
      obj.jointableForeignkey = message.jointableForeignkey;
    }
    if (message.associationJointableForeignkey !== "") {
      obj.associationJointableForeignkey = message.associationJointableForeignkey;
    }
    if (message.disableAssociationAutoupdate === true) {
      obj.disableAssociationAutoupdate = message.disableAssociationAutoupdate;
    }
    if (message.disableAssociationAutocreate === true) {
      obj.disableAssociationAutocreate = message.disableAssociationAutocreate;
    }
    if (message.associationSaveReference === true) {
      obj.associationSaveReference = message.associationSaveReference;
    }
    if (message.preload === true) {
      obj.preload = message.preload;
    }
    if (message.serializer !== "") {
      obj.serializer = message.serializer;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GormTag>, I>>(base?: I): GormTag {
    return GormTag.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GormTag>, I>>(object: I): GormTag {
    const message = createBaseGormTag();
    message.column = object.column ?? "";
    message.type = object.type ?? "";
    message.size = object.size ?? 0;
    message.precision = object.precision ?? 0;
    message.primaryKey = object.primaryKey ?? false;
    message.unique = object.unique ?? false;
    message.default = object.default ?? "";
    message.notNull = object.notNull ?? false;
    message.autoIncrement = object.autoIncrement ?? false;
    message.index = object.index ?? "";
    message.uniqueIndex = object.uniqueIndex ?? "";
    message.embedded = object.embedded ?? false;
    message.embeddedPrefix = object.embeddedPrefix ?? "";
    message.ignore = object.ignore ?? false;
    message.foreignkey = object.foreignkey ?? "";
    message.associationForeignkey = object.associationForeignkey ?? "";
    message.manyToMany = object.manyToMany ?? "";
    message.jointableForeignkey = object.jointableForeignkey ?? "";
    message.associationJointableForeignkey = object.associationJointableForeignkey ?? "";
    message.disableAssociationAutoupdate = object.disableAssociationAutoupdate ?? false;
    message.disableAssociationAutocreate = object.disableAssociationAutocreate ?? false;
    message.associationSaveReference = object.associationSaveReference ?? false;
    message.preload = object.preload ?? false;
    message.serializer = object.serializer ?? "";
    return message;
  },
};

function createBaseHasOneOptions(): HasOneOptions {
  return {
    foreignkey: "",
    foreignkeyTag: undefined,
    associationForeignkey: "",
    disableAssociationAutoupdate: false,
    disableAssociationAutocreate: false,
    associationSaveReference: false,
    preload: false,
    replace: false,
    append: false,
    clear: false,
  };
}

export const HasOneOptions = {
  encode(message: HasOneOptions, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.foreignkey !== "") {
      writer.uint32(10).string(message.foreignkey);
    }
    if (message.foreignkeyTag !== undefined) {
      GormTag.encode(message.foreignkeyTag, writer.uint32(18).fork()).ldelim();
    }
    if (message.associationForeignkey !== "") {
      writer.uint32(26).string(message.associationForeignkey);
    }
    if (message.disableAssociationAutoupdate === true) {
      writer.uint32(32).bool(message.disableAssociationAutoupdate);
    }
    if (message.disableAssociationAutocreate === true) {
      writer.uint32(40).bool(message.disableAssociationAutocreate);
    }
    if (message.associationSaveReference === true) {
      writer.uint32(48).bool(message.associationSaveReference);
    }
    if (message.preload === true) {
      writer.uint32(56).bool(message.preload);
    }
    if (message.replace === true) {
      writer.uint32(64).bool(message.replace);
    }
    if (message.append === true) {
      writer.uint32(72).bool(message.append);
    }
    if (message.clear === true) {
      writer.uint32(80).bool(message.clear);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): HasOneOptions {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseHasOneOptions();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.foreignkey = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.foreignkeyTag = GormTag.decode(reader, reader.uint32());
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.associationForeignkey = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.disableAssociationAutoupdate = reader.bool();
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.disableAssociationAutocreate = reader.bool();
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.associationSaveReference = reader.bool();
          continue;
        case 7:
          if (tag !== 56) {
            break;
          }

          message.preload = reader.bool();
          continue;
        case 8:
          if (tag !== 64) {
            break;
          }

          message.replace = reader.bool();
          continue;
        case 9:
          if (tag !== 72) {
            break;
          }

          message.append = reader.bool();
          continue;
        case 10:
          if (tag !== 80) {
            break;
          }

          message.clear = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): HasOneOptions {
    return {
      foreignkey: isSet(object.foreignkey) ? String(object.foreignkey) : "",
      foreignkeyTag: isSet(object.foreignkeyTag) ? GormTag.fromJSON(object.foreignkeyTag) : undefined,
      associationForeignkey: isSet(object.associationForeignkey) ? String(object.associationForeignkey) : "",
      disableAssociationAutoupdate: isSet(object.disableAssociationAutoupdate)
        ? Boolean(object.disableAssociationAutoupdate)
        : false,
      disableAssociationAutocreate: isSet(object.disableAssociationAutocreate)
        ? Boolean(object.disableAssociationAutocreate)
        : false,
      associationSaveReference: isSet(object.associationSaveReference)
        ? Boolean(object.associationSaveReference)
        : false,
      preload: isSet(object.preload) ? Boolean(object.preload) : false,
      replace: isSet(object.replace) ? Boolean(object.replace) : false,
      append: isSet(object.append) ? Boolean(object.append) : false,
      clear: isSet(object.clear) ? Boolean(object.clear) : false,
    };
  },

  toJSON(message: HasOneOptions): unknown {
    const obj: any = {};
    if (message.foreignkey !== "") {
      obj.foreignkey = message.foreignkey;
    }
    if (message.foreignkeyTag !== undefined) {
      obj.foreignkeyTag = GormTag.toJSON(message.foreignkeyTag);
    }
    if (message.associationForeignkey !== "") {
      obj.associationForeignkey = message.associationForeignkey;
    }
    if (message.disableAssociationAutoupdate === true) {
      obj.disableAssociationAutoupdate = message.disableAssociationAutoupdate;
    }
    if (message.disableAssociationAutocreate === true) {
      obj.disableAssociationAutocreate = message.disableAssociationAutocreate;
    }
    if (message.associationSaveReference === true) {
      obj.associationSaveReference = message.associationSaveReference;
    }
    if (message.preload === true) {
      obj.preload = message.preload;
    }
    if (message.replace === true) {
      obj.replace = message.replace;
    }
    if (message.append === true) {
      obj.append = message.append;
    }
    if (message.clear === true) {
      obj.clear = message.clear;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<HasOneOptions>, I>>(base?: I): HasOneOptions {
    return HasOneOptions.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<HasOneOptions>, I>>(object: I): HasOneOptions {
    const message = createBaseHasOneOptions();
    message.foreignkey = object.foreignkey ?? "";
    message.foreignkeyTag = (object.foreignkeyTag !== undefined && object.foreignkeyTag !== null)
      ? GormTag.fromPartial(object.foreignkeyTag)
      : undefined;
    message.associationForeignkey = object.associationForeignkey ?? "";
    message.disableAssociationAutoupdate = object.disableAssociationAutoupdate ?? false;
    message.disableAssociationAutocreate = object.disableAssociationAutocreate ?? false;
    message.associationSaveReference = object.associationSaveReference ?? false;
    message.preload = object.preload ?? false;
    message.replace = object.replace ?? false;
    message.append = object.append ?? false;
    message.clear = object.clear ?? false;
    return message;
  },
};

function createBaseBelongsToOptions(): BelongsToOptions {
  return {
    foreignkey: "",
    foreignkeyTag: undefined,
    associationForeignkey: "",
    disableAssociationAutoupdate: false,
    disableAssociationAutocreate: false,
    associationSaveReference: false,
    preload: false,
  };
}

export const BelongsToOptions = {
  encode(message: BelongsToOptions, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.foreignkey !== "") {
      writer.uint32(10).string(message.foreignkey);
    }
    if (message.foreignkeyTag !== undefined) {
      GormTag.encode(message.foreignkeyTag, writer.uint32(18).fork()).ldelim();
    }
    if (message.associationForeignkey !== "") {
      writer.uint32(26).string(message.associationForeignkey);
    }
    if (message.disableAssociationAutoupdate === true) {
      writer.uint32(32).bool(message.disableAssociationAutoupdate);
    }
    if (message.disableAssociationAutocreate === true) {
      writer.uint32(40).bool(message.disableAssociationAutocreate);
    }
    if (message.associationSaveReference === true) {
      writer.uint32(48).bool(message.associationSaveReference);
    }
    if (message.preload === true) {
      writer.uint32(56).bool(message.preload);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BelongsToOptions {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBelongsToOptions();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.foreignkey = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.foreignkeyTag = GormTag.decode(reader, reader.uint32());
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.associationForeignkey = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.disableAssociationAutoupdate = reader.bool();
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.disableAssociationAutocreate = reader.bool();
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.associationSaveReference = reader.bool();
          continue;
        case 7:
          if (tag !== 56) {
            break;
          }

          message.preload = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): BelongsToOptions {
    return {
      foreignkey: isSet(object.foreignkey) ? String(object.foreignkey) : "",
      foreignkeyTag: isSet(object.foreignkeyTag) ? GormTag.fromJSON(object.foreignkeyTag) : undefined,
      associationForeignkey: isSet(object.associationForeignkey) ? String(object.associationForeignkey) : "",
      disableAssociationAutoupdate: isSet(object.disableAssociationAutoupdate)
        ? Boolean(object.disableAssociationAutoupdate)
        : false,
      disableAssociationAutocreate: isSet(object.disableAssociationAutocreate)
        ? Boolean(object.disableAssociationAutocreate)
        : false,
      associationSaveReference: isSet(object.associationSaveReference)
        ? Boolean(object.associationSaveReference)
        : false,
      preload: isSet(object.preload) ? Boolean(object.preload) : false,
    };
  },

  toJSON(message: BelongsToOptions): unknown {
    const obj: any = {};
    if (message.foreignkey !== "") {
      obj.foreignkey = message.foreignkey;
    }
    if (message.foreignkeyTag !== undefined) {
      obj.foreignkeyTag = GormTag.toJSON(message.foreignkeyTag);
    }
    if (message.associationForeignkey !== "") {
      obj.associationForeignkey = message.associationForeignkey;
    }
    if (message.disableAssociationAutoupdate === true) {
      obj.disableAssociationAutoupdate = message.disableAssociationAutoupdate;
    }
    if (message.disableAssociationAutocreate === true) {
      obj.disableAssociationAutocreate = message.disableAssociationAutocreate;
    }
    if (message.associationSaveReference === true) {
      obj.associationSaveReference = message.associationSaveReference;
    }
    if (message.preload === true) {
      obj.preload = message.preload;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<BelongsToOptions>, I>>(base?: I): BelongsToOptions {
    return BelongsToOptions.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<BelongsToOptions>, I>>(object: I): BelongsToOptions {
    const message = createBaseBelongsToOptions();
    message.foreignkey = object.foreignkey ?? "";
    message.foreignkeyTag = (object.foreignkeyTag !== undefined && object.foreignkeyTag !== null)
      ? GormTag.fromPartial(object.foreignkeyTag)
      : undefined;
    message.associationForeignkey = object.associationForeignkey ?? "";
    message.disableAssociationAutoupdate = object.disableAssociationAutoupdate ?? false;
    message.disableAssociationAutocreate = object.disableAssociationAutocreate ?? false;
    message.associationSaveReference = object.associationSaveReference ?? false;
    message.preload = object.preload ?? false;
    return message;
  },
};

function createBaseHasManyOptions(): HasManyOptions {
  return {
    foreignkey: "",
    foreignkeyTag: undefined,
    associationForeignkey: "",
    positionField: "",
    positionFieldTag: undefined,
    disableAssociationAutoupdate: false,
    disableAssociationAutocreate: false,
    associationSaveReference: false,
    preload: false,
    replace: false,
    append: false,
    clear: false,
  };
}

export const HasManyOptions = {
  encode(message: HasManyOptions, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.foreignkey !== "") {
      writer.uint32(10).string(message.foreignkey);
    }
    if (message.foreignkeyTag !== undefined) {
      GormTag.encode(message.foreignkeyTag, writer.uint32(18).fork()).ldelim();
    }
    if (message.associationForeignkey !== "") {
      writer.uint32(26).string(message.associationForeignkey);
    }
    if (message.positionField !== "") {
      writer.uint32(34).string(message.positionField);
    }
    if (message.positionFieldTag !== undefined) {
      GormTag.encode(message.positionFieldTag, writer.uint32(42).fork()).ldelim();
    }
    if (message.disableAssociationAutoupdate === true) {
      writer.uint32(48).bool(message.disableAssociationAutoupdate);
    }
    if (message.disableAssociationAutocreate === true) {
      writer.uint32(56).bool(message.disableAssociationAutocreate);
    }
    if (message.associationSaveReference === true) {
      writer.uint32(64).bool(message.associationSaveReference);
    }
    if (message.preload === true) {
      writer.uint32(72).bool(message.preload);
    }
    if (message.replace === true) {
      writer.uint32(80).bool(message.replace);
    }
    if (message.append === true) {
      writer.uint32(88).bool(message.append);
    }
    if (message.clear === true) {
      writer.uint32(96).bool(message.clear);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): HasManyOptions {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseHasManyOptions();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.foreignkey = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.foreignkeyTag = GormTag.decode(reader, reader.uint32());
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.associationForeignkey = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.positionField = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.positionFieldTag = GormTag.decode(reader, reader.uint32());
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.disableAssociationAutoupdate = reader.bool();
          continue;
        case 7:
          if (tag !== 56) {
            break;
          }

          message.disableAssociationAutocreate = reader.bool();
          continue;
        case 8:
          if (tag !== 64) {
            break;
          }

          message.associationSaveReference = reader.bool();
          continue;
        case 9:
          if (tag !== 72) {
            break;
          }

          message.preload = reader.bool();
          continue;
        case 10:
          if (tag !== 80) {
            break;
          }

          message.replace = reader.bool();
          continue;
        case 11:
          if (tag !== 88) {
            break;
          }

          message.append = reader.bool();
          continue;
        case 12:
          if (tag !== 96) {
            break;
          }

          message.clear = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): HasManyOptions {
    return {
      foreignkey: isSet(object.foreignkey) ? String(object.foreignkey) : "",
      foreignkeyTag: isSet(object.foreignkeyTag) ? GormTag.fromJSON(object.foreignkeyTag) : undefined,
      associationForeignkey: isSet(object.associationForeignkey) ? String(object.associationForeignkey) : "",
      positionField: isSet(object.positionField) ? String(object.positionField) : "",
      positionFieldTag: isSet(object.positionFieldTag) ? GormTag.fromJSON(object.positionFieldTag) : undefined,
      disableAssociationAutoupdate: isSet(object.disableAssociationAutoupdate)
        ? Boolean(object.disableAssociationAutoupdate)
        : false,
      disableAssociationAutocreate: isSet(object.disableAssociationAutocreate)
        ? Boolean(object.disableAssociationAutocreate)
        : false,
      associationSaveReference: isSet(object.associationSaveReference)
        ? Boolean(object.associationSaveReference)
        : false,
      preload: isSet(object.preload) ? Boolean(object.preload) : false,
      replace: isSet(object.replace) ? Boolean(object.replace) : false,
      append: isSet(object.append) ? Boolean(object.append) : false,
      clear: isSet(object.clear) ? Boolean(object.clear) : false,
    };
  },

  toJSON(message: HasManyOptions): unknown {
    const obj: any = {};
    if (message.foreignkey !== "") {
      obj.foreignkey = message.foreignkey;
    }
    if (message.foreignkeyTag !== undefined) {
      obj.foreignkeyTag = GormTag.toJSON(message.foreignkeyTag);
    }
    if (message.associationForeignkey !== "") {
      obj.associationForeignkey = message.associationForeignkey;
    }
    if (message.positionField !== "") {
      obj.positionField = message.positionField;
    }
    if (message.positionFieldTag !== undefined) {
      obj.positionFieldTag = GormTag.toJSON(message.positionFieldTag);
    }
    if (message.disableAssociationAutoupdate === true) {
      obj.disableAssociationAutoupdate = message.disableAssociationAutoupdate;
    }
    if (message.disableAssociationAutocreate === true) {
      obj.disableAssociationAutocreate = message.disableAssociationAutocreate;
    }
    if (message.associationSaveReference === true) {
      obj.associationSaveReference = message.associationSaveReference;
    }
    if (message.preload === true) {
      obj.preload = message.preload;
    }
    if (message.replace === true) {
      obj.replace = message.replace;
    }
    if (message.append === true) {
      obj.append = message.append;
    }
    if (message.clear === true) {
      obj.clear = message.clear;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<HasManyOptions>, I>>(base?: I): HasManyOptions {
    return HasManyOptions.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<HasManyOptions>, I>>(object: I): HasManyOptions {
    const message = createBaseHasManyOptions();
    message.foreignkey = object.foreignkey ?? "";
    message.foreignkeyTag = (object.foreignkeyTag !== undefined && object.foreignkeyTag !== null)
      ? GormTag.fromPartial(object.foreignkeyTag)
      : undefined;
    message.associationForeignkey = object.associationForeignkey ?? "";
    message.positionField = object.positionField ?? "";
    message.positionFieldTag = (object.positionFieldTag !== undefined && object.positionFieldTag !== null)
      ? GormTag.fromPartial(object.positionFieldTag)
      : undefined;
    message.disableAssociationAutoupdate = object.disableAssociationAutoupdate ?? false;
    message.disableAssociationAutocreate = object.disableAssociationAutocreate ?? false;
    message.associationSaveReference = object.associationSaveReference ?? false;
    message.preload = object.preload ?? false;
    message.replace = object.replace ?? false;
    message.append = object.append ?? false;
    message.clear = object.clear ?? false;
    return message;
  },
};

function createBaseManyToManyOptions(): ManyToManyOptions {
  return {
    jointable: "",
    foreignkey: "",
    jointableForeignkey: "",
    associationForeignkey: "",
    associationJointableForeignkey: "",
    disableAssociationAutoupdate: false,
    disableAssociationAutocreate: false,
    associationSaveReference: false,
    preload: false,
    replace: false,
    append: false,
    clear: false,
  };
}

export const ManyToManyOptions = {
  encode(message: ManyToManyOptions, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.jointable !== "") {
      writer.uint32(10).string(message.jointable);
    }
    if (message.foreignkey !== "") {
      writer.uint32(18).string(message.foreignkey);
    }
    if (message.jointableForeignkey !== "") {
      writer.uint32(26).string(message.jointableForeignkey);
    }
    if (message.associationForeignkey !== "") {
      writer.uint32(34).string(message.associationForeignkey);
    }
    if (message.associationJointableForeignkey !== "") {
      writer.uint32(42).string(message.associationJointableForeignkey);
    }
    if (message.disableAssociationAutoupdate === true) {
      writer.uint32(48).bool(message.disableAssociationAutoupdate);
    }
    if (message.disableAssociationAutocreate === true) {
      writer.uint32(56).bool(message.disableAssociationAutocreate);
    }
    if (message.associationSaveReference === true) {
      writer.uint32(64).bool(message.associationSaveReference);
    }
    if (message.preload === true) {
      writer.uint32(72).bool(message.preload);
    }
    if (message.replace === true) {
      writer.uint32(80).bool(message.replace);
    }
    if (message.append === true) {
      writer.uint32(88).bool(message.append);
    }
    if (message.clear === true) {
      writer.uint32(104).bool(message.clear);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ManyToManyOptions {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseManyToManyOptions();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.jointable = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.foreignkey = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.jointableForeignkey = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.associationForeignkey = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.associationJointableForeignkey = reader.string();
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.disableAssociationAutoupdate = reader.bool();
          continue;
        case 7:
          if (tag !== 56) {
            break;
          }

          message.disableAssociationAutocreate = reader.bool();
          continue;
        case 8:
          if (tag !== 64) {
            break;
          }

          message.associationSaveReference = reader.bool();
          continue;
        case 9:
          if (tag !== 72) {
            break;
          }

          message.preload = reader.bool();
          continue;
        case 10:
          if (tag !== 80) {
            break;
          }

          message.replace = reader.bool();
          continue;
        case 11:
          if (tag !== 88) {
            break;
          }

          message.append = reader.bool();
          continue;
        case 13:
          if (tag !== 104) {
            break;
          }

          message.clear = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ManyToManyOptions {
    return {
      jointable: isSet(object.jointable) ? String(object.jointable) : "",
      foreignkey: isSet(object.foreignkey) ? String(object.foreignkey) : "",
      jointableForeignkey: isSet(object.jointableForeignkey) ? String(object.jointableForeignkey) : "",
      associationForeignkey: isSet(object.associationForeignkey) ? String(object.associationForeignkey) : "",
      associationJointableForeignkey: isSet(object.associationJointableForeignkey)
        ? String(object.associationJointableForeignkey)
        : "",
      disableAssociationAutoupdate: isSet(object.disableAssociationAutoupdate)
        ? Boolean(object.disableAssociationAutoupdate)
        : false,
      disableAssociationAutocreate: isSet(object.disableAssociationAutocreate)
        ? Boolean(object.disableAssociationAutocreate)
        : false,
      associationSaveReference: isSet(object.associationSaveReference)
        ? Boolean(object.associationSaveReference)
        : false,
      preload: isSet(object.preload) ? Boolean(object.preload) : false,
      replace: isSet(object.replace) ? Boolean(object.replace) : false,
      append: isSet(object.append) ? Boolean(object.append) : false,
      clear: isSet(object.clear) ? Boolean(object.clear) : false,
    };
  },

  toJSON(message: ManyToManyOptions): unknown {
    const obj: any = {};
    if (message.jointable !== "") {
      obj.jointable = message.jointable;
    }
    if (message.foreignkey !== "") {
      obj.foreignkey = message.foreignkey;
    }
    if (message.jointableForeignkey !== "") {
      obj.jointableForeignkey = message.jointableForeignkey;
    }
    if (message.associationForeignkey !== "") {
      obj.associationForeignkey = message.associationForeignkey;
    }
    if (message.associationJointableForeignkey !== "") {
      obj.associationJointableForeignkey = message.associationJointableForeignkey;
    }
    if (message.disableAssociationAutoupdate === true) {
      obj.disableAssociationAutoupdate = message.disableAssociationAutoupdate;
    }
    if (message.disableAssociationAutocreate === true) {
      obj.disableAssociationAutocreate = message.disableAssociationAutocreate;
    }
    if (message.associationSaveReference === true) {
      obj.associationSaveReference = message.associationSaveReference;
    }
    if (message.preload === true) {
      obj.preload = message.preload;
    }
    if (message.replace === true) {
      obj.replace = message.replace;
    }
    if (message.append === true) {
      obj.append = message.append;
    }
    if (message.clear === true) {
      obj.clear = message.clear;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ManyToManyOptions>, I>>(base?: I): ManyToManyOptions {
    return ManyToManyOptions.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ManyToManyOptions>, I>>(object: I): ManyToManyOptions {
    const message = createBaseManyToManyOptions();
    message.jointable = object.jointable ?? "";
    message.foreignkey = object.foreignkey ?? "";
    message.jointableForeignkey = object.jointableForeignkey ?? "";
    message.associationForeignkey = object.associationForeignkey ?? "";
    message.associationJointableForeignkey = object.associationJointableForeignkey ?? "";
    message.disableAssociationAutoupdate = object.disableAssociationAutoupdate ?? false;
    message.disableAssociationAutocreate = object.disableAssociationAutocreate ?? false;
    message.associationSaveReference = object.associationSaveReference ?? false;
    message.preload = object.preload ?? false;
    message.replace = object.replace ?? false;
    message.append = object.append ?? false;
    message.clear = object.clear ?? false;
    return message;
  },
};

function createBaseAutoServerOptions(): AutoServerOptions {
  return { autogen: false, txnMiddleware: false, withTracing: false };
}

export const AutoServerOptions = {
  encode(message: AutoServerOptions, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.autogen === true) {
      writer.uint32(8).bool(message.autogen);
    }
    if (message.txnMiddleware === true) {
      writer.uint32(16).bool(message.txnMiddleware);
    }
    if (message.withTracing === true) {
      writer.uint32(24).bool(message.withTracing);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AutoServerOptions {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAutoServerOptions();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.autogen = reader.bool();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.txnMiddleware = reader.bool();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.withTracing = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): AutoServerOptions {
    return {
      autogen: isSet(object.autogen) ? Boolean(object.autogen) : false,
      txnMiddleware: isSet(object.txnMiddleware) ? Boolean(object.txnMiddleware) : false,
      withTracing: isSet(object.withTracing) ? Boolean(object.withTracing) : false,
    };
  },

  toJSON(message: AutoServerOptions): unknown {
    const obj: any = {};
    if (message.autogen === true) {
      obj.autogen = message.autogen;
    }
    if (message.txnMiddleware === true) {
      obj.txnMiddleware = message.txnMiddleware;
    }
    if (message.withTracing === true) {
      obj.withTracing = message.withTracing;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<AutoServerOptions>, I>>(base?: I): AutoServerOptions {
    return AutoServerOptions.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<AutoServerOptions>, I>>(object: I): AutoServerOptions {
    const message = createBaseAutoServerOptions();
    message.autogen = object.autogen ?? false;
    message.txnMiddleware = object.txnMiddleware ?? false;
    message.withTracing = object.withTracing ?? false;
    return message;
  },
};

function createBaseMethodOptions(): MethodOptions {
  return { objectType: "" };
}

export const MethodOptions = {
  encode(message: MethodOptions, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.objectType !== "") {
      writer.uint32(10).string(message.objectType);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MethodOptions {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMethodOptions();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.objectType = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MethodOptions {
    return { objectType: isSet(object.objectType) ? String(object.objectType) : "" };
  },

  toJSON(message: MethodOptions): unknown {
    const obj: any = {};
    if (message.objectType !== "") {
      obj.objectType = message.objectType;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MethodOptions>, I>>(base?: I): MethodOptions {
    return MethodOptions.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MethodOptions>, I>>(object: I): MethodOptions {
    const message = createBaseMethodOptions();
    message.objectType = object.objectType ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
