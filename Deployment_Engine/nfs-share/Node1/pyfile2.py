from web3 import Web3
from eth_account import Account
import os
import json 

# Replace these values with your actual contract address and ABI
contract_address = '0x4242424242424242424242424242424242424242'
contract_abi = [

  {
    "constant": False,
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
    "payable": True,
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "constant": True,
    "inputs": [],
    "name": "get_deposit_root",
    "outputs": [
      {
        "name": "",
        "type": "bytes32"
      }
    ],
    "payable": False,
    "stateMutability": "view",
    "type": "function"
  },
  {
    "constant": True,
    "inputs": [],
    "name": "get_deposit_count",
    "outputs": [
      {
        "name": "",
        "type": "bytes"
      }
    ],
    "payable": True,
    "stateMutability": "view",
    "type": "function"
  }
]


# Replace these with your Ethereum RPC URL and private key
ethereum_rpc_url = 'http://192.168.253.108:8545'
private_key = 'f29faca0fe8bc2cdb89333fc05475e842a57b9f3be5c2c6a15648619d5b4daac'


w3 = Web3(Web3.HTTPProvider(ethereum_rpc_url))


account = Account.from_key(private_key)

w3.eth.default_account = account.address


directory = './validator_keys/'


files = os.listdir(directory)

deposit_data_files = [file for file in files if file.startswith('deposit_data-')]

# Process each deposit data file
for filename in deposit_data_files:
    file_path = os.path.join(directory, filename)
    

    with open(file_path, 'r') as json_file:
        deposit_data = json.load(json_file)
    

    pubkey = deposit_data[0]['pubkey']
    withdrawal_credentials = deposit_data[0]['withdrawal_credentials']
    signature = deposit_data[0]['signature']
    deposit_data_root = deposit_data[0]['deposit_data_root']




amount_eth = 32
amount_wei = w3.toWei(amount_eth, 'ether') 

contract = w3.eth.contract(address=contract_address, abi=contract_abi)

tx_data = contract.functions.deposit(pubkey, withdrawal_credentials, signature, deposit_data_root).buildTransaction({
    'chainId': 32304,  
    'gas': 2000000,
    'gasPrice': w3.toWei('5', 'gwei'),  
    # 'nonce': w3.eth.getTransactionCount(account.address),
    'value': amount_wei
})


signed_tx = w3.eth.account.signTransaction(tx_data, private_key)


tx_hash = w3.eth.sendRawTransaction(signed_tx.rawTransaction)

print(f'Transaction Hash: {tx_hash.hex()}')

