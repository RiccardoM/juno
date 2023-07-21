package messages

import (
	"fmt"
	"strings"

	"github.com/cosmos/gogoproto/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/juno/v5/types/config"

	"github.com/forbole/juno/v5/types"
)

// MessageNotSupported returns an error telling that the given message is not supported
func MessageNotSupported(msg sdk.Msg) error {
	return fmt.Errorf("message type not supported: %s", proto.MessageName(msg))
}

// MessageAddressesParser represents a function that extracts all the
// involved addresses from a provided message (both accounts and validators)
type MessageAddressesParser = func(tx *types.Tx) ([]string, error)

// CosmosMessageAddressesParser represents a MessageAddressesParser that parses a
// Chain message and returns all the involved addresses (both accounts and validators)
var CosmosMessageAddressesParser = DefaultMessagesParser

// DefaultMessagesParser represents the default messages parser that simply returns the list
// of all the signers of a message
func DefaultMessagesParser(tx *types.Tx) ([]string, error) {
	allAddressess := parseAddressesFromEvents(tx)
	return allAddressess, nil
}

// function to remove duplicate values
func removeDuplicates(s []string) []string {
	bucket := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := bucket[str]; !ok {
			bucket[str] = true
			result = append(result, str)
		}
	}
	return result
}

func parseAddressesFromEvents(tx *types.Tx) []string {
	var allAddressess []string
	chainPrefix := config.Cfg.Chain.Bech32Prefix

	for _, event := range tx.Events {
		for _, attribute := range event.Attributes {
			if strings.Contains(attribute.Value, "/") {
				continue
			}
			if strings.Contains(attribute.Value, chainPrefix) {
				allAddressess = append(allAddressess, attribute.Value)
			}
		}

	}
	allInvolvedAddresses := removeDuplicates(allAddressess)

	return allInvolvedAddresses
}
