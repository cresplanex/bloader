package config

import "fmt"

var (
	// ErrTypeInvalid is the error for the invalid config type.
	ErrTypeInvalid = fmt.Errorf("config type is invalid")
	// ErrTypeRequired is the error for the required config type.
	ErrTypeRequired = fmt.Errorf("config type is required")
	// ErrEnvRequired is the error for the required environment.
	ErrEnvRequired = fmt.Errorf("environment is required")
	// ErrLoaderRequired is the error for the required loader.
	ErrLoaderRequired = fmt.Errorf("loader is required")
	// ErrTargetsRequired is the error for the required targets.
	ErrTargetsRequired = fmt.Errorf("targets are required")
	// ErrOutputsRequired is the error for the required outputs.
	ErrOutputsRequired = fmt.Errorf("outputs are required")
	// ErrStoreRequired is the error for the required store.
	ErrStoreRequired = fmt.Errorf("store is required")
	// ErrEncryptsRequired is the error for the required encrypts.
	ErrEncryptsRequired = fmt.Errorf("encrypts are required")
	// ErrAuthRequired is the error for the required auth.
	ErrAuthRequired = fmt.Errorf("auth is required")
	// ErrServerRequired is the error for the required server.
	ErrServerRequired = fmt.Errorf("server is required")
	// ErrLoggingRequired is the error for the required logging.
	ErrLoggingRequired = fmt.Errorf("logging is required")
	// ErrClockRequired is the error for the required clock.
	ErrClockRequired = fmt.Errorf("clock is required")
	// ErrLanguageRequired is the error for the required language.
	ErrLanguageRequired = fmt.Errorf("language is required")
	// ErrOverrideRequired is the error for the required override.
	ErrOverrideRequired = fmt.Errorf("override is required")
	// ErrSlaveSettingRequired is the error for the required slave
	ErrSlaveSettingRequired = fmt.Errorf("slave setting is required")
	// ErrServerPortRequired is the error for the required server port.
	ErrServerPortRequired = fmt.Errorf("server port is required")
	// ErrLoggingOutputTypeRequired is the error for the required logging output type.
	ErrLoggingOutputTypeRequired = fmt.Errorf("logging output type is required")
	// ErrLoggingOutputFormatRequired is the error for the required logging output format.
	ErrLoggingOutputFormatRequired = fmt.Errorf("logging output format is required")
	// ErrLoggingOutputLevelRequired is the error for the required logging output level.
	ErrLoggingOutputLevelRequired = fmt.Errorf("logging output level is required")
	// ErrLoggingOutputTypeInvalid is the error for the invalid logging output type.
	ErrLoggingOutputTypeInvalid = fmt.Errorf("logging output type is invalid")
	// ErrLoggingOutputFormatInvalid is the error for the invalid logging output format.
	ErrLoggingOutputFormatInvalid = fmt.Errorf("logging output format is invalid")
	// ErrLoggingOutputLevelInvalid is the error for the invalid logging output level.
	ErrLoggingOutputLevelInvalid = fmt.Errorf("logging output level is invalid")
	// ErrLoggingOutputFilenameRequired is the error for the required logging output filename.
	ErrLoggingOutputFilenameRequired = fmt.Errorf("logging output filename is required")
	// ErrLoggingOutputAddressRequired is the error for the required logging output address.
	ErrLoggingOutputAddressRequired = fmt.Errorf("logging output address is required")
	// ErrLanguageDefaultRequired is the error for the required language default.
	ErrLanguageDefaultRequired = fmt.Errorf("language default is required")
	// ErrLanguageDefaultInvalid is the error for the invalid language default.
	ErrLanguageDefaultInvalid = fmt.Errorf("language default is invalid")
	// ErrEncryptIDRequired is the error for the required encrypt ID.
	ErrEncryptIDRequired = fmt.Errorf("encrypt ID is required")
	// ErrEncryptIDDuplicate is the error for the duplicate encrypt ID.
	ErrEncryptIDDuplicate = fmt.Errorf("encrypt ID is duplicate")
	// ErrEncryptTypeRequired is the error for the required encrypt type.
	ErrEncryptTypeRequired = fmt.Errorf("encrypt type is required")
	// ErrEncryptTypeUnsupportedOnSlave is the error for the unsupported encrypt type on the slave.
	ErrEncryptTypeUnsupportedOnSlave = fmt.Errorf("encrypt type is unsupported on the slave")
	// ErrEncryptTypeInvalid is the error for the invalid encrypt type.
	ErrEncryptTypeInvalid = fmt.Errorf("encrypt type is invalid")
	// ErrEncryptKeyRequired is the error for the required encrypt key.
	ErrEncryptKeyRequired = fmt.Errorf("encrypt key is required")
	// ErrEncryptMethodRequired is the error for the required encrypt method.
	ErrEncryptMethodRequired = fmt.Errorf("encrypt method is required")
	// ErrEncryptMethodInvalid is the error for the invalid encrypt method.
	ErrEncryptMethodInvalid = fmt.Errorf("encrypt method is invalid")
	// ErrEncryptRSAKeySizeInvalid is the error for the invalid encrypt RSA key size.
	ErrEncryptRSAKeySizeInvalid = fmt.Errorf("encrypt RSA key size is invalid, must be 16, 24, or 32 bytes")
	// ErrEncryptStoreRequired is the error for the required encrypt store.
	ErrEncryptStoreRequired = fmt.Errorf("encrypt store is required")
	// ErrClockFormatInvalid is the error for the invalid clock format.
	ErrClockFormatInvalid = fmt.Errorf("clock format is invalid")
	// ErrClockFakeTimeRequired is the error for the required clock fake time.
	ErrClockFakeTimeRequired = fmt.Errorf("clock fake time is required")
	// ErrClockFakeTimeInvalid is the error for the invalid clock fake time.
	ErrClockFakeTimeInvalid = fmt.Errorf("clock fake time is invalid")
	// ErrStoreFilePathRequired is the error for the required store filepath.
	ErrStoreFilePathRequired = fmt.Errorf("store filepath is required")
	// ErrStoreFileEnvRequired is the error for the required store file env.
	ErrStoreFileEnvRequired = fmt.Errorf("store file env is required")
	// ErrStoreBucketIDRequired is the error for the required store bucket ID.
	ErrStoreBucketIDRequired = fmt.Errorf("store bucket ID is required")
	// ErrStoreBucketIDDuplicate is the error for the duplicate store bucket ID.
	ErrStoreBucketIDDuplicate = fmt.Errorf("store bucket ID is duplicate")
	// ErrStoreKeyRequired is the error for the required store key.
	ErrStoreKeyRequired = fmt.Errorf("store key is required")
	// ErrStoreEncryptIDRequired is the error for the required store encrypt ID.
	ErrStoreEncryptIDRequired = fmt.Errorf("store encrypt ID is required")
	// ErrAuthIDRequired is the error for the required auth ID.
	ErrAuthIDRequired = fmt.Errorf("auth ID is required")
	// ErrAuthIDDuplicate is the error for the duplicate auth ID.
	ErrAuthIDDuplicate = fmt.Errorf("auth ID is duplicate")
	// ErrAuthTypeRequired is the error for the required auth type.
	ErrAuthTypeRequired = fmt.Errorf("auth type is required")
	// ErrAuthTypeInvalid is the error for the invalid auth type.
	ErrAuthTypeInvalid = fmt.Errorf("auth type is invalid")
	// ErrAuthDefaultDuplicate is the error for the duplicate auth default.
	ErrAuthDefaultDuplicate = fmt.Errorf("auth default is duplicate")
	// ErrAuthCredentialStoreRequired is the error for the required auth credential store.
	ErrAuthCredentialStoreRequired = fmt.Errorf("auth credential store is required")
	// ErrAuthOAuth2Required is the error for the required auth OAuth2.
	ErrAuthOAuth2Required = fmt.Errorf("auth OAuth2 is required")
	// ErrAuthOAuth2GrantTypeRequired is the error for the required auth OAuth2 grant type.
	ErrAuthOAuth2GrantTypeRequired = fmt.Errorf("auth OAuth2 grant type is required")
	// ErrAuthOAuth2GrantTypeInvalid is the error for the invalid auth OAuth2 grant type.
	ErrAuthOAuth2GrantTypeInvalid = fmt.Errorf("auth OAuth2 grant type is invalid")
	// ErrAuthOAuth2ClientIDRequired is the error for the required auth OAuth2 client ID.
	ErrAuthOAuth2ClientIDRequired = fmt.Errorf("auth OAuth2 client ID is required")
	// ErrAuthOAuth2ClientSecretRequired is the error for the required auth OAuth2 client secret.
	ErrAuthOAuth2ClientSecretRequired = fmt.Errorf("auth OAuth2 client secret is required")
	// ErrAuthOAuth2AuthURLRequired is the error for the required auth OAuth2 auth URL.
	ErrAuthOAuth2AuthURLRequired = fmt.Errorf("auth OAuth2 auth URL is required")
	// ErrAuthOAuth2TokenURLRequired is the error for the required auth OAuth2 token URL.
	ErrAuthOAuth2TokenURLRequired = fmt.Errorf("auth OAuth2 token URL is required")
	// ErrAuthOAuth2AccessTypeRequired is the error for the required auth OAuth2 access type.
	ErrAuthOAuth2AccessTypeRequired = fmt.Errorf("auth OAuth2 access type is required")
	// ErrAuthOAuth2AccessTypeInvalid is the error for the invalid auth OAuth2 access type.
	ErrAuthOAuth2AccessTypeInvalid = fmt.Errorf("auth OAuth2 access type is invalid")
	// ErrAuthOAuth2UsernameRequired is the error for the required auth OAuth2 username.
	ErrAuthOAuth2UsernameRequired = fmt.Errorf("auth OAuth2 username is required")
	// ErrAuthOAuth2PasswordRequired is the error for the required auth OAuth2 password.
	ErrAuthOAuth2PasswordRequired = fmt.Errorf("auth OAuth2 password is required")
	// ErrAuthOAuth2CredentialRequired is the error for the required auth OAuth2 credential.
	ErrAuthOAuth2CredentialRequired = fmt.Errorf("auth OAuth2 credential is required")
	// ErrAuthAPIKeyRequired is the error for the required auth API key.
	ErrAuthAPIKeyRequired = fmt.Errorf("auth API key is required")
	// ErrAuthAPIKeyHeaderNameRequired is the error for the required auth API key header name.
	ErrAuthAPIKeyHeaderNameRequired = fmt.Errorf("auth API key header name is required")
	// ErrAuthAPIKeyKeyRequired is the error for the required auth API key key.
	ErrAuthAPIKeyKeyRequired = fmt.Errorf("auth API key key is required")
	// ErrAuthBasicRequired is the error for the required auth basic.
	ErrAuthBasicRequired = fmt.Errorf("auth basic is required")
	// ErrAuthBasicUsernameRequired is the error for the required auth basic username.
	ErrAuthBasicUsernameRequired = fmt.Errorf("auth basic username is required")
	// ErrAuthBasicPasswordRequired is the error for the required auth basic password.
	ErrAuthBasicPasswordRequired = fmt.Errorf("auth basic password is required")
	// ErrAuthPrivateKeyRequired is the error for the required auth private key.
	ErrAuthPrivateKeyRequired = fmt.Errorf("auth private key is required")
	// ErrAuthPrivateKeyPrivateKeyRequired is the error for the required auth private key private key.
	ErrAuthPrivateKeyPrivateKeyRequired = fmt.Errorf("auth private key is required on private key auth")
	// ErrAuthJWTRequired is the error for the required auth JWT.
	ErrAuthJWTRequired = fmt.Errorf("auth JWT is required")
	// ErrAuthJWTCredentialRequired is the error for the required auth JWT credential.
	ErrAuthJWTCredentialRequired = fmt.Errorf("auth JWT credential is required")
	// ErrTargetIDRequired is the error for the required target ID.
	ErrTargetIDRequired = fmt.Errorf("target ID is required")
	// ErrTargetIDDuplicate is the error for the duplicate target ID.
	ErrTargetIDDuplicate = fmt.Errorf("target ID is duplicate")
	// ErrTargetTypeRequired is the error for the required target type.
	ErrTargetTypeRequired = fmt.Errorf("target type is required")
	// ErrTargetTypeInvalid is the error for the invalid target type.
	ErrTargetTypeInvalid = fmt.Errorf("target type is invalid")
	// ErrTargetValueEnvRequired is the error for the required target value env.
	ErrTargetValueEnvRequired = fmt.Errorf("target value env is required")
	// ErrTargetValueURLRequired is the error for the required target value URL.
	ErrTargetValueURLRequired = fmt.Errorf("target value URL is required")
	// ErrOutputValueEnvRequired is the error for the required output value env.
	ErrOutputValueEnvRequired = fmt.Errorf("output value env is required")
	// ErrOutputValueTypeRequired is the error for the required output value type.
	ErrOutputValueTypeRequired = fmt.Errorf("output value type is required")
	// ErrOutputValueTypeInvalid is the error for the invalid output value type.
	ErrOutputValueTypeInvalid = fmt.Errorf("output value type is invalid")
	// ErrOutputValueFormatRequired is the error for the required output value format.
	ErrOutputValueFormatRequired = fmt.Errorf("output value format is required")
	// ErrOutputValueFormatInvalid is the error for the invalid output value format.
	ErrOutputValueFormatInvalid = fmt.Errorf("output value format is invalid")
	// ErrOutputValueBasePathRequired is the error for the required output value base path.
	ErrOutputValueBasePathRequired = fmt.Errorf("output value base path is required")
	// ErrOutputValueIDRequired is the error for the required output value ID.
	ErrOutputValueIDRequired = fmt.Errorf("output value ID is required")
	// ErrOutputValueIDDuplicate is the error for the duplicate output value ID.
	ErrOutputValueIDDuplicate = fmt.Errorf("output value ID is duplicate")
	// ErrOverrideVarKeyRequired is the error for the required override var key.
	ErrOverrideVarKeyRequired = fmt.Errorf("override var key is required")
	// ErrOverrideVarValueRequired is the error for the required override var value.
	ErrOverrideVarValueRequired = fmt.Errorf("override var value is required")
	// ErrOverrideTypeRequired is the error for the required override type.
	ErrOverrideTypeRequired = fmt.Errorf("override type is required")
	// ErrOverrideTypeInvalid is the error for the invalid override type.
	ErrOverrideTypeInvalid = fmt.Errorf("override type is invalid")
	// ErrOverrideFileTypeRequired is the error for the required override file type.
	ErrOverrideFileTypeRequired = fmt.Errorf("override file type is required")
	// ErrOverrideFileTypeInvalid is the error for the invalid override file type.
	ErrOverrideFileTypeInvalid = fmt.Errorf("override file type is invalid")
	// ErrOverridePathRequired is the error for the required override path.
	ErrOverridePathRequired = fmt.Errorf("override path is required")
	// ErrOverrideStoreRequired is the error for the required override store.
	ErrOverrideStoreRequired = fmt.Errorf("override store is required")
	// ErrOverrideKeyRequired is the error for the required override key.
	ErrOverrideKeyRequired = fmt.Errorf("override key is required")
	// ErrOverrideValueRequired is the error for the required override value.
	ErrOverrideValueRequired = fmt.Errorf("override value is required")
	// ErrLoaderBasePathRequired is the error for the required loader base path.
	ErrLoaderBasePathRequired = fmt.Errorf("loader base path is required")
	// ErrSlaveSettingPortRequired is the error for the required slave setting port.
	ErrSlaveSettingPortRequired = fmt.Errorf("slave setting port is required")
	// ErrSlaveCertificateSlaveCertPathRequired is the error for the required slave certificate slave certificate path.
	ErrSlaveCertificateSlaveCertPathRequired = fmt.Errorf("slave certificate slave certificate path is required")
	// ErrSlaveCertificateSlaveKeyPathRequired is the error for the required slave certificate slave key path.
	ErrSlaveCertificateSlaveKeyPathRequired = fmt.Errorf("slave certificate slave key path is required")
	// ErrSlaveSettingEncryptIDRequired is the error for the required slave setting encrypt ID.
	ErrSlaveSettingEncryptIDRequired = fmt.Errorf("slave setting encrypt ID is required")
)
