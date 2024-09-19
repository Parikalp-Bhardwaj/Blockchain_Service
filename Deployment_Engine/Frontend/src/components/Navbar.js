import React from 'react'
import Validator from './Validator'
import Index from './Index'
const Navbar = () => {
  return (
    <div>
    <nav>
      <ul>
        <li><Link to="/">Home</Link></li>
        <li><Link to="/validator">Validator</Link></li>
        <li><Link to="/tns">Transaction</Link></li>
      </ul>
    </nav>
    </div>
  )
}

export default Navbar