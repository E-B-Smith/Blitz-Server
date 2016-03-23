// Generated by the protocol buffer compiler.  DO NOT EDIT!

#import <ProtocolBuffers/ProtocolBuffers.h>

// @@protoc_insertion_point(imports)

@class CMChatConnect;
@class CMChatConnectBuilder;
@class CMChatEnterRoom;
@class CMChatEnterRoomBuilder;
@class CMChatMessage;
@class CMChatMessageBuilder;
@class CMChatMessageType;
@class CMChatMessageTypeBuilder;
@class CMChatPresence;
@class CMChatPresenceBuilder;
@class CMChatResponse;
@class CMChatResponseBuilder;
@class CMChatRoom;
@class CMChatRoomBuilder;
@class CMChatUser;
@class CMChatUserBuilder;
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


typedef NS_ENUM(SInt32, CMStatusCode) {
  CMStatusCodeStatusSuccess = 1,
  CMStatusCodeStatusInputInvalid = 2,
  CMStatusCodeStatusNotAuthorized = 3,
  CMStatusCodeStatusServerError = 4,
};

BOOL CMStatusCodeIsValidValue(CMStatusCode value);
NSString *NSStringFromCMStatusCode(CMStatusCode value);


@interface CMChatRoot : NSObject {
}
+ (PBExtensionRegistry*) extensionRegistry;
+ (void) registerAllExtensions:(PBMutableExtensionRegistry*) registry;
@end

#define ChatMessage_senderID @"senderID"
#define ChatMessage_roomID @"roomID"
#define ChatMessage_timestamp @"timestamp"
#define ChatMessage_message @"message"
@interface CMChatMessage : PBGeneratedMessage<GeneratedMessageProtocol> {
@private
  BOOL hasTimestamp_:1;
  BOOL hasSenderID_:1;
  BOOL hasRoomID_:1;
  BOOL hasMessage_:1;
  Float64 timestamp;
  NSString* senderID;
  NSString* roomID;
  NSString* message;
}
- (BOOL) hasSenderID;
- (BOOL) hasRoomID;
- (BOOL) hasTimestamp;
- (BOOL) hasMessage;
@property (readonly, strong) NSString* senderID;
@property (readonly, strong) NSString* roomID;
@property (readonly) Float64 timestamp;
@property (readonly, strong) NSString* message;

+ (instancetype) defaultInstance;
- (instancetype) defaultInstance;

- (BOOL) isInitialized;
- (void) writeToCodedOutputStream:(PBCodedOutputStream*) output;
- (CMChatMessageBuilder*) builder;
+ (CMChatMessageBuilder*) builder;
+ (CMChatMessageBuilder*) builderWithPrototype:(CMChatMessage*) prototype;
- (CMChatMessageBuilder*) toBuilder;

+ (CMChatMessage*) parseFromData:(NSData*) data;
+ (CMChatMessage*) parseFromData:(NSData*) data extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (CMChatMessage*) parseFromInputStream:(NSInputStream*) input;
+ (CMChatMessage*) parseFromInputStream:(NSInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (CMChatMessage*) parseFromCodedInputStream:(PBCodedInputStream*) input;
+ (CMChatMessage*) parseFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
@end

@interface CMChatMessageBuilder : PBGeneratedMessageBuilder {
@private
  CMChatMessage* resultChatMessage;
}

- (CMChatMessage*) defaultInstance;

- (CMChatMessageBuilder*) clear;
- (CMChatMessageBuilder*) clone;

- (CMChatMessage*) build;
- (CMChatMessage*) buildPartial;

- (CMChatMessageBuilder*) mergeFrom:(CMChatMessage*) other;
- (CMChatMessageBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input;
- (CMChatMessageBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;

- (BOOL) hasSenderID;
- (NSString*) senderID;
- (CMChatMessageBuilder*) setSenderID:(NSString*) value;
- (CMChatMessageBuilder*) clearSenderID;

- (BOOL) hasRoomID;
- (NSString*) roomID;
- (CMChatMessageBuilder*) setRoomID:(NSString*) value;
- (CMChatMessageBuilder*) clearRoomID;

- (BOOL) hasTimestamp;
- (Float64) timestamp;
- (CMChatMessageBuilder*) setTimestamp:(Float64) value;
- (CMChatMessageBuilder*) clearTimestamp;

- (BOOL) hasMessage;
- (NSString*) message;
- (CMChatMessageBuilder*) setMessage:(NSString*) value;
- (CMChatMessageBuilder*) clearMessage;
@end

#define ChatUser_userID @"userID"
#define ChatUser_nickname @"nickname"
@interface CMChatUser : PBGeneratedMessage<GeneratedMessageProtocol> {
@private
  BOOL hasUserID_:1;
  BOOL hasNickname_:1;
  NSString* userID;
  NSString* nickname;
}
- (BOOL) hasUserID;
- (BOOL) hasNickname;
@property (readonly, strong) NSString* userID;
@property (readonly, strong) NSString* nickname;

+ (instancetype) defaultInstance;
- (instancetype) defaultInstance;

- (BOOL) isInitialized;
- (void) writeToCodedOutputStream:(PBCodedOutputStream*) output;
- (CMChatUserBuilder*) builder;
+ (CMChatUserBuilder*) builder;
+ (CMChatUserBuilder*) builderWithPrototype:(CMChatUser*) prototype;
- (CMChatUserBuilder*) toBuilder;

+ (CMChatUser*) parseFromData:(NSData*) data;
+ (CMChatUser*) parseFromData:(NSData*) data extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (CMChatUser*) parseFromInputStream:(NSInputStream*) input;
+ (CMChatUser*) parseFromInputStream:(NSInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (CMChatUser*) parseFromCodedInputStream:(PBCodedInputStream*) input;
+ (CMChatUser*) parseFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
@end

@interface CMChatUserBuilder : PBGeneratedMessageBuilder {
@private
  CMChatUser* resultChatUser;
}

- (CMChatUser*) defaultInstance;

- (CMChatUserBuilder*) clear;
- (CMChatUserBuilder*) clone;

- (CMChatUser*) build;
- (CMChatUser*) buildPartial;

- (CMChatUserBuilder*) mergeFrom:(CMChatUser*) other;
- (CMChatUserBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input;
- (CMChatUserBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;

- (BOOL) hasUserID;
- (NSString*) userID;
- (CMChatUserBuilder*) setUserID:(NSString*) value;
- (CMChatUserBuilder*) clearUserID;

- (BOOL) hasNickname;
- (NSString*) nickname;
- (CMChatUserBuilder*) setNickname:(NSString*) value;
- (CMChatUserBuilder*) clearNickname;
@end

#define ChatRoom_roomID @"roomID"
#define ChatRoom_roomName @"roomName"
@interface CMChatRoom : PBGeneratedMessage<GeneratedMessageProtocol> {
@private
  BOOL hasRoomID_:1;
  BOOL hasRoomName_:1;
  NSString* roomID;
  NSString* roomName;
}
- (BOOL) hasRoomID;
- (BOOL) hasRoomName;
@property (readonly, strong) NSString* roomID;
@property (readonly, strong) NSString* roomName;

+ (instancetype) defaultInstance;
- (instancetype) defaultInstance;

- (BOOL) isInitialized;
- (void) writeToCodedOutputStream:(PBCodedOutputStream*) output;
- (CMChatRoomBuilder*) builder;
+ (CMChatRoomBuilder*) builder;
+ (CMChatRoomBuilder*) builderWithPrototype:(CMChatRoom*) prototype;
- (CMChatRoomBuilder*) toBuilder;

+ (CMChatRoom*) parseFromData:(NSData*) data;
+ (CMChatRoom*) parseFromData:(NSData*) data extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (CMChatRoom*) parseFromInputStream:(NSInputStream*) input;
+ (CMChatRoom*) parseFromInputStream:(NSInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (CMChatRoom*) parseFromCodedInputStream:(PBCodedInputStream*) input;
+ (CMChatRoom*) parseFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
@end

@interface CMChatRoomBuilder : PBGeneratedMessageBuilder {
@private
  CMChatRoom* resultChatRoom;
}

- (CMChatRoom*) defaultInstance;

- (CMChatRoomBuilder*) clear;
- (CMChatRoomBuilder*) clone;

- (CMChatRoom*) build;
- (CMChatRoom*) buildPartial;

- (CMChatRoomBuilder*) mergeFrom:(CMChatRoom*) other;
- (CMChatRoomBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input;
- (CMChatRoomBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;

- (BOOL) hasRoomID;
- (NSString*) roomID;
- (CMChatRoomBuilder*) setRoomID:(NSString*) value;
- (CMChatRoomBuilder*) clearRoomID;

- (BOOL) hasRoomName;
- (NSString*) roomName;
- (CMChatRoomBuilder*) setRoomName:(NSString*) value;
- (CMChatRoomBuilder*) clearRoomName;
@end

#define ChatConnect_isConnecting @"isConnecting"
#define ChatConnect_user @"user"
#define ChatConnect_rooms @"rooms"
@interface CMChatConnect : PBGeneratedMessage<GeneratedMessageProtocol> {
@private
  BOOL hasIsConnecting_:1;
  BOOL hasUser_:1;
  BOOL isConnecting_:1;
  CMChatUser* user;
  NSMutableArray * roomsArray;
}
- (BOOL) hasIsConnecting;
- (BOOL) hasUser;
- (BOOL) isConnecting;
@property (readonly, strong) CMChatUser* user;
@property (readonly, strong) NSArray * rooms;
- (CMChatRoom*)roomsAtIndex:(NSUInteger)index;

+ (instancetype) defaultInstance;
- (instancetype) defaultInstance;

- (BOOL) isInitialized;
- (void) writeToCodedOutputStream:(PBCodedOutputStream*) output;
- (CMChatConnectBuilder*) builder;
+ (CMChatConnectBuilder*) builder;
+ (CMChatConnectBuilder*) builderWithPrototype:(CMChatConnect*) prototype;
- (CMChatConnectBuilder*) toBuilder;

+ (CMChatConnect*) parseFromData:(NSData*) data;
+ (CMChatConnect*) parseFromData:(NSData*) data extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (CMChatConnect*) parseFromInputStream:(NSInputStream*) input;
+ (CMChatConnect*) parseFromInputStream:(NSInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (CMChatConnect*) parseFromCodedInputStream:(PBCodedInputStream*) input;
+ (CMChatConnect*) parseFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
@end

@interface CMChatConnectBuilder : PBGeneratedMessageBuilder {
@private
  CMChatConnect* resultChatConnect;
}

- (CMChatConnect*) defaultInstance;

- (CMChatConnectBuilder*) clear;
- (CMChatConnectBuilder*) clone;

- (CMChatConnect*) build;
- (CMChatConnect*) buildPartial;

- (CMChatConnectBuilder*) mergeFrom:(CMChatConnect*) other;
- (CMChatConnectBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input;
- (CMChatConnectBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;

- (BOOL) hasIsConnecting;
- (BOOL) isConnecting;
- (CMChatConnectBuilder*) setIsConnecting:(BOOL) value;
- (CMChatConnectBuilder*) clearIsConnecting;

- (BOOL) hasUser;
- (CMChatUser*) user;
- (CMChatConnectBuilder*) setUser:(CMChatUser*) value;
- (CMChatConnectBuilder*) setUserBuilder:(CMChatUserBuilder*) builderForValue;
- (CMChatConnectBuilder*) mergeUser:(CMChatUser*) value;
- (CMChatConnectBuilder*) clearUser;

- (NSMutableArray *)rooms;
- (CMChatRoom*)roomsAtIndex:(NSUInteger)index;
- (CMChatConnectBuilder *)addRooms:(CMChatRoom*)value;
- (CMChatConnectBuilder *)setRoomsArray:(NSArray *)array;
- (CMChatConnectBuilder *)clearRooms;
@end

#define ChatEnterRoom_user @"user"
#define ChatEnterRoom_roomID @"roomID"
#define ChatEnterRoom_userIsEntering @"userIsEntering"
@interface CMChatEnterRoom : PBGeneratedMessage<GeneratedMessageProtocol> {
@private
  BOOL hasUserIsEntering_:1;
  BOOL hasRoomID_:1;
  BOOL hasUser_:1;
  BOOL userIsEntering_:1;
  NSString* roomID;
  CMChatUser* user;
}
- (BOOL) hasUser;
- (BOOL) hasRoomID;
- (BOOL) hasUserIsEntering;
@property (readonly, strong) CMChatUser* user;
@property (readonly, strong) NSString* roomID;
- (BOOL) userIsEntering;

+ (instancetype) defaultInstance;
- (instancetype) defaultInstance;

- (BOOL) isInitialized;
- (void) writeToCodedOutputStream:(PBCodedOutputStream*) output;
- (CMChatEnterRoomBuilder*) builder;
+ (CMChatEnterRoomBuilder*) builder;
+ (CMChatEnterRoomBuilder*) builderWithPrototype:(CMChatEnterRoom*) prototype;
- (CMChatEnterRoomBuilder*) toBuilder;

+ (CMChatEnterRoom*) parseFromData:(NSData*) data;
+ (CMChatEnterRoom*) parseFromData:(NSData*) data extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (CMChatEnterRoom*) parseFromInputStream:(NSInputStream*) input;
+ (CMChatEnterRoom*) parseFromInputStream:(NSInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (CMChatEnterRoom*) parseFromCodedInputStream:(PBCodedInputStream*) input;
+ (CMChatEnterRoom*) parseFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
@end

@interface CMChatEnterRoomBuilder : PBGeneratedMessageBuilder {
@private
  CMChatEnterRoom* resultChatEnterRoom;
}

- (CMChatEnterRoom*) defaultInstance;

- (CMChatEnterRoomBuilder*) clear;
- (CMChatEnterRoomBuilder*) clone;

- (CMChatEnterRoom*) build;
- (CMChatEnterRoom*) buildPartial;

- (CMChatEnterRoomBuilder*) mergeFrom:(CMChatEnterRoom*) other;
- (CMChatEnterRoomBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input;
- (CMChatEnterRoomBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;

- (BOOL) hasUser;
- (CMChatUser*) user;
- (CMChatEnterRoomBuilder*) setUser:(CMChatUser*) value;
- (CMChatEnterRoomBuilder*) setUserBuilder:(CMChatUserBuilder*) builderForValue;
- (CMChatEnterRoomBuilder*) mergeUser:(CMChatUser*) value;
- (CMChatEnterRoomBuilder*) clearUser;

- (BOOL) hasRoomID;
- (NSString*) roomID;
- (CMChatEnterRoomBuilder*) setRoomID:(NSString*) value;
- (CMChatEnterRoomBuilder*) clearRoomID;

- (BOOL) hasUserIsEntering;
- (BOOL) userIsEntering;
- (CMChatEnterRoomBuilder*) setUserIsEntering:(BOOL) value;
- (CMChatEnterRoomBuilder*) clearUserIsEntering;
@end

#define ChatPresence_room @"room"
#define ChatPresence_users @"users"
@interface CMChatPresence : PBGeneratedMessage<GeneratedMessageProtocol> {
@private
  BOOL hasRoom_:1;
  CMChatRoom* room;
  NSMutableArray * usersArray;
}
- (BOOL) hasRoom;
@property (readonly, strong) CMChatRoom* room;
@property (readonly, strong) NSArray * users;
- (CMChatUser*)usersAtIndex:(NSUInteger)index;

+ (instancetype) defaultInstance;
- (instancetype) defaultInstance;

- (BOOL) isInitialized;
- (void) writeToCodedOutputStream:(PBCodedOutputStream*) output;
- (CMChatPresenceBuilder*) builder;
+ (CMChatPresenceBuilder*) builder;
+ (CMChatPresenceBuilder*) builderWithPrototype:(CMChatPresence*) prototype;
- (CMChatPresenceBuilder*) toBuilder;

+ (CMChatPresence*) parseFromData:(NSData*) data;
+ (CMChatPresence*) parseFromData:(NSData*) data extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (CMChatPresence*) parseFromInputStream:(NSInputStream*) input;
+ (CMChatPresence*) parseFromInputStream:(NSInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (CMChatPresence*) parseFromCodedInputStream:(PBCodedInputStream*) input;
+ (CMChatPresence*) parseFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
@end

@interface CMChatPresenceBuilder : PBGeneratedMessageBuilder {
@private
  CMChatPresence* resultChatPresence;
}

- (CMChatPresence*) defaultInstance;

- (CMChatPresenceBuilder*) clear;
- (CMChatPresenceBuilder*) clone;

- (CMChatPresence*) build;
- (CMChatPresence*) buildPartial;

- (CMChatPresenceBuilder*) mergeFrom:(CMChatPresence*) other;
- (CMChatPresenceBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input;
- (CMChatPresenceBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;

- (BOOL) hasRoom;
- (CMChatRoom*) room;
- (CMChatPresenceBuilder*) setRoom:(CMChatRoom*) value;
- (CMChatPresenceBuilder*) setRoomBuilder:(CMChatRoomBuilder*) builderForValue;
- (CMChatPresenceBuilder*) mergeRoom:(CMChatRoom*) value;
- (CMChatPresenceBuilder*) clearRoom;

- (NSMutableArray *)users;
- (CMChatUser*)usersAtIndex:(NSUInteger)index;
- (CMChatPresenceBuilder *)addUsers:(CMChatUser*)value;
- (CMChatPresenceBuilder *)setUsersArray:(NSArray *)array;
- (CMChatPresenceBuilder *)clearUsers;
@end

#define ChatResponse_code @"code"
#define ChatResponse_message @"message"
@interface CMChatResponse : PBGeneratedMessage<GeneratedMessageProtocol> {
@private
  BOOL hasMessage_:1;
  BOOL hasCode_:1;
  NSString* message;
  CMStatusCode code;
}
- (BOOL) hasCode;
- (BOOL) hasMessage;
@property (readonly) CMStatusCode code;
@property (readonly, strong) NSString* message;

+ (instancetype) defaultInstance;
- (instancetype) defaultInstance;

- (BOOL) isInitialized;
- (void) writeToCodedOutputStream:(PBCodedOutputStream*) output;
- (CMChatResponseBuilder*) builder;
+ (CMChatResponseBuilder*) builder;
+ (CMChatResponseBuilder*) builderWithPrototype:(CMChatResponse*) prototype;
- (CMChatResponseBuilder*) toBuilder;

+ (CMChatResponse*) parseFromData:(NSData*) data;
+ (CMChatResponse*) parseFromData:(NSData*) data extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (CMChatResponse*) parseFromInputStream:(NSInputStream*) input;
+ (CMChatResponse*) parseFromInputStream:(NSInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (CMChatResponse*) parseFromCodedInputStream:(PBCodedInputStream*) input;
+ (CMChatResponse*) parseFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
@end

@interface CMChatResponseBuilder : PBGeneratedMessageBuilder {
@private
  CMChatResponse* resultChatResponse;
}

- (CMChatResponse*) defaultInstance;

- (CMChatResponseBuilder*) clear;
- (CMChatResponseBuilder*) clone;

- (CMChatResponse*) build;
- (CMChatResponse*) buildPartial;

- (CMChatResponseBuilder*) mergeFrom:(CMChatResponse*) other;
- (CMChatResponseBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input;
- (CMChatResponseBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;

- (BOOL) hasCode;
- (CMStatusCode) code;
- (CMChatResponseBuilder*) setCode:(CMStatusCode) value;
- (CMChatResponseBuilder*) clearCode;

- (BOOL) hasMessage;
- (NSString*) message;
- (CMChatResponseBuilder*) setMessage:(NSString*) value;
- (CMChatResponseBuilder*) clearMessage;
@end

#define ChatMessageType_chatMessage @"chatMessage"
#define ChatMessageType_chatConnect @"chatConnect"
#define ChatMessageType_chatEnterRoom @"chatEnterRoom"
#define ChatMessageType_chatPresence @"chatPresence"
#define ChatMessageType_chatResponse @"chatResponse"
@interface CMChatMessageType : PBGeneratedMessage<GeneratedMessageProtocol> {
@private
  BOOL hasChatMessage_:1;
  BOOL hasChatConnect_:1;
  BOOL hasChatEnterRoom_:1;
  BOOL hasChatPresence_:1;
  BOOL hasChatResponse_:1;
  CMChatMessage* chatMessage;
  CMChatConnect* chatConnect;
  CMChatEnterRoom* chatEnterRoom;
  CMChatPresence* chatPresence;
  CMChatResponse* chatResponse;
}
- (BOOL) hasChatMessage;
- (BOOL) hasChatConnect;
- (BOOL) hasChatEnterRoom;
- (BOOL) hasChatPresence;
- (BOOL) hasChatResponse;
@property (readonly, strong) CMChatMessage* chatMessage;
@property (readonly, strong) CMChatConnect* chatConnect;
@property (readonly, strong) CMChatEnterRoom* chatEnterRoom;
@property (readonly, strong) CMChatPresence* chatPresence;
@property (readonly, strong) CMChatResponse* chatResponse;

+ (instancetype) defaultInstance;
- (instancetype) defaultInstance;

- (BOOL) isInitialized;
- (void) writeToCodedOutputStream:(PBCodedOutputStream*) output;
- (CMChatMessageTypeBuilder*) builder;
+ (CMChatMessageTypeBuilder*) builder;
+ (CMChatMessageTypeBuilder*) builderWithPrototype:(CMChatMessageType*) prototype;
- (CMChatMessageTypeBuilder*) toBuilder;

+ (CMChatMessageType*) parseFromData:(NSData*) data;
+ (CMChatMessageType*) parseFromData:(NSData*) data extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (CMChatMessageType*) parseFromInputStream:(NSInputStream*) input;
+ (CMChatMessageType*) parseFromInputStream:(NSInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (CMChatMessageType*) parseFromCodedInputStream:(PBCodedInputStream*) input;
+ (CMChatMessageType*) parseFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
@end

@interface CMChatMessageTypeBuilder : PBGeneratedMessageBuilder {
@private
  CMChatMessageType* resultChatMessageType;
}

- (CMChatMessageType*) defaultInstance;

- (CMChatMessageTypeBuilder*) clear;
- (CMChatMessageTypeBuilder*) clone;

- (CMChatMessageType*) build;
- (CMChatMessageType*) buildPartial;

- (CMChatMessageTypeBuilder*) mergeFrom:(CMChatMessageType*) other;
- (CMChatMessageTypeBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input;
- (CMChatMessageTypeBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;

- (BOOL) hasChatMessage;
- (CMChatMessage*) chatMessage;
- (CMChatMessageTypeBuilder*) setChatMessage:(CMChatMessage*) value;
- (CMChatMessageTypeBuilder*) setChatMessageBuilder:(CMChatMessageBuilder*) builderForValue;
- (CMChatMessageTypeBuilder*) mergeChatMessage:(CMChatMessage*) value;
- (CMChatMessageTypeBuilder*) clearChatMessage;

- (BOOL) hasChatConnect;
- (CMChatConnect*) chatConnect;
- (CMChatMessageTypeBuilder*) setChatConnect:(CMChatConnect*) value;
- (CMChatMessageTypeBuilder*) setChatConnectBuilder:(CMChatConnectBuilder*) builderForValue;
- (CMChatMessageTypeBuilder*) mergeChatConnect:(CMChatConnect*) value;
- (CMChatMessageTypeBuilder*) clearChatConnect;

- (BOOL) hasChatEnterRoom;
- (CMChatEnterRoom*) chatEnterRoom;
- (CMChatMessageTypeBuilder*) setChatEnterRoom:(CMChatEnterRoom*) value;
- (CMChatMessageTypeBuilder*) setChatEnterRoomBuilder:(CMChatEnterRoomBuilder*) builderForValue;
- (CMChatMessageTypeBuilder*) mergeChatEnterRoom:(CMChatEnterRoom*) value;
- (CMChatMessageTypeBuilder*) clearChatEnterRoom;

- (BOOL) hasChatPresence;
- (CMChatPresence*) chatPresence;
- (CMChatMessageTypeBuilder*) setChatPresence:(CMChatPresence*) value;
- (CMChatMessageTypeBuilder*) setChatPresenceBuilder:(CMChatPresenceBuilder*) builderForValue;
- (CMChatMessageTypeBuilder*) mergeChatPresence:(CMChatPresence*) value;
- (CMChatMessageTypeBuilder*) clearChatPresence;

- (BOOL) hasChatResponse;
- (CMChatResponse*) chatResponse;
- (CMChatMessageTypeBuilder*) setChatResponse:(CMChatResponse*) value;
- (CMChatMessageTypeBuilder*) setChatResponseBuilder:(CMChatResponseBuilder*) builderForValue;
- (CMChatMessageTypeBuilder*) mergeChatResponse:(CMChatResponse*) value;
- (CMChatMessageTypeBuilder*) clearChatResponse;
@end


// @@protoc_insertion_point(global_scope)