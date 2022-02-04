import React, { useState } from "react";
import { Avatar } from "@material-ui/core";
import QuestionCard from "./Components/QuestionCard";
// import axios from "axios";
const QuestionBox = ({ profile, auth_status }) => {
    const [question, setQuestion] = useState("");

    const AskQuestion = async () => {
        const form_data = new FormData();
        form_data.append("question", question);

        // const url = "http://localhost:5000/api/ask-question";

        // try {
        //     const response = await axios.post(url, form_data, {
        //         withCredentials: true,
        //     });

        //     alert(response.data.msg);
        // } catch (error) {
        //     alert(error.response.data.msg);
        // }
    };

    return (
        <div>
            <QuestionCard title="Which college should I go to and why: FSU or UF? nbdg" firstCTA="Answer" secondCTA="follow" type="question" />
            <QuestionCard title="uf is awesome!!!" firstCTA="Comment" secondCTA="share" type="answer" />
        </div>
    );
};

export default QuestionBox;