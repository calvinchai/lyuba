package main

import (
    "github.com/hashgraph/hedera-sdk-go/v2"
    "time"
)

func main() {
    // Define your client, operator account ID and private key
    client := hedera.ClientForTestnet()
    operatorAccountID := hedera.AccountID{Account: 1234} // Replace with your account ID
    operatorPrivateKey, _ := hedera.PrivateKeyFromString("your-private-key") // Replace with your private key

    client.SetOperator(operatorAccountID, operatorPrivateKey)

    // Define the token ID and the node account ID you want to reward
    tokenID := hedera.TokenID{Token: 5678} // Replace with your token ID
    nodeAccountID := hedera.AccountID{Account: 9012} // Replace with the node account ID

    // Define the amount to reward
    rewardAmount := uint64(100) // Change to the amount you want to send

    // Create a transfer transaction
    transferTx, _ := hedera.NewTransferTransaction().
        AddTokenTransfer(tokenID, operatorAccountID, -rewardAmount).
        AddTokenTransfer(tokenID, nodeAccountID, rewardAmount).
        SetTransactionMemo("Rewarding node").
        SetTransactionValidDuration(120 * time.Second).
        FreezeWith(client)

    // Sign the transaction
    signedTx, _ := transferTx.Sign(operatorPrivateKey)

    // Submit the transaction to a Hedera network
    txResponse, _ := signedTx.Execute(client)

    // Request the receipt of the transaction
    receipt, _ := txResponse.GetReceipt(client)

    // Output the transaction status
    println("Transaction status:", receipt.Status)
}

