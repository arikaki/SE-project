
import * as React from 'react';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import { CardActionArea } from '@mui/material';
import './Grid.css';
import categories from '../Category';

function Grid() {
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

  const addItems = (e) => {
    if (buttons.find(e.target.value) === e.target.value) {
      buttons.push(e.target.value);
    }
    else {
      return;
    }
  }

  return (
    <div >
      <div className="container">
        {categories.map((item) => {
          // {console.log(item)}
          return (<Card sx={{ maxWidth: 150 }} className='card-icon' onClick={() => handleClick(item)}>
            <CardActionArea>
              <CardMedia
                component="img"
                height="100"
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
      </div>

      <Button className="sample" variant="contained" disabled={selected.length < 5}>
        Submit
      </Button>
    </div>
  )
}

export default Grid;