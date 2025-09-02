/*
I amde this file to create the interface of the items, so evrytime I pass the GetName to any struct
it will be cosiderd as an item
*/
package items

type Item interface {
	GetName() string
}
