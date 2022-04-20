import React ,{useState} from "react";
import DeleteIcon from "@material-ui/icons/Delete";
//import ArrowCircleDownIcon from '@mui/icons-material/ArrowCircleDown';
// import ArrowCircleUpIcon from '@mui/icons-material/ArrowCircleUp';
import "./css/addans.css";
function Addans(props) {

  // const [count, setCount] = useState(0);
  //
  //  function increase() {
  //    setCount(count + 1);
  //  }
  //
  //  function decrease() {
  //    setCount(count - 1);
  //  }

  const [voteCount, setVoteCount] = useState(0);
      const [isUpvoted, setIsUpvoted] = useState(false);
      const [isDownvoted, setIsDownvoted] = useState(false);

      const handleUpvote= () => {
          if(isDownvoted){
              setIsDownvoted(false);
              setIsUpvoted(true);
              setVoteCount(prevCount => prevCount+2)
          }
          else if(isUpvoted){
              setIsUpvoted(false);
              setIsDownvoted(false);
              setVoteCount(prevCount => prevCount-1)
          }
          else{
              setIsUpvoted(true);
              setVoteCount(prevCount => prevCount+1)
          }
      }

      const handleDownvote = () => {
          if(isDownvoted){
              setIsDownvoted(false);
              setIsUpvoted(false);
              setVoteCount(prevCount => prevCount+1)
          }
          else if(isUpvoted){
              setIsUpvoted(false);
              setIsDownvoted(true);
              setVoteCount(prevCount => prevCount-2)
          }
          else{
              setIsDownvoted(true);
              setVoteCount(prevCount => prevCount-1)
          }
      }




  function handleClick() {
    props.onDelete(props.id);
  }

  return (
    <div className="note">

      <p>{props.content}</p>

      <button className="voteup" onClick={handleUpvote} title="upvote">
      +</button>
      <p className="count">{voteCount}</p>
      <button className="votedown" onClick={handleDownvote} title="downvote">
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