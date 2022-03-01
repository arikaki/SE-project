import React, { useState } from "react";
import Avatar from "@material-ui/core/Avatar";
import "../StyleSheet/QuestionBox.css"; 
import { Grid } from '@mui/material';
import { Fab } from "@material-ui/core";
// import axios from "axios";
const QuestionBox = (props) => {
  const [question, setQuestion] = useState("");

  const AskQuestion = async () => {
    const form_data = new FormData();
    form_data.append("question", question);

    const url = "http://localhost:5000/api/ask-question";

    // try {
    //   const response = await axios.post(url, form_data, {
    //     withCredentials: true,
    //   });

    //   alert(response.data.msg);
    // } catch (error) {
    //   alert(error.response.data.msg);
    // }
  };

  return (
    <div className="QuestionBox">
      {/* <Stack sx={{ width: '100%' }} spacing={2}>
      <div className="QuestionBox__user">
        <Avatar alt="User Profile" />
        <h4 className="user__username">Junior</h4>
      </div>
      <Alert onClose={() => {}}></Alert>
       </Stack> */}
      <Grid container spacing={1}>
        <Grid item xs={11}>
        <div className="QuestionBox__user">
        <Avatar alt="User Profile" />
        <h4 className="user__username">Junior</h4>
      </div>
      </Grid>
      <Grid item xs={1}>
      {/* <CloseButton /> */}
      <Fab size='small' onClick={props.closeQuestion}>x</Fab>
      </Grid>
      </Grid>
      <div className="QuestionBox__inputField">
        <input
          type="text"
          placeholder="What is your question"
          className="QuestionBox__inputField"
          onChange={(e) => setQuestion(e.target.value)}
          value={question}
        />
        <button
          disabled={false}
          className="QuestionBox__btn"
          onClick={AskQuestion}
        >
          Ask Question
        </button>
      </div>
    </div>
  );
};

export default QuestionBox;