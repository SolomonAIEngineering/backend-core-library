/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import Long = require("long");

export const protobufPackage = "message_definition.v1";

/**
 * HeadlessAuthenticationServiceDeleteAccountMessageFormat: represents an sqs message format
 * for deleting an account via the headless authentication service
 */
export interface HeadlessAuthenticationServiceDeleteAccountMessageFormat {
  /**
   * authn id which is the id of the account from the vantage point of the
   * authentication service
   */
  authnId: number;
  /**
   * account email
   * Validations:
   * - must be an email and required
   */
  email: string;
}

/**
 * FinancialIntegrationServiceMessageFormat: represents an sqs message format
 * for deleting an account via the financial integration service
 */
export interface FinancialIntegrationServiceMessageFormat {
  /** user_id id from the vantage point of the user service */
  userId: number;
}

/**
 * SocialServiceMessageFormat: represents an sqs message format
 * for deleting an account via the financial integration service
 */
export interface SocialServiceMessageFormat {
  /** user_id id from the vantage point of the social service */
  userId: number;
}

function createBaseHeadlessAuthenticationServiceDeleteAccountMessageFormat(): HeadlessAuthenticationServiceDeleteAccountMessageFormat {
  return { authnId: 0, email: "" };
}

export const HeadlessAuthenticationServiceDeleteAccountMessageFormat = {
  encode(
    message: HeadlessAuthenticationServiceDeleteAccountMessageFormat,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.authnId !== 0) {
      writer.uint32(8).uint64(message.authnId);
    }
    if (message.email !== "") {
      writer.uint32(18).string(message.email);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): HeadlessAuthenticationServiceDeleteAccountMessageFormat {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseHeadlessAuthenticationServiceDeleteAccountMessageFormat();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.authnId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.email = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): HeadlessAuthenticationServiceDeleteAccountMessageFormat {
    return {
      authnId: isSet(object.authnId) ? Number(object.authnId) : 0,
      email: isSet(object.email) ? String(object.email) : "",
    };
  },

  toJSON(message: HeadlessAuthenticationServiceDeleteAccountMessageFormat): unknown {
    const obj: any = {};
    if (message.authnId !== 0) {
      obj.authnId = Math.round(message.authnId);
    }
    if (message.email !== "") {
      obj.email = message.email;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<HeadlessAuthenticationServiceDeleteAccountMessageFormat>, I>>(
    base?: I,
  ): HeadlessAuthenticationServiceDeleteAccountMessageFormat {
    return HeadlessAuthenticationServiceDeleteAccountMessageFormat.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<HeadlessAuthenticationServiceDeleteAccountMessageFormat>, I>>(
    object: I,
  ): HeadlessAuthenticationServiceDeleteAccountMessageFormat {
    const message = createBaseHeadlessAuthenticationServiceDeleteAccountMessageFormat();
    message.authnId = object.authnId ?? 0;
    message.email = object.email ?? "";
    return message;
  },
};

function createBaseFinancialIntegrationServiceMessageFormat(): FinancialIntegrationServiceMessageFormat {
  return { userId: 0 };
}

export const FinancialIntegrationServiceMessageFormat = {
  encode(message: FinancialIntegrationServiceMessageFormat, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FinancialIntegrationServiceMessageFormat {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFinancialIntegrationServiceMessageFormat();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): FinancialIntegrationServiceMessageFormat {
    return { userId: isSet(object.userId) ? Number(object.userId) : 0 };
  },

  toJSON(message: FinancialIntegrationServiceMessageFormat): unknown {
    const obj: any = {};
    if (message.userId !== 0) {
      obj.userId = Math.round(message.userId);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<FinancialIntegrationServiceMessageFormat>, I>>(
    base?: I,
  ): FinancialIntegrationServiceMessageFormat {
    return FinancialIntegrationServiceMessageFormat.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<FinancialIntegrationServiceMessageFormat>, I>>(
    object: I,
  ): FinancialIntegrationServiceMessageFormat {
    const message = createBaseFinancialIntegrationServiceMessageFormat();
    message.userId = object.userId ?? 0;
    return message;
  },
};

function createBaseSocialServiceMessageFormat(): SocialServiceMessageFormat {
  return { userId: 0 };
}

export const SocialServiceMessageFormat = {
  encode(message: SocialServiceMessageFormat, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SocialServiceMessageFormat {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSocialServiceMessageFormat();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SocialServiceMessageFormat {
    return { userId: isSet(object.userId) ? Number(object.userId) : 0 };
  },

  toJSON(message: SocialServiceMessageFormat): unknown {
    const obj: any = {};
    if (message.userId !== 0) {
      obj.userId = Math.round(message.userId);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SocialServiceMessageFormat>, I>>(base?: I): SocialServiceMessageFormat {
    return SocialServiceMessageFormat.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<SocialServiceMessageFormat>, I>>(object: I): SocialServiceMessageFormat {
    const message = createBaseSocialServiceMessageFormat();
    message.userId = object.userId ?? 0;
    return message;
  },
};

declare const self: any | undefined;
declare const window: any | undefined;
declare const global: any | undefined;
const tsProtoGlobalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new tsProtoGlobalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
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
