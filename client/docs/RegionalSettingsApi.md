# \RegionalSettingsApi

All URIs are relative to *https://platform.devtest.ringcentral.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCountries**](RegionalSettingsApi.md#ListCountries) | **Get** /restapi/v1.0/dictionary/country | Get Country List
[**ListLanguages**](RegionalSettingsApi.md#ListLanguages) | **Get** /restapi/v1.0/dictionary/language | Get Language List
[**ListLocations**](RegionalSettingsApi.md#ListLocations) | **Get** /restapi/v1.0/dictionary/location | Get Location List
[**ListTimezones**](RegionalSettingsApi.md#ListTimezones) | **Get** /restapi/v1.0/dictionary/timezone | Get Timezone List
[**LoadCountry**](RegionalSettingsApi.md#LoadCountry) | **Get** /restapi/v1.0/dictionary/country/{countryId} | Get Country
[**LoadLanguage**](RegionalSettingsApi.md#LoadLanguage) | **Get** /restapi/v1.0/dictionary/language/{languageId} | Get Language
[**LoadState**](RegionalSettingsApi.md#LoadState) | **Get** /restapi/v1.0/dictionary/state/{stateId} | Get State
[**LoadTimezone**](RegionalSettingsApi.md#LoadTimezone) | **Get** /restapi/v1.0/dictionary/timezone/{timezoneId} | Get Timezone
[**LoadlistStates**](RegionalSettingsApi.md#LoadlistStates) | **Get** /restapi/v1.0/dictionary/state | Get State List


# **ListCountries**
> GetCountryListResponse ListCountries(ctx, optional)
Get Country List

<p style='font-style:italic;'>Since 1.0.10 (Release 6.2)</p><p>Returns all the countries available for calling.</p><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListCountriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListCountriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginAllowed** | **optional.Bool**| Specifies whether login with the phone numbers of this country is enabled or not | 
 **signupAllowed** | **optional.Bool**| Indicates whether signup/billing is allowed for a country. If not specified all countries are returned (according to other filters specified if any) | 
 **numberSelling** | **optional.Bool**| Specifies if RingCentral sells phone numbers of this country | 
 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | 
 **perPage** | **optional.Int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | 
 **freeSoftphoneLine** | **optional.Bool**| Specifies if free phone line for softphone is available for a country or not | 

### Return type

[**GetCountryListResponse**](GetCountryListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLanguages**
> LanguageList ListLanguages(ctx, )
Get Language List

<p style='font-style:italic;'>Since 1.0.14 (Release 6.6)</p><p>Returns the information about supported languages.</p><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LanguageList**](LanguageList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLocations**
> GetLocationListResponse ListLocations(ctx, optional)
Get Location List

<p style='font-style:italic;'>Since 1.0.10 (Release 6.2)</p><p>Returns all the available locations for the certain state.</p><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListLocationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListLocationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | [**optional.Interface of []string**](string.md)| Sorts results by the specified property. The default value is &#39;City&#39; | 
 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39;. | 
 **perPage** | **optional.Int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default. | 
 **stateId** | **optional.String**| Internal identifier of a state | 
 **withNxx** | **optional.Bool**| Specifies if nxx codes are returned | 

### Return type

[**GetLocationListResponse**](GetLocationListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTimezones**
> GetTimezoneListResponse ListTimezones(ctx, optional)
Get Timezone List

<p style='font-style:italic;'>Since 1.0.10 (Release 6.2)</p><p>Returns all available timezones.</p><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListTimezonesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListTimezonesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | 
 **perPage** | **optional.String**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | 

### Return type

[**GetTimezoneListResponse**](GetTimezoneListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadCountry**
> GetCountryInfoDictionaryResponse LoadCountry(ctx, countryId)
Get Country

<p style='font-style:italic;'>Since 1.0.10 (Release 6.2)</p><p>Returns the information on the required country.</p><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryId** | **int32**| Internal identifier of a country | 

### Return type

[**GetCountryInfoDictionaryResponse**](GetCountryInfoDictionaryResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadLanguage**
> LanguageInfo LoadLanguage(ctx, languageId)
Get Language

<p style='font-style:italic;'>Since 1.0.14 (Release 6.6)</p><p>Returns language by its respective ID.</p><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **languageId** | **int32**| Internal identifier of a language | 

### Return type

[**LanguageInfo**](LanguageInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadState**
> GetStateInfoResponse LoadState(ctx, stateId)
Get State

<p style='font-style:italic;'>Since 1.0.10 (Release 6.2)</p><p>Returns the information on the required state.</p><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **stateId** | **int32**| Internal identifier of a state | 

### Return type

[**GetStateInfoResponse**](GetStateInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadTimezone**
> GetTimezoneInfoResponse LoadTimezone(ctx, timezoneId, optional)
Get Timezone

<p style='font-style:italic;'>Since 1.0.10 (Release 6.2)</p><p>Returns the information on a certain timezone.</p><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timezoneId** | **int32**| Internal identifier of a timezone | 
 **optional** | ***LoadTimezoneOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadTimezoneOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | 
 **perPage** | **optional.Int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | 

### Return type

[**GetTimezoneInfoResponse**](GetTimezoneInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadlistStates**
> GetStateListResponse LoadlistStates(ctx, optional)
Get State List

<p style='font-style:italic;'>Since 1.0.10 (Release 6.2)</p><p>Returns all the states for a certain country.</p><h4>Usage Plan Group</h4><p>Light</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LoadlistStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadlistStatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryId** | **optional.Int32**| Internal identifier of a country | 
 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39;. | 
 **perPage** | **optional.Int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | 
 **withPhoneNumbers** | **optional.Bool**| If &#39;True&#39;, the list of states with phone numbers available for buying is returned. The default value is &#39;False&#39; | 

### Return type

[**GetStateListResponse**](GetStateListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

