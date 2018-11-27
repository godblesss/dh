package main

// group id: 1 - RFC 2409 - First Oakley Group
//           2 - RFC 2409 - Second Oakley Group
//           14 - RFC 3526 - 2048-bit MODP Group
//           15 - RFC 3526 - 3072-bit MODP Group
//           16 - RFC 3526 - 4096-bit MODP Group
//  ...6144-bit MODP Group, 8192-bit MODP Group are included in RFC 3526,

// return DHGroup. Set the G and P.
// 	group, _ := GetGroup(id)
//
// Generate a private key from the group.
// return DHKEY. { x = random number[1, P), y = G ^ x mod P }
// 	aKey, _ := group.GeneratePrivateKey(nil)
//
// return DHKey.y.bytes
// 	pub := aKey.Bytes()
//
// Send the public key to B.
// 	Send("B", pub)
//
// Receive a slice of bytes from B, which contains B's public key
// 	b := Recv("B")
//
// Recover B's public key
// New DHKey, set DHKey.y
// 	bKey := NewPublicKey(b)
//
// Compute the key
// 	k, _ := group.ComputeKey(bKey, aKey)
//
// Get the key in the form of []byte
// 	key := k.Bytes()