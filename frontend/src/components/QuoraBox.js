import React, { useState } from "react";
import { useSelector } from "react-redux";
import { selectUser } from "../feature/userSlice";
import AddIcon from "@material-ui/icons/Add";
import Fab from "@material-ui/core/Fab";
import Zoom from "@material-ui/core/Zoom";
import axios from "axios";

function QuoraBox(props) {
  const [isExpanded, setExpanded] = useState(false);

  const [note, setNote] = useState({
    title: "",
    content: "",
  });

  function handleChange(event) {
    const { name, value } = event.target;

    setNote((prevNote) => {
      return {
        ...prevNote,
        [name]: value,
      };
    });
  }

  function submitNote(event) {
    axios.post(
      "http://localhost:8000/api/question/ask",
      {
        Topic: note.title,
        Question: note.content,
      },
      {
        "Access-Control-Allow-Origin": "*",
        "withCredentials": true,
      }
      )
      .then(function (response) {
        if (response.status == 200) {
          window.alert("Question Posted!")
          setExpanded(false);
        }
      })
      .catch(function (error) {
        console.log(error);
      });
    setNote({
      title: "",
      content: "",
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
      <form className="create-note">
        {isExpanded && (
          // <input
          //   name="title"
          //   onChange={handleChange}
          //   value={note.title}
          //   placeholder="Title"
          // >
          //   <option defaultValue="open">open</option>
          //   <option values="in progress">progress</option>
          //   <option values="completed">done</option>
          // </input>
          <select
            name="title"
            title="Title"
            onChange={handleChange}
            value={note.title}
            style={{
              padding: "10px",
              backgroundColor: "#3c6acccc",
              border: "none",
              color: "white",
              borderRadius: "10px",
              margin: "5px",
            }}
          >
            <option selected>Choose...</option>
            <option value="Science">Science</option>
            <option value="Music">Music</option>
            <option value="Technology">Technology</option>
            <option value="Computer">Computer</option>
            <option value="History">History</option>
            <option value="Movies">Movies</option>
            <option value="Cooking">Cooking</option>
            <option value="Health">Health</option>
            <option value="Psychology">Psychology</option>
            <option value="Education">Education</option>
            <option value="Business">Business</option>
            <option value="Finance">Finance</option>
          </select>
          // <Dropdown
          //   name="title"
          //   title="Title"
          //   onChange={handleChange}
          //   value={note.title}
          // >
          //   <Dropdown.Toggle variant="success" id="dropdown-basic">
          //     Dropdown Button
          //   </Dropdown.Toggle>

          //   <Dropdown.Menu>
          //     <Dropdown.Item href="#/action-1">Action</Dropdown.Item>
          //     <Dropdown.Item href="#/action-2">Another action</Dropdown.Item>
          //     <Dropdown.Item href="#/action-3">Something else</Dropdown.Item>
          //   </Dropdown.Menu>
          // </Dropdown>
        )}

        <textarea
          name="content"
          onClick={expand}
          onChange={handleChange}
          value={note.content}
          placeholder="Post a Question..."
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

export default QuoraBox;
