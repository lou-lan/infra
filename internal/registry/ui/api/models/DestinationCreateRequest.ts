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
import {
    DestinationKubernetes,
    DestinationKubernetesFromJSON,
    DestinationKubernetesFromJSONTyped,
    DestinationKubernetesToJSON,
} from './';

/**
 * 
 * @export
 * @interface DestinationCreateRequest
 */
export interface DestinationCreateRequest {
    /**
     * 
     * @type {string}
     * @memberof DestinationCreateRequest
     */
    name: string;
    /**
     * 
     * @type {DestinationKubernetes}
     * @memberof DestinationCreateRequest
     */
    kubernetes?: DestinationKubernetes;
}

export function DestinationCreateRequestFromJSON(json: any): DestinationCreateRequest {
    return DestinationCreateRequestFromJSONTyped(json, false);
}

export function DestinationCreateRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): DestinationCreateRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'name': json['name'],
        'kubernetes': !exists(json, 'kubernetes') ? undefined : DestinationKubernetesFromJSON(json['kubernetes']),
    };
}

export function DestinationCreateRequestToJSON(value?: DestinationCreateRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'name': value.name,
        'kubernetes': DestinationKubernetesToJSON(value.kubernetes),
    };
}

