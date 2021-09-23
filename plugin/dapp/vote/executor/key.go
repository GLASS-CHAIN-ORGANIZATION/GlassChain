package executor

/*
 *       kv   ，key           
 *  key = keyPrefix + userKey
 *          ，  ’-‘      
 */

var (
	//keyPrefixStateDB state db key    
	keyPrefixStateDB = "mavl-vote-"
	//keyPrefixLocalDB local db key    
	keyPrefixLocalDB = "LODB-vote-"
)

// groupID or voteID
func formatStateIDKey(id string) []byte {
	return []byte(keyPrefixStateDB + id)
}
