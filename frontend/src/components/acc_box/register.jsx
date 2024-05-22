import Lottie from 'lottie-react';
import animationData from "../../assets/hub-anim.json"
import React, { useState } from 'react';
import {Navigate, Link} from 'react-router-dom';

function Register() {
    const [passwordError, setPasswordError] = useState('');
    const [navigateToLogin, setNavigateToLogin] = useState(false);
  
    const checkpw = (e) => {
      e.preventDefault();
      const pw = document.getElementById("pw").value;
      const cpw = document.getElementById("cpw").value;
  
      if (pw !== cpw) {
        setPasswordError('Password does not match');
        setNavigateToLogin(false);
      } else {
        setPasswordError('');
        setNavigateToLogin(true);
      }
    }
  
    if (navigateToLogin) {
      return <Navigate to="/home" replace />;
    }
  
    return (
      <>
        <Lottie animationData={animationData} className='logo' />
        <div className='Login'>
          <h2>Register</h2>
          <form onSubmit={checkpw}>
            <label htmlFor='nickname'>Nickname:</label>
            <input type="text" name='nickname' />
  
            <label htmlFor='password'>Password:</label>
            <input type="password" name='password' id="pw" />
  
            <label htmlFor='confPassword'>Confirm your password:</label>
            <input type="password" name='confPassword' id="cpw" />
  
            <label htmlFor='confPassword'>Address:</label>
            <div className='ip'>
              <input className='Address' type="text" name='Address' />
              <input className='Port' type="text" name='Port' defaultValue='8080'  />
            </div>
  
  
            <p>{passwordError}</p>
            <p>Already have a SupaHub account? <Link to='/login'>sign in</Link></p>
  
            <input type="submit" value="Register" />
          </form>
        </div>
        
      </>
    );
  };

export default Register;