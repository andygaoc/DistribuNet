package main

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// go mod init golang-secp256k1-signature
// https://stackoverflow.com/questions/61682191/go-ethereum-sign-provides-different-signature-than-nodejs-ethers
func main() {
	// base64.StdEncoding.EncodeToString
	//dataStrBase64 := `eyJhZGRyIjoiMHhlODkzNWRhNjViMzU4YWZmNjk0M2EyMzk0MDMyNzU5NTc1ZGJkZWY1IiwiY3JlZGVudGlhbCI6IntcImNuXCI6XCJjcFwiLFwiY3RcIjpcImxKT0JzbXpPb1QxdEZackRyOW00OTVabkU0dmlPUi9xenphZjNYRDh3OWl5RnJVQXRoazZnekRHNnd2VDk3bkNjMWhSeWEramY3MW5zRm16Um13TjJ3QjNrakpQenpQVHF4MGFLY3hmTTJEY3ZVbDY2dlloamFiUW9iUW5jakQzTk5EZEQ2eTFVRzlzcU5tM2VUb1JDOUVMVTh2eTE3RlBYdk8rdTJsTGpZQlovZW9HMnJRc3djcG5Wc21GQ2FpUDVIdVhLeHJPOHdUUTZhQlhyNk90OHJ0MDAyeUlIQmdIWHRGeG5zM2tOSkZId3BQblJTZzNnTXc5Y1dUQ3drNmJ1MGRRZzlpZlgrT3ZoUFNFTGVSanE5WXdNNHdsM2tvci9GUkV4d0xIZnpJbkhBUW5BSjlEY01qMC90ZnY4bDhBL2wwaVwiLFwiaWRcIjpcIjEzMVwiLFwiaXZcIjpcImhvTVdaSFVmZnRiQm80SEZcIixcInRnXCI6XCJ6ejl5c0lEOGZmZll6c1I5V3UzYXd3PT1cIn0iLCJpZCI6Ijg4Mzc2YzU2LTdkNzItNDI2ZC1iYzQwLTRkNmY2MTNlZDcyYyIsIm5vbmNlIjoiMTU4OGRiZDMtYzMyMS00NGRmLTkwNjctYTA4YjhhMDdkNGMzIiwidGltZXN0YW1wIjoxNjkxNTc2MDM4fQ==`

	//dataStr, _ := base64.StdEncoding.DecodeString(dataStrBase64)
	dataStr := `{"addr":"0x5892034D0EeE2E4020ba314b6cfbE02a9a8bEe7a","nonce":"60d8f593-45d3-4cc1-8525-bece6570cb95","timestamp":1694261472,"personal_address":"0xafa05b51afc0deded22bad455ffbb4d0b30344ba"}`

	sigStr := "1a574d33dd7e4eee084645d14ed94484a8a9760df16cf9a8adbb6b85fe0dcb8f3d6aec06a5bb19babaad24ecbc144235d5de5445ce7780f45317652e2f2c84b11b"

	addrStr := "0x5892034D0EeE2E4020ba314b6cfbE02a9a8bEe7a"

	addrPrivateKey := "f738d96b723d800f764bea82199202177943bf7211539c883e5cf812dfbdd101"

	signature, _ := hex.DecodeString(sigStr)

	dataBytes := []byte(dataStr)
	dataHash := crypto.Keccak256Hash(dataBytes)

	if len(signature) != 65 {
		fmt.Errorf("invalid signature length: %d", len(signature))
	}

	if signature[64] != 27 && signature[64] != 28 {
		fmt.Errorf("invalid recovery id: %d", signature[64])
	}

	fmt.Printf("recovery id, signature[64] = %d\n", signature[64])
	signature[64] -= 27

	pubKeyRaw, err := crypto.Ecrecover(dataHash.Bytes(), signature)
	if err != nil {
		fmt.Errorf("invalid signature: %s", err.Error())
	}

	pubKey, err := crypto.UnmarshalPubkey(pubKeyRaw)
	if err != nil {
		fmt.Errorf("UnmarshalPubkey failed: %s", err.Error())
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	fmt.Println("raw addr", addrStr, ", recovery addr", recoveredAddr.String())

	privateKey, _ := crypto.HexToECDSA(addrPrivateKey)
	signature2, _ := crypto.Sign(dataHash.Bytes(), privateKey)
	if signature2[64] == 0 || signature2[64] == 1 {
		signature2[64] += 27
	}

	fmt.Println("signature: ", sigStr)
	hashSig := hexutil.Encode(signature2)
	fmt.Println("signature2: ", hashSig)
}
