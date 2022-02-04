import React from "react";
import Button from '@mui/material/Button';
import { getThemeProps } from "@material-ui/styles";
import {black } from "@material-ui/core/colors";


const CustomButton = (props) => {
return(
    <div>
        <Button >
            {props.children}
        </Button>
    </div>)
}

export default CustomButton;