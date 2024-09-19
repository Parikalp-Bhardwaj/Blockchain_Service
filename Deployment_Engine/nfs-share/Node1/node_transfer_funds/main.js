const Web3 = require('web3');
const fs = require("fs");
const path = require("path");
// Replace these values with your actual contract address and ABI
const contractAddress = 'YOUR_CONTRACT_ADDRESS';
const contractABI = [
  // Include the ABI of your contract here
  // Example:
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
];

// Replace these with your Ethereum RPC URL and private key
const web3 = new Web3('YOUR_ETHEREUM_RPC_URL');
const privateKey = 'YOUR_PRIVATE_KEY';

// Create an Ethereum account from the private key
const account = web3.eth.accounts.privateKeyToAccount(privateKey);

// Set the default account to your Ethereum account
web3.eth.accounts.wallet.add(account);
web3.eth.defaultAccount = account.address;

// Create a contract instance
const contract = new web3.eth.Contract(contractABI, contractAddress);

const directory = '';

// List all files in the directory
fs.readdir(directory, (err, files) => {
  if (err) {
    console.error('Error reading directory:', err);
    return;
  }

  // Filter files with the prefix "deposit_data-"
  const depositDataFiles = files.filter(file => file.startsWith('deposit_data-'));

  // Process each deposit data file
  depositDataFiles.forEach(filename => {
    const filePath = path.join(directory, filename);
    const jsonData = fs.readFileSync(filePath, 'utf-8');
    const depositData = JSON.parse(jsonData);

    // Extract the values from depositData and use them as needed
    const pubkey = depositData.pubkey;
    const withdrawalCredentials = depositData.withdrawalCredentials;
    const signature = depositData.signature;
    const depositDataRoot = depositData.depositDataRoot;

    console.log(`File: ${filename}`);
    console.log(`pubkey: ${pubkey}`);
    console.log(`withdrawalCredentials: ${withdrawalCredentials}`);
    console.log(`signature: ${signature}`);
    console.log(`depositDataRoot: ${depositDataRoot}`);
  });
});

// Define the transaction data
const txData = contract.methods.deposit(pubkey, withdrawalCredentials, signature, depositDataRoot).encodeABI();

// Replace with the amount you want to send (in Wei)
const value = web3.utils.toWei('32', 'ether');

// Build the transaction object
const txObject = {
  to: contractAddress,
  data: txData,
  value: value,
  gas: 2000000, // Adjust the gas limit as needed
};

// Sign and send the transaction
web3.eth.accounts.signTransaction(txObject, privateKey)
  .then(signedTx => {
    web3.eth.sendSignedTransaction(signedTx.rawTransaction)
      .on('transactionHash', hash => {
        console.log(`Transaction Hash: ${hash}`);
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

