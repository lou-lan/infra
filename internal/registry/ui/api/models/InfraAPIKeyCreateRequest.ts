/* tslint:disable */
/* eslint-disable */
/**
 * Infra API
 * Infra REST API
 *
 * The version of the OpenAPI document: 0.1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface InfraAPIKeyCreateRequest
 */
export interface InfraAPIKeyCreateRequest {
    /**
     * 
     * @type {string}
     * @memberof InfraAPIKeyCreateRequest
     */
    name: string;
}

export function InfraAPIKeyCreateRequestFromJSON(json: any): InfraAPIKeyCreateRequest {
    return InfraAPIKeyCreateRequestFromJSONTyped(json, false);
}

export function InfraAPIKeyCreateRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): InfraAPIKeyCreateRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'name': json['name'],
    };
}

export function InfraAPIKeyCreateRequestToJSON(value?: InfraAPIKeyCreateRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'name': value.name,
    };
}


