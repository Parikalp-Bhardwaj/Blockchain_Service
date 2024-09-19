const fs = require('fs');
const path = require('path');
const Wallet  = require('ethereumjs-wallet')
const Web3 = require('web3');
const ethereumRpcUrl = 'http://192.168.253.108:8545';


const keystoreFolderPath = '/Node-2/execution/keystore'; // Replace with the path to your keystore folder
const password = '';


let privateKey = ''

const contractAddress = '0x4242424242424242424242424242424242424242';
const contractABI = [
        {
    "constant": false,
    "inputs": [
      {
        "name": "pubkey",
        "type": "bytes"
      },
      {
        "name": "withdrawal_credentials",
        "type": "bytes"
      },
      {
        "name": "signature",
        "type": "bytes"
      },
      {
        "name": "deposit_data_root",
        "type": "bytes32"
      }
    ],
    "name": "deposit",
    "outputs": [],
    "payable": true,
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "constant": true,
    "inputs": [],
    "name": "get_deposit_root",
    "outputs": [
      {
        "name": "",
        "type": "bytes32"
      }
    ],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
  },
  {
    "constant": true,
    "inputs": [],
    "name": "get_deposit_count",
    "outputs": [
      {
        "name": "",
        "type": "bytes"
      }
    ],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
}
]



async function processKeystoreFile(keystoreFilePath) {
    const keystore = JSON.parse(fs.readFileSync(keystoreFilePath).toString());
    const wallet = Wallet.fromV3(keystore, password, true);
    const privateKey = wallet.getPrivateKeyString();
  
    const web3 = new Web3(ethereumRpcUrl);
    const account = web3.eth.accounts.privateKeyToAccount(privateKey);
    web3.eth.defaultAccount = account.address;
    const contract = new web3.eth.Contract(contractABI, contractAddress);
  
    // Process the keystore, perform any necessary actions
    // For example, you can access the contract here
    console.log(`Processing keystore for address: ${account.address}`);
    return [contract ,web3, account]
  }
  

  fs.readdir(keystoreFolderPath, async (err, files) => {
    if (err) {
      console.error('Error reading keystore folder:', err);
      return;
    }
  
    const keystoreFiles = files.filter((filename) => /^UTC--.*--.*$/.test(filename));
  
    if (keystoreFiles.length === 0) {
      console.error('No keystore files found in the folder.');
      return;
    }
  
    const keystoreFileName = keystoreFiles[0];
    const keystoreFilePath = path.join(keystoreFolderPath, keystoreFileName);
  
    let [contract, web3, account] = await processKeystoreFile(keystoreFilePath);

  
    const directory = '/Node-2/validator_keys/';
    fs.readdir(directory, async (err, files) => {
        if (err) {
          console.error('Error reading directory:', err);
          return;
        }
      
        const depositDataFiles = files.filter(file => file.startsWith('deposit_data-'));
      
        depositDataFiles.forEach(async (filename) => {
          const filePath = path.join(directory, filename);
          const jsonData = fs.readFileSync(filePath, 'utf-8');
          const depositData = JSON.parse(jsonData);
      
          const pubkey = '0x' + depositData[0].pubkey;
          const withdrawalCredentials = '0x' + depositData[0].withdrawal_credentials;
          const signature = '0x' + depositData[0].signature;
          const depositDataRoot = '0x' + depositData[0].deposit_data_root;
      
          console.log(`File: ${filename}`);
          console.log(`pubkey: ${pubkey}`);
          console.log(`withdrawalCredentials: ${withdrawalCredentials}`);
          console.log(`signature: ${signature}`);
          console.log(`depositDataRoot: ${depositDataRoot}`);
      
          let txData = await contract.methods.deposit(pubkey, withdrawalCredentials, signature, depositDataRoot).encodeABI();
      
          const value = web3.utils.toWei('32', 'ether');
          const txObject = {
            to: contractAddress,
            data: txData,
            value: value,
            gas: 2000000, // Adjust the gas limit as needed
          };
      
          account.signTransaction(txObject)
            .then(signedTx => {
              web3.eth.sendSignedTransaction(signedTx.rawTransaction)
                .on('transactionHash', hash => {
                  console.log(`Transaction Hash: ${hash}`)
                })
                .on('receipt', receipt => {
                  console.log('Transaction Receipt:', receipt);
                })
                .on('error', err => {
                  console.error('Transaction Error:', err);
                });
            })
            .catch(error => {
              console.error('Transaction Signing Error:', error);
            });
        });
      });
  })
