/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "message_definition.v1";

export enum AuthErrorCode {
  no_auth_error = 0,
  auth_failed_invalid_subject = 1001,
  auth_failed_invalid_audience = 1002,
  auth_failed_invalid_issuer = 1003,
  invalid_claims = 1004,
  auth_failed_invalid_bearer_token = 1005,
  bearer_token_missing = 1010,
  unauthenticated = 1500,
  UNRECOGNIZED = -1,
}

export function authErrorCodeFromJSON(object: any): AuthErrorCode {
  switch (object) {
    case 0:
    case "no_auth_error":
      return AuthErrorCode.no_auth_error;
    case 1001:
    case "auth_failed_invalid_subject":
      return AuthErrorCode.auth_failed_invalid_subject;
    case 1002:
    case "auth_failed_invalid_audience":
      return AuthErrorCode.auth_failed_invalid_audience;
    case 1003:
    case "auth_failed_invalid_issuer":
      return AuthErrorCode.auth_failed_invalid_issuer;
    case 1004:
    case "invalid_claims":
      return AuthErrorCode.invalid_claims;
    case 1005:
    case "auth_failed_invalid_bearer_token":
      return AuthErrorCode.auth_failed_invalid_bearer_token;
    case 1010:
    case "bearer_token_missing":
      return AuthErrorCode.bearer_token_missing;
    case 1500:
    case "unauthenticated":
      return AuthErrorCode.unauthenticated;
    case -1:
    case "UNRECOGNIZED":
    default:
      return AuthErrorCode.UNRECOGNIZED;
  }
}

export function authErrorCodeToJSON(object: AuthErrorCode): string {
  switch (object) {
    case AuthErrorCode.no_auth_error:
      return "no_auth_error";
    case AuthErrorCode.auth_failed_invalid_subject:
      return "auth_failed_invalid_subject";
    case AuthErrorCode.auth_failed_invalid_audience:
      return "auth_failed_invalid_audience";
    case AuthErrorCode.auth_failed_invalid_issuer:
      return "auth_failed_invalid_issuer";
    case AuthErrorCode.invalid_claims:
      return "invalid_claims";
    case AuthErrorCode.auth_failed_invalid_bearer_token:
      return "auth_failed_invalid_bearer_token";
    case AuthErrorCode.bearer_token_missing:
      return "bearer_token_missing";
    case AuthErrorCode.unauthenticated:
      return "unauthenticated";
    case AuthErrorCode.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum ErrorCode {
  no_error = 0,
  validation_error = 2000,
  authorization_model_not_found = 2001,
  authorization_model_resolution_too_complex = 2002,
  invalid_write_input = 2003,
  cannot_allow_duplicate_tuples_in_one_request = 2004,
  cannot_allow_duplicate_types_in_one_request = 2005,
  cannot_allow_multiple_references_to_one_relation = 2006,
  invalid_continuation_token = 2007,
  invalid_tuple_set = 2008,
  invalid_check_input = 2009,
  invalid_expand_input = 2010,
  unsupported_user_set = 2011,
  invalid_object_format = 2012,
  write_failed_due_to_invalid_input = 2017,
  authorization_model_assertions_not_found = 2018,
  latest_authorization_model_not_found = 2020,
  type_not_found = 2021,
  relation_not_found = 2022,
  empty_relation_definition = 2023,
  invalid_user = 2025,
  invalid_tuple = 2027,
  unknown_relation = 2028,
  store_id_invalid_length = 2030,
  assertions_too_many_items = 2033,
  id_too_long = 2034,
  authorization_model_id_too_long = 2036,
  tuple_key_value_not_specified = 2037,
  tuple_keys_too_many_or_too_few_items = 2038,
  page_size_invalid = 2039,
  param_missing_value = 2040,
  difference_base_missing_value = 2041,
  subtract_base_missing_value = 2042,
  object_too_long = 2043,
  relation_too_long = 2044,
  type_definitions_too_few_items = 2045,
  type_invalid_length = 2046,
  type_invalid_pattern = 2047,
  relations_too_few_items = 2048,
  relations_too_long = 2049,
  relations_invalid_pattern = 2050,
  object_invalid_pattern = 2051,
  query_string_type_continuation_token_mismatch = 2052,
  exceeded_entity_limit = 2053,
  invalid_contextual_tuple = 2054,
  duplicate_contextual_tuple = 2055,
  invalid_authorization_model = 2056,
  unsupported_schema_version = 2057,
  UNRECOGNIZED = -1,
}

export function errorCodeFromJSON(object: any): ErrorCode {
  switch (object) {
    case 0:
    case "no_error":
      return ErrorCode.no_error;
    case 2000:
    case "validation_error":
      return ErrorCode.validation_error;
    case 2001:
    case "authorization_model_not_found":
      return ErrorCode.authorization_model_not_found;
    case 2002:
    case "authorization_model_resolution_too_complex":
      return ErrorCode.authorization_model_resolution_too_complex;
    case 2003:
    case "invalid_write_input":
      return ErrorCode.invalid_write_input;
    case 2004:
    case "cannot_allow_duplicate_tuples_in_one_request":
      return ErrorCode.cannot_allow_duplicate_tuples_in_one_request;
    case 2005:
    case "cannot_allow_duplicate_types_in_one_request":
      return ErrorCode.cannot_allow_duplicate_types_in_one_request;
    case 2006:
    case "cannot_allow_multiple_references_to_one_relation":
      return ErrorCode.cannot_allow_multiple_references_to_one_relation;
    case 2007:
    case "invalid_continuation_token":
      return ErrorCode.invalid_continuation_token;
    case 2008:
    case "invalid_tuple_set":
      return ErrorCode.invalid_tuple_set;
    case 2009:
    case "invalid_check_input":
      return ErrorCode.invalid_check_input;
    case 2010:
    case "invalid_expand_input":
      return ErrorCode.invalid_expand_input;
    case 2011:
    case "unsupported_user_set":
      return ErrorCode.unsupported_user_set;
    case 2012:
    case "invalid_object_format":
      return ErrorCode.invalid_object_format;
    case 2017:
    case "write_failed_due_to_invalid_input":
      return ErrorCode.write_failed_due_to_invalid_input;
    case 2018:
    case "authorization_model_assertions_not_found":
      return ErrorCode.authorization_model_assertions_not_found;
    case 2020:
    case "latest_authorization_model_not_found":
      return ErrorCode.latest_authorization_model_not_found;
    case 2021:
    case "type_not_found":
      return ErrorCode.type_not_found;
    case 2022:
    case "relation_not_found":
      return ErrorCode.relation_not_found;
    case 2023:
    case "empty_relation_definition":
      return ErrorCode.empty_relation_definition;
    case 2025:
    case "invalid_user":
      return ErrorCode.invalid_user;
    case 2027:
    case "invalid_tuple":
      return ErrorCode.invalid_tuple;
    case 2028:
    case "unknown_relation":
      return ErrorCode.unknown_relation;
    case 2030:
    case "store_id_invalid_length":
      return ErrorCode.store_id_invalid_length;
    case 2033:
    case "assertions_too_many_items":
      return ErrorCode.assertions_too_many_items;
    case 2034:
    case "id_too_long":
      return ErrorCode.id_too_long;
    case 2036:
    case "authorization_model_id_too_long":
      return ErrorCode.authorization_model_id_too_long;
    case 2037:
    case "tuple_key_value_not_specified":
      return ErrorCode.tuple_key_value_not_specified;
    case 2038:
    case "tuple_keys_too_many_or_too_few_items":
      return ErrorCode.tuple_keys_too_many_or_too_few_items;
    case 2039:
    case "page_size_invalid":
      return ErrorCode.page_size_invalid;
    case 2040:
    case "param_missing_value":
      return ErrorCode.param_missing_value;
    case 2041:
    case "difference_base_missing_value":
      return ErrorCode.difference_base_missing_value;
    case 2042:
    case "subtract_base_missing_value":
      return ErrorCode.subtract_base_missing_value;
    case 2043:
    case "object_too_long":
      return ErrorCode.object_too_long;
    case 2044:
    case "relation_too_long":
      return ErrorCode.relation_too_long;
    case 2045:
    case "type_definitions_too_few_items":
      return ErrorCode.type_definitions_too_few_items;
    case 2046:
    case "type_invalid_length":
      return ErrorCode.type_invalid_length;
    case 2047:
    case "type_invalid_pattern":
      return ErrorCode.type_invalid_pattern;
    case 2048:
    case "relations_too_few_items":
      return ErrorCode.relations_too_few_items;
    case 2049:
    case "relations_too_long":
      return ErrorCode.relations_too_long;
    case 2050:
    case "relations_invalid_pattern":
      return ErrorCode.relations_invalid_pattern;
    case 2051:
    case "object_invalid_pattern":
      return ErrorCode.object_invalid_pattern;
    case 2052:
    case "query_string_type_continuation_token_mismatch":
      return ErrorCode.query_string_type_continuation_token_mismatch;
    case 2053:
    case "exceeded_entity_limit":
      return ErrorCode.exceeded_entity_limit;
    case 2054:
    case "invalid_contextual_tuple":
      return ErrorCode.invalid_contextual_tuple;
    case 2055:
    case "duplicate_contextual_tuple":
      return ErrorCode.duplicate_contextual_tuple;
    case 2056:
    case "invalid_authorization_model":
      return ErrorCode.invalid_authorization_model;
    case 2057:
    case "unsupported_schema_version":
      return ErrorCode.unsupported_schema_version;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ErrorCode.UNRECOGNIZED;
  }
}

export function errorCodeToJSON(object: ErrorCode): string {
  switch (object) {
    case ErrorCode.no_error:
      return "no_error";
    case ErrorCode.validation_error:
      return "validation_error";
    case ErrorCode.authorization_model_not_found:
      return "authorization_model_not_found";
    case ErrorCode.authorization_model_resolution_too_complex:
      return "authorization_model_resolution_too_complex";
    case ErrorCode.invalid_write_input:
      return "invalid_write_input";
    case ErrorCode.cannot_allow_duplicate_tuples_in_one_request:
      return "cannot_allow_duplicate_tuples_in_one_request";
    case ErrorCode.cannot_allow_duplicate_types_in_one_request:
      return "cannot_allow_duplicate_types_in_one_request";
    case ErrorCode.cannot_allow_multiple_references_to_one_relation:
      return "cannot_allow_multiple_references_to_one_relation";
    case ErrorCode.invalid_continuation_token:
      return "invalid_continuation_token";
    case ErrorCode.invalid_tuple_set:
      return "invalid_tuple_set";
    case ErrorCode.invalid_check_input:
      return "invalid_check_input";
    case ErrorCode.invalid_expand_input:
      return "invalid_expand_input";
    case ErrorCode.unsupported_user_set:
      return "unsupported_user_set";
    case ErrorCode.invalid_object_format:
      return "invalid_object_format";
    case ErrorCode.write_failed_due_to_invalid_input:
      return "write_failed_due_to_invalid_input";
    case ErrorCode.authorization_model_assertions_not_found:
      return "authorization_model_assertions_not_found";
    case ErrorCode.latest_authorization_model_not_found:
      return "latest_authorization_model_not_found";
    case ErrorCode.type_not_found:
      return "type_not_found";
    case ErrorCode.relation_not_found:
      return "relation_not_found";
    case ErrorCode.empty_relation_definition:
      return "empty_relation_definition";
    case ErrorCode.invalid_user:
      return "invalid_user";
    case ErrorCode.invalid_tuple:
      return "invalid_tuple";
    case ErrorCode.unknown_relation:
      return "unknown_relation";
    case ErrorCode.store_id_invalid_length:
      return "store_id_invalid_length";
    case ErrorCode.assertions_too_many_items:
      return "assertions_too_many_items";
    case ErrorCode.id_too_long:
      return "id_too_long";
    case ErrorCode.authorization_model_id_too_long:
      return "authorization_model_id_too_long";
    case ErrorCode.tuple_key_value_not_specified:
      return "tuple_key_value_not_specified";
    case ErrorCode.tuple_keys_too_many_or_too_few_items:
      return "tuple_keys_too_many_or_too_few_items";
    case ErrorCode.page_size_invalid:
      return "page_size_invalid";
    case ErrorCode.param_missing_value:
      return "param_missing_value";
    case ErrorCode.difference_base_missing_value:
      return "difference_base_missing_value";
    case ErrorCode.subtract_base_missing_value:
      return "subtract_base_missing_value";
    case ErrorCode.object_too_long:
      return "object_too_long";
    case ErrorCode.relation_too_long:
      return "relation_too_long";
    case ErrorCode.type_definitions_too_few_items:
      return "type_definitions_too_few_items";
    case ErrorCode.type_invalid_length:
      return "type_invalid_length";
    case ErrorCode.type_invalid_pattern:
      return "type_invalid_pattern";
    case ErrorCode.relations_too_few_items:
      return "relations_too_few_items";
    case ErrorCode.relations_too_long:
      return "relations_too_long";
    case ErrorCode.relations_invalid_pattern:
      return "relations_invalid_pattern";
    case ErrorCode.object_invalid_pattern:
      return "object_invalid_pattern";
    case ErrorCode.query_string_type_continuation_token_mismatch:
      return "query_string_type_continuation_token_mismatch";
    case ErrorCode.exceeded_entity_limit:
      return "exceeded_entity_limit";
    case ErrorCode.invalid_contextual_tuple:
      return "invalid_contextual_tuple";
    case ErrorCode.duplicate_contextual_tuple:
      return "duplicate_contextual_tuple";
    case ErrorCode.invalid_authorization_model:
      return "invalid_authorization_model";
    case ErrorCode.unsupported_schema_version:
      return "unsupported_schema_version";
    case ErrorCode.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum InternalErrorCode {
  no_internal_error = 0,
  internal_error = 4000,
  cancelled = 4003,
  deadline_exceeded = 4004,
  already_exists = 4005,
  resource_exhausted = 4006,
  failed_precondition = 4007,
  aborted = 4008,
  out_of_range = 4009,
  unavailable = 4010,
  data_loss = 4011,
  UNRECOGNIZED = -1,
}

export function internalErrorCodeFromJSON(object: any): InternalErrorCode {
  switch (object) {
    case 0:
    case "no_internal_error":
      return InternalErrorCode.no_internal_error;
    case 4000:
    case "internal_error":
      return InternalErrorCode.internal_error;
    case 4003:
    case "cancelled":
      return InternalErrorCode.cancelled;
    case 4004:
    case "deadline_exceeded":
      return InternalErrorCode.deadline_exceeded;
    case 4005:
    case "already_exists":
      return InternalErrorCode.already_exists;
    case 4006:
    case "resource_exhausted":
      return InternalErrorCode.resource_exhausted;
    case 4007:
    case "failed_precondition":
      return InternalErrorCode.failed_precondition;
    case 4008:
    case "aborted":
      return InternalErrorCode.aborted;
    case 4009:
    case "out_of_range":
      return InternalErrorCode.out_of_range;
    case 4010:
    case "unavailable":
      return InternalErrorCode.unavailable;
    case 4011:
    case "data_loss":
      return InternalErrorCode.data_loss;
    case -1:
    case "UNRECOGNIZED":
    default:
      return InternalErrorCode.UNRECOGNIZED;
  }
}

export function internalErrorCodeToJSON(object: InternalErrorCode): string {
  switch (object) {
    case InternalErrorCode.no_internal_error:
      return "no_internal_error";
    case InternalErrorCode.internal_error:
      return "internal_error";
    case InternalErrorCode.cancelled:
      return "cancelled";
    case InternalErrorCode.deadline_exceeded:
      return "deadline_exceeded";
    case InternalErrorCode.already_exists:
      return "already_exists";
    case InternalErrorCode.resource_exhausted:
      return "resource_exhausted";
    case InternalErrorCode.failed_precondition:
      return "failed_precondition";
    case InternalErrorCode.aborted:
      return "aborted";
    case InternalErrorCode.out_of_range:
      return "out_of_range";
    case InternalErrorCode.unavailable:
      return "unavailable";
    case InternalErrorCode.data_loss:
      return "data_loss";
    case InternalErrorCode.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum NotFoundErrorCode {
  no_not_found_error = 0,
  undefined_endpoint = 5000,
  store_id_not_found = 5002,
  unimplemented = 5004,
  UNRECOGNIZED = -1,
}

export function notFoundErrorCodeFromJSON(object: any): NotFoundErrorCode {
  switch (object) {
    case 0:
    case "no_not_found_error":
      return NotFoundErrorCode.no_not_found_error;
    case 5000:
    case "undefined_endpoint":
      return NotFoundErrorCode.undefined_endpoint;
    case 5002:
    case "store_id_not_found":
      return NotFoundErrorCode.store_id_not_found;
    case 5004:
    case "unimplemented":
      return NotFoundErrorCode.unimplemented;
    case -1:
    case "UNRECOGNIZED":
    default:
      return NotFoundErrorCode.UNRECOGNIZED;
  }
}

export function notFoundErrorCodeToJSON(object: NotFoundErrorCode): string {
  switch (object) {
    case NotFoundErrorCode.no_not_found_error:
      return "no_not_found_error";
    case NotFoundErrorCode.undefined_endpoint:
      return "undefined_endpoint";
    case NotFoundErrorCode.store_id_not_found:
      return "store_id_not_found";
    case NotFoundErrorCode.unimplemented:
      return "unimplemented";
    case NotFoundErrorCode.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface ValidationErrorMessageResponse {
  code: ErrorCode;
  message: string;
}

export interface InternalErrorMessageResponse {
  code: InternalErrorCode;
  message: string;
}

export interface PathUnknownErrorMessageResponse {
  code: NotFoundErrorCode;
  message: string;
}

export interface ErrorMessageRequest {
}

function createBaseValidationErrorMessageResponse(): ValidationErrorMessageResponse {
  return { code: 0, message: "" };
}

export const ValidationErrorMessageResponse = {
  encode(message: ValidationErrorMessageResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ValidationErrorMessageResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseValidationErrorMessageResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.code = reader.int32() as any;
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.message = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ValidationErrorMessageResponse {
    return {
      code: isSet(object.code) ? errorCodeFromJSON(object.code) : 0,
      message: isSet(object.message) ? globalThis.String(object.message) : "",
    };
  },

  toJSON(message: ValidationErrorMessageResponse): unknown {
    const obj: any = {};
    if (message.code !== 0) {
      obj.code = errorCodeToJSON(message.code);
    }
    if (message.message !== "") {
      obj.message = message.message;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ValidationErrorMessageResponse>, I>>(base?: I): ValidationErrorMessageResponse {
    return ValidationErrorMessageResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ValidationErrorMessageResponse>, I>>(
    object: I,
  ): ValidationErrorMessageResponse {
    const message = createBaseValidationErrorMessageResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    return message;
  },
};

function createBaseInternalErrorMessageResponse(): InternalErrorMessageResponse {
  return { code: 0, message: "" };
}

export const InternalErrorMessageResponse = {
  encode(message: InternalErrorMessageResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): InternalErrorMessageResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInternalErrorMessageResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.code = reader.int32() as any;
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.message = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): InternalErrorMessageResponse {
    return {
      code: isSet(object.code) ? internalErrorCodeFromJSON(object.code) : 0,
      message: isSet(object.message) ? globalThis.String(object.message) : "",
    };
  },

  toJSON(message: InternalErrorMessageResponse): unknown {
    const obj: any = {};
    if (message.code !== 0) {
      obj.code = internalErrorCodeToJSON(message.code);
    }
    if (message.message !== "") {
      obj.message = message.message;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<InternalErrorMessageResponse>, I>>(base?: I): InternalErrorMessageResponse {
    return InternalErrorMessageResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<InternalErrorMessageResponse>, I>>(object: I): InternalErrorMessageResponse {
    const message = createBaseInternalErrorMessageResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    return message;
  },
};

function createBasePathUnknownErrorMessageResponse(): PathUnknownErrorMessageResponse {
  return { code: 0, message: "" };
}

export const PathUnknownErrorMessageResponse = {
  encode(message: PathUnknownErrorMessageResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PathUnknownErrorMessageResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePathUnknownErrorMessageResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.code = reader.int32() as any;
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.message = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PathUnknownErrorMessageResponse {
    return {
      code: isSet(object.code) ? notFoundErrorCodeFromJSON(object.code) : 0,
      message: isSet(object.message) ? globalThis.String(object.message) : "",
    };
  },

  toJSON(message: PathUnknownErrorMessageResponse): unknown {
    const obj: any = {};
    if (message.code !== 0) {
      obj.code = notFoundErrorCodeToJSON(message.code);
    }
    if (message.message !== "") {
      obj.message = message.message;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<PathUnknownErrorMessageResponse>, I>>(base?: I): PathUnknownErrorMessageResponse {
    return PathUnknownErrorMessageResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<PathUnknownErrorMessageResponse>, I>>(
    object: I,
  ): PathUnknownErrorMessageResponse {
    const message = createBasePathUnknownErrorMessageResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    return message;
  },
};

function createBaseErrorMessageRequest(): ErrorMessageRequest {
  return {};
}

export const ErrorMessageRequest = {
  encode(_: ErrorMessageRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ErrorMessageRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseErrorMessageRequest();
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

  fromJSON(_: any): ErrorMessageRequest {
    return {};
  },

  toJSON(_: ErrorMessageRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<ErrorMessageRequest>, I>>(base?: I): ErrorMessageRequest {
    return ErrorMessageRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ErrorMessageRequest>, I>>(_: I): ErrorMessageRequest {
    const message = createBaseErrorMessageRequest();
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
