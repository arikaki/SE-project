import React, { useEffect, useState } from "react";
import QuoraBox from "./QuoraBox";
import "./css/Feed.css";
import Post from "./Post";
import axios from "axios";

import AddQue from "./AddQue"

function Feed() {
  const [posts, setPosts] = useState([]);
  const [notes, setNotes] = useState([]);
  useEffect(() => {
    axios.get("http://localhost:8000/api/question/getAll", {
          "withCredentials": true,
          "Access-control-Allow-Origin": "http://localhost:8000"
      })
      .then((res) => {
        setPosts(res.data);
      })
      .catch((e) => {
        console.log(e);
      });
  }, []);




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

      <QuoraBox onAdd={addNote} />

<div>
      {notes.map((noteItem, index) => {
        return (
          <AddQue
            key={index}
            id={index}
            title={noteItem.title}
            content={noteItem.content}
            onDelete={deleteNote}
          />
        );
      })}
      </div>
      {posts.map((post, index) => (
        <Post key={index} post={post} />
      ))}
      {/* <Post />
      <Post />
      <Post />
      <Post />
      <Post /> */}
    </div>
  );
}

export default Feed;
