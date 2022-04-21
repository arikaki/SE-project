import React from 'react';
import './css/NewCategories.css';
import Grid from './Grid';

function NewCategories(props) {
  return (
    <div className={`categories ${props.showFade? " fade-search": ""}`} onClick={() => props.setShowFade(false)}>
      <div className='text1'>Categories</div>
      {!props.notRegister && <div className='text2'>(Atleast 5)</div>}
      <div className="column3">
        <Grid setNewuser={props.setNewuser} selected={props.selected} notRegister={props.notRegister}/>
      </div>
    </div>
  );
}

export default NewCategories;