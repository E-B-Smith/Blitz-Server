// Generated by the protocol buffer compiler.  DO NOT EDIT!

#import <ProtocolBuffers/ProtocolBuffers.h>

#import "Types.pb.h"
#import "EntityTags.pb.h"
// @@protoc_insertion_point(imports)

@class BCoordinate;
@class BCoordinateBuilder;
@class BCoordinatePolygon;
@class BCoordinatePolygonBuilder;
@class BCoordinateRegion;
@class BCoordinateRegionBuilder;
@class BEntityTag;
@class BEntityTagBuilder;
@class BEntityTagList;
@class BEntityTagListBuilder;
@class BFeedPost;
@class BFeedPostBuilder;
@class BFeedPostFetchRequest;
@class BFeedPostFetchRequestBuilder;
@class BFeedPostFetchResponse;
@class BFeedPostFetchResponseBuilder;
@class BFeedPostUpdateRequest;
@class BFeedPostUpdateRequestBuilder;
@class BFeedPostUpdateResponse;
@class BFeedPostUpdateResponseBuilder;
@class BGlobals;
@class BGlobalsBuilder;
@class BKeyValue;
@class BKeyValueBuilder;
@class BLocation;
@class BLocationBuilder;
@class BSize;
@class BSizeBuilder;
@class BTimespan;
@class BTimespanBuilder;
@class BTimestamp;
@class BTimestampBuilder;
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


typedef NS_ENUM(SInt32, BFeedPostType) {
  BFeedPostTypeFPUnknown = 0,
  BFeedPostTypeFPOpenEndedQuestion = 1,
  BFeedPostTypeFPOpenEndedReply = 2,
  BFeedPostTypeFPSurveyQuestion = 3,
  BFeedPostTypeFPSurveyAnswer = 4,
};

BOOL BFeedPostTypeIsValidValue(BFeedPostType value);
NSString *NSStringFromBFeedPostType(BFeedPostType value);

typedef NS_ENUM(SInt32, BFeedPostScope) {
  BFeedPostScopeFPScopeUnknown = 0,
  BFeedPostScopeFPScopeLocalNetwork = 1,
  BFeedPostScopeFPScopeGlobalNetwork = 2,
};

BOOL BFeedPostScopeIsValidValue(BFeedPostScope value);
NSString *NSStringFromBFeedPostScope(BFeedPostScope value);

typedef NS_ENUM(SInt32, BFeedPostStatus) {
  BFeedPostStatusFPSUnknown = 0,
  BFeedPostStatusFPSActive = 1,
  BFeedPostStatusFPSDeleted = 2,
};

BOOL BFeedPostStatusIsValidValue(BFeedPostStatus value);
NSString *NSStringFromBFeedPostStatus(BFeedPostStatus value);

typedef NS_ENUM(SInt32, BUpdateVerb) {
  BUpdateVerbUVCreate = 1,
  BUpdateVerbUVUpdate = 2,
  BUpdateVerbUVDelete = 3,
};

BOOL BUpdateVerbIsValidValue(BUpdateVerb value);
NSString *NSStringFromBUpdateVerb(BUpdateVerb value);


@interface BFeedRoot : NSObject {
}
+ (PBExtensionRegistry*) extensionRegistry;
+ (void) registerAllExtensions:(PBMutableExtensionRegistry*) registry;
@end

#define FeedPost_postID @"postID"
#define FeedPost_parentID @"parentID"
#define FeedPost_postType @"postType"
#define FeedPost_postScope @"postScope"
#define FeedPost_userID @"userID"
#define FeedPost_anonymousPost @"anonymousPost"
#define FeedPost_timestamp @"timestamp"
#define FeedPost_timespanActive @"timespanActive"
#define FeedPost_headlineText @"headlineText"
#define FeedPost_bodyText @"bodyText"
#define FeedPost_postTags @"postTags"
#define FeedPost_replies_deprecated @"repliesDeprecated"
#define FeedPost_mayAddReply @"mayAddReply"
#define FeedPost_mayChooseMulitpleReplies @"mayChooseMulitpleReplies"
#define FeedPost_surveyAnswerSequence @"surveyAnswerSequence"
#define FeedPost_areMoreReplies @"areMoreReplies"
#define FeedPost_totalVoteCount @"totalVoteCount"
@interface BFeedPost : PBGeneratedMessage<GeneratedMessageProtocol> {
@private
  BOOL hasAnonymousPost_:1;
  BOOL hasMayAddReply_:1;
  BOOL hasMayChooseMulitpleReplies_:1;
  BOOL hasAreMoreReplies_:1;
  BOOL hasSurveyAnswerSequence_:1;
  BOOL hasTotalVoteCount_:1;
  BOOL hasPostID_:1;
  BOOL hasParentID_:1;
  BOOL hasUserID_:1;
  BOOL hasHeadlineText_:1;
  BOOL hasBodyText_:1;
  BOOL hasTimestamp_:1;
  BOOL hasTimespanActive_:1;
  BOOL hasPostType_:1;
  BOOL hasPostScope_:1;
  BOOL anonymousPost_:1;
  BOOL mayAddReply_:1;
  BOOL mayChooseMulitpleReplies_:1;
  BOOL areMoreReplies_:1;
  SInt32 surveyAnswerSequence;
  SInt32 totalVoteCount;
  NSString* postID;
  NSString* parentID;
  NSString* userID;
  NSString* headlineText;
  NSString* bodyText;
  BTimestamp* timestamp;
  BTimespan* timespanActive;
  BFeedPostType postType;
  BFeedPostScope postScope;
  NSMutableArray * postTagsArray;
  NSMutableArray * repliesDeprecatedArray;
}
- (BOOL) hasPostID;
- (BOOL) hasParentID;
- (BOOL) hasPostType;
- (BOOL) hasPostScope;
- (BOOL) hasUserID;
- (BOOL) hasAnonymousPost;
- (BOOL) hasTimestamp;
- (BOOL) hasTimespanActive;
- (BOOL) hasHeadlineText;
- (BOOL) hasBodyText;
- (BOOL) hasMayAddReply;
- (BOOL) hasMayChooseMulitpleReplies;
- (BOOL) hasSurveyAnswerSequence;
- (BOOL) hasAreMoreReplies;
- (BOOL) hasTotalVoteCount;
@property (readonly, strong) NSString* postID;
@property (readonly, strong) NSString* parentID;
@property (readonly) BFeedPostType postType;
@property (readonly) BFeedPostScope postScope;
@property (readonly, strong) NSString* userID;
- (BOOL) anonymousPost;
@property (readonly, strong) BTimestamp* timestamp;
@property (readonly, strong) BTimespan* timespanActive;
@property (readonly, strong) NSString* headlineText;
@property (readonly, strong) NSString* bodyText;
@property (readonly, strong) NSArray * postTags;
@property (readonly, strong) NSArray * repliesDeprecated;
- (BOOL) mayAddReply;
- (BOOL) mayChooseMulitpleReplies;
@property (readonly) SInt32 surveyAnswerSequence;
- (BOOL) areMoreReplies;
@property (readonly) SInt32 totalVoteCount;
- (BEntityTag*)postTagsAtIndex:(NSUInteger)index;
- (BFeedPost*)repliesDeprecatedAtIndex:(NSUInteger)index;

+ (instancetype) defaultInstance;
- (instancetype) defaultInstance;

- (BOOL) isInitialized;
- (void) writeToCodedOutputStream:(PBCodedOutputStream*) output;
- (BFeedPostBuilder*) builder;
+ (BFeedPostBuilder*) builder;
+ (BFeedPostBuilder*) builderWithPrototype:(BFeedPost*) prototype;
- (BFeedPostBuilder*) toBuilder;

+ (BFeedPost*) parseFromData:(NSData*) data;
+ (BFeedPost*) parseFromData:(NSData*) data extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (BFeedPost*) parseFromInputStream:(NSInputStream*) input;
+ (BFeedPost*) parseFromInputStream:(NSInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (BFeedPost*) parseFromCodedInputStream:(PBCodedInputStream*) input;
+ (BFeedPost*) parseFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
@end

@interface BFeedPostBuilder : PBGeneratedMessageBuilder {
@private
  BFeedPost* resultFeedPost;
}

- (BFeedPost*) defaultInstance;

- (BFeedPostBuilder*) clear;
- (BFeedPostBuilder*) clone;

- (BFeedPost*) build;
- (BFeedPost*) buildPartial;

- (BFeedPostBuilder*) mergeFrom:(BFeedPost*) other;
- (BFeedPostBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input;
- (BFeedPostBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;

- (BOOL) hasPostID;
- (NSString*) postID;
- (BFeedPostBuilder*) setPostID:(NSString*) value;
- (BFeedPostBuilder*) clearPostID;

- (BOOL) hasParentID;
- (NSString*) parentID;
- (BFeedPostBuilder*) setParentID:(NSString*) value;
- (BFeedPostBuilder*) clearParentID;

- (BOOL) hasPostType;
- (BFeedPostType) postType;
- (BFeedPostBuilder*) setPostType:(BFeedPostType) value;
- (BFeedPostBuilder*) clearPostType;

- (BOOL) hasPostScope;
- (BFeedPostScope) postScope;
- (BFeedPostBuilder*) setPostScope:(BFeedPostScope) value;
- (BFeedPostBuilder*) clearPostScope;

- (BOOL) hasUserID;
- (NSString*) userID;
- (BFeedPostBuilder*) setUserID:(NSString*) value;
- (BFeedPostBuilder*) clearUserID;

- (BOOL) hasAnonymousPost;
- (BOOL) anonymousPost;
- (BFeedPostBuilder*) setAnonymousPost:(BOOL) value;
- (BFeedPostBuilder*) clearAnonymousPost;

- (BOOL) hasTimestamp;
- (BTimestamp*) timestamp;
- (BFeedPostBuilder*) setTimestamp:(BTimestamp*) value;
- (BFeedPostBuilder*) setTimestampBuilder:(BTimestampBuilder*) builderForValue;
- (BFeedPostBuilder*) mergeTimestamp:(BTimestamp*) value;
- (BFeedPostBuilder*) clearTimestamp;

- (BOOL) hasTimespanActive;
- (BTimespan*) timespanActive;
- (BFeedPostBuilder*) setTimespanActive:(BTimespan*) value;
- (BFeedPostBuilder*) setTimespanActiveBuilder:(BTimespanBuilder*) builderForValue;
- (BFeedPostBuilder*) mergeTimespanActive:(BTimespan*) value;
- (BFeedPostBuilder*) clearTimespanActive;

- (BOOL) hasHeadlineText;
- (NSString*) headlineText;
- (BFeedPostBuilder*) setHeadlineText:(NSString*) value;
- (BFeedPostBuilder*) clearHeadlineText;

- (BOOL) hasBodyText;
- (NSString*) bodyText;
- (BFeedPostBuilder*) setBodyText:(NSString*) value;
- (BFeedPostBuilder*) clearBodyText;

- (NSMutableArray *)postTags;
- (BEntityTag*)postTagsAtIndex:(NSUInteger)index;
- (BFeedPostBuilder *)addPostTags:(BEntityTag*)value;
- (BFeedPostBuilder *)setPostTagsArray:(NSArray *)array;
- (BFeedPostBuilder *)clearPostTags;

- (NSMutableArray *)repliesDeprecated;
- (BFeedPost*)repliesDeprecatedAtIndex:(NSUInteger)index;
- (BFeedPostBuilder *)addRepliesDeprecated:(BFeedPost*)value;
- (BFeedPostBuilder *)setRepliesDeprecatedArray:(NSArray *)array;
- (BFeedPostBuilder *)clearRepliesDeprecated;

- (BOOL) hasMayAddReply;
- (BOOL) mayAddReply;
- (BFeedPostBuilder*) setMayAddReply:(BOOL) value;
- (BFeedPostBuilder*) clearMayAddReply;

- (BOOL) hasMayChooseMulitpleReplies;
- (BOOL) mayChooseMulitpleReplies;
- (BFeedPostBuilder*) setMayChooseMulitpleReplies:(BOOL) value;
- (BFeedPostBuilder*) clearMayChooseMulitpleReplies;

- (BOOL) hasSurveyAnswerSequence;
- (SInt32) surveyAnswerSequence;
- (BFeedPostBuilder*) setSurveyAnswerSequence:(SInt32) value;
- (BFeedPostBuilder*) clearSurveyAnswerSequence;

- (BOOL) hasAreMoreReplies;
- (BOOL) areMoreReplies;
- (BFeedPostBuilder*) setAreMoreReplies:(BOOL) value;
- (BFeedPostBuilder*) clearAreMoreReplies;

- (BOOL) hasTotalVoteCount;
- (SInt32) totalVoteCount;
- (BFeedPostBuilder*) setTotalVoteCount:(SInt32) value;
- (BFeedPostBuilder*) clearTotalVoteCount;
@end

#define FeedPostUpdateRequest_updateVerb @"updateVerb"
#define FeedPostUpdateRequest_feedPost_deprecated @"feedPostDeprecated"
#define FeedPostUpdateRequest_feedPosts @"feedPosts"
@interface BFeedPostUpdateRequest : PBGeneratedMessage<GeneratedMessageProtocol> {
@private
  BOOL hasFeedPostDeprecated_:1;
  BOOL hasUpdateVerb_:1;
  BFeedPost* feedPostDeprecated;
  BUpdateVerb updateVerb;
  NSMutableArray * feedPostsArray;
}
- (BOOL) hasUpdateVerb;
- (BOOL) hasFeedPostDeprecated;
@property (readonly) BUpdateVerb updateVerb;
@property (readonly, strong) BFeedPost* feedPostDeprecated;
@property (readonly, strong) NSArray * feedPosts;
- (BFeedPost*)feedPostsAtIndex:(NSUInteger)index;

+ (instancetype) defaultInstance;
- (instancetype) defaultInstance;

- (BOOL) isInitialized;
- (void) writeToCodedOutputStream:(PBCodedOutputStream*) output;
- (BFeedPostUpdateRequestBuilder*) builder;
+ (BFeedPostUpdateRequestBuilder*) builder;
+ (BFeedPostUpdateRequestBuilder*) builderWithPrototype:(BFeedPostUpdateRequest*) prototype;
- (BFeedPostUpdateRequestBuilder*) toBuilder;

+ (BFeedPostUpdateRequest*) parseFromData:(NSData*) data;
+ (BFeedPostUpdateRequest*) parseFromData:(NSData*) data extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (BFeedPostUpdateRequest*) parseFromInputStream:(NSInputStream*) input;
+ (BFeedPostUpdateRequest*) parseFromInputStream:(NSInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (BFeedPostUpdateRequest*) parseFromCodedInputStream:(PBCodedInputStream*) input;
+ (BFeedPostUpdateRequest*) parseFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
@end

@interface BFeedPostUpdateRequestBuilder : PBGeneratedMessageBuilder {
@private
  BFeedPostUpdateRequest* resultFeedPostUpdateRequest;
}

- (BFeedPostUpdateRequest*) defaultInstance;

- (BFeedPostUpdateRequestBuilder*) clear;
- (BFeedPostUpdateRequestBuilder*) clone;

- (BFeedPostUpdateRequest*) build;
- (BFeedPostUpdateRequest*) buildPartial;

- (BFeedPostUpdateRequestBuilder*) mergeFrom:(BFeedPostUpdateRequest*) other;
- (BFeedPostUpdateRequestBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input;
- (BFeedPostUpdateRequestBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;

- (BOOL) hasUpdateVerb;
- (BUpdateVerb) updateVerb;
- (BFeedPostUpdateRequestBuilder*) setUpdateVerb:(BUpdateVerb) value;
- (BFeedPostUpdateRequestBuilder*) clearUpdateVerb;

- (BOOL) hasFeedPostDeprecated;
- (BFeedPost*) feedPostDeprecated;
- (BFeedPostUpdateRequestBuilder*) setFeedPostDeprecated:(BFeedPost*) value;
- (BFeedPostUpdateRequestBuilder*) setFeedPostDeprecatedBuilder:(BFeedPostBuilder*) builderForValue;
- (BFeedPostUpdateRequestBuilder*) mergeFeedPostDeprecated:(BFeedPost*) value;
- (BFeedPostUpdateRequestBuilder*) clearFeedPostDeprecated;

- (NSMutableArray *)feedPosts;
- (BFeedPost*)feedPostsAtIndex:(NSUInteger)index;
- (BFeedPostUpdateRequestBuilder *)addFeedPosts:(BFeedPost*)value;
- (BFeedPostUpdateRequestBuilder *)setFeedPostsArray:(NSArray *)array;
- (BFeedPostUpdateRequestBuilder *)clearFeedPosts;
@end

#define FeedPostUpdateResponse_feedPost @"feedPost"
@interface BFeedPostUpdateResponse : PBGeneratedMessage<GeneratedMessageProtocol> {
@private
  BOOL hasFeedPost_:1;
  BFeedPost* feedPost;
}
- (BOOL) hasFeedPost;
@property (readonly, strong) BFeedPost* feedPost;

+ (instancetype) defaultInstance;
- (instancetype) defaultInstance;

- (BOOL) isInitialized;
- (void) writeToCodedOutputStream:(PBCodedOutputStream*) output;
- (BFeedPostUpdateResponseBuilder*) builder;
+ (BFeedPostUpdateResponseBuilder*) builder;
+ (BFeedPostUpdateResponseBuilder*) builderWithPrototype:(BFeedPostUpdateResponse*) prototype;
- (BFeedPostUpdateResponseBuilder*) toBuilder;

+ (BFeedPostUpdateResponse*) parseFromData:(NSData*) data;
+ (BFeedPostUpdateResponse*) parseFromData:(NSData*) data extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (BFeedPostUpdateResponse*) parseFromInputStream:(NSInputStream*) input;
+ (BFeedPostUpdateResponse*) parseFromInputStream:(NSInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (BFeedPostUpdateResponse*) parseFromCodedInputStream:(PBCodedInputStream*) input;
+ (BFeedPostUpdateResponse*) parseFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
@end

@interface BFeedPostUpdateResponseBuilder : PBGeneratedMessageBuilder {
@private
  BFeedPostUpdateResponse* resultFeedPostUpdateResponse;
}

- (BFeedPostUpdateResponse*) defaultInstance;

- (BFeedPostUpdateResponseBuilder*) clear;
- (BFeedPostUpdateResponseBuilder*) clone;

- (BFeedPostUpdateResponse*) build;
- (BFeedPostUpdateResponse*) buildPartial;

- (BFeedPostUpdateResponseBuilder*) mergeFrom:(BFeedPostUpdateResponse*) other;
- (BFeedPostUpdateResponseBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input;
- (BFeedPostUpdateResponseBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;

- (BOOL) hasFeedPost;
- (BFeedPost*) feedPost;
- (BFeedPostUpdateResponseBuilder*) setFeedPost:(BFeedPost*) value;
- (BFeedPostUpdateResponseBuilder*) setFeedPostBuilder:(BFeedPostBuilder*) builderForValue;
- (BFeedPostUpdateResponseBuilder*) mergeFeedPost:(BFeedPost*) value;
- (BFeedPostUpdateResponseBuilder*) clearFeedPost;
@end

#define FeedPostFetchRequest_timespan @"timespan"
#define FeedPostFetchRequest_feedScope @"feedScope"
#define FeedPostFetchRequest_parentID @"parentID"
@interface BFeedPostFetchRequest : PBGeneratedMessage<GeneratedMessageProtocol> {
@private
  BOOL hasParentID_:1;
  BOOL hasTimespan_:1;
  BOOL hasFeedScope_:1;
  NSString* parentID;
  BTimespan* timespan;
  BFeedPostScope feedScope;
}
- (BOOL) hasTimespan;
- (BOOL) hasFeedScope;
- (BOOL) hasParentID;
@property (readonly, strong) BTimespan* timespan;
@property (readonly) BFeedPostScope feedScope;
@property (readonly, strong) NSString* parentID;

+ (instancetype) defaultInstance;
- (instancetype) defaultInstance;

- (BOOL) isInitialized;
- (void) writeToCodedOutputStream:(PBCodedOutputStream*) output;
- (BFeedPostFetchRequestBuilder*) builder;
+ (BFeedPostFetchRequestBuilder*) builder;
+ (BFeedPostFetchRequestBuilder*) builderWithPrototype:(BFeedPostFetchRequest*) prototype;
- (BFeedPostFetchRequestBuilder*) toBuilder;

+ (BFeedPostFetchRequest*) parseFromData:(NSData*) data;
+ (BFeedPostFetchRequest*) parseFromData:(NSData*) data extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (BFeedPostFetchRequest*) parseFromInputStream:(NSInputStream*) input;
+ (BFeedPostFetchRequest*) parseFromInputStream:(NSInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (BFeedPostFetchRequest*) parseFromCodedInputStream:(PBCodedInputStream*) input;
+ (BFeedPostFetchRequest*) parseFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
@end

@interface BFeedPostFetchRequestBuilder : PBGeneratedMessageBuilder {
@private
  BFeedPostFetchRequest* resultFeedPostFetchRequest;
}

- (BFeedPostFetchRequest*) defaultInstance;

- (BFeedPostFetchRequestBuilder*) clear;
- (BFeedPostFetchRequestBuilder*) clone;

- (BFeedPostFetchRequest*) build;
- (BFeedPostFetchRequest*) buildPartial;

- (BFeedPostFetchRequestBuilder*) mergeFrom:(BFeedPostFetchRequest*) other;
- (BFeedPostFetchRequestBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input;
- (BFeedPostFetchRequestBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;

- (BOOL) hasTimespan;
- (BTimespan*) timespan;
- (BFeedPostFetchRequestBuilder*) setTimespan:(BTimespan*) value;
- (BFeedPostFetchRequestBuilder*) setTimespanBuilder:(BTimespanBuilder*) builderForValue;
- (BFeedPostFetchRequestBuilder*) mergeTimespan:(BTimespan*) value;
- (BFeedPostFetchRequestBuilder*) clearTimespan;

- (BOOL) hasFeedScope;
- (BFeedPostScope) feedScope;
- (BFeedPostFetchRequestBuilder*) setFeedScope:(BFeedPostScope) value;
- (BFeedPostFetchRequestBuilder*) clearFeedScope;

- (BOOL) hasParentID;
- (NSString*) parentID;
- (BFeedPostFetchRequestBuilder*) setParentID:(NSString*) value;
- (BFeedPostFetchRequestBuilder*) clearParentID;
@end

#define FeedPostFetchResponse_feedPosts @"feedPosts"
@interface BFeedPostFetchResponse : PBGeneratedMessage<GeneratedMessageProtocol> {
@private
  NSMutableArray * feedPostsArray;
}
@property (readonly, strong) NSArray * feedPosts;
- (BFeedPost*)feedPostsAtIndex:(NSUInteger)index;

+ (instancetype) defaultInstance;
- (instancetype) defaultInstance;

- (BOOL) isInitialized;
- (void) writeToCodedOutputStream:(PBCodedOutputStream*) output;
- (BFeedPostFetchResponseBuilder*) builder;
+ (BFeedPostFetchResponseBuilder*) builder;
+ (BFeedPostFetchResponseBuilder*) builderWithPrototype:(BFeedPostFetchResponse*) prototype;
- (BFeedPostFetchResponseBuilder*) toBuilder;

+ (BFeedPostFetchResponse*) parseFromData:(NSData*) data;
+ (BFeedPostFetchResponse*) parseFromData:(NSData*) data extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (BFeedPostFetchResponse*) parseFromInputStream:(NSInputStream*) input;
+ (BFeedPostFetchResponse*) parseFromInputStream:(NSInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (BFeedPostFetchResponse*) parseFromCodedInputStream:(PBCodedInputStream*) input;
+ (BFeedPostFetchResponse*) parseFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
@end

@interface BFeedPostFetchResponseBuilder : PBGeneratedMessageBuilder {
@private
  BFeedPostFetchResponse* resultFeedPostFetchResponse;
}

- (BFeedPostFetchResponse*) defaultInstance;

- (BFeedPostFetchResponseBuilder*) clear;
- (BFeedPostFetchResponseBuilder*) clone;

- (BFeedPostFetchResponse*) build;
- (BFeedPostFetchResponse*) buildPartial;

- (BFeedPostFetchResponseBuilder*) mergeFrom:(BFeedPostFetchResponse*) other;
- (BFeedPostFetchResponseBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input;
- (BFeedPostFetchResponseBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;

- (NSMutableArray *)feedPosts;
- (BFeedPost*)feedPostsAtIndex:(NSUInteger)index;
- (BFeedPostFetchResponseBuilder *)addFeedPosts:(BFeedPost*)value;
- (BFeedPostFetchResponseBuilder *)setFeedPostsArray:(NSArray *)array;
- (BFeedPostFetchResponseBuilder *)clearFeedPosts;
@end


// @@protoc_insertion_point(global_scope)
