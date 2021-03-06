// Generated by the protocol buffer compiler.  DO NOT EDIT!

#import "Device.pb.h"
// @@protoc_insertion_point(imports)

@implementation BDeviceRoot
static PBExtensionRegistry* extensionRegistry = nil;
+ (PBExtensionRegistry*) extensionRegistry {
  return extensionRegistry;
}

+ (void) initialize {
  if (self == [BDeviceRoot class]) {
    PBMutableExtensionRegistry* registry = [PBMutableExtensionRegistry registry];
    [self registerAllExtensions:registry];
    [ObjectivecDescriptorRoot registerAllExtensions:registry];
    [BTypesRoot registerAllExtensions:registry];
    extensionRegistry = registry;
  }
}
+ (void) registerAllExtensions:(PBMutableExtensionRegistry*) registry {
}
@end

BOOL BPlatformTypeIsValidValue(BPlatformType value) {
  switch (value) {
    case BPlatformTypePTUnknown:
    case BPlatformTypePTiOS:
    case BPlatformTypePTAndroid:
    case BPlatformTypePTWeb:
      return YES;
    default:
      return NO;
  }
}
NSString *NSStringFromBPlatformType(BPlatformType value) {
  switch (value) {
    case BPlatformTypePTUnknown:
      return @"BPlatformTypePTUnknown";
    case BPlatformTypePTiOS:
      return @"BPlatformTypePTiOS";
    case BPlatformTypePTAndroid:
      return @"BPlatformTypePTAndroid";
    case BPlatformTypePTWeb:
      return @"BPlatformTypePTWeb";
    default:
      return nil;
  }
}

@interface BDeviceInfo ()
@property (strong) NSString* vendorUID;
@property (strong) NSString* advertisingUID;
@property BPlatformType platformType;
@property (strong) NSString* modelName;
@property (strong) NSString* systemVersion;
@property (strong) NSString* language;
@property (strong) NSString* timezone;
@property (strong) NSString* phoneCountryCode;
@property (strong) BSize* screenSize;
@property Float32 screenScale;
@property (strong) NSString* appID;
@property (strong) NSString* appVersion;
@property (strong) NSString* notificationToken;
@property BOOL appIsReleaseVersion;
@property (strong) NSMutableArray * userTagsArray;
@property (strong) NSString* deviceUDID;
@property Float32 colorDepth;
@property (strong) NSString* iPAddress;
@property (strong) NSString* systemBuildVersion;
@property (strong) NSString* localIPAddress;
@end

@implementation BDeviceInfo

- (BOOL) hasVendorUID {
  return !!hasVendorUID_;
}
- (void) setHasVendorUID:(BOOL) _value_ {
  hasVendorUID_ = !!_value_;
}
@synthesize vendorUID;
- (BOOL) hasAdvertisingUID {
  return !!hasAdvertisingUID_;
}
- (void) setHasAdvertisingUID:(BOOL) _value_ {
  hasAdvertisingUID_ = !!_value_;
}
@synthesize advertisingUID;
- (BOOL) hasPlatformType {
  return !!hasPlatformType_;
}
- (void) setHasPlatformType:(BOOL) _value_ {
  hasPlatformType_ = !!_value_;
}
@synthesize platformType;
- (BOOL) hasModelName {
  return !!hasModelName_;
}
- (void) setHasModelName:(BOOL) _value_ {
  hasModelName_ = !!_value_;
}
@synthesize modelName;
- (BOOL) hasSystemVersion {
  return !!hasSystemVersion_;
}
- (void) setHasSystemVersion:(BOOL) _value_ {
  hasSystemVersion_ = !!_value_;
}
@synthesize systemVersion;
- (BOOL) hasLanguage {
  return !!hasLanguage_;
}
- (void) setHasLanguage:(BOOL) _value_ {
  hasLanguage_ = !!_value_;
}
@synthesize language;
- (BOOL) hasTimezone {
  return !!hasTimezone_;
}
- (void) setHasTimezone:(BOOL) _value_ {
  hasTimezone_ = !!_value_;
}
@synthesize timezone;
- (BOOL) hasPhoneCountryCode {
  return !!hasPhoneCountryCode_;
}
- (void) setHasPhoneCountryCode:(BOOL) _value_ {
  hasPhoneCountryCode_ = !!_value_;
}
@synthesize phoneCountryCode;
- (BOOL) hasScreenSize {
  return !!hasScreenSize_;
}
- (void) setHasScreenSize:(BOOL) _value_ {
  hasScreenSize_ = !!_value_;
}
@synthesize screenSize;
- (BOOL) hasScreenScale {
  return !!hasScreenScale_;
}
- (void) setHasScreenScale:(BOOL) _value_ {
  hasScreenScale_ = !!_value_;
}
@synthesize screenScale;
- (BOOL) hasAppID {
  return !!hasAppID_;
}
- (void) setHasAppID:(BOOL) _value_ {
  hasAppID_ = !!_value_;
}
@synthesize appID;
- (BOOL) hasAppVersion {
  return !!hasAppVersion_;
}
- (void) setHasAppVersion:(BOOL) _value_ {
  hasAppVersion_ = !!_value_;
}
@synthesize appVersion;
- (BOOL) hasNotificationToken {
  return !!hasNotificationToken_;
}
- (void) setHasNotificationToken:(BOOL) _value_ {
  hasNotificationToken_ = !!_value_;
}
@synthesize notificationToken;
- (BOOL) hasAppIsReleaseVersion {
  return !!hasAppIsReleaseVersion_;
}
- (void) setHasAppIsReleaseVersion:(BOOL) _value_ {
  hasAppIsReleaseVersion_ = !!_value_;
}
- (BOOL) appIsReleaseVersion {
  return !!appIsReleaseVersion_;
}
- (void) setAppIsReleaseVersion:(BOOL) _value_ {
  appIsReleaseVersion_ = !!_value_;
}
@synthesize userTagsArray;
@dynamic userTags;
- (BOOL) hasDeviceUDID {
  return !!hasDeviceUDID_;
}
- (void) setHasDeviceUDID:(BOOL) _value_ {
  hasDeviceUDID_ = !!_value_;
}
@synthesize deviceUDID;
- (BOOL) hasColorDepth {
  return !!hasColorDepth_;
}
- (void) setHasColorDepth:(BOOL) _value_ {
  hasColorDepth_ = !!_value_;
}
@synthesize colorDepth;
- (BOOL) hasIPAddress {
  return !!hasIPAddress_;
}
- (void) setHasIPAddress:(BOOL) _value_ {
  hasIPAddress_ = !!_value_;
}
@synthesize iPAddress;
- (BOOL) hasSystemBuildVersion {
  return !!hasSystemBuildVersion_;
}
- (void) setHasSystemBuildVersion:(BOOL) _value_ {
  hasSystemBuildVersion_ = !!_value_;
}
@synthesize systemBuildVersion;
- (BOOL) hasLocalIPAddress {
  return !!hasLocalIPAddress_;
}
- (void) setHasLocalIPAddress:(BOOL) _value_ {
  hasLocalIPAddress_ = !!_value_;
}
@synthesize localIPAddress;
- (instancetype) init {
  if ((self = [super init])) {
    self.vendorUID = @"";
    self.advertisingUID = @"";
    self.platformType = BPlatformTypePTUnknown;
    self.modelName = @"";
    self.systemVersion = @"";
    self.language = @"";
    self.timezone = @"";
    self.phoneCountryCode = @"";
    self.screenSize = [BSize defaultInstance];
    self.screenScale = 1;
    self.appID = @"";
    self.appVersion = @"";
    self.notificationToken = @"";
    self.appIsReleaseVersion = NO;
    self.deviceUDID = @"";
    self.colorDepth = 0;
    self.iPAddress = @"";
    self.systemBuildVersion = @"";
    self.localIPAddress = @"";
  }
  return self;
}
static BDeviceInfo* defaultBDeviceInfoInstance = nil;
+ (void) initialize {
  if (self == [BDeviceInfo class]) {
    defaultBDeviceInfoInstance = [[BDeviceInfo alloc] init];
  }
}
+ (instancetype) defaultInstance {
  return defaultBDeviceInfoInstance;
}
- (instancetype) defaultInstance {
  return defaultBDeviceInfoInstance;
}
- (NSArray *)userTags {
  return userTagsArray;
}
- (NSString*)userTagsAtIndex:(NSUInteger)index {
  return [userTagsArray objectAtIndex:index];
}
- (BOOL) isInitialized {
  return YES;
}
- (void) writeToCodedOutputStream:(PBCodedOutputStream*) output {
  if (self.hasVendorUID) {
    [output writeString:1 value:self.vendorUID];
  }
  if (self.hasAdvertisingUID) {
    [output writeString:2 value:self.advertisingUID];
  }
  if (self.hasPlatformType) {
    [output writeEnum:3 value:self.platformType];
  }
  if (self.hasModelName) {
    [output writeString:4 value:self.modelName];
  }
  if (self.hasSystemVersion) {
    [output writeString:5 value:self.systemVersion];
  }
  if (self.hasLanguage) {
    [output writeString:6 value:self.language];
  }
  if (self.hasTimezone) {
    [output writeString:7 value:self.timezone];
  }
  if (self.hasPhoneCountryCode) {
    [output writeString:8 value:self.phoneCountryCode];
  }
  if (self.hasScreenSize) {
    [output writeMessage:9 value:self.screenSize];
  }
  if (self.hasScreenScale) {
    [output writeFloat:10 value:self.screenScale];
  }
  if (self.hasAppID) {
    [output writeString:11 value:self.appID];
  }
  if (self.hasAppVersion) {
    [output writeString:12 value:self.appVersion];
  }
  if (self.hasNotificationToken) {
    [output writeString:13 value:self.notificationToken];
  }
  if (self.hasAppIsReleaseVersion) {
    [output writeBool:14 value:self.appIsReleaseVersion];
  }
  [self.userTagsArray enumerateObjectsUsingBlock:^(NSString *element, NSUInteger idx, BOOL *stop) {
    [output writeString:15 value:element];
  }];
  if (self.hasDeviceUDID) {
    [output writeString:16 value:self.deviceUDID];
  }
  if (self.hasColorDepth) {
    [output writeFloat:17 value:self.colorDepth];
  }
  if (self.hasIPAddress) {
    [output writeString:18 value:self.iPAddress];
  }
  if (self.hasSystemBuildVersion) {
    [output writeString:19 value:self.systemBuildVersion];
  }
  if (self.hasLocalIPAddress) {
    [output writeString:20 value:self.localIPAddress];
  }
  [self.unknownFields writeToCodedOutputStream:output];
}
- (SInt32) serializedSize {
  __block SInt32 size_ = memoizedSerializedSize;
  if (size_ != -1) {
    return size_;
  }

  size_ = 0;
  if (self.hasVendorUID) {
    size_ += computeStringSize(1, self.vendorUID);
  }
  if (self.hasAdvertisingUID) {
    size_ += computeStringSize(2, self.advertisingUID);
  }
  if (self.hasPlatformType) {
    size_ += computeEnumSize(3, self.platformType);
  }
  if (self.hasModelName) {
    size_ += computeStringSize(4, self.modelName);
  }
  if (self.hasSystemVersion) {
    size_ += computeStringSize(5, self.systemVersion);
  }
  if (self.hasLanguage) {
    size_ += computeStringSize(6, self.language);
  }
  if (self.hasTimezone) {
    size_ += computeStringSize(7, self.timezone);
  }
  if (self.hasPhoneCountryCode) {
    size_ += computeStringSize(8, self.phoneCountryCode);
  }
  if (self.hasScreenSize) {
    size_ += computeMessageSize(9, self.screenSize);
  }
  if (self.hasScreenScale) {
    size_ += computeFloatSize(10, self.screenScale);
  }
  if (self.hasAppID) {
    size_ += computeStringSize(11, self.appID);
  }
  if (self.hasAppVersion) {
    size_ += computeStringSize(12, self.appVersion);
  }
  if (self.hasNotificationToken) {
    size_ += computeStringSize(13, self.notificationToken);
  }
  if (self.hasAppIsReleaseVersion) {
    size_ += computeBoolSize(14, self.appIsReleaseVersion);
  }
  {
    __block SInt32 dataSize = 0;
    const NSUInteger count = self.userTagsArray.count;
    [self.userTagsArray enumerateObjectsUsingBlock:^(NSString *element, NSUInteger idx, BOOL *stop) {
      dataSize += computeStringSizeNoTag(element);
    }];
    size_ += dataSize;
    size_ += (SInt32)(1 * count);
  }
  if (self.hasDeviceUDID) {
    size_ += computeStringSize(16, self.deviceUDID);
  }
  if (self.hasColorDepth) {
    size_ += computeFloatSize(17, self.colorDepth);
  }
  if (self.hasIPAddress) {
    size_ += computeStringSize(18, self.iPAddress);
  }
  if (self.hasSystemBuildVersion) {
    size_ += computeStringSize(19, self.systemBuildVersion);
  }
  if (self.hasLocalIPAddress) {
    size_ += computeStringSize(20, self.localIPAddress);
  }
  size_ += self.unknownFields.serializedSize;
  memoizedSerializedSize = size_;
  return size_;
}
+ (BDeviceInfo*) parseFromData:(NSData*) data {
  return (BDeviceInfo*)[[[BDeviceInfo builder] mergeFromData:data] build];
}
+ (BDeviceInfo*) parseFromData:(NSData*) data extensionRegistry:(PBExtensionRegistry*) extensionRegistry {
  return (BDeviceInfo*)[[[BDeviceInfo builder] mergeFromData:data extensionRegistry:extensionRegistry] build];
}
+ (BDeviceInfo*) parseFromInputStream:(NSInputStream*) input {
  return (BDeviceInfo*)[[[BDeviceInfo builder] mergeFromInputStream:input] build];
}
+ (BDeviceInfo*) parseFromInputStream:(NSInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry {
  return (BDeviceInfo*)[[[BDeviceInfo builder] mergeFromInputStream:input extensionRegistry:extensionRegistry] build];
}
+ (BDeviceInfo*) parseFromCodedInputStream:(PBCodedInputStream*) input {
  return (BDeviceInfo*)[[[BDeviceInfo builder] mergeFromCodedInputStream:input] build];
}
+ (BDeviceInfo*) parseFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry {
  return (BDeviceInfo*)[[[BDeviceInfo builder] mergeFromCodedInputStream:input extensionRegistry:extensionRegistry] build];
}
+ (BDeviceInfoBuilder*) builder {
  return [[BDeviceInfoBuilder alloc] init];
}
+ (BDeviceInfoBuilder*) builderWithPrototype:(BDeviceInfo*) prototype {
  return [[BDeviceInfo builder] mergeFrom:prototype];
}
- (BDeviceInfoBuilder*) builder {
  return [BDeviceInfo builder];
}
- (BDeviceInfoBuilder*) toBuilder {
  return [BDeviceInfo builderWithPrototype:self];
}
- (void) writeDescriptionTo:(NSMutableString*) output withIndent:(NSString*) indent {
  if (self.hasVendorUID) {
    [output appendFormat:@"%@%@: %@\n", indent, @"vendorUID", self.vendorUID];
  }
  if (self.hasAdvertisingUID) {
    [output appendFormat:@"%@%@: %@\n", indent, @"advertisingUID", self.advertisingUID];
  }
  if (self.hasPlatformType) {
    [output appendFormat:@"%@%@: %@\n", indent, @"platformType", NSStringFromBPlatformType(self.platformType)];
  }
  if (self.hasModelName) {
    [output appendFormat:@"%@%@: %@\n", indent, @"modelName", self.modelName];
  }
  if (self.hasSystemVersion) {
    [output appendFormat:@"%@%@: %@\n", indent, @"systemVersion", self.systemVersion];
  }
  if (self.hasLanguage) {
    [output appendFormat:@"%@%@: %@\n", indent, @"language", self.language];
  }
  if (self.hasTimezone) {
    [output appendFormat:@"%@%@: %@\n", indent, @"timezone", self.timezone];
  }
  if (self.hasPhoneCountryCode) {
    [output appendFormat:@"%@%@: %@\n", indent, @"phoneCountryCode", self.phoneCountryCode];
  }
  if (self.hasScreenSize) {
    [output appendFormat:@"%@%@ {\n", indent, @"screenSize"];
    [self.screenSize writeDescriptionTo:output
                         withIndent:[NSString stringWithFormat:@"%@  ", indent]];
    [output appendFormat:@"%@}\n", indent];
  }
  if (self.hasScreenScale) {
    [output appendFormat:@"%@%@: %@\n", indent, @"screenScale", [NSNumber numberWithFloat:self.screenScale]];
  }
  if (self.hasAppID) {
    [output appendFormat:@"%@%@: %@\n", indent, @"appID", self.appID];
  }
  if (self.hasAppVersion) {
    [output appendFormat:@"%@%@: %@\n", indent, @"appVersion", self.appVersion];
  }
  if (self.hasNotificationToken) {
    [output appendFormat:@"%@%@: %@\n", indent, @"notificationToken", self.notificationToken];
  }
  if (self.hasAppIsReleaseVersion) {
    [output appendFormat:@"%@%@: %@\n", indent, @"appIsReleaseVersion", [NSNumber numberWithBool:self.appIsReleaseVersion]];
  }
  [self.userTagsArray enumerateObjectsUsingBlock:^(id obj, NSUInteger idx, BOOL *stop) {
    [output appendFormat:@"%@%@: %@\n", indent, @"userTags", obj];
  }];
  if (self.hasDeviceUDID) {
    [output appendFormat:@"%@%@: %@\n", indent, @"deviceUDID", self.deviceUDID];
  }
  if (self.hasColorDepth) {
    [output appendFormat:@"%@%@: %@\n", indent, @"colorDepth", [NSNumber numberWithFloat:self.colorDepth]];
  }
  if (self.hasIPAddress) {
    [output appendFormat:@"%@%@: %@\n", indent, @"iPAddress", self.iPAddress];
  }
  if (self.hasSystemBuildVersion) {
    [output appendFormat:@"%@%@: %@\n", indent, @"systemBuildVersion", self.systemBuildVersion];
  }
  if (self.hasLocalIPAddress) {
    [output appendFormat:@"%@%@: %@\n", indent, @"localIPAddress", self.localIPAddress];
  }
  [self.unknownFields writeDescriptionTo:output withIndent:indent];
}
- (void) storeInDictionary:(NSMutableDictionary *)dictionary {
  if (self.hasVendorUID) {
    [dictionary setObject: self.vendorUID forKey: @"vendorUID"];
  }
  if (self.hasAdvertisingUID) {
    [dictionary setObject: self.advertisingUID forKey: @"advertisingUID"];
  }
  if (self.hasPlatformType) {
    [dictionary setObject: @(self.platformType) forKey: @"platformType"];
  }
  if (self.hasModelName) {
    [dictionary setObject: self.modelName forKey: @"modelName"];
  }
  if (self.hasSystemVersion) {
    [dictionary setObject: self.systemVersion forKey: @"systemVersion"];
  }
  if (self.hasLanguage) {
    [dictionary setObject: self.language forKey: @"language"];
  }
  if (self.hasTimezone) {
    [dictionary setObject: self.timezone forKey: @"timezone"];
  }
  if (self.hasPhoneCountryCode) {
    [dictionary setObject: self.phoneCountryCode forKey: @"phoneCountryCode"];
  }
  if (self.hasScreenSize) {
   NSMutableDictionary *messageDictionary = [NSMutableDictionary dictionary]; 
   [self.screenSize storeInDictionary:messageDictionary];
   [dictionary setObject:[NSDictionary dictionaryWithDictionary:messageDictionary] forKey:@"screenSize"];
  }
  if (self.hasScreenScale) {
    [dictionary setObject: [NSNumber numberWithFloat:self.screenScale] forKey: @"screenScale"];
  }
  if (self.hasAppID) {
    [dictionary setObject: self.appID forKey: @"appID"];
  }
  if (self.hasAppVersion) {
    [dictionary setObject: self.appVersion forKey: @"appVersion"];
  }
  if (self.hasNotificationToken) {
    [dictionary setObject: self.notificationToken forKey: @"notificationToken"];
  }
  if (self.hasAppIsReleaseVersion) {
    [dictionary setObject: [NSNumber numberWithBool:self.appIsReleaseVersion] forKey: @"appIsReleaseVersion"];
  }
  [dictionary setObject:self.userTags forKey: @"userTags"];
  if (self.hasDeviceUDID) {
    [dictionary setObject: self.deviceUDID forKey: @"deviceUDID"];
  }
  if (self.hasColorDepth) {
    [dictionary setObject: [NSNumber numberWithFloat:self.colorDepth] forKey: @"colorDepth"];
  }
  if (self.hasIPAddress) {
    [dictionary setObject: self.iPAddress forKey: @"iPAddress"];
  }
  if (self.hasSystemBuildVersion) {
    [dictionary setObject: self.systemBuildVersion forKey: @"systemBuildVersion"];
  }
  if (self.hasLocalIPAddress) {
    [dictionary setObject: self.localIPAddress forKey: @"localIPAddress"];
  }
  [self.unknownFields storeInDictionary:dictionary];
}
- (BOOL) isEqual:(id)other {
  if (other == self) {
    return YES;
  }
  if (![other isKindOfClass:[BDeviceInfo class]]) {
    return NO;
  }
  BDeviceInfo *otherMessage = other;
  return
      self.hasVendorUID == otherMessage.hasVendorUID &&
      (!self.hasVendorUID || [self.vendorUID isEqual:otherMessage.vendorUID]) &&
      self.hasAdvertisingUID == otherMessage.hasAdvertisingUID &&
      (!self.hasAdvertisingUID || [self.advertisingUID isEqual:otherMessage.advertisingUID]) &&
      self.hasPlatformType == otherMessage.hasPlatformType &&
      (!self.hasPlatformType || self.platformType == otherMessage.platformType) &&
      self.hasModelName == otherMessage.hasModelName &&
      (!self.hasModelName || [self.modelName isEqual:otherMessage.modelName]) &&
      self.hasSystemVersion == otherMessage.hasSystemVersion &&
      (!self.hasSystemVersion || [self.systemVersion isEqual:otherMessage.systemVersion]) &&
      self.hasLanguage == otherMessage.hasLanguage &&
      (!self.hasLanguage || [self.language isEqual:otherMessage.language]) &&
      self.hasTimezone == otherMessage.hasTimezone &&
      (!self.hasTimezone || [self.timezone isEqual:otherMessage.timezone]) &&
      self.hasPhoneCountryCode == otherMessage.hasPhoneCountryCode &&
      (!self.hasPhoneCountryCode || [self.phoneCountryCode isEqual:otherMessage.phoneCountryCode]) &&
      self.hasScreenSize == otherMessage.hasScreenSize &&
      (!self.hasScreenSize || [self.screenSize isEqual:otherMessage.screenSize]) &&
      self.hasScreenScale == otherMessage.hasScreenScale &&
      (!self.hasScreenScale || self.screenScale == otherMessage.screenScale) &&
      self.hasAppID == otherMessage.hasAppID &&
      (!self.hasAppID || [self.appID isEqual:otherMessage.appID]) &&
      self.hasAppVersion == otherMessage.hasAppVersion &&
      (!self.hasAppVersion || [self.appVersion isEqual:otherMessage.appVersion]) &&
      self.hasNotificationToken == otherMessage.hasNotificationToken &&
      (!self.hasNotificationToken || [self.notificationToken isEqual:otherMessage.notificationToken]) &&
      self.hasAppIsReleaseVersion == otherMessage.hasAppIsReleaseVersion &&
      (!self.hasAppIsReleaseVersion || self.appIsReleaseVersion == otherMessage.appIsReleaseVersion) &&
      [self.userTagsArray isEqualToArray:otherMessage.userTagsArray] &&
      self.hasDeviceUDID == otherMessage.hasDeviceUDID &&
      (!self.hasDeviceUDID || [self.deviceUDID isEqual:otherMessage.deviceUDID]) &&
      self.hasColorDepth == otherMessage.hasColorDepth &&
      (!self.hasColorDepth || self.colorDepth == otherMessage.colorDepth) &&
      self.hasIPAddress == otherMessage.hasIPAddress &&
      (!self.hasIPAddress || [self.iPAddress isEqual:otherMessage.iPAddress]) &&
      self.hasSystemBuildVersion == otherMessage.hasSystemBuildVersion &&
      (!self.hasSystemBuildVersion || [self.systemBuildVersion isEqual:otherMessage.systemBuildVersion]) &&
      self.hasLocalIPAddress == otherMessage.hasLocalIPAddress &&
      (!self.hasLocalIPAddress || [self.localIPAddress isEqual:otherMessage.localIPAddress]) &&
      (self.unknownFields == otherMessage.unknownFields || (self.unknownFields != nil && [self.unknownFields isEqual:otherMessage.unknownFields]));
}
- (NSUInteger) hash {
  __block NSUInteger hashCode = 7;
  if (self.hasVendorUID) {
    hashCode = hashCode * 31 + [self.vendorUID hash];
  }
  if (self.hasAdvertisingUID) {
    hashCode = hashCode * 31 + [self.advertisingUID hash];
  }
  if (self.hasPlatformType) {
    hashCode = hashCode * 31 + self.platformType;
  }
  if (self.hasModelName) {
    hashCode = hashCode * 31 + [self.modelName hash];
  }
  if (self.hasSystemVersion) {
    hashCode = hashCode * 31 + [self.systemVersion hash];
  }
  if (self.hasLanguage) {
    hashCode = hashCode * 31 + [self.language hash];
  }
  if (self.hasTimezone) {
    hashCode = hashCode * 31 + [self.timezone hash];
  }
  if (self.hasPhoneCountryCode) {
    hashCode = hashCode * 31 + [self.phoneCountryCode hash];
  }
  if (self.hasScreenSize) {
    hashCode = hashCode * 31 + [self.screenSize hash];
  }
  if (self.hasScreenScale) {
    hashCode = hashCode * 31 + [[NSNumber numberWithFloat:self.screenScale] hash];
  }
  if (self.hasAppID) {
    hashCode = hashCode * 31 + [self.appID hash];
  }
  if (self.hasAppVersion) {
    hashCode = hashCode * 31 + [self.appVersion hash];
  }
  if (self.hasNotificationToken) {
    hashCode = hashCode * 31 + [self.notificationToken hash];
  }
  if (self.hasAppIsReleaseVersion) {
    hashCode = hashCode * 31 + [[NSNumber numberWithBool:self.appIsReleaseVersion] hash];
  }
  [self.userTagsArray enumerateObjectsUsingBlock:^(NSString *element, NSUInteger idx, BOOL *stop) {
    hashCode = hashCode * 31 + [element hash];
  }];
  if (self.hasDeviceUDID) {
    hashCode = hashCode * 31 + [self.deviceUDID hash];
  }
  if (self.hasColorDepth) {
    hashCode = hashCode * 31 + [[NSNumber numberWithFloat:self.colorDepth] hash];
  }
  if (self.hasIPAddress) {
    hashCode = hashCode * 31 + [self.iPAddress hash];
  }
  if (self.hasSystemBuildVersion) {
    hashCode = hashCode * 31 + [self.systemBuildVersion hash];
  }
  if (self.hasLocalIPAddress) {
    hashCode = hashCode * 31 + [self.localIPAddress hash];
  }
  hashCode = hashCode * 31 + [self.unknownFields hash];
  return hashCode;
}
@end

@interface BDeviceInfoBuilder()
@property (strong) BDeviceInfo* resultDeviceInfo;
@end

@implementation BDeviceInfoBuilder
@synthesize resultDeviceInfo;
- (instancetype) init {
  if ((self = [super init])) {
    self.resultDeviceInfo = [[BDeviceInfo alloc] init];
  }
  return self;
}
- (PBGeneratedMessage*) internalGetResult {
  return resultDeviceInfo;
}
- (BDeviceInfoBuilder*) clear {
  self.resultDeviceInfo = [[BDeviceInfo alloc] init];
  return self;
}
- (BDeviceInfoBuilder*) clone {
  return [BDeviceInfo builderWithPrototype:resultDeviceInfo];
}
- (BDeviceInfo*) defaultInstance {
  return [BDeviceInfo defaultInstance];
}
- (BDeviceInfo*) build {
  [self checkInitialized];
  return [self buildPartial];
}
- (BDeviceInfo*) buildPartial {
  BDeviceInfo* returnMe = resultDeviceInfo;
  self.resultDeviceInfo = nil;
  return returnMe;
}
- (BDeviceInfoBuilder*) mergeFrom:(BDeviceInfo*) other {
  if (other == [BDeviceInfo defaultInstance]) {
    return self;
  }
  if (other.hasVendorUID) {
    [self setVendorUID:other.vendorUID];
  }
  if (other.hasAdvertisingUID) {
    [self setAdvertisingUID:other.advertisingUID];
  }
  if (other.hasPlatformType) {
    [self setPlatformType:other.platformType];
  }
  if (other.hasModelName) {
    [self setModelName:other.modelName];
  }
  if (other.hasSystemVersion) {
    [self setSystemVersion:other.systemVersion];
  }
  if (other.hasLanguage) {
    [self setLanguage:other.language];
  }
  if (other.hasTimezone) {
    [self setTimezone:other.timezone];
  }
  if (other.hasPhoneCountryCode) {
    [self setPhoneCountryCode:other.phoneCountryCode];
  }
  if (other.hasScreenSize) {
    [self mergeScreenSize:other.screenSize];
  }
  if (other.hasScreenScale) {
    [self setScreenScale:other.screenScale];
  }
  if (other.hasAppID) {
    [self setAppID:other.appID];
  }
  if (other.hasAppVersion) {
    [self setAppVersion:other.appVersion];
  }
  if (other.hasNotificationToken) {
    [self setNotificationToken:other.notificationToken];
  }
  if (other.hasAppIsReleaseVersion) {
    [self setAppIsReleaseVersion:other.appIsReleaseVersion];
  }
  if (other.userTagsArray.count > 0) {
    if (resultDeviceInfo.userTagsArray == nil) {
      resultDeviceInfo.userTagsArray = [[NSMutableArray alloc] initWithArray:other.userTagsArray];
    } else {
      [resultDeviceInfo.userTagsArray addObjectsFromArray:other.userTagsArray];
    }
  }
  if (other.hasDeviceUDID) {
    [self setDeviceUDID:other.deviceUDID];
  }
  if (other.hasColorDepth) {
    [self setColorDepth:other.colorDepth];
  }
  if (other.hasIPAddress) {
    [self setIPAddress:other.iPAddress];
  }
  if (other.hasSystemBuildVersion) {
    [self setSystemBuildVersion:other.systemBuildVersion];
  }
  if (other.hasLocalIPAddress) {
    [self setLocalIPAddress:other.localIPAddress];
  }
  [self mergeUnknownFields:other.unknownFields];
  return self;
}
- (BDeviceInfoBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input {
  return [self mergeFromCodedInputStream:input extensionRegistry:[PBExtensionRegistry emptyRegistry]];
}
- (BDeviceInfoBuilder*) mergeFromCodedInputStream:(PBCodedInputStream*) input extensionRegistry:(PBExtensionRegistry*) extensionRegistry {
  PBUnknownFieldSetBuilder* unknownFields = [PBUnknownFieldSet builderWithUnknownFields:self.unknownFields];
  while (YES) {
    SInt32 tag = [input readTag];
    switch (tag) {
      case 0:
        [self setUnknownFields:[unknownFields build]];
        return self;
      default: {
        if (![self parseUnknownField:input unknownFields:unknownFields extensionRegistry:extensionRegistry tag:tag]) {
          [self setUnknownFields:[unknownFields build]];
          return self;
        }
        break;
      }
      case 10: {
        [self setVendorUID:[input readString]];
        break;
      }
      case 18: {
        [self setAdvertisingUID:[input readString]];
        break;
      }
      case 24: {
        BPlatformType value = (BPlatformType)[input readEnum];
        if (BPlatformTypeIsValidValue(value)) {
          [self setPlatformType:value];
        } else {
          [unknownFields mergeVarintField:3 value:value];
        }
        break;
      }
      case 34: {
        [self setModelName:[input readString]];
        break;
      }
      case 42: {
        [self setSystemVersion:[input readString]];
        break;
      }
      case 50: {
        [self setLanguage:[input readString]];
        break;
      }
      case 58: {
        [self setTimezone:[input readString]];
        break;
      }
      case 66: {
        [self setPhoneCountryCode:[input readString]];
        break;
      }
      case 74: {
        BSizeBuilder* subBuilder = [BSize builder];
        if (self.hasScreenSize) {
          [subBuilder mergeFrom:self.screenSize];
        }
        [input readMessage:subBuilder extensionRegistry:extensionRegistry];
        [self setScreenSize:[subBuilder buildPartial]];
        break;
      }
      case 85: {
        [self setScreenScale:[input readFloat]];
        break;
      }
      case 90: {
        [self setAppID:[input readString]];
        break;
      }
      case 98: {
        [self setAppVersion:[input readString]];
        break;
      }
      case 106: {
        [self setNotificationToken:[input readString]];
        break;
      }
      case 112: {
        [self setAppIsReleaseVersion:[input readBool]];
        break;
      }
      case 122: {
        [self addUserTags:[input readString]];
        break;
      }
      case 130: {
        [self setDeviceUDID:[input readString]];
        break;
      }
      case 141: {
        [self setColorDepth:[input readFloat]];
        break;
      }
      case 146: {
        [self setIPAddress:[input readString]];
        break;
      }
      case 154: {
        [self setSystemBuildVersion:[input readString]];
        break;
      }
      case 162: {
        [self setLocalIPAddress:[input readString]];
        break;
      }
    }
  }
}
- (BOOL) hasVendorUID {
  return resultDeviceInfo.hasVendorUID;
}
- (NSString*) vendorUID {
  return resultDeviceInfo.vendorUID;
}
- (BDeviceInfoBuilder*) setVendorUID:(NSString*) value {
  resultDeviceInfo.hasVendorUID = YES;
  resultDeviceInfo.vendorUID = value;
  return self;
}
- (BDeviceInfoBuilder*) clearVendorUID {
  resultDeviceInfo.hasVendorUID = NO;
  resultDeviceInfo.vendorUID = @"";
  return self;
}
- (BOOL) hasAdvertisingUID {
  return resultDeviceInfo.hasAdvertisingUID;
}
- (NSString*) advertisingUID {
  return resultDeviceInfo.advertisingUID;
}
- (BDeviceInfoBuilder*) setAdvertisingUID:(NSString*) value {
  resultDeviceInfo.hasAdvertisingUID = YES;
  resultDeviceInfo.advertisingUID = value;
  return self;
}
- (BDeviceInfoBuilder*) clearAdvertisingUID {
  resultDeviceInfo.hasAdvertisingUID = NO;
  resultDeviceInfo.advertisingUID = @"";
  return self;
}
- (BOOL) hasPlatformType {
  return resultDeviceInfo.hasPlatformType;
}
- (BPlatformType) platformType {
  return resultDeviceInfo.platformType;
}
- (BDeviceInfoBuilder*) setPlatformType:(BPlatformType) value {
  resultDeviceInfo.hasPlatformType = YES;
  resultDeviceInfo.platformType = value;
  return self;
}
- (BDeviceInfoBuilder*) clearPlatformType {
  resultDeviceInfo.hasPlatformType = NO;
  resultDeviceInfo.platformType = BPlatformTypePTUnknown;
  return self;
}
- (BOOL) hasModelName {
  return resultDeviceInfo.hasModelName;
}
- (NSString*) modelName {
  return resultDeviceInfo.modelName;
}
- (BDeviceInfoBuilder*) setModelName:(NSString*) value {
  resultDeviceInfo.hasModelName = YES;
  resultDeviceInfo.modelName = value;
  return self;
}
- (BDeviceInfoBuilder*) clearModelName {
  resultDeviceInfo.hasModelName = NO;
  resultDeviceInfo.modelName = @"";
  return self;
}
- (BOOL) hasSystemVersion {
  return resultDeviceInfo.hasSystemVersion;
}
- (NSString*) systemVersion {
  return resultDeviceInfo.systemVersion;
}
- (BDeviceInfoBuilder*) setSystemVersion:(NSString*) value {
  resultDeviceInfo.hasSystemVersion = YES;
  resultDeviceInfo.systemVersion = value;
  return self;
}
- (BDeviceInfoBuilder*) clearSystemVersion {
  resultDeviceInfo.hasSystemVersion = NO;
  resultDeviceInfo.systemVersion = @"";
  return self;
}
- (BOOL) hasLanguage {
  return resultDeviceInfo.hasLanguage;
}
- (NSString*) language {
  return resultDeviceInfo.language;
}
- (BDeviceInfoBuilder*) setLanguage:(NSString*) value {
  resultDeviceInfo.hasLanguage = YES;
  resultDeviceInfo.language = value;
  return self;
}
- (BDeviceInfoBuilder*) clearLanguage {
  resultDeviceInfo.hasLanguage = NO;
  resultDeviceInfo.language = @"";
  return self;
}
- (BOOL) hasTimezone {
  return resultDeviceInfo.hasTimezone;
}
- (NSString*) timezone {
  return resultDeviceInfo.timezone;
}
- (BDeviceInfoBuilder*) setTimezone:(NSString*) value {
  resultDeviceInfo.hasTimezone = YES;
  resultDeviceInfo.timezone = value;
  return self;
}
- (BDeviceInfoBuilder*) clearTimezone {
  resultDeviceInfo.hasTimezone = NO;
  resultDeviceInfo.timezone = @"";
  return self;
}
- (BOOL) hasPhoneCountryCode {
  return resultDeviceInfo.hasPhoneCountryCode;
}
- (NSString*) phoneCountryCode {
  return resultDeviceInfo.phoneCountryCode;
}
- (BDeviceInfoBuilder*) setPhoneCountryCode:(NSString*) value {
  resultDeviceInfo.hasPhoneCountryCode = YES;
  resultDeviceInfo.phoneCountryCode = value;
  return self;
}
- (BDeviceInfoBuilder*) clearPhoneCountryCode {
  resultDeviceInfo.hasPhoneCountryCode = NO;
  resultDeviceInfo.phoneCountryCode = @"";
  return self;
}
- (BOOL) hasScreenSize {
  return resultDeviceInfo.hasScreenSize;
}
- (BSize*) screenSize {
  return resultDeviceInfo.screenSize;
}
- (BDeviceInfoBuilder*) setScreenSize:(BSize*) value {
  resultDeviceInfo.hasScreenSize = YES;
  resultDeviceInfo.screenSize = value;
  return self;
}
- (BDeviceInfoBuilder*) setScreenSizeBuilder:(BSizeBuilder*) builderForValue {
  return [self setScreenSize:[builderForValue build]];
}
- (BDeviceInfoBuilder*) mergeScreenSize:(BSize*) value {
  if (resultDeviceInfo.hasScreenSize &&
      resultDeviceInfo.screenSize != [BSize defaultInstance]) {
    resultDeviceInfo.screenSize =
      [[[BSize builderWithPrototype:resultDeviceInfo.screenSize] mergeFrom:value] buildPartial];
  } else {
    resultDeviceInfo.screenSize = value;
  }
  resultDeviceInfo.hasScreenSize = YES;
  return self;
}
- (BDeviceInfoBuilder*) clearScreenSize {
  resultDeviceInfo.hasScreenSize = NO;
  resultDeviceInfo.screenSize = [BSize defaultInstance];
  return self;
}
- (BOOL) hasScreenScale {
  return resultDeviceInfo.hasScreenScale;
}
- (Float32) screenScale {
  return resultDeviceInfo.screenScale;
}
- (BDeviceInfoBuilder*) setScreenScale:(Float32) value {
  resultDeviceInfo.hasScreenScale = YES;
  resultDeviceInfo.screenScale = value;
  return self;
}
- (BDeviceInfoBuilder*) clearScreenScale {
  resultDeviceInfo.hasScreenScale = NO;
  resultDeviceInfo.screenScale = 1;
  return self;
}
- (BOOL) hasAppID {
  return resultDeviceInfo.hasAppID;
}
- (NSString*) appID {
  return resultDeviceInfo.appID;
}
- (BDeviceInfoBuilder*) setAppID:(NSString*) value {
  resultDeviceInfo.hasAppID = YES;
  resultDeviceInfo.appID = value;
  return self;
}
- (BDeviceInfoBuilder*) clearAppID {
  resultDeviceInfo.hasAppID = NO;
  resultDeviceInfo.appID = @"";
  return self;
}
- (BOOL) hasAppVersion {
  return resultDeviceInfo.hasAppVersion;
}
- (NSString*) appVersion {
  return resultDeviceInfo.appVersion;
}
- (BDeviceInfoBuilder*) setAppVersion:(NSString*) value {
  resultDeviceInfo.hasAppVersion = YES;
  resultDeviceInfo.appVersion = value;
  return self;
}
- (BDeviceInfoBuilder*) clearAppVersion {
  resultDeviceInfo.hasAppVersion = NO;
  resultDeviceInfo.appVersion = @"";
  return self;
}
- (BOOL) hasNotificationToken {
  return resultDeviceInfo.hasNotificationToken;
}
- (NSString*) notificationToken {
  return resultDeviceInfo.notificationToken;
}
- (BDeviceInfoBuilder*) setNotificationToken:(NSString*) value {
  resultDeviceInfo.hasNotificationToken = YES;
  resultDeviceInfo.notificationToken = value;
  return self;
}
- (BDeviceInfoBuilder*) clearNotificationToken {
  resultDeviceInfo.hasNotificationToken = NO;
  resultDeviceInfo.notificationToken = @"";
  return self;
}
- (BOOL) hasAppIsReleaseVersion {
  return resultDeviceInfo.hasAppIsReleaseVersion;
}
- (BOOL) appIsReleaseVersion {
  return resultDeviceInfo.appIsReleaseVersion;
}
- (BDeviceInfoBuilder*) setAppIsReleaseVersion:(BOOL) value {
  resultDeviceInfo.hasAppIsReleaseVersion = YES;
  resultDeviceInfo.appIsReleaseVersion = value;
  return self;
}
- (BDeviceInfoBuilder*) clearAppIsReleaseVersion {
  resultDeviceInfo.hasAppIsReleaseVersion = NO;
  resultDeviceInfo.appIsReleaseVersion = NO;
  return self;
}
- (NSMutableArray *)userTags {
  return resultDeviceInfo.userTagsArray;
}
- (NSString*)userTagsAtIndex:(NSUInteger)index {
  return [resultDeviceInfo userTagsAtIndex:index];
}
- (BDeviceInfoBuilder *)addUserTags:(NSString*)value {
  if (resultDeviceInfo.userTagsArray == nil) {
    resultDeviceInfo.userTagsArray = [[NSMutableArray alloc]init];
  }
  [resultDeviceInfo.userTagsArray addObject:value];
  return self;
}
- (BDeviceInfoBuilder *)setUserTagsArray:(NSArray *)array {
  resultDeviceInfo.userTagsArray = [[NSMutableArray alloc] initWithArray:array];
  return self;
}
- (BDeviceInfoBuilder *)clearUserTags {
  resultDeviceInfo.userTagsArray = nil;
  return self;
}
- (BOOL) hasDeviceUDID {
  return resultDeviceInfo.hasDeviceUDID;
}
- (NSString*) deviceUDID {
  return resultDeviceInfo.deviceUDID;
}
- (BDeviceInfoBuilder*) setDeviceUDID:(NSString*) value {
  resultDeviceInfo.hasDeviceUDID = YES;
  resultDeviceInfo.deviceUDID = value;
  return self;
}
- (BDeviceInfoBuilder*) clearDeviceUDID {
  resultDeviceInfo.hasDeviceUDID = NO;
  resultDeviceInfo.deviceUDID = @"";
  return self;
}
- (BOOL) hasColorDepth {
  return resultDeviceInfo.hasColorDepth;
}
- (Float32) colorDepth {
  return resultDeviceInfo.colorDepth;
}
- (BDeviceInfoBuilder*) setColorDepth:(Float32) value {
  resultDeviceInfo.hasColorDepth = YES;
  resultDeviceInfo.colorDepth = value;
  return self;
}
- (BDeviceInfoBuilder*) clearColorDepth {
  resultDeviceInfo.hasColorDepth = NO;
  resultDeviceInfo.colorDepth = 0;
  return self;
}
- (BOOL) hasIPAddress {
  return resultDeviceInfo.hasIPAddress;
}
- (NSString*) iPAddress {
  return resultDeviceInfo.iPAddress;
}
- (BDeviceInfoBuilder*) setIPAddress:(NSString*) value {
  resultDeviceInfo.hasIPAddress = YES;
  resultDeviceInfo.iPAddress = value;
  return self;
}
- (BDeviceInfoBuilder*) clearIPAddress {
  resultDeviceInfo.hasIPAddress = NO;
  resultDeviceInfo.iPAddress = @"";
  return self;
}
- (BOOL) hasSystemBuildVersion {
  return resultDeviceInfo.hasSystemBuildVersion;
}
- (NSString*) systemBuildVersion {
  return resultDeviceInfo.systemBuildVersion;
}
- (BDeviceInfoBuilder*) setSystemBuildVersion:(NSString*) value {
  resultDeviceInfo.hasSystemBuildVersion = YES;
  resultDeviceInfo.systemBuildVersion = value;
  return self;
}
- (BDeviceInfoBuilder*) clearSystemBuildVersion {
  resultDeviceInfo.hasSystemBuildVersion = NO;
  resultDeviceInfo.systemBuildVersion = @"";
  return self;
}
- (BOOL) hasLocalIPAddress {
  return resultDeviceInfo.hasLocalIPAddress;
}
- (NSString*) localIPAddress {
  return resultDeviceInfo.localIPAddress;
}
- (BDeviceInfoBuilder*) setLocalIPAddress:(NSString*) value {
  resultDeviceInfo.hasLocalIPAddress = YES;
  resultDeviceInfo.localIPAddress = value;
  return self;
}
- (BDeviceInfoBuilder*) clearLocalIPAddress {
  resultDeviceInfo.hasLocalIPAddress = NO;
  resultDeviceInfo.localIPAddress = @"";
  return self;
}
@end


// @@protoc_insertion_point(global_scope)
