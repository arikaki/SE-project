import { Avatar } from "@material-ui/core";
import React,{useState} from "react";
import { useSelector } from "react-redux";
import { selectUser } from "../feature/userSlice";
import "./css/QuoraBox.css";
import AddIcon from "@material-ui/icons/Add";
import Fab from "@material-ui/core/Fab";
import Zoom from "@material-ui/core/Zoom";


function Answerbox(props) {
  const [isExpanded, setExpanded] = useState(false);

  const [note, setNote] = useState({
    // title: "",
    content: ""
  });

  function handleChange(event) {
    const { name, value } = event.target;

    setNote(prevNote => {
      return {
        ...prevNote,
        [name]: value
      };
    });
  }

  function submitNote(event) {
    props.onAdd(note);
    setNote({
      // title: "",
      content: ""
    });
    event.preventDefault();
  }

  function expand() {
     setExpanded(true);
   }




  const user = useSelector(selectUser);
  return (
  //   <div className="quoraBox">

   // <div className="quoraBox__info">
  // <Avatar src={user?.photo} />
  // </div>
  //     <div className="quoraBox__quora">
  //       <h5>What is your question or link?</h5>
  //     </div>
  //   </div>
  // );

<div>
        <form className="create-ans">
          {isExpanded}



          <textarea
          name="content"
          onClick={expand}
          onChange={handleChange}
          value={note.content}
          placeholder="Post your answer..."
          rows={isExpanded ? 3 : 1}
        />
        <Zoom in={isExpanded}>
          <Fab onClick={submitNote}>
            <AddIcon />
          </Fab>
        </Zoom>
      </form>
    </div>
  );

}

export default Answerbox;
