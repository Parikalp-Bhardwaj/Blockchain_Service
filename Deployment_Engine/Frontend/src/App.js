
import React,{useState} from "react"
import './App.css';
import Index from './components/Index';
import Config from './components/Config';
import Validator from './components/Validator';
// import { Route, BrowserRouter, Switch,Routes  } from 'react-router-dom';
import ServerConfig from "./components/ServerConfig";
import { Button, Container, Typography } from '@mui/material';
import { styled } from '@mui/system';

function App() {
  const [currentSlide, setCurrentSlide] = useState(1);

  const nextSlide = () => {
    setCurrentSlide((prevSlide) => (prevSlide === 1 ? 2 : 1));
  };



  return (
    // <BrowserRouter>
    // <div className="App">
    
    //   <Switch>
    //     <Route exact path="/" component={Config} />
    //     <Route path="/validator" component={Validator} />
   
    //   </Switch>

    // </div>
    // </BrowserRouter>
    <div>
    {currentSlide === 1 ? <ServerConfig /> : <Config />}
    <button onClick={nextSlide} style={styles.button}>
      {currentSlide === 1 ? 
      <Button
      color="primary"
      variant="contained"
      m={1}
      className="mt-5"
    >
    Node Configuration
    </Button>
     : 
     <Button
     color="primary"
     variant="contained"
     m={1}
     className="mt-5"
   >
    Blockchain Configuration
    </Button>
    }
      
    </button>
  </div>
  );
}

const styles = {
  button: {
    position: 'fixed',
    bottom: '20px',
    right: '20px',
    padding: '10px 20px',
    fontSize: '1em',
  },
};

export default App;
