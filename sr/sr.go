// Description: Avro utils
// Author: Pixie79
// ============================================================================
// package tinygo

package tinygo

import (
	"fmt"
	"strconv"

	utils "github.com/pixie79/tiny-utils/utils"
	sr "github.com/redpanda-data/redpanda/src/go/transform-sdk/sr"
)

// GetSchema retrieves the schema with the given ID from the specified URL.
//
// Parameters:
// - id: The ID of the schema (as a string).
// - url: The URL of the Schema Registry.
//
// Returns:
// - The retrieved schema (as a string).
func GetSchema(id string) string {
	registry := sr.NewClient()
	schemaIdInt, err := strconv.Atoi(id)
	utils.Print("DEBUG", fmt.Sprintf("Schema ID: %s", id))
	utils.MaybeDie(err, fmt.Sprintf("SCHEMA_ID not an integer: %s", id))
	schema, err := registry.LookupSchemaById(schemaIdInt)
	utils.MaybeDie(err, fmt.Sprintf("Unable to retrieve schema for ID: %s", id))
	return schema.Schema
}

// ExtractID extracts the ID from a byte slice.
//
// Takes a byte slice as input.
// Returns an integer.
func ExtractID(msg []byte) (int, error) {
	schemaID, err := sr.ExtractID(msg)
	return schemaID, err
}
