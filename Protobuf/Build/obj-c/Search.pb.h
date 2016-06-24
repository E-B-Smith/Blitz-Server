// Generated by the protocol buffer compiler.  DO NOT EDIT!

#import <ProtocolBuffers/ProtocolBuffers.h>

#import "UserProfiles.pb.h"
// @@protoc_insertion_point(imports)

@class BAutocompleteRequest;
@class BAutocompleteRequestBuilder;
@class BAutocompleteResponse;
@class BAutocompleteResponseBuilder;
@class BConfirmationRequest;
@class BConfirmationRequestBuilder;
@class BContactInfo;
@class BContactInfoBuilder;
@class BCoordinate;
@class BCoordinateBuilder;
@class BCoordinatePolygon;
@class BCoordinatePolygonBuilder;
@class BCoordinateRegion;
@class BCoordinateRegionBuilder;
@class BEditProfile;
@class BEditProfileBuilder;
@class BEducation;
@class BEducationBuilder;
@class BEmployment;
@class BEmploymentBuilder;
@class BEntityTag;
@class BEntityTagBuilder;
@class BEntityTagList;
@class BEntityTagListBuilder;
@class BFriendUpdate;
@class BFriendUpdateBuilder;
@class BImageData;
@class BImageDataBuilder;
@class BImageUpload;
@class BImageUploadBuilder;
@class BKeyValue;
@class BKeyValueBuilder;
@class BLocation;
@class BLocationBuilder;
@class BPoint;
@class BPointBuilder;
@class BProfilesFromContactInfo;
@class BProfilesFromContactInfoBuilder;
@class BSearchCategories;
@class BSearchCategoriesBuilder;
@class BSearchCategory;
@class BSearchCategoryBuilder;
@class BSize;
@class BSizeBuilder;
@class BSocialIdentity;
@class BSocialIdentityBuilder;
@class BTimespan;
@class BTimespanBuilder;
@class BTimestamp;
@class BTimestampBuilder;
@class BUserInvite;
@class BUserInviteBuilder;
@class BUserProfile;
@class BUserProfileBuilder;
@class BUserProfileQuery;
@class BUserProfileQueryBuilder;
@class BUserProfileUpdate;
@class BUserProfileUpdateBuilder;
@class BUserReview;
@class BUserReviewBuilder;
@class BUserSearchRequest;
@class BUserSearchRequestBuilder;
@class BUserSearchResponse;
@class BUserSearchResponseBuilder;
@class BVoid;
@class BVoidBuilder;
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
#define AutocompleteRequest_searchType @"searchType"
@interface BAutocompleteRequest : PBGeneratedMessage<GeneratedMessageProtocol> {
@private
  BOOL hasQuery_:1;
  BOOL hasSearchType_:1;
  NSString* query;
  BSearchType searchType;
}
- (BOOL) hasQuery;
- (BOOL) hasSearchType;
@property (readonly, strong) NSString* query;
@property (readonly) BSearchType searchType;

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

- (BOOL) hasSearchType;
- (BSearchType) searchType;
- (BAutocompleteRequestBuilder*) setSearchType:(BSearchType) value;
- (BAutocompleteRequestBuilder*) clearSearchType;
@end

#define AutocompleteResponse_query @"query"
#define AutocompleteResponse_suggestions @"suggestions"
@interface BAutocompleteResponse : PBGeneratedMessage<GeneratedMessageProtocol> {
@private
  BOOL hasQuery_:1;
  NSString* query;
  NSMutableArray * suggestionsArray;
}
- (BOOL) hasQuery;
@property (readonly, strong) NSString* query;
@property (readonly, strong) NSArray * suggestions;
- (NSString*)suggestionsAtIndex:(NSUInteger)index;

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

- (BOOL) hasQuery;
- (NSString*) query;
- (BAutocompleteResponseBuilder*) setQuery:(NSString*) value;
- (BAutocompleteResponseBuilder*) clearQuery;

- (NSMutableArray *)suggestions;
- (NSString*)suggestionsAtIndex:(NSUInteger)index;
- (BAutocompleteResponseBuilder *)addSuggestions:(NSString*)value;
- (BAutocompleteResponseBuilder *)setSuggestionsArray:(NSArray *)array;
- (BAutocompleteResponseBuilder *)clearSuggestions;
@end

#define UserSearchRequest_query @"query"
@interface BUserSearchRequest : PBGeneratedMessage<GeneratedMessageProtocol> {
@private
  BOOL hasQuery_:1;
  NSString* query;
}
- (BOOL) hasQuery;
@property (readonly, strong) NSString* query;

+ (instancetype) defaultInstance;
- (instancetype) defaultInstance;

- (BOOL) isInitialized;
- (void) writeToCodedOutputStream:(PBCodedOutputStream*) output;
- (BUserSearchRequestBuilder*) builder;
+ (BUserSearchRequestBuilder*) builder;
+ (BUserSearchRequestBuilder*) builderWithPrototype:(BUserSearchRequest*) prototype;
- (BUserSearchRequestBuilder*) toBuilder;

+ (BUserSearchRequest*) parseFromData:(NSData*) data;
+ (BUserSearchRequest*) parseFromData:(NSData*) data extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (BUserSearchRequest*) parseFromInputStream:(NSInputStream*) input;
+ (BUserSearchRequest*) parseFromInputStream:(NSInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (BUserSearchRequest*) parseFromCodedInputStream:(PBCodedInputStream*) input;
+ (BUserSearchRequest*) parseFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
@end

@interface BUserSearchRequestBuilder : PBGeneratedMessageBuilder {
@private
  BUserSearchRequest* resultUserSearchRequest;
}

- (BUserSearchRequest*) defaultInstance;

- (BUserSearchRequestBuilder*) clear;
- (BUserSearchRequestBuilder*) clone;

- (BUserSearchRequest*) build;
- (BUserSearchRequest*) buildPartial;

- (BUserSearchRequestBuilder*) mergeFrom:(BUserSearchRequest*) other;
- (BUserSearchRequestBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input;
- (BUserSearchRequestBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;

- (BOOL) hasQuery;
- (NSString*) query;
- (BUserSearchRequestBuilder*) setQuery:(NSString*) value;
- (BUserSearchRequestBuilder*) clearQuery;
@end

#define UserSearchResponse_query @"query"
#define UserSearchResponse_profiles @"profiles"
@interface BUserSearchResponse : PBGeneratedMessage<GeneratedMessageProtocol> {
@private
  BOOL hasQuery_:1;
  NSString* query;
  NSMutableArray * profilesArray;
}
- (BOOL) hasQuery;
@property (readonly, strong) NSString* query;
@property (readonly, strong) NSArray * profiles;
- (BUserProfile*)profilesAtIndex:(NSUInteger)index;

+ (instancetype) defaultInstance;
- (instancetype) defaultInstance;

- (BOOL) isInitialized;
- (void) writeToCodedOutputStream:(PBCodedOutputStream*) output;
- (BUserSearchResponseBuilder*) builder;
+ (BUserSearchResponseBuilder*) builder;
+ (BUserSearchResponseBuilder*) builderWithPrototype:(BUserSearchResponse*) prototype;
- (BUserSearchResponseBuilder*) toBuilder;

+ (BUserSearchResponse*) parseFromData:(NSData*) data;
+ (BUserSearchResponse*) parseFromData:(NSData*) data extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (BUserSearchResponse*) parseFromInputStream:(NSInputStream*) input;
+ (BUserSearchResponse*) parseFromInputStream:(NSInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (BUserSearchResponse*) parseFromCodedInputStream:(PBCodedInputStream*) input;
+ (BUserSearchResponse*) parseFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
@end

@interface BUserSearchResponseBuilder : PBGeneratedMessageBuilder {
@private
  BUserSearchResponse* resultUserSearchResponse;
}

- (BUserSearchResponse*) defaultInstance;

- (BUserSearchResponseBuilder*) clear;
- (BUserSearchResponseBuilder*) clone;

- (BUserSearchResponse*) build;
- (BUserSearchResponse*) buildPartial;

- (BUserSearchResponseBuilder*) mergeFrom:(BUserSearchResponse*) other;
- (BUserSearchResponseBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input;
- (BUserSearchResponseBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;

- (BOOL) hasQuery;
- (NSString*) query;
- (BUserSearchResponseBuilder*) setQuery:(NSString*) value;
- (BUserSearchResponseBuilder*) clearQuery;

- (NSMutableArray *)profiles;
- (BUserProfile*)profilesAtIndex:(NSUInteger)index;
- (BUserSearchResponseBuilder *)addProfiles:(BUserProfile*)value;
- (BUserSearchResponseBuilder *)setProfilesArray:(NSArray *)array;
- (BUserSearchResponseBuilder *)clearProfiles;
@end

#define SearchCategory_item @"item"
#define SearchCategory_parent @"parent"
#define SearchCategory_isLeaf @"isLeaf"
#define SearchCategory_descriptionText @"descriptionText"
@interface BSearchCategory : PBGeneratedMessage<GeneratedMessageProtocol> {
@private
  BOOL hasIsLeaf_:1;
  BOOL hasItem_:1;
  BOOL hasParent_:1;
  BOOL hasDescriptionText_:1;
  BOOL isLeaf_:1;
  NSString* item;
  NSString* parent;
  NSString* descriptionText;
}
- (BOOL) hasItem;
- (BOOL) hasParent;
- (BOOL) hasIsLeaf;
- (BOOL) hasDescriptionText;
@property (readonly, strong) NSString* item;
@property (readonly, strong) NSString* parent;
- (BOOL) isLeaf;
@property (readonly, strong) NSString* descriptionText;

+ (instancetype) defaultInstance;
- (instancetype) defaultInstance;

- (BOOL) isInitialized;
- (void) writeToCodedOutputStream:(PBCodedOutputStream*) output;
- (BSearchCategoryBuilder*) builder;
+ (BSearchCategoryBuilder*) builder;
+ (BSearchCategoryBuilder*) builderWithPrototype:(BSearchCategory*) prototype;
- (BSearchCategoryBuilder*) toBuilder;

+ (BSearchCategory*) parseFromData:(NSData*) data;
+ (BSearchCategory*) parseFromData:(NSData*) data extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (BSearchCategory*) parseFromInputStream:(NSInputStream*) input;
+ (BSearchCategory*) parseFromInputStream:(NSInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (BSearchCategory*) parseFromCodedInputStream:(PBCodedInputStream*) input;
+ (BSearchCategory*) parseFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
@end

@interface BSearchCategoryBuilder : PBGeneratedMessageBuilder {
@private
  BSearchCategory* resultSearchCategory;
}

- (BSearchCategory*) defaultInstance;

- (BSearchCategoryBuilder*) clear;
- (BSearchCategoryBuilder*) clone;

- (BSearchCategory*) build;
- (BSearchCategory*) buildPartial;

- (BSearchCategoryBuilder*) mergeFrom:(BSearchCategory*) other;
- (BSearchCategoryBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input;
- (BSearchCategoryBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;

- (BOOL) hasItem;
- (NSString*) item;
- (BSearchCategoryBuilder*) setItem:(NSString*) value;
- (BSearchCategoryBuilder*) clearItem;

- (BOOL) hasParent;
- (NSString*) parent;
- (BSearchCategoryBuilder*) setParent:(NSString*) value;
- (BSearchCategoryBuilder*) clearParent;

- (BOOL) hasIsLeaf;
- (BOOL) isLeaf;
- (BSearchCategoryBuilder*) setIsLeaf:(BOOL) value;
- (BSearchCategoryBuilder*) clearIsLeaf;

- (BOOL) hasDescriptionText;
- (NSString*) descriptionText;
- (BSearchCategoryBuilder*) setDescriptionText:(NSString*) value;
- (BSearchCategoryBuilder*) clearDescriptionText;
@end

#define SearchCategories_categories @"categories"
@interface BSearchCategories : PBGeneratedMessage<GeneratedMessageProtocol> {
@private
  NSMutableArray * categoriesArray;
}
@property (readonly, strong) NSArray * categories;
- (BSearchCategory*)categoriesAtIndex:(NSUInteger)index;

+ (instancetype) defaultInstance;
- (instancetype) defaultInstance;

- (BOOL) isInitialized;
- (void) writeToCodedOutputStream:(PBCodedOutputStream*) output;
- (BSearchCategoriesBuilder*) builder;
+ (BSearchCategoriesBuilder*) builder;
+ (BSearchCategoriesBuilder*) builderWithPrototype:(BSearchCategories*) prototype;
- (BSearchCategoriesBuilder*) toBuilder;

+ (BSearchCategories*) parseFromData:(NSData*) data;
+ (BSearchCategories*) parseFromData:(NSData*) data extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (BSearchCategories*) parseFromInputStream:(NSInputStream*) input;
+ (BSearchCategories*) parseFromInputStream:(NSInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (BSearchCategories*) parseFromCodedInputStream:(PBCodedInputStream*) input;
+ (BSearchCategories*) parseFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
@end

@interface BSearchCategoriesBuilder : PBGeneratedMessageBuilder {
@private
  BSearchCategories* resultSearchCategories;
}

- (BSearchCategories*) defaultInstance;

- (BSearchCategoriesBuilder*) clear;
- (BSearchCategoriesBuilder*) clone;

- (BSearchCategories*) build;
- (BSearchCategories*) buildPartial;

- (BSearchCategoriesBuilder*) mergeFrom:(BSearchCategories*) other;
- (BSearchCategoriesBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input;
- (BSearchCategoriesBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;

- (NSMutableArray *)categories;
- (BSearchCategory*)categoriesAtIndex:(NSUInteger)index;
- (BSearchCategoriesBuilder *)addCategories:(BSearchCategory*)value;
- (BSearchCategoriesBuilder *)setCategoriesArray:(NSArray *)array;
- (BSearchCategoriesBuilder *)clearCategories;
@end


// @@protoc_insertion_point(global_scope)
