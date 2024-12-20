// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.3.0
//   protoc               unknown
// source: message_definition/v1/message.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export const protobufPackage = "message_definition.v1";

/**
 * DeleteAccountMessageFormat: represents an sqs message format
 * for deleting an accoun
 */
export interface DeleteAccountMessageFormat {
  /** auth_zero_id which is the id indicating the user and is the source of truth across all backend services */
  authZeroId: string;
  /**
   * account email
   * Validations:
   * - must be an email and required
   */
  email: string;
  /** user_id id from the vantage point of the user service */
  userId: number;
  /** the profile type of the given account */
  profileType: DeleteAccountMessageFormat_ProfileType;
}

export enum DeleteAccountMessageFormat_ProfileType {
  PROFILE_TYPE_UNSPECIFIED = 0,
  PROFILE_TYPE_USER = 1,
  PROFILE_TYPE_BUSINESS = 2,
  UNRECOGNIZED = -1,
}

export function deleteAccountMessageFormat_ProfileTypeFromJSON(object: any): DeleteAccountMessageFormat_ProfileType {
  switch (object) {
    case 0:
    case "PROFILE_TYPE_UNSPECIFIED":
      return DeleteAccountMessageFormat_ProfileType.PROFILE_TYPE_UNSPECIFIED;
    case 1:
    case "PROFILE_TYPE_USER":
      return DeleteAccountMessageFormat_ProfileType.PROFILE_TYPE_USER;
    case 2:
    case "PROFILE_TYPE_BUSINESS":
      return DeleteAccountMessageFormat_ProfileType.PROFILE_TYPE_BUSINESS;
    case -1:
    case "UNRECOGNIZED":
    default:
      return DeleteAccountMessageFormat_ProfileType.UNRECOGNIZED;
  }
}

export function deleteAccountMessageFormat_ProfileTypeToJSON(object: DeleteAccountMessageFormat_ProfileType): string {
  switch (object) {
    case DeleteAccountMessageFormat_ProfileType.PROFILE_TYPE_UNSPECIFIED:
      return "PROFILE_TYPE_UNSPECIFIED";
    case DeleteAccountMessageFormat_ProfileType.PROFILE_TYPE_USER:
      return "PROFILE_TYPE_USER";
    case DeleteAccountMessageFormat_ProfileType.PROFILE_TYPE_BUSINESS:
      return "PROFILE_TYPE_BUSINESS";
    case DeleteAccountMessageFormat_ProfileType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface AlgoliaSearchRecordFormat {
  name: string;
  userId: string;
  profileId: number;
  communityId: number;
  tags: string[];
  createdAt: string;
  profileImageUrl: string;
}

function createBaseDeleteAccountMessageFormat(): DeleteAccountMessageFormat {
  return { authZeroId: "", email: "", userId: 0, profileType: 0 };
}

export const DeleteAccountMessageFormat: MessageFns<DeleteAccountMessageFormat> = {
  encode(message: DeleteAccountMessageFormat, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.authZeroId !== "") {
      writer.uint32(10).string(message.authZeroId);
    }
    if (message.email !== "") {
      writer.uint32(18).string(message.email);
    }
    if (message.userId !== 0) {
      writer.uint32(24).uint64(message.userId);
    }
    if (message.profileType !== 0) {
      writer.uint32(32).int32(message.profileType);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): DeleteAccountMessageFormat {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteAccountMessageFormat();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.authZeroId = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.email = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 24) {
            break;
          }

          message.userId = longToNumber(reader.uint64());
          continue;
        }
        case 4: {
          if (tag !== 32) {
            break;
          }

          message.profileType = reader.int32() as any;
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteAccountMessageFormat {
    return {
      authZeroId: isSet(object.authZeroId) ? globalThis.String(object.authZeroId) : "",
      email: isSet(object.email) ? globalThis.String(object.email) : "",
      userId: isSet(object.userId) ? globalThis.Number(object.userId) : 0,
      profileType: isSet(object.profileType) ? deleteAccountMessageFormat_ProfileTypeFromJSON(object.profileType) : 0,
    };
  },

  toJSON(message: DeleteAccountMessageFormat): unknown {
    const obj: any = {};
    if (message.authZeroId !== "") {
      obj.authZeroId = message.authZeroId;
    }
    if (message.email !== "") {
      obj.email = message.email;
    }
    if (message.userId !== 0) {
      obj.userId = Math.round(message.userId);
    }
    if (message.profileType !== 0) {
      obj.profileType = deleteAccountMessageFormat_ProfileTypeToJSON(message.profileType);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteAccountMessageFormat>, I>>(base?: I): DeleteAccountMessageFormat {
    return DeleteAccountMessageFormat.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DeleteAccountMessageFormat>, I>>(object: I): DeleteAccountMessageFormat {
    const message = createBaseDeleteAccountMessageFormat();
    message.authZeroId = object.authZeroId ?? "";
    message.email = object.email ?? "";
    message.userId = object.userId ?? 0;
    message.profileType = object.profileType ?? 0;
    return message;
  },
};

function createBaseAlgoliaSearchRecordFormat(): AlgoliaSearchRecordFormat {
  return { name: "", userId: "", profileId: 0, communityId: 0, tags: [], createdAt: "", profileImageUrl: "" };
}

export const AlgoliaSearchRecordFormat: MessageFns<AlgoliaSearchRecordFormat> = {
  encode(message: AlgoliaSearchRecordFormat, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.userId !== "") {
      writer.uint32(18).string(message.userId);
    }
    if (message.profileId !== 0) {
      writer.uint32(24).uint64(message.profileId);
    }
    if (message.communityId !== 0) {
      writer.uint32(32).uint64(message.communityId);
    }
    for (const v of message.tags) {
      writer.uint32(42).string(v!);
    }
    if (message.createdAt !== "") {
      writer.uint32(50).string(message.createdAt);
    }
    if (message.profileImageUrl !== "") {
      writer.uint32(58).string(message.profileImageUrl);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): AlgoliaSearchRecordFormat {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAlgoliaSearchRecordFormat();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.userId = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 24) {
            break;
          }

          message.profileId = longToNumber(reader.uint64());
          continue;
        }
        case 4: {
          if (tag !== 32) {
            break;
          }

          message.communityId = longToNumber(reader.uint64());
          continue;
        }
        case 5: {
          if (tag !== 42) {
            break;
          }

          message.tags.push(reader.string());
          continue;
        }
        case 6: {
          if (tag !== 50) {
            break;
          }

          message.createdAt = reader.string();
          continue;
        }
        case 7: {
          if (tag !== 58) {
            break;
          }

          message.profileImageUrl = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): AlgoliaSearchRecordFormat {
    return {
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      userId: isSet(object.userId) ? globalThis.String(object.userId) : "",
      profileId: isSet(object.profileId) ? globalThis.Number(object.profileId) : 0,
      communityId: isSet(object.communityId) ? globalThis.Number(object.communityId) : 0,
      tags: globalThis.Array.isArray(object?.tags) ? object.tags.map((e: any) => globalThis.String(e)) : [],
      createdAt: isSet(object.createdAt) ? globalThis.String(object.createdAt) : "",
      profileImageUrl: isSet(object.profileImageUrl) ? globalThis.String(object.profileImageUrl) : "",
    };
  },

  toJSON(message: AlgoliaSearchRecordFormat): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.userId !== "") {
      obj.userId = message.userId;
    }
    if (message.profileId !== 0) {
      obj.profileId = Math.round(message.profileId);
    }
    if (message.communityId !== 0) {
      obj.communityId = Math.round(message.communityId);
    }
    if (message.tags?.length) {
      obj.tags = message.tags;
    }
    if (message.createdAt !== "") {
      obj.createdAt = message.createdAt;
    }
    if (message.profileImageUrl !== "") {
      obj.profileImageUrl = message.profileImageUrl;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<AlgoliaSearchRecordFormat>, I>>(base?: I): AlgoliaSearchRecordFormat {
    return AlgoliaSearchRecordFormat.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<AlgoliaSearchRecordFormat>, I>>(object: I): AlgoliaSearchRecordFormat {
    const message = createBaseAlgoliaSearchRecordFormat();
    message.name = object.name ?? "";
    message.userId = object.userId ?? "";
    message.profileId = object.profileId ?? 0;
    message.communityId = object.communityId ?? 0;
    message.tags = object.tags?.map((e) => e) || [];
    message.createdAt = object.createdAt ?? "";
    message.profileImageUrl = object.profileImageUrl ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(int64: { toString(): string }): number {
  const num = globalThis.Number(int64.toString());
  if (num > globalThis.Number.MAX_SAFE_INTEGER) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  if (num < globalThis.Number.MIN_SAFE_INTEGER) {
    throw new globalThis.Error("Value is smaller than Number.MIN_SAFE_INTEGER");
  }
  return num;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export interface MessageFns<T> {
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  fromJSON(object: any): T;
  toJSON(message: T): unknown;
  create<I extends Exact<DeepPartial<T>, I>>(base?: I): T;
  fromPartial<I extends Exact<DeepPartial<T>, I>>(object: I): T;
}
