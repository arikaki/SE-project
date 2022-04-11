import React, { useEffect, useState } from "react";
import Answerbox from "./Answerbox";
import QuoraHeader from "./QuoraHeader";
import Addans from "./Addans";
import "./css/Feed.css";



function Answer() {

  const [notes, setNotes] = useState([]);

  function addNote(newNote) {
    setNotes(prevNotes => {
      return [...prevNotes, newNote];
    });
  }

  function deleteNote(id) {
    setNotes(prevNotes => {
      return prevNotes.filter((noteItem, index) => {
        return index !== id;
      });
    });
  }


  return (
    <div className="feed">
  
      <Answerbox onAdd={addNote} />

<div>
      {notes.map((noteItem, index) => {
        return (
          <Addans
            key={index}
            id={index}

            content={noteItem.content}
            onDelete={deleteNote}
          />
        );
      })}
      </div>


    </div>
  );
}

export default Answer;
