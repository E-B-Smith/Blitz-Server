// Generated by the protocol buffer compiler.  DO NOT EDIT!

#import <ProtocolBuffers/ProtocolBuffers.h>

// @@protoc_insertion_point(imports)

@class BAutocompleteRequest;
@class BAutocompleteRequestBuilder;
@class BAutocompleteResponse;
@class BAutocompleteResponseBuilder;
@class ObjectiveCFileOptions;
@class ObjectiveCFileOptionsBuilder;
@class PBDescriptorProto;
@class PBDescriptorProtoBuilder;
@class PBDescriptorProtoExtensionRange;
@class PBDescriptorProtoExtensionRangeBuilder;
@class PBDescriptorProtoReservedRange;
@class PBDescriptorProtoReservedRangeBuilder;
@class PBEnumDescriptorProto;
@class PBEnumDescriptorProtoBuilder;
@class PBEnumOptions;
@class PBEnumOptionsBuilder;
@class PBEnumValueDescriptorProto;
@class PBEnumValueDescriptorProtoBuilder;
@class PBEnumValueOptions;
@class PBEnumValueOptionsBuilder;
@class PBFieldDescriptorProto;
@class PBFieldDescriptorProtoBuilder;
@class PBFieldOptions;
@class PBFieldOptionsBuilder;
@class PBFileDescriptorProto;
@class PBFileDescriptorProtoBuilder;
@class PBFileDescriptorSet;
@class PBFileDescriptorSetBuilder;
@class PBFileOptions;
@class PBFileOptionsBuilder;
@class PBMessageOptions;
@class PBMessageOptionsBuilder;
@class PBMethodDescriptorProto;
@class PBMethodDescriptorProtoBuilder;
@class PBMethodOptions;
@class PBMethodOptionsBuilder;
@class PBOneofDescriptorProto;
@class PBOneofDescriptorProtoBuilder;
@class PBServiceDescriptorProto;
@class PBServiceDescriptorProtoBuilder;
@class PBServiceOptions;
@class PBServiceOptionsBuilder;
@class PBSourceCodeInfo;
@class PBSourceCodeInfoBuilder;
@class PBSourceCodeInfoLocation;
@class PBSourceCodeInfoLocationBuilder;
@class PBUninterpretedOption;
@class PBUninterpretedOptionBuilder;
@class PBUninterpretedOptionNamePart;
@class PBUninterpretedOptionNamePartBuilder;


typedef NS_ENUM(SInt32, BSearchType) {
  BSearchTypeSTSearchAll = 0,
  BSearchTypeSTUsers = 1,
  BSearchTypeSTTopics = 2,
};

BOOL BSearchTypeIsValidValue(BSearchType value);
NSString *NSStringFromBSearchType(BSearchType value);


@interface BSearchRoot : NSObject {
}
+ (PBExtensionRegistry*) extensionRegistry;
+ (void) registerAllExtensions:(PBMutableExtensionRegistry*) registry;
@end

#define AutocompleteRequest_query @"query"
#define AutocompleteRequest_type @"type"
@interface BAutocompleteRequest : PBGeneratedMessage<GeneratedMessageProtocol> {
@private
  BOOL hasQuery_:1;
  BOOL hasType_:1;
  NSString* query;
  BSearchType type;
}
- (BOOL) hasQuery;
- (BOOL) hasType;
@property (readonly, strong) NSString* query;
@property (readonly) BSearchType type;

+ (instancetype) defaultInstance;
- (instancetype) defaultInstance;

- (BOOL) isInitialized;
- (void) writeToCodedOutputStream:(PBCodedOutputStream*) output;
- (BAutocompleteRequestBuilder*) builder;
+ (BAutocompleteRequestBuilder*) builder;
+ (BAutocompleteRequestBuilder*) builderWithPrototype:(BAutocompleteRequest*) prototype;
- (BAutocompleteRequestBuilder*) toBuilder;

+ (BAutocompleteRequest*) parseFromData:(NSData*) data;
+ (BAutocompleteRequest*) parseFromData:(NSData*) data extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (BAutocompleteRequest*) parseFromInputStream:(NSInputStream*) input;
+ (BAutocompleteRequest*) parseFromInputStream:(NSInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (BAutocompleteRequest*) parseFromCodedInputStream:(PBCodedInputStream*) input;
+ (BAutocompleteRequest*) parseFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
@end

@interface BAutocompleteRequestBuilder : PBGeneratedMessageBuilder {
@private
  BAutocompleteRequest* resultAutocompleteRequest;
}

- (BAutocompleteRequest*) defaultInstance;

- (BAutocompleteRequestBuilder*) clear;
- (BAutocompleteRequestBuilder*) clone;

- (BAutocompleteRequest*) build;
- (BAutocompleteRequest*) buildPartial;

- (BAutocompleteRequestBuilder*) mergeFrom:(BAutocompleteRequest*) other;
- (BAutocompleteRequestBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input;
- (BAutocompleteRequestBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;

- (BOOL) hasQuery;
- (NSString*) query;
- (BAutocompleteRequestBuilder*) setQuery:(NSString*) value;
- (BAutocompleteRequestBuilder*) clearQuery;

- (BOOL) hasType;
- (BSearchType) type;
- (BAutocompleteRequestBuilder*) setType:(BSearchType) value;
- (BAutocompleteRequestBuilder*) clearType;
@end

#define AutocompleteResponse_items @"items"
@interface BAutocompleteResponse : PBGeneratedMessage<GeneratedMessageProtocol> {
@private
  NSMutableArray * itemsArray;
}
@property (readonly, strong) NSArray * items;
- (NSString*)itemsAtIndex:(NSUInteger)index;

+ (instancetype) defaultInstance;
- (instancetype) defaultInstance;

- (BOOL) isInitialized;
- (void) writeToCodedOutputStream:(PBCodedOutputStream*) output;
- (BAutocompleteResponseBuilder*) builder;
+ (BAutocompleteResponseBuilder*) builder;
+ (BAutocompleteResponseBuilder*) builderWithPrototype:(BAutocompleteResponse*) prototype;
- (BAutocompleteResponseBuilder*) toBuilder;

+ (BAutocompleteResponse*) parseFromData:(NSData*) data;
+ (BAutocompleteResponse*) parseFromData:(NSData*) data extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (BAutocompleteResponse*) parseFromInputStream:(NSInputStream*) input;
+ (BAutocompleteResponse*) parseFromInputStream:(NSInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (BAutocompleteResponse*) parseFromCodedInputStream:(PBCodedInputStream*) input;
+ (BAutocompleteResponse*) parseFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
@end

@interface BAutocompleteResponseBuilder : PBGeneratedMessageBuilder {
@private
  BAutocompleteResponse* resultAutocompleteResponse;
}

- (BAutocompleteResponse*) defaultInstance;

- (BAutocompleteResponseBuilder*) clear;
- (BAutocompleteResponseBuilder*) clone;

- (BAutocompleteResponse*) build;
- (BAutocompleteResponse*) buildPartial;

- (BAutocompleteResponseBuilder*) mergeFrom:(BAutocompleteResponse*) other;
- (BAutocompleteResponseBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input;
- (BAutocompleteResponseBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;

- (NSMutableArray *)items;
- (NSString*)itemsAtIndex:(NSUInteger)index;
- (BAutocompleteResponseBuilder *)addItems:(NSString*)value;
- (BAutocompleteResponseBuilder *)setItemsArray:(NSArray *)array;
- (BAutocompleteResponseBuilder *)clearItems;
@end


// @@protoc_insertion_point(global_scope)
