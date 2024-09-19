import React, { useState } from 'react';
import { Container, TextField, Button, IconButton, Typography, Grid, Paper, Select, MenuItem, InputLabel, FormControl } from '@mui/material';
import { Add, Delete } from '@mui/icons-material';
import { isDisabled } from '@testing-library/user-event/dist/utils';

const ServerConfig = () => {
  const [hosts, setHosts] = useState([{ name: '', ip: '', port: '', type: 'master', ansible_user: '', ansible_ssh_private_key_file: '', password: '' }]);
  const [isDisabledButton, setDisabledButton] = useState(true)


  const handleInputChange = (index, event) => {
    const { name, value } = event.target;
    const updatedHosts = [...hosts];
    updatedHosts[index][name] = value;
    setHosts(updatedHosts);
  };

  const addHost = () => {
    setHosts([...hosts, { name: '', ip: '', port: '', type: 'worker', ansible_user: '', ansible_ssh_private_key_file: '', password: '' }]);
  };

  const removeHost = (index) => {
    setHosts(hosts.filter((_, i) => i !== index));
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    const payload = { inventory: hosts };

    try {
      const response = await fetch('/api/saveInventory', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(payload),
      });

      if (response.ok) {
        setDisabledButton(true)
        alert('Inventory saved successfully.');
      } else {
        alert('Failed to save inventory.');
      }
    } catch (error) {
      console.error('Error:', error);
      alert('An error occurred while saving the inventory.');
    }
  };

  const createCluster = () =>{
    console.log("--------------")
  }

  return (
    <Container style={{backgroundColor: "#597688"}}>
      <Paper elevation={3} style={{ padding: 24, marginTop: 32 ,backgroundColor: '#e5eaec' }}>
      <Typography variant="h6" gutterBottom >
        Create Ansible Hosts Inventory
      </Typography>
      <form onSubmit={handleSubmit}>
        {hosts.map((host, index) => (
          <Paper key={index} style={{ padding: 16, marginBottom: 16 }}>
            <Grid container spacing={2}>
              <Grid item xs={12} sm={2}>
                <TextField
                  label="Host Name"
                  name="name"
                  fullWidth
                  variant="outlined"
                  value={host.name}
                  onChange={(event) => handleInputChange(index, event)}
                  required
                />
              </Grid>
              <Grid item xs={12} sm={2}>
                <TextField
                  label="IP Address"
                  name="ip"
                  fullWidth
                  variant="outlined"
                  value={host.ip}
                  onChange={(event) => handleInputChange(index, event)}
                  required
                />
              </Grid>
              <Grid item xs={12} sm={2}>
                <TextField
                  label="Port"
                  name="port"
                  type="number"
                  fullWidth
                  variant="outlined"
                  value={host.port}
                  onChange={(event) => handleInputChange(index, event)}
                  required
                />
              </Grid>
              <Grid item xs={12} sm={2}>
                <TextField
                  label="Ansible User"
                  name="ansible_user"
                  fullWidth
                  variant="outlined"
                  value={host.ansible_user}
                  onChange={(event) => handleInputChange(index, event)}
                  required
                />
              </Grid>
              <Grid item xs={12} sm={2}>
                <TextField
                  label="SSH Key Path"
                  name="ansible_ssh_private_key_file"
                  fullWidth
                  variant="outlined"
                  value={host.ansible_ssh_private_key_file}
                  onChange={(event) => handleInputChange(index, event)}
                  required
                />
              </Grid>
              <Grid item xs={12} sm={2}>
                <TextField
                  label="Password"
                  name="password"
                  type="password"
                  fullWidth
                  variant="outlined"
                  value={host.password}
                  onChange={(event) => handleInputChange(index, event)}
                  required
                />
              </Grid>
              <Grid item xs={12} sm={2}>
                <FormControl fullWidth variant="outlined" required>
                  <InputLabel id={`type-label-${index}`}>Type</InputLabel>
                  <Select
                    labelId={`type-label-${index}`}
                    name="type"
                    value={host.type}
                    onChange={(event) => handleInputChange(index, event)}
                    label="Type"
                  >
                    <MenuItem value="master">Master</MenuItem>
                    <MenuItem value="worker">Worker</MenuItem>
                  </Select>
                </FormControl>
              </Grid>
              <Grid item xs={12} sm={1} style={{ display: 'flex', alignItems: 'center' }}>
                <IconButton color="error" onClick={() => removeHost(index)}>
                  <Delete />
                </IconButton>
              </Grid>
            </Grid>
          </Paper>
        ))}
        <Button
          variant="contained"
          color="primary"
          startIcon={<Add />}
          onClick={addHost}
          style={{ marginLeft: 8 }}
        >
          Add Host
        </Button>
        <Button variant="contained" color="success" type="submit" style={{ margin: 8 }}>
          Save Inventory
        </Button>
        <Button variant="outlined" color="success" type="submit" disabled={isDisabledButton}
        style={{ margin: 8, backgroundColor: '#d52649', color: "#fff" }} onClick={createCluster}>
          Create Cluster
        </Button>
      </form>
      </Paper>
    </Container>
  );
};

export default ServerConfig;
