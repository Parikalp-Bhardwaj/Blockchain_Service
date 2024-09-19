#!/usr/bin/expect -f

# sudo apt-get install xclip

# Move to the staking-deposit-cli directory
spawn mkdir -p /shared-files/validator_keys
expect eof
cd /shared-files/staking-deposit-cli


# Upgrade setuptools and wheel
spawn pip install --upgrade setuptools wheel
expect eof

# Install Python dependencies
spawn pip3 install -r requirements.txt
expect eof

# Install Python package
spawn python3 setup.py install
expect eof

# Run deposit.sh script
# If you are getting error on related this bash: ./deposit.sh: /bin/bash^M: bad interpreter: No such file or directory
# Run these command
# 1 sudo apt install dos2unix
# 2 dos2unix deposit.sh
spawn ./deposit.sh install
expect eof

# Run deposit.py script with specified parameters
spawn python3 ./staking_deposit/deposit.py new-mnemonic --num_validators=1 --mnemonic_language=english --chain=devnet --folder=/shared-files/


set lang ""
send "$lang\r"

set password "12345678"
expect "Password:"
send "$password\r"
expect "Repeat password:"
send "$password\r"

sleep 5
# Prompt the user to press any key
send_user "Press any key.\n"
expect_user -re {.*}
send "\r"


sleep 5
spawn chmod 777 -R /shared-files/validator_keys/
expect eof

spawn mv /shared-files/validator_keys /Node-4/
expect eof
