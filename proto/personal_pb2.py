# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/personal.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from proto import date_pb2 as proto_dot_date__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/personal.proto',
  package='person',
  syntax='proto3',
  serialized_options=b'Z\n./personpb',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x14proto/personal.proto\x12\x06person\x1a\x10proto/date.proto\"\xfc\x02\n\x08Personal\x12\x12\n\nfirst_name\x18\x01 \x01(\t\x12\x13\n\x0bsecond_name\x18\x02 \x01(\t\x12\x18\n\x10linkedin_profile\x18\x03 \x01(\t\x12\x15\n\rphone_numbers\x18\x04 \x03(\t\x12\x1d\n\x15personal_website_link\x18\x05 \x01(\t\x12\x13\n\x0bgithub_link\x18\x06 \x01(\t\x12\x17\n\x0fheader_finished\x18\x07 \x01(\x08\x12\x17\n\x0fprofile_picture\x18\x08 \x01(\x0c\x12\x13\n\x0bpage_number\x18\t \x01(\x05\x12*\n\x0f\x65\x64ucation_level\x18\n \x01(\x0e\x32\x11.person.Education\x12\"\n\x0egraduation_day\x18\x0b \x01(\x0b\x32\n.date.Date\x12\"\n\taddresses\x18\x0c \x03(\x0b\x32\x0f.person.Address\x12\'\n\x0esingle_address\x18\r \x01(\x0b\x32\x0f.person.Address\"o\n\x07\x41\x64\x64ress\x12\x13\n\x0b\x61\x64\x64ress_one\x18\x01 \x01(\t\x12\x13\n\x0b\x61\x64\x64ress_two\x18\x02 \x01(\t\x12\x13\n\x0bpostal_code\x18\x03 \x01(\t\x12\x17\n\x0fprefecture_name\x18\x04 \x01(\t\x12\x0c\n\x04ward\x18\x05 \x01(\t*4\n\tEducation\x12\x0b\n\x07\x44\x45\x46\x41ULT\x10\x00\x12\x0b\n\x07\x44IPLOMA\x10\x01\x12\r\n\tBACHELORS\x10\x02\x42\x0cZ\n./personpbb\x06proto3'
  ,
  dependencies=[proto_dot_date__pb2.DESCRIPTOR,])

_EDUCATION = _descriptor.EnumDescriptor(
  name='Education',
  full_name='person.Education',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='DEFAULT', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='DIPLOMA', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='BACHELORS', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=546,
  serialized_end=598,
)
_sym_db.RegisterEnumDescriptor(_EDUCATION)

Education = enum_type_wrapper.EnumTypeWrapper(_EDUCATION)
DEFAULT = 0
DIPLOMA = 1
BACHELORS = 2



_PERSONAL = _descriptor.Descriptor(
  name='Personal',
  full_name='person.Personal',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='first_name', full_name='person.Personal.first_name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='second_name', full_name='person.Personal.second_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='linkedin_profile', full_name='person.Personal.linkedin_profile', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='phone_numbers', full_name='person.Personal.phone_numbers', index=3,
      number=4, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='personal_website_link', full_name='person.Personal.personal_website_link', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='github_link', full_name='person.Personal.github_link', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='header_finished', full_name='person.Personal.header_finished', index=6,
      number=7, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='profile_picture', full_name='person.Personal.profile_picture', index=7,
      number=8, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='page_number', full_name='person.Personal.page_number', index=8,
      number=9, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='education_level', full_name='person.Personal.education_level', index=9,
      number=10, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='graduation_day', full_name='person.Personal.graduation_day', index=10,
      number=11, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='addresses', full_name='person.Personal.addresses', index=11,
      number=12, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='single_address', full_name='person.Personal.single_address', index=12,
      number=13, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=51,
  serialized_end=431,
)


_ADDRESS = _descriptor.Descriptor(
  name='Address',
  full_name='person.Address',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='address_one', full_name='person.Address.address_one', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='address_two', full_name='person.Address.address_two', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='postal_code', full_name='person.Address.postal_code', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='prefecture_name', full_name='person.Address.prefecture_name', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ward', full_name='person.Address.ward', index=4,
      number=5, type=9, cpp_type=9, label=1,
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
  serialized_start=433,
  serialized_end=544,
)

_PERSONAL.fields_by_name['education_level'].enum_type = _EDUCATION
_PERSONAL.fields_by_name['graduation_day'].message_type = proto_dot_date__pb2._DATE
_PERSONAL.fields_by_name['addresses'].message_type = _ADDRESS
_PERSONAL.fields_by_name['single_address'].message_type = _ADDRESS
DESCRIPTOR.message_types_by_name['Personal'] = _PERSONAL
DESCRIPTOR.message_types_by_name['Address'] = _ADDRESS
DESCRIPTOR.enum_types_by_name['Education'] = _EDUCATION
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Personal = _reflection.GeneratedProtocolMessageType('Personal', (_message.Message,), {
  'DESCRIPTOR' : _PERSONAL,
  '__module__' : 'proto.personal_pb2'
  # @@protoc_insertion_point(class_scope:person.Personal)
  })
_sym_db.RegisterMessage(Personal)

Address = _reflection.GeneratedProtocolMessageType('Address', (_message.Message,), {
  'DESCRIPTOR' : _ADDRESS,
  '__module__' : 'proto.personal_pb2'
  # @@protoc_insertion_point(class_scope:person.Address)
  })
_sym_db.RegisterMessage(Address)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
