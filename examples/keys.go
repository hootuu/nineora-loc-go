package examples

import "github.com/hootuu/nineorai/keys"

var gKeysDict map[string]*keys.Key

func init() {
	gKeysDict = make(map[string]*keys.Key)
}

func GetKey(keyName string) *keys.Key {
	key, ok := gKeysDict[keyName]
	if !ok {
		key, _ = keys.NewKey()
		gKeysDict[keyName] = key
	}
	return key
}
