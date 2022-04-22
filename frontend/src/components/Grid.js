
import * as React from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Typography from '@mui/material/Typography';
import { CardActionArea } from '@mui/material';
import './css/Grid.css';
import axios from 'axios';
import categories from './Category';
import { useNavigate } from "react-router-dom";

function Grid(props) {
  let navigate = useNavigate();
  const [selected, setSelected] = React.useState([]);

  React.useEffect(() => {
    props.selected && setSelected(props.selected);
  }, [props.selected])

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

    axios.post('http://localhost:8000/api/user/setUserCategory', {
      Topic: selected
    }, {
      "withCredentials": true,
      "Access-control-Allow-Origin": "http://localhost:8000"
    })
      .then((response) => {
        if (props.notRegister) {
          navigate("/");
        } else {
          props.setNewuser(false);
          localStorage.setItem('NewUser', false);
        }
      })
      .catch((error) => {
        console.log(error);
      });
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