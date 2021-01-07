# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: media.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='media.proto',
  package='protobuf',
  syntax='proto3',
  serialized_options=b'Z\"github.com/bic4907/webrtc/protobuf',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0bmedia.proto\x12\x08protobuf\"\x1d\n\nReqMessage\x12\x0f\n\x07\x43ontent\x18\x01 \x01(\t\"\x1d\n\nRplMessage\x12\x0f\n\x07\x43ontent\x18\x01 \x01(\t\"N\n\nVideoChunk\x12\x0e\n\x06RoomId\x18\x01 \x01(\t\x12\x0e\n\x06UserId\x18\x02 \x01(\t\x12\r\n\x05\x43hunk\x18\x03 \x01(\x0c\x12\x11\n\tCreatedAt\x18\x04 \x01(\x03\"\x1e\n\x0cReceiveReply\x12\x0e\n\x06Result\x18\x01 \x01(\t2\x83\x01\n\x07Service\x12\x39\n\x0bSendMessage\x12\x14.protobuf.ReqMessage\x1a\x14.protobuf.RplMessage\x12=\n\x0bStreamVideo\x12\x14.protobuf.VideoChunk\x1a\x16.protobuf.ReceiveReply(\x01\x42$Z\"github.com/bic4907/webrtc/protobufb\x06proto3'
)




_REQMESSAGE = _descriptor.Descriptor(
  name='ReqMessage',
  full_name='protobuf.ReqMessage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='Content', full_name='protobuf.ReqMessage.Content', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
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
  serialized_start=25,
  serialized_end=54,
)


_RPLMESSAGE = _descriptor.Descriptor(
  name='RplMessage',
  full_name='protobuf.RplMessage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='Content', full_name='protobuf.RplMessage.Content', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
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
  serialized_start=56,
  serialized_end=85,
)


_VIDEOCHUNK = _descriptor.Descriptor(
  name='VideoChunk',
  full_name='protobuf.VideoChunk',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='RoomId', full_name='protobuf.VideoChunk.RoomId', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='UserId', full_name='protobuf.VideoChunk.UserId', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Chunk', full_name='protobuf.VideoChunk.Chunk', index=2,
      number=3, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='CreatedAt', full_name='protobuf.VideoChunk.CreatedAt', index=3,
      number=4, type=3, cpp_type=2, label=1,
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
  serialized_start=87,
  serialized_end=165,
)


_RECEIVEREPLY = _descriptor.Descriptor(
  name='ReceiveReply',
  full_name='protobuf.ReceiveReply',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='Result', full_name='protobuf.ReceiveReply.Result', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
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
  serialized_start=167,
  serialized_end=197,
)

DESCRIPTOR.message_types_by_name['ReqMessage'] = _REQMESSAGE
DESCRIPTOR.message_types_by_name['RplMessage'] = _RPLMESSAGE
DESCRIPTOR.message_types_by_name['VideoChunk'] = _VIDEOCHUNK
DESCRIPTOR.message_types_by_name['ReceiveReply'] = _RECEIVEREPLY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ReqMessage = _reflection.GeneratedProtocolMessageType('ReqMessage', (_message.Message,), {
  'DESCRIPTOR' : _REQMESSAGE,
  '__module__' : 'media_pb2'
  # @@protoc_insertion_point(class_scope:protobuf.ReqMessage)
  })
_sym_db.RegisterMessage(ReqMessage)

RplMessage = _reflection.GeneratedProtocolMessageType('RplMessage', (_message.Message,), {
  'DESCRIPTOR' : _RPLMESSAGE,
  '__module__' : 'media_pb2'
  # @@protoc_insertion_point(class_scope:protobuf.RplMessage)
  })
_sym_db.RegisterMessage(RplMessage)

VideoChunk = _reflection.GeneratedProtocolMessageType('VideoChunk', (_message.Message,), {
  'DESCRIPTOR' : _VIDEOCHUNK,
  '__module__' : 'media_pb2'
  # @@protoc_insertion_point(class_scope:protobuf.VideoChunk)
  })
_sym_db.RegisterMessage(VideoChunk)

ReceiveReply = _reflection.GeneratedProtocolMessageType('ReceiveReply', (_message.Message,), {
  'DESCRIPTOR' : _RECEIVEREPLY,
  '__module__' : 'media_pb2'
  # @@protoc_insertion_point(class_scope:protobuf.ReceiveReply)
  })
_sym_db.RegisterMessage(ReceiveReply)


DESCRIPTOR._options = None

_SERVICE = _descriptor.ServiceDescriptor(
  name='Service',
  full_name='protobuf.Service',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=200,
  serialized_end=331,
  methods=[
  _descriptor.MethodDescriptor(
    name='SendMessage',
    full_name='protobuf.Service.SendMessage',
    index=0,
    containing_service=None,
    input_type=_REQMESSAGE,
    output_type=_RPLMESSAGE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='StreamVideo',
    full_name='protobuf.Service.StreamVideo',
    index=1,
    containing_service=None,
    input_type=_VIDEOCHUNK,
    output_type=_RECEIVEREPLY,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_SERVICE)

DESCRIPTOR.services_by_name['Service'] = _SERVICE

# @@protoc_insertion_point(module_scope)