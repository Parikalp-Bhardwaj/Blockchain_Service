import React, { useEffect, useState } from "react";
import { Button } from "@material-ui/core";
import { useHistory } from "react-router-dom";
import Grid from "@mui/material/Grid";
import DeleteOutlineIcon from "@mui/icons-material/DeleteOutline";
import TextField from "@mui/material/TextField";
import CircularProgress from "@mui/material/CircularProgress";
import AddCircleOutlineIcon from "@mui/icons-material/AddCircleOutline";
import Web3 from "web3";

import PropTypes from "prop-types";
import LinearProgress from "@mui/material/LinearProgress";
import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import IconButton from '@mui/material/IconButton';
import Alert from '@mui/material/Alert';
import CloseIcon from '@mui/icons-material/Close';
import Collapse from '@mui/material/Collapse';
import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
} from '@material-ui/core';
import { Close } from '@material-ui/icons';

// import {ConfirmDialog} from "./ConfirDialog"

function LinearProgressWithLabel(props) {
  return (
    <Box sx={{ display: "flex", alignItems: "center" }}>
      <Box sx={{ width: "100%", mr: 1 }}>
        <LinearProgress variant="determinate" {...props} />
      </Box>
      <Box sx={{ minWidth: 35 }}>
        <Typography variant="body2" color="text.secondary">{`${Math.round(
          props.value
        )}%`}</Typography>
      </Box>
    </Box>
  );
}

LinearProgressWithLabel.propTypes = {
  /**
   * The value of the progress indicator for the determinate and buffer variants.
   * Value between 0 and 100.
   */
  value: PropTypes.number.isRequired,
};



// export const ConfirmDialog = () => {
//   return (
//     <Dialog open={true} maxWidth="sm" fullWidth>
//       <DialogTitle>Confirm the action</DialogTitle>
//       <Box position="absolute" top={0} right={0}>
//         <IconButton>
//           <Close />
//         </IconButton>
//       </Box>
//       <DialogContent>
//         <Typography>Blockchain</Typography>
//       </DialogContent>
//       <DialogActions>
//         <Button color="secondary" variant="contained">
//           Confirm
//         </Button>
//       </DialogActions>
//     </Dialog>
//   );
// };



const BLOCK_BATCH = 1;
let FROM_BLOCK = 10;
let TO_BLOCK = 20;
const web3 = new Web3("http://192.168.253.100:8545/");

const Index = () => {

  let history = useHistory();
  const [chainid, setChainId] = useState(null);

  const [gaslimit, setGasLimit] = useState("");

  const [secondPerSlot, setSecondPerSlot] = useState(null);
  const [slotPerEpoch, setSlotPerEpoch] = useState(null);
  const [count, setCount] = useState(1);

  const [blockInfo, setBlockInfo] = useState(0);

  const [isLoading, setIsLoading] = useState(false);
  const [radiobutton, setRadioButton] = useState(false);
  const [blockNum, setBlockNum] = useState(0);
  const [totalTransaction, setTotalTransaction] = useState(0);
  const [valTransaction, setValTransaction] = useState(null);

  const [addresses, setAddresses] = useState(Array(count).fill(null));
  const [balances, setBalances] = useState(Array(count).fill(null));
  const [statusData, setStatusData] = useState();

  // const [progress, setProgress] = React.useState(20);
  const [isButtonDisabled, setIsButtonDisabled] = useState(false);
  const [isButtonDisabled2, setIsButtonDisabled2] = useState(true);

  const [progress, setProgress] = useState(0);

  const [issuccessfull, setSuccessfull] = useState(false);
  const [isthroughput, setThroughput] = useState(0);
  const [numValidator, setGetValidator] = useState(0)
  const [blockchainUp,setBlockchainUp] = useState(false)


  // function CircularProgressWithLabel(props) {
  //   return (
  //     <Box sx={{ position: "relative", display: "inline-flex" }}>
  //       <CircularProgress variant="determinate" {...props} />
  //       <Box
  //         sx={{
  //           top: 0,
  //           left: 0,
  //           bottom: 0,
  //           right: 0,
  //           position: "absolute",
  //           display: "flex",
  //           alignItems: "center",
  //           justifyContent: "center",
  //         }}
  //       >
  //         <Typography variant="caption" component="div" color="text.secondary">
  //           {`${Math.round(props.value)}%`}
  //         </Typography>
  //       </Box>
  //     </Box>
  //   );
  // }

  const updateAddressAndBalance = (index, newAddress, newBalance) => {
    console.log("newBalance ", newBalance, typeof (newBalance))
    setData((prevData) =>
      prevData.map((row, i) =>
        i === index ? { ...row, address: newAddress, balance: newBalance } : row
      )
    );
  };


  useEffect(() => {

    const fetchData = async () => {
      try {
        const response = await fetch(
          `http://192.168.253.110:5000/block/${secondPerSlot}`
        );
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        if (response.status === 200)  {
          const data = await response.json();

          if(data["getValidator"] != undefined)
            setGetValidator(data["getValidator"])
          
          if(data["data"] != null)
            setBlockInfo(data["data"]);

          if(data["blockNum"] != undefined)
            setBlockNum(data["blockNum"]);

          setTotalTransaction(data["toalTransaction"]);
          setBlockchainUp(data["UpBlockchain"])

          console.log("data ----->", data["data"], data["blockNum"], data["getValidator"],data["UpBlockchain"]);


        } else {
          setBlockInfo({ error: "Block not found" });
        }
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };
    console.log("ValTransaction ", valTransaction);

    const getThroughput = async () => {
      console.log("getting throughput")
      try {
        const response = await fetch(
          `http://192.168.253.109:8081/throughput`
        );

        if (response.status === 200) {
          const data = await response.json();
          setThroughput(data["Throughput"]*3)
          console.log("throughput--> ", data*3)
        }

        // const unicodeDecimalNumber = parseInt(hexadecimalString, 16);
        // const unicodeCharacter = String.fromCodePoint(unicodeDecimalNumber);
        // console.log("Unicode Character:", unicodeCharacter);


        // } else {
        //   setBlockInfo({ error: "Block not found" });
        // }

      } catch (error) {
        console.log("Error fetching data:", error)
      }
    }

    // const getValidator = async () => {
    //   try {
    //     const response = await fetch(
    //       `http://192.168.253.108:3500/eth/v1/beacon/states/head/validators?status=active_ongoing`
    //     );
    //     if (!response.ok) {
    //       throw new Error("Network response was not ok");
    //     }
    //     if (response.status === 200) {
    //       const data = await response.json();

    //       console.log("num of Validator ", data["data"].length)
    //       setGetValidator(data["data"].length)


    //       // const unicodeDecimalNumber = parseInt(hexadecimalString, 16);
    //       // const unicodeCharacter = String.fromCodePoint(unicodeDecimalNumber);
    //       // console.log("Unicode Character:", unicodeCharacter);


    //     } else {
    //       setBlockInfo({ error: "Block not found" });
    //     }

    //   } catch (error) {
    //     console.log("Error fetching data:", error)
    //   }
    // }

    fetchData();
    // getValidator()
    getThroughput()
    if (blockchainUp === true) {
      setRadioButton(true);
      setIsButtonDisabled(true)
      setIsButtonDisabled2(false)
      // setRadioButton(true)
    }else if ((blockchainUp === false)){
      // setRadioButton(false);
      setIsButtonDisabled(false)
      setIsButtonDisabled2(true)
      // setRadioButton(false)
    }
    // || parseInt(numValidator) > 1
    // Set up an interval to fetch data every 5 seconds
    const interval = setInterval(() => {
      console.log("numValidator blockNum----------- ",numValidator,blockNum)
      if (blockchainUp === true )  {
        setRadioButton(true);
        setIsButtonDisabled(true)
        setIsButtonDisabled2(false)
        // setRadioButton(true)
      }else if (blockchainUp === false){
        // setRadioButton(false);
        setIsButtonDisabled(false)
        setIsButtonDisabled2(true)
        // setRadioButton(false)
      }

      // getValidator()
      fetchData();
      getThroughput()
    }, 2000);

    return () => clearInterval(interval);


  }, [secondPerSlot, valTransaction, isthroughput, blockNum, blockchainUp,radiobutton]);




  // "proxy": "http://127.0.0.1:8000",

  useEffect(() => {
    const fetchData2 = async () => {
      try {
        const response = await fetch("/status");
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        if (response.status === 200) {
          const data = await response.json();
          setStatusData(data["Status"]);

          console.log("data Stat ->", data["Status"]);
        } else {
          setBlockInfo({ error: "Block not found" });
        }
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };
    console.log("ValTransaction ", valTransaction);
    // Initially, fetch data
    fetchData2();

    console.log("isButtonDisabled ", isButtonDisabled)
    console.log("isButtonDisabled2 ", isButtonDisabled2)


    const intervalId = setInterval(() => {
      fetchData2(); // Fetch data periodically, e.g., every 5 seconds
    }, 3000); // Adjust the interval as needed (in milliseconds)

    if (statusData === "Initializing...." && progress === 0) {
      setProgress((prevProgress) => prevProgress + 15);
    } else if (statusData === "Genesis Created Config Created" && progress === 15) {
      setProgress((prevProgress) => prevProgress + 15);
    } else if (statusData === "Node 1 is Up and Running" && progress === 30) {
      setProgress((prevProgress) => prevProgress + 15);
    } else if (statusData === "Node 2 is Up and Running" && progress === 45) {
      setProgress((prevProgress) => prevProgress + 15);
    } else if (statusData === "Node 3 is Up and Running" && progress === 60) {
      setProgress((prevProgress) => prevProgress + 10);
    } 
    // else if (statusData === "Node 4 is Up and Running" && progress === 70) {
    //   setProgress((prevProgress) => prevProgress + 10);
    // } 
    else if (statusData === "Staking Deposit Processing..." && progress === 70) {
      setProgress((prevProgress) => prevProgress + 20);
    } else if (statusData === "Blockchain Launched" && progress === 90) {
      setProgress((prevProgress) => prevProgress + 10);
    } else if (statusData === "Blockchain Launched" && progress === 100) {
      setSuccessfull(true);
    }

    return () => {
      // Clean up the interval when the component unmounts
      clearInterval(intervalId);
    };




  }, [statusData, progress]);

  const callApi = async (e) => {
    e.preventDefault();
    setIsButtonDisabled(true)


    const alloc = {};

   
    data.forEach((row) => {
   
      const { address, balance } = row;
      const stringWithout0x = address.replace(/^0x/, '');
      console.log("stringWithout0x ", stringWithout0x)
 
      const weiValue = Web3.utils.toWei(balance, 'ether');
      console.log("weiValue gaslimit", weiValue, gaslimit)

      // Replace with your decimal number





      if (stringWithout0x) {
        // Assign the balance to the 'address' key in 'alloc'
        alloc[stringWithout0x] = {
          balance: weiValue.toString(),
        };
      }
    });

    setIsLoading(true);
    try {
      const response = await fetch("/createnetwork", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          SECONDS_PER_SLOT: parseInt(secondPerSlot),
          SLOTS_PER_EPOCH: parseInt(slotPerEpoch),
          config: {
            chainId: parseInt(chainid),
          },
          alloc: alloc,
          gasLimit: gaslimit.toString(),
        }),
      });

      if (!response.ok) {
        console.log("error ..");
        throw new Error("Network response was not ok");
      }

      if (response.status === 200) {
        const rawData = await response.text();
        const firstJsonObject = JSON.parse(rawData.substring(0, rawData.indexOf('}') + 1));
        const secondJsonObject = JSON.parse(rawData.substring(rawData.lastIndexOf('{')));
        const newtworkNotStarted = secondJsonObject["networkNotStarted"]

        console.log("Data from API:", rawData);
        console.log("First JSON object:", firstJsonObject);
        console.log("Second JSON object:", secondJsonObject["networkNotStarted"]);
        // setIsLoading(true);
        setIsButtonDisabled2(newtworkNotStarted)

      }


   
      setRadioButton(true);
    } catch (error) {
      console.error("Fetch error:", error.message);
    } finally {
      setIsLoading(false);
    }
  };

  const downBlockchain = async (e) => {
    e.preventDefault();
    // setIsButtonDisabled2(true)
    // setIsButtonDisabled(false)
    try {
      const response = await fetch("/delete") // Replace with your API URL

      if (!response.ok) {
        console.log("error ..");
        throw new Error("Network response was not ok");
      }

      const data = await response.json()
      console.log("data delete ", data);


      setIsButtonDisabled(data["networkNotStarted"])
      setIsLoading(false);
      setRadioButton(false);
      setProgress(0)
    } catch (error) {
      // setError(err);
      setIsLoading(false);
      setRadioButton(false);
      console.error("Fetch error:", error);
    }

  };

  const NumberOfTransaction = async (e) => {
    e.preventDefault();
    console.log("getting num");
    const each = Math.floor(valTransaction / 3); // Get the base amount for each RPC
    const remainder = valTransaction % 3; // Get the remainder after dividing by 3
    // Distribute the base amount and handle the remainder
    const rpc1 = each + (remainder >= 1 ? 1 : 0);
    const rpc2 = each + (remainder >= 2 ? 1 : 0);
    const rpc3 = each;
    const data1 = {
      "count": rpc1,
      "rpcs": ["http://192.168.253.108:8545"],
      "chainid": parseInt(chainid)
    };

    const data2 = {
      "count": rpc2,
      "rpcs": ["http://192.168.253.106:8545"],
      "chainid": parseInt(chainid)
    };

    const data3 = {
      "count": rpc3,
      "rpcs": ["http://192.168.253.107:8545"],
      "chainid": parseInt(chainid)
    };

    console.log("data start*************", data2);

    

    try {
      const response1 = await fetch(
        "http://192.168.253.109:8081/start",
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(data1),
        });


      const res = await response1.json();
      setBlockInfo(res);
      console.log("res ", res);


    } catch (error) {
      console.error("Error fetching block data:", error.msg);
      if (error.response) {
        console.error("Response status:", error.response.status);
        console.error("Response body:", error.response.body);
      }
    }

    try {
      const response2 = await fetch(
        "http://192.168.253.102:8081/start",
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(data2),
        });


      const res = await response2.json()
      console.log("res ", res);


    } catch (error) {
      console.error("Error fetching block data:", error.msg);
      if (error.response) {
        console.error("Response status:", error.response.status);
        console.error("Response body:", error.response.body);
      }
    }

    try {
      const response3 = await fetch(
        "http://192.168.253.103:8081/start",
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(data3),
        });


      const res = await response3.json()
      console.log("res ", res);


    } catch (error) {
      console.error("Error fetching block data:", error.msg);
      if (error.response) {
        console.error("Response status:", error.response.status);
        console.error("Response body:", error.response.body);
      }
    }

    //  const data4 = {
    //   "count": parseInt(valTransaction),
    //   "rpcs": ["http://192.168.253.106:8545"],
    //   "chainid": parseInt(chainid)
    // };


    // try {
    //   const response3 = await fetch(
    //     "/start",
    //     {
    //       method: "POST",
    //       headers: { "Content-Type": "application/json" },
    //       body: JSON.stringify(data4),
    //       });

          
    //       const res = await response3.json();
    //         // setBlockInfo(res);
    //       console.log("res ", res);
      

    //     } catch (error) {
    //       console.error("Error fetching block data:", error.msg);
    //       if (error.response) {
    //         console.error("Response status:", error.response.status);
    //         console.error("Response body:", error.response.body);
    //       }
    //     }


  }


  const fundBalance = async (e) => {
    e.preventDefault();
    let rpc = "http://192.168.253.108:8545".toString()
    const data = {
      "rpc": "http://192.168.253.108:8545".toString(),
      "chainid": parseInt(chainid),
      "amount": parseInt(10000),
    }
    console.log("data ", data)
    console.log("rpc", rpc)
    try {
      const response = await fetch("/fund", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
      }
      );

      if (!response.ok) {
        console.log("error .. ", response);
        throw new Error("Network response was not ok");
      }

      const responseData = await response.json(); // Parse the response data as JSON
      console.log("Response Data:", responseData);

    } catch (error) {
      console.error("Error fetching block data:", error.msg);
      if (error.response) {
        console.error("Response status:", error.response.status);
        console.error("Response body:", error.response.body);
      }
    }

  }


  const stopFunding = async (e) => {
    e.preventDefault();
    const apiUrlList = [
      "http://192.168.253.109:8081/stop",
      "http://192.168.253.102:8081/stop",
      "http://192.168.253.103:8081/stop"
    ];

    const fetchPromises = apiUrlList.map(async (url) => {
      try {
        const response = await fetch(url, {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({}),
        });

        const res = await response.json();
        return res;
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    });

    try {
      const responseDataList = await Promise.all(fetchPromises);
    } catch (error) {
      console.error('Error fetching data:', error);
    }

  }

  const [data, setData] = useState([]);

  const addNewRow = () => {
    const newRow = {
      id: data.length + 1,
      address: "",
      balance: "",
      name: `Row ${data.length + 1}`,
    };
    setCount(count + 1);
    setData((prevData) => [...prevData, newRow]);
    console.log("count ", count);
  };

  const removeRow = (id) => {
    const updatedData = data.filter((row) => row.id !== id);
    setData(updatedData);
    setCount(count - 1);
    console.log("count ", count);
  };

  const [open, setOpen] = React.useState(true);

  return (
    <div className="" style={{backgroundColor: '#425361'}}>
      {isLoading ? (


        <div>

          {issuccessfull ? (<div className="blockChainDone">
            <Collapse in={open}>
              <Alert
                action={
                  <IconButton
                    aria-label="close"
                    color="inherit"
                    size="small"
                    onClick={() => {
                      setOpen(false);
                    }}
                  >
                    <CloseIcon fontSize="inherit" />
                  </IconButton>
                }
                sx={{ mb: 2 }}
              >
                Blockchain created successfully.
              </Alert>
            </Collapse>
          </div>

          )


            : (<></>)}


          <div className="justify-center items-center m-10 mainProgressBar">
            <h2>Creating Blockchain</h2>
            <Box sx={{ width: "100%" }}>
              <LinearProgressWithLabel value={progress} />
            </Box>
            <h3><span>{statusData}</span></h3>
          </div>
        </div>

      ) : (




        <div className="container mainContainer">
          <Grid container spacing={2}>
            <Grid item sm={12} md={6}>
              <div className="bg-gray-100 p-4 configWrap">
                <div className="bg-white rounded-lg shadow-md p-4">
                  <Grid container spacing={2}>
                    <Grid item md={12}>
                      <h3 className="text-lg font-medium  text-left">
                        Configuration
                      </h3>
                    </Grid>
                    <Grid item sm={12} md={6}>
                      <TextField
                        id="chain-id"
                        value={chainid}
                        //   placeholder='Chain Id'
                        onChange={(e) => setChainId(e.target.value)}
                        className="inputMain"
                        label="Chain Id"
                        variant="outlined"
                      />
                    </Grid>

                    <Grid item sm={12} md={6}>
                      <TextField
                        id="gas-limit"
                        value={gaslimit}
                        // placeholder='Gas limit'
                        onChange={(e) => setGasLimit(e.target.value)}
                        className="inputMain"
                        label="Gas limit"
                        variant="outlined"
                      />
                    </Grid>

                    <Grid item sm={12} md={6}>
                      <TextField
                        id="second-per-slot"
                        value={secondPerSlot}
                        // placeholder="Second Per Slot"
                        onChange={(e) => setSecondPerSlot(e.target.value)}
                        className="inputMain"
                        label="Second Per Slot"
                        variant="outlined"
                      />
                    </Grid>

                    <Grid item sm={12} md={6}>
                      <TextField
                        id="slot-per-epoch"
                        value={slotPerEpoch}
                        //    placeholder="Slot Per Epoch"
                        className="inputMain"
                        onChange={(e) => setSlotPerEpoch(e.target.value)}
                        label="Slot Per Epoch"
                        variant="outlined"
                      />
                    </Grid>

                    <Grid item xs="12" md="12">
                      <Button
                        className="addNewRow"
                        endIcon={<AddCircleOutlineIcon />}
                        onClick={addNewRow}
                      >
                        Add New Address
                      </Button>
                    </Grid>
                  </Grid>

                  <div className="addressWrap">

                    <Grid conatiner spacing={3}>
                      <Grid item md={12}>
                        {data.map((row, index) => (
                          <div className="newAddressItem" key={row.id}>
                            <Grid container spacing={2}>
                              <Grid item xs={5}>
                                <TextField
                                  id={`address-${row.id}`}
                                  value={row.address}
                                  className="inputMain"
                                  label={`Address ${row.id}`}
                                  variant="outlined"
                                  onChange={(e) =>
                                    updateAddressAndBalance(index, e.target.value, row.balance)
                                  }
                                />
                              </Grid>

                              <Grid item xs={5}>
                                <TextField
                                  id={`balance-${row.id}`}
                                  value={row.balance}
                                  className="inputMain"
                                  label={`Balance ${row.id}`}
                                  variant="outlined"
                                  onChange={(e) =>
                                    updateAddressAndBalance(index, row.address, e.target.value)
                                  }
                                />
                              </Grid>
                              <Grid item xs={2}>
                                <Button
                                  spacing={5}
                                  variant="outlined"
                                  className="removeRowInput"
                                  endIcon={<DeleteOutlineIcon />}
                                  onClick={() => removeRow(row.id)}
                                >
                                </Button>
                              </Grid>
                            </Grid>
                          </div>
                        ))}

                      </Grid>
                    </Grid>

                  </div>

                  <div className="m-3">
                    <Grid item xs={12}>
                      <Button
                        color="primary"
                        variant="contained"
                        m={1}
                        onClick={callApi}
                        className="mt-5"
                        disabled={isButtonDisabled}
                      >
                        Create Blockchain
                      </Button>
                    </Grid>
                  </div>
                </div>
              </div>
            </Grid>

            <Grid item sm={12} md={6}>
              <div className="bg-gray-100 p-4 statusSide">
                <div className="bg-white rounded-lg shadow-md p-6">
                  <div className="mt-3 flex justify-between">
                    <h3 className="text-xl font-semibold text-gray-900 dark:text-black">
                      Status
                    </h3>
                    {radiobutton ? (
                      <div className="flex items-center text-green-500">
                        <span className="w-4 h-4 rounded-full bg-green-500 mr-2"></span>
                        <span className="font-semibold">Running</span>
                      </div>
                    ) : (
                      <div className="flex items-center text-red-500">
                        <span className="w-4 h-4 rounded-full bg-red-500 mr-2"></span>
                        <span className="font-semibold">Stopped</span>
                      </div>
                    )}
                    <Button
                      color="secondary"
                      variant="contained"
                      m={1}
                      onClick={downBlockchain}
                      disabled={isButtonDisabled2}

                    >
                      Down Blockchain
                    </Button>
                  </div>

                  <div className="blockCountWrap">
                    <Grid container spacing={2}>
                      <Grid item md={6}>
                        <div className="blockCount">
                          <h5 className="text-l font-semibold text-gray-900 dark:text-black m-4">
                            Tx Per Sec
                          </h5>
                          <h3 className="text-xl font-bold text-blue-500 dark:text-blue-400  m-4">
                            {blockInfo}
                          </h3>
                        </div>
                      </Grid>
                      <Grid item md={6}>
                        <div className="blockCount">
                          <h5 className="text-l font-semibold text-gray-900 dark:text-black m-4">
                            Block Number
                          </h5>
                          <h3 className="text-xl font-bold text-blue-500 dark:text-blue-400  m-4">
                            {blockNum}
                          </h3>
                        </div>
                      </Grid>
                      <Grid item md={6}>
                        <div className="blockCount">
                          <h5 className="text-l font-semibold text-gray-900 dark:text-black m-4">
                            Total Transaction
                          </h5>
                          <h3 className="text-xl font-bold text-blue-500 dark:text-blue-400  m-4">
                            {parseInt(isthroughput)}
                          </h3>
                        </div>
                      </Grid>
                      <Grid item md={6}>
                        <div className="blockCount">
                          <h5 className="text-l font-semibold text-gray-900 dark:text-black m-4">
                            Validators
                          </h5>
                          <h3 className="text-xl font-bold text-blue-500 dark:text-blue-400  m-4">
                            {numValidator}
                          </h3>
                        </div>
                      </Grid>
                    </Grid>
                  </div>
                </div>

                <div className="bg-white rounded-lg shadow-md p-6 mt-4 md:mt-8">
                  <h3 className="text-xl font-semibold text-gray-900 dark:text-black text-left">
                    Test Console
                  </h3>

                  <div className="testConsoleWrap">
                    <h3>Number of transaction:</h3>

                    <div className="consoleForm">
                      <input
                        placeholder="Enter number"
                        value={valTransaction}
                        onChange={(e) => setValTransaction(e.target.value)}
                      />
                      <Button
                        style={{
                          backgroundColor: "#4AD246"
                        }}
                        color="primary"
                        variant="contained"
                        m={1}
                        onClick={NumberOfTransaction}
                      >
                        Start
                      </Button>
                      {/* <Button
                        style={{
                          backgroundColor: "#2C78CE"
                        }}
                        color="secondary"
                        variant="contained"
                        m={1}
                        onClick={fundBalance}
                      >
                        Fund
                      </Button> */}
                      <Button
                        color="secondary"
                        variant="contained"
                        m={1}
                        onClick={stopFunding}
                      >
                        Stop
                      </Button>
                    </div>
                  </div>
                </div>
              </div>
            </Grid>
          </Grid>
        </div>
      )}
    </div>
  );
};

export default Index;
