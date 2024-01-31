/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import Long = require("long");

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

function createBaseDeleteAccountMessageFormat(): DeleteAccountMessageFormat {
  return { authZeroId: "", email: "", userId: 0, profileType: 0 };
}

export const DeleteAccountMessageFormat = {
  encode(message: DeleteAccountMessageFormat, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteAccountMessageFormat {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteAccountMessageFormat();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.authZeroId = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.email = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.profileType = reader.int32() as any;
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
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

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(globalThis.Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
