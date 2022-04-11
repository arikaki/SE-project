import React from 'react';
import './NewCategories.css';
import Grid from './Grid';
// import CustomButton from "./CustomButton";

function NewCategories() {
  
    console.log('New Category rendered')
  return (
    <div>
    <div className='text'> Select atleast five categories</div>
    <div className="container">
      <Grid />
    </div>
    </div>
  );
}

export default NewCategories;