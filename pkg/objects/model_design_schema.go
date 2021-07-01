/*
 * Fledge REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package objects

import "go.mongodb.org/mongo-driver/bson/primitive"

// DesignSchema - Schema to define the roles and their connections
type DesignSchema struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`

	IsCurrentlyUsed bool `bson:"isCurrentlyUsed" json:"isCurrentlyUsed"`

	Name string `json:"name"`

	Description string `json:"description,omitempty"`

	Roles []Role `bson:"roles" json:"roles"`

	Channels []Channel `json:"channels"`

	Connectors []Connector `json:"connectors,omitempty"`
}