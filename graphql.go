package uuid

//GetGraphQLType returns the associated GraphQL type
func (_ UUID) GetGraphQLType() string {
	return "uuid"
}
