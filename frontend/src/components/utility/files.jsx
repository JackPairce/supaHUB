import React, { useState, useEffect } from 'react';
import MultiFiles from './multiFiles';
import Sidebar from '../main_comp/sidebar';

const Files = () => {
  return (
    <>
      <MultiFiles />
      <Sidebar />
    </>
  )
};

export default Files;