import React,{useEffect,useState} from 'react'
// import {Data} from "/home/rasin/go-Apis/backend/staking-deposit-cli/validator_keys/validator_keys/deposit_data-1693303071.json"
import {Button} from "@material-ui/core"
// import Data from "../validator_keys/validator_keys/deposit_data.json"
import { ethers } from 'ethers';


const contractABI = [
  {
    "inputs": [{
        "internalType": "bytes",
        "name": "pubkey",
        "type": "bytes"
    }, {
        "internalType": "bytes",
        "name": "withdrawal_credentials",
        "type": "bytes"
    }, {
        "internalType": "bytes",
        "name": "signature",
        "type": "bytes"
    }, {
        "internalType": "bytes32",
        "name": "deposit_data_root",
        "type": "bytes32"
    }],
    "name": "deposit",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
}
    ]

    
    

const contractAddress = '0x4242424242424242424242424242424242424242';

const Validator = () => {

    const [provider, setProvider] = useState(null);
    const [contract, setContract] = useState(null);
    const [account,setAccount] = useState(null)

    const [data, setData] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
  
      
    useEffect(() => {
        const initialize = async () => {
            if (typeof window.ethereum !== "undefined") {
              try {
                const accounts = await window.ethereum.request({ method: "eth_requestAccounts" });
                const ethereumProvider = new ethers.providers.Web3Provider(window.ethereum);
                setAccount(accounts[0]);
                setProvider(ethereumProvider);
        
                window.ethereum.on('chainChanged', (chainId) => {
                  window.location.reload();
                });
        
                window.ethereum.on('accountsChanged', async function (accounts) {
                  setAccount(accounts[0]);
                  await connectWeb();
                });
        
                const chainId = await ethereumProvider.getSigner().getChainId();
                if (chainId !== 32382) {
                  alert("Please change your network to Localhost hardhat");
                }
        
                const contractInstance = new ethers.Contract(contractAddress, contractABI, ethereumProvider.getSigner());
                setContract(contractInstance);
              } catch (error) {
                console.error("Error initializing:", error);
              }
            }
            
          };


        
          initialize();
          console.log("account ",account)

          const apiUrl = 'http://10.101.245.13:3500/eth/v1/node/identity';

          // Make the GET request
          fetch(apiUrl)
            .then((response) => {
              if (!response.ok) {
                throw new Error('Network response was not ok');
              }
              return response.json();
            })
            .then((json) => {
              console.log("json",json["data"]["p2p_addresses"][0])
              setData(json);
              setLoading(false);
            })
            .catch((error) => {
              setError(error);
              setLoading(false);
            });

           }, [account]);
    
      const connectWeb = async () => {
        if (provider) {
          const signer = provider.getSigner();
          const chainId = await signer.getChainId();
          if (chainId === 32382) {
            const contractInstance = new ethers.Contract(contractAddress, contractABI, signer);
            setContract(contractInstance);
          } else {
            alert("Please change your network to Localhost hardhat");
          }
        }
      };
    
  
    // Function to interact with the contract
    const depositToContract = async (e) => {
        e.preventDefault()
      if (contract) {
        const signer = provider.getSigner();
        
        try {
     
            const tx = await contract.deposit(
              "0x913c27ae930472680fcf59243a097c8447f25da8656177f4732e1308a33304143c0278c107e72b27560817ea8f149e04",
              "0x00e25dd36097cfb360bfaf1e6ebc27454489ce32740aebdc6782c0633486f2ae",
              "0xa071b515580e2b96e5ac121e67809ddd18e1471e0d639d593a3eedbf945099c13f8a33ec6b9769f4f5a8aad8f05432e109dcb6c6a174c5cf2a99646d278d795bf3d70a35e266baf6dd6a5afab329f34dd85e9fb51a181b119ee0dac30eef3836",
              "0x42559b9a39c53b9350635bbc3ed428121e2b6980db617bc585b370b581d08663",
              { gasLimit: 200000, value: ethers.utils.parseEther("32.0") }
            );
          // Wait for the transaction to be mined
          const receipt = await tx.wait();
    
          if (receipt.status === 1) {
            console.log('Transaction confirmed:', receipt);
          } else {
            console.error('Transaction failed:', receipt);
          }
        } catch (error) {
          console.error('Transaction error:', error);
        }
      }
    }
  

    
    return (
            <div>
    {/*             
                {Data.forEach(element => {
                    return (
                    <h1>element.pubkey</h1>
                    <h1>element.withdrawal_credentials</h1>
                    <h1>element.signature</h1>
                    <h1>element.deposit_data_root</h1>
                    )})} */}

            
            {/* <div>
            {Data.map((element, index) => (
                <div key={index}>
                    <p>Pubkey: {element.pubkey}</p>
                    <p>Withdrawal Credentials: {element.withdrawal_credentials}</p>
                    <p>signature: {element.signature}</p>
                    <p>deposit_data_root: {element.deposit_data_root}</p>
                </div>
            ))} 
        </div>*/}

        <div className="App">
           <p>{account}</p>
            <Button 
            color='primary' 
            variant="contained"
            onClick={depositToContract}>Deposit to Contract</Button>
           
        </div>

            </div>
    )
}

export default Validator