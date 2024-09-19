import json
from eth_keys import keys
from eth_account import Account

file_path = "./keystore/UTC--2023-09-28T11-51-18.436328606Z--77807cc4c75ef540f0db4c741b5e7774b7678ad3.json"
# Load the UTC (JSON) file
with open("/Node-1/execution/keystore/UTC--2023-11-01T06-41-48.000446396Z--abe3f423136650def254f8467d294114e91818b9", "r") as file:
    utc_data = json.load(file)

# Prompt for the password to decrypt the UTC file
password = ""

# Decrypt the UTC file using the password
private_key = Account.decrypt(utc_data, password)

# Print the private key in hexadecimal format
print("Private Key:", private_key.hex())
