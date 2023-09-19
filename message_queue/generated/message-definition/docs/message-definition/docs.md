# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [message_definition/v1/errors_ignore.proto](#message_definition_v1_errors_ignore-proto)
    - [ErrorMessageRequest](#message_definition-v1-ErrorMessageRequest)
    - [InternalErrorMessageResponse](#message_definition-v1-InternalErrorMessageResponse)
    - [PathUnknownErrorMessageResponse](#message_definition-v1-PathUnknownErrorMessageResponse)
    - [ValidationErrorMessageResponse](#message_definition-v1-ValidationErrorMessageResponse)
  
    - [AuthErrorCode](#message_definition-v1-AuthErrorCode)
    - [ErrorCode](#message_definition-v1-ErrorCode)
    - [InternalErrorCode](#message_definition-v1-InternalErrorCode)
    - [NotFoundErrorCode](#message_definition-v1-NotFoundErrorCode)
  
- [message_definition/v1/message.proto](#message_definition_v1_message-proto)
    - [FinancialIntegrationServiceMessageFormat](#message_definition-v1-FinancialIntegrationServiceMessageFormat)
    - [HeadlessAuthenticationServiceDeleteAccountMessageFormat](#message_definition-v1-HeadlessAuthenticationServiceDeleteAccountMessageFormat)
    - [SocialServiceMessageFormat](#message_definition-v1-SocialServiceMessageFormat)
  
- [Scalar Value Types](#scalar-value-types)



<a name="message_definition_v1_errors_ignore-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## message_definition/v1/errors_ignore.proto



<a name="message_definition-v1-ErrorMessageRequest"></a>

### ErrorMessageRequest







<a name="message_definition-v1-InternalErrorMessageResponse"></a>

### InternalErrorMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [InternalErrorCode](#message_definition-v1-InternalErrorCode) |  |  |
| message | [string](#string) |  |  |






<a name="message_definition-v1-PathUnknownErrorMessageResponse"></a>

### PathUnknownErrorMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [NotFoundErrorCode](#message_definition-v1-NotFoundErrorCode) |  |  |
| message | [string](#string) |  |  |






<a name="message_definition-v1-ValidationErrorMessageResponse"></a>

### ValidationErrorMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [ErrorCode](#message_definition-v1-ErrorCode) |  |  |
| message | [string](#string) |  |  |





 


<a name="message_definition-v1-AuthErrorCode"></a>

### AuthErrorCode


| Name | Number | Description |
| ---- | ------ | ----------- |
| no_auth_error | 0 |  |
| auth_failed_invalid_subject | 1001 |  |
| auth_failed_invalid_audience | 1002 |  |
| auth_failed_invalid_issuer | 1003 |  |
| invalid_claims | 1004 |  |
| auth_failed_invalid_bearer_token | 1005 |  |
| bearer_token_missing | 1010 |  |
| unauthenticated | 1500 |  |



<a name="message_definition-v1-ErrorCode"></a>

### ErrorCode


| Name | Number | Description |
| ---- | ------ | ----------- |
| no_error | 0 |  |
| validation_error | 2000 |  |
| authorization_model_not_found | 2001 |  |
| authorization_model_resolution_too_complex | 2002 |  |
| invalid_write_input | 2003 |  |
| cannot_allow_duplicate_tuples_in_one_request | 2004 |  |
| cannot_allow_duplicate_types_in_one_request | 2005 |  |
| cannot_allow_multiple_references_to_one_relation | 2006 |  |
| invalid_continuation_token | 2007 |  |
| invalid_tuple_set | 2008 |  |
| invalid_check_input | 2009 |  |
| invalid_expand_input | 2010 |  |
| unsupported_user_set | 2011 |  |
| invalid_object_format | 2012 |  |
| write_failed_due_to_invalid_input | 2017 |  |
| authorization_model_assertions_not_found | 2018 |  |
| latest_authorization_model_not_found | 2020 |  |
| type_not_found | 2021 |  |
| relation_not_found | 2022 |  |
| empty_relation_definition | 2023 |  |
| invalid_user | 2025 |  |
| invalid_tuple | 2027 |  |
| unknown_relation | 2028 |  |
| store_id_invalid_length | 2030 |  |
| assertions_too_many_items | 2033 |  |
| id_too_long | 2034 |  |
| authorization_model_id_too_long | 2036 |  |
| tuple_key_value_not_specified | 2037 |  |
| tuple_keys_too_many_or_too_few_items | 2038 |  |
| page_size_invalid | 2039 |  |
| param_missing_value | 2040 |  |
| difference_base_missing_value | 2041 |  |
| subtract_base_missing_value | 2042 |  |
| object_too_long | 2043 |  |
| relation_too_long | 2044 |  |
| type_definitions_too_few_items | 2045 |  |
| type_invalid_length | 2046 |  |
| type_invalid_pattern | 2047 |  |
| relations_too_few_items | 2048 |  |
| relations_too_long | 2049 |  |
| relations_invalid_pattern | 2050 |  |
| object_invalid_pattern | 2051 |  |
| query_string_type_continuation_token_mismatch | 2052 |  |
| exceeded_entity_limit | 2053 |  |
| invalid_contextual_tuple | 2054 |  |
| duplicate_contextual_tuple | 2055 |  |
| invalid_authorization_model | 2056 |  |
| unsupported_schema_version | 2057 |  |



<a name="message_definition-v1-InternalErrorCode"></a>

### InternalErrorCode


| Name | Number | Description |
| ---- | ------ | ----------- |
| no_internal_error | 0 |  |
| internal_error | 4000 |  |
| cancelled | 4003 |  |
| deadline_exceeded | 4004 |  |
| already_exists | 4005 |  |
| resource_exhausted | 4006 |  |
| failed_precondition | 4007 |  |
| aborted | 4008 |  |
| out_of_range | 4009 |  |
| unavailable | 4010 |  |
| data_loss | 4011 |  |



<a name="message_definition-v1-NotFoundErrorCode"></a>

### NotFoundErrorCode


| Name | Number | Description |
| ---- | ------ | ----------- |
| no_not_found_error | 0 |  |
| undefined_endpoint | 5000 |  |
| store_id_not_found | 5002 |  |
| unimplemented | 5004 |  |


 

 

 



<a name="message_definition_v1_message-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## message_definition/v1/message.proto



<a name="message_definition-v1-FinancialIntegrationServiceMessageFormat"></a>

### FinancialIntegrationServiceMessageFormat
FinancialIntegrationServiceMessageFormat: represents an sqs message format
for deleting an account via the financial integration service


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint64](#uint64) |  | user_id id from the vantage point of the user service |






<a name="message_definition-v1-HeadlessAuthenticationServiceDeleteAccountMessageFormat"></a>

### HeadlessAuthenticationServiceDeleteAccountMessageFormat
HeadlessAuthenticationServiceDeleteAccountMessageFormat: represents an sqs message format
for deleting an account via the headless authentication service


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| authn_id | [uint64](#uint64) |  | authn id which is the id of the account from the vantage point of the authentication service |
| email | [string](#string) |  | account email Validations: - must be an email and required |






<a name="message_definition-v1-SocialServiceMessageFormat"></a>

### SocialServiceMessageFormat
SocialServiceMessageFormat: represents an sqs message format
for deleting an account via the financial integration service


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint64](#uint64) |  | user_id id from the vantage point of the social service |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

