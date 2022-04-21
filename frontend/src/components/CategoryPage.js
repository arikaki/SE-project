import React, { useEffect } from "react";
import userSlice from "../feature/userSlice";
import NewCategories from "./NewCategories";
import axios from "axios";

function CategoryPage(props) {
    const [selected, setSelected] = React.useState([]);
    const getSelected = () => {
        const {user} = props;
        user && axios.post('http://localhost:8000/api/user/fetch-user', {
            UserName: user.userName,
        }, {
            "withCredentials": true,
            "Access-control-Allow-Origin": "http://localhost:8000"
        })
            .then((response) => {
                setSelected(response.data.Topics);
            })
            .catch((error) => {
                console.log(error);
            });
    }
    useEffect(() => {
        getSelected();
    }, [props.user]);

    return <NewCategories selected={selected} notRegister={props.notRegister} showFade={props.showFade} setShowFade={props.setShowFade? props.setShowFade: ()=>{}}/>
}

export default CategoryPage;