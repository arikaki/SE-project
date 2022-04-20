import React from 'react';
import './css/NewCategories.css';
import Grid from './Grid';

function NewCategories(props) {
  return (
    <div className='categories'>
      <div className='text1'>Categories</div>
      <div className='text2'>(Atleast 5)</div>
      <div className="column3">
        <Grid setNewuser={props.setNewuser}/>
      </div>
    </div>
  );
}

export default NewCategories;