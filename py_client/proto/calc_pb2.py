# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/calc.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/calc.proto',
  package='proto',
  syntax='proto3',
  serialized_options=b'Z example.com/sandeepkp/calc;proto',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x10proto/calc.proto\x12\x05proto\x1a\x1cgoogle/api/annotations.proto\"\x1f\n\x07Request\x12\t\n\x01\x61\x18\x01 \x01(\x03\x12\t\n\x01\x62\x18\x02 \x01(\x03\"\x1a\n\x08Response\x12\x0e\n\x06result\x18\x01 \x01(\x03\x32\x91\x02\n\x0b\x43\x61lcService\x12?\n\x03\x41\x64\x64\x12\x0e.proto.Request\x1a\x0f.proto.Response\"\x17\x82\xd3\xe4\x93\x02\x11\"\x0c/v1/calc/add:\x01*\x12?\n\x03Sub\x12\x0e.proto.Request\x1a\x0f.proto.Response\"\x17\x82\xd3\xe4\x93\x02\x11\"\x0c/v1/calc/sub:\x01*\x12?\n\x03Mul\x12\x0e.proto.Request\x1a\x0f.proto.Response\"\x17\x82\xd3\xe4\x93\x02\x11\"\x0c/v1/calc/mul:\x01*\x12?\n\x03\x44iv\x12\x0e.proto.Request\x1a\x0f.proto.Response\"\x17\x82\xd3\xe4\x93\x02\x11\"\x0c/v1/calc/div:\x01*B\"Z example.com/sandeepkp/calc;protob\x06proto3'
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,])




_REQUEST = _descriptor.Descriptor(
  name='Request',
  full_name='proto.Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='a', full_name='proto.Request.a', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='b', full_name='proto.Request.b', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=57,
  serialized_end=88,
)


_RESPONSE = _descriptor.Descriptor(
  name='Response',
  full_name='proto.Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='result', full_name='proto.Response.result', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=90,
  serialized_end=116,
)

DESCRIPTOR.message_types_by_name['Request'] = _REQUEST
DESCRIPTOR.message_types_by_name['Response'] = _RESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Request = _reflection.GeneratedProtocolMessageType('Request', (_message.Message,), {
  'DESCRIPTOR' : _REQUEST,
  '__module__' : 'proto.calc_pb2'
  # @@protoc_insertion_point(class_scope:proto.Request)
  })
_sym_db.RegisterMessage(Request)

Response = _reflection.GeneratedProtocolMessageType('Response', (_message.Message,), {
  'DESCRIPTOR' : _RESPONSE,
  '__module__' : 'proto.calc_pb2'
  # @@protoc_insertion_point(class_scope:proto.Response)
  })
_sym_db.RegisterMessage(Response)


DESCRIPTOR._options = None

_CALCSERVICE = _descriptor.ServiceDescriptor(
  name='CalcService',
  full_name='proto.CalcService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=119,
  serialized_end=392,
  methods=[
  _descriptor.MethodDescriptor(
    name='Add',
    full_name='proto.CalcService.Add',
    index=0,
    containing_service=None,
    input_type=_REQUEST,
    output_type=_RESPONSE,
    serialized_options=b'\202\323\344\223\002\021\"\014/v1/calc/add:\001*',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='Sub',
    full_name='proto.CalcService.Sub',
    index=1,
    containing_service=None,
    input_type=_REQUEST,
    output_type=_RESPONSE,
    serialized_options=b'\202\323\344\223\002\021\"\014/v1/calc/sub:\001*',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='Mul',
    full_name='proto.CalcService.Mul',
    index=2,
    containing_service=None,
    input_type=_REQUEST,
    output_type=_RESPONSE,
    serialized_options=b'\202\323\344\223\002\021\"\014/v1/calc/mul:\001*',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='Div',
    full_name='proto.CalcService.Div',
    index=3,
    containing_service=None,
    input_type=_REQUEST,
    output_type=_RESPONSE,
    serialized_options=b'\202\323\344\223\002\021\"\014/v1/calc/div:\001*',
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_CALCSERVICE)

DESCRIPTOR.services_by_name['CalcService'] = _CALCSERVICE

# @@protoc_insertion_point(module_scope)
