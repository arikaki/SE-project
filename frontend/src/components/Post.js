import { Avatar } from "@material-ui/core";
import {
  ArrowDownwardOutlined,
  ArrowUpwardOutlined,
  ChatBubbleOutlined,
  MoreHorizOutlined,
  RepeatOneOutlined,
  ShareOutlined,
} from "@material-ui/icons";
import React, { useState } from "react";
import "./css/Post.css";
import { Modal } from "react-responsive-modal";
import "react-responsive-modal/styles.css";
import CloseIcon from "@material-ui/icons/Close";
import ReactQuill from "react-quill";
import "react-quill/dist/quill.snow.css";
import ReactTimeAgo from "react-time-ago";
import axios from "axios";
import ReactHtmlParser from "html-react-parser";
import { useSelector } from "react-redux";
import { selectUser } from "../feature/userSlice";

function LastSeen({ date }) {
  return (
    <div>
      <ReactTimeAgo date={date} locale="en-US" timeStyle="round" />
    </div>
  );
}
function Post({ post }) {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [answer, setAnswer] = useState("");
  const Close = <CloseIcon />;

  const user = useSelector(selectUser);

  const handleQuill = (value) => {
    setAnswer(value);
  };

  const handleSubmit = async () => {
    if (post?._id && answer !== "") {
      const config = {
        headers: {
          "Content-Type": "application/json",
        },
      };
      const body = {
        answer: answer,
        questionId: post?._id,
        user: user,
      };
      await axios
        .post("/api/answers", body, config)
        .then((res) => {
          alert("Answer added succesfully");
          setIsModalOpen(false);
          window.location.href = "/";
        })
        .catch((e) => {
          console.log(e);
        });
    }
  };
  return (
    <div className="post">
      <div className="post__info">
        <h4>{post?.Username}</h4>
      </div>
      <div className="post__body">
        <div className="post__question">
          <p>{post?.Question}</p>
          <button
            onClick={() => {
              setIsModalOpen(true);
              console.log(post?._id);
            }}
            className="post__btnAnswer"
          >
            Answer
          </button>
          <Modal
            open={isModalOpen}
            closeIcon={Close}
            onClose={() => setIsModalOpen(false)}
            closeOnEsc
            center
            closeOnOverlayClick={false}
            styles={{
              overlay: {
                height: "auto",
              },
            }}
          >
            <div className="modal__question">
              <h1>{post?.questionName}</h1>
              <p>
                asked by <span className="name">{post?.user?.userName}</span> on{" "}
                <span className="name">
                  {new Date(post?.createdAt).toLocaleString()}
                </span>
              </p>
            </div>
            <div className="modal__answer">
              <ReactQuill
                value={answer}
                onChange={handleQuill}
                placeholder="Enter your answer"
              />
            </div>
            <div className="modal__button">
              <button className="cancel" onClick={() => setIsModalOpen(false)}>
                Cancel
              </button>
              <button onClick={handleSubmit} type="submit" className="add">
                Add Answer
              </button>
            </div>
          </Modal>
        </div>
      </div>
      <div className="post__footer">
        <div className="post__footerAction">
          <ArrowUpwardOutlined />
          <ArrowDownwardOutlined />
        </div>
        <div className="post__footerLeft">
          <button>Report</button>
        </div>
      </div>
    </div>
  );
}

export default Post;