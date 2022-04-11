import React ,{useState} from "react";
import DeleteIcon from "@material-ui/icons/Delete";
//import ArrowCircleDownIcon from '@mui/icons-material/ArrowCircleDown';
// import ArrowCircleUpIcon from '@mui/icons-material/ArrowCircleUp';
import "./css/addans.css";
function Addans(props) {

  const [count, setCount] = useState(0);

   function increase() {
     setCount(count + 1);
   }

   function decrease() {
     setCount(count - 1);
   }
  function handleClick() {
    props.onDelete(props.id);
  }

  return (
    <div className="note">

      <p>{props.content}</p>

      <button className="voteup" onClick={increase} title="upvote">
      +</button>
      <p className="count">{count}</p>
      <button className="votedown" onClick={decrease} title="downvote">
      -</button>

      <button className="report" title="report">
      !</button>

      <button className="delete" onClick={handleClick} title="delete">
        <DeleteIcon />
      </button>
    </div>
  );
}

export default Addans;
