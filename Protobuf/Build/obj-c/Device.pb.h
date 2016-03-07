// Generated by the protocol buffer compiler.  DO NOT EDIT!

#import <ProtocolBuffers/ProtocolBuffers.h>

#import "Types.pb.h"
// @@protoc_insertion_point(imports)

@class BCoordinate;
@class BCoordinateBuilder;
@class BCoordinatePolygon;
@class BCoordinatePolygonBuilder;
@class BCoordinateRegion;
@class BCoordinateRegionBuilder;
@class BDeviceInfo;
@class BDeviceInfoBuilder;
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


typedef NS_ENUM(SInt32, BPlatformType) {
  BPlatformTypePTUnknown = 0,
  BPlatformTypePTiOS = 1,
  BPlatformTypePTAndroid = 2,
  BPlatformTypePTWeb = 3,
};

BOOL BPlatformTypeIsValidValue(BPlatformType value);
NSString *NSStringFromBPlatformType(BPlatformType value);


@interface BDeviceRoot : NSObject {
}
+ (PBExtensionRegistry*) extensionRegistry;
+ (void) registerAllExtensions:(PBMutableExtensionRegistry*) registry;
@end

#define DeviceInfo_vendorUID @"vendorUID"
#define DeviceInfo_advertisingUID @"advertisingUID"
#define DeviceInfo_platformType @"platformType"
#define DeviceInfo_modelName @"modelName"
#define DeviceInfo_systemVersion @"systemVersion"
#define DeviceInfo_language @"language"
#define DeviceInfo_timezone @"timezone"
#define DeviceInfo_phoneCountryCode @"phoneCountryCode"
#define DeviceInfo_screenSize @"screenSize"
#define DeviceInfo_screenScale @"screenScale"
#define DeviceInfo_appID @"appID"
#define DeviceInfo_appVersion @"appVersion"
#define DeviceInfo_notificationToken @"notificationToken"
#define DeviceInfo_lastContentRefresh_Deprecated @"lastContentRefreshDeprecated"
#define DeviceInfo_userTags @"userTags"
#define DeviceInfo_deviceUDID @"deviceUDID"
#define DeviceInfo_appIsReleaseVersion @"appIsReleaseVersion"
#define DeviceInfo_colorDepth @"colorDepth"
#define DeviceInfo_IPAddress @"iPAddress"
#define DeviceInfo_systemBuildVersion @"systemBuildVersion"
#define DeviceInfo_localIPAddress @"localIPAddress"
@interface BDeviceInfo : PBGeneratedMessage<GeneratedMessageProtocol> {
@private
  BOOL hasAppIsReleaseVersion_:1;
  BOOL hasScreenScale_:1;
  BOOL hasColorDepth_:1;
  BOOL hasVendorUID_:1;
  BOOL hasAdvertisingUID_:1;
  BOOL hasModelName_:1;
  BOOL hasSystemVersion_:1;
  BOOL hasLanguage_:1;
  BOOL hasTimezone_:1;
  BOOL hasPhoneCountryCode_:1;
  BOOL hasAppID_:1;
  BOOL hasAppVersion_:1;
  BOOL hasNotificationToken_:1;
  BOOL hasDeviceUDID_:1;
  BOOL hasIPAddress_:1;
  BOOL hasSystemBuildVersion_:1;
  BOOL hasLocalIPAddress_:1;
  BOOL hasScreenSize_:1;
  BOOL hasLastContentRefreshDeprecated_:1;
  BOOL hasPlatformType_:1;
  BOOL appIsReleaseVersion_:1;
  Float32 screenScale;
  Float32 colorDepth;
  NSString* vendorUID;
  NSString* advertisingUID;
  NSString* modelName;
  NSString* systemVersion;
  NSString* language;
  NSString* timezone;
  NSString* phoneCountryCode;
  NSString* appID;
  NSString* appVersion;
  NSString* notificationToken;
  NSString* deviceUDID;
  NSString* iPAddress;
  NSString* systemBuildVersion;
  NSString* localIPAddress;
  BSize* screenSize;
  BTimestamp* lastContentRefreshDeprecated;
  BPlatformType platformType;
  NSMutableArray * userTagsArray;
}
- (BOOL) hasVendorUID;
- (BOOL) hasAdvertisingUID;
- (BOOL) hasPlatformType;
- (BOOL) hasModelName;
- (BOOL) hasSystemVersion;
- (BOOL) hasLanguage;
- (BOOL) hasTimezone;
- (BOOL) hasPhoneCountryCode;
- (BOOL) hasScreenSize;
- (BOOL) hasScreenScale;
- (BOOL) hasAppID;
- (BOOL) hasAppVersion;
- (BOOL) hasNotificationToken;
- (BOOL) hasLastContentRefreshDeprecated;
- (BOOL) hasDeviceUDID;
- (BOOL) hasAppIsReleaseVersion;
- (BOOL) hasColorDepth;
- (BOOL) hasIPAddress;
- (BOOL) hasSystemBuildVersion;
- (BOOL) hasLocalIPAddress;
@property (readonly, strong) NSString* vendorUID;
@property (readonly, strong) NSString* advertisingUID;
@property (readonly) BPlatformType platformType;
@property (readonly, strong) NSString* modelName;
@property (readonly, strong) NSString* systemVersion;
@property (readonly, strong) NSString* language;
@property (readonly, strong) NSString* timezone;
@property (readonly, strong) NSString* phoneCountryCode;
@property (readonly, strong) BSize* screenSize;
@property (readonly) Float32 screenScale;
@property (readonly, strong) NSString* appID;
@property (readonly, strong) NSString* appVersion;
@property (readonly, strong) NSString* notificationToken;
@property (readonly, strong) BTimestamp* lastContentRefreshDeprecated;
@property (readonly, strong) NSArray * userTags;
@property (readonly, strong) NSString* deviceUDID;
- (BOOL) appIsReleaseVersion;
@property (readonly) Float32 colorDepth;
@property (readonly, strong) NSString* iPAddress;
@property (readonly, strong) NSString* systemBuildVersion;
@property (readonly, strong) NSString* localIPAddress;
- (NSString*)userTagsAtIndex:(NSUInteger)index;

+ (instancetype) defaultInstance;
- (instancetype) defaultInstance;

- (BOOL) isInitialized;
- (void) writeToCodedOutputStream:(PBCodedOutputStream*) output;
- (BDeviceInfoBuilder*) builder;
+ (BDeviceInfoBuilder*) builder;
+ (BDeviceInfoBuilder*) builderWithPrototype:(BDeviceInfo*) prototype;
- (BDeviceInfoBuilder*) toBuilder;

+ (BDeviceInfo*) parseFromData:(NSData*) data;
+ (BDeviceInfo*) parseFromData:(NSData*) data extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (BDeviceInfo*) parseFromInputStream:(NSInputStream*) input;
+ (BDeviceInfo*) parseFromInputStream:(NSInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
+ (BDeviceInfo*) parseFromCodedInputStream:(PBCodedInputStream*) input;
+ (BDeviceInfo*) parseFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;
@end

@interface BDeviceInfoBuilder : PBGeneratedMessageBuilder {
@private
  BDeviceInfo* resultDeviceInfo;
}

- (BDeviceInfo*) defaultInstance;

- (BDeviceInfoBuilder*) clear;
- (BDeviceInfoBuilder*) clone;

- (BDeviceInfo*) build;
- (BDeviceInfo*) buildPartial;

- (BDeviceInfoBuilder*) mergeFrom:(BDeviceInfo*) other;
- (BDeviceInfoBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input;
- (BDeviceInfoBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry;

- (BOOL) hasVendorUID;
- (NSString*) vendorUID;
- (BDeviceInfoBuilder*) setVendorUID:(NSString*) value;
- (BDeviceInfoBuilder*) clearVendorUID;

- (BOOL) hasAdvertisingUID;
- (NSString*) advertisingUID;
- (BDeviceInfoBuilder*) setAdvertisingUID:(NSString*) value;
- (BDeviceInfoBuilder*) clearAdvertisingUID;

- (BOOL) hasPlatformType;
- (BPlatformType) platformType;
- (BDeviceInfoBuilder*) setPlatformType:(BPlatformType) value;
- (BDeviceInfoBuilder*) clearPlatformType;

- (BOOL) hasModelName;
- (NSString*) modelName;
- (BDeviceInfoBuilder*) setModelName:(NSString*) value;
- (BDeviceInfoBuilder*) clearModelName;

- (BOOL) hasSystemVersion;
- (NSString*) systemVersion;
- (BDeviceInfoBuilder*) setSystemVersion:(NSString*) value;
- (BDeviceInfoBuilder*) clearSystemVersion;

- (BOOL) hasLanguage;
- (NSString*) language;
- (BDeviceInfoBuilder*) setLanguage:(NSString*) value;
- (BDeviceInfoBuilder*) clearLanguage;

- (BOOL) hasTimezone;
- (NSString*) timezone;
- (BDeviceInfoBuilder*) setTimezone:(NSString*) value;
- (BDeviceInfoBuilder*) clearTimezone;

- (BOOL) hasPhoneCountryCode;
- (NSString*) phoneCountryCode;
- (BDeviceInfoBuilder*) setPhoneCountryCode:(NSString*) value;
- (BDeviceInfoBuilder*) clearPhoneCountryCode;

- (BOOL) hasScreenSize;
- (BSize*) screenSize;
- (BDeviceInfoBuilder*) setScreenSize:(BSize*) value;
- (BDeviceInfoBuilder*) setScreenSizeBuilder:(BSizeBuilder*) builderForValue;
- (BDeviceInfoBuilder*) mergeScreenSize:(BSize*) value;
- (BDeviceInfoBuilder*) clearScreenSize;

- (BOOL) hasScreenScale;
- (Float32) screenScale;
- (BDeviceInfoBuilder*) setScreenScale:(Float32) value;
- (BDeviceInfoBuilder*) clearScreenScale;

- (BOOL) hasAppID;
- (NSString*) appID;
- (BDeviceInfoBuilder*) setAppID:(NSString*) value;
- (BDeviceInfoBuilder*) clearAppID;

- (BOOL) hasAppVersion;
- (NSString*) appVersion;
- (BDeviceInfoBuilder*) setAppVersion:(NSString*) value;
- (BDeviceInfoBuilder*) clearAppVersion;

- (BOOL) hasNotificationToken;
- (NSString*) notificationToken;
- (BDeviceInfoBuilder*) setNotificationToken:(NSString*) value;
- (BDeviceInfoBuilder*) clearNotificationToken;

- (BOOL) hasLastContentRefreshDeprecated;
- (BTimestamp*) lastContentRefreshDeprecated;
- (BDeviceInfoBuilder*) setLastContentRefreshDeprecated:(BTimestamp*) value;
- (BDeviceInfoBuilder*) setLastContentRefreshDeprecatedBuilder:(BTimestampBuilder*) builderForValue;
- (BDeviceInfoBuilder*) mergeLastContentRefreshDeprecated:(BTimestamp*) value;
- (BDeviceInfoBuilder*) clearLastContentRefreshDeprecated;

- (NSMutableArray *)userTags;
- (NSString*)userTagsAtIndex:(NSUInteger)index;
- (BDeviceInfoBuilder *)addUserTags:(NSString*)value;
- (BDeviceInfoBuilder *)setUserTagsArray:(NSArray *)array;
- (BDeviceInfoBuilder *)clearUserTags;

- (BOOL) hasDeviceUDID;
- (NSString*) deviceUDID;
- (BDeviceInfoBuilder*) setDeviceUDID:(NSString*) value;
- (BDeviceInfoBuilder*) clearDeviceUDID;

- (BOOL) hasAppIsReleaseVersion;
- (BOOL) appIsReleaseVersion;
- (BDeviceInfoBuilder*) setAppIsReleaseVersion:(BOOL) value;
- (BDeviceInfoBuilder*) clearAppIsReleaseVersion;

- (BOOL) hasColorDepth;
- (Float32) colorDepth;
- (BDeviceInfoBuilder*) setColorDepth:(Float32) value;
- (BDeviceInfoBuilder*) clearColorDepth;

- (BOOL) hasIPAddress;
- (NSString*) iPAddress;
- (BDeviceInfoBuilder*) setIPAddress:(NSString*) value;
- (BDeviceInfoBuilder*) clearIPAddress;

- (BOOL) hasSystemBuildVersion;
- (NSString*) systemBuildVersion;
- (BDeviceInfoBuilder*) setSystemBuildVersion:(NSString*) value;
- (BDeviceInfoBuilder*) clearSystemBuildVersion;

- (BOOL) hasLocalIPAddress;
- (NSString*) localIPAddress;
- (BDeviceInfoBuilder*) setLocalIPAddress:(NSString*) value;
- (BDeviceInfoBuilder*) clearLocalIPAddress;
@end


// @@protoc_insertion_point(global_scope)
