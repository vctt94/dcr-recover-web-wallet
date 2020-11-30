# dcr-recover-web-wallet
This repository is a tool for recoverying backups from the decred web wallet, which is deprecated.

steps

 #### OBS: You will need the password used for the wallet backup

 - Use https://bitwiseshiftleft.github.io/sjcl/demo/ for decrypting the backup wallet file
 - Now if the wallet is not encrypted, you will have a file like the following:

```
{"network":"livenet","xPrivKey":"xprv9s21ZrQH143K4bya4AzybKvpMVZs7pRK1CSJAnjmpL2XrwVYjKKVbdce88FMz3YAhYumcQarKbjGTNtRdm42BsuvADPa9be65p6xmbkrVWi","xPubKey":"xpub6DXmUpTCZxfVv7gABZvrccd3QoTGwjDLEETRpXteanLAQZ3sqz85ReM9AAR7ZgpQx1GPqA2YFZsTKA2PuUtTbWommvivjf1MsXXKKL7Lyhf","requestPrivKey":"4115bcdd322e2747b9ddbb96cba9ee729ccb8523c704bd87271daf995d7dd635","requestPubKey":"02b6a798adae3d087b66062fc3b4b3f01040acb5f7a7881aebe2f5490b1536ca16","copayerId":"8f9ae4d67e00c816a5dcd1c39cd698710035fd81f3671328a61d63fb9ef978df","publicKeyRing":[{"xPubKey":"xpub6DXmUpTCZxfVv7gABZvrccd3QoTGwjDLEETRpXteanLAQZ3sqz85ReM9AAR7ZgpQx1GPqA2YFZsTKA2PuUtTbWommvivjf1MsXXKKL7Lyhf","requestPubKey":"02b6a798adae3d087b66062fc3b4b3f01040acb5f7a7881aebe2f5490b1536ca16"}],"walletId":"950fcf52-f846-423a-90b4-b024fd6e42ed","walletName":"teste","m":1,"n":1,"walletPrivKey":"77ca3c930530bcee73d7f43981e50ca5a0166ac86dfec18ca5907c2993313572","personalEncryptingKey":"HMY4xVt8g57iRRjLRG2THQ==","sharedEncryptingKey":"p9e+Ck7Si/+ci5dHd4DpDg==","copayerName":"eu","mnemonic":"still violin symptom dinosaur report vapor session crazy chunk veteran travel mad","entropySource":"2bb7d38e5eb5790f986cb09965314da04bc74eb70376b9f4c5f7ad7ab9301947","mnemonicHasPassphrase":false,"derivationStrategy":"BIP44","account":0,"compliantDerivation":true,"addressType":"P2PKH"}
```

 - If the wallet is also encrypted, the `xPrivKey` will be encrypted, and it will also need to be decrypted.
 - Now run this repository on a terminal and it will ask for the xprivkey
 - After entering a valid xPrivKey, a hex will be returned, which can be used for recoverying the wallet (at decrediton, for example)
 
