
import * as React from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Typography from '@mui/material/Typography';
import { CardActionArea } from '@mui/material';
import './css/Grid.css';
import categories from './Category';

function Grid(props) {
  var buttons = [];
  const [selected, setSelected] = React.useState([]);

  const handleClick = (item) => {
    let selectedCopy;
    const index = selected.indexOf(item)
    if (index >= 0) {
      selectedCopy = [...selected]
      selectedCopy.splice(index, 1)
    } else {
      selectedCopy = [...selected]
      selectedCopy.push(item)
    }
    setSelected(selectedCopy)
  };

  const onCategorySubmit = () => {
    props.setNewuser(false);
    localStorage.setItem('NewUser', false);
  }

  return (
    <div className="container1">
      {categories.map((item) => {
        return (<Card className='card-icon' onClick={() => handleClick(item)}>
          <CardActionArea>
            <CardMedia
              component="img"
              height="130"
              image={`/images/${item}.jpeg`}
              // alt="green iguana"
            />
            <CardContent className={`${selected.indexOf(item) >= 0 ? ' selected' : ''}`}>
              <Typography gutterBottom variant="h6" component="div">
                {item}
              </Typography>
            </CardContent>
          </CardActionArea>
        </Card>)
      })}
      <button className="card-icon sample" variant="contained" disabled={selected.length < 5} onClick={onCategorySubmit}>
        Continue
      </button>
    </div>
  )
}

export default Grid;