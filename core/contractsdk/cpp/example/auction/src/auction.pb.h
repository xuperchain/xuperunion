// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: auction.proto

#ifndef PROTOBUF_INCLUDED_auction_2eproto
#define PROTOBUF_INCLUDED_auction_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3007000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3007001 < PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers. Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/port_undef.inc>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/inlined_string_field.h>
#include <google/protobuf/metadata_lite.h>
#include <google/protobuf/message_lite.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/generated_enum_util.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_auction_2eproto

// Internal implementation detail -- do not use these members.
struct TableStruct_auction_2eproto {
  static const ::google::protobuf::internal::ParseTableField entries[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::AuxillaryParseTableField aux[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::ParseTable schema[2]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::FieldMetadata field_metadata[];
  static const ::google::protobuf::internal::SerializationTable serialization_table[];
  static const ::google::protobuf::uint32 offsets[];
};
namespace auction {
class Chip;
class ChipDefaultTypeInternal;
extern ChipDefaultTypeInternal _Chip_default_instance_;
class Lot;
class LotDefaultTypeInternal;
extern LotDefaultTypeInternal _Lot_default_instance_;
}  // namespace auction
namespace google {
namespace protobuf {
template<> ::auction::Chip* Arena::CreateMaybeMessage<::auction::Chip>(Arena*);
template<> ::auction::Lot* Arena::CreateMaybeMessage<::auction::Lot>(Arena*);
}  // namespace protobuf
}  // namespace google
namespace auction {

enum Status {
  INVALID = 0,
  PROGRESS = 1,
  ENDED = 2,
  Status_INT_MIN_SENTINEL_DO_NOT_USE_ = std::numeric_limits<::google::protobuf::int32>::min(),
  Status_INT_MAX_SENTINEL_DO_NOT_USE_ = std::numeric_limits<::google::protobuf::int32>::max()
};
bool Status_IsValid(int value);
const Status Status_MIN = INVALID;
const Status Status_MAX = ENDED;
const int Status_ARRAYSIZE = Status_MAX + 1;

// ===================================================================

class Lot :
    public ::google::protobuf::MessageLite /* @@protoc_insertion_point(class_definition:auction.Lot) */ {
 public:
  Lot();
  virtual ~Lot();

  Lot(const Lot& from);

  inline Lot& operator=(const Lot& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  Lot(Lot&& from) noexcept
    : Lot() {
    *this = ::std::move(from);
  }

  inline Lot& operator=(Lot&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const Lot& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const Lot* internal_default_instance() {
    return reinterpret_cast<const Lot*>(
               &_Lot_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  void Swap(Lot* other);
  friend void swap(Lot& a, Lot& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline Lot* New() const final {
    return CreateMaybeMessage<Lot>(nullptr);
  }

  Lot* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<Lot>(arena);
  }
  void CheckTypeAndMergeFrom(const ::google::protobuf::MessageLite& from)
    final;
  void CopyFrom(const Lot& from);
  void MergeFrom(const Lot& from);
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  #if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  static const char* _InternalParse(const char* begin, const char* end, void* object, ::google::protobuf::internal::ParseContext* ctx);
  ::google::protobuf::internal::ParseFunc _ParseFunc() const final { return _InternalParse; }
  #else
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) final;
  #endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const final;
  void DiscardUnknownFields();
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const;
  void InternalSwap(Lot* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return nullptr;
  }
  inline void* MaybeArenaPtr() const {
    return nullptr;
  }
  public:

  ::std::string GetTypeName() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // string id = 1;
  void clear_id();
  static const int kIdFieldNumber = 1;
  const ::std::string& id() const;
  void set_id(const ::std::string& value);
  #if LANG_CXX11
  void set_id(::std::string&& value);
  #endif
  void set_id(const char* value);
  void set_id(const char* value, size_t size);
  ::std::string* mutable_id();
  ::std::string* release_id();
  void set_allocated_id(::std::string* id);

  // string auctioneer = 2;
  void clear_auctioneer();
  static const int kAuctioneerFieldNumber = 2;
  const ::std::string& auctioneer() const;
  void set_auctioneer(const ::std::string& value);
  #if LANG_CXX11
  void set_auctioneer(::std::string&& value);
  #endif
  void set_auctioneer(const char* value);
  void set_auctioneer(const char* value, size_t size);
  ::std::string* mutable_auctioneer();
  ::std::string* release_auctioneer();
  void set_allocated_auctioneer(::std::string* auctioneer);

  // string data = 4;
  void clear_data();
  static const int kDataFieldNumber = 4;
  const ::std::string& data() const;
  void set_data(const ::std::string& value);
  #if LANG_CXX11
  void set_data(::std::string&& value);
  #endif
  void set_data(const char* value);
  void set_data(const char* value, size_t size);
  ::std::string* mutable_data();
  ::std::string* release_data();
  void set_allocated_data(::std::string* data);

  // string bidder = 5;
  void clear_bidder();
  static const int kBidderFieldNumber = 5;
  const ::std::string& bidder() const;
  void set_bidder(const ::std::string& value);
  #if LANG_CXX11
  void set_bidder(::std::string&& value);
  #endif
  void set_bidder(const char* value);
  void set_bidder(const char* value, size_t size);
  ::std::string* mutable_bidder();
  ::std::string* release_bidder();
  void set_allocated_bidder(::std::string* bidder);

  // uint64 floor = 3;
  void clear_floor();
  static const int kFloorFieldNumber = 3;
  ::google::protobuf::uint64 floor() const;
  void set_floor(::google::protobuf::uint64 value);

  // uint64 bid = 6;
  void clear_bid();
  static const int kBidFieldNumber = 6;
  ::google::protobuf::uint64 bid() const;
  void set_bid(::google::protobuf::uint64 value);

  // .auction.Status status = 7;
  void clear_status();
  static const int kStatusFieldNumber = 7;
  ::auction::Status status() const;
  void set_status(::auction::Status value);

  // @@protoc_insertion_point(class_scope:auction.Lot)
 private:
  class HasBitSetters;

  ::google::protobuf::internal::InternalMetadataWithArenaLite _internal_metadata_;
  ::google::protobuf::internal::ArenaStringPtr id_;
  ::google::protobuf::internal::ArenaStringPtr auctioneer_;
  ::google::protobuf::internal::ArenaStringPtr data_;
  ::google::protobuf::internal::ArenaStringPtr bidder_;
  ::google::protobuf::uint64 floor_;
  ::google::protobuf::uint64 bid_;
  int status_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  friend struct ::TableStruct_auction_2eproto;
};
// -------------------------------------------------------------------

class Chip :
    public ::google::protobuf::MessageLite /* @@protoc_insertion_point(class_definition:auction.Chip) */ {
 public:
  Chip();
  virtual ~Chip();

  Chip(const Chip& from);

  inline Chip& operator=(const Chip& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  Chip(Chip&& from) noexcept
    : Chip() {
    *this = ::std::move(from);
  }

  inline Chip& operator=(Chip&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const Chip& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const Chip* internal_default_instance() {
    return reinterpret_cast<const Chip*>(
               &_Chip_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  void Swap(Chip* other);
  friend void swap(Chip& a, Chip& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline Chip* New() const final {
    return CreateMaybeMessage<Chip>(nullptr);
  }

  Chip* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<Chip>(arena);
  }
  void CheckTypeAndMergeFrom(const ::google::protobuf::MessageLite& from)
    final;
  void CopyFrom(const Chip& from);
  void MergeFrom(const Chip& from);
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  #if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  static const char* _InternalParse(const char* begin, const char* end, void* object, ::google::protobuf::internal::ParseContext* ctx);
  ::google::protobuf::internal::ParseFunc _ParseFunc() const final { return _InternalParse; }
  #else
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) final;
  #endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const final;
  void DiscardUnknownFields();
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const;
  void InternalSwap(Chip* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return nullptr;
  }
  inline void* MaybeArenaPtr() const {
    return nullptr;
  }
  public:

  ::std::string GetTypeName() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // string bidder = 2;
  void clear_bidder();
  static const int kBidderFieldNumber = 2;
  const ::std::string& bidder() const;
  void set_bidder(const ::std::string& value);
  #if LANG_CXX11
  void set_bidder(::std::string&& value);
  #endif
  void set_bidder(const char* value);
  void set_bidder(const char* value, size_t size);
  ::std::string* mutable_bidder();
  ::std::string* release_bidder();
  void set_allocated_bidder(::std::string* bidder);

  // uint64 amount = 3;
  void clear_amount();
  static const int kAmountFieldNumber = 3;
  ::google::protobuf::uint64 amount() const;
  void set_amount(::google::protobuf::uint64 value);

  // uint64 bid = 4;
  void clear_bid();
  static const int kBidFieldNumber = 4;
  ::google::protobuf::uint64 bid() const;
  void set_bid(::google::protobuf::uint64 value);

  // .auction.Status status = 8;
  void clear_status();
  static const int kStatusFieldNumber = 8;
  ::auction::Status status() const;
  void set_status(::auction::Status value);

  // @@protoc_insertion_point(class_scope:auction.Chip)
 private:
  class HasBitSetters;

  ::google::protobuf::internal::InternalMetadataWithArenaLite _internal_metadata_;
  ::google::protobuf::internal::ArenaStringPtr bidder_;
  ::google::protobuf::uint64 amount_;
  ::google::protobuf::uint64 bid_;
  int status_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  friend struct ::TableStruct_auction_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// Lot

// string id = 1;
inline void Lot::clear_id() {
  id_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& Lot::id() const {
  // @@protoc_insertion_point(field_get:auction.Lot.id)
  return id_.GetNoArena();
}
inline void Lot::set_id(const ::std::string& value) {
  
  id_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:auction.Lot.id)
}
#if LANG_CXX11
inline void Lot::set_id(::std::string&& value) {
  
  id_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:auction.Lot.id)
}
#endif
inline void Lot::set_id(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  
  id_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:auction.Lot.id)
}
inline void Lot::set_id(const char* value, size_t size) {
  
  id_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:auction.Lot.id)
}
inline ::std::string* Lot::mutable_id() {
  
  // @@protoc_insertion_point(field_mutable:auction.Lot.id)
  return id_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* Lot::release_id() {
  // @@protoc_insertion_point(field_release:auction.Lot.id)
  
  return id_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void Lot::set_allocated_id(::std::string* id) {
  if (id != nullptr) {
    
  } else {
    
  }
  id_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), id);
  // @@protoc_insertion_point(field_set_allocated:auction.Lot.id)
}

// string auctioneer = 2;
inline void Lot::clear_auctioneer() {
  auctioneer_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& Lot::auctioneer() const {
  // @@protoc_insertion_point(field_get:auction.Lot.auctioneer)
  return auctioneer_.GetNoArena();
}
inline void Lot::set_auctioneer(const ::std::string& value) {
  
  auctioneer_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:auction.Lot.auctioneer)
}
#if LANG_CXX11
inline void Lot::set_auctioneer(::std::string&& value) {
  
  auctioneer_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:auction.Lot.auctioneer)
}
#endif
inline void Lot::set_auctioneer(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  
  auctioneer_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:auction.Lot.auctioneer)
}
inline void Lot::set_auctioneer(const char* value, size_t size) {
  
  auctioneer_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:auction.Lot.auctioneer)
}
inline ::std::string* Lot::mutable_auctioneer() {
  
  // @@protoc_insertion_point(field_mutable:auction.Lot.auctioneer)
  return auctioneer_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* Lot::release_auctioneer() {
  // @@protoc_insertion_point(field_release:auction.Lot.auctioneer)
  
  return auctioneer_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void Lot::set_allocated_auctioneer(::std::string* auctioneer) {
  if (auctioneer != nullptr) {
    
  } else {
    
  }
  auctioneer_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), auctioneer);
  // @@protoc_insertion_point(field_set_allocated:auction.Lot.auctioneer)
}

// uint64 floor = 3;
inline void Lot::clear_floor() {
  floor_ = PROTOBUF_ULONGLONG(0);
}
inline ::google::protobuf::uint64 Lot::floor() const {
  // @@protoc_insertion_point(field_get:auction.Lot.floor)
  return floor_;
}
inline void Lot::set_floor(::google::protobuf::uint64 value) {
  
  floor_ = value;
  // @@protoc_insertion_point(field_set:auction.Lot.floor)
}

// string data = 4;
inline void Lot::clear_data() {
  data_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& Lot::data() const {
  // @@protoc_insertion_point(field_get:auction.Lot.data)
  return data_.GetNoArena();
}
inline void Lot::set_data(const ::std::string& value) {
  
  data_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:auction.Lot.data)
}
#if LANG_CXX11
inline void Lot::set_data(::std::string&& value) {
  
  data_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:auction.Lot.data)
}
#endif
inline void Lot::set_data(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  
  data_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:auction.Lot.data)
}
inline void Lot::set_data(const char* value, size_t size) {
  
  data_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:auction.Lot.data)
}
inline ::std::string* Lot::mutable_data() {
  
  // @@protoc_insertion_point(field_mutable:auction.Lot.data)
  return data_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* Lot::release_data() {
  // @@protoc_insertion_point(field_release:auction.Lot.data)
  
  return data_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void Lot::set_allocated_data(::std::string* data) {
  if (data != nullptr) {
    
  } else {
    
  }
  data_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), data);
  // @@protoc_insertion_point(field_set_allocated:auction.Lot.data)
}

// string bidder = 5;
inline void Lot::clear_bidder() {
  bidder_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& Lot::bidder() const {
  // @@protoc_insertion_point(field_get:auction.Lot.bidder)
  return bidder_.GetNoArena();
}
inline void Lot::set_bidder(const ::std::string& value) {
  
  bidder_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:auction.Lot.bidder)
}
#if LANG_CXX11
inline void Lot::set_bidder(::std::string&& value) {
  
  bidder_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:auction.Lot.bidder)
}
#endif
inline void Lot::set_bidder(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  
  bidder_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:auction.Lot.bidder)
}
inline void Lot::set_bidder(const char* value, size_t size) {
  
  bidder_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:auction.Lot.bidder)
}
inline ::std::string* Lot::mutable_bidder() {
  
  // @@protoc_insertion_point(field_mutable:auction.Lot.bidder)
  return bidder_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* Lot::release_bidder() {
  // @@protoc_insertion_point(field_release:auction.Lot.bidder)
  
  return bidder_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void Lot::set_allocated_bidder(::std::string* bidder) {
  if (bidder != nullptr) {
    
  } else {
    
  }
  bidder_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), bidder);
  // @@protoc_insertion_point(field_set_allocated:auction.Lot.bidder)
}

// uint64 bid = 6;
inline void Lot::clear_bid() {
  bid_ = PROTOBUF_ULONGLONG(0);
}
inline ::google::protobuf::uint64 Lot::bid() const {
  // @@protoc_insertion_point(field_get:auction.Lot.bid)
  return bid_;
}
inline void Lot::set_bid(::google::protobuf::uint64 value) {
  
  bid_ = value;
  // @@protoc_insertion_point(field_set:auction.Lot.bid)
}

// .auction.Status status = 7;
inline void Lot::clear_status() {
  status_ = 0;
}
inline ::auction::Status Lot::status() const {
  // @@protoc_insertion_point(field_get:auction.Lot.status)
  return static_cast< ::auction::Status >(status_);
}
inline void Lot::set_status(::auction::Status value) {
  
  status_ = value;
  // @@protoc_insertion_point(field_set:auction.Lot.status)
}

// -------------------------------------------------------------------

// Chip

// string bidder = 2;
inline void Chip::clear_bidder() {
  bidder_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& Chip::bidder() const {
  // @@protoc_insertion_point(field_get:auction.Chip.bidder)
  return bidder_.GetNoArena();
}
inline void Chip::set_bidder(const ::std::string& value) {
  
  bidder_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:auction.Chip.bidder)
}
#if LANG_CXX11
inline void Chip::set_bidder(::std::string&& value) {
  
  bidder_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:auction.Chip.bidder)
}
#endif
inline void Chip::set_bidder(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  
  bidder_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:auction.Chip.bidder)
}
inline void Chip::set_bidder(const char* value, size_t size) {
  
  bidder_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:auction.Chip.bidder)
}
inline ::std::string* Chip::mutable_bidder() {
  
  // @@protoc_insertion_point(field_mutable:auction.Chip.bidder)
  return bidder_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* Chip::release_bidder() {
  // @@protoc_insertion_point(field_release:auction.Chip.bidder)
  
  return bidder_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void Chip::set_allocated_bidder(::std::string* bidder) {
  if (bidder != nullptr) {
    
  } else {
    
  }
  bidder_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), bidder);
  // @@protoc_insertion_point(field_set_allocated:auction.Chip.bidder)
}

// uint64 amount = 3;
inline void Chip::clear_amount() {
  amount_ = PROTOBUF_ULONGLONG(0);
}
inline ::google::protobuf::uint64 Chip::amount() const {
  // @@protoc_insertion_point(field_get:auction.Chip.amount)
  return amount_;
}
inline void Chip::set_amount(::google::protobuf::uint64 value) {
  
  amount_ = value;
  // @@protoc_insertion_point(field_set:auction.Chip.amount)
}

// uint64 bid = 4;
inline void Chip::clear_bid() {
  bid_ = PROTOBUF_ULONGLONG(0);
}
inline ::google::protobuf::uint64 Chip::bid() const {
  // @@protoc_insertion_point(field_get:auction.Chip.bid)
  return bid_;
}
inline void Chip::set_bid(::google::protobuf::uint64 value) {
  
  bid_ = value;
  // @@protoc_insertion_point(field_set:auction.Chip.bid)
}

// .auction.Status status = 8;
inline void Chip::clear_status() {
  status_ = 0;
}
inline ::auction::Status Chip::status() const {
  // @@protoc_insertion_point(field_get:auction.Chip.status)
  return static_cast< ::auction::Status >(status_);
}
inline void Chip::set_status(::auction::Status value) {
  
  status_ = value;
  // @@protoc_insertion_point(field_set:auction.Chip.status)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace auction

namespace google {
namespace protobuf {

template <> struct is_proto_enum< ::auction::Status> : ::std::true_type {};

}  // namespace protobuf
}  // namespace google

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // PROTOBUF_INCLUDED_auction_2eproto
