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
 * @interface CredRequest
 */
export interface CredRequest {
    /**
     * 
     * @type {string}
     * @memberof CredRequest
     */
    destination?: string;
}

export function CredRequestFromJSON(json: any): CredRequest {
    return CredRequestFromJSONTyped(json, false);
}

export function CredRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): CredRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'destination': !exists(json, 'destination') ? undefined : json['destination'],
    };
}

export function CredRequestToJSON(value?: CredRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'destination': value.destination,
    };
}


