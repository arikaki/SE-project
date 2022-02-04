import React from 'react';
import { Grid, Typography, Paper, Toolbar, Button, Menu, MenuItem, Fade, IconButton } from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';
import MenuIcon from '@material-ui/icons/Menu';
import { padding } from '@mui/system';

const useStyles = makeStyles((theme) => ({

    root: {
        margin: '7px 7px 7px 7px',
        borderRadius: '35px',
        display: 'flex',
        background: '#1DC0F4'
    },
    content: {
        flex: '1 1 auto',
        padding: '0px 15px 15px',
        width: '100%',
        height: '100vh',
        overflow: 'hidden',
    },
    title: {
        marginLeft: '40%',
        marginTop: '50%',
        borderRadius: '15px',
        background: '#148DB5',
        padding: '20px',
    },
    logout: {
        // borderRadius: '15px',
        // marginRight: '40%',
        marginLeft: '80%',
        marginTop: '2%',
        // justifyContent: 'flex-end'
        // background: '#ad0014'
    }
}));